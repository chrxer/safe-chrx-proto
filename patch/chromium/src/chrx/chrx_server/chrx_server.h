#pragma once

#include <string>
#include "base/files/file_path.h"
#include "base/process/launch.h"
#include "base/process/process.h"
#include "base/rand_util.h"

base::FilePath GetExecutablePath();

class CryptServerLauncher {
public:
    static CryptServerLauncher& Instance();

    int GetPort();
    const std::string& GetKey();

    // Process's exit code (-1 if still running)
    int GetExitCode();

private:
    CryptServerLauncher();
    bool Start();
    bool LaunchChild();

    bool started_ = false;
    int port_;
    std::string aes_key_;
    base::FilePath executable_path_;
    base::Process server_process_;
};
