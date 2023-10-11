package models

import (
	"fmt"
	"net/http"
	"sync"
	"time"

	"github.com/gorilla/mux"
)

const maxLoginAttempts = 5
const lockoutDuration = 5 * time.Minute

var (
	users = map[string]string{
		"user1": "password1",
		"user2": "password2",
	}
	failedLoginAttempts = map[string]int{}
	lockoutTimers       = map[string]*time.Timer{}
	lockoutMutex        sync.Mutex
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/login", handleLogin).Methods("POST")

	http.Handle("/", r)
	fmt.Println("Server started on :8080")
	http.ListenAndServe(":8080", nil)
}

func handleLogin(w http.ResponseWriter, r *http.Request) {
	username := r.FormValue("username")
	password := r.FormValue("password")

	if isValidLogin(username, password) {
		fmt.Fprintf(w, "Login successful for %s\n", username)
		resetFailedAttempts(username)
	} else {
		handleFailedAttempt(username)
		fmt.Fprintf(w, "Login failed for %s\n", username)
	}
}

func isValidLogin(username, password string) bool {
	storedPassword, exists := users[username]
	return exists && storedPassword == password
}

func handleFailedAttempt(username string) {
	lockoutMutex.Lock()
	defer lockoutMutex.Unlock()

	failedLoginAttempts[username]++

	if failedLoginAttempts[username] >= maxLoginAttempts {
		if _, exists := lockoutTimers[username]; !exists {
			lockoutTimers[username] = time.AfterFunc(lockoutDuration, func() {
				resetFailedAttempts(username)
			})
		}
	}
}

func resetFailedAttempts(username string) {
	lockoutMutex.Lock()
	defer lockoutMutex.Unlock()

	delete(failedLoginAttempts, username)
	if timer, exists := lockoutTimers[username]; exists {
		timer.Stop()
		delete(lockoutTimers, username)
	}
}
