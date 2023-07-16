package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/mrKrabsmr/infokeeper/internal/app/dto"
	"github.com/mrKrabsmr/infokeeper/internal/app/services"
	"net/http"
)

// Get @Summary
// @Description get info by key and unique id
// @Tags info-keeper
// @Accept json
// @Produce json
// @Param input query dto.GetInfoDTO true "info"
// @Success 200 {object} dto.Response
// @Failure 400,404,422,500 {object} dto.ErrorResponse
// @Router /api/v1/info-keeper [get]
func Get(c *gin.Context) {
	var dtoInfo dto.GetInfoDTO

	if err := c.ShouldBindQuery(&dtoInfo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   true,
			"message": "the necessary parameters is missing",
		})
		return
	}

	if err := dtoInfo.Validate(); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"error":   true,
			"message": "check again the correctness of the transmitted data",
		})
	}

	service, err := services.NewInfoService()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   true,
			"message": "server error",
		})
		return
	}

	value, err := service.GetValue(&dtoInfo)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error":   true,
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"error":  false,
		"result": value,
	})
}

// Post @Summary
// @Description send the key and the information you want to save and be sure to save the resulting id along with the key
// @Tags info-keeper
// @Accept json
// @Produce json
// @Param input body dto.CreateInfoDTO true "info"
// @Success 201 {object} dto.Response
// @Failure 400,422,500 {object} dto.ErrorResponse
// @Router /api/v1/info-keeper [post]
func Post(c *gin.Context) {
	var dtoInfo *dto.CreateInfoDTO

	if err := c.ShouldBindJSON(&dtoInfo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   true,
			"message": "the necessary data is missing",
		})
		return
	}

	if err := dtoInfo.Validate(); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"error":   true,
			"message": "check again the correctness of the transmitted data",
		})
		return
	}

	service, err := services.NewInfoService()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   true,
			"message": err.Error(),
		})
		return
	}

	ip := c.ClientIP()

	id, err := service.Save(dtoInfo, ip)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   true,
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"error":  false,
		"result": id,
	})
}

// Patch @Summary
// @Description send the request body with the mandatory content of id and key (they are immutable) and the modified/unchanged value and read-only
// @Tags info-keeper
// @Accept json
// @Produce json
// @Param input body dto.UpdateInfoDTO true "info"
// @Success 200 {object} dto.Response
// @Failure 400,404,422,500 {object} dto.ErrorResponse
// @Router /api/v1/info-keeper [patch]
func Patch(c *gin.Context) {
	var dtoInfo *dto.UpdateInfoDTO

	if err := c.ShouldBindJSON(&dtoInfo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   true,
			"message": "the necessary data is missing",
		})
		return
	}

	if err := dtoInfo.Validate(); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"error":   true,
			"message": "check again the correctness of the transmitted data",
		})
		return
	}

	service, err := services.NewInfoService()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   true,
			"message": "server error",
		})
		return
	}

	ip := c.ClientIP()

	if err := service.PartialUpdate(dtoInfo, ip); err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error":   true,
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"error":  false,
		"result": "resource updated successfully",
	})
}

// Delete @Summary
// @Description delete info by key and unique id
// @Tags info-keeper
// @Accept json
// @Produce json
// @Param input body dto.DeleteInfoDTO true "info"
// @Success 200 {object} dto.Response
// @Failure 400,404,422,500 {object} dto.ErrorResponse
// @Router /api/v1/info-keeper [delete]
func Delete(c *gin.Context) {
	var dtoInfo *dto.DeleteInfoDTO

	if err := c.ShouldBindJSON(&dtoInfo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   true,
			"message": "the necessary data is missing",
		})
		return
	}

	if err := dtoInfo.Validate(); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"error":   true,
			"message": "check again the correctness of the transmitted data",
		})
		return
	}

	service, err := services.NewInfoService()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   true,
			"message": "server error",
		})
		return
	}

	if err := service.Delete(dtoInfo); err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error":   true,
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"error":  false,
		"result": "resource deleted successfully",
	})
}

// GetCount @Summary
// @Description send a request and get the amount of information registered on your ip address
// @Tags info-keeper
// @Accept json
// @Produce json
// @Success 200 {object} dto.Response
// @Failure 400,500 {object} dto.ErrorResponse
// @Router /api/v1/info-keeper/count [get]
func GetCount(c *gin.Context) {
	service, err := services.NewInfoService()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   true,
			"message": "server error",
		})
		return
	}

	ip := c.ClientIP()

	count, err := service.GetCountInfo(ip)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   true,
			"message": "error receiving data",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"error":  false,
		"result": count,
	})
}
