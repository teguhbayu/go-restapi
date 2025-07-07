package controllers

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/teguhbayu/go-restapi/initializers"
	"github.com/teguhbayu/go-restapi/models"
	"github.com/teguhbayu/go-restapi/utils"
	"gorm.io/gorm/clause"
)

type pathParams struct {
	ID string `uri:"id"`
}

type Body struct {
	Title         string
	Year_released int
	Synopsis      string
}

func CreateBook(c *gin.Context) {
	var body Body
	c.Bind(&body)

	book := models.Books{
		Title:        body.Title,
		YearReleased: body.Year_released,
		Synopsis:     body.Synopsis,
	}

	result := initializers.DB.Create(&book)

	if result.Error != nil {
		log.Fatal("Failed to create!")
		utils.Response(c, nil, "Failed to create book!", http.StatusBadRequest)
		return
	}

	utils.Response(c, book, "Created a book!", 200)
}

func GetBooks(c *gin.Context) {
	var books []models.Books
	res := initializers.DB.Find(&books)

	if res.Error != nil {
		log.Fatal("Failed to create!")
		utils.Response(c, nil, "err!", http.StatusBadRequest)
		return
	}

	utils.Response(c, books, "Found Books!", 200)
}

func GetBookById(c *gin.Context) {
	var params pathParams
	c.ShouldBind(&params)

	book := models.Books{}
	initializers.DB.First(&book, params.ID)

	utils.Response(c, book, "Found Book!", 200)
}

func UpdateBook(c *gin.Context) {
	var body Body
	id := c.Param("id")

	c.Bind(&body)

	var book models.Books

	initializers.DB.Model(&book).Clauses(clause.Returning{}).Where("id = ?", id).Updates(models.Books{
		UpdatedAt:    time.Now(),
		Title:        body.Title,
		YearReleased: body.Year_released,
		Synopsis:     body.Synopsis,
	})

	utils.Response(c, book, "Updated Book!", 200)
}

func DeleteBook(c *gin.Context) {
	id := c.Param("id")

	initializers.DB.Where("id = ?", id).Delete(&models.Books{})
	utils.Response(c, nil, "Deleted Book!", 200)
}
