package service
import (
  "context"
  "log"
  "fmt"
  "strconv"
  "project01/app/model"
  "project01/app/config"
  // "encoding/json"
)
var ArrKategori = map[int32]string{1:"elektronik",2:"non elektronik"}
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
  if res.IdItem == 0 {
    fmt.Println("ID Item Yang anda masukan tidak tersedia")
  }else {
    fmt.Println("Item Anda adalah ", res)
  }

}
func AddItem (conn model.InventoryClient) {

  var (
    idItem int32 = 0
    namaItem string
    jumlah int32
    kategori int32
    idUser int32 = config.IdUser
  )
  kategori = SelectKategori()
  fmt.Println("Masukkan nama item")
  fmt.Scan(&namaItem)

  fmt.Println("Masukkan jumlah")
  fmt.Scan(&jumlah)

  // fmt.Println("Masukkan kategori")
  // fmt.Scan(&kategori)

  req:= &model.Item{
    IdItem : idItem,
    NamaItem: namaItem,
    Jumlah: jumlah,
    Kategori: kategori,
    IdUser: idUser,
  }
  if kategori > 2 {
    fmt.Println("Kategori Yang Anda Masukan Salah")
    return
  }

  res,err:= conn.AddItem(context.Background(), req)

  if err != nil {
    log.Fatalf("Tidak bisa menerima response terkait Add item", err)
  }

  // fmt.Println("Status Anda adalah ", res.GetStatus())
  fmt.Println(res.GetMessage() + "\nNama Barang :",req.NamaItem,"\nJumlah :",req.Jumlah,"\nKategori :",ArrKategori[req.Kategori] )
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
  fmt.Println(config.IdUser)
  req:= &model.Item{
    IdUser : config.IdUser,
  }

  res, err:= conn.ShowAll(context.Background(), req)
  result := res.ItemList
  if err != nil {
    log.Fatalf("Tidak bisa menerima response terkait Show All", err)
  }
  fmt.Println("List item Anda adalah :\n")
  // for _, val :=range result{
  //   fmt.Println(val)
  // }
  if result == nil {
    fmt.Println("Anda belum menyimpan barangg")
  }else {
    for _,val :=range result{
      fmt.Println(val)
    }
  }
  // rsp,_ := json.Marshal(result)
  // fmt.Println(string(rsp))
}
func SelectKategori() int32 {

  var res string
  fmt.Println("Pilih Kategori :")
  for key,val :=range ArrKategori{
    fmt.Println(key,val)
  }
  fmt.Println("===============")
  fmt.Scan(&res)
  res2,_ := strconv.Atoi(res)
  return int32(res2)
}
