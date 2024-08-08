package data

import (
	"context"
	"log"
	"os"
	"time"

	"task-manager/models"

	"github.com/golang-jwt/jwt"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)


func HashPassword(password string) (string, error){
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err

	}

	
func CheckHashedPassword(password,hashed string) bool{
	err := bcrypt.CompareHashAndPassword([]byte(hashed), []byte(password))
	return err == nil

}

func GenerateToken(user models.User) (string, error){
	expirationTime := time.Now().Add(24*time.Hour)
	claims := &models.Claim{
		Email: user.Email,
		Role: user.Role,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},

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
	
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte(Secretkey))

	if err != nil {
		return "", err
	}
	return tokenString, err
}
	

func FindUserByEmail(email string) (models.User,error){

	var curUser models.User
	err := userCollection.FindOne(context.TODO(), bson.D{{"email",email}}).Decode(&curUser)

	return curUser, err
}

func CreateUser(user models.User) (models.User, string) {
	_, err := FindUserByEmail(user.Email)
	if err == nil {
		return models.User{}, "user name already taken"
	}
	hashedPassword, err := HashPassword(user.Password)

	if err != nil{
		return models.User{}, "failed to hash the password"
	}

	count, err := userCollection.CountDocuments(context.TODO(), bson.D{})
    if err != nil {
        return models.User{}, "failed to check user count"
    }

    // Assign role based on user count
    if count == 0 {
        user.Role = "admin"
    } else {
        user.Role = "user"
    }

	user.Password = hashedPassword
	user.ID = primitive.NewObjectID()
	_, err = userCollection.InsertOne(context.TODO(),user)
	return user,""
}

func Login(user models.User) (models.User,string){


	existingUser, err := FindUserByEmail(user.Email)
	if err != nil {
		return models.User{}, "User doesn't exist"
	}

	if !CheckHashedPassword(user.Password, existingUser.Password){
		return models.User{}, "invalid user name or password"
	}

	signedToken, error := GenerateToken(existingUser)

	if error != nil{
		return models.User{}, "failed to generate token"
	}

	return existingUser, signedToken

}

func PromoteUser(id string) (models.User, error) { 
	objectID, err := primitive.ObjectIDFromHex(id)
    if err != nil {
        return models.User{}, err
    }

    filter := bson.D{{"_id", objectID}}

    updateFields := bson.M{"role": "admin"}

    update := bson.D{{"$set", updateFields}}

    _, err = userCollection.UpdateOne(context.TODO(), filter, update)
    if err != nil {
        return models.User{}, err
    }

    // Fetch the updated user from the database
    var result models.User
    err = userCollection.FindOne(context.TODO(), filter).Decode(&result)
    if err != nil {
        return models.User{}, err
    }

    return result, nil
	
}