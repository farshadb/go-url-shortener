package routes

import (
	"go-url-shortener/database"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/F4r5h4d/go-url-shortener/database"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/limiter"
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
	r2 := database.CreateClient(1)
	defer r2.Close()
	val,err := r2.Get(database.Ctx, c.IP()).Result()
	if err != nil {
		_ = r2.Set(database.Ctx, c.IP(), os.Getenv("API_QUOTA"), 30*60*time.Second).Err()	
	} else {
		val, _ = r2.Get(database.Ctx, c.IP()).Result()
		valInt, _ := strconv.Atoi(val)
		if valInt <= 0 {
			limit, _ := r2.TTL(database.Ctx, c.IP()).Result()
			return c.Status(fiber.StatusServiceUnavailable).JSON(fiber.Map{
				 "error": "rate limit exceeded",
				 "rate_limit_reset": limit / time.Nanosecond / time.Minute })
		}

	}

	// check if the input is actually a valid URL
	if ! govalicator.IsURL(body.URL) {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{ "error": "invalid URL" })
	}


	// check for domain errors
	if ! helpers.RemoveDomainError(body.URL) {
		return c.Status(fiber.StatusServiceUnavailable).JSON(fiber.Map{ "error": "You can't access it" })
	}

	// enforce https, SSL 
	body.URL = helpers.EnforceHTTPS(body.URL){

		var id string
		// give abbility to user to set his/her own short URL
		var body.CustomShort == "" {
			id == uuid.New().String()[:6]
		} else {
			id = body.CustomShort
		}
		
		r := database.CreateClient(0)
		defer r.Close()

		val, _ := r.Get(database.Ctx, id).Result()
		if val != "" {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "short URL already exists", 
			})
		}
		if body.Expiry == 0 {
			body.Expiry = 24

		}
		err = r.Set(database.Ctx, id, body.URL, body.Expiry*3600*time.Second).Err()

		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "error saving to database", 
			})
		}

		r2.Decer(database.Ctx, c.IP())
	}

}