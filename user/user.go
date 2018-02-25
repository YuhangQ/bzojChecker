package user

type User struct {
	Username string
	Email    string
	Receive  bool
}

func NewUser(username string, email string, receive bool) User {
	User := new(User)
	User.Username = username
	User.Email = email
	User.Receive = receive
	return *User
}
