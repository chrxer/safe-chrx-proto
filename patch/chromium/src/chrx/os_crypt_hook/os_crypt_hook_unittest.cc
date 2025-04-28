#include "gtest/gtest.h"
#include "os_crypt.h"

TEST(ChrxEncryptDecryptTest, EncryptDecryptTest) {
    std::string plaintext = "This is a test message!";
    std::string encrypted_text;
    std::string decrypted_text;

    ASSERT_TRUE(ChrxEncrypt(plaintext, encrypted_text));
    ASSERT_TRUE(ChrxDecrypt(decrypted_text, encrypted_text));
    ASSERT_EQ(plaintext, decrypted_text);
}

TEST(ChrxEncryptDecryptTest, DecryptWithIncorrectCiphertext) {
    std::string incorrect_ciphertext = "InvalidCiphertext!";
    std::string decrypted_text;
    ASSERT_FALSE(ChrxDecrypt(decrypted_text, incorrect_ciphertext));
}

TEST(ChrxEncryptDecryptTest, EncryptEmptyPlaintext) {
    std::string plaintext = "";
    std::string encrypted_text;
    
    ASSERT_TRUE(ChrxEncrypt(plaintext, encrypted_text));
    ASSERT_TRUE(encrypted_text.empty());
}

TEST(ChrxEncryptDecryptTest, DecryptEmptyCiphertext) {
    std::string empty_ciphertext = "";
    std::string decrypted_text;

    ASSERT_TRUE(ChrxDecrypt(decrypted_text, empty_ciphertext));
    ASSERT_TRUE(decrypted_text.empty());
}
