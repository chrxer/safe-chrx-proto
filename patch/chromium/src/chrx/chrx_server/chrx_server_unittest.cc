#include "gtest/gtest.h"
#include "chrx_server.h"
#include "base/logging.h"

TEST_F(CryptServerLauncher, GetPortReturnsValidPort) {
    int port = CryptServerLauncher::Instance().GetPort();
    EXPECT_GT(port, 0) << "Port should be greater than zero.";
    EXPECT_LT(port, 65536) << "Port should be a valid TCP port.";
}

TEST_F(CryptServerLauncher, GetKeyReturnsNonEmptyKey) {
    const std::string& key = CryptServerLauncher::Instance().GetKey();
    EXPECT_FALSE(key.empty()) << "AES key should not be empty.";
    EXPECT_GT(key.length(), 8u) << "AES key should be reasonably long.";
}

TEST_F(CryptServerLauncher, GetPortAndGetKeyConsistency) {
    auto& launcher = CryptServerLauncher::Instance();
    int port1 = launcher.GetPort();
    int port2 = launcher.GetPort();
    EXPECT_EQ(port1, port2) << "Port should be consistent across calls.";

    const std::string& key1 = launcher.GetKey();
    const std::string& key2 = launcher.GetKey();
    EXPECT_EQ(key1, key2) << "AES key should be consistent across calls.";
}
