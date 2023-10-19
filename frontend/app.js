// app.js

const express = require('express');
const axios = require('axios');
const app = express();
const PORT = process.env.PORT || 3000;

app.use(express.static('public'));

// Define the API URL
const API_URL = 'https://api.howtofixcve.com';

// Set up routes
app.get('/', (req, res) => {
    res.sendFile(__dirname + '/public/index.html');
});

// Route to search for a CVE by ID
app.get('/search/:cveId', async (req, res) => {
    try {
        const cveId = req.params.cveId;
        const response = await axios.get(`${API_URL}/info/${cveId}`);
        const data = response.data;
        res.status(200).json(data);
    } catch (error) {
        console.error(error);
        res.status(500).json({ error: 'Internal server error' });
    }
});

// Route to get the latest CVE
app.get('/latest', async (req, res) => {
    try {
        const response = await axios.get(`${API_URL}/latest`);
        const data = response.data;
        res.json(data);

    } catch (error) {
        console.error(error);
        res.status(500).json({ error: 'Internal server error' });
    }
});

// Start the server
app.listen(PORT, () => {
    console.log(`Server is running on port ${PORT}`);
});

