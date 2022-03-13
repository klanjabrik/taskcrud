// controllers/task.go

package controllers

import (
	"net/http"
	"os"

	api "taskcrud/api"
	"taskcrud/helpers"
	"taskcrud/models"
	"taskcrud/utils/messaging"
	"taskcrud/utils/token"

	"github.com/gin-gonic/gin"

	"net/url"

	"golang.org/x/crypto/bcrypt"
)

type CreateUserInput struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

type UpdateUserInput struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

type LoginInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type UpdateTokenInput struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

// GET /users
// Get all tasks
func FindUsers(c *gin.Context) {
	var users []models.User

	m := map[string]interface{}{}
	if err := models.GetAllUser(&users, m); err != nil {
		api.RespondError(c, http.StatusNotFound,
			helpers.GetFormattedMessage("empty_notfound_formatted", "User"))
	}

	api.RespondSuccess(c, http.StatusOK, users)
}

// POST /register
// Create new task
func CreateUser(c *gin.Context) {
	// Validate input
	var input CreateUserInput
	if err := c.ShouldBindJSON(&input); err != nil {
		api.RespondError(c, http.StatusBadRequest, err.Error())
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		api.RespondError(c, http.StatusBadRequest, err.Error())
		return
	}

	// create token verification
	emailToken, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		api.RespondError(c, http.StatusBadRequest, err.Error())
		return
	}

	urlEmail := os.Getenv("API_URL") + "users/verification/" + url.QueryEscape(string(emailToken))

	data := `
	{
		"subject":"Welcome",
		"template":"user_register.html",
		"rcpt": ["` + input.Email + `"],
		"data":{
			"Name": "` + input.Username + `",
			"Url": "` + urlEmail + `"
		}
	}`

	messaging.Send(c, "send_email", data)

	// Create user
	user := models.User{Username: input.Username, Password: string(hashedPassword), Email: input.Email, EmailToken: string(emailToken)}

	if models.AddNewUser(&user) != nil {
		api.RespondError(c, http.StatusBadRequest, err.Error())
		return
	}

	api.RespondSuccess(c, http.StatusOK, user)
}

// POST /login
// Login a user
func Login(c *gin.Context) { // Get model if exist
	var user models.User
	var input LoginInput
	if err := c.ShouldBindJSON(&input); err != nil {
		api.RespondError(c, http.StatusBadRequest, err.Error())
		return
	}

	m := map[string]interface{}{"username": input.Username}
	if err := models.GetOneUser(&user, m); err != nil {
		api.RespondError(c, http.StatusNotFound,
			helpers.GetFormattedMessage("empty_notfound_formatted", "User"))
		return
	}

	if user.EmailToken != "" {
		api.RespondError(c, http.StatusForbidden,
			helpers.GetMessage("email_not_verified"))
		return
	}

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password))
	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		api.RespondError(c, http.StatusBadRequest, err.Error())
		return
	}

	generatedToken, err := token.GenerateToken(user.ID)
	if err != nil {
		api.RespondError(c, http.StatusBadRequest, err.Error())
		return
	}

	api.RespondLoginSuccess(c, http.StatusOK, user, generatedToken)
}

// GET /resetpassword
// Reset password
func ResetPassword(c *gin.Context) { // Get model if exist
	var user models.User

	// db := c.MustGet("db").(*gorm.DB)
	// if err := db.Where("id = ?", c.Param("id")).First(&task).Error; err != nil {
	// 	api.RespondError(c, http.StatusBadRequest,
	// 		helpers.GetFormattedMessage("empty_notfound"))
	// 	return
	// }

	api.RespondSuccess(c, http.StatusOK, user)
}

// GET /users/:id
// Find a user
func FindUser(c *gin.Context) { // Get model if exist
	var user models.User

	if err := models.GetOneUserId(&user, c.Param("id")); err != nil {
		api.RespondError(c, http.StatusNotFound,
			helpers.GetFormattedMessage("empty_notfound_formatted", "User"))
		return
	}

	api.RespondSuccess(c, http.StatusOK, user)
}

// PATCH /users/:id
// Update a task
func UpdateUser(c *gin.Context) {

	// Get model if exist
	var user models.User
	if err := models.GetOneUserId(&user, c.Param("id")); err != nil {
		api.RespondError(c, http.StatusNotFound,
			helpers.GetFormattedMessage("empty_notfound_formatted", "User"))
		return
	}

	// Validate input
	var input UpdateUserInput
	if err := c.ShouldBindJSON(&input); err != nil {
		api.RespondError(c, http.StatusBadRequest, err.Error())
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		api.RespondError(c, http.StatusBadRequest, err.Error())
		return
	}

	updatedInput := map[string]interface{}{
		"username": input.Username,
		"password": string(hashedPassword),
		"email":    input.Email,
	}

	if err := models.PutOneUser(&user, updatedInput); err != nil {
		api.RespondError(c, http.StatusBadRequest, err.Error())
	}

	api.RespondSuccess(c, http.StatusOK, user)
}

// DELETE /users/:id
// Delete a user
func DeleteUser(c *gin.Context) {
	// Get model if exist
	var user models.User
	if err := models.GetOneUserId(&user, c.Param("id")); err != nil {
		api.RespondError(c, http.StatusNotFound,
			helpers.GetFormattedMessage("empty_notfound_formatted", "User"))
		return
	}

	if err := models.DeleteUser(&user); err != nil {
		api.RespondError(c, http.StatusBadRequest, err.Error())
	}

	api.RespondSuccess(c, http.StatusOK, true)
}

func VerificationUser(c *gin.Context) {
	var user models.User
	token := c.Param("token")

	m := map[string]interface{}{"email_token": token[1:]}
	if err := models.GetOneUser(&user, m); err != nil {
		api.RespondError(c, http.StatusBadRequest,
			helpers.GetFormattedMessage("empty_notfound_formatted", "Token"))
		return
	}

	updatedInput := map[string]interface{}{"email_token": ""}

	if err := models.PutOneUser(&user, updatedInput); err != nil {
		api.RespondError(c, http.StatusBadRequest, err.Error())
	}

	api.RespondSuccess(c, http.StatusOK, true)
}
