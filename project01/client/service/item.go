package service
import (
  "context"
  "log"
  "fmt"
  "project01/app/model"
  "project01/app/config"
)

func AmbilItem(conn model.InventoryClient){
  var (
    idItem int32
  )

  fmt.Println("Masukkan id item")
  fmt.Scan(&idItem)

  req:= &model.Item{
    IdItem: idItem,
  }

  res,err:= conn.GetItem(context.Background(),req)

  if err != nil {
    log.Fatalf("Tidak bisa menerima response terkait Get Item", err)
  }

  fmt.Println("Item Anda adalah ", res)
}
func AddItem (conn model.InventoryClient) {

  var (
    idItem int32 = 0
    namaItem string
    jumlah int32
    kategori int32
    idUser int32 = config.IdUser
  )
  fmt.Println("Masukkan nama item")
  fmt.Scan(&namaItem)

  fmt.Println("Masukkan jumlah")
  fmt.Scan(&jumlah)

  fmt.Println("Masukkan kategori")
  fmt.Scan(&kategori)

  req:= &model.Item{
    IdItem : idItem,
    NamaItem: namaItem,
    Jumlah: jumlah,
    Kategori: kategori,
    IdUser: idUser,
  }

  res,err:= conn.AddItem(context.Background(), req)

  if err != nil {
    log.Fatalf("Tidak bisa menerima response terkait Add item", err)
  }

  // fmt.Println("Status Anda adalah ", res.GetStatus())
  fmt.Println("Message Anda adalah ", res.GetMessage())
}
func ShowPerItem(conn model.InventoryClient){
  var idItem int32

  fmt.Println("Masukkan id item")
  fmt.Scan(&idItem)

  req:= &model.Item{
    IdItem:idItem,
  }

  res,err := conn.Show(context.Background(),req)

  if err != nil {
    log.Fatalf("Tidak bisa menerima response terkait Show", err)
  }
  fmt.Println("Item yang Anda cari Adalah ", res)
}
func ShowAll(conn model.InventoryClient) {
  req:= &model.Item{
    IdUser : config.IdUser,
  }

  res, err:= conn.ShowAll(context.Background(), req)

  if err != nil {
    log.Fatalf("Tidak bisa menerima response terkait Show All", err)
  }

  fmt.Println("List item Anda adalah ", res)
}
