#include "testing/gtest/include/gtest/gtest.h"

// regular unittests executable (headless) for chrx
int main(int argc, char **argv) {
    ::testing::InitGoogleTest(&argc, argv);
    return RUN_ALL_TESTS();
}
