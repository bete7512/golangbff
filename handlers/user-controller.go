package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/go-self-learning/types"
	"gorm.io/gorm"
)

type UserController struct {
	DB *gorm.DB // Add a field to hold the database instance
}

func NewUserController(db *gorm.DB) UserController {
	return UserController{DB: db}
}

// @Summary Create a new user
// @Description Create a new user.
// @Accept json
// @Param user body CreateUser true "User data"
// @Success 200 {object} UserResponse
// @Router /users [post]
func (u *UserController) CreateUser(c *gin.Context) {
	var user types.Users
	c.BindJSON(&user)
	u.DB.Create(&user)
	c.JSON(200, gin.H{
		"user": user,
	})
}

type UserResponse struct {
	ID        uint   `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Age       int    `json:"age"`
}
type CreateUserRequest struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Age       int    `json:"age"`
}
type CreateUser struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Age       int    `json:"age"`
}
type UpdateUserRequest struct {
	ID        uint   `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Age       int    `json:"age"`
}

// @Summary Get all users
// @Description Get all users.
// @Accept json
// @Success 200 {array} UserResponse
// @Router /users [get]
func (u *UserController) GetUsers(c *gin.Context) {
	var users []types.Users
	u.DB.Find(&users)
	c.JSON(200, gin.H{
		"users": users,
	})
}

// @Summary Get a user
// @Description Get a user by ID.
// @Accept json
// @Param id path int true "User ID"
// @Success 200 {object} UserResponse
// @Router /users/{id} [get]
func (u *UserController) GetUser(c *gin.Context) {
	var user types.Users
	id := c.Param("id")
	u.DB.First(&user, id)
	c.JSON(200, gin.H{
		"user": user,
	})
}

// @Summary Update a user
// @Description Update a user by ID.
// @Accept json
// @Param id path int true "User ID" // User ID to update
// @Param user body UpdateUserRequest true "User object" // User object to update
// @Success 200 {object} UserResponse
// @Router /users/{id} [put]
func (u *UserController) UpdateUser(c *gin.Context) {
	var user types.Users
	id := c.Param("id")
	u.DB.First(&user, id)
	c.BindJSON(&user)
	u.DB.Save(&user)
	c.JSON(200, gin.H{
		"user": user,
	})
}

// @Summary Delete a user
// @Description Delete a user by ID.
// @Accept json
// @Param id path int true "User ID" // User ID to delete
// @Success 200 {object} UserResponse
// @Router /users/{id} [delete]
func (u *UserController) DeleteUser(c *gin.Context) {
	var user types.Users
	id := c.Param("id")
	u.DB.Delete(&user, id)
	c.JSON(200, gin.H{
		"user": user,
	})
}
