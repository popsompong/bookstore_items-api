package services

import (
	items2 "github.com/popsompong/bookstore_items-api/domain/items"
	queries2 "github.com/popsompong/bookstore_items-api/domain/queries"
	"github.com/popsompong/bookstore_utils-go/rest_errors"
)

var (
	ItemsService itemsServiceInterface = &itemsService{}
)

type itemsServiceInterface interface {
	Create(items2.Item) (*items2.Item, rest_errors.RestErr)
	Get(string) (*items2.Item, rest_errors.RestErr)
	Search(queries2.EsQuery) ([]items2.Item, rest_errors.RestErr)
}

type itemsService struct {
}

func (s *itemsService) Create(item items2.Item) (*items2.Item, rest_errors.RestErr) {
	if err := item.Save(); err != nil {
		return nil, err
	}

	return &item, nil
}

func (s *itemsService) Get(id string) (*items2.Item, rest_errors.RestErr) {
	item := items2.Item{Id: id}

	if err := item.Get(); err != nil {
		return nil, err
	}

	return &item, nil
}

func (s *itemsService) Search(query queries2.EsQuery) ([]items2.Item, rest_errors.RestErr) {
	dao := items2.Item{}
	return dao.Search(query)
}
