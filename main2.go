package main

import (
	"fmt"
	"unicode/utf8"
)

var words = []string{"in", "input", "output", "out", "puton", "one"}

type flag struct {
	word  string
	start int
	end   int
}

func (f *flag) endp() {
	(*f).end = (*f).start + utf8.RuneCountInString((*f).word)
}

var flags []flag

func compare(substr []rune, start, end int) bool {
	var count, fcount int
	var res bool
	var f flag
	fmt.Printf("%c \n", substr)
	for _, word := range words {

		temp := []rune(word)
		i := 0
	Flag:
		for i = 0; i < len(temp); {
			if len(substr) < len(temp) {
				break Flag
			}
			if substr[i] == temp[i] {
				fmt.Printf("Sub= %c word= %c i= %d count= %d, W= %v \n", substr[i], temp[i], i, count, word)
				i = i + 1
				count = count + 1
			} else {
				fmt.Printf(" !!! Sub= %c word= %c i= %d count= %d, W= %v \n", substr[i], temp[i], i, count, word)
				break Flag
			}
		}
		if count == len(temp) {
			f = flag{word, start, end}
			f.endp()
			flags = append(flags, f)
			count = 0
			fcount = fcount + 1
		} else {
			count = 0
		}

	}

	if fcount > 0 {
		res = true
	} else {
		res = false
	}

	fmt.Println(res)
	return res
}

func main() {
	//var i int
	var start, end int = 0, 6
	str := []rune("oneputonininputoutoutput")
	// fmt.Println(flags)
	for {
		if len(str) == 0 {
			break
		}
		compare(str[start:end], start, end)
		//compare(str[10:15], 10, 15)
		fmt.Println(flags)
		// start = end
		// end = end + 6
		compare(str[start:end], start, end)
		fmt.Println(flags)
		break
	}

}
