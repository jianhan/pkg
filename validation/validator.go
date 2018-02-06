package validation

import (
	"fmt"

	"github.com/asaskevich/govalidator"
	"github.com/gosimple/slug"
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

// ValidateSliceSlugs checks if a slice of slugs are all valid.
func ValidateSliceSlugs(slugs []string) error {
	for _, v := range slugs {
		if !slug.IsSlug(v) {
			return fmt.Errorf("%s is not a valid slug", v)
		}
	}
	return nil
}

// CheckSlug is a wrapper for checking slug, will be used for customized
// validation package.
func CheckSlug(value interface{}) error {
	s, _ := value.(string)
	if !slug.IsSlug(s) {
		return fmt.Errorf("%s must be a valid slug", s)
	}
	return nil
}
