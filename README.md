# LiteDB Rate Limiter

LiteDB is a simple rate-limiting server built with Node.js, Express, and EJS. It limits client requests to a specified threshold within a given time window, using LiteDB for efficient storage and retrieval of request counts and timestamps. This project demonstrates basic rate-limiting techniques and dynamic EJS rendering for notifications and status updates.

## Features

- **Rate Limiting**: Restricts clients to a maximum number of requests within a defined time window.
- **LiteDB Integration**: Uses LiteDB as a backend for fast, temporary storage of client request data.
- **Frontend Feedback**: Provides a dynamic EJS frontend that displays a "Too many requests" message and a countdown for when clients can make new requests.

## Getting Started

### Prerequisites

- **Node.js** (v18 or later recommended)
- **LiteDB** server running locally (default port: 6379)
- **Go** installed for running the LiteDB server

### Installation

1. Clone the repository:
   ```bash
   git clone https://github.com/busybrowsensei1/LiteDB.git
2. Navigate into the project directory:
   ```bash
   cd LiteDB
3. Install the dependencies for the Node.js server:
   ```bash
   npm install

### Running the LiteDB Server (server.go)

LiteDB is built with Go, so to run the server, you need to start the LiteDB server separately:

1. Ensure you have Go installed.
2. Navigate to the directory containing server.go.
3. Run the Go server:

   ```bash
   go run server.go

This will start the LiteDB server. The Node.js app communicates with this server to handle rate limiting.

### Running the Rate Limiter Server (index.js)

After the LiteDB server is running, you can start the rate-limiting server by running:

   ```bash
   node index.js

This will start the server on http://localhost:3000.
   
   
