package service
import (
  "context"
  "log"
  "fmt"
  "project01/app/model"
  "project01/app/config"
  "strconv"
)

func GetSingle(conn model.InventoryClient, id_user int32) *model.User {
  user := &model.User{
    IdUser: id_user,
  }
  resp,err := conn.GetUser(context.Background(),user)
  if err != nil {
    log.Fatalf(err.Error())
  }
  return resp
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

func ShowUser(conn model.InventoryClient) {
  resp,err := conn.ShowUser(context.Background(), new(model.Empty))
  if err != nil {
    log.Fatalf(err.Error())
  }
  if len(resp.UserList) < 1 {
    fmt.Println("Belum ada user yang mendaftar")
  }else {
    fmt.Println("User yang telah mendaftar :")
    for _,val :=range resp.UserList{
      fmt.Println("========================")
      fmt.Println("ID User :",val.IdUser)
      fmt.Println("Nama User :",val.NamaLengkap)
      fmt.Println("Username :",val.Username)
    }
  }
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
    id,_ := strconv.Atoi(resp.Message)
    config.Status = true
    config.IdUser = int32(id)
    config.Tipe = "user"
    if (user.Username == "admin") && (user.Password == "admin") {
      config.Tipe = "admin"
    }
    resp.Message = "Berhasil Login"
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
