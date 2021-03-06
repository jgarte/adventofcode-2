load("@bazel_tools//tools/build_defs/repo:git.bzl", "git_repository")

def cpp_dependencies():
    git_repository(
        name = "googletest",
        remote = "https://github.com/google/googletest",
        # `commit` and `shallow_since` was given by first specifying:
        #     tag = "release-1.10.0"
        # and then following the debug messages given by Bazel.
        commit = "703bd9caab50b139428cea1aaff9974ebee5742e",
        shallow_since = "1570114335 -0400",
    )

    git_repository(
        name = "googlebenchmark",
        remote = "https://github.com/google/benchmark",
        # `commit` and `shallow_since` was given by first specifying:
        #     tag = "v1.5.0"
        # and then following the debug messages given by Bazel.
        commit = "090faecb454fbd6e6e17a75ef8146acb037118d4",
        shallow_since = "1557776538 +0300",
    )

    git_repository(
        name = "abseil",
        remote = "https://github.com/abseil/abseil-cpp",
        # `commit` and `shallow_since` was given by first specifying:
        #     tag = "20190808"
        # and then following the debug messages given by Bazel.
        commit = "aa844899c937bde5d2b24f276b59997e5b668bde",
        shallow_since = "1565288385 -0400",
    )
