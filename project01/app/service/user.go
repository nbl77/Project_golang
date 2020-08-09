package service

import (
  "context"
  "project01/app/model"
)
func (Service) Register(ctx context.Context, item *model.Item) (*model.Status, error) {
  return new(model.Status), nil
}
