package routes

import (
	"skycrypt/src/constants"
	"skycrypt/src/lib"

	"github.com/gofiber/fiber/v2"
)

// ItemHandlers godoc
// @Summary Render and return an item image
// @Description Returns a PNG image of an item for the given texture ID
// @Tags item
// @Produce  png
// @Param itemId path string true "Item ID"
// @Success 200 {file} binary "PNG image of the item"
// @Failure 400 {object} models.ProcessingError
// @Failure 500 {string} string "Failed to render item"
// @Router /api/item/{itemId} [get]
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
		return c.JSON(constants.InvalidItemProvidedError)
	}

	c.Type("png")
	// fmt.Printf("Returning /api/item/%s in %s\n", textureId, time.Since(timeNow))
	return c.Send(textureBytes)
}
