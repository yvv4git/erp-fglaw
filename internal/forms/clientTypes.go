package forms

// ClientTypes form.
type ClientTypes struct {
	ID         int64  `valid:"type(int64)" json:"id"`
	ClientType string `valid:"length(0|20)" json:"ctype"`
	ActingAs   string `valid:"length(0|20)" json:"actas"`
}
