package auth

import (
	"crypto/sha256"
	"net/http"
	"github.com/gorilla/securecookie"
	"io"
)

const (
	secret = "askn210987sdfy8u0sdf&T^UYGHIBOJP()*"
	secretHashKey = "@#$567*(O)lp;[}{_POIUYG"
	secretBlockKey = "UYFKLHOIghjbkh;y7i6uhjuoi8%^&"

	cookieKey = "gopy-auth"
)

var (
	sCookie = securecookie.New([]byte(secretHashKey), []byte(secretBlockKey))
)


// This type is used to identify any users and for any auth related
// tasks
type User struct {
	Email        string
	passwordHash []byte
	IsSuperUser  bool
	IsActive     bool
}

// Authenticate will return true if the password matches. false if not
func (u *User) Authenticate(pass string) bool {
	sha := generateHash(pass)
	if string(sha) == string(u.passwordHash) {
		return true
	}
	return false
}

// Set generate the password hash and store it into the user item
func (u *User) SetPassword(password string) {
	sha := generateHash(password)
	u.passwordHash = sha
}

// Generates hash from the password
func generateHash(pass string) []byte {
	pcrypt := sha256.New()
	salt := pass + secret
	io.WriteString(pcrypt, salt)
	output := pcrypt.Sum(nil)
	pcrypt.Reset()
	return output
}

// Logs in a specific email id to a session
func Login(w http.ResponseWriter, email string) error {
	// create data to write
	value := map[string]string{
		colEmail:email,
	}

	// encode and write the cookie if there is no error.
	if encoded, err := sCookie.Encode(cookieKey, value); err != nil {
		c := &http.Cookie{
			Name:cookieKey,
			Value:encoded,
			Path:"/",
		}
		http.SetCookie(w, c)
		return nil
	} else {
		return err
	}
}

// Returns a User and if he is authenticated
func Authenticate(r *http.Request) (User, bool) {
	cookie, err := r.Cookie(cookieKey)
	if err != nil {
		u := User{}
		return u, false
	}
	cookie_val := make(map[string]string)
	sCookie.Decode(cookieKey, cookie.Value, &cookie_val)
	email, found := cookie_val[colEmail]
	if !found {
		return User{}, false
	}
	u, err := GetUser(email)
	if err != nil {
		return User{}, false
	}
	if !u.IsActive {
		return u, false
	}
	return u, true
}