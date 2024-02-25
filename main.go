package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2"
	_ "github.com/jackc/pgx/v4/stdlib"
)

type cars struct {
	brand string
	model string
	year  int
}
type IPQuery struct {
	// Status string `json:"status"`

	Country string `json:"country"`
	// CountryCode string `json:"countryCode"`
	// Region      string `json:"region"`
	// RegionName  string `json:"regionName"`
	// City        string `json:"city"`
	// Zip         string `json:"zip"`
	// Lat         int64    `json:"lat"`
	// Lon         int    `json:"lon"`
	// Timezone string      `json:"timezone"`
	// Isp      string      `json:"isp"`
	// Org      interface{} `json:"org"`
	// As       string      `json:"as"`
	Query string `json:"query"`
}

func GetInfo(ctx context.Context, conn *sql.DB, IP string) (IPQuery, error) {
	url := fmt.Sprintf("SELECT coutnry FROM IPInfoo where ip = '%s' ;", IP)
	// const query = `SELECT coutnry FROM IPInfoo where ip = '$1' ;`
	output := IPQuery{Query: IP}
	err := conn.QueryRow( url).Scan(&output.Country)
	if err == nil {
		log.Printf("ip: %-15s coutnry: %-15s is query successfully",output.Query , output.Country)
	}
	return output, err

}

// func GetUser(ctx context.Context, conn *sql.DB, brand string) (cars, error) {
// 	const query = `SELECT brand , model , year FROM cars  where brand = $1 ;`
// 	u := cars{brand: brand}
// 	err := conn.QueryRowContext(ctx, query, brand).Scan(&u.brand, &u.model, &u.year)
// 	return u, err

// }
var Database_name string

func main() {
	Database_name = "ipinfoo"
	conn, err := sql.Open("pgx", "postgres://a:123@localhost:5432/ali")
	// _, err = conn.Exec("CREATE DATABASE " + Database_name)
	// _, err = conn.Exec("CREATE TABLE IPInfo ( IP varchar(35), country varchar(32) )")
	
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
	defer conn.Close()
	ctx, _ := context.WithTimeout(context.Background(), 2*time.Second)
	// if err := conn.PingContext(ctx); err != nil {
	// 	fmt.Println("46")
	// }
	// car, _ := GetUser(ctx, conn, "Ford")
	// fmt.Println(car)
	// cancel()

	app := fiber.New()
	app.Get("/getip/:ip", func(c *fiber.Ctx) error {
		output, err := GetInfo(ctx, conn, c.Params("ip"))
		if err != nil {

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
			var result IPQuery
			// Unmarshal JSON data into the struct
			err = json.Unmarshal(bodyBytes, &result)
			if err != nil {
				return c.Status(http.StatusInternalServerError).SendString(err.Error())
			}
			var insert string = fmt.Sprintf("INSERT INTO ipinfoo (ip , coutnry) values ('%s','%s');", result.Query, result.Country)
			_, _ = conn.Exec(insert)
			output, _ = GetInfo(ctx, conn, c.Params("ip"))
		} 
		return c.JSON(output)

	})

	app.Listen(":1111")

}
