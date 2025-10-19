# Benchmarks

This document provides setup instructions and usage details for running benchmarks in this repository.

## Prerequisites

Before running benchmarks, ensure you have the following tools installed:

### 1. Docker Compose (or Podman Compose)

- **Docker Compose**:  
    Install Docker Desktop, which includes Docker Compose.  
    [Docker Compose Install Guide](https://docs.docker.com/compose/install/)

- **Podman Compose** (alternative):  
    Install Podman and Podman Compose if you prefer a rootless container runtime.  
    [Podman Install Guide](https://podman.io/getting-started/installation)  
    [Podman Compose Install Guide](https://github.com/containers/podman-compose)

- **Verify Installation**:
    ```sh
    docker-compose --version
    # or
    podman-compose --version
    ```

### 2. ghz

- **ghz** is a benchmarking and load testing tool for gRPC services.

- **Install via Homebrew (macOS/Linux):**
    ```sh
    brew install ghz
    ```

- **Install via Go:**
    ```sh
    go install github.com/bojand/ghz/cmd/ghz@latest
    ```

- **Verify Installation:**
    ```sh
    ghz --version
    ```

---

## Environment Setup

Before running any benchmarks, make sure the gRPC server and any required services are running. You can start the environment using Docker Compose or Podman Compose:

```sh
docker-compose -f deploy/docker-compose/docker-compose.yml up -d --build
# or
podman-compose -f deploy/docker-compose/docker-compose.yml up -d --build
```

---

## Directory Structure

```
benchmarks/
├── ghz/
│   ├── next_snowflake.toml
│   ├── batch_next_snowflake.toml
│   └── README.md
├── scripts/
│   └── benchmark.sh
├── reports/
│   ├── next_snowflake_report.log
│   └── batch_next_snowflake_report.log
└── README.md
```

- `ghz/` – Benchmark configuration files for each RPC.
- `scripts/` – Helper scripts to automate benchmarking.
- `reports/` – Output and reports from benchmark runs.

---

## Running Benchmarks

To run all benchmarks and generate reports, use the provided script:

```sh
cd benchmarks/scripts
./benchmark.sh
```

This will run the benchmarks defined in the `ghz/` configs and save summary reports to the `reports/` directory.

---

## Customizing Benchmarks

- Edit the `.toml` files in `benchmarks/ghz/` to change parameters such as concurrency, total requests, or target host.
- You can run individual benchmarks manually with:
    ```sh
    ghz --config ../ghz/next_snowflake.toml --output ../reports/next_snowflake_report.log --format summary
    ```

---

## Viewing Results

- Reports are saved in the `benchmarks/reports/` directory.
- Example: `next_snowflake_report.log` contains a summary of the NextSnowflake RPC benchmark.

---

## References

- [ghz Documentation](https://ghz.sh/docs/)
- [gRPC](https://grpc.io/)

---

Contributions and improvements to the benchmarks are welcome!