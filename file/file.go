package file

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

var (
	ErrNotJSONExt = errors.New("file: расширение файла должно быть .json")
	ErrFileNotExist = errors.New("file: файл не существует")
	ErrInvalidJSON  = errors.New("file: содержимое не является валидным JSON")
)

func ReadJSON(path string) ([]byte, error) {
	
	ext := strings.ToLower(filepath.Ext(path))
	if ext != ".json" {
		return nil, fmt.Errorf("%w: получено %q", ErrNotJSONExt, ext)
	}

	data, err := os.ReadFile(path)
	if err != nil {
		
		if errors.Is(err, os.ErrNotExist) {
			return nil, fmt.Errorf("%w: %s", ErrFileNotExist, path)
		}
		return nil, fmt.Errorf("file: ошибка чтения %s: %w", path, err)
	}

	if !json.Valid(data) {
		return nil, fmt.Errorf("%w: %s", ErrInvalidJSON, path)
	}
	return data, nil
}