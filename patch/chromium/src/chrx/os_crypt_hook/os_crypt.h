#pragma once
#include <string>

bool ChrxEncrypt(const std::string& plaintext, std::string& ciphertext);
bool ChrxDecrypt(std::string& plaintext, const std::string& ciphertext);