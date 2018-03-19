package route

import (
  "net/http"
  "github.com/dgrijalva/jwt-go"
  "helloworld_web/app/util"
  "time"
)

/* Set up a global string for our secret */
var mySigningKey = []byte(util.GetConfigurationValue("secret_key"))

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("Token")
		if err != nil {
			http.Redirect(w, r, "/login", http.StatusFound)
		} else {
      info, result := ExtractClaims(cookie.Value)
      if result {
        cookie    :=    http.Cookie{ Name: "Token", Value: GenerateToken(info["username"].(string))}
        http.SetCookie(w, &cookie)
      	next.ServeHTTP(w, r)
      } else {
      	http.Redirect(w, r, "/login", http.StatusFound)
      }
		}

	})
}

func GenerateToken(username string) string{
  token := jwt.New(jwt.SigningMethodHS256)

  claims := token.Claims.(jwt.MapClaims)
  claims["username"] = username
  claims["exp"] = time.Now().Add(time.Hour * 1).Unix()

  /* Sign the token with our secret */
  tokenString, _ := token.SignedString(mySigningKey)
  return tokenString
}

func ExtractClaims(tokenStr string) (jwt.MapClaims, bool) {
  token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
       return mySigningKey, nil
  })

  if err != nil {
      return nil, false
  }

  if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
      return claims, true
  } else {
      return nil, false
  }
}
