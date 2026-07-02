package image

import (
	"encoding/json"
	"fmt"

	appimage "github.com/JSYoo5B/SandStack/internal/app/image"
)

type imagePatchOperation struct {
	Op    string          `json:"op"`
	Path  string          `json:"path"`
	Value json.RawMessage `json:"value"`
}

func toPatchImage(operations []imagePatchOperation) (appimage.PatchImage, error) {
	var patch appimage.PatchImage
	for _, operation := range operations {
		if operation.Op != "replace" {
			return appimage.PatchImage{}, fmt.Errorf("unsupported patch operation")
		}

		if err := applyPatchOperation(&patch, operation); err != nil {
			return appimage.PatchImage{}, err
		}
	}

	return patch, nil
}

func applyPatchOperation(
	patch *appimage.PatchImage,
	operation imagePatchOperation,
) error {
	switch operation.Path {
	case "/name":
		var value string
		if err := json.Unmarshal(operation.Value, &value); err != nil {
			return err
		}
		patch.Name = &value
	case "/min_disk":
		var value int
		if err := json.Unmarshal(operation.Value, &value); err != nil {
			return err
		}
		patch.MinDisk = &value
	case "/min_ram":
		var value int
		if err := json.Unmarshal(operation.Value, &value); err != nil {
			return err
		}
		patch.MinRAM = &value
	case "/protected":
		var value bool
		if err := json.Unmarshal(operation.Value, &value); err != nil {
			return err
		}
		patch.Protected = &value
	case "/visibility":
		var value string
		if err := json.Unmarshal(operation.Value, &value); err != nil {
			return err
		}
		patch.Visibility = &value
	case "/tags":
		var value []string
		if err := json.Unmarshal(operation.Value, &value); err != nil {
			return err
		}
		patch.Tags = value
		patch.HasTags = true
	default:
		return fmt.Errorf("unsupported patch path")
	}

	return nil
}
