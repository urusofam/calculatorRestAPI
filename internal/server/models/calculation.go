package models

type Calculation struct {
	ID         string `json:"id"`
	Expression string `json:"expression"`
	Result     string `json:"result"`
}

type CalculationsRequest struct {
	Expression string `json:"expression"`
}
