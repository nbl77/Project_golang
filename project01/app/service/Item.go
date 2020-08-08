package service

import (
  "context"
  "project01/app/model"
)

var Inventory *model.Item
var InventoryList *model.ItemList


func init()  {
  Inventory = new(model.Item)
  InventoryList = new(model.ItemList)
  InventoryList.ItemList = make([]*model.Item, 0)
}

type ServiceItem struct{}

func (ServiceItem) AddItem(ctx context.Context, item *model.Item) (*model.Empty, error) {
  items := &model.Item{
    IdItem : 1,
    NamaItem : "Komputer",
    Jumlah : 10,
    Kategori : 5,
  }
  InventoryList.ItemList = append(InventoryList.ItemList, items)
  return new(model.Empty), nil
}

func (ServiceItem) ShowAll(ctx context.Context, empty *model.Empty) (*model.ItemList, error) {
  return InventoryList, nil
}
