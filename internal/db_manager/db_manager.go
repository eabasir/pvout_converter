package db_manager

import (
	"database/sql"
	"fmt"
	"pvout_converter/internal/configs"
	"pvout_converter/internal/types"

	_ "github.com/lib/pq"
)

func ConnectDB(configs *configs.Configs) (*sql.DB, error) {
	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		configs.DB_host, configs.DB_port, configs.DB_user, configs.DB_pass, configs.DB_name)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	fmt.Println("Connected to the database")

	return db, nil
}

func CloseDB(db *sql.DB) {
	err := db.Close()
	if err != nil {
		fmt.Println("Error closing database connection:", err)
	} else {
		fmt.Println("Database connection closed")
	}
}

func Insert(db *sql.DB, pv_data []types.PVData) {

	tx, err := db.Begin()
	if err != nil {
		panic(err.Error())
	}

	for _, pv := range pv_data {

		stmt, err := tx.Prepare("INSERT INTO pv_data (geom, value, data_month) VALUES (ST_GeomFromText($1), $2, $3)")

		if err != nil {
			panic(err.Error())
		}

		defer stmt.Close()

		_, err = stmt.Exec(latLongToString(pv.Latitude, pv.Longitude), pv.Value, pv.Month)
		if err != nil {
			tx.Rollback()
			panic(err.Error())
		}

	}

	fmt.Printf("Committing %d records\n", len(pv_data))

	if err := tx.Commit(); err != nil {
		panic(err.Error())
	}
}

func latLongToString(lat float64, lng float64) string {
	return fmt.Sprintf("Point(%f %f)", lat, lng)
}
