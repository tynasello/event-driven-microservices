package entity

type User struct {
	Id       int
	Username string
	Password string
}

func (u User) Equals(other User) bool {
	return u.Id == other.Id && u.Username == other.Username && u.Password == other.Password
}
