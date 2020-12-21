const express = require('express');
var bodyParser = require('body-parser');
const { exit } = require('process');
const { fork } = require('child_process');
const app = express();
const port = 3000;

if(process.argv.length != 4) {
    console.log(`Usage: ${process.argv[0]} ${process.argv[1]} ANSWER FLAG`);
    exit(1);
}
const answer = process.argv[2];
const flag = process.argv[3];

app.use(bodyParser.json());

app.post('/stopPollution', (req, res) => {
    res.setHeader('Content-Type', 'application/json');
    const stopPollution = fork('stopPollution.js');
    stopPollution.on('close', (code) => {
        console.log(`child process exited with code ${code}`);
    });
    input = req.body;
    input.correctAnswer = answer;
    input.flag = flag;
    stopPollution.send(input);
    stopPollution.on('message', result => {
        res.end(JSON.stringify(result, null, 2));
    });
});

app.use(express.static('.'));

app.listen(port, () => {
  console.log(`santas-christmas-factory listening at http://localhost:${port}`)
});

