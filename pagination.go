package fortnox

type Pagination struct {
	Limit  int `schema:"limit"`
	Offset int `schema:"offset"`
	Page   int `schema:"page"`
}
