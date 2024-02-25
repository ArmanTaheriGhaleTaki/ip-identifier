package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	// "time"

	"github.com/gofiber/fiber/v2"
	_ "github.com/jackc/pgx/v4/stdlib"
	// "github.com/Masterminds/structable"
)

type cars struct {
	brand string
	model string
	year  int
}
type IPQuery struct {
	Status      string `json:"status"`
	Country     string `json:"country"`
	CountryCode string `json:"countryCode"`
	Region      string `json:"region"`
	RegionName  string `json:"regionName"`
	City        string `json:"city"`
	Zip         string `json:"zip"`
	// Lat         int64    `json:"lat"`
	// Lon         int    `json:"lon"`
	Timezone string      `json:"timezone"`
	Isp      string      `json:"isp"`
	Org      interface{} `json:"org"`
	As       string      `json:"as"`
	Query    string      `json:"query"`
}

func GetUser(ctx context.Context, conn *sql.DB, brand string) (cars, error) {
	const query = `SELECT brand , model , year FROM cars  where brand = $1 ;`
	u := cars{brand: brand}
	err := conn.QueryRowContext(ctx, query, brand).Scan(&u.brand, &u.model, &u.year)
	return u, err

}

var Database_name string 

func main() {
Database_name = "IPQuery"
	conn, err := sql.Open("pgx", "postgres://a:123@localhost:5432/ali")
	_,err = conn.Exec("CREATE DATABASE "+Database_name)
//    if err != nil {
    //    if(err =fmt.Errorf("ERROR: database "ipquery" already exists (SQLSTATE 42P04)")){
		// fmt.Println("it's ok")
	//    }
//    }
//    _,err = conn.Exec("USE "+Database_name)
//    if err != nil {
//        panic(err)
//    }

   _,err = conn.Exec("CREATE TABLE example ( id integer, data varchar(32) )")
   if err != nil {
       panic(err)
   }
  


	// if err != nil {
	// 	fmt.Println(err)
	// }
	// err = conn.Ping()
	// if err != nil {
	// 	fmt.Println("mmd2")
	// }
	// defer conn.Close()
	// ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	// if err := conn.PingContext(ctx); err != nil {
	// 	fmt.Println("46")
	// }
	// car, _ := GetUser(ctx, conn, "Ford")
	// fmt.Println(car)
	// cancel()

	app := fiber.New()
	app.Get("/getip/:ip", func(c *fiber.Ctx) error {
		url := fmt.Sprintf("http://ip-api.com/json/%s", c.Params("ip"))
		resp, err := http.Get(url)
		if err != nil {
			return c.Status(http.StatusInternalServerError).SendString(err.Error())
		}
		defer resp.Body.Close()
		// Read the response body

		bodyBytes, err := io.ReadAll(resp.Body)
		// bodyString := string(bodyBytes)
		if err != nil {
			return c.Status(http.StatusInternalServerError).SendString(err.Error())
		}
		// fmt.Println(string(bodyString))
		//#################### for debug ##############################
		// fmt.Println("Raw JSON response:", string(bodyBytes))
		//#################### for debug ##############################

		// Create an instance of the struct to hold the response
		var result IPQuery
		// Unmarshal JSON data into the struct
		err = json.Unmarshal(bodyBytes, &result)
		if err != nil {
			return c.Status(http.StatusInternalServerError).SendString(err.Error())
		}
		fmt.Println(result.Query)
		// Send the response as JSON

		return c.JSON(result)
		// Send the response from the external API
		// return c.SendString(string(bodyString))
	})

	app.Listen(":1111")

}
