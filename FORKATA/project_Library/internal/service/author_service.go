package service

import (
	"context"
	"golang/project_Library/internal/models"
)

func (service *SuperService) CreateAuthor(ctx context.Context, author models.Author) error{
	return service.AuthorService.Create(ctx, author)	
}

func (service *SuperService) GetAuthorByID(ctx context.Context, id string) (models.Author, error){
	return service.AuthorService.GetByID(ctx, id)
}

func (service *SuperService) GetAllAuthorsBooks(ctx context.Context, author_id string) ([]models.Book, error){
	return service.AuthorService.GetAllBooks(ctx, author_id)
}

func (service *SuperService) DeleteAuthor(ctx context.Context, id string) error{
	return service.AuthorService.Delete(ctx, id)
}

func (service *SuperService) GetTopAuthors(ctx context.Context) ([]models.Author, error){
	return service.AuthorService.GetTop(ctx)
}