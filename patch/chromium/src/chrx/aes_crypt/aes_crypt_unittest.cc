#include "aes_crypt.h"
#include "testing/gtest/include/gtest/gtest.h"
#include <string>
#include <vector>

// Test for AES-GCM encryption and decryption
TEST(AesCryptTest, EncryptDecryptTest) {
    std::string plaintext = "Secret message for AES-GCM encryption!";
    std::string passphrase = "strong-passphrase";

    // Derive key from passphrase
    std::vector<uint8_t> key = NewSHA256(passphrase);
    ASSERT_EQ(key.size(), kKeySize) << "Key size must be 32 bytes (AES-256)";

    std::vector<uint8_t> plain_bytes(plaintext.begin(), plaintext.end());
    std::vector<uint8_t> ciphertext;
    std::vector<uint8_t> decrypted;

    // Test encryption
    bool enc_ok = EncryptAESGCM(plain_bytes, key, ciphertext);
    ASSERT_TRUE(enc_ok) << "Encryption failed";

    // Ensure ciphertext is not empty
    ASSERT_FALSE(ciphertext.empty()) << "Ciphertext is empty";

    // Test decryption
    bool dec_ok = DecryptAESGCM(ciphertext, key, decrypted);
    ASSERT_TRUE(dec_ok) << "Decryption failed";

    // Convert decrypted bytes back to string
    std::string decrypted_str(decrypted.begin(), decrypted.end());

    // Check that the decrypted string matches the original plaintext
    EXPECT_EQ(plaintext, decrypted_str) << "Decrypted text does not match original";
}