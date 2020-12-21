const { merge } = require('lodash');

function checkAnswer(input, result) {
    if(input.answer === input.correctAnswer) {
        result.success = true;
    }
}

function addFlag(input, result) {
    if(result.success) {
        result.flag = input.flag;
    } else {
        result.hint = "Nope 👎 Try again!";
    }
}

process.on('message', (msg) => {
    try {
        let input = {
            answer: "",
            correctAnswer: "exampleAnswer",
            flag: "exampleFlag"
        };
        merge(input, msg);
        input = msg;
        let result = {};
        
        checkAnswer(input, result);
        addFlag(input, result);
        process.send(result, () => process.exit(0));
    } catch(e) {
        process.send({hint: "Exception occurred: "+ e});
        process.exit(1);
    }
});