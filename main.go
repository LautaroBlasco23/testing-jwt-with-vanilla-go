package main

import (
	"fmt"
	"jwt-go/controllers"
	"jwt-go/databases"
	"net/http"
)

// To handle the favicon
func doNothing(w http.ResponseWriter, r *http.Request){}

func main() {
  databases.ConnectDatabase()

  http.HandleFunc("/favicon.ico", doNothing)

  // User API Routes
  http.HandleFunc("/login", controllers.Login)
  http.HandleFunc("/register", controllers.Register)
  http.HandleFunc("/user", controllers.GetUserData)

  // Admin API Routes
  http.HandleFunc("/data", controllers.GetAllData)

  fmt.Println("Server running")
  http.ListenAndServe(":8000", nil)
}
