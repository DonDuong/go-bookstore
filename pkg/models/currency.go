package domain

import (
	"currency-api/config"
	"database/sql"
	"fmt"
	"net/http"
	"strings"

	"github.com/labstack/echo"
)

var db *sql.DB
var err error

type CurrencyInter interface {
	GetCurrencys(c echo.Context) error
	GetCurrency(c echo.Context) error
	GetCurrencyNotEcho() echo.HandlerFunc
}
type Currency struct {
	Currency_id *string
	Currency_pg *string
	Currency_jg *string
	Currency_nr *int
}

type CurrencyResponse struct {
	Message  string     `json:"message"`
	Currency []Currency `json:"currency"`
}

func init() {
	config.Connect()
	db = config.GetDB()
}
func (currency Currency) GetCurrencys(c echo.Context) error {
	var currencys []Currency
	rows, err := db.Query("SELECT * FROM currency_pg")
	if err != nil {
		fmt.Println("Select nil")
	}
	defer rows.Close()

	for rows.Next() {

		if err := rows.Scan(&currency.Currency_id, &currency.Currency_pg, &currency.Currency_jg, &currency.Currency_nr); err != nil {
			fmt.Println("Scan faild")
		}
		currencys = append(currencys, currency)
	}
	response := CurrencyResponse{
		Message:  "Success",
		Currency: currencys,
	}
	return c.JSON(http.StatusOK, response)
}

func (currency Currency) GetCurrency(c echo.Context) error {
	currency_id := c.Param("currency_id")
	fmt.Println()
	//check id isExist
	existsQuery := `SELECT EXISTS(SELECT 1 FROM currency_pg WHERE currency_id = $1)`
	var exists bool
	err := db.QueryRow(existsQuery, strings.ToUpper(currency_id)).Scan(&exists)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	if !exists {
		response := CurrencyResponse{
			Message:  "ID not found",
			Currency: nil,
		}
		return c.JSON(404, response)
	}
	//Select id
	row := db.QueryRow(`SELECT * FROM currency_pg WHERE currency_id = $1`, strings.ToUpper(currency_id))

	if err := row.Scan(&currency.Currency_id, &currency.Currency_pg, &currency.Currency_jg, &currency.Currency_nr); err != nil {
		fmt.Println(err.Error())
	}
	response := CurrencyResponse{
		Message:  "Success",
		Currency: []Currency{currency},
	}
	return c.JSON(http.StatusOK, response)
}

func (currency Currency) GetCurrencyNotEcho() echo.HandlerFunc {
	return func(c echo.Context) error {
		currency_id := c.Param("currency_id")
		fmt.Println()
		//check id isExist
		existsQuery := `SELECT EXISTS(SELECT 1 FROM currency_pg WHERE currency_id = $1)`
		var exists bool
		err := db.QueryRow(existsQuery, strings.ToUpper(currency_id)).Scan(&exists)
		if err != nil {
			fmt.Println(err.Error())
			return err
		}
		if !exists {
			response := CurrencyResponse{
				Message:  "ID not found",
				Currency: nil,
			}
			return c.JSON(404, response)
		}
		//Select id
		row := db.QueryRow(`SELECT * FROM currency_pg WHERE currency_id = $1`, strings.ToUpper(currency_id))

		if err := row.Scan(&currency.Currency_id, &currency.Currency_pg, &currency.Currency_jg, &currency.Currency_nr); err != nil {
			fmt.Println(err.Error())
		}
		response := CurrencyResponse{
			Message:  "Success",
			Currency: []Currency{currency},
		}
		return c.JSON(http.StatusOK, response)
	}
}
