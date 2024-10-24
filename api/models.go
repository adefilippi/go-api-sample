package api

import (
	"encoding/json"
	"example/web-service-gin/entity"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/datatypes"
	"net/http"

	"example/web-service-gin/repository"
	"example/web-service-gin/service/utils"
)

//	@Summary		Show all models
//	@Description	get all models
//	@Tags			Model
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	entity.Model
//	@Failure		400	{object}	ApiError
//	@Failure		404	{object}	ApiError
//	@Failure		500	{object}	ApiError
//	@Router			/models [get]
func GetModels(c *gin.Context) {
	models, err := repository.GetAllModels()
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}
	c.IndentedJSON(http.StatusOK, models)
}

//	@Summary		Show an account
//	@Description	get string by ID
//	@Tags			Model
//	@Accept			json
//	@Produce		json
//	@Param			id	path		string	true	"Model ID"
//	@Success		200	{object}	entity.Model
//	@Failure		400	{object}	ApiError
//	@Failure		404	{object}	ApiError
//	@Failure		500	{object}	ApiError
//	@Router			/models/{id} [get]
func GetModelById(c *gin.Context) {
	uuid, error := uuid.Parse(c.Param("id"))
	if error != nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}
	model, error := repository.GetModelById(uuid)
	c.IndentedJSON(http.StatusOK, model)
}

//	@Summary		Add a model
//	@Description	Add by json model
//	@Tags			Model
//	@Accept			json
//	@Produce		json
//	@Param			model	body		entity.Model	true	"Add account"
//	@Success		200		{object}	entity.Model
//	@Failure		400		{object}	ApiError
//	@Failure		404		{object}	ApiError
//	@Failure		500		{object}	ApiError
//	@Router			/models [post]
func CreateModel(c *gin.Context) {
	var model entity.Model
	c.BindJSON(&model)
	result, error := repository.CreateModel(model)
	if error != nil {

		code, messages := HandleError(error)
		c.IndentedJSON(code, messages)
	} else {
		c.IndentedJSON(http.StatusCreated, result)
	}
}

//	@Summary		Update a Model
//	@Description	Update by json Model
//	@Tags			Model
//	@Accept			json
//	@Produce		json
//	@Param			id		path		string					true	"Model ID"
//	@Param			model	body		entity.Model	true	"Update model"
//	@Success		200		{object}	entity.Model
//	@Failure		400		{object}	ApiError
//	@Failure		404		{object}	ApiError
//	@Failure		500		{object}	ApiError
//	@Router			/models/{id} [patch]
func UpdateModel(c *gin.Context) {
	var model entity.Model
	c.BindJSON(&model)
	result, error := repository.UpdateModel(model)
	if error != nil {
		code, messages := HandleError(error)
		c.IndentedJSON(code, messages)
	} else {
		c.IndentedJSON(http.StatusOK, result)
	}
}

//	@Summary		Delete a model
//	@Description	Delete by model ID
//	@Tags			Model
//	@Accept			json
//	@Produce		json
//	@Param			id	path	string	true	"Model ID"
//	@Success		204	{object}	entity.Model
//	@Failure		400	{object}	ApiError
//	@Failure		404	{object}	ApiError
//	@Failure		500	{object}	ApiError
//	@Router			/models/{id} [delete]
func DeleteModel(c *gin.Context) {
	uuid, error := uuid.Parse(c.Param("id"))
	if error != nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}
	error = repository.DeleteModel(uuid)
	if error != nil {
		code, messages := HandleError(error)
		c.IndentedJSON(code, messages)
	} else {
		c.AbortWithStatus(http.StatusNoContent)
	}
}

func GetMdelImage(c *gin.Context) {
	modelId, error := uuid.Parse(c.Param("id"))
	if error != nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}
	model, error := repository.GetModelById(modelId)
	if model.ID == nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}

	c.IndentedJSON(http.StatusCreated, model)
}

func CreateModelImage(c *gin.Context) {
	modelId, error := uuid.Parse(c.Param("id"))
	if error != nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}
	model, error := repository.GetModelById(modelId)
	if model.ID == nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}

	filepath, filename := utils.HandleFile(c)
	mimetype, _ := utils.GetFileMimeType(filepath)
	size, _ := utils.GetFileSize(filepath)

	mo := entity.MediaObject{
		ModelID:      *model.ID,
		Name:         &filename,
		OriginalName: &filename,
		MimeType:     &mimetype,
		Size:         &size,
	}
	if utils.IsImageMimeType(mimetype) {
		width, height, err := utils.GetImageDimensions(filepath)
		dimensions := [2]int{width, height}
		if err == nil {
			dimensionsJSON, err := json.Marshal(dimensions)
			if err == nil {
				mo.Dimensions = datatypes.JSON(dimensionsJSON)
			}
		}
	}

	mo, error = repository.CreateMediaObject(mo)
	if error != nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}
	c.IndentedJSON(http.StatusCreated, mo)
}

func DeleteModelImage(c *gin.Context) {

}
