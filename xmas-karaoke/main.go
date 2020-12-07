package main

import (
	"bufio"
	"bytes"
	"encoding/binary"
	"fmt"
	"io"
	"log"
	"math/rand"
	"net"
	"os"
	"strings"
	"time"
	"unicode/utf16"

	"golang.org/x/text/encoding/unicode"
	"golang.org/x/text/encoding/unicode/utf32"
)

var vengaboys = []string{
	"If you're alone and you need a friend",
	"Someone to make you forget your problems",
	"Just come along baby",
	"Take my hand",
	"I'll be your lover tonight",
	"Whoa oh oh oh",
	"This is what I wanna do",
	"Let's have some fun",
	"What I want is me and you",
	"BOM BOM BOM BOM",
	"I want you in my room",
	"Let's spend the night together",
	"From now until forever",
	"BOM BOM BOM BOM",
	"I wanna go BOM BOM",
	"Let's spend the night together",
	"Together in my room",
	"Everybody get on down",
	"The Vengaboys are back in town",
	"This is what I wanna do",
	"Let's have some fun",
	"What I want is me and you",
	"BOM BOM BOM BOM",
	"I want you in my room",
	"Let's spend the night together",
	"From now until forever",
	"BOM BOM BOM BOM",
	"I wanna go BOM BOM",
	"Let's spend the night together",
}

func encodeMessageUTF16(message string, big bool) []byte {
	runeByte := []rune(message)
	encodedByte := utf16.Encode(runeByte)
	var buf []byte
	var ta []byte
	for _, num := range encodedByte {
		ta = make([]byte, 2)
		if big {
			binary.BigEndian.PutUint16(ta, num)
		} else {
			binary.LittleEndian.PutUint16(ta, num)
		}
		buf = append(buf, ta...)
	}
	return buf
}

func encodeMessageUTF32(message string, big bool) []byte {
	endian := unicode.LittleEndian
	if big {
		endian = unicode.BigEndian
	}
	e := utf32.UTF32(utf32.Endianness(endian), utf32.BOMPolicy(unicode.IgnoreBOM))
	encoder := e.NewEncoder()
	var outbuf bytes.Buffer
	writer := encoder.Writer(&outbuf)
	inReader := strings.NewReader(message)
	_, err := io.Copy(writer, inReader)
	if err != nil {
		log.Fatal(err)
	}
	return outbuf.Bytes()
}

var BOMList = []string{
	"EFBBBF",   // UTF8
	"FEFF",     // UTF16-BE
	"FFFE",     // UTF16-LE
	"0000FEFF", // UTF32-BE
	"FFFE00",   // UTF32-LE
}

func handleConnection(c net.Conn) {
	fmt.Printf("Serving %s\n", c.RemoteAddr().String())

	c.Write([]byte("Welcome to Santa's XMAS Karaoke! "))
	c.Write([]byte("Please choose your tune:\n[1] Vengaboys - BOM BOM BOM BOM\n[2] Shaggy - BOMbastic\n[3] BOMfunk MC's - Freestyler\nSong: "))

	chosenSong := 0
	tries := 0

	for {
		inData, err := bufio.NewReader(c).ReadString('\n')
		if err != nil {
			fmt.Println(err)
			return
		}

		temp := strings.TrimSpace(string(inData))
		switch temp {
		case "1":
			c.Write([]byte("You chose BOM BOM BOM BOM, let's go!\n\n"))
			chosenSong = 1
			break
		case "2":
			c.Write([]byte("BOMbastic is currently not available :(\n"))
			break
		case "3":
			c.Write([]byte("Freestyler is currently not available :(\n"))
			break
		default:
			c.Write([]byte("Unknown song\n"))
			tries++
		}

		if chosenSong != 0 {
			break
		}
		if tries > 2 {
			c.Write([]byte("You broke Santa's jukebox!\n"))
			c.Close()
			return
		}
	}

	var random int

	for _, v := range vengaboys {
		rand.Seed(time.Now().UnixNano())
		random = rand.Intn(5)
		//fmt.Printf("Random: %d, BOM sent: %s\n", random, BOMList[random])

		c.Write([]byte(v + "\n"))
		c.Write([]byte(BOMList[random] + ": "))

		inData, err := bufio.NewReader(c).ReadBytes('\n')
		if err != nil {
			fmt.Println(err)
			return
		}

		inData = bytes.TrimSpace(inData)
		//fmt.Printf("Input data: %s (%x)\n", inData, inData)

		switch random {
		case 0:
			// UTF8
			temp := strings.TrimSpace(string(inData))
			if temp != v {
				c.Write([]byte("Your voice encoding is wrong.\n"))
				//fmt.Printf("Invalid input for UTF8: %s != %s\n", temp, v)
				c.Close()
				return
			}
		case 1:
			// UTF16-BE
			encoded := encodeMessageUTF16(v, true)
			//fmt.Printf("UTF16-LE encoded: %x\n", encoded)

			if bytes.Compare(inData, encoded) != 0 {
				//fmt.Printf("Invalid input for UTF16-BE:\n\t%s (%x)\n\t%s (%x)\n", inData, inData, encoded, encoded)
				c.Write([]byte("Your voice encoding is wrong.\n"))
				c.Close()
				return
			}
		case 2:
			// UTF16-LE
			encoded := encodeMessageUTF16(v, false)
			//fmt.Printf("UTF16-BE encoded: %x\n", encoded)

			if bytes.Compare(inData, encoded) != 0 {
				//fmt.Printf("Invalid input for UTF16-LE:\n\t%s (%x)\n\t%s (%x)\n", inData, inData, encoded, encoded)
				c.Write([]byte("Your voice encoding is wrong.\n"))
				c.Close()
				return
			}
		case 3:
			// UTF32-BE
			encoded := encodeMessageUTF32(v, true)
			//fmt.Printf("UTF32-BE encoded: %x\n", encoded)

			if bytes.Compare(inData, encoded) != 0 {
				fmt.Printf("Invalid input for UTF32-BE:\n\t%s (%x)\n\t%s (%x)\n", inData, inData, encoded, encoded)
				c.Write([]byte("Your voice encoding is wrong.\n"))
				c.Close()
				return
			}
		case 4:
			// UTF32-LE
			encoded := encodeMessageUTF32(v, false)
			//fmt.Printf("UTF32-LE encoded: %x\n", encoded)

			if bytes.Compare(inData, encoded) != 0 {
				fmt.Printf("Invalid input for UTF32-LE:\n\t%s (%x)\n\t%s (%x)\n", inData, inData, encoded, encoded)
				c.Write([]byte("Your voice encoding is wrong.\n"))
				c.Close()
				return
			}
		}

	}

	c.Write([]byte(fmt.Sprintf("Wow, you are an amazing singer! Take this prize: %s\n", os.Getenv("XMAS_SECRET"))))
	fmt.Printf("Client got solution: %s\n", c.RemoteAddr().String())
	c.Close()
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Please provide a port number!")
		return
	}
	PORT := ":" + os.Args[1]
	l, err := net.Listen("tcp4", PORT)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer l.Close()

	for {
		c, err := l.Accept()
		if err != nil {
			fmt.Println(err)
			return
		}
		go handleConnection(c)
	}
}
