package validation

import (
	"fmt"

	"github.com/asaskevich/govalidator"
)

// ValidateSliceUUID performs validation to check if all IDs are valid uuids.
func ValidateSliceUUID(uuids []string) error {
	for _, v := range uuids {
		if !govalidator.IsUUID(v) {
			return fmt.Errorf("ID: %s is not a valid UUID", v)
		}
	}
	return nil
}
