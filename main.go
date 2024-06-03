package main

import (
	"context"
	"encoding/json"
	"io"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/spf13/viper"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

// user google information struct
type User struct {
	Email          string `json:"email"`
	Family_name    string `json:"family_name"`
	Given_name     string `json:"given_name"`
	Id             string `json:"id"`
	Locale         string `json:"locale"`
	Name           string `json:"name"`
	Picture        string `json:"picture"`
	Verified_email bool   `json:"verified_email"`
}

func main() {
	app := fiber.New()
	viper := viper.New()
	viper.SetConfigType("env")
	viper.SetConfigName(".env")
	viper.AddConfigPath("./")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal(err)
	}
	oauthConf := &oauth2.Config{ //setup oauth
		ClientID:     viper.GetString("CLIENT_ID"),
		ClientSecret: viper.GetString("CLIENT_SECRET"),
		RedirectURL:  viper.GetString("REDIRECT_URL"),
		Scopes: []string{
			"https://www.googleapis.com/auth/userinfo.email",
			"https://www.googleapis.com/auth/userinfo.profile",
		},
		Endpoint: google.Endpoint,
	}

	app.Get("/oauth/google", func(c *fiber.Ctx) error {
		url := oauthConf.AuthCodeURL("state") //get auth URL
		return c.Redirect(url)                //redirect to google auth url
	})

	app.Get("/oauth/redirect", func(c *fiber.Ctx) error {
		code := c.Query("code") //get code from query params for generating token
		if code == "" {
			return c.Status(fiber.StatusInternalServerError).SendString("Failed to exchange token: " + err.Error())
		}
		token, err := oauthConf.Exchange(context.Background(), code) //get token
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString("Failed to exchange token: " + err.Error())
		}
		client := oauthConf.Client(context.Background(), token)                      //set client for getting user info like email, name, etc.
		response, err := client.Get("https://www.googleapis.com/oauth2/v2/userinfo") //get user info
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString("Failed to get user info: " + err.Error())
		}

		defer response.Body.Close()
		var user User                           //user variable
		bytes, err := io.ReadAll(response.Body) //reading response body from client
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString("Error reading response body: " + err.Error())
		}
		err = json.Unmarshal(bytes, &user) //unmarshal user info
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString("Error unmarshal json body " + err.Error())
		}

		/*
			Do everything you want
			like storing user information into database, etc.....
		*/
		return c.Status(fiber.StatusOK).JSON(user) //return user info

	})

	app.Listen(":8080")

}
