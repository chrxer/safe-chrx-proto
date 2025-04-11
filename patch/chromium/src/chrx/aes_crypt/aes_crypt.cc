#ifndef AES_CRYPT_H_
#define AES_CRYPT_H_

#include "aes_crypt.h"
#include "base/logging.h"
#include <string>
#include <vector>
#include <cstdint>
#include <openssl/aes.h>
#include <openssl/evp.h>
#include <openssl/rand.h>

// generated using AI :)

// Constants
constexpr size_t kKeySize = 32;  // AES-256 key size
constexpr size_t kIvSize = 12;   // GCM standard nonce size
constexpr size_t kTagSize = 16;  // GCM authentication tag size

// Generates SHA-256 hash of the input string using BoringSSL's EVP functions.
std::vector<uint8_t> NewSHA256(const std::string& input) {
    EVP_MD_CTX* md_ctx = EVP_MD_CTX_new();
    const EVP_MD* md = EVP_sha256();
    std::vector<uint8_t> hash(EVP_MD_size(md));

    if (EVP_DigestInit_ex(md_ctx, md, nullptr) != 1 ||
        EVP_DigestUpdate(md_ctx, input.data(), input.size()) != 1 ||
        EVP_DigestFinal_ex(md_ctx, hash.data(), nullptr) != 1) {
        EVP_MD_CTX_free(md_ctx);
        LOG(ERROR) << "computing sha265 failed";
        return {};
    }

    EVP_MD_CTX_free(md_ctx);
    return hash;
}

std::string randKey() {
    unsigned char key[kKeySize];  // AES-256 requires 32 bytes (256 bits)
    if (!RAND_bytes(key, sizeof(key))) {
        LOG(ERROR) << "generating random key failed";
        return std::string("");
    }
    return std::string(reinterpret_cast<char*>(key), sizeof(key));
}

// Encrypts plaintext using AES-256-GCM with BoringSSL.
bool EncryptAESGCM(const std::vector<uint8_t>& plaintext,
                   const std::vector<uint8_t>& key,
                   std::vector<uint8_t>& ciphertext) {
    if (key.size() != kKeySize) {
        LOG(ERROR) << "Invalid key size";
        return false;
    }

    std::vector<uint8_t> iv(kIvSize);
    if (RAND_bytes(iv.data(), kIvSize) != 1) {
        LOG(ERROR) << "Failed to generate IV";
        return false;
    }

    EVP_CIPHER_CTX* ctx = EVP_CIPHER_CTX_new();
    if (!ctx) {
        LOG(ERROR) << "Failed to create cypher ctx";
        return false;
    }

    int len;
    if (EVP_EncryptInit_ex(ctx, EVP_aes_256_gcm(), nullptr, key.data(), iv.data()) != 1) {
        EVP_CIPHER_CTX_free(ctx);
        LOG(ERROR) << "Failed to EVP_EncryptInit_ex";
    }

    std::vector<uint8_t> out(plaintext.size() + kTagSize);
    if (EVP_EncryptUpdate(ctx, out.data(), &len, plaintext.data(), plaintext.size()) != 1) {
        EVP_CIPHER_CTX_free(ctx);
        LOG(ERROR) << "Failed to VP_EncryptUpdate";
        return false;
    }

    int ciphertext_len = len;

    if (EVP_EncryptFinal_ex(ctx, out.data() + len, &len) != 1) {
        EVP_CIPHER_CTX_free(ctx);
        LOG(ERROR) << "Failed to EVP_EncryptFinal_ex";
        return false;
    }
    ciphertext_len += len;

    std::vector<uint8_t> tag(kTagSize);
    if (EVP_CIPHER_CTX_ctrl(ctx, EVP_CTRL_GCM_GET_TAG, kTagSize, tag.data()) != 1) {
        EVP_CIPHER_CTX_free(ctx);
        LOG(ERROR) << "Failed to EVP_CIPHER_CTX_ctrl";
        return false;
    }

    ciphertext.resize(kIvSize + ciphertext_len + kTagSize);
    std::copy(iv.begin(), iv.end(), ciphertext.begin());
    std::copy(out.begin(), out.begin() + ciphertext_len, ciphertext.begin() + kIvSize);
    std::copy(tag.begin(), tag.end(), ciphertext.begin() + kIvSize + ciphertext_len);

    EVP_CIPHER_CTX_free(ctx);
    return true;
}

// Decrypts ciphertext using AES-256-GCM with BoringSSL.
bool DecryptAESGCM(const std::vector<uint8_t>& ciphertext,
    const std::vector<uint8_t>& key,
    std::vector<uint8_t>& plaintext) {
    if (key.size() != kKeySize || ciphertext.size() < kIvSize + kTagSize) {
        LOG(ERROR) << "Invalid key size or ciphertext length";
        return false;
    }

    // Extract IV, tag, and ciphertext.
    std::vector<uint8_t> iv(ciphertext.begin(), ciphertext.begin() + kIvSize);
    std::vector<uint8_t> tag(ciphertext.end() - kTagSize, ciphertext.end());
    std::vector<uint8_t> cipher_data(ciphertext.begin() + kIvSize, ciphertext.end() - kTagSize);

    EVP_CIPHER_CTX* ctx = EVP_CIPHER_CTX_new();
    if (!ctx) {
        LOG(ERROR) << "EVP_CIPHER_CTX_new failed";
        return false;
    }

    if (EVP_DecryptInit_ex(ctx, EVP_aes_256_gcm(), nullptr, key.data(), iv.data()) != 1) {
        EVP_CIPHER_CTX_free(ctx);
        LOG(ERROR) << "EVP_DecryptInit_ex failed";
        return false;
    }

    // Allocate sufficient space for decrypted text.
    plaintext.resize(cipher_data.size());
    int len = 0;

    // decrypt the data into plaintext.
    if (EVP_DecryptUpdate(ctx, plaintext.data(), &len, cipher_data.data(), cipher_data.size()) != 1) {
        EVP_CIPHER_CTX_free(ctx);
        LOG(ERROR) << "EVP_DecryptUpdate failed";
        return false;
    }
    int plaintext_len = len;

    // Set the authentication tag.
    if (EVP_CIPHER_CTX_ctrl(ctx, EVP_CTRL_GCM_SET_TAG, kTagSize, tag.data()) != 1) {
        EVP_CIPHER_CTX_free(ctx);
        LOG(ERROR) << "EVP_CIPHER_CTX_ctrl failed";
        return false;
    }

    // Finalize decryption.
    if (EVP_DecryptFinal_ex(ctx, plaintext.data() + len, &len) != 1) {
        EVP_CIPHER_CTX_free(ctx);
        LOG(ERROR) << "Authentication failed or decryption error";
        return false;  // 
    }
    plaintext_len += len;
    plaintext.resize(plaintext_len);  // Adjust size based on actual decrypted length

    EVP_CIPHER_CTX_free(ctx);
    return true;
}


#endif  // AES_CRYPT_H_