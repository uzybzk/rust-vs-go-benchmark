# Rust vs Go Performance Comparison

Comparing performance between Rust and Go implementations of common algorithms.

## Benchmarks

### Fibonacci Sequence

Implementations:
- Recursive approach (inefficient but demonstrates raw performance)
- Iterative approach (efficient)

**Running the benchmarks:**

Go:
```bash
cd go
go run fibonacci.go
```

Rust:
```bash
cd rust
cargo run --bin fibonacci --release
```

### Results (Preliminary)

Test system: Local development machine
Input: Fibonacci(40)

| Implementation | Language | Time (Debug) | Time (Release) |
|----------------|----------|--------------|----------------|
| Recursive      | Go       | ~1.2s        | N/A            |
| Recursive      | Rust     | ~2.1s        | ~0.8s          |
| Iterative      | Go       | <1ms         | N/A            |
| Iterative      | Rust     | <1ms         | <1ms           |

### Observations

1. **Recursive Fibonacci**: Rust with `--release` flag significantly outperforms Go
2. **Iterative Fibonacci**: Both languages perform similarly for efficient algorithms
3. **Compilation**: Rust's release builds are heavily optimized
4. **Development**: Go's faster compilation makes development iteration quicker

### Next Benchmarks

- [ ] HTTP server performance
- [ ] JSON parsing/serialization
- [ ] Concurrent workloads
- [ ] Memory usage comparison
- [ ] Binary size comparison

## Notes

This is a learning exercise comparing two languages I'm working with. Each has strengths:

- **Go**: Simple, fast compilation, great for microservices
- **Rust**: Zero-cost abstractions, memory safety, excellent performance

Both are excellent choices for backend development, depending on requirements.