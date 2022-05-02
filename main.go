package main

import (
	"fmt"
	"strings"
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
	for i,_ := range word {
		newWord += string(word[len(word)-i-1])
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

func frequenzimetro (parole []string) []int {
	var output []int
	for _,parola := range parole {
		output = append(output,len(parola))
	}
	return output
}

func associates (word string, listOfWords []string) bool {
	for _,test := range listOfWords {
		if word == test {
			return true
		} else {
			continue
		}
	}
	return false
}

func rovarspraket (word string) string {
	const alfabeth string = "aeiou"
	var newWord string
	for _,value := range(word) {
		if strings.Contains(alfabeth,string(value)) {
			newWord += string(value)
		} else {
			newWord += "o" + string(value) + "o"
		}
	}
	return newWord
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
	fmt.Printf("Calling frequenzimetro on %v returns this lenghts: %v\n",[]string{"ciao","Mamma","Nicola"},frequenzimetro([]string{"ciao","Mamma","Nicola"}))
	fmt.Printf("Calling associates with param %v on %v returns %v\n","mamma",[]string{"mamma","papà"},associates("mamma",[]string{"mamma","papà"}))
	fmt.Printf("Calling rovarspraket on %v returns %v\n","ma",rovarspraket("ma"))
}