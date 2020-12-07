package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io"
	"log"
	"net"
	"strings"
	"time"
	"unicode/utf16"

	"golang.org/x/text/encoding/unicode"
	"golang.org/x/text/encoding/unicode/utf32"
)

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
func main() {
	addr, err := net.ResolveTCPAddr("tcp", "localhost:11")
	if err != nil {
		log.Fatal(err)
	}
	con, err := net.DialTCP("tcp", nil, addr)
	if err != nil {
		log.Fatal(err)
	}

	tmp := make([]byte, 256)
	for {
		_, err := con.Read(tmp)
		if err != nil {
			log.Fatal(err)
		}
		if strings.Contains(string(tmp), "Song: ") {
			break
		}
	}

	_, err = con.Write([]byte("1\n"))
	if err != nil {
		log.Fatal(err)
	}

	time.Sleep(300 * time.Millisecond)

	var initialLine, line, initialBOM, bom string
	tmp = make([]byte, 256)
	for {
		_, err := con.Read(tmp)
		if err != nil {
			log.Fatal(err)
		}
		if strings.Contains(string(tmp), "let's go!\n\n") {
			tmp = bytes.TrimSpace(tmp)
			tmp = bytes.Trim(tmp, "\x00")
			parts := strings.Split(string(tmp), "\n\n")
			others := strings.Split(parts[1], "\n")
			initialLine = strings.TrimSpace(others[0])
			initialBOM = strings.ReplaceAll(others[1], ":", "")
			initialBOM = strings.TrimSpace(initialBOM)
			break
		}
	}

	for {
		if line == "" && bom == "" {
			line = initialLine
			bom = initialBOM
		}

		fmt.Printf("---\nLine: %s\nBOM: %s (%x)\n---\n", line, bom, bom)

		switch bom {
		case "EFBBBF":
			fmt.Printf("Processing UTF8\n")
			con.Write([]byte(line))
			con.Write([]byte("\n"))
		case "FFFE":
			fmt.Printf("Processing UTF16-LE\n")
			bla := encodeMessageUTF16(line, false)
			con.Write(bla)
			con.Write([]byte("\n"))
		case "FEFF":
			fmt.Printf("Processing UTF16-BE\n")
			bla := encodeMessageUTF16(line, true)
			con.Write(bla)
			con.Write([]byte("\n"))
		case "0000FEFF":
			fmt.Printf("Processing UTF32-BE\n")
			bla := encodeMessageUTF32(line, true)
			con.Write(bla)
			con.Write([]byte("\n"))
		case "FFFE00":
			fmt.Printf("Processing UTF32-LE\n")
			bla := encodeMessageUTF32(line, false)
			con.Write(bla)
			con.Write([]byte("\n"))
		}

		time.Sleep(500 * time.Millisecond)
		fmt.Println("Reading response...")

		resp := make([]byte, 1024)
		_, err = con.Read(resp)
		if err != nil {
			log.Fatal(err)
		}

		line = ""
		bom = ""

		resp = bytes.Trim(resp, "\x00")
		fmt.Printf("Response: %s\n", resp)
		if strings.Contains(string(resp), "encoding is wrong") {
			fmt.Println("Our encoding was wrong")
			return
		}
		parts := strings.Split(string(resp), "\n")
		line = strings.TrimSpace(parts[0])
		bom = strings.ReplaceAll(parts[1], ":", "")
		bom = strings.TrimSpace(bom)

		if bom == "" {
			fmt.Printf("BOM is empty?\n")
			more := make([]byte, 1024)
			_, err = con.Read(more)
			if err != nil {
				log.Fatal(err)
			}
			fmt.Printf("More data: %s\n", more)

		} else {
			fmt.Printf("Using BOM: %s (%x)\n", bom, bom)
		}

	}

}
