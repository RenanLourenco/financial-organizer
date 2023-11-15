package delete_expense

type DeleteExpenseInput struct {
	ID string `json:"id"`
}

type DeleteExpenseOutput struct {
	Msg string `json:"msg"`
}