load("@rules_antlr//antlr:antlr4.bzl", "antlr")
load("@io_bazel_rules_go//go:def.bzl", "go_binary", "go_library")

antlr(
    name = "generated",
    srcs = glob(["*.g4"]),
    language = "Go",
    package = "json",
)

go_library(
    name = "json",
    srcs = [
        ":generated",
    ],
    importpath = "github.com/example/com/parser/json",
    visibility = ["//visibility:public"],
    deps = [
        "@antlr4_runtimes//:go",
    ],
)

go_binary(
    name = "main",
    srcs = [
        "main.go",
    ],
    deps = [
        ":json",
        "@antlr4_runtimes//:go",
    ],
)
