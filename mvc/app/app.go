package app

import "net/http"

// StartApp func
func StartApp() {
	http.HandleFunc("/users", controllers.GetUser)
}
