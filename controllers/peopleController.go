package controllers

import (
	"Practical_1/initializers"
	"Practical_1/models"

	"github.com/gin-gonic/gin"
)

func PostsCreate(c *gin.Context) {
	// Get data off req body
	var body struct {
		id         string
		First_Name string
		Last_Name  string
		Email      string
		Phone      string
	}

	c.Bind(&body)

	// Create a post
	people := models.People{Id: "", First_Name: body.First_Name, Last_Name: body.Last_Name, Email: body.Email, Phone: body.Phone}
	result := initializers.DB.Create(&people)

	if result.Error != nil {
		c.String(400, "Error creating post ", people)
		return
	}

	// Return it
	c.JSON(200, gin.H{
		"people": people,
	})
}

func PeopleAll(c *gin.Context) {
	// Get the post
	var people []models.People
	initializers.DB.Find(&people)

	// Respond with them
	c.JSON(200, gin.H{
		"people": people,
	})
}

func PeopleIndex(c *gin.Context) {
	// Get id off url
	id := c.Param("id")

	// Get the post
	var people models.People
	peoplde := initializers.DB.First(&people, "id = ?", id)

	if peoplde.Error != nil {
		c.String(200, "Không có people id này")
		return
	}

	// Respod with them
	c.JSON(200, gin.H{
		"people": people,
	})
}

func PeopleUpdate(c *gin.Context) {
	// Get the id off the url
	id := c.Param("id")

	// Get the data off req body
	var body struct {
		First_Name string
		Last_Name  string
		Email      string
		Phone      string
	}

	c.Bind(&body)

	// Lấy models people
	var people models.People
	initializers.DB.First(&people, "id = ?", id)

	// Update it
	initializers.DB.Model(&people).Updates(models.People{
		First_Name: body.First_Name,
		Last_Name:  body.Last_Name,
		Email:      body.Email,
		Phone:      body.Phone,
	})

	// Respod with it
	c.JSON(200, gin.H{
		"people": people,
	})
}

func PeopleDelete(c *gin.Context) {
	// Get the id off the url
	id := c.Param("id")

	// Delete the posts
	initializers.DB.Delete(&models.People{}, "id = ?", id)

	// Respond
	c.String(200, "Delete people success")
}
