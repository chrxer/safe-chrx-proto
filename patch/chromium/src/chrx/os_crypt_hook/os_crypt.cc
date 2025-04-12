#include "third_party/crashpad/crashpad/third_party/cpp-httplib/cpp-httplib/httplib.h"
#include "../chrx_server/chrx_server.h"
#include "../aes_crypt/aes_crypt.h"
#include "os_crypt.h"
#include <vector>
#include <string>
#include <cstdint>
#include <chrono>
#include <thread>
#include "base/logging.h"

using ByteVector = std::vector<uint8_t>;

bool ChrxSendRequest(const std::string& endpoint, const std::string& input, std::string& output) {
  // Get server port and key from the launcher
  auto& launcher = CryptServerLauncher::Instance();
  int port = launcher.GetPort();
  const std::string& key = launcher.GetKey();

  std::string url = "http://localhost:" + std::to_string(port);
  LOG(INFO) << "sending request to: " << url << endpoint;

  httplib::Client client(url);
  client.set_connection_timeout(1);
  client.set_read_timeout(1);
  client.set_write_timeout(1);

  // Convert strings to byte vectors
  ByteVector input_vec(input.begin(), input.end());
  ByteVector key_vec(key.begin(), key.end());
  ByteVector encrypted_vec;

  // Encrypt the input data before sending
  if (!EncryptAESGCM(input_vec, key_vec, encrypted_vec)) {
      LOG(ERROR) << "EncryptAESGCM failed";
      return false;
  }

  // Convert encrypted vector back to string for sending
  std::string encrypted_input(encrypted_vec.begin(), encrypted_vec.end());

  // Retry mechanism
  const int max_duration_seconds = 20;
  const int retry_interval_ms = 500;
  auto start_time = std::chrono::steady_clock::now();

  while (std::chrono::steady_clock::now() - start_time < std::chrono::seconds(max_duration_seconds)) {
      auto response = client.Post(endpoint.c_str(), encrypted_input, "application/octet-stream");

      if (response) {
          if (response->status == 200) {
              // Convert response to vector for decryption
              ByteVector response_vec(response->body.begin(), response->body.end());
              ByteVector output_vec;

              // Decrypt the response
              if (!DecryptAESGCM(response_vec, key_vec, output_vec)) {
                  LOG(ERROR) << "DecryptAESGCM failed";
                  return false;
              }

              // Convert decrypted vector back to string
              output.assign(output_vec.begin(), output_vec.end());
              return true;
          } else {
              LOG(ERROR) << "Received HTTP status: " << response->status;
              return false;  // Exit early if server responded but with an error
          }
      }

      // Sleep a bit before retrying
      std::this_thread::sleep_for(std::chrono::milliseconds(retry_interval_ms));
  }

  LOG(ERROR) << "No response received from server after 20 seconds\n. Server process exit code (-1 if running): " << launcher.GetExitCode();
  return false;
}

bool ChrxEncrypt(const std::string& plaintext, std::string& ciphertext) {
    bool ok = ChrxSendRequest("/encrypt", plaintext, ciphertext);
    if(!ok){
      LOG(FATAL) << "ChrxSendRequest to /encrypt failed";
    }
    return ok;
}

bool ChrxDecrypt(std::string& plaintext, const std::string& ciphertext) {
    bool ok = ChrxSendRequest("/decrypt", ciphertext, plaintext);
    if(!ok){
      LOG(FATAL) << "ChrxSendRequest to /decrypt failed";
    }
    return ok;
}
