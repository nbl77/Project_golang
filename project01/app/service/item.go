package service

import (
  "context"
  "project01/app/model"
  "log"
)
var Inventory *model.Item
var InventoryList *model.ItemList

var InventoryListShowAll *model.ItemList


func init()  {
  Inventory = new(model.Item)
  InventoryList = new(model.ItemList)

  InventoryList.ItemList = make([]*model.Item, 0)
}
func idFilterInventory (idItem int32) *model.Item {
  for _,val :=range InventoryList.ItemList{
    if idItem == val.IdItem {
      return val
    }
  }
  return nil
}

func IdItemFilter (idItem int32) int32 {
  for i := 0; i < len(InventoryList.ItemList); i++ {
    InventoryDariList := InventoryList.ItemList[i]
      if idItem == InventoryDariList.IdItem {
        return 1
      } else {
        return 2
      }

  }
  return 0
}

func removeInventory (slice []*model.Item, s int32) []*model.Item {
  return append(slice[:s], slice[s+1:]...)

}

func (*Service) AddItem(ctx context.Context, item *model.Item) (*model.Status, error) {
  item.IdItem = int32(len(InventoryList.ItemList) + 1)
  InventoryList.ItemList = append(InventoryList.ItemList, item)

  res := &model.Status {
    Status : 200,
    Message : "Barang Berhasil Disimpan",
  }
  log.Println("Menambahkan Item...")
  log.Println("ID Item :", item.IdItem,", Nama Item:",item.NamaItem)

  return res, nil
}

func (*Service) TakeItem(ctx context.Context, item *model.Item) (*model.Item, error) {
  InventoryDariClient := idFilterInventory(item.IdItem)
  res := &model.Item {
    IdItem: 0,
  }
    if InventoryDariClient != nil {
      res = InventoryDariClient
      if len(InventoryList.ItemList) == 1 {
        InventoryList.ItemList = make([]*model.Item, 0)
        }else {
          var flag int32
          for key,val:=range InventoryList.ItemList{
            if val.IdItem == item.IdItem {
                flag = int32(key)
                break
            }
          }
          InventoryList.ItemList = removeInventory(InventoryList.ItemList, flag)
        }
    }
  return res,nil

}

func(*Service) Show(ctx context.Context, item *model.Item) (*model.Item, error) {
  InventoryDariClient := idFilterInventory(item.GetIdItem())
  var err error
  if InventoryDariClient != nil {
    return InventoryDariClient, nil
  } else {
    res :=&model.Item {
      IdItem: 0,
    }
    return res, err
  }
}

func(*Service) ShowAll(ctx context.Context, item *model.Item) (*model.ItemList, error) {
  idUser := item.IdUser
  InventoryListShowAll = new(model.ItemList)
  InventoryListShowAll.ItemList =make([]*model.Item, 0)
  for _,val :=range InventoryList.ItemList{
    if val.IdUser == idUser {
      InventoryListShowAll.ItemList = append(InventoryListShowAll.ItemList, val)
    }
  }
  log.Println("Menampilkan Barang User...")
  return InventoryListShowAll,nil
}

func(*Service) ChangeStatus(ctx context.Context, item *model.Item) (*model.Item, error) {
  idItem := item.IdItem
  itemNow := &model.Item{
    IdUser:0,
  }
  for key,val :=range InventoryList.ItemList{
    if val.IdItem == idItem {
      InventoryList.ItemList[key].Status = item.Status
      itemNow = InventoryList.ItemList[key]
    }
  }
  log.Println("Merubah Status Barang User...")
  return itemNow,nil
}

func (*Service) ShowItem(ctx context.Context, empty *model.Empty) (*model.ItemList, error) {
  return InventoryList,nil
}
