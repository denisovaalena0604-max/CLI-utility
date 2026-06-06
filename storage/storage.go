package storage

import (
	"encoding/json"
	"errors"
	"os"

	"github.com/denisovaalena0604-max/CLI-utility/bins"
)

const fileName = "storage_bin.json"

func Save(list bins.BinList) error {
	
	data, err := json.MarshalIndent(list, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(fileName, data, 0644)
}

func Load() (bins.BinList, error) {
	var list bins.BinList

	data, err := os.ReadFile(fileName)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return bins.BinList{Bins: []bins.Bin{}}, nil
		}
		return list, err
	}

	if len(data) == 0 {
		return bins.BinList{Bins: []bins.Bin{}}, nil
	}

	if err := json.Unmarshal(data, &list); err != nil {
		return list, err
	}

	return list, nil
}