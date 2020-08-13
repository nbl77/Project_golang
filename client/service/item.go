package service
import (
  "context"
  "log"
  "fmt"
  "strconv"
  "project01/app/model"
  "project01/app/config"
)

func AmbilItem(conn model.InventoryClient){
  var ArrKategori = GetKategori().KategoriList
  var idItem int32
  fmt.Println("Masukkan id item")
  fmt.Scan(&idItem)
  req:= &model.Item{
    IdItem: idItem,
  }
  itemReq,_ := conn.Show(context.Background(),req)
  if (itemReq.IdUser != config.IdUser) || (itemReq.Status != 2) {
    fmt.Println("-----------------------------------")
    fmt.Println("ID Item Yang anda masukan tidak tersedia")
    fmt.Println("-----------------------------------")
  }else {
    res,err:= conn.TakeItem(context.Background(),req)
    if err != nil {
      log.Fatalf("Tidak bisa menerima response terkait Item", err)
    }
    if res.IdItem == 0 {
      fmt.Println("-----------------------------------")
      fmt.Println("ID Item Yang anda masukan tidak tersedia")
      fmt.Println("-----------------------------------")
      }else {
        fmt.Println("-----------------------------------")
        fmt.Println("Item yang Anda Ambil Adalah :")
        fmt.Println("-----------------------------------")
        fmt.Println("ID Barang :",res.IdItem)
        fmt.Println("Nama Barang :",res.NamaItem)
        fmt.Println("Jumlah Barang :",res.Jumlah)
        fmt.Println("Kategori Barang :",ArrKategori[res.Kategori].NamaKategori)
        fmt.Println("-----------------------------------")
      }
  }
}

func AddItem (conn model.InventoryClient) {
  var ArrKategori = GetKategori().KategoriList
  var (
    idItem int32 = 0
    namaItem string
    jumlah int32
    kategori int32
    status int32 = 0
    idUser int32 = config.IdUser
  )
  kategori = SelectKategori()
  fmt.Println("Masukkan nama item")
  fmt.Scan(&namaItem)

  fmt.Println("Masukkan jumlah")
  fmt.Scan(&jumlah)

  req:= &model.Item{
    IdItem : idItem,
    NamaItem: namaItem,
    Jumlah: jumlah,
    Kategori: kategori,
    Status: status,
    IdUser: idUser,
  }
  if int(kategori) > len(ArrKategori) {
    fmt.Println("Kategori Yang Anda Masukan Salah")
    return
  }
  res,err:= conn.AddItem(context.Background(), req)
  if err != nil {
    log.Fatalf("Tidak bisa menerima response terkait Add item", err)
  }

  fmt.Println("-----------------------------------")
  fmt.Println(res.GetMessage() + "\nNama Barang :",req.NamaItem,"\nJumlah :",req.Jumlah,"\nKategori :",ArrKategori[req.Kategori].NamaKategori )
  fmt.Println("Menunggu Persetujuan Admin.")
  fmt.Println("-----------------------------------")
}
func ShowPerItem(conn model.InventoryClient){
  var ArrKategori = GetKategori().KategoriList
  var idItem int32
  fmt.Println("Masukkan id item")
  fmt.Scan(&idItem)

  req:= &model.Item{
    IdItem:idItem,
  }

  // hahahhahahahahahahaa
  res,err := conn.Show(context.Background(),req)

  if err != nil {
    log.Fatalf("Tidak bisa menerima response terkait Show", err)
  }
  if res.IdItem == 0 {
    fmt.Println("-----------------------------------")
    fmt.Println("ID Item Yang anda masukan tidak tersedia")
    }else {
      fmt.Println("-----------------------------------")
      fmt.Println("Item yang Anda cari Adalah :")
      fmt.Println("-----------------------------------")
      fmt.Println("ID Barang :",res.IdItem)
      fmt.Println("Nama Barang :",res.NamaItem)
      fmt.Println("Jumlah Barang :",res.Jumlah)
      fmt.Println("Status Barang :",GetStatusItem(res.Status))
      fmt.Println("Kategori Barang :",ArrKategori[res.Kategori].NamaKategori)
    }
}
func ShowAll(conn model.InventoryClient) {
  var ArrKategori = GetKategori().KategoriList
  req:= &model.Item{
    IdUser : config.IdUser,
  }
  res, err:= conn.ShowAll(context.Background(), req)
  result := res.ItemList
  if err != nil {
    log.Fatalf("Tidak bisa menerima response terkait Show All", err)
  }
  fmt.Println("-----------------------------------")
  fmt.Println("List item Anda adalah :\n")
  if result == nil {
    fmt.Println("Anda belum menyimpan barangg")
  }else {
    fmt.Println("Barang Anda : ")
    for _,val :=range result{
      fmt.Println("------------------------------")
      fmt.Println("ID Barang :",val.IdItem)
      fmt.Println("Nama Barang :",val.NamaItem)
      fmt.Println("Jumlah Barang :",val.Jumlah)
      fmt.Println("Status Barang :",GetStatusItem(val.Status))
      fmt.Println("Kategori Barang :",ArrKategori[val.Kategori].NamaKategori)
    }
  }
}
func GetAllItem(conn model.InventoryClient)  {
  var ArrKategori = GetKategori().KategoriList
  resp, err:= conn.ShowItem(context.Background(), new(model.Empty))
  if err != nil {
    log.Fatalf("Tidak bisa menerima response terkait Show All", err)
  }
  if len(resp.ItemList) < 1 {
    fmt.Println("-----------------------------------")
    fmt.Println("Belum ada Item Yang Di Simpan")
    fmt.Println("-----------------------------------")
  }else {
    fmt.Println("-----------------------------------")
    fmt.Println("Barang yang telah Di Simpan :")
    for _,val :=range resp.ItemList{
      fmt.Println("-----------------------------------")
      fmt.Println("ID Item :",val.IdItem)
      fmt.Println("Nama Item :",val.NamaItem)
      fmt.Println("Jumlah :",val.Jumlah)
      fmt.Println("Status Barang :",GetStatusItem(val.Status))
      fmt.Println("Kategori :",ArrKategori[val.Kategori].NamaKategori)
      fmt.Println("Pemilik :",GetSingle(conn, val.IdUser).NamaLengkap)
    }
  }
}
func SelectKategori() int32 {
  var ArrKategori = GetKategori().KategoriList
  var res string
  fmt.Println("Pilih Kategori :")
  for key,val :=range ArrKategori{
    fmt.Println(strconv.Itoa(key + 1)+".",val.NamaKategori)
  }
  fmt.Println("-----------------------------------")
  fmt.Scan(&res)
  res2,_ := strconv.Atoi(res)
  return int32(res2 - 1)
}

func FilterItemByKat(conn model.InventoryClient)  {
  var ArrKategori = GetKategori().KategoriList
  kategori := SelectKategori()
  resp, err:= conn.ShowItem(context.Background(), new(model.Empty))
  if err != nil {
    log.Fatalf("Tidak bisa menerima response terkait Show All", err)
  }
  if len(resp.ItemList) < 1 {
    fmt.Println("-----------------------------------")
    fmt.Println("Belum ada Item Yang Di Simpan")
    fmt.Println("-----------------------------------")
  }else {
    fmt.Println("-----------------------------------")
    fmt.Println("Barang yang telah Di Simpan :")
    for _,val :=range resp.ItemList{
      if val.Kategori == kategori {
        fmt.Println("-----------------------------------")
        fmt.Println("ID Item :",val.IdItem)
        fmt.Println("Nama Item :",val.NamaItem)
        fmt.Println("Jumlah :",val.Jumlah)
        fmt.Println("Status Barang :",GetStatusItem(val.Status))
        fmt.Println("Kategori :",ArrKategori[val.Kategori].NamaKategori)
        fmt.Println("Pemilik :",GetSingle(conn, val.IdUser).NamaLengkap)
      }
    }
  }
}

func MenuChangeStatus(conn model.InventoryClient)  {
  var key string
  var id string
  var flag = true
  for flag{
    fmt.Println("Pilih Aksi :")
    fmt.Println("1.Terima Barang")
    fmt.Println("2.Terima Semua Barang User")
    fmt.Println("3.Tolak Barang")
    fmt.Println("4.Tolak Semua Barang User")
    fmt.Println("5.Kembali")
    fmt.Scan(&key)
    switch key {
    case "1":
      fmt.Println("Pilih Barang Yang Akan Diterima :")
      ShowWaiting(conn,2)
      fmt.Println("Masukan Id Item :")
      fmt.Scan(&id)
      ChangeStatus(conn, id, 2)
      break
    case "2":
      ShowUser(conn)
      fmt.Println("Pilih ID User :")
      fmt.Scan(&id)
      ChangeStatusUser(conn,id,2)
      break
    case "3":
      fmt.Println("Pilih Barang Yang Akan Ditolak :")
      ShowWaiting(conn,1)
      fmt.Println("Masukan Id Item :")
      fmt.Scan(&id)
      ChangeStatus(conn, id, 1)
      break
    case "4":
      ShowUser(conn)
      fmt.Println("Pilih ID User :")
      fmt.Scan(&id)
      ChangeStatusUser(conn,id,1)
      break
    case "5":
      flag = false
      break
    }
  }

}



func ChangeStatus(conn model.InventoryClient,IdItem string, stat int32)  {
  Id,_ := strconv.Atoi(IdItem)
  item := &model.Item{
    IdItem : int32(Id),
    Status: stat,
  }
  res,_ := conn.Show(context.Background(),item)
  if (item.IdItem != res.IdItem) {
    fmt.Println("-----------------------------------")
    fmt.Println("ID Item Yang anda masukan tidak tersedia")
    fmt.Println("-----------------------------------")
  }else {
    resp, err:= conn.ChangeStatus(context.Background(), item)
    if err != nil {
      log.Fatalf("Tidak bisa menerima response", err)
    }
    fmt.Println("--------------------------------")
    fmt.Println("Barang",GetStatusItem(resp.Status))
    fmt.Println("-----------------------------------")
  }
}

func ChangeStatusUser(conn model.InventoryClient,IdUser string, stat int32)  {
  Id,_ := strconv.Atoi(IdUser)
  item := &model.Item{
    IdUser : int32(Id),
    Status: stat,
  }
  res,_ := conn.ShowAll(context.Background(),item)
  if len(res.ItemList) < 1 {
    fmt.Println("-----------------------------------")
    fmt.Println("Item tidak tersedia")
    fmt.Println("-----------------------------------")
  }else {
    for _,val :=range res.ItemList{
      item = &model.Item{
        IdItem:val.IdItem,
        Status:stat,
      }
      conn.ChangeStatus(context.Background(), item)
    }
    fmt.Println("--------------------------------")
    fmt.Println("Barang",GetStatusItem(stat))
    fmt.Println("-----------------------------------")
  }
}

func ShowWaiting(conn model.InventoryClient, stat int32)  {
  var ArrKategori = GetKategori().KategoriList
  resp, err:= conn.ShowItem(context.Background(), new(model.Empty))
  if err != nil {
    log.Fatalf("Tidak bisa menerima response terkait Show All", err)
  }
  for _,val :=range resp.ItemList{
    if val.Status != stat {
      fmt.Println("-----------------------------------")
      fmt.Println("ID Item :",val.IdItem)
      fmt.Println("Nama Item :",val.NamaItem)
      fmt.Println("Jumlah :",val.Jumlah)
      fmt.Println("Status Barang :",GetStatusItem(val.Status))
      fmt.Println("Kategori :",ArrKategori[val.Kategori].NamaKategori)
      fmt.Println("Pemilik :",GetSingle(conn, val.IdUser).NamaLengkap)
      fmt.Println("-----------------------------------")
    }
  }
}

func GetStatusItem(status int32) string {
  strStatus := ""
  if status == 0 {
    strStatus = "Menunggu Persetujuan"
  }else if status == 1 {
    strStatus = "Ditolak"
  }else if status == 2 {
    strStatus = "Diterima"
  }
  return strStatus
}
