#include "testing/gtest/include/gtest/gtest.h"
#include <vector>
#include <string>

int main(int argc, char** argv) {
  // Create a modifiable copy of arguments
  std::vector<std::string> args(argv, argv + argc);

  // Check if --no-headless exists
  bool has_no_headless = false;
  for (const auto& arg : args) {
    if (arg == "--no-headless") {
      has_no_headless = true;
      break;
    }
  }

  // Add --no-headless if missing
  if (!has_no_headless) {
    args.push_back("--no-headless");
  }

  // Prepare arguments for InitGoogleTest
  std::vector<char*> arg_ptrs;
  for (auto& arg : args) {
    arg_ptrs.push_back(&arg[0]);  // Get pointer to string data
  }
  int new_argc = static_cast<int>(arg_ptrs.size());

  // Initialize Google Test
  testing::InitGoogleTest(&new_argc, arg_ptrs.data());

  return RUN_ALL_TESTS();
}
