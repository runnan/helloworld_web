package util

import (
  "errors"
)

type ValidationError error

var (
  ErrNoUsername = ValidationError(errors.New("You must supply a username"))
  ErrEmail = ValidationError(errors.New("You must supply a valid email"))
  ErrNoPassword = ValidationError(errors.New("You must supply a password"))
  ErrPasswordTooShort = ValidationError(errors.New("Your password is too short"))
  ErrUsernameOrEmailExists = ValidationError(errors.New("That username or email is taken"))
  ErrCredentialsIncorrect = ValidationError(errors.New("We couldn’t find a user with the supplied username and password combination"))
)

func IsValidationError(err error) bool {_, ok := err.(ValidationError)
  return ok
}
