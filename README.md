# Waiter: A Load Simulation Image

This is a very image to be able to simulate load inside a docker environment. I
developed this to test my Kubernetes Operator implementations and simulating
long running jobs. Feel free to use it.

## Usage

```bash
docker run --rm ghcr.io/flohansen/waiter:latest -timeout 30s
```
