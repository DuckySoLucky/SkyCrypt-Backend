package main

import (
	"encoding/json"
	"fmt"
	"image"
	_ "image/png"
	"io/fs"
	"os"
	"path/filepath"
	"strings"

	"github.com/kettek/apng"
)

type McMeta struct {
	Animation McMetaAnimation `json:"animation"`
}

type McMetaAnimation struct {
	Frametime int `json:"frametime"`
}

func main() {
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

		if packDir.Name() == "Vanilla" {
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

			if !strings.Contains(path, "/textures/item/") {
				return nil
			}

			if strings.HasSuffix(path, ".json") {
				return nil
			}

			if strings.HasSuffix(path, ".mcmeta") {
				return nil
			}

			file, err := os.Open(path)
			if err != nil {
				fmt.Printf("Failed to open %s: %v\n", path, err)
				return nil
			}
			defer file.Close()

			img, _, err := image.Decode(file)
			if err != nil {
				fmt.Printf("Failed to decode %s: %v\n", path, err)
				return nil
			}

			width := img.Bounds().Dx()
			height := img.Bounds().Dy()
			if width != height && height%width == 0 {
				mcmetaPath := path + ".mcmeta"
				if _, err := os.Stat(mcmetaPath); os.IsNotExist(err) {
					return nil
				}

				mcmetaData, err := os.ReadFile(mcmetaPath)
				if err != nil {
					fmt.Printf("Failed to read %s: %v\n", mcmetaPath, err)
					return nil
				}

				var mcMeta McMeta
				if err := json.Unmarshal(mcmetaData, &mcMeta); err != nil {
					fmt.Printf("Failed to parse %s: %v\n", mcmetaPath, err)
					return nil
				}

				frameCount := height / width
				frames := make([]image.Image, frameCount)
				for i := 0; i < frameCount; i++ {
					frameRect := image.Rect(0, i*width, width, (i+1)*width)
					subImg := img.(interface {
						SubImage(r image.Rectangle) image.Image
					}).SubImage(frameRect)
					frames[i] = subImg
				}

				delay := uint16(mcMeta.Animation.Frametime * 50 / 10) // APNG delay is in 1/100s

				delays := make([]uint16, frameCount)
				for i := 0; i < frameCount; i++ {
					frameRect := image.Rect(0, i*width, width, (i+1)*width)
					subImg := img.(interface {
						SubImage(r image.Rectangle) image.Image
					}).SubImage(frameRect)
					frames[i] = subImg
					delays[i] = delay
				}

				apngImg := apng.APNG{}
				for i := 0; i < frameCount; i++ {
					frameRect := image.Rect(0, i*width, width, (i+1)*width)
					subImg := img.(interface {
						SubImage(r image.Rectangle) image.Image
					}).SubImage(frameRect)
					apngImg.Frames = append(apngImg.Frames, apng.Frame{
						Image:            subImg,
						DelayNumerator:   delay,
						DelayDenominator: 100,
					})
				}

				outFile, err := os.Create(path)
				if err != nil {
					fmt.Printf("Failed to create APNG %s: %v\n", path, err)
					return nil
				}
				defer outFile.Close()
				if err := apng.Encode(outFile, apngImg); err != nil {
					fmt.Printf("Failed to encode APNG %s: %v\n", path, err)
					return nil
				}

				fmt.Printf("Created APNG: %s\n", path)
			}

			return nil
		})
	}
}
