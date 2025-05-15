package web

// NewSchema generates the required Query and Mutation objects
func NewSchema() (qraphql.Schema, error) {
	return q.NewSchema(qraphql.SchemaConfig{
		Query:    RootQuery,
		Mutation: nil,
	})
}
