package auth

import (
	"testing"
	"os"
	"github.com/Extremus-io/gopy/db"
	"github.com/Extremus-io/gopy/log"
)

const testEmail = "kittuov@gmail.com"
const testPass = "kittu1234"

func TestUser_Authenticate(t *testing.T) {
	const password = "mypassword"
	pass_enc := generateHash(password)
	log.Verbosef("generated hash ''%s''",pass_enc)
	u := User{Email:"kittuov@gmail.com", passwordHash:pass_enc}
	if !u.Authenticate(password) {
		t.Error("hash mechanism produced different results")
	}
}

func TestUser(t *testing.T) {
	t.Run("create-user", TestCreateUser)
	t.Run("get-user", TestGetUser)
	t.Run("update-user", TestUpdateUser)
	t.Run("delete-user", TestDeleteUser)
	os.Remove(db.DBName)
}

func TestCreateUser(t *testing.T) {
	u := User{
		Email: testEmail,
		IsSuperUser:true,
		IsActive:true,
	}
	u.SetPassword(testPass)
	err := CreateUser(u)
	if err != nil {
		t.Error("unable to create user")
		t.Errorf("%v", err)
		t.Fail()
	}
}
func TestGetUser(t *testing.T) {
	user, err := GetUser(testEmail)
	if err != nil {
		t.Error("unable to get user")
		t.Error(err.Error())
		t.Fail()
	}
	if user.Email != testEmail || !user.Authenticate(testPass) {
		t.Error("wrong data received")
		t.Fail()
	}
}
func TestUpdateUser(t *testing.T) {
	user, err := GetUser(testEmail)
	if err != nil {
		t.Error(err.Error())
		t.Fail()
	}
	user.IsSuperUser = false
	err = UpdateUser(user)
	if err != nil {
		t.Error(err.Error())
		t.Fail()
	}
	user, _ = GetUser(testEmail)
	if user.IsSuperUser {
		t.Error("Data not updated properly")
		t.Fail()
	}
}
func TestDeleteUser(t *testing.T) {
	err := DeleteUser(testEmail)
	if err != nil {
		t.Errorf("delete Email failed `%s`", err.Error())
		t.Fail()
	}
	_, err = GetUser(testEmail)
	if err == nil {
		t.Error("user not deleted")
		t.Fail()
	}
}