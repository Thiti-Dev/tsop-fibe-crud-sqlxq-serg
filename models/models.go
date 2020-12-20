package models

import "database/sql"

type UserModelDb struct {
	UserID		sql.NullString `json:"UserID"`
	FirstName	sql.NullString	`json:"FirstName"`
	LastName	sql.NullString	`json:"LastName"`
	Age			sql.NullInt32	`json:"Age"`
	Status		sql.NullInt32	`json:"Status"`
	CreatedAt	sql.NullString	`json:"CreatedAt"`
}


// UserModelReq is for the request
type UserModelReq struct {
	FirstName	string	`json:"FirstName"`
	LastName	string	`json:"LastName"`
	Age			int	`json:"Age"`
	Status		int	`json:"Status"`	
}

type UserModelUpdate struct {
	UserID		string `json:"UserID"`
	FirstName	string	`json:"FirstName"`
	LastName	string	`json:"LastName"`
	Age			int	`json:"Age"`
	Status		int	`json:"Status"`	
}

type UserModel struct {
	UserID		string `json:"UserID"`
	FirstName	string	`json:"FirstName"`
	LastName	string	`json:"LastName"`
	Age			int	`json:"Age"`
	Status		int	`json:"Status"`
	CreatedAt	string	`json:"CreatedAt"`
}