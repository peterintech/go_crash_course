package user

import "fmt"

type User struct {
	Name, Email, password string
}

func (u *User) SetPassword(password string) string {
	u.password = password

	return fmt.Sprintf("\nPassword for %s has been set to %s", u.Name, u.password)
}

func (u *User) CheckPassword(password string) bool {
	if password == u.password {
		return true
	} else {
		return false
	}
}
