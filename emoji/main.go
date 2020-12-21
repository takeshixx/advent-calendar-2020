package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"net"
	"os"
	"strings"
	"time"
)

const port = 8080

var token = "PLEASE PASS THE TOKEN"

type emojiConnectionHandler struct {
	c net.Conn
}

func (e *emojiConnectionHandler) printLn(line string) bool {
	_, err := e.c.Write([]byte(line + "\n"))
	return err == nil
}

func (e *emojiConnectionHandler) readString() (line string, success bool) {
	readLine, err := bufio.NewReader(e.c).ReadString('\n')
	readLine = strings.TrimSpace(string(readLine))
	return readLine, err == nil
}

func (e *emojiConnectionHandler) handleConnection() {
	fmt.Printf("Serving %s\n", e.c.RemoteAddr().String())

	e.quiz()

	success := e.welcome() &&
		e.scissorsStonePaper() &&
		e.quiz()

	if success {
		e.printToken()
	}

	e.c.Close()
}

func (e *emojiConnectionHandler) welcome() bool {
	success := true
	for {
		answers := []string{
			"🎅",
			"🤶",
		}
		question := answers[rand.Intn(len(answers))]

		success = e.printLn(question)
		if !success {
			break
		}
		var input string
		input, success = e.readString()
		if !success {
			break
		}

		if input == question {
			break
		}
	}

	e.printLn("👍\n")
	return success
}

func (e *emojiConnectionHandler) scissorsStonePaper() bool {
	success := true
	for {
		answers := []string{
			"✂️",
			"🪨",
			"📃",
		}
		question := answers[rand.Intn(len(answers))]

		success = e.printLn(question)
		if !success {
			break
		}
		var input string
		input, success = e.readString()
		if !success {
			break
		}

		if input == "✂️" && question == "📃" ||
			input == "🪨" && question == "✂️" ||
			input == "📃" && question == "🪨" {
			break
		}
	}

	e.printLn("👍\n")
	return success
}

func (e *emojiConnectionHandler) quiz() bool {
	success := true
	emojis := []string{
		"🎅",
		"🎄",
		"☃️",
		"🎁",
		"❄️",
		"⛄",
		"🎍",
		"🛷",
		"🦌",
		"🤶",
	}
	var answer [8]string
	for i := 0; i < len(answer); i++ {
		answer[i] = emojis[rand.Intn(len(emojis))]
	}

	for {
		var input string
		input, success = e.readString()
		if !success {
			break
		}

		var check [len(answer)]string
		correctAnswer := true
		for i := 0; i < len(answer); i++ {
			// This code is wrong atm, see https://stackoverflow.com/a/12668840/2073799
			e.printLn(fmt.Sprintf("%d", len(answer[i])))
			if len(input) > i && string(input[i]) == answer[i] {
				check[i] = answer[i]
			} else {
				check[i] = "🚫"
				correctAnswer = false
			}
		}
		if correctAnswer {
			break
		}

		success = e.printLn(strings.Join(check[:], ""))
		if !success {
			break
		}
	}

	e.printLn("👍\n")
	return success
}

func (e *emojiConnectionHandler) printToken() {
	log.Printf("%v solved the challenge", e.c.RemoteAddr())
	e.printLn("🪙" + token + "🪙")
}

func main() {
	rand.Seed(time.Now().UnixNano())

	if len(os.Args) != 2 {
		log.Fatalf("Please pass the token as the first argument.")
	}
	token = os.Args[1]

	port := 8080
	log.Printf("Listening on :%d...\n", port)
	l, err := net.Listen("tcp", fmt.Sprintf(":%v", port))
	if err != nil {
		log.Fatal(err)
	}
	defer l.Close()

	for {
		c, err := l.Accept()

		if err != nil {
			log.Fatal(err)
		}

		e := emojiConnectionHandler{
			c,
		}
		go e.handleConnection()
	}
}
