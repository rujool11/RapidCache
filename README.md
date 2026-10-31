# RapidCache

A high-performance, thread-safe in-memory cache service written in Go. RapidCache features lock-striped map sharding, dual eviction policies (TTL + LRU), disk persistence, and an HTTP-based API—complete with a native Go SDK.

## Features

- **Lock-Striped Sharding:** Eliminates global mutex bottlenecks under heavy parallel workloads.
- **Dual Eviction Policies:** Automatic background TTL expiration and Least Recently Used (LRU) capacity management.
- **Persistence:** Append-Only File (AOF) logging to survive server restarts.
- **HTTP Redis-Like API:** Exposes clean command endpoints over HTTP with API key authentication.
- **Native Go SDK:** Idiomatic client library to easily integrate RapidCache into any Go application.
