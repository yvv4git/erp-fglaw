package forms

// Pagination is used for embding pagination struct to form.
type Pagination struct {
	Page  int64 `valid:"type(int64)" json:"page"`
	Limit int64 `valid:"type(int64)" json:"limit"`
}
