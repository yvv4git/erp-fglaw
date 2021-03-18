package forms

// Pagination is used for embding pagination struct to form.
type Pagination struct {
	Page  int64 `valid:"type(int64)" json:"page" query:"page"`
	Limit int64 `valid:"type(int64)" json:"limit" query:"limit"`
}

// Offset is used to get offset value.
func (p Pagination) Offset() int {
	return int(p.Page) * int(p.Limit)
}
