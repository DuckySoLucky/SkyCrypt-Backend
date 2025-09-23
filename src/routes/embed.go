package routes

import (
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
)

func EmbedHandler(c *fiber.Ctx) error {
	timeNow := time.Now()

	uuid := c.Params("uuid")
	profileId := c.Params("profileId")
	/*embed, err := redis.Get(fmt.Sprintf("embed:%s:%s", uuid, profileId))
	if err != nil {
		c.Status(400)
		return c.JSON(constants.InvalidUserError)
	}

	var embedData models.EmbedData
	if err := json.Unmarshal([]byte(embed), &embedData); err != nil {
		c.Status(500)
		return c.JSON(constants.InternalServerError)
	}*/

	fmt.Printf("Returning /api/embed/%s in %s\n", profileId, time.Since(timeNow))


	return c.JSON(fiber.Map{
		"embed": uuid,
	})
}
