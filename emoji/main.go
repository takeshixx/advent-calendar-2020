package main

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"math/rand"
	"net"
	"os"
	"strings"
	"time"

	"github.com/rivo/uniseg"
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

	e.printLn("❓❓❓❓❓❓❓❓")
	// Randomly generate correct answer:
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
	ab := strings.Builder{}
	for i := 0; i < 8; i++ {
		ab.WriteString(emojis[rand.Intn(len(emojis))])
	}
	answerGr := uniseg.NewGraphemes(ab.String())

	for {
		var input string
		input, success = e.readString()
		if !success {
			break
		}

		inputGr := uniseg.NewGraphemes(input)
		answerGr.Reset()
		correctAnswer := true
		output := strings.Builder{}
		for answerGr.Next() {
			hasNext := inputGr.Next()
			if hasNext && bytes.Equal(answerGr.Bytes(), inputGr.Bytes()) {
				output.WriteString(answerGr.Str())
			} else {
				output.WriteString("🚫")
				correctAnswer = false
			}
		}

		if correctAnswer {
			break
		}

		success = e.printLn(output.String())
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
