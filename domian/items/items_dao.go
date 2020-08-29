package items

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/shakilbd009/go-items-api/clients/elasticsearch"
	"github.com/shakilbd009/go-items-api/domian/queries"
	"github.com/shakilbd009/go-utils-lib/rest_errors"
)

const (
	indexItem = "items"
	typeItem  = "_doc"
)

func (i *Item) Save() rest_errors.RestErr {
	result, err := elasticsearch.Client.Index(indexItem, typeItem, i)
	if err != nil {
		return rest_errors.NewInternalServerError("error when trying to save item", fmt.Errorf("database_error"))
	}
	i.Id = result.Id
	return nil
}

func (i *Item) Get() rest_errors.RestErr {

	itemID := i.Id
	result, err := elasticsearch.Client.Get(indexItem, typeItem, i.Id)
	if err != nil {
		if strings.Contains(err.Error(), "404") {
			return rest_errors.NewNotFoundError(fmt.Sprintf("no item found with id %s", i.Id))
		}
		return rest_errors.NewInternalServerError(fmt.Sprintf("error when trying to get id %s", i.Id), fmt.Errorf("database_error"))
	}
	data, err := result.Source.MarshalJSON()
	if err != nil {
		return rest_errors.NewInternalServerError("error marshalling to bytes", err)
	}

	if err := json.Unmarshal(data, i); err != nil {
		return rest_errors.NewInternalServerError("error when trying to parse database response", err)
	}
	i.Id = itemID
	return nil
}

func (i *Item) Search(querys queries.EsQuery) ([]Item, rest_errors.RestErr) {

	result, err := elasticsearch.Client.Search(indexItem, querys.Build())
	if err != nil {
		return nil, rest_errors.NewInternalServerError("error when trying to search documents", fmt.Errorf("database error"))
	}
	items := make([]Item, result.TotalHits())
	for index, hit := range result.Hits.Hits {
		bytes, _ := hit.Source.MarshalJSON()
		var item Item
		if err := json.Unmarshal(bytes, &item); err != nil {
			return nil, rest_errors.NewInternalServerError("error when trying to parse response", fmt.Errorf("database error"))
		}
		items[index] = item
	}
	if len(items) == 0 {
		return nil, rest_errors.NewNotFoundError("no items found matching given criteria")
	}
	return items, nil
}

func (i *Item) Delete() rest_errors.RestErr {

	result, err := elasticsearch.Client.Delete(indexItem, typeItem, i.Id)
	if err != nil {
		if strings.Contains(err.Error(), "404") {
			return rest_errors.NewNotFoundError(fmt.Sprintf("no document found with given id %s", i.Id))
		}
		return rest_errors.NewInternalServerError("error when trying to delete document", fmt.Errorf("database error"))
	}
	i.Status = result
	return nil
}
