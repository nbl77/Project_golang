package main

import (
  "context"
  "log"
  "fmt"
  "project01/app/model"
  "google.golang.org/grpc"
  "project01/app/config"
  "encoding/json"
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
      fmt.Println("1.Logout")
    }
    fmt.Println("99.Exit")
    fmt.Scanln(&key)
    switch key {
    case "1":
      if !config.Status {
        var(
          namaLengkap,
          username,
          password string
        )
        fmt.Print("Masukan Nama Lengkap Anda :")
        fmt.Scanln(&namaLengkap)
        fmt.Print("Masukan Username :")
        fmt.Scanln(&username)
        fmt.Print("Masukan Password :")
        fmt.Scanln(&password)
        fmt.Println(Register(conn, namaLengkap, username, password))
      }else {
        fmt.Println(Logout(conn))
      }
    case "2":
      if !config.Status {
        var(
          username,
          password string
        )
        fmt.Print("Masukan Username :")
        fmt.Scanln(&username)
        fmt.Print("Masukan Password :")
        fmt.Scanln(&password)
        fmt.Println(Login(conn, username, password))
      }
    case "":
    default:
      fmt.Println("Opsi yang anda masukan salah!")
    }
  }

}

func Register(conn model.InventoryClient, namaLengkap,username, password string) string {
  user := &model.User{
    IdUser : 0,
    NamaLengkap : namaLengkap,
    Username: username,
    Password: password,
  }
  resp,err := conn.Register(context.Background(), user)
  if err != nil {
    log.Fatalf(err.Error())
  }
  // res,_ := json.Marshal(resp)
 return resp.Message
}

func ShowUser(conn model.InventoryClient) string {
  resp,err := conn.ShowUser(context.Background(), new(model.Empty))
  if err != nil {
    log.Fatalf(err.Error())
  }
  res,_ := json.Marshal(resp)
  return string(res)
}

func Login(conn model.InventoryClient, username, password string) string {
  user := &model.User{
    IdUser : 0,
    NamaLengkap : "",
    Username: username,
    Password: password,
  }
  resp,err := conn.Login(context.Background(), user)
  if err != nil {
    log.Fatalf(err.Error())
  }
  if resp.Status == 200 {
    config.Status = true
    config.IdUser = user.IdUser
  }
  return resp.Message
}

func Logout(conn model.InventoryClient) string {
  config.Status = false
  config.IdUser = 0
  if !config.Status {
    return "Berhasil Logout"
  }else {
    return "Kesalahan Pada sistem"
  }
}
