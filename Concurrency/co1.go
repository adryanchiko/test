package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"encoding/csv"
	"log"
	"net/http"
    "time"
	"bytes"
	"os"
)

type Wilayah struct {
    MuseumID          string      `json:"museum_id"`
    KodePengelolaan   string      `json:"kode_pengelolaan"`
    Nama              string      `json:"nama"`
    Sdm               string      `json:"sdm"`
    AlamatJalan       string      `json:"alamat_jalan"`
    DesaKelurahan     string      `json:"desa_kelurahan"`
    Kecamatan         string      `json:"kecamatan"`
    KabupatenKota     string      `json:"kabupaten_kota"`
    Propinsi          string      `json:"propinsi"`
    Lintang           string      `json:"lintang"`
    Bujur             string      `json:"bujur"`
    Koleksi           string      `json:"koleksi"`
    SumberDana        string      `json:"sumber_dana"`
    Pengelola         string      `json:"pengelola"`
    Tipe              string      `json:"tipe"`
    Standar           string      `json:"standar"`
    TahunBerdiri      string      `json:"tahun_berdiri"`
    Bangunan          string      `json:"bangunan"`
    LuasTanah         string      `json:"luas_tanah"`
    StatusKepemilikan interface{} `json:"status_kepemilikan"`
}

var path string

func main() {

	url := "http://jendela.data.kemdikbud.go.id/api/index.php/CcariMuseum/searchGET?kode_kab_kota=016200"

	spaceClient := http.Client{
		Timeout: time.Second * 60, // Maximum of 1 Minute
	}

    //Build Request
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Set("User-Agent", "spacecount-tutorial")

	res, getErr := spaceClient.Do(req)
	if getErr != nil {
		log.Fatal(getErr)
	}

	body, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		log.Fatal(readErr)
    }
    body = bytes.TrimPrefix(body, []byte("\xef\xbb\xbf"))

    resp := struct {
		Data []Wilayah `json:"data"`
	}{}
    
	jerr := json.Unmarshal(body, &resp)
	if jerr != nil {
		fmt.Println("error:", jerr)
	}

	var kabkota string
	
	for _,d := range resp.Data {
		var f *os.File
		var err error
		if kabkota != d.KabupatenKota {
			f, err = os.Create("./"+d.KabupatenKota+".csv")
			if err != nil {
				fmt.Println(err)
			}
		} else {
			f, err = os.OpenFile("./"+d.KabupatenKota+".csv", os.O_WRONLY|os.O_APPEND, 0644)
			if err != nil {
				log.Fatalf("failed opening file: %s", err)
			}
			defer f.Close()
		}
		
		w := csv.NewWriter(f)
		var record []string
		record = append(record, "Museum ID : "+d.MuseumID+"\n")
		record = append(record, "Kode Pengelolaan : "+d.KodePengelolaan+"\n")
		record = append(record, "Nama : "+d.Nama+"\n")
		record = append(record, "SDM : "+d.Sdm+"\n")
		record = append(record, "Alamat : "+d.AlamatJalan+"\n")
		record = append(record, "Desa/Kelurahan : "+d.DesaKelurahan+"\n")
		record = append(record, "Kecamatan : "+d.Kecamatan+"\n")
		record = append(record, "Kabupaten/Kota : "+d.KabupatenKota+"\n")
		record = append(record, "Propinsi : "+d.Propinsi+"\n")
		record = append(record, "Lintang : "+d.Lintang+"\n")
		record = append(record, "Bujur : "+d.Bujur+"\n")
		record = append(record, "Koleksi : "+d.Koleksi+"\n")
		record = append(record, "Sumber Dana : "+d.SumberDana+"\n")
		record = append(record, "Pengelola : "+d.Pengelola+"\n")
		record = append(record, "Tipe : "+d.Tipe+"\n")
		record = append(record, "Standar : "+d.Standar+"\n")
		record = append(record, "Tahun Berdiri : "+d.TahunBerdiri+"\n")
		record = append(record, "Bangunan : "+d.Bangunan+"\n")
		record = append(record, "Luas Tanah : "+d.LuasTanah+"\n")
		record = append(record, "-----------------------------------------------------------------------")
		w.Write(record)
		w.Flush()

		kabkota = d.KabupatenKota
	}
	
}
