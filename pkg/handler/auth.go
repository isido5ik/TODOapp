package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/isido5ik/todo-app"
)

// @Summary SignUp
// @Tags auth
// @Description create account
// @ID create-account
// @Accept  json
// @Produce  json
// @Param input body todo.User true "account info"
// @Success 200 {integer} integer 1
// @Failure 400,404 {newErrorResponse} errorResponse
// @Failure 500 {newErrorResponse} errorResponse
// @Failure default {newErrorResponse} errorResponse
// @Router /auth/sign-up [post]
func (h *Handler) signUp(c *gin.Context) {

	var input todo.User
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.services.Authorization.CreateUser(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})

}

type signInInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// @Summary SignIn
// @Tags auth
// @Description login
// @ID login
// @Accept  json
// @Produce  json
// @Param input body signInInput true "credentials"
// @Success 200 {string} string "token"
// @Router /auth/sign-in [post]
func (h *Handler) signIn(c *gin.Context) {
	var input signInInput
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	token, err := h.services.Authorization.GenerateToken(input.Username, input.Password)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"token": token,
	})
}

// func (h *Handler) getAllUsers(c *gin.Context) {
// 	users, err := h.services.GetUsers()
// 	if err != nil {
// 		newErrorResponse(c, http.StatusInternalServerError, err.Error())
// 	}
// 	c.JSON(http.StatusOK, map[string]interface{}{
// 		"users": users,
// 	})

// }
