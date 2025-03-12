#include "third_party/crashpad/crashpad/third_party/cpp-httplib/cpp-httplib/httplib.h"

#include <string>

bool ChrxEncrypt(const std::string& plaintext, std::string& ciphertext) {
  httplib::Client client("http://localhost:3333");
  auto response = client.Post("/encrypt", plaintext, "application/json");

  if (response && response->status == 200) {
      ciphertext = response->body;
      return true;
  }
  return false;
}
