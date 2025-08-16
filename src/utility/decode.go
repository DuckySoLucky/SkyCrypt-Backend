package utility

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"fmt"
	"skycrypt/src/models"
	"sync"

	"github.com/Tnze/go-mc/nbt"
)

var gzipReaderPool = sync.Pool{
	New: func() any {
		return &gzip.Reader{}
	},
}

func DecodeInventory(inventoryData *string) (*models.DecodedInventory, error) {
	if *inventoryData == "" {
		return &models.DecodedInventory{}, nil
	}

	decodedData, err := base64.StdEncoding.DecodeString(*inventoryData)
	if err != nil {
		return nil, fmt.Errorf("failed to decode base64: %w", err)
	}

	reader := gzipReaderPool.Get().(*gzip.Reader)
	defer gzipReaderPool.Put(reader)

	err = reader.Reset(bytes.NewReader(decodedData))
	if err != nil {
		return nil, fmt.Errorf("failed to reset gzip reader: %w", err)
	}

	var nbtData models.DecodedInventory
	decoder := nbt.NewDecoder(reader)
	_, err = decoder.Decode(&nbtData)
	if err != nil {
		return nil, fmt.Errorf("failed to parse NBT data: %w", err)
	}

	return &nbtData, nil
}
