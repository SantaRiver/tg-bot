package models

type User struct {
	ID        int64
	Name      string
	Timestamp Timestamp
}

type CreateUsersSharesRequest struct {
	UserID    int64
	Shares    int64
	Timestamp Timestamp
}
