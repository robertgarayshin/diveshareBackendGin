package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
	"strconv"
)

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
