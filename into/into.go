// package into

// import(
// 	_ "database/sql"

// 	"lotto_stats/conn"
// 	"lotto_stats/writeJson"

// 	_ "github.com/go-sql-driver/mysql"
// )

// func Into(lotto Lotto) {
// 	db := conn.Conn()

// 	defer db.Close()

// 	stmt, err := db.Prepare("insert into lstats(draw_date, p, o, w, e, r, ball, temp20, temp21, temp22, temp23, RelativeHumidity20, RelativeHumidity21, RelativeHumidity22, RelativeHumidity23, DewPoint20, DewPoint21, DewPoint22, DewPoint23, Precipitation20, Precipitation21, Precipitation22, Precipitation23, Weathercode20, WeatherDesc20, Weathercode21, WeatherDesc21, Weathercode22, WeatherDesc22, Weathercode23, WeatherDesc23, PressureMSL20, PressureMSL21, PressureMSL22, PressureMSL23, SurfacePressure20, SurfacePressure21, SurfacePressure22, SurfacePressure23, CloudCover20, CloudCover21, CloudCover22, CloudCover23, Evapotranspiration20, Evapotranspiration21, Evapotranspiration22, Evapotranspiration23) values(?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)")
// 	if err != nil {
// 		panic(err.Error())
// 	}

// 	_, err = stmt.Exec(lotto.DrawDate, lotto.DrawNums0, lotto.DrawNums1, lotto.DrawNums2, lotto.DrawNums3, lotto.DrawNums4, lotto.DrawNums5, lotto.T20temp, lotto.T21temp, lotto.T22temp, lotto.T23temp, lotto.RelativeHumidity20, lotto.RelativeHumidity21, lotto.RelativeHumidity22, lotto.RelativeHumidity23, lotto.DewPoint20, lotto.DewPoint21, lotto.DewPoint22, lotto.DewPoint23, lotto.Precipitation20, lotto.Precipitation21, lotto.Precipitation22, lotto.Precipitation23, lotto.Weathercode20, lotto.WeatherDesc20, lotto.Weathercode21, lotto.WeatherDesc21, lotto.Weathercode22, lotto.WeatherDesc22, lotto.Weathercode23, lotto.WeatherDesc23, lotto.PressureMSL20, lotto.PressureMSL21, lotto.PressureMSL22, lotto.PressureMSL23, lotto.SurfacePressure20, lotto.SurfacePressure21, lotto.SurfacePressure22, lotto.SurfacePressure23, lotto.CloudCover20, lotto.CloudCover21, lotto.CloudCover22, lotto.CloudCover23, lotto.Evapotranspiration20, lotto.Evapotranspiration21, lotto.Evapotranspiration22, lotto.Evapotranspiration23)
// 	if err != nil {
// 		panic(err.Error())
// 	}
// }