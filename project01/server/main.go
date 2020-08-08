package main

import (
  "net"
  "log"
  "google.golang.org/grpc"
  "project01/app/model"
  "project01/app/config"
  "project01/app/service"
)

func main()  {
  var InventoryServ *service.ServiceItem
  srv := grpc.NewServer()
  model.RegisterInventoryServer(srv, InventoryServ)
  log.Println("Starting RPC Server at ", config.SERVER_PORT)
  l, err := net.Listen("tcp",config.SERVER_PORT)
  if err != nil {
    log.Fatalf("Could not listen to %s %v :",config.SERVER_PORT,err)
  }
  log.Fatal(srv.Serve(l))
}
