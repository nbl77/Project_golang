package service

import(
  "project01/app/model"
	"project01/app/config"
  "google.golang.org/grpc"
	"fmt"
	"log"
	"context"
)


func AddKategori(conn model.InventoryClient) {
	var (
		idKategori int32 =0
		namaKategori string
	)

	fmt.Println("Masukkan nama kategori")
	fmt.Scan(&namaKategori)


	req := &model.Kategori{
		IdKategori :idKategori,
		NamaKategori: namaKategori,
	}

	res,err := conn.AddKategori(context.Background(), req)
	if err != nil {
		log.Fatalf("Tidak bisa menerima response terkait add Kategori karena error: ", err)
	}
  fmt.Println("----------------------------")
	fmt.Println(res.GetMessage(),"!!!")
  fmt.Println("----------------------------")
}

func EditKategori(conn model.InventoryClient){
	var(
		idKategori int32
		namaKategori string
	)

	fmt.Println("Masukkan id kategori")
	fmt.Scan(&idKategori)

	fmt.Println("Masukkan nama Kategori")
	fmt.Scan(&namaKategori)

	req := &model.Kategori{
		IdKategori: idKategori,
		NamaKategori : namaKategori,
	}

	res,err := conn.EditKategori(context.Background(), req)
	if err != nil {
		log.Fatalf("Tidak bisa menerima response terkait Edit Kategori dikarenaka karena error ", err)
	}
  fmt.Println("------------------------")
	fmt.Println(res.GetMessage(),"!!!")
  fmt.Println("------------------------")
}

func DeleteKategori(conn model.InventoryClient) {
	var (
		idKategori int32
	)

	fmt.Println("Masukkan id Kategori")
	fmt.Scan(&idKategori)

	req := &model.Kategori {
		IdKategori: idKategori,
	}

	res,err := conn.DeleteKategori(context.Background(), req)
	if err != nil {
		log.Fatalf("Tidak bisa menerima response terkait Delete Kategori dikarenakan error ", err)
	}

  fmt.Println("--------------------------")
	fmt.Println(res.GetMessage())
  fmt.Println("--------------------------")
}

func ShowKategori(conn model.InventoryClient) {
	res,err := conn.ShowKategori(context.Background(),new(model.Empty))
	if err != nil {
		log.Fatalf("Tidak bisa menerima response terkait Show Kategori ", err)
	}


	if len(res.KategoriList) < 1 {
		fmt.Println("Belum ada kategori yang disimpan")
	} else {
    fmt.Println("--------------------------------")
		fmt.Println("Kategori : ")
		for _,val:= range res.KategoriList {
      fmt.Println("--------------------------------")
			fmt.Println("Id Kategori: ", val.IdKategori)
			fmt.Println("Nama Kategori: ", val.NamaKategori)
      fmt.Println("--------------------------------")
		}
	}

}

func MenuKategori (conn model.InventoryClient) {

	var flag bool = false
	var angkaSwitch int

	for !flag {
    Menu()
		fmt.Println("Pilih Opsi :")
		fmt.Scan(&angkaSwitch)

		switch angkaSwitch {
		case 1:
			AddKategori(conn)
			break
		case 2:
			EditKategori(conn)
			break
		// case 3:
		// 	DeleteKategori(conn)
		// 	break
    case 3:
			ShowKategori(conn)
			break
		case 4:
			flag = true
			break

		}
	}
}

func Menu(){
	// fmt.Println("0. Menu Kategori")
  fmt.Println("=====================")
  fmt.Println("Konfigurasi Kategori :")
  fmt.Println("=====================")
	fmt.Println("1. Add Kategori")
	fmt.Println("2. Edit Kategori")
	// fmt.Println("3. Delete Kategori")
	fmt.Println("3. Show Kategori")
	fmt.Println("4. Kembali")
}
func GetKategori() *model.KategoriList {
  port := config.SERVER_PORT
  conn2, err := grpc.Dial(port, grpc.WithInsecure())
  if err != nil {
    log.Fatalf("Could not connect to ",port, err)
  }
  conn := model.NewInventoryClient(conn2)
  resp,err := conn.ShowKategori(context.Background(), new(model.Empty))
  if err != nil {
		log.Fatalf("Terdapat Error Pada Server!\nError :", err)
	}
  return resp
}
