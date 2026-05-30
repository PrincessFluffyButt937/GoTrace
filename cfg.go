package main

import (
	"database/sql"
	"fmt"
	"os"
	"strconv"

	"github.com/PrincessFluffyButt937/GoTrace/internal/database"
	_ "github.com/mattn/go-sqlite3"
)

type apiConfig struct {
	db         *database.Queries
	parseLimit int
}

func InitApiConig() (apiConfig, error) {
	cfg := apiConfig{}

	dbPath := os.Getenv("DB_PATH")
	if dbPath == "" {
		return apiConfig{}, fmt.Errorf("DB_PATH variable empty or missing in .env file.")
	}
	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		return apiConfig{}, err
	}
	dbQueries := database.New(db)
	cfg.db = dbQueries

	ParLim := os.Getenv("PARSE_LIMIT")
	if ParLim == "" {
		cfg.parseLimit = 25
		fmt.Printf("PARSE_LIMIT empty or missing, parsing limited to 25 by default.")
	} else {
		limit, err := strconv.Atoi(ParLim)
		if err != nil {
			cfg.parseLimit = 25
			fmt.Printf("Non number PARSE_LIMIT not allowed, parsing limited to 25 by default.")
		} else {
			if limit <= 0 {
				cfg.parseLimit = 25
				fmt.Println("Negative or 0 PARSE_LIMIT not allowed, parsing limited to 25 by default.")
			} else {
				cfg.parseLimit = limit
			}
		}
	}
	return cfg, nil
}
