// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.18.0

package db

import (
	"context"
)

type Querier interface {
	CreateAccount(ctx context.Context, arg CreateAccountParams) (Account, error)
	CreateCategory(ctx context.Context, arg CreateCategoryParams) (Category, error)
	CreateUser(ctx context.Context, arg CreateUserParams) (User, error)
	DeleteAccount(ctx context.Context, id int32) error
	DeleteCategories(ctx context.Context, id int32) error
	GetAccount(ctx context.Context, id int32) (Account, error)
	GetAccounts(ctx context.Context, arg GetAccountsParams) ([]GetAccountsRow, error)
	GetAccountsGraph(ctx context.Context, arg GetAccountsGraphParams) (int64, error)
	GetAccountsReports(ctx context.Context, arg GetAccountsReportsParams) (interface{}, error)
	GetCategories(ctx context.Context, arg GetCategoriesParams) ([]Category, error)
	GetCategory(ctx context.Context, id int32) (Category, error)
	GetUser(ctx context.Context, username string) (User, error)
	GetUserById(ctx context.Context, id int32) (User, error)
	UpdateAccount(ctx context.Context, arg UpdateAccountParams) (Account, error)
	UpdateCategories(ctx context.Context, arg UpdateCategoriesParams) (Category, error)
}

var _ Querier = (*Queries)(nil)
