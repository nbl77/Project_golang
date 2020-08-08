package main

import (
  "context"
  "log"
  "fmt"
  "project01/app/model"
  "github.com/golang/protobuf/ptypes/empty"
  "google.golang.org/grpc"
)

func GetItem() model.InventoryClient {
  port := ":9000"
  conn, err := grpc.Dial(port, grpc.WithInsecure())
  if err != nil {
    log.Fatalf("Could not connect to ",port, err)
  }
  return model.NewInventoryClient(conn)
}

func main()  {
  gt := GetItem()
  resp,err := gt.ShowAll(context.Background(), new(empty.Empty))
  if err != nil {
    log.Fatalf(err.Error())
  }
  fmt.Println(resp)

}
