package dto

type CreateAccountRequest struct {
	Name   string `json:"name"`
	Amount int    `json:"amount"`
}

type ChangeAccountAmountRequest struct {
	Name      string `json:"name"`
	NewAmount int    `json:"newAmount"`
}

type ChangeAccountNameRequest struct {
	Name    string `json:"name"`
	NewName string `json:"newName"`
}

type DeleteAccountRequest struct {
	Name string `json:"name"`
}
