package lib

import (
	"encoding/json"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"regexp"
	"skycrypt/src/constants"
	"skycrypt/src/models"
	"skycrypt/src/utility"
	"slices"
	"strings"
)

func GetTexturePath(texturePath string, textureString string) string {
	textureId := textureString[strings.Index(textureString, "/")+1:]
	formattedPath := ""
	if texturePath == "Vanilla" {
		formattedPath = fmt.Sprintf("resourcepacks/%s/assets/firmskyblock/models/item/%s", texturePath, textureId)
	} else {
		if after, ok := strings.CutPrefix(textureId, "firmskyblock:item"); ok {
			textureId = after
		}

		formattedPath = fmt.Sprintf("resourcepacks/%s/assets/cittofirmgenerated/textures/item/%s.png", texturePath, textureId)
	}

	if os.Getenv("DEV") != "true" {
		return fmt.Sprintf("/assets/%s", formattedPath)
	}

	return "http://localhost:8080/assets/" + formattedPath
}

func GetTexture(item models.TextureItem, disabledPacksParam ...[]string) AppliedItemTexture {
	textures := ITEM_MAP[strings.ToLower(item.Tag.ExtraAttributes["id"].(string))]
	if len(textures) == 0 {
		return AppliedItemTexture{}
	}

	disabledPacks := disabledPacksParam[0]
	for _, disabledPack := range disabledPacks {
		textures = slices.DeleteFunc(textures, func(t models.ItemTexture) bool {
			return t.ResourcePackId == disabledPack
		})
	}

	if len(textures) == 0 {
		return AppliedItemTexture{}
	}

	// First, check all overrides with 'firmament:all' predicate
	var evalPredicate func(key string, value interface{}) bool
	evalPredicate = func(key string, value interface{}) bool {
		switch key {
		case "firmament:display_name":
			switch v := value.(type) {
			case map[string]interface{}:
				if regexVal, ok := v["regex"]; ok {
					if regexStr, ok := regexVal.(string); ok {
						matched, err := regexp.MatchString(regexStr, item.Tag.Display.Name)
						return err == nil && matched
					}
				}
			case string:
				return v == item.Tag.Display.Name
			}
		case "firmament:lore":
			switch v := value.(type) {
			case map[string]interface{}:
				if regexVal, ok := v["regex"]; ok {
					if regexStr, ok := regexVal.(string); ok {
						for _, line := range item.Tag.Display.Lore {
							matched, err := regexp.MatchString(regexStr, line)
							if err == nil && matched {
								return true
							}
						}
					}
				}
			case string:
				for _, line := range item.Tag.Display.Lore {
					if v == line {
						return true
					}
				}
			}
			return false
		case "firmament:extra_attributes":
			if m, ok := value.(map[string]interface{}); ok {
				if path, ok := m["path"].(string); ok {
					attrVal, exists := item.Tag.ExtraAttributes[path]
					if !exists {
						return false
					}

					intVal, ok := attrVal.(int)
					if !ok {
						// Try float64 conversion (just in case)
						if f, ok := attrVal.(float64); ok {
							intVal = int(f)
						} else {
							return false
						}
					}

					if intMap, ok := m["int"].(map[string]interface{}); ok {
						if minVal, ok := intMap["min"].(float64); ok {
							if intVal < int(minVal) {
								return false
							}
						}
					}
					return true
				}
			}
			return false
		case "firmament:all":
			// value is expected to be []interface{} of predicate maps
			if arr, ok := value.([]interface{}); ok {
				for _, sub := range arr {
					if subMap, ok := sub.(map[string]interface{}); ok {
						for k, v := range subMap {
							if !evalPredicate(k, v) {
								return false
							}
						}
					} else {
						return false
					}
				}
				return true
			}
			return false
		case "firmament:not":
			// value is a predicate map or array of predicate maps
			switch v := value.(type) {
			case map[string]interface{}:
				for k, val := range v {
					if evalPredicate(k, val) {
						return false
					}
				}
				return true
			case []interface{}:
				for _, sub := range v {
					if subMap, ok := sub.(map[string]interface{}); ok {
						for k, val := range subMap {
							if evalPredicate(k, val) {
								return false
							}
						}
					}
				}
				return true
			}
			return false
		}
		return false
	}

	for _, texture := range textures {
		// For each override, all predicates must match (AND logic)
		for i := len(texture.Overrides) - 1; i >= 0; i-- {
			override := texture.Overrides[i]
			allMatch := true
			for k, v := range override.Predicate {
				if k == "firmament:not" {
					// firmament:not must be true for the override to match
					if !evalPredicate(k, v) {
						allMatch = false
						break
					}
				} else {
					if !evalPredicate(k, v) {
						allMatch = false
						break
					}
				}
			}
			if allMatch {
				return AppliedItemTexture{
					Texture:     override.Texture,
					TexturePack: texture.ResourcePackId,
				}
			}
		}

		if tex, ok := texture.Textures["layer0"]; ok {
			return AppliedItemTexture{
				Texture:     tex,
				TexturePack: texture.ResourcePackId,
			}
		}

		for _, tex := range texture.Textures {
			return AppliedItemTexture{
				Texture:     tex,
				TexturePack: texture.ResourcePackId,
			}
		}

	}

	return AppliedItemTexture{}
}

var VANILLA_ITEM_MAP = map[string]models.ItemTexture{}
var ITEM_MAP = map[string][]models.ItemTexture{}

func init() {
	assetsRoot := "assets/resourcepacks"
	packDirs, err := os.ReadDir(assetsRoot)
	if err != nil {
		fmt.Printf("Failed to read assets directory: %v\n", err)
		return
	}

	for _, packDir := range packDirs {
		if !packDir.IsDir() {
			continue
		}

		packAssetsPath := filepath.Join(assetsRoot, packDir.Name(), "assets")
		if _, err := os.Stat(packAssetsPath); os.IsNotExist(err) {
			continue
		}

		configPath := filepath.Join(assetsRoot, packDir.Name(), "config.json")
		if _, err := os.Stat(configPath); err != nil {
			fmt.Printf("No config.json found for pack %s, skipping\n", packDir.Name())
			continue
		}

		data, err := os.ReadFile(configPath)
		if err != nil {
			fmt.Printf("Failed to read config.json for pack %s: %v\n", packDir.Name(), err)
		}

		var config models.ResourcePackConfig
		if err := json.Unmarshal(data, &config); err != nil {
			fmt.Printf("Failed to parse config.json for pack %s: %v\n", packDir.Name(), err)
		}

		if config.Disabled {
			fmt.Printf("Skipping disabled resource pack: %s\n", packDir.Name())
			continue
		}

		filepath.WalkDir(packAssetsPath, func(path string, d fs.DirEntry, err error) error {
			if err != nil {
				return err
			}

			if d.IsDir() {
				return nil
			}

			if !strings.Contains(path, "/models/item/") {
				return nil
			}

			if !strings.HasSuffix(path, ".json") {
				return nil
			}

			data, err := os.ReadFile(path)
			if err != nil {
				fmt.Printf("Failed to read %s: %v\n", path, err)
				return nil
			}

			if packDir.Name() != "Vanilla" {
				var model models.ItemTexture = models.ItemTexture{ResourcePackId: config.Id}
				if err := json.Unmarshal(data, &model); err != nil {
					fmt.Printf("Failed to parse %s: %v\n", path, err)
					return nil
				}

				// Skip 3D models for now
				if len(model.Elements) > 0 || model.HeadModel != "" {
					return nil
				}

				fileName := filepath.Base(path)
				itemName := fileName[:len(fileName)-len(filepath.Ext(fileName))]
				if _, exists := ITEM_MAP[itemName]; !exists {
					ITEM_MAP[itemName] = []models.ItemTexture{}
				}

				for i := range model.Overrides {
					if model.Overrides[i].Texture != "" {
						model.Overrides[i].Texture = GetTexturePath(packDir.Name(), model.Overrides[i].Texture)
					}
				}

				for key, texture := range model.Textures {
					if texture != "" {
						model.Textures[key] = GetTexturePath(packDir.Name(), texture)
					}
				}

				ITEM_MAP[itemName] = append(ITEM_MAP[itemName], model)
				return nil
			} else {
				var model models.VanillaTexture
				if err := json.Unmarshal(data, &model); err != nil {
					fmt.Printf("Failed to parse %s: %v\n", path, err)
					return nil
				}

				textureId := fmt.Sprintf("%s:%d", model.VanillaId, model.Damage)
				fileName := strings.ReplaceAll(filepath.Base(path), ".json", ".png")
				VANILLA_ITEM_MAP[textureId] = models.ItemTexture{
					Parent:    "item/generated",
					Textures:  map[string]string{"layer0": GetTexturePath(packDir.Name(), fileName)},
					Overrides: []models.Override{},
				}
			}

			return nil
		})
	}
}

type AppliedItemTexture struct {
	Texture     string
	TexturePack string
}

func ApplyTexture(item models.TextureItem, disabledPacksParam ...[]string) AppliedItemTexture {
	// ? NOTE: we're ignoring enchanted books because they're quite expensive to render and not really worth the performance hit
	if item.Tag.ExtraAttributes == nil || item.Tag.ExtraAttributes["id"] == "ENCHANTED_BOOK" {
		if os.Getenv("DEV") == "true" {
			return AppliedItemTexture{Texture: "http://localhost:8080/assets/resourcepacks/Vanilla/assets/firmskyblock/models/item/enchanted_book.png"}
		}

		return AppliedItemTexture{Texture: "/assets/resourcepacks/Vanilla/assets/firmskyblock/models/item/enchanted_book.png"}
	}

	disabledPacks := []string{}
	if len(disabledPacksParam) > 0 {
		disabledPacks = disabledPacksParam[0]
	}

	customTexture := GetTexture(item, disabledPacks)
	if customTexture.Texture != "" {
		if !strings.Contains(customTexture.Texture, "Vanilla") && !strings.Contains(customTexture.Texture, "skull") {
			return customTexture
		}
	}

	if item.Tag.SkullOwner != nil && item.Tag.SkullOwner.Properties.Textures[0].Value != "" {
		skinHash := utility.GetSkinHash(item.Tag.SkullOwner.Properties.Textures[0].Value)
		if os.Getenv("DEV") != "true" {
			return AppliedItemTexture{Texture: fmt.Sprintf("/api/head/%s", skinHash)}
		}

		return AppliedItemTexture{Texture: fmt.Sprintf("http://localhost:8080/api/head/%s", skinHash)}
	}

	// Preparsed texture from /api/item endpoint
	if item.Texture != "" {
		if os.Getenv("DEV") != "true" {
			return AppliedItemTexture{Texture: fmt.Sprintf("/api/head/%s", item.Texture)}
		}

		return AppliedItemTexture{Texture: fmt.Sprintf("http://localhost:8080/api/head/%s", item.Texture)}
	}

	if *item.ID >= 298 && *item.ID <= 301 {
		armorType := constants.ARMOR_TYPES[*item.ID-298]

		armorColor := fmt.Sprintf("%06X", item.Tag.Display.Color)
		if item.Tag.ExtraAttributes["dye_item"] != "" {
			idStr, ok := item.Tag.ExtraAttributes["id"].(string)
			if ok {
				defaultHexColor := constants.ITEMS[idStr].Color
				if defaultHexColor != "" {
					armorColor = defaultHexColor
				}

				if defaultHexColor != "" {
					armorColor = defaultHexColor
				}
			}

		}

		if os.Getenv("DEV") != "true" {
			return AppliedItemTexture{Texture: fmt.Sprintf("/api/leather/%s/%s", armorType, armorColor)}
		}

		return AppliedItemTexture{Texture: fmt.Sprintf("http://localhost:8080/api/leather/%s/%s", armorType, armorColor)}
	}

	textureId := fmt.Sprintf("%d:%d", *item.ID, *item.Damage)
	if texture, ok := VANILLA_ITEM_MAP[textureId]; ok {
		if tex, ok := texture.Textures["layer0"]; ok && tex != "" {
			return AppliedItemTexture{Texture: tex}
		}

		for _, tex := range texture.Textures {
			if tex == "" {
				continue
			}

			return AppliedItemTexture{Texture: tex}
		}
	}

	vanillaPath := fmt.Sprintf("assets/resourcepacks/Vanilla/assets/firmskyblock/models/item/%s.png", strings.ToLower(item.RawId))
	if _, err := os.Stat(vanillaPath); err == nil {
		if os.Getenv("DEV") != "true" {
			return AppliedItemTexture{Texture: "/" + vanillaPath}
		}

		return AppliedItemTexture{Texture: "http://localhost:8080/" + vanillaPath}
	}

	fmt.Printf("[CUSTOM_RESOURCES] No custom texture found for item %s, returning default barrier texture\n", item.Tag.ExtraAttributes["id"])
	if os.Getenv("DEV") != "true" {
		return AppliedItemTexture{Texture: "/assets/resourcepacks/Vanilla/assets/firmskyblock/models/item/barrier.png"}
	}

	return AppliedItemTexture{Texture: "http://localhost:8080/assets/resourcepacks/Vanilla/assets/firmskyblock/models/item/barrier.png"}
}
