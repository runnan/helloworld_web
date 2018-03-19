package store
import (
  "helloworld_web/app/model"
  "golang.org/x/crypto/bcrypt"
  "helloworld_web/app/db"
  "helloworld_web/app/util"
)

var GlobalUserStore UserStore

type UserStore struct {

}

func (store *UserStore) Save(user *model.User) error {
  _,
  err := db.GlobalMySQLDB.Exec(`REPLACE INTO users (id, username, email, password) VALUES (?, ?, ?, ?)`,
  user.ID,
  user.Username,
  user.Email,
  user.HashedPassword,
)
return err
}

func (store *UserStore) FindByEmail(email string) (*model.User, error) {
  row := db.GlobalMySQLDB.QueryRow(
    `
    SELECT id, username, email, password
    FROM users
    WHERE email = ?`,
    email,
  )
  user := model.User{}
  err := row.Scan(
    &user.ID,
    &user.Username,
    &user.Email,
    &user.HashedPassword,
  )
  return &user, err
}

func (store *UserStore) FindByUsername(username string) (*model.User, error) {
  row := db.GlobalMySQLDB.QueryRow(
    `
    SELECT id, username, email, password
    FROM users
    WHERE username = ?`,
    username,
  )
  user := model.User{}
  err := row.Scan(
    &user.ID,
    &user.Username,
    &user.Email,
    &user.HashedPassword,
  )
  return &user, err
}

func (store *UserStore) FindByUsernameOrEmail(username string, email string) (*model.User, error) {
  row := db.GlobalMySQLDB.QueryRow(
    `
    SELECT id, username, email, password
    FROM users
    WHERE username = ? or email = ?`,
    username, email,
  )
  user := model.User{}
  err := row.Scan(
    &user.ID,
    &user.Username,
    &user.Email,
    &user.HashedPassword,
  )
  return &user, err
}

func (store *UserStore) FindUser(username, password string) (*model.User, error) {
  out := &model.User{
    Username: username,
  }
  existingUser, err := store.FindByUsername(username)
  if err != nil || existingUser == nil || bcrypt.CompareHashAndPassword([]byte(existingUser.HashedPassword),[]byte(password),) != nil{
    return out, util.ErrCredentialsIncorrect
  }
  return existingUser, nil
}
