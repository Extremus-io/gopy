package auth

// Retrieve a user from the database
// error returned will be the database select query error
func GetUser(email string) (User, error) {
	u := User{}
	row := stmtGetUserByEmail.QueryRow(email)
	err := row.Scan(&u.Email, &u.passwordHash, &u.IsSuperUser, &u.IsActive)
	return u, err
}

// Create a new user entry in database
// error returned will be the database insert query error
func CreateUser(u User) error {
	_, err := stmtInsertUser.Exec(u.Email, u.passwordHash, u.IsSuperUser, u.IsActive)
	return err
}

// Delete user entry in database
// error returned will be the database delete query error
func DeleteUser(email string) error {
	_, err := stmtDeleteUser.Exec(email)
	return err
}

// Update the status of current user
// error returned will be the database update query error
func UpdateUser(u User) error {
	_, err := stmtUpdateUser.Exec(u.passwordHash, u.IsSuperUser, u.IsActive, u.Email)
	return err
}