package math

type IntResult int64

type IntOperation struct {
	ValueA int
	ValueB int
}

func (io *IntOperation) Add() IntResult {
	return IntResult(io.ValueA + io.ValueB)
}

func (io *IntOperation) Subtract() IntResult {
	return IntResult(io.ValueA - io.ValueB)
}

func (io *IntOperation) Multiply() IntResult {
	return IntResult(io.ValueA * io.ValueB)
}

func (io *IntOperation) Divide() IntResult {
	if io.ValueB == 0 {
		return 0
	}
	return IntResult(io.ValueA / io.ValueB)
}

func (io *IntOperation) FromString(str string) IntResult {
	switch str {
	case "+":
		return io.Add()
	case "-":
		return io.Subtract()
	case "*":
		return io.Multiply()
	case "/":
		return io.Divide()
	default:
		return 0
	}
}
