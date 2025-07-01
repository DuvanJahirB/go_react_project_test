package middleware

import (
    "net/http"
    "strings"
    "os"
    "log"
    "github.com/dgrijalva/jwt-go"
    "github.com/gin-gonic/gin"
)

var jwtKey []byte

func init() {
    key := os.Getenv("JWT_SECRET_KEY")
    if key == "" {
        log.Fatal("JWT_SECRET_KEY environment variable not set")
    }
    jwtKey = []byte(key)
}

// AuthMiddleware es un middleware de Gin para la autenticación basada en JWT.
// Verifica la presencia de un token en la cabecera Authorization, lo valida y, si es válido,
// extrae el email del usuario y lo añade al contexto de la solicitud.
func AuthMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        // Obtiene la cabecera Authorization.
        authHeader := c.GetHeader("Authorization")
        // Comprueba si la cabecera está vacía o no tiene el prefijo "Bearer ".
        if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "Token missing"})
            c.Abort()
            return
        }

        // Extrae el token de la cabecera.
        tokenString := strings.TrimPrefix(authHeader, "Bearer ")

        // Parsea y valida el token.
        token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
            // La clave secreta debe ser la misma que se usó para firmar el token.
            return jwtKey, nil
        })

        // Si hay un error en el parseo o el token no es válido, devuelve un error.
        if err != nil || !token.Valid {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
            c.Abort()
            return
        }

        // Extrae la informacion del jwt y obtiene email
        claims := token.Claims.(jwt.MapClaims)
        c.Set("email", claims["email"])
        // Continua con el siguiente manejador en la cadena.
        c.Next()
    }
}
