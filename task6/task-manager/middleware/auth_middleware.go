package middleware

import (
	"log"
	"net/http"
	"os"
	"strings"
	"task-manager/models"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/joho/godotenv"
)


func AuthMiddleware(ctx *gin.Context) {
		auth := ctx.GetHeader("Authorization")
		if auth == ""{
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message":"missing token"})
			return
		}

		authParts := strings.Split(auth, " ")
		if len(authParts) != 2 || authParts[0] != "bearer" {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message":"invalid authorization token"})
			return

	}
	// Load the .env file
    err := godotenv.Load()
    if err != nil {
        log.Fatal("Error loading .env file")
    }

    // Get the MongoDB URI from environment variables
    Secretkey := os.Getenv("SECRET_KEY")
    if Secretkey == "" {
        log.Fatal("SECRET_KEY not set in environment")
    }
	claims := &models.Claim{}
	token, err := jwt.ParseWithClaims(authParts[1], claims, func(t *jwt.Token) (interface{}, error) {

			return []byte(Secretkey),nil
		})
	if err != nil || !token.Valid{
		ctx.AbortWithStatusJSON(http.StatusForbidden, gin.H{"message":"Unauthorized"})
		return
	}
	ctx.Set("name",claims.Email)
	ctx.Set("role", claims.Role)
	ctx.Next()
	}

	

func AdminMidleware(ctx *gin.Context){
	role, exists := ctx.Get("role")
	if !exists || role != "admin" {
		ctx.AbortWithStatusJSON(http.StatusForbidden, gin.H{"message":"Unauthorized"})
		return
	}
	ctx.Next()
}