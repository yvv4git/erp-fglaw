package forms

// Clients ...
type Clients struct {
	Pagination   Pagination
	ID           int64  `valid:"type(int64)" json:"id"`
	Number       int64  `valid:"type(int64)" json:"num"`
	Address      string `valid:"length(0|50)" json:"addr"`
	CuitCustomer string `valid:"length(0|20)" json:"cuit"`
	ClientPhone  string `valid:"length(0|20)" json:"phone"`
	ClientTypeID int64  `valid:"type(int64)" json:"typeid"`
}
