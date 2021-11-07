package Soal1

import (
	"CadItTest/models"
	"sort"
	"strconv"

	"github.com/montanaflynn/stats"
)

func (b UsecaseSoal1Struct) Soal2Func2(byDay string, roomArea string, traceHeader map[string]string) (models.DataSensorRespone, error) {
	//Get data data
	var response = models.DataSensorRespone{}
	res, err := b.testRepo.GetDataSensor(traceHeader)
	if err != nil {
		b.log.Error(err, "usecase/Soal1: cnnot execute get data from repo", traceHeader)
		return models.DataSensorRespone{}, err
	}
	//Declare Variable
	TempData := 0.0
	HumData := 0.0
	modusTemp := 0.0
	modusHum := 0.0
	medianTemp := 0.0
	medianHumid := 0.0
	arrayTemp := []float64{}
	arrayHumid := []float64{}
	lengthArr := len(res)
	totalData := 0
	day, _ := strconv.ParseInt(byDay, 10, 64)
	for i := 0; i < lengthArr; i++ {
		if res[i].Timestamp == day {
			if res[i].RoomArea == roomArea {
				arrayTemp = append(arrayTemp, res[i].Temperature)
				arrayHumid = append(arrayHumid, res[i].Humidity)
				TempData = TempData + res[i].Temperature
				HumData = HumData + res[i].Humidity
				totalData = totalData + 1
			}
		}
	}
	sort.Float64s(arrayTemp)
	sort.Float64s(arrayHumid)

	//Modus Data
	modusTemp, _ = stats.Mode(arrayTemp)
	modusHum, _ = stats.Mode(arrayHumid)
	//Median Data
	if totalData%2 == 0 {
		medianTemp = (arrayTemp[totalData/2] + arrayTemp[(totalData/2)+1]) / 2
		medianHumid = (arrayHumid[totalData/2] + arrayHumid[(totalData/2)+1]) / 2
	} else {
		medianTemp = arrayTemp[(totalData+1)/2]
		medianHumid = arrayHumid[totalData+1/2]
	}
	//Mean data
	meanTemp := TempData / float64(totalData)
	meanHum := HumData / float64(totalData)
	//parsing data
	response.Temperature.TemperatureMean = meanTemp
	response.Temperature.TemperatureMedian = medianTemp
	response.Temperature.TemperatureModus = modusTemp
	response.Humidity.HumidityMean = meanHum
	response.Humidity.HumidityMedian = medianHumid
	response.Humidity.HumidityModus = modusHum
	return response, nil
}
