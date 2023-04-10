package controller

import (
	"fmt"
	"log"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
	"github.com/pratiksha/blogbackend/database"
	"github.com/pratiksha/blogbackend/models"
	"github.com/pratiksha/blogbackend/util"
)

func validateEmail(email string) bool {
	emailRegex := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)
	return emailRegex.MatchString(email)
}

func Register(c *fiber.Ctx) error {
	var data map[string]interface{}
	var userData models.User
	if err := c.BodyParser(&data); err != nil {
		fmt.Println("Unable to pass body")
	}
	//check if password is lesser then 6 chracter
	if len(data["password"].(string)) <= 6 {
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": "password must be greater then 6 chracter",
		})
	}
	if !validateEmail(strings.TrimSpace(data["email"].(string))) {
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": "Please give the valid email",
		})
	}

	//check if email alredy exist in db
	database.DB.Where("email=?", strings.TrimSpace(data["email"].(string))).First(&userData)

	if userData.Id != 0 {
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": "Email alrady exist",
		})
	}

	user := models.User{
		FirstName: data["fisrt_name"].(string),
		LastName:  data["last_name"].(string),
		Phone:     data["phone"].(string),
		Email:     strings.TrimSpace(data["email"].(string)),
	}

	user.SetPassword(data["password"].(string))
	err := database.DB.Create(&user)

	if err != nil {
		log.Print("error in saving user", err)
	}

	c.Status(200)
	return c.JSON(fiber.Map{
		"user":    user,
		"message": "User Account Created Successfully",
	})
}

func Login(c *fiber.Ctx) error {
	var data map[string]string
	if err := c.BodyParser(&data); err != nil {
		fmt.Println("Unable to pass body")
	}
	var user models.User
	database.DB.Where("email=?", data["email"]).First(&user)
	if user.Id == 0 {
		c.Status(404)
		return c.JSON(fiber.Map{
			"message": "email id is not exist,kindly create an account",
		})
	}
	if err := user.ComparePassword(data["password"]); err != nil {
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": "incorrect password",
		})
	}
	token, err := util.GenerateJwt(strconv.Itoa(int(user.Id)))
	if err != nil {
		c.Status(fiber.StatusInternalServerError)
		return nil
	}
	cookie := fiber.Cookie{
		Name:     "jwt",
		Value:    token,
		Expires:  time.Now().Add(time.Hour * 24),
		HTTPOnly: true,
	}
	c.Cookie(&cookie)
	c.Status(200)
	return c.JSON(fiber.Map{
		"message": "User Authenticate successful",
		"user":    user,
	})

}

type Claims struct {
	jwt.StandardClaims
}
