package services

import (
	"github.com/shakilbd009/go-items-api/domian/items"
	"github.com/shakilbd009/go-items-api/domian/queries"
	"github.com/shakilbd009/go-utils-lib/rest_errors"
)

var (
	ItemsService itemsServiceInterface = &itemsService{}
)

type itemsServiceInterface interface {
	Create(items.Item) (*items.Item, rest_errors.RestErr)
	Get(string) (*items.Item, rest_errors.RestErr)
	Search(queries.EsQuery) ([]items.Item, rest_errors.RestErr)
	Delete(id string) (*items.Item, rest_errors.RestErr)
}

type itemsService struct{}

func (s *itemsService) Create(item items.Item) (*items.Item, rest_errors.RestErr) {

	if err := item.Save(); err != nil {
		return nil, err
	}
	return &item, nil
}

func (s *itemsService) Get(id string) (*items.Item, rest_errors.RestErr) {

	item := &items.Item{
		Id: id,
	}
	if err := item.Get(); err != nil {
		return nil, err
	}
	return item, nil
}

func (s *itemsService) Search(querys queries.EsQuery) ([]items.Item, rest_errors.RestErr) {
	dao := items.Item{}
	return dao.Search(querys)
}

func (s *itemsService) Delete(id string) (*items.Item, rest_errors.RestErr) {

	item := &items.Item{Id: id}
	if err := item.Delete(); err != nil {
		return nil, err
	}
	return item, nil
}
