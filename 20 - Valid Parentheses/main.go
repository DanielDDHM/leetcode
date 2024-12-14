package main

import "fmt"

func isValid(s string) bool {
	// Create a stack to store opening brackets
	stack := []rune{}

	// Create a map to match closing brackets with their corresponding opening brackets
	matchingBrackets := map[rune]rune{
		')': '(',
		'}': '{',
		']': '[',
	}

	// Loop through each character in the string
	for _, char := range s {
		// If the character is a closing bracket
		if openBracket, exists := matchingBrackets[char]; exists {
			// Check if the stack is empty or if the top of the stack is not the corresponding opening bracket
			if len(stack) == 0 || stack[len(stack)-1] != openBracket {
				return false
			}
			// Pop the top element from the stack
			stack = stack[:len(stack)-1]
		} else {
			// If it's an opening bracket, push it onto the stack
			stack = append(stack, char)
		}
	}

	// If the stack is empty, all brackets are balanced
	return len(stack) == 0
}

func main() {
	fmt.Println(isValid("()"))     // true
	fmt.Println(isValid("()[]{}")) // true
	fmt.Println(isValid("(]"))     // false
	fmt.Println(isValid("([)]"))   // false
	fmt.Println(isValid("{[]}"))   // true
}
