package api

import (
	"context"
	"github.com/stashapp/stash/pkg/models"
	"time"
)

func (s stashIDResolver) UpdatedAt(ctx context.Context, obj *models.StashID) (*time.Time, error) {
	return obj.UpdatedAt.Ptr(), nil
}

func (s stashIDInputResolver) UpdatedAt(ctx context.Context, obj *models.StashID, data *time.Time) error {
	return nil
}
