# Carbon - A Lightweight In-Memory Cache for High Throughput  

**Carbon** is a lightweight, high-performance in-memory cache library for Go, designed to meet the demands of APIs requiring fast and efficient data caching. I created Carbon as an alternative to Redis when it proved too slow for some of my APIs and when managing invalidation in a `sync.Map` became overly complex.  

At its core, Carbon leverages a thread-safe `sync.Map` to ensure simplicity and reliability. The cache operates locally, meaning only the application instance that initialized it can access the data, ensuring isolation and high performance.  

### Key Invalidation  
Carbon provides two efficient mechanisms for invalidating expired keys:  
- **Background Expiry Check (Optional):** A configurable Go routine runs in the background, periodically checking for expired items and automatically removing them.  
- **On-Demand Expiry Check:** When retrieving a key, Carbon checks if it is expired. If expired, the key is deleted immediately, and `nil` is returned.  

### Why Use Carbon?  
- **When Redis is Too Slow:** If your application requires low-latency caching and Redis introduces unnecessary overhead, Carbon offers a streamlined alternative.  
- **Simplicity and Expiry Management:** Carbon provides a straightforward in-memory cache solution with built-in support for key expiry and flexible invalidation strategies.  
- **Local Access Only:** Perfect for use cases where the cache doesn't need to be shared across multiple instances of your application.  

**When Not to Use Carbon:**  
If your cache needs to be accessible by multiple instances or services, Redis or a distributed caching solution would be a better fit.

Thank you for providing the code! Based on your code, here's a detailed usage section for the README documentation:

---

## Installation

To get started with **Carbon**, install it via `go get`:

```bash
go get github.com/scott-mescudi/carbon
```

## Basic Usage

Carbon provides an easy-to-use in-memory caching solution with built-in expiry handling. Here's how to use it:

### Create a New Cache Store

To create a new instance of the **CarbonStore**, use the `NewCarbonStore` function. You can optionally pass in a `cleanFrequency` to periodically clean expired items from the store. If you do not want the cleaner routine to run, pass `carbon.NoClean`.

```go
package main

import (
    "fmt"
    "github.com/scott-mescudi/carbon"
    "time"
)

func main() {
    // Create a new Carbon store with a 5-minute cleanup frequency
    cache := carbon.NewCarbonStore(5 * time.Minute)

    // Set a cache key with a 10-second expiry
    cache.Set("user:123", "John Doe", 10 * time.Second)

    // Retrieve a cached value
    value, err := cache.Get("user:123")
    if err != nil {
        fmt.Println("Error:", err)
        return
    } 
    
    fmt.Println("Cached Value:", value)
    
}
```

Alternatively, to create a store without the cleaner:

```go
// Create a new Carbon store without the cleaner routine
cache := carbon.NewCarbonStore(carbon.NoClean)
```

### Importing a Cache Store from a File

You can also initialize the **CarbonStore** by loading data from a file. The `ImportStoreFromFile` function parses a file, matches key-value pairs using a regular expression, and loads them into the cache.

```go
cache, err := carbon.ImportStoreFromFile("cache_data.txt", 5 * time.Minute)
if err != nil {
    fmt.Println("Error loading cache:", err)
} else {
    fmt.Println("Cache loaded successfully!")
}
```

### Setting and Getting Cached Values

You can store values in the cache with a specified expiration time. If the expiration time is not set, the key will persist indefinitely until it is manually deleted. 

#### Set a Cache Key

You can set keys with the following options:
- **Expiry time:** Specify a duration after which the key will expire.
- **No expiry:** Use the `carbon.NoExpiry` flag to keep the key indefinitely.

```go
// Set a cache key that expires in 10 seconds
cache.Set("user:123", "John Doe", 10 * time.Second)

// Set a cache key with no expiry (keeps the key forever, or until manually deleted)
cache.Set("user:124", "Jane Doe", carbon.NoExpiry)

// Set a cache key with no expiry and no cleaner
cache.Set("user:125", "Alice", carbon.NoExpiry)
```

#### Get a Cache Key

```go
value, err := cache.Get("user:123")
if err != nil {
    fmt.Println("Error:", err)
    return
} 

fmt.Println("Value:", value)

```

If the key is expired, it will be removed from the cache, and an error will be returned:

```go
// If the key has expired
value, err := cache.Get("user:123")
if err != nil {
    fmt.Println("Error:", err)  // Output: 'user:123' is expired
}
```

### Deleting Keys

You can delete keys manually from the cache:

```go
cache.Delete("user:123")
```

### Clearing the Cache Store

To clear all keys in the cache:

```go
cache.ClearStore()
```

### Stopping the Cache Cleaner

The background cleaner is responsible for removing expired items periodically. If you want to stop the cleaner, you can call the `StopCleaner` method:

```go
cache.StopCleaner()
```

### Backup and Restore Cache

To back up your cache to a file, use the `BackupToFile` method:

```go
err := cache.BackupToFile("backup.txt")
if err != nil {
    fmt.Println("Error backing up cache:", err)
} else {
    fmt.Println("Cache backed up successfully!")
}
```

To import the backup into a new **CarbonStore** instance, use `ImportStoreFromFile`.

---

## Key Invalidation Mechanisms

Carbon provides two mechanisms for invalidating expired keys:

- **Background Expiry Check (Optional):** A background Go routine runs at a configurable interval (`cleanFrequency`) and checks for expired items. Expired items are deleted automatically.
   
   Example of setting up the cleaner with a 5-minute interval:
   ```go
   cache := carbon.NewCarbonStore(5 * time.Minute)
   ```

- **On-Demand Expiry Check:** Each time you attempt to retrieve a cached value, Carbon will check if it is expired. If expired, the key is deleted, and `nil` is returned.

---

## Examples

To see more use cases and examples of how to implement **Carbon** in your Go applications, check out the [examples folder](./examples). There, you'll find several example programs that demonstrate different features of the Carbon cache, such as basic caching, using a file-based store, key expiry, and more advanced features like cache backups and cleaner configurations.

---

## Todo

- **Dynamic TTL (Time-To-Live) Adjustments**  
  Allow developers to extend the TTL of cache keys dynamically, keeping frequently accessed items alive longer without requiring a full cache reset.

- **Atomic Operations and Transactions**  
  Support atomic operations like compare-and-swap (CAS) for cache keys, enabling safe, conditional updates to cache values based on their current state.

- **Automatic Expiry Based on Events**  
  Add event-driven expiry, where cache keys automatically expire when certain conditions or external events (e.g., database updates) are triggered.

- **Cache Size Limit and Monitoring**  
  Introduce a maximum cache size (either by memory or number of entries) with configurable eviction policies to automatically manage cache size, along with monitoring features to track cache performance.