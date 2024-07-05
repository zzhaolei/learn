const std = @import("std");

pub fn Lockable(comptime T: type) type {
    return struct {
        const Self = @This();
        mu: std.Thread.Mutex,
        value: T,

        pub fn init(val: T) Self {
            return Self{
                .mu = std.Thread.Mutex{},
                .value = val,
            };
        }

        pub fn set(self: *Self, value: T) void {
            self.mu.lock();
            defer self.mu.unlock();

            self.value = value;
        }

        pub fn get(self: *Self) T {
            self.mu.lock();
            defer self.mu.unlock();

            return self.value;
        }
    };
}
