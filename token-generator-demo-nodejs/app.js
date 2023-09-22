const express = require('express');
const jwt = require('jsonwebtoken');
const fs = require('fs');

const app = express();
const port = 3000;

// Read private key for signing
const privateKey = fs.readFileSync('private.pem', 'utf8');

app.get('/get-token', (req, res) => {
    const now = Math.floor(Date.now() / 1000); // Current time in seconds
    const payload = {
        sub: "1234567890",           // Replace with your user ID 
        iat: now,                    // Issued at
        nbf: now,                    // Not Before
        exp: now + (60 * 60 * 6)     // Expires in 6 hours
    };
    
    const signOptions = {
        algorithm: "RS256"
    };

    const token = jwt.sign(payload, privateKey, signOptions);

    res.json({ token });
});

app.listen(port, () => {
    console.log(`Get token at http://localhost:${port}/get-token`);
});