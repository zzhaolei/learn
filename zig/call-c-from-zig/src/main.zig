const std = @import("std");
const arithmetic = @cImport({
    @cInclude("arithmetic.c");
});

fn add(x: i32, y: i32) i32 {
    return arithmetic.add(x, y);
}

pub fn main() !void {
    const x = 5;
    const y = 12;
    const z = add(x, y);

    const stdout_file = std.io.getStdOut().writer();
    var bw = std.io.bufferedWriter(stdout_file);
    const stdout = bw.writer();

    try stdout.print("{} + {} = {}\n", .{ x, y, z });

    try bw.flush();
}

test "test add" {
    try std.testing.expectEqual(@as(i32, 17), add(5, 12));
}
