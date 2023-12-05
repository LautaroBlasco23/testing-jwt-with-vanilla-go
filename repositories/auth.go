package repositories

import (
  "errors"
  "os"
  "fmt"

	"github.com/golang-jwt/jwt/v5"
)

// Middlewares are handled inside controllers
func UserTokenRequired(token_string string, user_email string) bool {
  token, err := jwt.Parse(token_string, func(token *jwt.Token) (interface{}, error) {
    if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
      return nil, errors.New("wrong algorithm")
    }

    return []byte(os.Getenv("JWT_SECRET")), nil
  })
  if err != nil {
    return false
  }

  if claims, ok := token.Claims.(jwt.MapClaims); ok {
    if claims["email"] != user_email {
      return false
    }
  } else {
    fmt.Println(err)
  }
  return true
}

func AdminTokenRequired(token_string string) bool {
  _, err := jwt.Parse(token_string, func(token *jwt.Token) (interface{}, error) {
    if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
      return nil, errors.New("wrong algorithm")
    }

    return []byte(os.Getenv("JWT_ADMIN")), nil
  })
  if err != nil {
    return false
  }

  return true
}

func CreateUserToken(email string) string {
  token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims {
    "email": email,
  })
  signedToken, err :=  token.SignedString([]byte(os.Getenv("JWT_SECRET")))
  if err != nil {
    fmt.Println(err)
    fmt.Println("error signing token")
    return ""
  }

  return signedToken
}

func CreateAdminToken(email string) string {
  token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims {
    "email": email,
  })
  signedToken, err :=  token.SignedString([]byte(os.Getenv("JWT_ADMIN")))
  if err != nil {
    fmt.Println(err)
    fmt.Println("error signing token")
    return ""
  }

  return signedToken
}

func ValidAdminSecret(admin_secret string) bool {
  if admin_secret == os.Getenv("ADMIN_PASS") {
    return true
  }
  return false
}
