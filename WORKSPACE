workspace(name = "rules_antlr")

load("@bazel_tools//tools/build_defs/repo:git.bzl", "git_repository")
load("@bazel_tools//tools/build_defs/repo:http.bzl", "http_archive", "http_jar")

local_repository(
    name = "examples",
    path = "./examples",
)

http_jar(
    name = "junit",
    sha256 = "59721f0805e223d84b90677887d9ff567dc534d7c502ca903c0c2b17f05c116a",
    url = "https://jcenter.bintray.com/junit/junit/4.12/junit-4.12.jar",
)

http_jar(
    name = "jimfs",
    sha256 = "c4828e28d7c0a930af9387510b3bada7daa5c04d7c25a75c7b8b081f1c257ddd",
    url = "https://jcenter.bintray.com/com/google/jimfs/jimfs/1.1/jimfs-1.1.jar",
)

http_jar(
    name = "guava",
    sha256 = "4a5aa70cc968a4d137e599ad37553e5cfeed2265e8c193476d7119036c536fe7",
    url = "https://jcenter.bintray.com/com/google/guava/guava/27.1-jre/guava-27.1-jre.jar",
)

load("//antlr:repositories.bzl", "rules_antlr_dependencies")

rules_antlr_dependencies(2, 3, 4)

http_archive(
    name = "io_bazel_stardoc",
    sha256 = "aa814dae0ac400bbab2e8881f9915c6f47c49664bf087c409a15f90438d2c23e",
    urls = [
        "https://mirror.bazel.build/github.com/bazelbuild/stardoc/releases/download/0.5.1/stardoc-0.5.1.tar.gz",
        "https://github.com/bazelbuild/stardoc/releases/download/0.5.1/stardoc-0.5.1.tar.gz",
    ],
)

load("@io_bazel_stardoc//:setup.bzl", "stardoc_repositories")

stardoc_repositories()
