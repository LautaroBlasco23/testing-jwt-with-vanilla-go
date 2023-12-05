package controllers

import (
	"encoding/json"
	"jwt-go/models"
	"jwt-go/repositories"
	"net/http"
	"strconv"
	"strings"
)

// USER HANDLERS
func Login(w http.ResponseWriter, req *http.Request) {
  if req.Method == "POST" {
    var new_user models.User
    json.NewDecoder(req.Body).Decode(&new_user)

    type UserAndToken struct {
      User models.User
      Token models.Token
    }

    // User
    user := repositories.GetUserByEmailFromDatabase(new_user.Email)

    var signedToken string
    if user.Admin {
      signedToken = repositories.CreateAdminToken(new_user.Email)
    } else {
      signedToken = repositories.CreateUserToken(new_user.Email)
    }

    if len(signedToken) == 0 {
      w.Write([]byte("error creating token"))
    }

    return_value := UserAndToken {
      User: user,
      Token: models.Token{Token: signedToken},
    }

    json.NewEncoder(w).Encode(return_value)
  } else {
    w.WriteHeader(http.StatusBadRequest)
    w.Write([]byte("invalid method"))
  }
}

func Register(w http.ResponseWriter, req *http.Request) {
  if req.Method == "POST" {
    var user_input models.UserRegisterInput
    json.NewDecoder(req.Body).Decode(&user_input)

    new_user := models.User{
      Email: user_input.Email,
      Password: user_input.Password,
    }

    // If Admin Secret is correct, the user will be an admin.
    // This is the only way to create an admin.
    // You can modify user's admin value in the Database if you want.
    if repositories.ValidAdminSecret(user_input.AdminSecret) {
      new_user.Admin = true
    }

    err := repositories.InsertUserInDatabase(&new_user)
    if err != nil {
      w.WriteHeader(http.StatusInternalServerError)
      if strings.Contains(err.Error(), "duplicate key") {
        w.Write([]byte("email already in use"))
      } else {
        w.Write([]byte("error registering user"))
      }
      return
    }

    user := repositories.GetUserByEmailFromDatabase(new_user.Email)

    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(user)
  } else {
    w.WriteHeader(http.StatusBadRequest)
    w.Write([]byte("invalid method"))
  }
}

func GetAllData(w http.ResponseWriter, req *http.Request) {
  if req.Method == "GET" {
    token_result := false
    // Getting Authorization value
    for _, cookies := range req.Header.Values("Cookie") {
      list_of_cookies := strings.Split(cookies, ";")

      for _, cookie := range list_of_cookies {
        cookie_name := strings.Split(cookie, "=")[0]
        cookie_name = strings.Trim(cookie_name, " ")
        if cookie_name == "Authorization" {
          auth_value := strings.Trim(strings.Split(cookie, "=")[1], "\"")
          // checking token 
          token_result =  repositories.AdminTokenRequired(auth_value)
        }
      }
    } 
    if token_result {
      list_of_users := repositories.GetUsersFromDatabase()
      json.NewEncoder(w).Encode(list_of_users)
    } else {
      w.WriteHeader(http.StatusUnauthorized)
      w.Write([]byte("invalid token"))
    }
  } else {
    w.WriteHeader(http.StatusBadRequest)
    w.Write([]byte("invalid method"))
  }
}

func GetUserData(w http.ResponseWriter, req *http.Request) {
  if req.Method == "GET" {
    id := req.URL.Query().Get("id")
    id_int, err := strconv.Atoi(id)
    if err != nil {
      w.Write([]byte("invalid ID"))
      return
    }
    user := repositories.GetUserByIdFromDatabase(id_int)

    token_result := false
    // Getting Authorization value
    for _, cookies := range req.Header.Values("Cookie") {
      list_of_cookies := strings.Split(cookies, ";")

      for _, cookie := range list_of_cookies {
        cookie_name := strings.Split(cookie, "=")[0]
        cookie_name = strings.Trim(cookie_name, " ")
        if cookie_name == "Authorization" {
          auth_value := strings.Trim(strings.Split(cookie, "=")[1], "\"")
          // checking token 
          if repositories.AdminTokenRequired(auth_value) || repositories.UserTokenRequired(auth_value, user.Email) {
            token_result = true
          } else {
            token_result = false
          }
        }
      }
    } 
    if token_result {
      id, err := strconv.Atoi(id)
      if err != nil {
        w.Write([]byte("invalid ID"))
        return
      }

      user := repositories.GetUserByIdFromDatabase(id)
      if user == (models.User{}) {
        w.WriteHeader(http.StatusBadRequest)
        w.Write([]byte("user not found"))
      } else {
        json.NewEncoder(w).Encode(user)
      }
    } else {
      w.WriteHeader(http.StatusUnauthorized)
      w.Write([]byte("invalid token"))
    }
  } else {
    w.WriteHeader(http.StatusBadRequest)
    w.Write([]byte("invalid method"))
  }
}
