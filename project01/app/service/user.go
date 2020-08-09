package service

import (
  "context"
  "log"
  "project01/app/model"
  "regexp"
)
var storageUserList *model.UserList
var session *model.Session

type ServiceUser struct {}

func init()  {
  storageUserList = new(model.UserList)
  session = new(model.Session)
  storageUserList.UserList = make([]*model.User, 0)
}
func (ServiceUser) Register(ctx context.Context, user *model.User) (*model.Status, error) {
  cek := true
  user.IdUser = int32(len(storageUserList.UserList))
  status := new(model.Status)
  length := len(storageUserList.UserList)
  for _,val :=range storageUserList.UserList{
    if user.Username == val.Username {
      cek = false
    }
  }

  userRegex,_ := regexp.MatchString(`^(?=.*[0-9])(?=.*[a-z]).{8,}$`,user.Username)
  passwordRegex,_ := regexp.MatchString(`^(?=.*[0-9])(?=.*[a-z])(?=.*[A-Z]).{10,}$`,user.Password)
  if !userRegex {
    status = &model.Status{
      Status : 400,
      Message : "Username Minimal 8 karakter dan 1 angka",
    }
    return status, nil
  }

  if !passwordRegex {
    status = &model.Status{
      Status : 400,
      Message : "Password Minimal 10 karakter, 1 Kapital dan 1 angka",
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
func (ServiceUser) ShowUser(ctx context.Context, empty *model.Empty) (*model.UserList, error) {
  return storageUserList,nil
}
func (ServiceUser) Login(ctx context.Context, user *model.User) (*model.Status, error){
  status := new(model.Status)
  login := false
  for _, val:=range storageUserList.UserList{
      if (val.Username == user.Username) && (val.Password == user.Password) {
        login = true
        break
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
        Message : "Berhasil Login",
      }
      log.Println("Success Registered!")
    }
    return status, nil
  }
