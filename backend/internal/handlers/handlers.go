package handlers

import (
    "context"
    "log"
    "net/http"
    "os"
    "time"

    "backend/internal/models"
    "backend/internal/utils"
    "github.com/gin-gonic/gin"
    "github.com/go-playground/validator/v10"
    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/mongo"
)

var validate = validator.New()

var UserCollection *mongo.Collection
var jwtKey []byte

func init() {
    key := os.Getenv("JWT_SECRET_KEY")
    if key == "" {
        log.Fatal("JWT_SECRET_KEY environment variable not set")
    }
    jwtKey = []byte(key)
}


// RegisterUser godoc
// @Summary Register a new user
// @Description Register a new user with name, email, and password
// @Tags users
// @Accept json
// @Produce json
// @Param user body models.User true "User registration data"
// @Success 201 {object} models.MessageResponse
// @Failure 400 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
// @Router /user [post]
func RegisterUser(c *gin.Context) {
    var user models.User
    if err := c.BindJSON(&user); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
        return
    }

    if err := validate.Struct(user); err != nil {
        errorMessage := utils.FormatValidationError(err)
        c.JSON(http.StatusBadRequest, gin.H{"error": errorMessage})
        return
    }

    var existing models.User
    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()
    err := UserCollection.FindOne(ctx, bson.M{"email": user.Email}).Decode(&existing)
    if err == nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Email already registered"})
        return
    }

    hashedPassword, err := utils.HashPassword(user.Password)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Error hashing password"})
        return
    }

    user.Password = hashedPassword

    _, err = UserCollection.InsertOne(context.Background(), user)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Error saving user"})
        return
    }

    c.JSON(http.StatusCreated, gin.H{"message": "User registered"})
}


// LoginUser godoc
// @Summary Log in a user
// @Description Authenticate user and return a JWT token
// @Tags auth
// @Accept json
// @Produce json
// @Param loginData body models.LoginRequest true "User login credentials"
// @Success 200 {object} models.TokenResponse
// @Failure 400 {object} models.ErrorResponse
// @Failure 401 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
// @Router /auth/login [post]
func LoginUser(c *gin.Context) {

    var loginData struct {
        Email    string `json:"email" validate:"required,email"`
        Password string `json:"password" validate:"required"`
    }
    
    if err := c.BindJSON(&loginData); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
        return
    }

    if err := validate.Struct(loginData); err != nil {
        errorMessage := utils.FormatValidationError(err)
        c.JSON(http.StatusBadRequest, gin.H{"error": errorMessage})
        return
    }

    var found models.User
    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()

    err := UserCollection.FindOne(ctx, bson.M{"email": loginData.Email}).Decode(&found)
    if err != nil {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
        return
    }

    if !utils.CheckPasswordHash(loginData.Password, found.Password) {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
        return
    }

    token, err := utils.GenerateToken(found.Email)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not generate token"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"token": token})
}


// GetProfile godoc
// @Summary Get user profile
// @Description Get the profile information of the authenticated user
// @Tags users
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Success 200 {object} models.ProfileResponse
// @Failure 401 {object} models.ErrorResponse
// @Failure 404 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
// @Router /profile [get]
func GetProfile(c *gin.Context) {
    email, exists := c.Get("email")
    if !exists {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Email not found in token"})
        return
    }

    var user models.User
    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()

    err := UserCollection.FindOne(ctx, bson.M{"email": email}).Decode(&user)
    if err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
        return
    }

    c.JSON(http.StatusOK, gin.H{
        "email": user.Email,
        "name":  user.Name,
    })
}