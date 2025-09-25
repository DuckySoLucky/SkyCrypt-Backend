package routes

import (
	"skycrypt/src/constants"
	"skycrypt/src/lib"
	"skycrypt/src/utility"

	"github.com/gofiber/fiber/v2"
)

func ItemHandlers(c *fiber.Ctx) error {
	// timeNow := time.Now()
	textureId := c.Params("itemId")
	if textureId == "" {
		c.Status(400)
		return c.JSON(constants.InvalidItemProvidedError)
	}

	textureBytes, err := lib.RenderItem(textureId)
	if err != nil {
		if redirectErr, ok := err.(lib.RedirectError); ok {
			return c.Redirect(redirectErr.URL, 302)
		}

		c.Status(500)
		// return c.JSON(constants.InvalidItemProvidedError)
		utility.SendWebhook("error", err.Error(), []byte(nil))
		return c.JSON(fiber.Map{
			"error": err.Error(),
		})

	}

	c.Type("png")
	// fmt.Printf("Returning /api/item/%s in %s\n", textureId, time.Since(timeNow))
	return c.Send(textureBytes)
}
