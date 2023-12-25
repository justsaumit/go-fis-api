# Golang Backend API for FIS (File Integrity Surveillance) Application

Developing a simple Golang backend API with the [Echo framework](https://github.com/labstack/echo) for FIS(File Integrity Surveillance) application which can be found [here](https://github.com/ayato91/Fair-Files).  
This API stores IDs and their corresponding hashes in a SQL server and provides functionality to verify if a given hash matches the stored hash for a specific ID. All the communication between the Application and API is secured using TLS encryption(HTTPS).  
Thereby providing both confidentiality and integrity service that aligns with the CIA (Confidentiality, Integrity, Availability) triad for data security.

## Getting Started

To get started with this project, follow these steps:

### Prerequisites

- Go installed on your machine.
- A SQL database server (sqlite) to store IDs and hashes.

### Installation

1. Clone this repository to your local machine:

  ```bash
  git clone https://github.com/justsaumit/go-fis-api.git
  cd go-fis-api
  ```

2. Initialize and install project dependencies using Go modules:

  ```bash
  go mod vendor
  ```

3. Create a `.env` file at the root of the project and specify the environment variables. Default values will be used if not set.  
   For example (refer to env.example):

   ```dotenv
   ENVIRONMENT=production
   PORT=3000

   # SSL/TLS Configuration (if running in container, make it /certs/<cert/key>.pem and mount the certificate directory as a volume to /certs)
   CERTPATH=<path to>/fullchain.pem
   KEYPATH=<path to>/privkey.pem
   ```

   If the `.env` file is not set, default values such as `development` and `localhost` will be used where TLS is not required.

4. Run the server using one of the following methods:

   a. **Run with Go (Development Mode):**

      ```bash
      go run server.go
      ```
      This command starts the server using the Go runtime.

   b. **Build and Run:**

      ```bash
      CGO_ENABLED=1 go build -o main .
      ./main
      ```
      This sequence of commands builds the application and then runs the compiled binary `main`.

   c. **Using Docker:**

      - With the help of Dockerfile:
        ```bash
        docker build -t myapp .
      - Run the image
        - For `development` environment:
        ```bash
        docker run -p 3000:3000 myapp
          ```
        - For `production` environment with specified certificate and private key paths:
        ```bash
        docker run -p 3000:3000 -v /path/to/certifcates:/certs myapp
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
- [ ] Optimize Docker Image size (Currently 509MB)
## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
