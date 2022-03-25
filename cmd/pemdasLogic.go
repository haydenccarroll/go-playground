package cmd

import (
	"fmt"

	"os"

	"strconv"
)

func IsOperatorToLeft(str string, strLoc int) bool {
	if strLoc < 0 || strLoc >= len(str) {
		fmt.Println("Error in IsOperatorToLeft. strLoc is not in string.")
		os.Exit(1)
	}
	if strLoc == 0 {
		return false
	}
	if str[strLoc-1] == '*' || str[strLoc-1] == '+' {
		return true
	}
	return false
}

func IsOperand(str string) bool {
	_, err := strconv.ParseFloat(str, 64)

	if err != nil {
		if err == strconv.ErrSyntax {
			return false
		}
		fmt.Println("Unknown error checking if string is operand. Exiting now...")
		os.Exit(1)
	}
	return true
}

func FindLeftRightOperands(expression string, operatorLoc int) (int, int) {
	if 0 > operatorLoc || operatorLoc >= len(expression) {
		fmt.Println("Error occured in FindLeftRightOperands. operator location" +
			" is out of bounds.")
		return -1, -1
	}

	leftPtr, rightPtr := operatorLoc, operatorLoc
	leftEnd, rightEnd := -1, -1
	for leftEnd == -1 || rightEnd == -1 {
		if leftEnd == -1 {
			leftPtr -= 1
			// if it isnt the operand anymore
			if leftPtr < 0 || expression[leftPtr] == '+' ||
				expression[leftPtr] == '(' || expression[leftPtr] == ')' ||
				expression[leftPtr] == '/' || expression[leftPtr] == '*' {
				leftPtr += 1
				leftEnd = leftPtr
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
	return leftEnd, rightEnd
}

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

func IsBinarySubOperator(expression string, subLoc int) bool {
	return (expression[subLoc] == '-' && subLoc != 0 &&
		!IsOperatorToLeft(expression, subLoc))
}

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
		replaceSubExp := RecursivePemdas(subExpression)
		expression = ReplaceSubExp(expression, replaceSubExp, leftParenLoc, rightParenLoc)
	}
	return
}

func RecursivePemdas(expression string) string {
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
