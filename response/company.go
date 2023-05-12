package response

type CompanyResp struct {
	Id        int    `json:"id"`
	Name      string `json:"name"`
	Address   string `json:"address"`
	Balance   int    `json:"balance"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}
