package handler

type InquiryAccountRequest struct {
	AccountNumber string `json:"accountNumber"`
}

type InquiryAccountResponse struct {
	AccountNumber string `json:"accountNumber"`
	AccountName   string `json:"accountName"`
}
