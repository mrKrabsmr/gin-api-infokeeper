package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/mrKrabsmr/infokeeper/internal/app/dto"
	"github.com/mrKrabsmr/infokeeper/internal/app/services"
	"net/http"
)

type InfoController struct {
	Service *services.InfoService
}

// Get @Summary
// @Description get info by key and unique id
// @Tags info-keeper
// @Accept json
// @Produce json
// @Param input query dto.NewInfoDTO true "info"
// @Success 200 {object} dto.Response
// @Failure 400,404,422,500 {object} dto.ErrorResponse
// @Router /api/v1/info-keeper [get]
func (r *InfoController) Get(c *gin.Context) {
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

	value, err := r.Service.GetValue(&dtoInfo)
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
func (r *InfoController) Post(c *gin.Context) {
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

	ip := c.ClientIP()

	id, err := r.Service.Save(dtoInfo, ip)
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
func (r *InfoController) Patch(c *gin.Context) {
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

	ip := c.ClientIP()

	if err := r.Service.PartialUpdate(dtoInfo, ip); err != nil {
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
func (r *InfoController) Delete(c *gin.Context) {
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

	if err := r.Service.Delete(dtoInfo); err != nil {
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
func (r *InfoController) GetCount(c *gin.Context) {
	ip := c.ClientIP()

	count, err := r.Service.GetCountInfo(ip)
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
