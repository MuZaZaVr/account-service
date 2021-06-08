package repository

import (
	"context"
	"github.com/MuZaZaVr/account-service/internal/model"
	"go.mongodb.org/mongo-driver/mongo"
)

// Repositories contain other repository interfaces
type Repositories struct {
	AccountRepository    Account
	CompanyRepository    Company
	CredentialRepository Credential
}

// NewRepositories is a constructor for Repositories struct
func NewRepositories(db *mongo.Database) *Repositories {
	return &Repositories{
		AccountRepository:    NewAccountRepository(db),
		CompanyRepository:    NewCompanyRepository(db),
		CredentialRepository: NewCredentialRepository(db),
	}
}

// Company represents CRUD-repo for model.CompanyDTO
type Company interface {
	Create(ctx context.Context, company model.CompanyDTO) (string, error)

	FindByName(ctx context.Context, name string) (*model.CompanyDTO, error)
	FindByURL(ctx context.Context, url string) (*model.CompanyDTO, error)

	UpdateName(ctx context.Context, id string, newName string) (string, error)
	UpdateDescription(ctx context.Context, id string, newDescription string) (string, error)
	UpdateURL(ctx context.Context, id string, newUrl string) (string, error)

	Delete(ctx context.Context, id string) (string, error)

	IsExist(ctx context.Context, name string) (bool, error)
}

// Credential represents CRUD-repo for model.CredentialDTO
type Credential interface {
	Create(ctx context.Context, credential model.CredentialDTO) (string, error)

	FindByLogin(ctx context.Context, login string) ([]model.CredentialDTO, error)
	FindByEmail(ctx context.Context, email string) ([]model.CredentialDTO, error)
	FindByPhone(ctx context.Context, phone string) ([]model.CredentialDTO, error)
	FindByName(ctx context.Context, name string) ([]model.CredentialDTO, error)
	FindByMiddleName(ctx context.Context, middlename string) ([]model.CredentialDTO, error)
	FindBySurname(ctx context.Context, surname string) ([]model.CredentialDTO, error)

	UpdateLogin(ctx context.Context, newLogin string) (string, error)
	UpdateEmail(ctx context.Context, newEmail string) (string, error)
	UpdatePhone(ctx context.Context, newPhone string) (string, error)
	UpdateName(ctx context.Context, newName string) (string, error)
	UpdateMiddleName(ctx context.Context, newMiddlename string) (string, error)
	UpdateSurname(ctx context.Context, newSurname string) (string, error)

	Delete(ctx context.Context, id string) (bool, error)
}

// Account represents CRUD-repo for model.AccountDTO
type Account interface {
	Create(ctx context.Context, account model.AccountDTO) (string, error)

	FindByName(ctx context.Context, name string) (account model.AccountDTO, err error)
	FindAllByCompanyId(ctx context.Context, id string) ([]model.AccountDTO, error)
	FindAllByUserID(ctx context.Context, id int) ([]model.AccountDTO, error)

	UpdateName(ctx context.Context, newName string) (string, error)
	UpdateDescription(ctx context.Context, newDescription string) (string, error)
	UpdateCompanyId(ctx context.Context, newCompanyId int) (string, error)

	Delete(ctx context.Context, id string) (bool, error)

	IsExist(ctx context.Context, name string) (bool, error)
}