package mysql

import "fmt"

// GenerateSQLStr generate sort sql.
func (s *Sort) GenerateSQLStr() string {
	if s.Field != "" {
		return fmt.Sprintf("ORDER BY %s %s", s.Field, s.Direction.String())
	}
	return ""
}
