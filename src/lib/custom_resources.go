package lib

import (
	"encoding/json"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"regexp"
	"skycrypt/src/models"
	"skycrypt/src/utility"
	"strings"

	"golang.org/x/exp/slices"
)

func GetTexturePath(texturePath string, textureString string) string {
	textureId := textureString[strings.Index(textureString, "/")+1:]
	formattedPath := ""
	if texturePath == "Vanilla" {
		formattedPath = fmt.Sprintf("%s/assets/firmskyblock/models/item/%s", texturePath, textureId)
	} else {
		if after, ok := strings.CutPrefix(textureId, "firmskyblock:item"); ok {
			textureId = after
		}

		formattedPath = fmt.Sprintf("%s/assets/cittofirmgenerated/textures/item/%s.png", texturePath, textureId)
	}

	return "http://localhost:8080/assets/" + formattedPath
}

func GetTexture(item models.TextureItem) string {
	textures := ITEM_MAP[strings.ToLower(item.Tag.ExtraAttributes["id"].(string))]
	if len(textures) == 0 {
		textureId := fmt.Sprintf("%d:%d", *item.ID, *item.Damage)
		// if it's a SKULL, and custom texture is not found, return the skin texture
		if textureId == "397:3" {
			if item.Tag.SkullOwner != nil && item.Tag.SkullOwner.Properties.Textures[0].Value != "" {
				skinHash := utility.GetSkinHash(item.Tag.SkullOwner.Properties.Textures[0].Value)
				return fmt.Sprintf("http://localhost:8080/api/head/%s", skinHash)
			}

			return ""
		}

		if texture, ok := VANILLA_ITEM_MAP[textureId]; ok {
			if tex, ok := texture.Textures["layer0"]; ok && tex != "" {
				return tex
			}

			for _, tex := range texture.Textures {
				if tex != "" {
					return tex
				}
			}

			fmt.Printf("[CUSTOM_RESOURCES] No textures found for vanilla item: %s %+v\n", textureId, VANILLA_ITEM_MAP[textureId])
			return ""
		} else {
			vanillaPath := fmt.Sprintf("assets/resourcepacks/Vanilla/assets/firmskyblock/models/item/%s.png", strings.ToLower(item.RawId))
			if _, err := os.Stat(vanillaPath); err == nil {
				return "http://localhost:8080/" + vanillaPath
			}

			fmt.Printf("[CUSTOM_RESOURCES] Texture not found for item: %s\n", textureId)
		}

		return ""
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
				return override.Texture
			}
		}

		if tex, ok := texture.Textures["layer0"]; ok {
			return tex
		}

		for _, tex := range texture.Textures {
			return tex
		}

	}

	return ""
}

var VANILLA_ITEM_MAP = map[string]models.ItemTexture{}
var ITEM_MAP = map[string][]models.ItemTexture{}

var ALLOWED_PARENTS = []string{
	"minecraft:item/handheld",
	"cittofirmgenerated:item/skyblock/item",
	"minecraft:item/fishing_rod",
	"cittofirmgenerated:item/skyblock/vacuum",
	"cittofirmgenerated:item/skyblock/gun",
	"cittofirmgenerated:item/metal_detector",
}

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
				var model models.ItemTexture
				if err := json.Unmarshal(data, &model); err != nil {
					fmt.Printf("Failed to parse %s: %v\n", path, err)
					return nil
				}

				if !slices.Contains(ALLOWED_PARENTS, model.Parent) {
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
