# Golang Backend API using Echo for FSI(File System Integrity) Application
Developing a simple Golang backend API using the [Echo framework](https://github.com/labstack/echo). This API stores IDs and their corresponding hashes in a SQL server and provides functionality to verify if a given hash matches the stored hash for a specific ID, thereby providing integrity service that aligns with the CIA (Confidentiality, Integrity, Availability) triad for data security.

## Getting Started

To get started with this project, follow these steps:

### Prerequisites

- Go installed on your machine.
- A SQL database server to store IDs and hashes.

### Installation

1. Clone this repository to your local machine:

  ```bash
  git clone https://github.com/justsaumit/go-fsi-api.git
  cd go-fsi-api
  ```

2. Initialize and install project dependencies using Go modules:

  ```bash
  go mod vendor
  ```

3. Create a configuration file (e.g., config.yaml) to specify your database settings and other configurations.

4. Run the server:

  ```bash
  go run main.go
  ```

### Usage

Once the server is running, you can access the API endpoints to add file hashes and verify them.

- To add a file hash, make a POST request to `/add` with JSON data containing the ID and hash.
- To verify a file hash, make a POST request to `/verify` with JSON data containing the ID and hash.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
