package app_auth

type InsertUser struct {
	Name     string `gorm:"column:name;not null" json:"name"`
	Email    string `gorm:"column:email;not null;unique" json:"email"`
	Password string `gorm:"column:password;not null" json:"password"`
}

type UserLogIn struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserLogedIn struct {
	Token string `json:"token"`
}
