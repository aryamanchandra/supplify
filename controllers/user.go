package controllers

import (
	// "context"
	// "fmt"
	// "log"
	// "strconv"
	// "net/http"
	// "time"
	// "github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	// helper "github.com/aryamanchandra/supplify/helpers"
	// "github.com/aryamanchandra/supplify/models"
	database "github.com/aryamanchandra/supplify/database"
	"go.mongodb.org/mongo-driver/mongo"
	// "golang.org/x/crypto/bcrypt"
)

var userCollection *mongo.Collection = database.OpenCollection(database.Client, "user")
var validate = validator.New()

func HashPassword()
func VerifyPassword()
func Signup()
func Login()
func GetUsers()
func GetUser()
