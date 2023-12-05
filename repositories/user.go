package repositories

import (
	"fmt"
	"jwt-go/databases"
	"jwt-go/models"
)

func GetUsersFromDatabase() []models.User {
  db := databases.DB.DB
  var list_of_users []models.User

  err := db.Find(&list_of_users).Error
  if err != nil {
    fmt.Println(err)
    return []models.User{}
  }

  return list_of_users
}

func InsertUserInDatabase(new_user *models.User) error {
  db := databases.DB.DB

  err := db.Create(new_user).Error
  if err != nil {
    fmt.Println(err.Error())
    return err
  }

  return nil
}

func GetUserByEmailFromDatabase(user_email string) models.User {
  db := databases.DB.DB
  var user models.User

  err := db.Find(&user, "email = ?", user_email).Error
  if err != nil {
    fmt.Println(err)
    return models.User{}
  }

  return user 
}

func GetUserByIdFromDatabase(user_id int) models.User {
  db := databases.DB.DB
  var user models.User

  err := db.Find(&user, "ID = ?", user_id).Error
  if err != nil {
    fmt.Println(err)
    return models.User{}
  }

  return user 
}
