use serde::{Deserialize, Serialize};
use std::time::Instant;

#[derive(Serialize, Deserialize, Debug)]
struct Person {
    name: String,
    age: u32,
    email: String,
    address: Address,
    hobbies: Vec<String>,
}

#[derive(Serialize, Deserialize, Debug)]
struct Address {
    street: String,
    city: String,
    country: String,
    zip_code: String,
}

fn generate_test_data(count: usize) -> Vec<Person> {
    (0..count)
        .map(|i| Person {
            name: format!("Person {}", i),
            age: 20 + (i as u32 % 50),
            email: format!("person{}@example.com", i),
            address: Address {
                street: format!("{} Main Street", i + 1),
                city: "Example City".to_string(),
                country: "Example Country".to_string(),
                zip_code: format!("{:05}", 10000 + i),
            },
            hobbies: vec![
                "reading".to_string(),
                "coding".to_string(),
                "gaming".to_string(),
            ],
        })
        .collect()
}

fn main() {
    const DATA_SIZE: usize = 10000;
    
    println!("Rust JSON Benchmark with {} records", DATA_SIZE);
    
    // Generate test data
    let people = generate_test_data(DATA_SIZE);
    
    // Serialize benchmark
    let start = Instant::now();
    let json_data = serde_json::to_string(&people).expect("Serialization failed");
    let serialize_duration = start.elapsed();
    
    println!("Serialization time: {:?}", serialize_duration);
    println!("JSON size: {} bytes", json_data.len());
    
    // Deserialize benchmark
    let start = Instant::now();
    let deserialized_people: Vec<Person> = 
        serde_json::from_str(&json_data).expect("Deserialization failed");
    let deserialize_duration = start.elapsed();
    
    println!("Deserialization time: {:?}", deserialize_duration);
    println!("Deserialized {} records", deserialized_people.len());
}