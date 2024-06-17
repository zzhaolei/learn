const std = @import("std");

pub fn build(b: *std.Build) void {
    const target = b.standardTargetOptions(.{});
    const optimize = b.standardOptimizeOption(.{});

    const sqlite = b.addStaticLibrary(.{
        .name = "sqlite",
        .target = target,
        .optimize = optimize,
    });
    sqlite.addIncludePath(b.path("sqlite3"));
    sqlite.addCSourceFiles(.{
        .files = &.{
            "sqlite3/sqlite3.c",
        },
        .flags = &.{
            "-std=c99",
        },
    });
    sqlite.installHeader(b.path("sqlite3/sqlite3.h"), "sqlite3.h");
    sqlite.linkLibC();
    b.installArtifact(sqlite);

    const exe = b.addExecutable(.{
        .name = "demo",
        .root_source_file = b.path("src/main.zig"),
        .target = target,
        .optimize = optimize,
    });
    exe.linkLibrary(sqlite);

    b.installArtifact(exe);

    const run_cmd = b.addRunArtifact(exe);
    run_cmd.step.dependOn(b.getInstallStep());

    if (b.args) |args| {
        run_cmd.addArgs(args);
    }

    const run_step = b.step("run", "Run the app");
    run_step.dependOn(&run_cmd.step);

    const exe_unit_tests = b.addTest(.{
        .root_source_file = b.path("src/main.zig"),
        .target = target,
        .optimize = optimize,
    });

    const run_exe_unit_tests = b.addRunArtifact(exe_unit_tests);

    const test_step = b.step("test", "Run unit tests");
    test_step.dependOn(&run_exe_unit_tests.step);
}
