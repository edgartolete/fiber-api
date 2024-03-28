package handlers

import "github.com/gofiber/fiber/v2"


type User struct {
    ID    int    `json:"id"`
    Name  string `json:"name"`
    Married bool `json:"married"`
}




	func GetUsersHandler(c *fiber.Ctx) error {
		var users = []User{{
			ID:    1,
            Name:  "Edgar",
            Married: true,
		}, {
			ID:    1,
            Name:  "John",
            Married: false,
		}};
		
		return c.JSON(users)
	}
	

