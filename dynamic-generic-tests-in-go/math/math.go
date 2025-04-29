package math

type IntOperation struct {
	ValueA int
	ValueB int
}

type IntResult int64

func (io *IntOperation) Add() IntResult {
	return IntResult(io.ValueA + io.ValueB)
}

func (io *IntOperation) Multiply() IntResult {
	return IntResult(io.ValueA * io.ValueB)
}

func (io *IntOperation) Subtract() IntResult {
	return IntResult(io.ValueA - io.ValueB)
}

func (io *IntOperation) Divide() IntResult {
	if io.ValueB == 0 {
		return IntResult(0)
	}
	return IntResult(io.ValueA / io.ValueB)
}

func (io *IntOperation) FromString(operator string) IntResult {
	switch operator {
	case "+":
		return io.Add()
	case "-":
		return io.Subtract()
	case "*":
		return io.Multiply()
	case "/":
		return io.Divide()
	default:
		return IntResult(0)
	}
}
