package cmd

import (
	"fmt"

	"os"

	"strconv"
)

// returns true if there is an operator to the left of strLoc
func IsOperatorToLeft(str string, strLoc int) bool {
	// strLoc isnt in string, exit the program
	if strLoc < 0 || strLoc >= len(str) {
		fmt.Println("Error in IsOperatorToLeft. strLoc is not in string.")
		os.Exit(1)
	}
	// operator is in leftmost location of string, no operator can be to left
	if strLoc == 0 {
		return false
	}
	// if there is an operator to the left
	if str[strLoc-1] == '*' || str[strLoc-1] == '+' ||
		(str[strLoc-1] == '-' && IsBinarySubOperator(str, strLoc-1)) {
		return true
	}
	return false
}

// returns true if a string converts successfully to a float, else false
func IsOperand(str string) bool {
	_, err := strconv.ParseFloat(str, 64)

	// if error converting to float
	if err != nil {
		if err == strconv.ErrSyntax { // its a syntax error, return false
			return false
		}
		// an unknown issue occured when converting the string, exit program
		fmt.Println("Unknown error checking if string is operand. Exiting now...")
		os.Exit(1)
	}
	return true
}

// returns location of left and right operand around an operator, given an expression string
// and an operatorLoc index. leftStart's value should be the leftmost index of the left operand,
// while rightEnd's index should be the rightmost index of the right operand
// Note: left or right operand may be -1 if that operand is not found.
func FindLeftRightOperands(expression string, operatorLoc int) (leftStart int, rightEnd int) {
	leftPtr, rightPtr := operatorLoc, operatorLoc
	leftStart, rightEnd = -1, -1

	// operatorLoc is out of bounds
	if 0 > operatorLoc || operatorLoc >= len(expression) {
		fmt.Println("Error occured in FindLeftRightOperands. operator location" +
			" is out of bounds.")
		return
	}
	// branch out left and right from operatorLoc until we find a character that isnt
	// 	
	for leftStart == -1 || rightEnd == -1 {
		if leftStart == -1 {
			leftPtr -= 1
			// if it isnt the operand anymore
			if leftPtr < 0 || expression[leftPtr] == '+' ||
				expression[leftPtr] == '(' || expression[leftPtr] == ')' ||
				expression[leftPtr] == '/' || expression[leftPtr] == '*' {
				leftPtr += 1
				leftStart = leftPtr
			}
		}

		if rightEnd == -1 {
			rightPtr += 1
			// if it isnt the operand anymore
			if rightPtr >= len(expression) || expression[rightPtr] == '+' ||
				expression[rightPtr] == '(' || expression[rightPtr] == ')' ||
				expression[rightPtr] == '/' || expression[rightPtr] == '*' ||
				(expression[rightPtr] == '-' && rightPtr != operatorLoc+1) {
				rightPtr -= 1
				rightEnd = rightPtr
			}
		}
	}
	return
}

// returns a new expression string with the newSubExpression inserted from letLoc to
// rightLoc (inclusive)
func ReplaceSubExp(expression string, newSubExpression string, leftLoc int, rightLoc int) string {
	// put result in expression
	tempExpression := ""
	if leftLoc != 0 {
		tempExpression += expression[:leftLoc]
	}
	tempExpression += newSubExpression
	if rightLoc != len(expression)-1 {
		tempExpression += expression[rightLoc+1:]
	}
	expression = tempExpression
	return expression
}

// returns true if the - sign is a binary operator, false if it is simply a unary
// negation operator: e.g., -5, and not 3-2
func IsBinarySubOperator(expression string, subLoc int) bool {
	return (expression[subLoc] == '-' && subLoc != 0 &&
		!IsOperatorToLeft(expression, subLoc))
}

// calculates only one pass of multiplication/division in a string, and returns
// the new, more simplified expression, and a bool determining if it changed the expression
// or not. example shown below:
// "3*4*5" -> ("12*5", true) 	first call of MulDivPemdas
// "12*5"  -> ("60", true) 		second call of MulDivPemdas
// "60"    -> ("60", false) 	third call of MulDivPemdas
func MulDivPemdas(inputExpression string) (expression string, didRun bool) {
	expression = inputExpression
	didRun = false

	// find first mul/div location
	operatorLoc := -1
	isMul, isDiv := false, false
	for i, c := range expression {
		if c == '*' {
			operatorLoc = i
			isMul = true
			break
		} else if c == '/' {
			operatorLoc = i
			isDiv = true
			break
		}
	}
	if !isMul && !isDiv {
		return
	}

	var leftOperandStr, rightOperandStr string
	var leftOperandLoc, rightOperandLoc int
	var result string

	leftOperandLoc, rightOperandLoc = FindLeftRightOperands(expression, operatorLoc)
	leftOperandStr = expression[leftOperandLoc:operatorLoc]
	rightOperandStr = expression[operatorLoc+1 : rightOperandLoc+1]

	leftOperandFloat, err := strconv.ParseFloat(leftOperandStr, 64)
	if err != nil {
		fmt.Printf("Error converting %s to float\n", leftOperandStr)
	}
	rightOperandFloat, err := strconv.ParseFloat(rightOperandStr, 64)
	if err != nil {
		fmt.Printf("Error converting %s to float\n", rightOperandStr)
	}
	if isMul {
		result = fmt.Sprintf("%f", leftOperandFloat*rightOperandFloat)
	} else {
		result = fmt.Sprintf("%f", leftOperandFloat/rightOperandFloat)
	}
	expression = ReplaceSubExp(expression, result, leftOperandLoc, rightOperandLoc)
	didRun = true
	return
}

// calculates only one pass of addition/subtraction in a string, and returns
// the new, more simplified expression, and a bool determining if it changed the expression
// or not. Note that the - sign can be both a binary operator OR unary operator in
// mathematics, so this function does a check to make sure every - is in fact subtraction and
// not negation.
// Example of some calls to this function shown below:
// "3+4-5" -> ("7-5", true) 	first call of AddSubPemdas
// "7-5"   -> ("2", true) 		second call of AddSubPemdas
// "2"     -> ("2", false) 		third call of AddSubPemdas
func AddSubPemdas(inputExpression string) (expression string, didRun bool) {
	didRun = false
	expression = inputExpression
	// find first mul/div location
	operatorLoc := -1
	isAdd, isSub := false, false
	for i, c := range expression {
		if c == '+' {
			operatorLoc = i
			isAdd = true
			break
		} else if IsBinarySubOperator(expression, i) { // if its binary subtraction operator
			operatorLoc = i
			isSub = true
			break
		}
	}
	if !isAdd && !isSub {
		return
	}

	var leftOperandStr, rightOperandStr string
	var leftOperandLoc, rightOperandLoc int
	var result string

	leftOperandLoc, rightOperandLoc = FindLeftRightOperands(expression, operatorLoc)
	leftOperandStr = expression[leftOperandLoc:operatorLoc]
	rightOperandStr = expression[operatorLoc+1 : rightOperandLoc+1]

	leftOperandFloat, err := strconv.ParseFloat(leftOperandStr, 64)
	if err != nil {
		fmt.Printf("Error converting %s to float\n", leftOperandStr)
	}
	rightOperandFloat, err := strconv.ParseFloat(rightOperandStr, 64)
	if err != nil {
		fmt.Printf("Error converting %s to float\n", rightOperandStr)
	}
	if isAdd {
		result = fmt.Sprintf("%f", leftOperandFloat+rightOperandFloat)
	} else {
		result = fmt.Sprintf("%f", leftOperandFloat-rightOperandFloat)
	}
	expression = ReplaceSubExp(expression, result, leftOperandLoc, rightOperandLoc)
	didRun = true
	return
}

// calculates only one pass of parenthesis recursively. The new, more simplified expression, and
// a bool determining if it changed the expression or not. This function finds the first instance of a
// paranthesis group and recursively solves that whole group by calling Pemdas() on everything
// inside the parenthesis and then removing the parenthesis.
// This function will not calculate a second parenthesis group. Example shown below.
// "(3+(4+7))+(5-6)" -> ("14+(5-6)", true)	First call of ParenthesisPemdas
// "14+(5-6)" 		 -> ("14+-1", true)		Second call of ParenthesisPemdas
// "14+-1"			 -> ("14+-1", false)	Third call of ParenthesisPemdas
func ParenthesisPemdas(inputExpression string) (expression string, didRun bool) {
	expression = inputExpression
	didRun = false
	parenCount := 0
	leftParenLoc, rightParenLoc := -1, -1
	for i, c := range expression {
		if c == '(' {
			if parenCount == 0 {
				leftParenLoc = i
			}
			parenCount += 1
		} else if c == ')' {
			if parenCount == 1 {
				rightParenLoc = i
				break
			}
			parenCount -= 1
		}
	}

	if leftParenLoc != -1 && rightParenLoc == -1 {
		fmt.Println("An error occurred. Mising )")
		os.Exit(1)
	} else if leftParenLoc == -1 && rightParenLoc != -1 {
		fmt.Println("An error occured. Unexpected )")
		os.Exit(1)
	} else if leftParenLoc != -1 && rightParenLoc != -1 {
		didRun = true
		subExpression := expression[leftParenLoc+1 : rightParenLoc]
		replaceSubExp := Pemdas(subExpression)
		expression = ReplaceSubExp(expression, replaceSubExp, leftParenLoc, rightParenLoc)
	}
	return
}

// given a string expression, calculates the value of said expression. This first
// calculates all values in Parenthesis (recursively), then Multiplication/Division
// simultaneously, then Addition/Subtraction simultaneously.
// Note that this function does NOT guarantee the expression will be evaluated correctly.
// There are some limitations currently, and bad input syntax in expression
// may lead to undefined behavior.
func Pemdas(expression string) string {
	var didRun bool = true

	for didRun {
		expression, didRun = ParenthesisPemdas(expression)
	}
	didRun = true
	for didRun {
		expression, didRun = MulDivPemdas(expression)
	}
	didRun = true
	for didRun {
		expression, didRun = AddSubPemdas(expression)
	}

	return expression
}
