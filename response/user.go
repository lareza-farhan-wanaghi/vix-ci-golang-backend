package response

type UserResp struct {
	Id         int           `json:"id"`
	SecretId   string        `json:"secret_id"`
	Name       string        `json:"name"`
	Email      string        `json:"email"`
	Phone      string        `json:"phone"`
	Address    string        `json:"address"`
	PositionId int           `json:"position_id"`
	Position   *PositionResp `json:"position"`
	CreatedAt  string        `json:"created_at"`
	UpdatedAt  string        `json:"updated_at"`
}
