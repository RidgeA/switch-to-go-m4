package behavioral

import (
	"strconv"
	"strings"
)

type (
	operand interface {
		evaluate() int
	}

	numberOperand struct {
		value int
	}

	plusOperand struct {
		left, right int
	}
)

func (n *numberOperand) evaluate() int {
	return n.value
}

func (p *plusOperand) evaluate() int {
	return p.left + p.right
}

func Evaluate(expression string) int {
	tokens := strings.Split(expression, " ")
	operands := make([]operand, 0, len(tokens))
	for _, token := range tokens {
		switch token {
		default:
			v, _ := strconv.Atoi(token)
			operands = append(operands, &numberOperand{v})
		case "+":
			tail := operands[len(operands)-2:]
			operands = append(operands[:len(operands)-2], &plusOperand{tail[0].evaluate(), tail[1].evaluate()})
		}
	}
	return operands[0].evaluate()
}
