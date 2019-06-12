package models

type Operator string

const (
	EQUALS Operator = "EQUALS"
	IN Operator = "IN"
)

type Predicate struct {
	Field string
	Operator Operator
	Values []string
}

type Selector struct {
	Fields []string	`json:"fields,omitempty"`
	Predicates []Predicate	`json:"predicates,omitempty"`
	// TODO:
}
