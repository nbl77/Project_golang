package service

import (
  "context"
  "log"
  "project01/app/model"
  "regexp"
  "strconv"
)
var storageUserList *model.UserList
var session *model.Session

type Service struct {
}


// func ConnectUser(srv *grpc.Server){
//   var service *Service
//   model.RegisterInventoryServer(srv, service)
// }

func init()  {
  storageUserList = new(model.UserList)
  session = new(model.Session)
  storageUserList.UserList = make([]*model.User, 0)
}
func (*Service) Register(ctx context.Context, user *model.User) (*model.Status, error) {
  cek := true
  user.IdUser = int32(len(storageUserList.UserList))
  status := new(model.Status)
  length := len(storageUserList.UserList)
  for _,val :=range storageUserList.UserList{
    if user.Username == val.Username {
      cek = false
    }
  }

  userRegex,_ := regexp.MatchString(`^\w.{7,}$`,user.Username)
  passwordRegex,_ := regexp.MatchString(`[\w].{9,}$`,user.Password)
  if !userRegex {
    status = &model.Status{
      Status : 400,
      Message : "Username Minimal 8 karakter",
    }
    return status, nil
  }

  if !passwordRegex {
    status = &model.Status{
      Status : 400,
      Message : "Password Minimal 10 karakter",
    }
    return status, nil
  }

  if !cek {
    status = &model.Status{
      Status : 400,
      Message : "Username Telah Di Gunakan",
    }
    return status, nil
  }
  if user.Username == "" {
    status = &model.Status{
      Status : 400,
      Message : "Username Tidak Boleh Kosong",
    }
    return status, nil
  }
  if user.Password == "" {
    status = &model.Status{
      Status : 400,
      Message : "Password Tidak Boleh Kosong",
    }
    return status, nil
  }
  storageUserList.UserList = append(storageUserList.UserList, user)
  if len(storageUserList.UserList) > length {
    status = &model.Status{
      Status : 200,
      Message : "Berhasil Register",
    }
  }else {
    status = &model.Status{
      Status : 400,
      Message : "Gagal Register",
    }
  }
  log.Println("Success Registered!")
  return status, nil
}
func (*Service) ShowUser(ctx context.Context, empty *model.Empty) (*model.UserList, error) {
  return storageUserList,nil
}
func (*Service) Login(ctx context.Context, user *model.User) (*model.Status, error){
  status := new(model.Status)
  login := false
  kun := 0
  if (user.Username == "admin") && (user.Password == "admin") {
    login = true
  }else {
    for key, val:=range storageUserList.UserList{
      if (val.Username == user.Username) && (val.Password == user.Password) {
        login = true
        kun = key
        break
      }
    }
  }
    if !login {
      status = &model.Status{
        Status : 400,
        Message : "Username Atau Password Salah",
      }
    }else {
      status = &model.Status{
        Status : 200,
        Message : strconv.Itoa(kun),
      }
      log.Println("Success Login!")
    }
    return status, nil
  }
func(*Service) GetUser(context context.Context, user *model.User) (*model.User, error) {
  flag := false
  id_user := 0
  for key,val:=range storageUserList.UserList{
    if user.IdUser == val.IdUser {
      flag = true
      id_user = key
      break
    }
  }
  user = &model.User{}
  if flag {
    user= storageUserList.UserList[id_user]
  }
  return user,nil
}
