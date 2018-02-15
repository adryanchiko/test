package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
    "time"
    "bytes"
)

// type Museum map[string] interface{}
// type Museum struct {
//     Data map[string] interface{}
//   }
// type Museum struct {
// 	Data []struct {
// 		MuseumID          string      `json:"museum_id"`
// 		KodePengelolaan   string      `json:"kode_pengelolaan"`
// 		Nama              string      `json:"nama"`
// 		Sdm               string      `json:"sdm"`
// 		AlamatJalan       string      `json:"alamat_jalan"`
// 		DesaKelurahan     string      `json:"desa_kelurahan"`
// 		Kecamatan         string      `json:"kecamatan"`
// 		KabupatenKota     string      `json:"kabupaten_kota"`
// 		Propinsi          string      `json:"propinsi"`
// 		Lintang           string      `json:"lintang"`
// 		Bujur             string      `json:"bujur"`
// 		Koleksi           string      `json:"koleksi"`
// 		SumberDana        string      `json:"sumber_dana"`
// 		Pengelola         string      `json:"pengelola"`
// 		Tipe              string      `json:"tipe"`
// 		Standar           string      `json:"standar"`
// 		TahunBerdiri      string      `json:"tahun_berdiri"`
// 		Bangunan          string      `json:"bangunan"`
// 		LuasTanah         string      `json:"luas_tanah"`
// 		StatusKepemilikan interface{} `json:"status_kepemilikan"`
// 	} `json:"data"`
// }

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

    var Museum map[string] interface{}
    
	jerr := json.Unmarshal(body, &Museum)
	if jerr != nil {
		fmt.Println("error:", err)
	}

	fmt.Println(Museum)
}