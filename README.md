# Golang Backend API for FIS (File Integrity Surveillance) Application

[![Docker Image CI for GHCR](https://github.com/justsaumit/go-fis-api/actions/workflows/publish-ghcr.yml/badge.svg)](https://github.com/justsaumit/go-fis-api/actions/workflows/publish-ghcr.yml)
[![Go-Releaser](https://github.com/justsaumit/go-fis-api/actions/workflows/go-releaser.yml/badge.svg)](https://github.com/justsaumit/go-fis-api/actions/workflows/go-releaser.yml)

Developing a simple Golang backend API with the [Echo framework](https://github.com/labstack/echo) for FIS(File Integrity Surveillance) application which can be found [here](https://github.com/ayato91/Fair-Files).  
This API stores IDs and their corresponding hashes in a SQL server and provides functionality to verify if a given hash matches the stored hash for a specific ID. All the communication between the Application and API is secured using TLS encryption(HTTPS).  
Thereby providing both confidentiality and integrity service that aligns with the CIA (Confidentiality, Integrity, Availability) triad for data security.

## Getting Started

To get started with this project, follow these steps:

### Prerequisites

- Go installed on your machine.
- A SQL database server (sqlite) to store IDs and hashes.

### Installation

**1. Clone this repository:**

```bash
git clone https://github.com/justsaumit/go-fis-api.git
cd go-fis-api
```

**2. Initialize project dependencies:**

```bash
go mod vendor
```

**3. Configure environment variables (optional):**

- Create a `.env` file at the root of the project and specify environment variables.
- For example (refer to env.example):

```dotenv
ENVIRONMENT=<development or production>
PORT=3000

# SSL/TLS Configuration (if running in container, make it /certs/<cert/key>.pem and mount the certificate directory as a volume to /certs)
CERTPATH=<path to>/fullchain.pem
KEYPATH=<path to>/privkey.pem
```

**4. Run the server:**

**Option A: Run with Go (Development Mode):**

```bash
go run server.go
```

**Option B: Build and Run:**

```bash
go build -o main .
./main
```

**Option C: Using Docker:**

**a. Build the image (if using Dockerfile):**

```bash
docker build -t myapp .
```

**b. Pull the prebuilt image (alternative):**

```bash
docker pull ghcr.io/justsaumit/go-fis-api:latest
```

**c. Run the container:**

- **Development:**

```bash
docker run -p 3000:3000 myapp
```

- **Production (with certificates and persistent storage):**

```bash
docker run -p 3000:3000 --env-file .env -v docker-dbvolume:/app -v /path/to/certifcates:/certs myapp
```

### Usage

Once the server is running, you can access the API endpoints to upload a file for hashing and verify them.

- **Adding a File Hash:**
  - To add a file hash to the DB, make a POST request to `/upload` by uploading the file using multipart form. The server will generate an ID and hash for the uploaded file, which will be returned as a JSON response.

- **Verifying a File Hash:**
  - To verify a file hash, make a POST request to `/verify` with the previously generated ID and the file using multipart form. The server will respond with a JSON message indicating whether the hash of the uploaded file (calculatedHash) matches with the hash in the database (storedHash) corresponding to the given ID. It also indicates whether an ID is present in the database or not.

## To-Do-List

- [x] Handle Uploaded files (API)
- [x] Perform short ID Generation (API)
- [x] Perform Hashing (API)
- [x] Connect with DB
- [x] Store JSON data to DB
- [x] Perform verification
- [x] Containerize the API
- [x] Optimize Docker Image size (Previously 509MB, Optimized: ~~22.2MB~~ 20.8MB)
- [x] Replace the [go-sqlite3](https://pkg.go.dev/github.com/mattn/go-sqlite3) driver; with a CGO-free port, written in pure Go - [modernc-sqlite](https://pkg.go.dev/modernc.org/sqlite) (to avoid CGO cross-compilation errors)
- [x] Create a Github workflow for binary releases and docker image packages

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
