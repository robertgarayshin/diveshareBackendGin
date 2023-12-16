package handler

import (
	"diveshareBackendGin/model"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
	"strconv"
)

// GetUserById
//
// @Summary Get User By Id
// @Tags users
// @Description get user by id
// @ID get-user-by-id
// @Accept  json
// @Produce  json
// @Param 	id path int true "id"
// @Success 200 {object} model.User
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /user/{id} [get]
func (h *Handler) GetUserById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		logrus.Print(err)
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}
	user, err := h.services.User.GetById(id)
	if err != nil {
		return
	}
	c.JSON(http.StatusOK, user)
}

// GetAllUsers
//
// @Summary Get All Users
// @Tags users
// @Description get all users
// @ID get-users
// @Accept  json
// @Produce  json
// @Success 200 {array} model.User
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /user/ [get]
func (h *Handler) GetAllUsers(c *gin.Context) {
	users, err := h.services.User.GetAll()
	if err != nil {
		logrus.Print(err)
		newErrorResponse(c, http.StatusInternalServerError, "internal error")
		return
	}
	c.JSON(http.StatusOK, users)
}

// EditProfile
//
// @Summary Edit Profile
// @Tags users
// @Description edit profile
// @ID edit-profile
// @Accept  json
// @Produce  json
// @Param 	id path int true "id"
// @Param input body model.UpdateUserInput true "account info"
// @Success 200 {object} statusResponse
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /user/{id} [put]
func (h *Handler) EditProfile(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}
	var input model.UpdateUserInput
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.services.User.Update(id, input); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, statusResponse{"ok"})
}

// DeleteProfile
//
// @Summary Delete User
// @Tags users
// @Description delete user by id
// @ID delete-user-by-id
// @Accept  json
// @Produce  json
// @Param 	id path int true "id"
// @Success 200 {object} statusResponse
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /user/{id} [delete]
func (h *Handler) DeleteProfile(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}
	if err := h.services.User.Delete(id); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, statusResponse{"ok"})
}
