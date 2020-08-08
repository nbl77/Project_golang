package main

import (
  "net"
  "log"
  "strconv"
  "context"
  "time"
  "math/rand"
  "google.golang.org/grpc"
  "github.com/golang/protobuf/ptypes/empty"
  "Day5/model"
  "strings"
)

var storageQueue *model.QueueList //nil
var storageParkir *model.Parkir //nil

// func init()  {
//   storageQueue = new(model.QueueList)
//   storageQueue.QList = make([]*model.Queue, 0)
//   storageParkir = new(model.Parkir)
// }

type ParkirService struct{}

func (ParkirService) GetID(ctx context.Context, empty *empty.Empty) (*model.Queue, error) {
  queue := &model.Queue{
    Id:strconv.Itoa(rand.Intn(1000000)),
    Time: time.Now().Unix(),
  }
  storageQueue.QList = append(storageQueue.QList, queue)
  log.Println("Add Queue :",queue.String())
  return queue, nil
}

func (ParkirService) GetList(ctx context.Context, empty *empty.Empty) (*model.QueueList, error) {
  return storageQueue,nil
}

func (ParkirService) Exit(ctx context.Context, parkir *model.Parkir) (*model.Result, error) {
  res := false
  key := 0
  for k,val :=range storageQueue.QList{
    if parkir.Id == val.Id {
      res = true
      key = k
      break
    }
  }
  if !res {
    result := &model.Result{
      Status: 404,
      Message: "ID Tidak Di Temukan",
      Second: "0",
      Total: "0",
    }
    return result,nil
  }
  timeQueue := storageQueue.QList[key]
  now := int(time.Now().Unix()) - int(timeQueue.Time)
  msg,total := getTotal(parkir.Tipe, int(now))

  if total == 0 {
    result := &model.Result{
      Status: 404,
      Message: msg,
      Second: "0",
      Total: "0",
    }
    return result,nil
  }

  result := &model.Result{
    Status: 200,
    Message: "Anda Parkir Selama",
    Second: strconv.Itoa(now) + " detik",
    Total: strconv.Itoa(total),
  }
  RemoveIndex(storageQueue, key)
  return result,nil
}

func RemoveIndex(ql *model.QueueList, index int)  {
  ql.QList = append(ql.QList[:index],ql.QList[index + 1:]...)
}

func getTotal(tipe string,times int) (string, int) {
  tipe = strings.ToLower(tipe)
  total := 0
  txt := ""
  switch tipe {
  case "mobil":
    for i := 0; i < times; i++ {
        if i < 1 {
          total += 5000
          continue
        }
        total += 3000
    }
  case "motor":
    for i := 0; i < times; i++ {
        if i < 1 {
          total += 3000
          continue
        }
        total += 2000
    }
  default:
    txt = "Tipe kendaraan yang anda masukan tidak tersedia"
  }
  return txt,total
}

func main()  {
  srv := grpc.NewServer()
  var parkServ ParkirService
  model.RegisterParkirServiceServer(srv, parkServ)
  log.Println("Starting RPC Server at ", "9000")
  l, err := net.Listen("tcp",":9000")
  if err != nil {
    log.Fatalf("Could not listen to %s %v :",":9000",err)
  }
  log.Fatal(srv.Serve(l))
}
