package src

import (
	"fmt"
	"log"
	"os"
	notenoughupdates "skycrypt/src/NotEnoughUpdates"
	"skycrypt/src/api"
	redis "skycrypt/src/db"
	"skycrypt/src/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/joho/godotenv"
)

func SetupApplication() error {
	err := godotenv.Load()
	if err != nil && os.Getenv("FIBER_PREFORK_CHILD") == "" {
		log.Println("No .env file found, using environment variables")
	}

	redisHost := os.Getenv("REDIS_HOST")
	if redisHost == "" {
		redisHost = "localhost"
	}

	redisPort := os.Getenv("REDIS_PORT")
	if redisPort == "" {
		redisPort = "6379"
	}

	redisPassword := os.Getenv("REDIS_PASSWORD")

	redisAddr := fmt.Sprintf("%s:%s", redisHost, redisPort)

	err = redis.InitRedis(redisAddr, redisPassword, 0)
	if err != nil {
		return fmt.Errorf("failed to connect to Redis: %v", err)
	}

	if err := api.LoadSkyBlockItems(); err != nil {
		return fmt.Errorf("error loading SkyBlock items: %v", err)
	}

	/*
		if err := notenoughupdates.InitializeNEURepository(); err != nil {
			return fmt.Errorf("error initializing repository: %v", err)
		}

		if err := notenoughupdates.UpdateNEURepository(); err != nil {
			return fmt.Errorf("error updating repository: %v", err)
		}
	*/

	err = notenoughupdates.ParseNEURepository()
	if err != nil {
		return fmt.Errorf("error parsing NEU repository: %v", err)
	}

	if os.Getenv("FIBER_PREFORK_CHILD") == "" {
		fmt.Print("[SKYCRYPT] SkyCrypt initialized successfully\n")
	}

	return nil
}

func SetupRoutes(app *fiber.App) {
	app.Use(compress.New(compress.Config{
		Level: compress.LevelBestSpeed,
	}))

	// Assets folder
	app.Static("/assets", "assets")

	if os.Getenv("DEV") != "true" {
		fmt.Println("[ENVIROMENT] Running in production mode")

		/*
			app.Use(etag.New())
			app.Use("/api", cache.New(cache.Config{
				Expiration:   5 * time.Minute,
				CacheControl: true,
			})
		*/
	}

	api := app.Group("/api")

	// USERNAME AND UUID RESOLVING
	api.Get("/uuid/:username", routes.UUIDHandler)
	api.Get("/username/:uuid", routes.UsernameHandler)

	// HYPIXEL API ENDPOINTS
	api.Get("/profiles/:uuid", routes.ProfilesHandler)
	api.Get("/player/:uuid", routes.PlayerHandler)
	api.Get("/museum/:profileId", routes.MuseumHandler)
	api.Get("/garden/:profileId", routes.GardenHandler)

	// STATS ENDPOINTS
	api.Get("/stats/:uuid/:profileId", routes.StatsHandler)
	api.Get("/stats/:uuid", routes.StatsHandler)

	api.Get("/networth/:uuid/:profileId", routes.NetworthHandler)

	api.Get("/gear/:uuid/:profileId", routes.GearHandler)

	api.Get("/accessories/:uuid/:profileId", routes.AccessoriesHandler)

	api.Get("/pets/:uuid/:profileId", routes.PetsHandler)

	api.Get("/inventory/:uuid/:profileId/:inventoryId", routes.InventoryHandler)
	api.Get("/inventory/:uuid/:profileId/:inventoryId/:search", routes.InventoryHandler)

	api.Get("/skills/:uuid/:profileId", routes.SkillsHandler)

	api.Get("/dungeons/:uuid/:profileId", routes.DungeonsHandler)

	api.Get("/slayer/:uuid/:profileId", routes.SlayersHandler)

	api.Get("/minions/:uuid/:profileId", routes.MinionsHandler)

	api.Get("/bestiary/:uuid/:profileId", routes.BestiaryHandler)

	api.Get("/collections/:uuid/:profileId", routes.CollectionsHandler)

	api.Get("/crimson_isle/:uuid/:profileId", routes.CrimsonIsleHandler)

	api.Get("/rift/:uuid/:profileId", routes.RiftHandler)

	api.Get("/misc/:uuid/:profileId", routes.MiscHandler)

	api.Get("/embed/:uuid/:profileId", routes.EmbedHandler)

	// RENDERING ENDPOINTS
	api.Get("/head/:textureId", routes.HeadHandlers)

	api.Get("/item/:itemId", routes.ItemHandlers)

	api.Get("/potion/:type/:color", routes.PotionHandlers)

	api.Get("/leather/:type/:color", routes.LeatherHandlers)

}
