package queries

type (
	EsQuery struct {
		Equals []FieldValue `json:"equals "`
	}
	FieldValue struct {
		Field string      `json:"field"`
		Value interface{} `json:"value"`
	}
)
