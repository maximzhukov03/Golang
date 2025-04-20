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
	return json.MarshalIndent(r, "", "")
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

// WeatherProvider интерфейс для получения погоды
type WeatherGeter interface {
	GetWeather(city string) (Weather, error)
}

// MockWeatherProvider - мок-реализация WeatherProvider
type MockWeather struct {
	data map[string]Weather
}

// NewMockWeatherProvider создает новый мок-объект с данными для 3 городов
func NewMockWeather() *MockWeather {
	return &MockWeather{
		data: map[string]Weather{
			"Уфа": {
				Temperature: "+24 °C",
				Wind:        "19 km/h",
				Description: "Sunny",
				Forecast: []Forecast{
					{Day: "1", Temperature: "+24 °C", Wind: "17 km/h"},
					{Day: "2", Temperature: "+23 °C", Wind: "15 km/h"},
					{Day: "3", Temperature: "+27 °C", Wind: "17 km/h"},
				},
			},
			"Стерлитамак": {
				Temperature: "+18 °C",
				Wind:        "12 km/h",
				Description: "Partly cloudy",
				Forecast: []Forecast{
					{Day: "1", Temperature: "+18 °C", Wind: "12 km/h"},
					{Day: "2", Temperature: "+16 °C", Wind: "10 km/h"},
					{Day: "3", Temperature: "+20 °C", Wind: "15 km/h"},
				},
			},
			"СПб": {
				Temperature: "+15 °C",
				Wind:        "22 km/h",
				Description: "Rainy",
				Forecast: []Forecast{
					{Day: "1", Temperature: "+15 °C", Wind: "22 km/h"},
					{Day: "2", Temperature: "+14 °C", Wind: "20 km/h"},
					{Day: "3", Temperature: "+16 °C", Wind: "18 km/h"},
				},
			},
		},
	}
}

// GetWeather реализует интерфейс WeatherProvider для мок-объекта
func (m *MockWeather) GetWeather(city string) (Weather, error) {
	data, ok := m.data[city]
	if !ok {
		return Weather{}, fmt.Errorf("city %s not found", city)
	}
	return data, nil
}

func main(){
	// Создаем мок-провайдер
	mock := NewMockWeather()

	// Получаем данные для Минска
	weather, err := mock.GetWeather("Уфа")
	if err != nil {
		fmt.Println(err)
	}
  weath, err := weather.Marshal()
  if err != nil{
    fmt.Println(err)
  }
	// Выводим результат
	fmt.Println(string(weath))

	// Демонстрация работы с другими городами
	cities := []string{"Уфа", "Стерлитамак", "СПб",}
	for _, city := range cities {
		weather, err := mock.GetWeather(city)
		if err != nil {
			fmt.Println(err)
    }
    weath, err := weather.Marshal()
    if err != nil{
      fmt.Println(err)
    }
    fmt.Printf("%s:\n%s\n", city, string(weath))
	}
}