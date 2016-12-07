// Ingress Decoder

package main

import (
	"flag"
	"fmt"
	"strings"
)

var (
	input     string
	method    string
	morseCode map[string]string
)

func init() {
	flag.StringVar(&input, "i", "", "You must specify a code")
	flag.StringVar(&method, "m", "", "atbash/hexatbash/rot13/frommorse/tomorse/swapmorse/fence. default fence")
	flag.Parse()

	morseCode = make(map[string]string)
	morseCode["A"] = ".-"
	morseCode["B"] = "-..."
	morseCode["C"] = "-.-."
	morseCode["D"] = "-.."
	morseCode["E"] = "."
	morseCode["F"] = "..-."
	morseCode["G"] = "--."
	morseCode["H"] = "...."
	morseCode["I"] = ".."
	morseCode["J"] = ".---"
	morseCode["K"] = "-.-"
	morseCode["L"] = ".-.."
	morseCode["M"] = "--"
	morseCode["N"] = "-."
	morseCode["O"] = "---"
	morseCode["P"] = ".--."
	morseCode["Q"] = "--.-"
	morseCode["R"] = ".-."
	morseCode["S"] = "..."
	morseCode["T"] = "-"
	morseCode["U"] = "..-"
	morseCode["V"] = "...-"
	morseCode["W"] = ".--"
	morseCode["X"] = "-..-"
	morseCode["Y"] = "-.--"
	morseCode["Z"] = "--.."
	morseCode["0"] = "-----"
	morseCode["1"] = ".----"
	morseCode["2"] = "..---"
	morseCode["3"] = "...--"
	morseCode["4"] = "....-"
	morseCode["5"] = "....."
	morseCode["6"] = "-...."
	morseCode["7"] = "--..."
	morseCode["8"] = "---.."
	morseCode["9"] = "----."

}

func keyboardMirror(d string) string {
	var keyboard = []string{
		"!@#$%^&*()",
		"1234567890",
		"QWERTYUIOP",
		"qwertyuiop",
		"ASDFGHJKL:",
		"asdfghjkl;",
		"ZXCVBNM<>?",
		"zxcvbnm,./",
	}
	var bs string
	for _, s := range d {
		for _, v1 := range keyboard {
			for i, v2 := range v1 {
				if s == v2 {
					bs += string(v1[9-i])
				}
			}
		}
	}
	return bs
}

func fromMorse(d string) string {
	var ds = strings.Split(d, "/")
	var bs string
	for _, s := range ds {
		for k, v := range morseCode {
			if v == s {
				bs += k
			}
		}
	}
	return bs
}

func toMorse(d string) string {
	var bs string
	for _, s := range d {
		bs = fmt.Sprintf("%s/%s", bs, morseCode[string(s)])
	}
	return bs
}

func swapMorse(d string) string {
	bs := strings.Replace(d, ".", ",", -1)
	bs = strings.Replace(bs, "-", ".", -1)
	bs = strings.Replace(bs, ",", "-", -1)
	return bs
}

func atbash(d string) string {
	var b = []byte(d)
	for i, v := range b {
		if v >= 49 && v <= 57 {
			b[i] = (49 + 57) - v
		} else if v >= 65 && v <= 90 {
			b[i] = (65 + 90) - v
		} else if v >= 97 && v <= 122 {
			b[i] = (97 + 122) - v
		}
	}
	return string(b)
}

func rot13(d string) string {
	var b = []byte(d)
	for i, v := range b {
		if v >= 97 && v <= 122 {
			if 122-v >= 13 {
				b[i] = v + 13
			} else {
				b[i] = 13 + v - 122 + 97 - 1
			}
		} else if v >= 65 && v <= 90 {
			if 90-v >= 13 {
				b[i] = v + 13
			} else {
				b[i] = 13 + v - 90 + 65 - 1
			}
		}
	}
	return string(b)
}

func hexAtbash(d string) string {
	var b = []byte(strings.ToLower(d))
	for i, v := range b {
		if v >= 97 && v <= 102 {
			b[i] = (97 + 102) - v
		} else if v >= 48 && v <= 57 {
			b[i] = (48 + 57) - v
		}
	}
	return string(b)
}

func fence(d string) {
	var l = len(d)
	fmt.Println("lenth =", l)
	for i := 2; i < l; i++ {
		if l%i == 0 {
			fmt.Printf("try %dx%d:\n", i, l/i)
			for j := 0; j < l; j += i {
				fmt.Println(d[j : j+i])
			}
			fmt.Println("")
		}
	}
}

func main() {
	switch strings.ToLower(method) {
	case "atbash":
		fmt.Println(atbash(input))
	case "hexatbash":
		fmt.Println(hexAtbash(input))
	case "rot13":
		fmt.Println(rot13(input))
	case "frommorse":
		fmt.Println(fromMorse(input))
	case "tomorse":
		fmt.Println(toMorse(input))
	case "swapmorse":
		fmt.Println(
			fromMorse(
				swapMorse(
					toMorse(input),
				),
			),
		)
	case "keymirror":
		fmt.Println(keyboardMirror(input))
	case "fence":
		fallthrough
	default:
		fence(input)
	}
}
