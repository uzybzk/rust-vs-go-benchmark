package main

import (
    "encoding/json"
    "fmt"
    "time"
)

type Person struct {
    Name     string `json:"name"`
    Age      int    `json:"age"`
    Email    string `json:"email"`
    Address  Address `json:"address"`
    Hobbies  []string `json:"hobbies"`
}

type Address struct {
    Street  string `json:"street"`
    City    string `json:"city"`
    Country string `json:"country"`
    ZipCode string `json:"zip_code"`
}

func generateTestData(count int) []Person {
    people := make([]Person, count)
    
    for i := 0; i < count; i++ {
        people[i] = Person{
            Name:  fmt.Sprintf("Person %d", i),
            Age:   20 + (i % 50),
            Email: fmt.Sprintf("person%d@example.com", i),
            Address: Address{
                Street:  fmt.Sprintf("%d Main Street", i+1),
                City:    "Example City",
                Country: "Example Country",
                ZipCode: fmt.Sprintf("%05d", 10000+i),
            },
            Hobbies: []string{"reading", "coding", "gaming"},
        }
    }
    
    return people
}

func main() {
    const dataSize = 10000
    
    fmt.Printf("Go JSON Benchmark with %d records\n", dataSize)
    
    // Generate test data
    people := generateTestData(dataSize)
    
    // Serialize benchmark
    start := time.Now()
    jsonData, err := json.Marshal(people)
    if err != nil {
        panic(err)
    }
    serializeDuration := time.Since(start)
    
    fmt.Printf("Serialization time: %v\n", serializeDuration)
    fmt.Printf("JSON size: %d bytes\n", len(jsonData))
    
    // Deserialize benchmark
    var deserializedPeople []Person
    start = time.Now()
    err = json.Unmarshal(jsonData, &deserializedPeople)
    if err != nil {
        panic(err)
    }
    deserializeDuration := time.Since(start)
    
    fmt.Printf("Deserialization time: %v\n", deserializeDuration)
    fmt.Printf("Deserialized %d records\n", len(deserializedPeople))
}