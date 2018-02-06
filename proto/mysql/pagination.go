package mysql

import "fmt"

// GenerateSQLStr generates sql for pagination.
func (p *Pagination) GenerateSQLStr() string {
	if p.Page == 0 || p.PerPage == 0 {
		return fmt.Sprintf("LIMIT %d OFFSET %d", 20, 0)
	}
	return fmt.Sprintf("LIMIT %d OFFSET %d", p.PerPage, (p.Page-1)*p.PerPage)
}
