import net from 'net';
import express from 'express';

const app = express();
const RATE_LIMIT = 10;
const TIME_WINDOW = 20; // Time window in seconds

app.set('view engine', 'ejs'); // Set up EJS for templating

// Connect to Redis (LiteDB) function
function connectLiteDB() {
    const client = new net.Socket();
    client.connect(6379, '127.0.0.1');
    return client;
}

async function rateLimit(req, res, next) {
    const client = connectLiteDB();
    const key = `rate_limit:${req.ip}`;
    console.log(`Handling request from IP: ${req.ip}`);

    // Check if the rate limit key exists and get its value
    client.write(`*2\r\n$3\r\nGET\r\n$${key.length}\r\n${key}\r\n`);

    client.once('data', (data) => {
        const response = data.toString().trim();

        if (response === "$-1") {
            // Key does not exist; initialize it with the first request timestamp and count of 1
            const now = Date.now();
            console.log(`New IP detected. Initializing rate limit for IP: ${req.ip}`);
            client.write(`*3\r\n$3\r\nSET\r\n$${key.length}\r\n${key}\r\n$${`${now}:1`.length}\r\n${now}:1\r\n`);
            client.write(`*3\r\n$4\r\nEXPIRE\r\n$${key.length}\r\n${key}\r\n$${TIME_WINDOW}\r\n`);
            client.once('data', () => {
                client.end();
                next();
            });
        } else {
            // Parse the value to get the timestamp and current count
            const value = response.split("\r\n")[1];
            const [timestampStr, countStr] = value.split(':');
            const timestamp = parseInt(timestampStr, 10);
            const currentCount = parseInt(countStr, 10);

            console.log(`IP: ${req.ip} - Current request count: ${currentCount}`);

            if (currentCount >= RATE_LIMIT) {
                // Calculate the remaining time based on the first request timestamp
                const elapsedTime = Math.floor((Date.now() - timestamp) / 1000);
                const remainingTime = Math.max(TIME_WINDOW - elapsedTime, 0);

                if (remainingTime > 0) {
                    // If time is left in the window, block the request and show remaining time
                    console.log(`IP: ${req.ip} - Rate limit exceeded. Time remaining: ${remainingTime} seconds`);
                    client.end();
                    return res.render('rate_limited', { remainingTime });
                } else {
                    // If time window has expired, reset the rate limit
                    console.log(`IP: ${req.ip} - Time window expired. Resetting rate limit.`);
                    const now = Date.now();
                    client.write(`*3\r\n$3\r\nSET\r\n$${key.length}\r\n${key}\r\n$${`${now}:1`.length}\r\n${now}:1\r\n`);
                    client.write(`*3\r\n$4\r\nEXPIRE\r\n$${key.length}\r\n${key}\r\n$${TIME_WINDOW}\r\n`);
                    client.once('data', () => {
                        client.end();
                        next();
                    });
                }
            } else {
                // Increment the count within the allowed window
                const updatedCount = currentCount + 1;
                console.log(`IP: ${req.ip} - Incrementing request count to: ${updatedCount}`);
                client.write(`*3\r\n$3\r\nSET\r\n$${key.length}\r\n${key}\r\n$${`${timestamp}:${updatedCount}`.length}\r\n${timestamp}:${updatedCount}\r\n`);
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
    res.render('index');
});

app.listen(3000, () => console.log('Server running on http://localhost:3000'));
