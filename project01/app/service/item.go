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

// func ConnectItem(srv *grpc.Server){
//   var service **Service
//   model.RegisterInventoryServer(srv, service)
// }

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
  // hasilFilterId := idFilterInventory(item.GetIdItem())
  item.IdItem = int32(len(InventoryList.ItemList) + 1)
  // if hasilFilterId != nil {
  //   res := &model.Status {
  //     Status: 400,
  //     Message: "Id Item sudah ada yang sama",
  //   }
  //
  //   return res, nil
  //
  // } else if hasilFilterId == nil{
  //   // Inventory.NamaItem = item.NamaItem
  //   // Inventory.IdItem = item.IdItem
  //   // Inventory.Jumlah = item.Jumlah
  //   // Inventory.Kategori = item.Kategori
  // }
  // res := &model.Status {
  //   Status : 404,
  //   Message : "Not Found",
  // }
  // return res, nil
  InventoryList.ItemList = append(InventoryList.ItemList, item)

  res := &model.Status {
    Status : 200,
    Message : "Barang Berhasil Disimpan",
  }
  log.Println("Menambahkan Item...")
  log.Println("ID Item :", item.IdItem,", Nama Item:",item.NamaItem)

  return res, nil
}

func (*Service) GetItem(ctx context.Context, item *model.Item) (*model.Item, error) {
  InventoryDariClient := idFilterInventory(item.IdItem)
  // InventoryMauDihapus:= InventoryDariClient
  // var err error
  // if InventoryDariClient != nil {
  //
  //
  //   return res, nil
  // } else {
  //   res:= &model.Item {
  //     IdItem:0,
  //   }
  //   return res, err
  // }
  // if len(InventoryList.ItemList) == 1 {
  //   InventoryList.ItemList = make([]*model.Item, 0)
  // }else {
  // }
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
  // for i := 0; i < len(InventoryList.ItemList); i++ {
  //
  //   InventoryDariClient:= InventoryList.ItemList[i]
  //   if idUser == InventoryDariClient.IdUser {
  //     InventoryListShowAll.ItemList = append(InventoryListShowAll.ItemList, InventoryDariClient)
  //   }
  //
  // }
  // res:= &model.ItemList{
  //   ItemList:InventoryListShowAll.ItemList,
  // }
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

func (*Service) ShowItem(ctx context.Context, empty *model.Empty) (*model.ItemList, error) {
  return InventoryList,nil
}
