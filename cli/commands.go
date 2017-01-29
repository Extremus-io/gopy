package cli

import (
	"encoding/json"
	"github.com/Extremus-io/gopy/auth"
	"github.com/Extremus-io/gopy/log"
	"errors"
)

var commands = map[string]func(data *json.RawMessage) error{
	"create_user":func(user_json *json.RawMessage) error {
		if user_json == nil {
			return errors.New("no user data received for create_user")
		}
		user_raw := []byte(*user_json)
		user := struct {
			Email       string `json:"email"`
			Password    string `json:"password"`
			IsSuperUser bool `json:"is_superuser"`
		}{}
		log.Verbosef("%s", user_raw)
		err := json.Unmarshal(user_raw, &user)
		if err != nil {
			return err
		}
		u := auth.User{
			Email:user.Email,
			IsSuperUser:user.IsSuperUser,
			IsActive:true,
		}
		u.SetPassword(user.Password)
		err = auth.CreateUser(u)
		return err
	},
}