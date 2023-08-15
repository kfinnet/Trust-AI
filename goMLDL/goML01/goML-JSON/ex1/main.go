package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// citiBikeURL citiBike 자전거 공유 정류장의 상황을 알려준다
const citiBikeURL = "https://gbfs.citibikenyc.com/gbfs/en/station_status.json"

// stationData는 citiBikeURL 로부터 변환된 JSON문서의 구문을 분석하는 데 사용된다.
type stationData struct {
	LastUpdated int     `json:"last_updated"`
	TTL         int     `json:"ttl"`
	Data        struct {
	Stations []station  `json:"stations"`
	}                   `json:"data"`
}

// station은 stationData 안의 각 station 문서의 구문을 분석하는 데 사용된다
type station struct {
	ID                string `json:"station_id"`
	NumBikesAvailable int    `json:"num_bikes_available"`
	NumBikesDisabled  int    `json:"num_bike_disabled"`
	NumDocksAvailable int    `json:"num_docks_available"`
	NumDocksDisabled  int    `json:"num_docks_disabled"`
	IsInstalled       int    `json:"is_installed"`
	IsRenting         int    `json:"is_renting"`
	IsReturning       int    `json:"is_returning"`
	LastReported      int    `json:"last_reported"`
	HasAvailableKeys  bool   `json:"eightd_has_available_keys"`
}

func main() {

	// URL 로부터 JSON 응답을 얻는다.
	response, err := http.Get(citiBikeURL) //REST API(GET/POST/DELETE/PUT-UPDATE) Get은 read로 반환된다. REST API
	if err != nil {
		log.Fatal(err)
	}

	// Defer closing the response body.
	defer response.Body.Close()

	// 응답의 Body를  []byte 읽는다.
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	// stationData 유형의 변수를 선언한다.
	var sd stationData

	// stationData 변수로 JSON 데이터를 읽는다.
	if err := json.Unmarshal(body, &sd); err != nil {
		log.Fatal(err)
		return
	}

	// 첫번째 정류장 정보를 출력한다. 이 코드를 실행하면 
	// URL로 부터 읽어온 데이터가 저장된 구조체를 확인할 수있다.
	fmt.Printf("%+v\n\n", sd.Data.Stations[10])
}
