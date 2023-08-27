package ports

import "context"

type Repository interface {
	GetDocuments(ctx context.Context, resources string, offset, limit int64) (interface{}, error)
	GetDocumentById(ctx context.Context, resources string, id string) (interface{}, error)
	CreateDocument(ctx context.Context, resources string, resource interface{}) (interface{}, error)
	UpdateDocument(ctx context.Context, resources string, resource interface{}, id string) (interface{}, error)
	DeleteDocument(ctx context.Context, resources string, id string) (interface{}, error)
}
