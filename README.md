# NATS Example

This repository showcases various use cases of using NATS (NATS.io) as a powerful messaging system. NATS is a high-performance messaging system used for building distributed systems, microservices, and cloud-native applications. This repository aims to provide a comprehensive set of examples demonstrating how to leverage NATS for different messaging patterns and scenarios.

## Examples

1. nats-websocket
   - Leveraging NATS to broadcast messages across multiple server instances that use WebSocket to deliver real-time information to clients.
   - In my use case, I have multiple servers running and deployed via Kubernetes for load balancing. NATS is used to ensure all server instances receive the same broadcast messages to update the client side in real-time.
