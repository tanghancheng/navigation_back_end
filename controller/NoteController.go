package controller

import (
	"Navigation-Web/models"
	"Navigation-Web/models/dto"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type noteController struct {
}

var NoteController = new(noteController)

func (noteController *noteController) GetOne(c *gin.Context) {
	id, ok := c.Params.Get("id")
	if !ok {
		c.JSON(http.StatusInternalServerError, "参数无效")
	}
	idInt, _ := strconv.Atoi(id)
	note, err := models.NoteFunc.GetOne(idInt)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
	} else {
		c.JSON(http.StatusOK, note)
	}
}

func (noteController *noteController) GetAll(c *gin.Context) {
	pageInfo := dto.NewPageInfo()
	c.ShouldBind(pageInfo)
	page, err := models.NoteFunc.GetAll(pageInfo)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
	} else {
		c.JSON(http.StatusOK, page)
	}
}
func (noteController *noteController) GetListByQueryDto(c *gin.Context) {
	noteDto := dto.NewNoteNoteQueryDto()
	// c.BindJSON(&noteDto)
	c.ShouldBind(noteDto)
	page, err := models.NoteFunc.GetListByQueryDto(noteDto)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
	} else {
		c.JSON(http.StatusOK, page)
	}
}

func (noteController *noteController) Create(c *gin.Context) {
	var note models.Note
	c.BindJSON(&note)
	err := models.NoteFunc.Create(&note)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
	} else {
		c.JSON(http.StatusOK, "success")
	}
}

func (noteController *noteController) Update(c *gin.Context) {
	id, ok := c.Params.Get("id")
	if !ok {
		c.JSON(http.StatusInternalServerError, "参数无效")
	}
	idInt, _ := strconv.Atoi(id)
	note, err := models.NoteFunc.GetOne(idInt)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
	}
	c.BindJSON(&note)
	err = models.NoteFunc.Update(note)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
	} else {
		c.JSON(http.StatusOK, "success")
	}
}

func (noteController *noteController) Delete(c *gin.Context) {
	id, ok := c.Params.Get("id")
	if !ok {
		c.JSON(http.StatusInternalServerError, "参数无效")
	}
	idInt, _ := strconv.Atoi(id)
	note, err := models.NoteFunc.GetOne(idInt)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
	}
	note.DeleteStatus = 1
	err = models.NoteFunc.Update(note)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
	} else {
		c.JSON(http.StatusOK, "success")
	}
}
