package interfaces

type User struct {
	ID       int32   `db:"id"`
	Nickname string  `db:"nickname"`
	Password []byte  `db:"password"`
	RoleId   int32   `db:"role_id"`
	Token    *string `db:"token"`
}

type Users interface {
	CreateUser(user User) error
	TokenExists(token string) (bool, error)
	GetUserByToken(token string) (User, error)
	GetUserByNickname(nickname string) (User, error)
	SeTokenByNickname(token, nickname string) error
	NicknameExists(nickname string) (bool, error)
}
