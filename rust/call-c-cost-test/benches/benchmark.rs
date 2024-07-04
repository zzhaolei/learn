use criterion::{criterion_group, criterion_main, Criterion};
use call_c_cost_test::call_rand;

fn benchmark(c: &mut Criterion) {
    c.bench_function("Benchmark", |b| {
        b.iter(|| {
            call_rand();
        })
    });
}

criterion_group!(benches, benchmark);
criterion_main!(benches);
