package entity

type AccountEntity struct {
	ID      uint  `json:"id"`
	UserID  uint  `json:"user_id"`
	Balance int64 `json:"balance"`
}
