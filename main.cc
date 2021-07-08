#include <gtest/gtest.h>

int main(int argc, char** argv) {
    // testing::internal::CaptureStdout();
    testing::InitGoogleTest(&argc, argv);

    // Runs all tests using Google Test.
    return RUN_ALL_TESTS();
 }