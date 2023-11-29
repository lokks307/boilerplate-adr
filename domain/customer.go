package domain

import (
	"github.com/lokks307/adr-boilerplate/models"
)

// CustomerUsecase 인터페이스를 정의합니다.
// 서비스 (또는 테이블)에 대한 동작을 관리할 수 있습니다.
// 복잡한 도메인 비즈니스 로직은 별도의 타입을 정의해서 인터페이스로 관리할 수 있을 것입니다.
// 코드와 인터페이스를 분리하고 어떤 기능이 있는지 쉽게 파악할 수 있습니다.
type CustomerUsecase interface {
	ReadCustomerByID(cid int64) (*models.Customer, error)

	// Fetch(exec boil.ContextExecutor, cursor string, num int64) ([]*models.Customer, string, error)
	// Update(exec boil.ContextExecutor, customer *models.Customer) error
	// GetByTitle(exec boil.ContextExecutor, title string) (*models.Customer, error)
	// Store(exec boil.ContextExecutor, customer *models.Customer) error
	// Delete(exec boil.ContextExecutor, id int64) error
}

// join이나 예외상황이 생겼을 때 어떻게 될까 모르겠다
