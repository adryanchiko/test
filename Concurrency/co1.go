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

type Provinsi struct {
	KodeWilayah    string `json:"kode_wilayah"`
	Nama           string `json:"nama"`
	MstKodeWilayah string `json:"mst_kode_wilayah"`
}

type Kabukota struct {
	KodeWilayah    string `json:"kode_wilayah"`
	Nama           string `json:"nama"`
	MstKodeWilayah string `json:"mst_kode_wilayah"`
}

type Museum struct {
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

func findProv() {
	url := "http://jendela.data.kemdikbud.go.id/api/index.php/CWilayah/wilayahGET"
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
		Data []Provinsi `json:"data"`
	}{}
    
	jerr := json.Unmarshal(body, &resp)
	if jerr != nil {
		fmt.Println("error:", jerr)
	}

	for _,d := range resp.Data {
		findKabukota(d.KodeWilayah)
	}
}

func findKabukota(kodewil string) {
	url := "http://jendela.data.kemdikbud.go.id/api/index.php/CWilayah/wilayahGET?mst_kode_wilayah="+kodewil
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
		Data []Kabukota `json:"data"`
	}{}
    
	jerr := json.Unmarshal(body, &resp)
	if jerr != nil {
		fmt.Println("error:", jerr)
	}

	for _,d := range resp.Data {
		findMuseum(d.KodeWilayah)
	}
}

func findMuseum(kodewil string) {
	url := "http://jendela.data.kemdikbud.go.id/api/index.php/CcariMuseum/searchGET?kode_kab_kota="+kodewil
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
		Data []Museum `json:"data"`
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
			w := csv.NewWriter(f)
			var record []string
			record = append(record, "Museum ID", "Kode Pengelolaan", "Nama", "SDM", "Alamat", "Desa/Kelurahan",
			"Kecamatan", "Kabupaten/Kota", "Propinsi", "Lintang", "Bujur", "Koleksi", "Sumber Dana", "Pengelola",
			"Tipe", "Standar", "Tahun Berdiri", "Bangunan", "Luas Tanah")
			w.Write(record)
			w.Flush()
		} else {
			f, err = os.OpenFile("./"+d.KabupatenKota+".csv", os.O_WRONLY|os.O_APPEND, 0644)
			if err != nil {
				log.Fatalf("failed opening file: %s", err)
			}
			defer f.Close()
		}
		
		w := csv.NewWriter(f)
		var record []string
		record = append(record, d.MuseumID)
		record = append(record, d.KodePengelolaan)
		record = append(record, d.Nama)
		record = append(record, d.Sdm)
		record = append(record, d.AlamatJalan)
		record = append(record, d.DesaKelurahan)
		record = append(record, d.Kecamatan)
		record = append(record, d.KabupatenKota)
		record = append(record, d.Propinsi)
		record = append(record, d.Lintang)
		record = append(record, d.Bujur)
		record = append(record, d.Koleksi)
		record = append(record, d.SumberDana)
		record = append(record, d.Pengelola)
		record = append(record, d.Tipe)
		record = append(record, d.Standar)
		record = append(record, d.TahunBerdiri)
		record = append(record, d.Bangunan)
		record = append(record, d.LuasTanah)
		w.Write(record)
		w.Flush()

		kabkota = d.KabupatenKota
	}
}

func main() {

	//Prov -> Kab/Kota -> Museum
	findProv()
	
}
