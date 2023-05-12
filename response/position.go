package response

type PositionResp struct {
	Id        string `json:"id"`
	Name      string `json:"name"`
	Salary    int    `json:"salary"`
	CreatedAt string `json:"created_at" `
	UpdatedAt string `json:"updated_at"`
}
