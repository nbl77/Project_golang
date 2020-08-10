package main

import (
  // "context"
  "log"
  "fmt"
  "project01/app/model"
  "project01/client/service"
  "google.golang.org/grpc"
  "project01/app/config"
  // "encoding/json"
)

func ConnectInven() model.InventoryClient {
  port := config.SERVER_PORT
  conn, err := grpc.Dial(port, grpc.WithInsecure())
  if err != nil {
    log.Fatalf("Could not connect to ",port, err)
  }
  return model.NewInventoryClient(conn)
}

func main()  {
  conn := ConnectInven()
  var key string
  for key != "99"{
    fmt.Println("===================================")
    fmt.Println("Pilih Menu : ")
    fmt.Println("===================================")
    if !config.Status {
      fmt.Println("1.Register")
      fmt.Println("2.Login")
    }else {
      fmt.Println("1.Tambah Barang")
      fmt.Println("2.Ambil Barang")
      fmt.Println("3.Lihat Barang")
      fmt.Println("4.Lihat Semua Barang")
      fmt.Println("5.Logout")
    }
    fmt.Println("99.Exit")
    key = ""
    fmt.Scan(&key)
    fmt.Println(key)
    switch key {
    case "1":
      if !config.Status {
        var(
          namaLengkap,
          username,
          password string
        )
        fmt.Print("Masukan Nama Lengkap Anda :")
        fmt.Scan(&namaLengkap)
        fmt.Print("Masukan Username :")
        fmt.Scan(&username)
        fmt.Print("Masukan Password :")
        fmt.Scan(&password)
        fmt.Println(service.Register(conn, namaLengkap, username, password))
      }else {
        service.AddItem(conn)
      }
      break
    case "2":
      if !config.Status {
        var(
          username,
          password string
        )
        fmt.Print("Masukan Username :")
        fmt.Scan(&username)
        fmt.Print("Masukan Password :")
        fmt.Scan(&password)
        fmt.Println(service.Login(conn, username, password))
      }else {
        service.AmbilItem(conn)
      }
      break
    case "3":
      if config.Status {
        service.ShowPerItem(conn)
      }
    case "4":
      if config.Status {
        service.ShowAll(conn)
      }
      break
    case "5":
      if config.Status {
        fmt.Println(service.Logout(conn))
      }
      break
    default:
      fmt.Println("Opsi yang anda masukan salah!")
      break
    }
  }

}
