package handlers

import (
    "github.com/gin-gonic/gin"
    "github.com/go-self-learning/types"
    "gorm.io/gorm"
)

type PostController struct {
    DB *gorm.DB // Add a field to hold the database instance
}

func NewPostController(db *gorm.DB) PostController {
    return PostController{DB: db}
}

// @Summary Create a new post
// @Description Create a new post.
// @Accept json
// @Param post body CreatePostRequest true "Post data"
// @Success 200 {object} PostResponse
// @Router /posts [post]
func (u *PostController) CreatePost(c *gin.Context) {
    var post types.Posts
    c.BindJSON(&post)
    u.DB.Create(&post)
    c.JSON(200, gin.H{
        "post": post,
    })
}

type PostResponse struct {
	ID     uint   `json:"id"`
	Title  string `json:"title"`
    Body   string `json:"body"`
    UserID uint   `json:"user_id"`
}

// @Summary Get all posts
// @Description Get all posts.
// @Accept json
// @Success 200 {array} PostResponse
// @Router /posts [get]
func (u *PostController) GetPosts(c *gin.Context) {
    var posts []types.Posts
    u.DB.Find(&posts)
    c.JSON(200, gin.H{
        "posts": posts,
    })
}

// @Summary Get a post
// @Description Get a post by ID.
// @Accept json
// @Param id path int true "Post ID"
// @Success 200 {object} PostResponse
// @Router /posts/{id} [get]
func (u *PostController) GetPost(c *gin.Context) {
    var post types.Posts
    id := c.Param("id")
    u.DB.First(&post, id)
    c.JSON(200, gin.H{
        "post": post,
    })
}

// @Summary Update a post
// @Description Update a post by ID.
// @Accept json
// @Param id path int true "Post ID"
// @Param post body UpdatePostRequest true "Post object"
// @Success 200 {object} PostResponse
// @Router /posts/{id} [put]
func (u *PostController) UpdatePost(c *gin.Context) {
    var post types.Posts
    id := c.Param("id")
    u.DB.First(&post, id)
    c.BindJSON(&post)
    u.DB.Save(&post)
    c.JSON(200, gin.H{
        "post": post,
    })
}

// @Summary Delete a post
// @Description Delete a post by ID.
// @Accept json
// @Param id path int true "Post ID"
// @Success 200 {object} PostResponse
// @Router /posts/{id} [delete]
func (u *PostController) DeletePost(c *gin.Context) {
    var post types.Posts
    id := c.Param("id")
    u.DB.Delete(&post, id)
    c.JSON(200, gin.H{
        "post": post,
    })
}

type CreatePostRequest struct {
    Title  string `json:"title"`
    Body   string `json:"body"`
    UserID uint   `json:"user_id"`
}

type UpdatePostRequest struct {
    Title string `json:"title"`
    Body  string `json:"body"`
}
