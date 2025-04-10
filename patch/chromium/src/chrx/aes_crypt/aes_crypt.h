#ifndef AES_CRYPT_H_
#define AES_CRYPT_H_

#include <string>
#include <vector>
#include <cstdint>

// Constants
constexpr size_t kKeySize = 32;  // AES-256 key size
constexpr size_t kIvSize = 12;   // GCM standard nonce size
constexpr size_t kTagSize = 16;  // GCM authentication tag size

// Generates SHA-256 hash of the input string.
std::vector<uint8_t> NewSHA256(const std::string& input);

// Encrypts plaintext using AES-256-GCM.
// Returns true on success; false on failure.
bool EncryptAESGCM(const std::vector<uint8_t>& plaintext,
                   const std::vector<uint8_t>& key,
                   std::vector<uint8_t>& ciphertext);

// Decrypts ciphertext using AES-256-GCM.
// Returns true on success; false on failure.
bool DecryptAESGCM(const std::vector<uint8_t>& ciphertext,
                   const std::vector<uint8_t>& key,
                   std::vector<uint8_t>& plaintext);

#endif  // AES_CRYPT_H_
