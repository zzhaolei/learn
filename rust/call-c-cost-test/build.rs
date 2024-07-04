fn main() {
    cc::Build::new()
        .file("src/rand.c")
        .compile("rand");
}
