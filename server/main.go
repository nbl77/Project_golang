package main

import (
  "net"
  "log"
  "google.golang.org/grpc"
  "project01/app/config"
  "project01/app/service"
  "project01/app/model"
)

func main()  {
  srv := grpc.NewServer()
  var service *service.Service
  model.RegisterInventoryServer(srv, service)
  // service.ConnectUser(srv)
  // service.ConnectItem(srv)
  // service.NyambunginItem()
  log.Println("Starting RPC Server at ", config.SERVER_PORT)
  l, err := net.Listen("tcp",config.SERVER_PORT)
  if err != nil {
    log.Fatalf("Could not listen to %s %v :",config.SERVER_PORT,err)
  }
  log.Fatal(srv.Serve(l))
}
