# nats-websocket

To showcase the capability of using NATS in delivering messages to all connected apps/clients

# nats-docker

```
docker run --name nats --rm -p 4222:4222 -p 8222:8222 -p 9090:9090 -v $(pwd)/nats-server.conf:/etc/nats/nats-server.conf nats --http_port 8222 -c /etc/nats/nats-server.conf

```