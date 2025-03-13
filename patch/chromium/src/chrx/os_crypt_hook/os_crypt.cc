#include "third_party/crashpad/crashpad/third_party/cpp-httplib/cpp-httplib/httplib.h"
// TODO(@kaliiiiiiiiii) use chrx/third_party/HTTPRequest or /net instead

#include <string>

const std::string url = "http://localhost:3333";

bool ChrxSendRequest(const std::string& endpoint, const std::string& input, std::string& output) {
    httplib::Client client(url);
    auto response = client.Post(endpoint.c_str(), input, "application/json");

    if (response && response->status == 200) {
        output = response->body;
        return true;
    }
    return false;
}

bool ChrxEncrypt(const std::string& plaintext, std::string& ciphertext) {
    return ChrxSendRequest("/encrypt", plaintext, ciphertext);
}

bool ChrxDecrypt(std::string& plaintext, const std::string& ciphertext) {
    return ChrxSendRequest("/decrypt", ciphertext, plaintext);
}
