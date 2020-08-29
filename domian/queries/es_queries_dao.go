package queries

import "github.com/olivere/elastic"

func (q EsQuery) Build() elastic.Query {

	equalQueries := make([]elastic.Query, 0)
	query := elastic.NewBoolQuery()
	for _, eq := range q.Equals {
		equalQueries = append(equalQueries, elastic.NewMatchQuery(eq.Field, eq.Value))
	}
	query.Must(equalQueries...)
	return query
}
