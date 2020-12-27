package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/ultra-utsav/Book-Rentals/backend/api/models"
	repo "github.com/ultra-utsav/Book-Rentals/backend/api/repositories"
	"go.mongodb.org/mongo-driver/mongo"
	"net/http"
)

//CartController struct
type CartController struct {
	cartRepo *repo.CartRepository
}

//NewCartController new user controller
func NewCartController(db *mongo.Database) *CartController {
	return &CartController{
		cartRepo: repo.GetCartRepository(db),
	}
}

//FetchCart fetch cart of a user
func (u *CartController) FetchCart(ctx *gin.Context) {
	userID := ctx.Params.ByName("userID")

	if userID == "" {
		ctx.JSON(http.StatusNotFound, gin.H{"message": "User ID not found"})
	}

	cart, err := u.cartRepo.FetchCart(userID)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Unable to fetch cart!", "error": err.Error()})
	} else {
		ctx.JSON(http.StatusOK, gin.H{"cart": cart, "message": "Cart Fetched Successfully"})
	}
}

//AddBookToCart add book to the cart
func (u *CartController) AddBookToCart(ctx *gin.Context) {
	var books models.Books
	_ = ctx.BindJSON(&books)

	if books.UserID == "" {
		ctx.JSON(http.StatusNotFound, gin.H{"message": "User ID not found"})
		return
	}

	_, err := u.cartRepo.AddBook(books)

	if err != nil && err != mongo.ErrNoDocuments {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Unable to add book to the cart!", "error": err.Error()})
	} else {
		ctx.JSON(http.StatusOK, gin.H{"message": "Book added Successfully to the cart"})
	}
}

//RemoveBookFromCart add book to the cart
func (u *CartController) RemoveBookFromCart(ctx *gin.Context) {
	var books models.Books
	_ = ctx.BindJSON(&books)

	if books.UserID == "" {
		ctx.JSON(http.StatusNotFound, gin.H{"message": "User ID not found"})
		return
	}

	err := u.cartRepo.RemoveBook(books)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Unable to remove book from the cart!", "error": err.Error()})
	} else {
		ctx.JSON(http.StatusOK, gin.H{"message": "Book removed Successfully from the cart"})
	}
}
