package main

import (
    "net/http"
    "encoding/json"
    "strings"
    "log"
    )

func main() {
    http.HandleFunc("/", hello)
    http.HandleFunc("/weather/", weatherFunc)
    http.ListenAndServe(":8080", nil)
}

func hello(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("Hello From GO http server"))
}

func weatherFunc(w http.ResponseWriter, r *http.Request) {
    city := strings.SplitN(r.URL.Path, "/", 3)[2]
    data, err := query(city)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json; charset=utf-8")
    json.NewEncoder(w).Encode(data)
}

func query(city string) (weatherData, error) {
    resp, err := http.Get("http://api.openweathermap.org/data/2.5/weather?q=" + city + "&appid=2de143494c0b295cca9337e1e96b00e0")
    if err != nil {
        return weatherData{}, err
    }

    defer resp.Body.Close()
    var d weatherData

    if err := json.NewDecoder(resp.Body).Decode(&d); err != nil {
        return weatherData{}, err
    }
    log.Printf("City : %s, Temp : %.2f Kelvin", city, d.Main.Kelvin)
    return d, nil
}

type weatherData struct {
    Name string `json:"name"`
    Main struct {
        Kelvin float64 `json:"temp"`
    } `json:"main"`
}
