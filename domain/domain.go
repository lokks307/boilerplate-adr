package domain

import (
	"context"

	"github.com/lokks307/adr-boilerplate/models"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

func ReadCustomerByID(exec boil.ContextExecutor, cid int64) (*models.Customer, error) {
	ctx := context.Background()
	found, err := models.FindCustomer(ctx, exec, cid)
	if err != nil {
		return nil, err
	}

	return found, nil
}
