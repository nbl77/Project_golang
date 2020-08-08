package main

import (
  "net"
  "log"
  "context"
  "google.golang.org/grpc"
  "github.com/golang/protobuf/ptypes/empty"
  "project01/app/model"
)

var Inventory *model.Item //nil


func init()  {
  Inventory = new(model.Item)
}

type ServiceItem struct{}

func (ServiceItem) ShowAll(ctx context.Context, empty *empty.Empty) (*model.Item, error)  {
  item := &model.Item{
    IdItem : 1,
    NamaItem : "Komputer",
    Jumlah : 10,
    Kategori : 5,
  }
  Inventory = item
  return item, nil
}

func main()  {
  srv := grpc.NewServer()
  var InventoryServ ServiceItem
  model.RegisterInventoryServer(srv, InventoryServ)
  log.Println("Starting RPC Server at ", "9000")
  l, err := net.Listen("tcp",":9000")
  if err != nil {
    log.Fatalf("Could not listen to %s %v :",":9000",err)
  }
  log.Fatal(srv.Serve(l))
}
