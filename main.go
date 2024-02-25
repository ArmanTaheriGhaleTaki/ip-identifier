package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2"
	_ "github.com/jackc/pgx/v4/stdlib"
)

type IPQuery struct {
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

func IPCheck(ip string) bool {
	if net.ParseIP(ip) == nil {
		return false
	} else {
		return true
	}

}

func GetInfo(ctx context.Context, conn *sql.DB, IP string) (IPQuery, error) {
	url := fmt.Sprintf("SELECT coutnry FROM %s where ip = '%s' ;", Database_name, IP)
	output := IPQuery{Query: IP}
	err := conn.QueryRow(url).Scan(&output.Country)
	if err == nil {
		log.Printf("ip: %-15s coutnry: %-15s is query successfully", output.Query, output.Country)
	}
	return output, err

}

var Database_name string

func main() {
	Database_name = "ipinfoo"
	conn, _ := sql.Open("pgx", "postgres://a:123@localhost:5432/ali")
	_, _ = conn.Exec("CREATE DATABASE " + Database_name)
	_, _ = conn.Exec("CREATE TABLE IPInfo ( IP varchar(35), country varchar(32) )")
	defer conn.Close()
	ctx, _ := context.WithTimeout(context.Background(), 2*time.Second)

	app := fiber.New()
	app.Get("/getip/:ip", func(c *fiber.Ctx) error {
		if IPCheck(c.Params("ip")) {
			log.Printf("input is valid: %-15s ", c.Params("ip"))

			output, err := GetInfo(ctx, conn, c.Params("ip"))
			if err != nil {

				url := fmt.Sprintf("http://ip-api.com/json/%s", c.Params("ip"))
				resp, err := http.Get(url)
				if err != nil {
					return c.Status(http.StatusInternalServerError).SendString(err.Error())
				}
				defer resp.Body.Close()
				bodyBytes, err := io.ReadAll(resp.Body)
				if err != nil {
					return c.Status(http.StatusInternalServerError).SendString(err.Error())
				}
				var result IPQuery
				err = json.Unmarshal(bodyBytes, &result)
				if err != nil {
					return c.Status(http.StatusInternalServerError).SendString(err.Error())
				}
				var insert string = fmt.Sprintf("INSERT INTO ipinfoo (ip , coutnry) values ('%s','%s');", result.Query, result.Country)
				_, _ = conn.Exec(insert)
				output, _ = GetInfo(ctx, conn, c.Params("ip"))
			}
			return c.JSON(output)
		} else {
			log.Printf("invalid input: %-15s ", c.Params("ip"))
			return c.JSON("invalid parameter")
		}

	})

	app.Listen(":1111")

}
