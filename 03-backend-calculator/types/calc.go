package types

type CalculatorRequest struct {
	First  int `json:"number1" validate:"required"`
	Second int `json:"number2" validate:"required"`
}

type CalculatorResponse struct {
	Result int `json:"result"`
}
