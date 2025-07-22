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

	// Fetch(cursor string, num int64) (models.CustomerSlice, string, error)
	// Update(customer *models.Customer) error
	// GetByTitle(title string) (*models.Customer, error)
	// Store(customer *models.Customer) error
	// Delete(id int64) error
}

// TODO: 각 서로 다른 속성의 객체에 접근하는 도메인 함수들끼리의 참조가 필요하면
//  전체 도메인에 대해 단일 인터페이스를 사용할 수도 있음

// usercase에 비즈니스 로직을 넣을 예정이라면, 이 domain 레벨에는 각 CRUD 함수의 유닛들을 위치시키고
// biz에는 각 모듈 함수들을 가져다가 비즈니스 로직을 만들도록 구현하는것이 덜 복잡할 수 있음
