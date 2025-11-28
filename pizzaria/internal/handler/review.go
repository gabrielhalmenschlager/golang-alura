package handler

import (
	"net/http"
	"strconv"

	"github.com/gabrielhalmenschlager/curso-golang-alura/pizzaria/internal/data"
	"github.com/gabrielhalmenschlager/curso-golang-alura/pizzaria/internal/models"
	"github.com/gabrielhalmenschlager/curso-golang-alura/pizzaria/internal/service"
	"github.com/gin-gonic/gin"
)

func PostReview(c *gin.Context) {
	pizzaID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var newReview models.Review
	if err := c.ShouldBindJSON(&newReview); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := service.ValidateReviewRating(newReview); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	for i, p := range data.Pizzas {
		if p.ID == pizzaID {
			p.Review = append(p.Review, newReview)
			data.Pizzas[i] = p
			data.SavePizza()
			c.JSON(http.StatusCreated, p)
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"message": "Pizza Not Found"})
}
