package cmd

import (
	"fmt"

	"os"

	"strconv"
)

type operands struct {
	str string
	val int
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
				expression[rightPtr] == '/' || expression[rightPtr] == '*' {
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

func MulDivPemdas(expression string) string {
	// calculate addition
	mulLoc, divLoc := -1, -1
	for i, c := range expression {
		if c == '*' {
			mulLoc = i
			break
		} else if c == '/' {
			divLoc = i
			break
		}
	}
	if mulLoc == -1 && divLoc == -1 {
		return expression
	}

	var leftOperandStr, rightOperandStr string
	var leftOperandLoc, rightOperandLoc int
	var result string
	if mulLoc != -1 {
		leftOperandLoc, rightOperandLoc = FindLeftRightOperands(expression, mulLoc)
		leftOperandStr = expression[leftOperandLoc:mulLoc]
		rightOperandStr = expression[mulLoc+1 : rightOperandLoc+1]
	} else {
		leftOperandLoc, rightOperandLoc = FindLeftRightOperands(expression, divLoc)
		leftOperandStr = expression[leftOperandLoc:divLoc]
		rightOperandStr = expression[divLoc+1 : rightOperandLoc+1]
	}
	leftOperandFloat, err := strconv.ParseFloat(leftOperandStr, 64)
	if err != nil {
		fmt.Printf("Error converting %s to float", leftOperandStr)
	}
	rightOperandFloat, err := strconv.ParseFloat(rightOperandStr, 64)
	if err != nil {
		fmt.Printf("Error converting %s to float", rightOperandStr)
	}
	if mulLoc != -1 {
		result = fmt.Sprintf("%f", leftOperandFloat*rightOperandFloat)
	} else {
		result = fmt.Sprintf("%f", leftOperandFloat/rightOperandFloat)
	}
	expression = ReplaceSubExp(expression, result, leftOperandLoc, rightOperandLoc)
	return expression
}

func AdditionPemdas(expression string) string {
	// calculate addition
	additionLoc := -1
	for i, c := range expression {
		if c == '+' {
			additionLoc = i
			break
		}
	}
	if additionLoc != -1 {
		leftOperandLoc, rightOperandLoc := FindLeftRightOperands(expression, additionLoc)
		leftOperandStr := expression[leftOperandLoc:additionLoc]
		rightOperandStr := expression[additionLoc+1 : rightOperandLoc+1]

		leftOperandFloat, err := strconv.ParseFloat(leftOperandStr, 64)
		if err != nil {
			fmt.Printf("Error converting %s to float", leftOperandStr)
		}
		rightOperandFloat, err := strconv.ParseFloat(rightOperandStr, 64)
		if err != nil {
			fmt.Printf("Error converting %s to float", rightOperandStr)
		}

		result := fmt.Sprintf("%f", leftOperandFloat+rightOperandFloat)
		expression = ReplaceSubExp(expression, result, leftOperandLoc, rightOperandLoc)
	}
	return expression
}

func ParenthesisPemdas(expression string) string {
	leftParenCount, rightParenCount := 0, 0
	leftParen, rightParen := -1, -1
	for i, c := range expression {
		if c == '(' && leftParenCount == 0 {
			leftParenCount += 1
			leftParen = i
		} else if c == ')' && (leftParenCount-rightParenCount) == 1 {
			rightParen = i
			break
		}
	}

	// calculate parenthesis
	if leftParen != -1 && rightParen != -1 {
		subexpression := expression[leftParen+1 : rightParen]
		replaceSubExp := RecursivePemdas(subexpression)
		expression = ReplaceSubExp(expression, replaceSubExp, leftParen, rightParen)

	} else if leftParen != -1 && rightParen == -1 {
		fmt.Println("An error occurred. Mising )")
		os.Exit(1)
	} else if leftParen == -1 && rightParen != -1 {
		fmt.Println("An error occured. Unexpected )")
		os.Exit(1)
	}
	return expression
}

func RecursivePemdas(expression string) string {
	expression = ParenthesisPemdas(expression)
	expression = MulDivPemdas(expression)
	expression = AdditionPemdas(expression)

	_, err := strconv.ParseFloat(expression, 64)
	if err == nil {
		return expression
	}
	return RecursivePemdas(expression)
}