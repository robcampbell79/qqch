package pull

import(
	_ "database/sql"

	"qqch/conn"

	_ "github.com/go-sql-driver/mysql"
)

type Lt struct {
	TheDate string
	Month int
	Year int
	TheNums0 int
	TheNums1 int
	TheNums2 int
	TheNums3 int
	TheNums4 int
	TheNums5 int
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

func Pull() []Lt {
	//47
	// var a, y, aa, cc, ee string
	// var b, c, d, e, f, g, x, z, bb, dd, nn, oo, pp, qq int
	// var h, i, j, k, l, m, n, o, p, q, r, s, t, u, v, w, ff, gg, hh, ii, jj, kk, ll, mm, rr, ss, tt, uu float64
	
	db := conn.Conn()

	defer db.Close()

	rows, err := db.Query("SELECT the_date, the_month, the_year, b1, b2, b3, b4, b5, b6, temp20, temp21, temp22, temp23, RelativeHumidity20, RelativeHumidity21, RelativeHumidity22, RelativeHumidity23, DewPoint20, DewPoint21, DewPoint22, DewPoint23, Precipitation20, Precipitation21, Precipitation22, Precipitation23, Weathercode20, WeatherDesc20, Weathercode21, WeatherDesc21, Weathercode22, WeatherDesc22, Weathercode23, WeatherDesc23, PressureMSL20, PressureMSL21, PressureMSL22, PressureMSL23, SurfacePressure20, SurfacePressure21, SurfacePressure22, SurfacePressure23, CloudCover20, CloudCover21, CloudCover22, CloudCover23, Evapotranspiration20, Evapotranspiration21, Evapotranspiration22, Evapotranspiration23 FROM lstats")
	if err != nil {
		panic(err.Error())
	}

	result := []Lt{}

	for rows.Next() {
		lt := Lt{}
		err = rows.Scan(&lt.TheDate, &lt.Month, &lt.Year, &lt.TheNums0, &lt.TheNums1, &lt.TheNums2, &lt.TheNums3, &lt.TheNums4, &lt.TheNums5, &lt.T20temp, &lt.T21temp, &lt.T22temp, &lt.T23temp, &lt.RelativeHumidity20, &lt.RelativeHumidity21, &lt.RelativeHumidity22, &lt.RelativeHumidity23, &lt.DewPoint20, &lt.DewPoint21, &lt.DewPoint22, &lt.DewPoint23, &lt.Precipitation20, &lt.Precipitation21, &lt.Precipitation22, &lt.Precipitation23, &lt.Weathercode20, &lt.WeatherDesc20, &lt.Weathercode21, &lt.WeatherDesc21, &lt.Weathercode22, &lt.WeatherDesc22, &lt.Weathercode23, &lt.WeatherDesc23, &lt.PressureMSL20, &lt.PressureMSL21, &lt.PressureMSL22, &lt.PressureMSL23, &lt.SurfacePressure20, &lt.SurfacePressure21, &lt.SurfacePressure22, &lt.SurfacePressure23, &lt.CloudCover20, &lt.CloudCover21, &lt.CloudCover22, &lt.CloudCover23, &lt.Evapotranspiration20, &lt.Evapotranspiration21, &lt.Evapotranspiration22, &lt.Evapotranspiration23)
		if err != nil {
			panic(err.Error())
		}
		result = append(result, lt)
	}

	return result
}