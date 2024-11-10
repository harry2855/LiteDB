
# LiteDB Project

## Overview

LiteDB is a lightweight, file-based NoSQL database created for easy storage and retrieval of structured data. The project is designed to be simple and flexible, with three primary components that demonstrate LiteDB's capabilities in various contexts:

1. **Database** - A standalone database engine written in Go, supporting basic CRUD operations and custom commands.
2. **Rate Limiter** - A middleware implemented in Node.js that uses LiteDB to track request counts and limit user access to resources within specified time windows.
3. **Chat App** - A basic real-time chat application that uses LiteDB as a backend to store and retrieve chat messages.

This project is ideal for learning about file-based databases, integrating custom databases with applications, and using Go and Node.js to build modular components. 

## Features

### Database
- **File-based storage** - Stores data in a lightweight file format, making it easy to deploy and operate without complex setup.
- **Custom commands** - Supports commands beyond typical CRUD operations, allowing for flexible data interactions (e.g., incrementing values, transactions).
- **Backup and restore** - Provides commands to save and load data, supporting data persistence and migration.
  
### Rate Limiter
- **Request rate limiting** - Limits the frequency of requests to prevent abuse, with customizable rates and time windows.
- **Easy integration** - Simple setup allows you to add rate limiting to any Express-based app.
- **Real-time feedback** - Displays information about remaining requests and wait times.

### Chat App
- **Real-time messaging** - Supports real-time chat with data storage in LiteDB for easy message retrieval.
- **Lightweight backend** - Uses LiteDB as the backend database, providing a simple setup without the need for complex database infrastructure.
- **Scalable** - Designed to handle small-scale chat functionality, ideal for testing and learning.

---

## Tech Stack

- **Go** - Database engine and command handling.
- **Node.js** - Backend framework for rate limiter and chat application.
- **Express** - Web server framework used in the rate limiter and chat app.
- **Redis** - CLI for testing and interacting with the rate limiter.

---

## Prerequisites

To run this project, you need the following software installed:

1. **Go** - [Download and install Go](https://go.dev/doc/install)
2. **Node.js** - [Download and install Node.js](https://nodejs.org/)
3. **Redis CLI** - [Install Redis CLI](https://redis.io/docs/latest/operate/oss_and_stack/install/install-redis/)

---

## Command Syntax and Function

*Note: All commands should be written in **uppercase**.*

1. **CONFIGGET** - Retrieves the database configuration settings.
2. **DELETE** - Deletes a specified key-value pair from the database.
3. **ECHO** - Outputs a provided message to verify database connectivity.
4. **INCR** - Increments a numeric value by one; initializes to 0 if the key does not exist.
5. **KEYS** - Lists all keys currently stored in the database.
6. **LIST** - Retrieves all items stored in a specified list.
7. **LOAD** - Loads data from a backup file into the database.
8. **MULTI** - Starts a transaction block for atomic command execution.
9. **SAVE** - Saves the current database state to a file for persistence.
10. **SETGET** - Sets a value for a specified key and retrieves it immediately.

---

## Setup Instructions

### Database Setup

1. **Run the LiteDB Server**  
   - Navigate to the `database/app` directory:
     ```bash
     cd database/app
     ```
   - Start the server:
     ```bash
     go run server.go
     ```
   - This starts the LiteDB server, which will handle requests from the rate limiter and chat app.

2. **Install Redis CLI**  
   - Follow the instructions on [Redis installation page](https://redis.io/docs/latest/operate/oss_and_stack/install/install-redis/) to install the Redis CLI for testing and managing rate limits.

### Rate Limiter

The rate limiter middleware limits the number of requests allowed within a time window. This helps prevent abuse and control server load. The rate and time window are adjustable based on application needs.

#### Prerequisites
- Ensure the LiteDB server is running as outlined above.

#### Installation and Setup
1. Clone the repository.
2. Navigate to the rate limiter directory:
   ```bash
   cd LiteDB/rate-limiter
   ```
3. Install dependencies:
   ```bash
   npm install
   ```
4. Start the server:
   ```bash
   node index.js
   ```

5. Access the rate limiter at `http://localhost:3000`. You can adjust the request rate and time window in the configuration.

---

### Chat App

The chat app enables real-time messaging and uses LiteDB as the backend storage for messages. 

#### Prerequisites
- Ensure the LiteDB server is running.

#### Installation and Setup
1. Clone the repository if not already done.
2. Navigate to the chat app directory:
   ```bash
   cd LiteDB/chatbox
   ```
3. Install dependencies:
   ```bash
   npm install
   ```
4. Start the development server:
   ```bash
   npm run dev
   ```

The chat app should now be accessible at `http://localhost:3000`.

---
