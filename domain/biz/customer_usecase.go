package biz

import (
	"context"
	"sync"

	"github.com/sirupsen/logrus"
	"github.com/volatiletech/sqlboiler/v4/boil"
	. "github.com/volatiletech/sqlboiler/v4/queries/qm"

	"github.com/lokks307/adr-boilerplate/domain"
	"github.com/lokks307/adr-boilerplate/e"
	"github.com/lokks307/adr-boilerplate/models"
)

var lock = sync.Mutex{}

type customerUsecase struct{}

var customerInstance *customerUsecase

func getInstance() *customerUsecase {
	if customerInstance == nil {
		lock.Lock()
		defer lock.Unlock()

		if customerInstance == nil {
			logrus.Info("Create CustomerUsecase Instance")
			customerInstance = &customerUsecase{}
		} else {
			// "CustomerUsecase Instance is already created"
		}
	}

	return customerInstance
}

func Customer() domain.CustomerUsecase {
	return getInstance()
}

func (a *customerUsecase) InsertCustomer(
	firstName string,
	lastName string,
	email string,
) error {
	ctx := context.Background()

	newCustomer := &models.Customer{
		FirstName: firstName,
		LastName:  lastName,
		Email:     email,
	}

	err := newCustomer.Insert(ctx, boil.GetContextDB(), boil.Infer())
	if err != nil {
		return e.ErrorWrap(err)
	}

	return nil
}

func (a *customerUsecase) ReadCustomerByID(cid int64) (*models.Customer, error) {
	ctx := context.Background()

	customerModel, err := models.Customers(
		Select(),
		models.CustomerWhere.CustomerId.EQ(cid),
	).One(ctx, boil.GetContextDB())

	// noRow를 에러로 처리하지 않을 경우
	// if errors.Is(err, sql.ErrNoRows) {
	// 	return nil, e.ErrorWrap(err)
	// }
	if err != nil {
		return nil, e.ErrorWrap(err)
	}

	return customerModel, nil
}
