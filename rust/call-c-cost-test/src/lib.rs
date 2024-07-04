#[link(name = "rand")]
extern "C" {
    fn rand() -> i32;
}

pub fn call_rand() -> i32 {
    unsafe { rand() }
}
