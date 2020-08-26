package security

import (
	"context"
	models "go-rest-api/modules/security/secumodel"
)

func (s *securityStorage) FindByUsername(ctx context.Context, userName string) (models.Security, error) {
	var user models.Security

	if ok := s.db.Find(&user).Where("username = ?", userName).Error; ok != nil {
		//TODO
		return user, ok
	}

	return user, nil
}
