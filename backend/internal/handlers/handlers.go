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


func RegisterUser(c *gin.Context) {
    var user models.User
    if err := c.BindJSON(&user); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
        return
    }

    if err := validate.Struct(user); err != nil {
        errors := utils.FormatValidationError(err)
        c.JSON(http.StatusBadRequest, gin.H{"errors": errors})
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
        errors := utils.FormatValidationError(err)
        c.JSON(http.StatusBadRequest, gin.H{"errors": errors})
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