package services

import (
	"github.com/popsompong/bookstore_items-api/domain/items"
	"github.com/popsompong/bookstore_utils-go/rest_errors"
	"net/http"
)

var (
	ItemsService itemsServiceInterface = &itemsService{}
)

type itemsServiceInterface interface {
	Create(items.Item) (*items.Item, rest_errors.RestErr)
	Get(string) (*items.Item, rest_errors.RestErr)
}

type itemsService struct {
}

func (s *itemsService) Create(items.Item) (*items.Item, rest_errors.RestErr) {

	return nil, rest_errors.NewRestError("implement me!", http.StatusNotImplemented, "not_implemented", nil)
}

func (s *itemsService) Get(string) (*items.Item, rest_errors.RestErr) {

	return nil, rest_errors.NewRestError("implement me!", http.StatusNotImplemented, "not_implemented", nil)
}
