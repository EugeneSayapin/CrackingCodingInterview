package main

import "fmt"

func countSpaces(input string, length int) int {
	runes := []rune(input);
	result := 0;
	for i:=0; i < length; i++ {
		if runes[i] == ' '  {
			result++;
		}
	}
	return result;
}

func URLify (input string, length int) string {
	inputPos := length - 1;
	outputPos := inputPos + countSpaces(input, length)*2;
	runes := []rune(input);
	for  inputPos >=0 {
		if runes[inputPos] != ' '  {
			runes[outputPos] = runes[inputPos];
			outputPos--;
			inputPos--;
		} else {
			runes[outputPos-2] = '%';
			runes[outputPos-1] = '2';
			runes[outputPos] = '0';
			outputPos -= 3;			
			inputPos--;
		}
	}
	return string(runes)
}

func main (){
	var s1 = "Hello, Mr John Smith       ";	
	fmt.Println(URLify(s1, 20));

}