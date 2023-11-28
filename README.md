# Golang Backend API for FIS (File Integrity Surveillance) Application
Developing a simple Golang backend API with the [Echo framework](https://github.com/labstack/echo) for FIS(File Integrity Surveillance) application which can be found [here](https://github.com/ayato91/Fair-Files).  
This API stores IDs and their corresponding hashes in a SQL server and provides functionality to verify if a given hash matches the stored hash for a specific ID. All the communication between the Application and API is secured using TLS encryption(HTTPS).  
Thereby providing both confidentiality and integrity service that aligns with the CIA (Confidentiality, Integrity, Availability) triad for data security.

## Getting Started

To get started with this project, follow these steps:

### Prerequisites

- Go installed on your machine.
- A SQL database server(sqlite) to store IDs and hashes.

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

3. Create a configuration file (e.g., config.yaml) to specify your database settings and other configurations.

4. Run the server:

  ```bash
  go run server.go
  ```

### Usage

Once the server is running, you can access the API endpoints to add file hashes and verify them.

- To add a file hash, make a POST request to `/upload` by uploading the file using multipart form, an generated ID and hash would be returned.
- To verify a file hash, make a POST request to `/verify` with previously generated ID and the file, a JSON message with Hashes matching or not will be recieved as response .

### Usage

Once the server is operational, you can interact with the API endpoints to add and verify file hashes.

- **Adding a File Hash:**
  - To add a file hash, send a POST request to the `/upload` endpoint. You'll need to upload the file using multipart/form-data. The server will generate an ID and a hash for the uploaded file, which will be returned in the response.

- **Verifying a File Hash:**
  - To verify a file hash, send a POST request to the `/verify` endpoint. Include the previously generated ID as a form field and upload the file using multipart/form-data. The server will respond with a JSON message indicating whether the hash of the uploaded file matches the stored hash for the given ID.

## To-Do-List
- [x] Handle Uploaded files (API)
- [x] Perform short ID Generation (API)
- [x] Perform Hashing (API)
- [x] Connect with DB
- [x] Store JSON data to DB
- [x] Perform verification

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
