package repositories

import (
	"net/http"

	"example.com/domain/repositories"
	"example.com/services"
	"example.com/utils/errors"
	"github.com/gin-gonic/gin"
)

func CreateRepo(c *gin.Context) {
	var request repositories.CreateRepoRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		apiErr := errors.NewBadRequestError("invalid json body")
		c.JSON(apiErr.Status(), apiErr)
		return
	}

	result, err := services.RepositoryService.CreateRepo(request)

	if err != nil {
		c.JSON(err.Status(), err)
		return
	}

	c.JSON(http.StatusCreated, result)
}
