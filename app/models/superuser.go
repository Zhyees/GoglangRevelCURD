package models

type Superuser struct {
	Id int `db:"id" json:"id"`
	Email string `db:"email" json:"email"`
	FullName string `db:"full_name" json:"full_name"`
	HashedPassword string `db:"password" json:"password"`
	LoginToken string `db:"login_token" json:"login_token"`
	RoleId int `db:"role_id" json:"role_id"`
	LastLogin string `db:"last_login" json:"last_login"`
	IsDeleted int `db:"is_deleted" json:"is_deleted"`
}

type SuperuserRecord struct {
	Id int `db:"id" json:"id"`
	Email string `db:"email" json:"email"`
	FullName string `db:"full_name" json:"full_name"`
	RoleId int `db:"role_id" json:"role_id"`
	RoleName string  `db:"role_name" json:"role_name"`
	IsDeleted int `db:"is_deleted" json:"is_deleted"`
}