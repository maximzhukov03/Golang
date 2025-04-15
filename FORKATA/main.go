package main


import (
  "encoding/json"
  "fmt"
)

func UnmarshalWeather(data []byte) (Weather, error) {
	var r Weather
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *Weather) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type Weather struct {
	Temperature string     `json:"temperature"`
	Wind        string     `json:"wind"`
	Description string     `json:"description"`
	Forecast    []Forecast `json:"forecast"`
}

type Forecast struct {
	Day         string `json:"day"`
	Temperature string `json:"temperature"`
	Wind        string `json:"wind"`
}

type WeatherData interface{
  Get(city string) (Weather, error)
}

type DataWeather struct{
  data map[string]Weather
}

func NewMock() *DataWeather{
  return &DataWeather{    
    data: map[string]Weather{
      "Уфа":{
        Temperature:"20",
	      Wind:"23",
	      Description:"23",
	      Forecast:[]Forecast{
          {	
            Day: "1",
	          Temperature: "23",
	          Wind:"25",
          },
        },
      },
    },
  }
}

func (db *DataWeather) Get(c string) (Weather, error){
  data, ok := db.data[c]
  if ok {
    return data, nil
  }
  
  return data, fmt.Errorf("Ошибка получения")
}

func main(){
  db := DataWeather{
    data: map[string]Weather{
      "Уфа":{
        Temperature:"20",
	      Wind:"23",
	      Description:"23",
	      Forecast:[]Forecast{
          {	
            Day: "1",
	          Temperature: "23",
	          Wind:"25",
          },
        },
      },
    },
  }
  
  data, _ := db.Get("Уфа")
  fmt.Println(data)
}