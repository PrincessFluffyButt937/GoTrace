package main

import (
	"context"
	"database/sql"
	"fmt"
	"os"
	"strconv"

	"github.com/PrincessFluffyButt937/GoTrace/internal/cache"
	"github.com/PrincessFluffyButt937/GoTrace/internal/database"
	"github.com/PrincessFluffyButt937/GoTrace/internal/structure"
	_ "github.com/mattn/go-sqlite3"
)

type apiConfig struct {
	db         *database.Queries
	cache      *cache.Cache
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

func (cfg *apiConfig) saveToDatabase(traceData structure.FormatedXMLdata) error {
	cachedSN := cfg.cache.SN.Contains(traceData.SerialNumber)
	if !cachedSN {
		boardData := database.StoreBoardDataParams{
			Sn:  traceData.SerialNumber,
			Pb:  traceData.Project,
			Rev: traceData.Revision,
		}
		if _, err := cfg.db.StoreBoardData(context.Background(), boardData); err != nil {
			fmt.Println(err.Error())
		}
	}
	for _, compEntry := range traceData.Components {
		cashedHU := cfg.cache.HU.Contains(compEntry.HandlingUnit)
		if !cashedHU {
			hadnlingUnitData := database.StoreCompDataParams{
				Hu:  compEntry.HandlingUnit,
				Pn:  compEntry.PartNumber,
				Lot: compEntry.LotCode,
			}
			if _, err := cfg.db.StoreCompData(context.Background(), hadnlingUnitData); err != nil {
				fmt.Println(err.Error())
			}
		}

		trace := database.StoreTraceDataParams{
			Sn:      traceData.SerialNumber,
			Hu:      compEntry.HandlingUnit,
			RefList: refListToString(compEntry.RefereceList, true),
			Placed:  traceData.PlacedIn,
		}
		if _, err := cfg.db.StoreTraceData(context.Background(), trace); err != nil {
			fmt.Println(err.Error())
		}
	}
	return nil
}

func (cfg *apiConfig) loadSerialNumberData(serialNumber string) (structure.DatabaseSerialData, error) {
	rows, err := cfg.db.FetchBoardBySN(context.Background(), serialNumber)
	if err != nil {
		return structure.DatabaseSerialData{}, err
	}

	formattedData := structure.DatabaseSerialData{
		Components: make([]structure.DatabaseComponentData, len(rows)),
	}

	for i, row := range rows {
		if i == 0 {
			formattedData.SerialNumber = row.Sn
			formattedData.Project = row.Pb
			formattedData.Revision = row.Rev
			formattedData.PlacedIn = row.Placed
		}
		formattedData.Components[i].HandlingUnit = row.Hu
		formattedData.Components[i].PartNumber = row.Pn
		formattedData.Components[i].LotCode = row.Lot
		formattedData.Components[i].RefereceList = row.RefList
	}
	return formattedData, nil
}

func (cfg *apiConfig) loadHandlingUnitData(HandlingUnit string) (structure.DatabaseHandlingData, error) {
	rows, err := cfg.db.FetchHandlingUnitData(context.Background(), HandlingUnit)
	if err != nil {
		return structure.DatabaseHandlingData{}, nil
	}

	HandlingUnitData := structure.DatabaseHandlingData{
		SerialNumbers: make([]structure.DatabaseHandlingSerialMatch, len(rows)),
	}

	for i, row := range rows {
		if i == 0 {
			HandlingUnitData.HandlingUnit = row.Hu
			HandlingUnitData.PartNumber = row.Pn
			HandlingUnitData.LotCode = row.Lot
		}
		HandlingUnitData.SerialNumbers[i].SerialNumber = row.Sn
		HandlingUnitData.SerialNumbers[i].Project = row.Pb
		HandlingUnitData.SerialNumbers[i].Revision = row.Rev
		HandlingUnitData.SerialNumbers[i].PlacedIn = row.Placed
		HandlingUnitData.SerialNumbers[i].RefereceList = row.RefList
	}

	return HandlingUnitData, nil
}

func (cfg *apiConfig) fetchSerialNumberMatches(HU string) ([]string, error) {
	return nil, nil
	//to be implemented
}
