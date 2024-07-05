const std = @import("std");
const Lockable = @import("lockable.zig").Lockable;

pub fn main() !void {
    var lock = Lockable(u8).init(1);
    const v = lock.get();
    std.debug.print("{}\n", .{v});
}

test "simple test" {
    var lock = Lockable(u8).init(1);
    try std.testing.expectEqual(@as(u8, 1), lock.get());
    lock.set(10);
    try std.testing.expectEqual(@as(u8, 10), lock.get());

    const t = struct {
        a: u8 = 1,
        b: u8 = 2,
    };
    var lock1 = Lockable(t).init(.{});
    try std.testing.expectEqual(@as(u8, 2), lock1.get().b);
}
