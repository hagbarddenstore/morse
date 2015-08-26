package main

import (
	"bufio"
	"bytes"
	"flag"
	"fmt"
	"os"
	"strings"
)

const (
	letterSpace = "   "
	wordSpace   = "     "
)

var (
	shouldDecode bool
	morseToAscii = make(map[string]string, 0)
	asciiToMorse = map[string]string{
		"A":  "- ---",
		"B":  "--- - - -",
		"C":  "--- - --- -",
		"D":  "--- - -",
		"E":  "-",
		"F":  "- - --- -",
		"G":  "--- --- -",
		"H":  "- - - -",
		"I":  "- -",
		"J":  "- --- --- ---",
		"K":  "--- - ---",
		"L":  "- --- - -",
		"M":  "--- ---",
		"N":  "--- -",
		"O":  "--- --- ---",
		"P":  "- --- --- -",
		"Q":  "--- --- - ---",
		"R":  "- --- -",
		"S":  "- - -",
		"T":  "---",
		"U":  "- - ---",
		"V":  "- - - ---",
		"W":  "- --- ---",
		"X":  "--- - - ---",
		"Y":  "--- - --- ---",
		"Z":  "--- --- - -",
		"1":  "- --- --- --- ---",
		"2":  "- - --- --- ---",
		"3":  "- - - --- ---",
		"4":  "- - - - ---",
		"5":  "- - - - -",
		"6":  "--- - - - -",
		"7":  "--- --- - - -",
		"8":  "--- --- --- - -",
		"9":  "--- --- --- --- -",
		"0":  "--- --- --- --- ---",
		"?":  "- - -- --- - -",
		"!":  "- - --- --- -",
		",":  "--- --- - - --- ---",
		".":  "- --- - --- - ---",
		"-":  "--- - - - - ---",
		"(":  "--- - --- --- -",
		")":  "--- - --- --- - ---",
		"@":  "- --- --- - --- -",
		"/":  "--- - - --- -",
		"%":  "- --- --- - -",
		"\"": "- --- - - --- -",
		"'":  "- --- --- --- --- -",
		";":  "--- - --- - --- -",
		":":  "--- --- --- - - -",
	}
)

func init() {
	flag.BoolVar(&shouldDecode, "decode", false, "")

	for key, value := range asciiToMorse {
		morseToAscii[value] = key
	}
}

func main() {
	flag.Parse()

	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		text := scanner.Text()

		if shouldDecode {
			fmt.Println(decode(text))
		} else {
			fmt.Println(encode(text))
		}
	}
}

func decode(text string) string {
	var buffer bytes.Buffer

	words := strings.Split(text, wordSpace)

	for i, word := range words {
		for _, morse := range strings.Split(word, letterSpace) {
			if letter, ok := morseToAscii[morse]; ok {
				buffer.WriteString(letter)
			}
		}

		if i != len(words)-1 {
			buffer.WriteString(" ")
		}
	}
	return buffer.String()
}

func encode(text string) string {
	var buffer bytes.Buffer
	writeSpace := true

	for i, r := range text {
		letter := strings.ToUpper(string(r))

		if morse, ok := asciiToMorse[letter]; ok {
			buffer.WriteString(morse)

			writeSpace = true
		}

		if letter == " " {
			buffer.WriteString(wordSpace)
			writeSpace = false
		}

		if i == len(text)-1 {
			writeSpace = false
		}

		if writeSpace {
			buffer.WriteString(letterSpace)
		}
	}

	return buffer.String()
}
