package service

import (
	"context"
	"golang/project_Library/internal/models"
)

func (service *SuperService) CreateBook(ctx context.Context, book models.Book) error{
	return service.BookService.Create(ctx, book)
}

func (service *SuperService) GetBookByID(ctx context.Context, id string) (models.Book, error){
	return service.BookService.GetByID(ctx, id)
}

func (service *SuperService) Delete(ctx context.Context, id string) error{
	return service.BookService.Delete(ctx, id)
}