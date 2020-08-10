package service
import (
  "context"
  "log"
  "project01/app/model"
  "project01/app/config"
  "encoding/json"
  "strconv"
)
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
    id,_ := strconv.Atoi(resp.Message)
    config.Status = true
    config.IdUser = int32(id)
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
