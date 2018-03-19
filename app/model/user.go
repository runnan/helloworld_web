package model
import (
  "golang.org/x/crypto/bcrypt"
  "github.com/badoux/checkmail"
  "helloworld_web/app/util"
)

type User struct{
  Username string
  Email string
  HashedPassword string
  ID string
}

const (
  passwordLength = 6
  hashCost = 10
)

func NewUser(username, email, password string) (User, error){
  user := User{
    ID: util.GenerateID("user", 10),
    Email: email,
    Username: username,
  }
  if username == "" {
    return user, util.ErrNoUsername
  }
  if checkmail.ValidateFormat(email) != nil {
    return user, util.ErrEmail
  }
  if password == "" {
    return user, util.ErrNoPassword
  }
  if len(password) < passwordLength {
    return user, util.ErrPasswordTooShort
  }

  hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), hashCost)
  user.HashedPassword = string(hashedPassword)
  return user, err
}
