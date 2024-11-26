package writeJson

import(
	"fmt"
	"os"
	"io/ioutil"
	"encoding/json"
	"strconv"
	"strings"
	"net/http"

	_ "database/sql"
	_ "github.com/go-sql-driver/mysql"

	"qqch/weatherCode"
	"qqch/conn"
)

func checkE(e error) {
	if e != nil {
		panic(e.Error())
	}
}

type LtFrom struct {
	TheDate string
	TheNums string
}

type Temp2m struct {
	T2m []float64 `json:"temperature_2m"`
	Rh2m []float64 `json:"relative_humidity_2m"`
	Dp2m []float64 `json:"dew_point_2m"`
	Prec []float64 `json:"precipitation"`
	Wc []int `json:"weather_code"`
	Pmsl []float64 `json:"pressure_msl"`
	Sp []float64 `json:"surface_pressure"`
	Cc []int `json:"cloud_cover"`
	Eva []float64 `json:"et0_fao_evapotranspiration"`
}

type Temps struct {
	Ts Temp2m `json:"hourly"`
}

type LtTo struct {
	TheDate string
	Month int
	Year int
	TheNums []int
	T20temp float64
	T21temp float64
	T22temp float64
	T23temp float64
	RelativeHumidity20 float64
	RelativeHumidity21 float64
	RelativeHumidity22 float64
	RelativeHumidity23 float64
	DewPoint20 float64
	DewPoint21 float64
	DewPoint22 float64
	DewPoint23 float64
	Precipitation20 float64
	Precipitation21 float64
	Precipitation22 float64
	Precipitation23 float64
	Weathercode20 int
	WeatherDesc20 string
	Weathercode21 int
	WeatherDesc21 string
	Weathercode22 int
	WeatherDesc22 string
	Weathercode23 int
	WeatherDesc23 string
	PressureMSL20 float64
	PressureMSL21 float64
	PressureMSL22 float64
	PressureMSL23 float64
	SurfacePressure20 float64
	SurfacePressure21 float64
	SurfacePressure22 float64
	SurfacePressure23 float64
	CloudCover20 int
	CloudCover21 int
	CloudCover22 int
	CloudCover23 int
	Evapotranspiration20 float64
	Evapotranspiration21 float64
	Evapotranspiration22 float64
	Evapotranspiration23 float64
}

func WriteJson() {
	jfile, err := os.Open("rows.json")
	checkE(err)

	defer jfile.Close()

	bval, err := ioutil.ReadAll(jfile)

	var data1 []LtFrom
	var ltArr []LtTo

	_ = json.Unmarshal(bval, &data1)

	var theNums []int

	for i := 0; i < len(data1); i++ {
		snums := strings.Fields(data1[i].TheNums)

		for j := 0; j < len(snums); j++ {
			newNum, _ := strconv.Atoi(snums[j])
			theNums = append(theNums, newNum)
		}

		getDate := strings.Split(data1[i].TheDate, "-")

		yearStr := getDate[0]
		year, _ := strconv.Atoi(yearStr)
		monthStr := getDate[1]
		month, _ := strconv.Atoi(monthStr)

		lottArr = append(lottArr, getDeets(data1[i].TheDate, month, year, theNums))

		Into(getDeets(data1[i].TheDate, month, year, theNums))

		theNums = nil
		
		fmt.Println(data1[i].TheDate)
	}

	newLt, _ := json.MarshalIndent(ltArr, "", " ")
	_ = ioutil.WriteFile("nWRows.json", newLt, 0777)
}

func getDeets(theDate string, theMonth int, theYear int, theNums []int) LtTo {

	resp, err := http.Get("https://archive-api.open-meteo.com/v1/archive?latitude=30.4382&longitude=-84.2806&start_date="+theDate+"&end_date="+theDate+"&hourly=temperature_2m,relative_humidity_2m,dew_point_2m,apparent_temperature,precipitation,weather_code,pressure_msl,surface_pressure,cloud_cover,et0_fao_evapotranspiration&temperature_unit=fahrenheit&timezone=America%2FNew_York")
	checkE(err)

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}

	var w Temps

	json.Unmarshal(body, &w)

	data2 := LtTo{}
	data2.TheDate = theDate
	data2.Month = theMonth
	data2.Year = theYear
	data2.TheNums = theNums
	data2.T20temp = w.Ts.T2m[20]
	data2.T21temp = w.Ts.T2m[21]
	data2.T22temp = w.Ts.T2m[22]
	data2.T23temp = w.Ts.T2m[23]
	data2.RelativeHumidity20 = w.Ts.Rh2m[20]
	data2.RelativeHumidity21 = w.Ts.Rh2m[21]
	data2.RelativeHumidity22 = w.Ts.Rh2m[22]
	data2.RelativeHumidity23 = w.Ts.Rh2m[23]
	data2.DewPoint20 = w.Ts.Dp2m[20]
	data2.DewPoint21 = w.Ts.Dp2m[21]
	data2.DewPoint22 = w.Ts.Dp2m[22]
	data2.DewPoint23 = w.Ts.Dp2m[23]
	data2.Precipitation20 = w.Ts.Prec[20]
	data2.Precipitation21 = w.Ts.Prec[21]
	data2.Precipitation22 = w.Ts.Prec[22]
	data2.Precipitation23 = w.Ts.Prec[23]
	data2.Weathercode20 = w.Ts.Wc[20]
	data2.WeatherDesc20 = getWeaDesc(w.Ts.Wc[20])
	data2.Weathercode21 = w.Ts.Wc[21]
	data2.WeatherDesc21 = getWeaDesc(w.Ts.Wc[21])
	data2.Weathercode22 = w.Ts.Wc[22]
	data2.WeatherDesc22 = getWeaDesc(w.Ts.Wc[22])
	data2.Weathercode23 = w.Ts.Wc[23]
	data2.WeatherDesc23 = getWeaDesc(w.Ts.Wc[23])
	data2.PressureMSL20 = w.Ts.Pmsl[20]
	data2.PressureMSL21 = w.Ts.Pmsl[21]
	data2.PressureMSL22 = w.Ts.Pmsl[22]
	data2.PressureMSL23 = w.Ts.Pmsl[23]
	data2.SurfacePressure20 = w.Ts.Sp[20]
	data2.SurfacePressure21 = w.Ts.Sp[21]
	data2.SurfacePressure22 = w.Ts.Sp[22]
	data2.SurfacePressure23 = w.Ts.Sp[23]
	data2.CloudCover20 = w.Ts.Cc[20]
	data2.CloudCover21 = w.Ts.Cc[21]
	data2.CloudCover22 = w.Ts.Cc[22]
	data2.CloudCover23 = w.Ts.Cc[23]
	data2.Evapotranspiration20 = w.Ts.Eva[20]
	data2.Evapotranspiration21 = w.Ts.Eva[21]
	data2.Evapotranspiration22 = w.Ts.Eva[22]
	data2.Evapotranspiration23 = w.Ts.Eva[23]

	return data2
}

func getWeaDesc(n int) string {
	return weatherCode.GetCode(n)
}

func Into(lt LtTo) {
	db := conn.Conn()

	defer db.Close()

	stmt, err := db.Prepare("insert into lstats(the_date, the_month, the_year, b1, b2, b3, b4, b5, b6, temp20, temp21, temp22, temp23, RelativeHumidity20, RelativeHumidity21, RelativeHumidity22, RelativeHumidity23, DewPoint20, DewPoint21, DewPoint22, DewPoint23, Precipitation20, Precipitation21, Precipitation22, Precipitation23, Weathercode20, WeatherDesc20, Weathercode21, WeatherDesc21, Weathercode22, WeatherDesc22, Weathercode23, WeatherDesc23, PressureMSL20, PressureMSL21, PressureMSL22, PressureMSL23, SurfacePressure20, SurfacePressure21, SurfacePressure22, SurfacePressure23, CloudCover20, CloudCover21, CloudCover22, CloudCover23, Evapotranspiration20, Evapotranspiration21, Evapotranspiration22, Evapotranspiration23) values(?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)")
	if err != nil {
		panic(err.Error())
	}

	_, err = stmt.Exec(lt.TheDate, lt.Month, lt.Year, lt.TheNums[0], lt.TheNums[1], lt.TheNums[2], lt.TheNums[3], lt.TheNums[4], lt.TheNums[5], lt.T20temp, lt.T21temp, lt.T22temp, lt.T23temp, lt.RelativeHumidity20, lt.RelativeHumidity21, lt.RelativeHumidity22, lt.RelativeHumidity23, lt.DewPoint20, lt.DewPoint21, lt.DewPoint22, lt.DewPoint23, lt.Precipitation20, lt.Precipitation21, lt.Precipitation22, lt.Precipitation23, lt.Weathercode20, lt.WeatherDesc20, lt.Weathercode21, lt.WeatherDesc21, lt.Weathercode22, lt.WeatherDesc22, lt.Weathercode23, lt.WeatherDesc23, lt.PressureMSL20, lt.PressureMSL21, lt.PressureMSL22, lt.PressureMSL23, lt.SurfacePressure20, lt.SurfacePressure21, lt.SurfacePressure22, lt.SurfacePressure23, lt.CloudCover20, lt.CloudCover21, lt.CloudCover22, lt.CloudCover23, lt.Evapotranspiration20, lt.Evapotranspiration21, lt.Evapotranspiration22, lt.Evapotranspiration23)
	if err != nil {
		panic(err.Error())
	}
}