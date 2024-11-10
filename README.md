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
