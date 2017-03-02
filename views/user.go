package views

import (
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/markcheno/go-gin-app/models"
)

// generateSessionToken create session token
func generateSessionToken() string {
	// DO NOT USE THIS IN PRODUCTION
	return strconv.FormatInt(rand.Int63(), 16)
}

// UserView -
type UserView struct {
	DB *models.DB
}

// ShowLoginPage render login page
func (u *UserView) ShowLoginPage(c *gin.Context) {
	render(c,
		gin.H{"title": "Login"},
		"login.html")
}

// PerformLogin handle login
func (u *UserView) PerformLogin(c *gin.Context) {

	username := c.PostForm("username")
	password := c.PostForm("password")

	if models.IsUserValid(username, password) {

		token := generateSessionToken()
		c.SetCookie("token", token, 3600, "", "", false, true)
		c.Set("is_logged_in", true)
		render(c,
			gin.H{"title": "Successful Login"},
			"login-successful.html")

	} else {
		c.HTML(http.StatusBadRequest,
			"login.html",
			gin.H{"ErrorTitle": "Login Failed",
				"ErrorMessage": "Invalid credentials provided"})
	}
}

// Logout handle logout
func (u *UserView) Logout(c *gin.Context) {
	// Clear the cookie
	c.SetCookie("token", "", -1, "", "", false, true)
	c.Redirect(http.StatusTemporaryRedirect, "/")
}

// ShowRegistrationPage render registration page
func (u *UserView) ShowRegistrationPage(c *gin.Context) {
	render(c,
		gin.H{"title": "Register"},
		"register.html")
}

// Register render register page
func (u *UserView) Register(c *gin.Context) {

	// Obtain the POSTed username and password values
	username := c.PostForm("username")
	password := c.PostForm("password")

	if _, err := models.RegisterNewUser(username, password); err == nil {
		// If the user is created, set the token in a cookie and log the user in
		token := generateSessionToken()
		c.SetCookie("token", token, 3600, "", "", false, true)
		c.Set("is_logged_in", true)

		render(c,
			gin.H{"title": "Successful registration & Login"},
			"login-successful.html")

	} else {
		c.HTML(http.StatusBadRequest,
			"register.html",
			gin.H{"ErrorTitle": "Registration Failed",
				"ErrorMessage": err.Error()})
	}
}
