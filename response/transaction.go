package response

type TransactionResp struct {
	Id        int    `json:"id"`
	Amount    int    `json:"amount"`
	Note      string `json:"note"`
	Type      string `json:"type"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}
