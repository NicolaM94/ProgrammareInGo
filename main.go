package main

import (
	"fmt"
	"math/rand"
	"strings"
	"os"
	"sysinfo"
)

func maxOfTwo (a int, b int) int {
	if a > b {
		return a
	}
	if b > a {
		return b
	}
	return 0
}

func maxOfThree (a,b,c int) int {


	if a > b && a > c {
		return a 
	}
	if b > a && b > c {
		return b
	}
	if c > a && c > b {
		return c
	}
	return 0
}

func maxOfAll (listOfNumbers []int) int {
	var partialResult int
	for _,value := range listOfNumbers {
		if value > partialResult {
			partialResult = value
		}
	}
	return partialResult
}

func areUAVowel (letter string) bool {
	for _,c := range("aeiou") {
		if string(c) == letter {
			return true
		}
	}
	return false
}

func sommatriceInarrestabile (numeri []int) int {
	var result int 
	for _,value := range numeri {
		result += value
	}
	return result
}

func moltiplicatriceInarrestabile (numeri []int) int {
	var result int = 1
	for _,value := range numeri {
		result = result * value
	}
	return result
}

func reverser (word string) string {
	var newWord string
	for i := range word {		newWord += string(word[len(word)-i-1])
	}
	return newWord
}

func palindromoOrNot (word string) bool {
	var newWord string 
	for i := range word {
		newWord += string(word[len(word)-i-1])
	}
	return newWord == word
}

func personalizedLen (word string) int {
	var counter int 
	for range word {
		counter += 1
	}
	return counter
}

func istogramma (arrayOfIntegers []int) {
	for _,value := range( arrayOfIntegers) {
		fmt.Println(strings.Repeat("*",value))
	}
}

func everybodysOwns (arrayOfWords []string) []int {
	var output []int
	for _,value := range(arrayOfWords) {
		output = append(output, len(value))
	}
	return output
}

func frequenzimetro (word string) map[string]int {
	output := make(map[string]int)
	for _,value := range(word){
		
		_,ok := output[string(value)]
		if ok {
			output[string(value)] += 1
		} else {
			output[string(value)] = 1
		}
	}
	return output
}

func onlyForAssociates (word string, listOfWords []string) bool {
	for _,value := range(listOfWords) {
		if value == word {
			return true
		}
	}
	return false
}

func rovarSpraket (word string) string {
	var output string
	for _,char := range(word) {
		if strings.Contains("aeiou",string(char)) {
			output += string(char)
		} else {
			output += string(char)+"o"+string(char)
		}
	}
	return output
}

func timeLord (days, hours, minutes int) int {
	daysToSeconds := days*24*60*60
	hoursToSeconds := hours*60*60
	minutesToSeconds := minutes*60
	return daysToSeconds+hoursToSeconds+minutesToSeconds
}

func passWordGenerator (numberOfChars int) string {
	var output string
	const alfabeth string = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890!£$%&/()=?^*"
	pool := []rune(alfabeth)
	
	for i:=0;i<=numberOfChars;i++ {
		randomIndex := rand.Intn(len(pool))
		output += string(pool[randomIndex])
	}
	return output
}

func macGenerator () string {
	const alfa string = "ABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890"
	var generatedMac string
	for i:=0;i<=11;i++ {
		pool := []rune(alfa)
		randomIndex := rand.Intn(len(pool))
		generatedMac += string(pool[randomIndex])
		if i%2 == 1 {
			generatedMac += ":"
		}
	}
	return generatedMac
}

func rimario (word string, possibleRymes []string) string {
	for _,psb := range possibleRymes {
		if word[3:] == psb[3:] {
			return psb
		}
	}
	return "None"
}

func rot13 (word string) string {
	var cipher string
	const alfabeth string = "ABCDEFGHIJKLMNOPQRSTUVWXYZ ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	for _,letter := range(word) {
		fetchIndex := strings.Index(alfabeth,string(letter))
		cipher += string(alfabeth[fetchIndex+13])
	}
	return cipher
}

func factorial (number int) int {
	if number == 1 || number == 0 {
		return 1
	} else {
		return number * factorial(number-1)
	}
}

func fibonacci (position int) int {
	if position == 1 || position == 2 {
		return 1
	} else {
		return fibonacci(position-1) + fibonacci(position-2)
	}
}



func main () {
	fmt.Printf("Calling maxOfTwo with values %v and %v returns a value of %v\n",4,3,maxOfTwo(4,3))
	fmt.Printf("Calling maxOfThree with values %v, %v e %v returns a value of %v\n",4,3,5,maxOfThree(4,3,5))
	fmt.Printf("Calling maxOfAll with values %v return a value of %v\n",[]int{2,3,5},maxOfAll([]int{2,3,5}))
	fmt.Printf("Calling areUAVowel with value '%v' returns %v\n","a",areUAVowel("a"))
	fmt.Printf("Calling sommatriceInarrestabile with values %v returns a result of %v\n",[]int{1,2,3,4,5,6,7,8,9,10},sommatriceInarrestabile([]int{1,2,3,4,5,6,7,8,9,10}))
	fmt.Printf("Calling moltiplicatriceInarrestabile with values %v return a result of %v\n",[]int{1,2,3,4,5,6,7,8,9,10},moltiplicatriceInarrestabile([]int{1,2,3,4,5,6,7,8,9,10}))
	fmt.Printf("Calling reverser on '%v' return a result of '%v'\n","craicraispesafantastica",reverser("craicraispesafantastica"))
	fmt.Printf("Calling palindromoOrNot on %v return a result of %v\n","axooxa",palindromoOrNot("axooxa"))
	fmt.Printf("Calling personalizedLen on '%v' returns a lenght of %v\n","word",personalizedLen("word"))
	fmt.Printf("Calling istogramma on %v returns this:\n",[]int{3,4,5,1,2})
	istogramma([]int{3,4,5,1,2})
	fmt.Printf("Calling everybodysOwn on %v return this: %v\n",[]string{"Ciao","Mamma","Cane"},everybodysOwns([]string{"Ciao","Mamma","Cane"}))
	fmt.Printf("Calling frequenzimetro on %v returns this: %v\n","mamma",frequenzimetro("mamma"))
	fmt.Printf("Calling onlyForAssociates on %v and %v returns %v\n","word",[]string{"word","mamma","cane"},onlyForAssociates("word",[]string{"word","mamma","cane"}))
	fmt.Printf("Calling rovarSpraket on %v returns %v\n","mangiare",rovarSpraket("mangiare"))
	fmt.Printf("Calling timeLord on %v days,%v hrs,%v minutes returns %v\n",24,60,60,timeLord(24,60,60))
	fmt.Printf("Calling passWordGen on %v returns this: %v\n",130,passWordGenerator(130))
	fmt.Printf("Calling macGenerator on returns this: %v\n",macGenerator())
	fmt.Printf("Calling rimario on %v with %v returns %v","pane",[]string{"cane","mamma","presto"},rimario("pane",[]string{"cane","mamma","presto"}))
	fmt.Printf("Calling cipher on %v return a value of %v\n","TROPPOSTROPPIA", rot13("TROPPOSTROPPIA"))
	fmt.Printf("Calling factorial on %v returns a vlaue of %v\n",4,factorial(4))
	fmt.Printf("Calling fibonacci on %v returns a value of %v\n",7,fibonacci(7))
}

