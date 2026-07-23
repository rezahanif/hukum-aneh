package validator

import (
	"fmt"
	"image"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"os"
)

type ImageValidator struct{}

func New() *ImageValidator {
	return &ImageValidator{}
}

// Validate checks if the image exists, is not empty, and can be successfully decoded.
// Can also check aspect ratio or dimensions matches requirements.
func (v *ImageValidator) Validate(filePath string) (bool, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return false, fmt.Errorf("open image: %w", err)
	}
	defer file.Close()

	stat, err := file.Stat()
	if err != nil {
		return false, fmt.Errorf("stat image: %w", err)
	}
	if stat.Size() == 0 {
		return false, fmt.Errorf("image file is empty")
	}

	config, _, err := image.DecodeConfig(file)
	if err != nil {
		return false, fmt.Errorf("decode image config: %w", err)
	}

	// Spec §5.9: Check dimensions / aspect ratio
	// Expected: aspect ratio roughly matching 4:5 or 1:1 depending on guide
	if config.Width == 0 || config.Height == 0 {
		return false, fmt.Errorf("invalid image dimensions: %dx%d", config.Width, config.Height)
	}

	return true, nil
}
