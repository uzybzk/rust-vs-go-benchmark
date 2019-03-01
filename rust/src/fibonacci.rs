use std::time::Instant;

fn fibonacci(n: u32) -> u32 {
    if n <= 1 {
        return n;
    }
    fibonacci(n - 1) + fibonacci(n - 2)
}

fn fibonacci_iterative(n: u32) -> u32 {
    if n <= 1 {
        return n;
    }
    
    let mut a = 0;
    let mut b = 1;
    
    for _ in 2..=n {
        let temp = a + b;
        a = b;
        b = temp;
    }
    
    b
}

fn main() {
    let n = 40;
    
    // Recursive version
    let start = Instant::now();
    let result = fibonacci(n);
    let recursive_duration = start.elapsed();
    
    println!("Rust Recursive Fibonacci({}) = {}", n, result);
    println!("Time: {:?}", recursive_duration);
    
    // Iterative version
    let start = Instant::now();
    let result = fibonacci_iterative(n);
    let iterative_duration = start.elapsed();
    
    println!("Rust Iterative Fibonacci({}) = {}", n, result);
    println!("Time: {:?}", iterative_duration);
}