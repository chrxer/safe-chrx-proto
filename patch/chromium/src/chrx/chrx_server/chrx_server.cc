#include "base/base_paths.h"
#include "base/files/file_util.h"
#include "base/path_service.h"
#include "base/strings/string_number_conversions.h"
#include "base/base64.h"
#include "base/logging.h"
#include "ports.h"
#include "chrx_server.h"

#if defined(OS_WIN)
#include <windows.h>
constexpr char kServerBinary[] = "chrxCryptServer.exe";
#else
#include <fcntl.h>
#include <unistd.h>
#include <errno.h>
constexpr char kServerBinary[] = "chrxCryptServer";
#endif

CryptServerLauncher& CryptServerLauncher::Instance() {
  static CryptServerLauncher instance;
  return instance;
}

CryptServerLauncher::CryptServerLauncher()
  : port_(findFreePort()),
    aes_key_(base::RandBytesAsString(32)) {
  base::FilePath exe_path;
  if (base::PathService::Get(base::FILE_EXE, &exe_path)) {
    executable_path_ = exe_path.DirName().AppendASCII(kServerBinary);
  }
}

int CryptServerLauncher::GetPort() {
  if (started_) {
    return port_;
  }
  if (!Start()) {
    LOG(FATAL) << "Failed to start crypt server";
  }
  return port_;
}

const std::string& CryptServerLauncher::GetKey() {
  if (started_) {
    return aes_key_;
  }
  if (!Start()) {
    LOG(FATAL) << "Failed to start crypt server";
  }
  return aes_key_;
}

int CryptServerLauncher::GetExitCode() {
  if (!started_ || !server_process_.IsValid()) {
    return -1;
  }

  int exit_code = -1;
  if (server_process_.WaitForExitWithTimeout(base::Seconds(0), &exit_code)) {
    return exit_code;
  }

  return -1;  // Process still running
}


bool CryptServerLauncher::Start() {
  if (started_ && (!server_process_.IsValid() || server_process_.WaitForExitWithTimeout(base::Seconds(0), nullptr))) return true;
  return LaunchChild();
}

bool CryptServerLauncher::LaunchChild() {
  if (!base::PathExists(executable_path_)) {
    LOG(ERROR) << "Server binary not found";
    return false;
  }

  base::CommandLine cmd(executable_path_);
  cmd.AppendSwitchASCII("port", base::NumberToString(port_));

  base::LaunchOptions options;
  options.wait = false;
  options.kill_on_parent_death = true;
  options.new_process_group = true;

#if defined(OS_WIN) // not tested yet
  HANDLE stdin_read, stdin_write;
  HANDLE stdout_read, stdout_write;
  SECURITY_ATTRIBUTES sa = {sizeof(SECURITY_ATTRIBUTES), NULL, TRUE};

  if (!CreatePipe(&stdin_read, &stdin_write, &sa, 0) ||
      !CreatePipe(&stdout_read, &stdout_write, &sa, 0)) {
    LOG(ERROR) << "Failed to create pipes";
    return false;
  }
  base::ScopedFD parent_write(stdin_write);
  base::ScopedFD parent_read(stdout_read);
  options.stdin_handle = stdin_read;
  options.stdout_handle = stdout_write;
#else
  int stdin_pipe[2], stdout_pipe[2];
  if (pipe(stdin_pipe) != 0 || pipe(stdout_pipe) != 0) {
    LOG(ERROR) << "Failed to create pipes";
    return false;
  }
  base::ScopedFD parent_write(stdin_pipe[1]);
  base::ScopedFD parent_read(stdout_pipe[0]);
  options.fds_to_remap.emplace_back(stdin_pipe[0], STDIN_FILENO);
  options.fds_to_remap.emplace_back(stdout_pipe[1], STDOUT_FILENO);
  fcntl(parent_read.get(), F_SETFL, O_NONBLOCK);
  fcntl(parent_write.get(), F_SETFL, O_NONBLOCK);
#endif

  LOG(INFO) << "Starting process: " << cmd.GetCommandLineString();
  server_process_ = base::LaunchProcess(cmd, options);
  if (!server_process_.IsValid()) {
    LOG(ERROR) << "Failed to launch server process";
    return false;
  }

  std::string key_b64 = base::Base64Encode(aes_key_) + "\n";
  ssize_t written = 0;
  constexpr int kMaxAttempts = 3;
  int attempts = 0;

  while (written < static_cast<ssize_t>(key_b64.size()) && attempts < kMaxAttempts) {
    ssize_t result = write(parent_write.get(),
                          key_b64.data() + written,
                          key_b64.size() - written);

    if (result > 0) {
      written += result;
      attempts = 0;
    } else if (errno == EPIPE) {
      LOG(ERROR) << "Pipe broken";
      return false;
    } else if (errno == EAGAIN || errno == EWOULDBLOCK) {
      base::PlatformThread::Sleep(base::Milliseconds(100));
      attempts++;
    } else {
      LOG(ERROR) << "Write error: " << strerror(errno);
      return false;
    }
  }

  if (written != static_cast<ssize_t>(key_b64.size())) {
    LOG(ERROR) << "Failed to send complete key";
    return false;
  }

  int exit_code;
  if (server_process_.WaitForExitWithTimeout(base::Seconds(0), &exit_code)) {
    LOG(ERROR) << "Server died immediately: " << exit_code;
    return false;
  }

  started_ = true;
  return true;
}
