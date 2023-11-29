package usecase

import (
	"context"
	"sync"

	"github.com/lokks307/adr-boilerplate/domain"
	"github.com/lokks307/adr-boilerplate/models"
	"github.com/sirupsen/logrus"
)

// Notice - 싱글톤보다 더 나은 방법이 있다면 수정해주세요.
// 이 객체를 선언할 곳을 마땅히 찾지 못해 싱글톤으로 구현합니다.
// customerUsecase := usecase.NewCustomerUsecase()

var lock = sync.Mutex{}

type customerUsecase struct {
}

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

func NewCustomerUsecase() domain.CustomerUsecase {
	return getInstance()
}

func (a *customerUsecase) ReadCustomerByID(cid int64) (*models.Customer, error) {
	// exec boil.ContextExecutor
	exec := domain.Conn()
	ctx := context.Background()
	found, err := models.FindCustomer(ctx, exec, cid)
	if err != nil {
		return nil, err
	}

	return found, nil
}
