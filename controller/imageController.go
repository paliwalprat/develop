package controller

import (
	"fmt"
	"math/rand"

	"github.com/gofiber/fiber/v2"
)

var letters = []rune("abcdefghijklmnopqrstuvwxyz")

func randLetter(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

func Upload(c *fiber.Ctx) error {
	form, err := c.MultipartForm()
	if err != nil {
		return err
	}
	files := form.File["image"]
	fileName := ""

	for _, file := range files {

		fileName = randLetter(5) + "-" + file.Filename

		fmt.Println(fileName, "fileNamefileNamefileNamefileNamefileName")

		if err := c.SaveFile(file, "./upload/"+fileName); err != nil {
			fmt.Println(err, "errerrerrerrerrerrerrerrerr")
			return nil
		}
	}
	return c.JSON(fiber.Map{
		"url":     "http://localhsot:3000/api/upload/" + fileName,
		"message": "upload done",
	})
}
