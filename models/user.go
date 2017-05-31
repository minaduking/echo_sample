package models

type(
	User struct{
		ID int `json:"id"`
		Name string `json:"name"`
	}
)

//本来DBに置き換わる領域。
var(
	Users = map[int]*User{}
	Seq = 1
)