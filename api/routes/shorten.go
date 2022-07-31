package routes

import (
	"net/http"
	"time"
	"github.com/F4r5h4d/go-url-shortener/database"
	"github.com/gofiber/fiber/v2"
)

type request struct {
	URL 		        string				`json:"url"`
	CustomShort		    string				`json:"short"`
	Expiry			    time.Duration		`json:"expiry"`
}

type response struct {
	URL				    string				`json:"url"`	
	CustomShort			string				`json:"short"`
	Expiry				time.Duration		`json:"expiry"`
	XRateRemaning		int					`json:"xrate_remaining"`
	XRateLInitRest	    time.Duration		`json:"xrate_limit_init_rest"`
}


func ShortenURL(c *fiber.Ctx) error {
	body := new(request)

	if err := c.BodyParser(body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{ "error": "can't parse JSON" })
	}

	// Implement rate limiting


	
	// check if the input is actually a valid URL
	if ! govalicator.IsURL(body.URL) {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{ "error": "invalid URL" })
	}


	// check for domain errors

	if ! helpers.RemoveDomainError(body.URL) {
		return c.Status(fiber.StatusServiceUnavailable).JSON(fiber.Map{ "error": "You can't access it" })
	}

	// enforce https, ssl, tls

	body.URL = helpers.EnforceHTTPS(body.URL)

}