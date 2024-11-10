import net from 'net';
import express from 'express';

const app = express();
const RATE_LIMIT = 10; // Set your rate limit here
const TIME_WINDOW = 20; // Rate limit time window in seconds

// Helper function to connect to LiteDB
function connectLiteDB() {
    const client = new net.Socket();
    client.connect(6379, '127.0.0.1');
    return client;
}

async function rateLimit(req, res, next) {
    const client = connectLiteDB();
    const key = `rate_limit:${req.ip}`;

    // Check if the key exists with a GET command
    client.write(`*2\r\n$3\r\nGET\r\n$${key.length}\r\n${key}\r\n`);
    
    client.once('data', (data) => {
        const response = data.toString().trim();
        console.log(response)

        // If the key does not exist, initialize it
        if (response === "$-1") {
            // Set the key with an initial value of 1 and set an expiry
            client.write(`*5\r\n$3\r\nSET\r\n$${key.length}\r\n${key}\r\n$1\r\n1\r\n$2\r\npx\r\n$${TIME_WINDOW.toString().length+3}\r\n${TIME_WINDOW*1000}\r\n`);
            console.log(`Updated request count for ${req.ip}: ${1}`);
            client.once('data', () => {
                client.end();
                next(); // Allow the request since it's the first time for this IP
            });
        } else {
            // Parse the current count from the response
            const currentCount = parseInt(response.split("\r\n")[1], 10);
            console.log(`Current request count for ${req.ip}: ${currentCount}`);

            if (currentCount >= RATE_LIMIT) {
                client.end();
                return res.status(429).json({ message: "Too many requests, try again later." });
            } else {
                // Increment the key since it's within the limit
                client.write(`*2\r\n$4\r\nINCR\r\n$${key.length}\r\n${key}\r\n`);
                console.log(`Updated request count for ${req.ip}: ${currentCount+1}`);
                client.once('data', () => {
                    client.end();
                    next();
                });
            }
        }
    });
}

app.use(rateLimit);

app.get('/', (req, res) => {
    res.send('Welcome! You are within the rate limit.');
});

app.listen(3000, () => console.log('Server running on http://localhost:3000'));
