const std = @import("std");
const overloading = @import("overloading");

fn test_a(name: []const u8, age: u8) u8 {
    std.debug.print("name: {s}, age: {d}\n", .{ name, age });
    return age;
}

fn test_b(age: u8, name: []const u8) u8 {
    std.debug.print("name: {s}, age: {d}\n", .{ name, age });
    return age;
}

fn add0() i32 {
    return 0;
}
fn add1(a: i32) i32 {
    return a;
}
fn add2(a: i32, b: i32) i32 {
    return a + b;
}
fn add3(a: i32, b: i32, c: i32) i32 {
    return a + b + c;
}

pub fn main() !void {
    const add = comptime add: {
        break :add overloading.make(.{
            add0,
            add1,
            add2,
            add3,
        });
    };

    const one = add({});
    const two = add(2);
    const three = add(.{ 50, 2 });
    std.debug.print("{any}, {any}, {any}\n", .{ one, two, three });

    const f = comptime f: {
        break :f overloading.make(.{
            test_a,
            test_b,
        });
    };
    const none1 = f(.{ "name", 18 });
    std.debug.print("{any}\n", .{none1});
}
