package cli

import (
	"encoding/json"
	"github.com/Extremus-io/gopy/auth"
)

var commands = map[string]func(data []byte) error{
	"create_user":func(user_raw []byte) error {
		user := struct {
			Email       string `json:"email"`
			Password    string `json:"password"`
			IsSuperUser bool `json:"is_superuser"`
		}{}

		json.Unmarshal(user_raw, &user)
		u := auth.User{
			Email:user.Email,
			IsSuperUser:user.IsSuperUser,
			IsActive:true,
		}
		u.SetPassword(user.Password)
		err := auth.CreateUser(u)
		return err
	},
}