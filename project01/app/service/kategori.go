package service



import (
	"project01/app/model"
	"context"
	"fmt"
)

var Category *model.Kategori
var CategoryList *model.KategoriList


func init() {
	Category = new(model.Kategori)
	CategoryList = new(model.KategoriList)

	CategoryList.KategoriList = make([]*model.Kategori,0)
  kat := &model.Kategori{
    IdKategori:1,
    NamaKategori:"Un Kategori",
  }
  CategoryList.KategoriList = append(CategoryList.KategoriList, kat)
}

func idFilterKategory (idKategori int32) *model.Kategori {
	for i := 0; i < len(CategoryList.KategoriList); i++ {
	  KategoriDariList := CategoryList.KategoriList[i]
	  if idKategori == KategoriDariList.IdKategori {
		return KategoriDariList
	  } else {
		return nil
	  }
	}
	return nil
  }

  func IdItemFilterKategori (idKategori int32) int32 {
	for i := 0; i < len(CategoryList.KategoriList); i++ {
	  KategoryDariList := CategoryList.KategoriList[i]
		if idKategori == KategoryDariList.IdKategori {
		  return 1
		} else {
		  return 2
		}

	}
	return 0
  }

  func removeKategori (slice []*model.Kategori, s int32) []*model.Kategori {
	return append(slice[:s], slice[s+1:]...)

  }



func(*Service) AddKategori(ctx context.Context, kategori *model.Kategori) (*model.Status, error) {
	kategori.IdKategori = int32(len(CategoryList.KategoriList) + 1)
	CategoryList.KategoriList = append(CategoryList.KategoriList, kategori)


	res := &model.Status{
		Status:200,
		Message:"Kategori berhasil disimpan",
	}
	fmt.Println("Menambahkan kategori...")
	fmt.Println("ID Kategori : ", kategori.IdKategori, ", Kategori : ", kategori.NamaKategori)

	return res,nil
}

func(*Service) EditKategori(ctx context.Context, kategori *model.Kategori) (*model.Status, error) {
	// KategoriDariClient:= idFilterKategory(kategori.IdKategori)
	flag := false
	res := &model.Status {
		Status:0,
		Message:"",
	}



	for key,val := range CategoryList.KategoriList {
		if val.IdKategori == kategori.IdKategori {
			CategoryList.KategoriList[key] = kategori
			flag = true
			break
		}
	}
	if !flag {
		res = &model.Status {
			Status: 404,
			Message: "Barang tidak ada di database",
		}
			return res, nil

		}

	res = &model.Status {
		Status: 200,
		Message: "Barang berhasil di edit",
	}


	return res, nil
}

func (*Service) DeleteKategori(ctx context.Context, kategori *model.Kategori) (*model.Status, error) {
	KategoriDariClient := idFilterKategory(kategori.IdKategori)
	res := &model.Status {
		Status:0,
		Message:"",
	}

	if KategoriDariClient == nil {
		res = &model.Status {
			Status: 404,
			Message: "Barang tidak ada di database",
		}
			return res, nil

		}

	var flag int32

	for key,val := range CategoryList.KategoriList {
		if val.IdKategori == kategori.IdKategori {
			flag = int32(key)
			break
		}
	}

	CategoryList.KategoriList = removeKategori(CategoryList.KategoriList, flag)


	res = &model.Status {
		Status: 200,
		Message: "Barang berhasil di edit",
	}


	return res, nil
}

func (*Service) ShowKategori(ctx context.Context, empty *model.Empty) (*model.KategoriList, error) {
	return CategoryList,nil
}
