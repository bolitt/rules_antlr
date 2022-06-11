load("@rules_antlr//antlr:antlr4.bzl", "antlr")
load("@io_bazel_rules_go//go:def.bzl", "go_binary", "go_library")

antlr(
    name = "generated",
    srcs = glob(["*.g4"]),
    language = "Go",
)

go_library(
    name = "parser",
    srcs = [
        ":generated",
    ],
    importpath = "github.com/example/com/json/parser",
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
        ":parser",
        "@antlr4_runtimes//:go",
    ],
)
