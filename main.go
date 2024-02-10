package main

import (
	"fmt"
	"io"
	"net/http"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	app.Get("/getip/:ip", func(c *fiber.Ctx) error {
	url := fmt.Sprintf("http://ip-api.com/json/%s", c.Params("ip"))
    resp, err := http.Get(url)
    if err != nil {
        return c.Status(http.StatusInternalServerError).SendString(err.Error())
    }
    defer resp.Body.Close()
    // Read the response body
	bodyBytes, err := io.ReadAll(resp.Body)
	bodyString := string(bodyBytes)
    if err != nil {
        return c.Status(http.StatusInternalServerError).SendString(err.Error())
    }
    // Send the response from the external API
    return c.SendString(string(bodyString))
})

	app.Listen(":1111")

}
