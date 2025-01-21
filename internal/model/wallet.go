package model

import "time"

type WalletResponse struct {
	UserID    int32     `json:"user_id"`
	Balance   float32   `json:"balance"`
	CreatedAt time.Time `json:"created_at"`
}
