# nats-websocket

This example leverages NATS to broadcast messages across multiple server instances that use WebSocket to deliver real-time information to clients. For instance, if you have multiple servers running and deployed via Kubernetes for load balancing, NATS ensures that all server instances receive the same broadcast messages to update the client side in real-time.

## Getting Started

1. Clone the repository
2. Navigate to the nats-websocket directory

    ```sh
    cd nats-websocket
    ```

3. Start the nats-server via docker

   ```sh
    docker run --name nats --rm -p 4222:4222 -p 8222:8222 nats --http_port 8222
   ```

4. Start the server 1 instance

    ```sh
    go run main.go server1
    ```

5. Start the server 2 instance

    ```sh
    go run main.go server2
    ```

6. Connect to both server instances websocket endpoints
7. Publish message via nats client cli tool

   ```sh
   nats pub foo "hello, world!"
   ```

8. Both server instances should receive the same message
