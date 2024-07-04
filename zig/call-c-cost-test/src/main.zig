const std = @import("std");
const print = std.debug.print;

extern "c" fn rand() c_int;

fn bench() !void {
    var timer = try std.time.Timer.start();
    var total: u64 = 0;

    var i: usize = 0;
    const count = 2_000_000;
    timer.reset();
    while (i < count) : (i += 1) {
        _ = rand();
    }
    total = timer.read();
    print("bench: \ncount: {}, total: {}, single: {}\n", .{ count, std.fmt.fmtDuration(total), std.fmt.fmtDuration(total / count) });
}

pub fn main() !void {
    try bench();
}
