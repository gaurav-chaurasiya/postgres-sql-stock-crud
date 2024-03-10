package middleware

import (
	"database/sql"

	_ "github.com/lib/pq"

	"encoding/json"
	"fmt"
	"go-postgres-sql-stock/models"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

type responce struct {
	ID      int64  `json:"id,omitempty"`
	Message string `json:"message,omitempty"`
}

func createConnection() *sql.DB {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Errpr loading .env file")

	}
	db, err := sql.Open("postgres", os.Getenv("POSTGRES_URL"))
	if err != nil {
		panic(err)
	}
	err = db.Ping()
	if err != nil {
		panic(err)

	}
	fmt.Println("Successfully connected to postgresr")
	return db
}
func Test(w http.ResponseWriter, r *http.Request) {
	insertID := createConnection()
	fmt.Printf("connnt %v", insertID)

}
func CreateStock(w http.ResponseWriter, r *http.Request) {

	var stock models.Stock
	err := json.NewDecoder(r.Body).Decode(&stock)
	if err != nil {
		log.Fatal("Unable to decode the request body. %v", err)
	}
	fmt.Printf("stock %v", stock.Price)
	insertID := insertStock(stock)
	res := responce{
		ID:      insertID,
		Message: "Stock create successfully",
	}
	json.NewEncoder(w).Encode(res)
}

func GetStock(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		log.Fatal("unable to convert the string")

	}
	fmt.Printf("this is id %v", id)
	stock, err := getStock(int64(id))
	if err != nil {
		log.Fatal("Unble to get stock %v", err)
	}
	json.NewEncoder(w).Encode(stock)
}

func GetAllStock(w http.ResponseWriter, r *http.Request) {
	stock, err := getAllStocks()
	if err != nil {
		log.Fatal("unble to get all stock %v", err)
	}
	json.NewEncoder(w).Encode(stock)

}
func UpdateStock(w http.ResponseWriter, r *http.Request) {
	parsers := mux.Vars(r)
	id, err := strconv.Atoi(parsers["id"])
	if err != nil {

		log.Fatalf("Unable to convert the string into int %v", err)

	}
	var stock models.Stock
	err = json.NewDecoder(r.Body).Decode(&stock)
	if err != nil {
		log.Fatalf("unable to decode  the request body %v", err)

	}
	updatedRows := updateStock(int64(id), stock)
	msg := fmt.Sprintf("stock is updated successfully . Total rows/records affected %v", updatedRows)

	res := responce{
		ID:      int64(id),
		Message: msg,
	}

	json.NewEncoder(w).Encode(res)
}
func DeleteStock(w http.ResponseWriter, r *http.Request) {
	parsers := mux.Vars(r)
	id, err := strconv.Atoi(parsers["id"])
	if err != nil {
		log.Fatalf("unable to convert string to int %v ", err)
	}

	deletedRows := deleteStock(int64(id))
	msg := fmt.Sprintf("Stock delete successfully total rows/records %v", deletedRows)
	res := responce{
		ID:      int64(id),
		Message: msg,
	}
	json.NewEncoder(w).Encode(res)
}

func insertStock(stock models.Stock) int64 {
	db := createConnection()
	defer db.Close()
	fmt.Print(stock)
	sqlStatement := `insert into stocks(name,price,company) values($1,$2,$3) RETURNING stockid`
	var id int64
	err := db.QueryRow(sqlStatement, stock.Name, stock.Price, stock.Company).Scan(&id)
	if err != nil {
		log.Fatalf("unable to execute the query %v", err)

	}
	fmt.Printf("Inserted a single record %v", id)
	return id
}
func getStock(id int64) (models.Stock, error) {
	db := createConnection()
	defer db.Close()
	var stock models.Stock
	sqlStatement := `SELECT * from stocks where stockid=$1`
	row := db.QueryRow(sqlStatement, id)
	// fmt.Println("Now rows were returned! %v",)

	err := row.Scan(&stock.StockID, &stock.Name, &stock.Price, &stock.Company)
	switch err {
	case sql.ErrNoRows:
		fmt.Println("Now rows were returned!")
		return stock, nil
	case nil:
		return stock, nil
	default:
		log.Fatalf("Unable to scan the ro. %v", err)

	}
	return stock, err
}
func getAllStocks() ([]models.Stock, error) {
	db := createConnection()
	defer db.Close()
	var stocks []models.Stock
	sqlStatement := `SELECT * FROM stocks`
	rows, err := db.Query(sqlStatement)
	if err != nil {
		log.Fatalf("Unable to execute Query %v", err)

	}
	defer rows.Close()

	for rows.Next() {
		var stock models.Stock

		err := rows.Scan(&stock.StockID, &stock.Name, &stock.Price, &stock.Company)
		if err != nil {
			log.Fatalf("unable to scan the row %v", err)

		}
		stocks = append(stocks, stock)

	}
	return stocks, err

}
func updateStock(id int64, stock models.Stock) int64 {
	db := createConnection()
	defer db.Close()
	sqlStatement := `update stocks SET name=$2,price=$3,company=$4 where stockid=$1`
	res, err := db.Exec(sqlStatement, id, stock.Name, stock.Price, stock.Company)

	if err != nil {
		log.Fatalf("Unable to execute the query %v,err")
	}
	rowsAffected, err := res.RowsAffected()
	if err != nil {
		log.Fatalf("Error while checking the affected roes %v ", err)

	}
	fmt.Printf("Total rows/records affected %v", rowsAffected)
	return rowsAffected
}
func deleteStock(id int64) int64 {
	db := createConnection()
	defer db.Close()
	sqlStatement := `DELETE FROM stocks where stockid=$1`
	res, err := db.Exec(sqlStatement, id)
	if err != nil {
		log.Fatalf("Unable to execute the query %v,err")
	}
	rowsAffected, err := res.RowsAffected()
	if err != nil {
		log.Fatalf("Error while checking the affected roes %v ", err)

	}
	fmt.Printf("Total rows/records affected %v", rowsAffected)
	return rowsAffected
}
