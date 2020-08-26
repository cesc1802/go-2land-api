package security

import (
	"context"
	"errors"
	models "go-rest-api/modules/security/secumodel"
)

func (s *securityStorage) Create(ctx context.Context, security models.Security) error {
	db := s.db.New().Begin()

	if err := db.Table(security.TableName()).Create(&security).Error; err != nil {
		db.Rollback()
		return errors.New("cannot insert new user")
	}

	if err := db.Commit().Error; err != nil {
		db.Rollback()
		return errors.New("commit fail")
	}

	return nil
}
