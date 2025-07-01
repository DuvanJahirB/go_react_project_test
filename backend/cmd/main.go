package main

import (
    "context"
    "log"
    "os"
    "time"

    "backend/internal/handlers"
    "backend/internal/middleware"
    "backend/internal/models"
    "backend/internal/utils"
    "github.com/gin-contrib/cors"
    "github.com/gin-gonic/gin"
    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
)


func main() {
    mongoURI := os.Getenv("MONGO_URI")
    if mongoURI == "" {
        mongoURI = "mongodb://mongo:27017"
    }

    client, err := mongo.NewClient(options.Client().ApplyURI(mongoURI))
    if err != nil {
        log.Fatal(err)
    }
    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()
    err = client.Connect(ctx)
    if err != nil {
        log.Fatal(err)
    }

    handlers.UserCollection = client.Database("mydb").Collection("users")

    seedDatabase(handlers.UserCollection)

    router := gin.Default()

    // Configuración de CORS
    router.Use(cors.New(cors.Config{
        AllowOrigins:     []string{"*"}, // Permite todos los orígenes para desarrollo
        AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD", "OPTIONS"},
        AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type", "Authorization"},
        ExposeHeaders:    []string{"Content-Length"},
        AllowCredentials: true,
        MaxAge:           12 * time.Hour,
    }))

    router.POST("/user", handlers.RegisterUser)
    router.POST("/auth/login", handlers.LoginUser)

    protected := router.Group("/")
    protected.Use(middleware.AuthMiddleware())
    protected.GET("/profile", handlers.GetProfile)

    

    port := os.Getenv("PORT")
    if port == "" {
        port = "8080"
    }
    router.Run(":" + port)
}

func seedDatabase(userCollection *mongo.Collection) {
    // Valores por defecto
    name  := "Administrador"
    email  := "admin@example.com"
    password := "admin123"

    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()

    // Comprueba si el usuario por defecto ya existe.
    var existing models.User
    err := userCollection.FindOne(ctx, bson.M{"email": email}).Decode(&existing)
    if err == nil {
        return
    }

    // Si el usuario no existe, se crea
    if err == mongo.ErrNoDocuments {
        hashedPassword, err := utils.HashPassword(password)
        if err != nil {
            log.Fatalf("Error al hashear la contraseña por defecto: %v", err)
        }

        defaultUser := models.User{
            Name:     name,
            Email:    email,
            Password: hashedPassword,
        }

        _, err = userCollection.InsertOne(context.Background(), defaultUser)
        if err != nil {
            log.Fatalf("Error al insertar el usuario por defecto: %v", err)
        }
        log.Println("Usuario por defecto creado.")
    } else if err != nil {
        log.Fatalf("Error al comprobar el usuario por defecto: %v", err)
    }
}