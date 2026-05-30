package main

import (
	"context"
	"fmt"
	"os"

	"github.com/PrincessFluffyButt937/GoTrace/internal/database"
	"github.com/joho/godotenv"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	fmt.Println(".env loading...")
	if err := godotenv.Load(); err != nil {
		fmt.Printf(".env load error: %s\n", err.Error())
		return
	}
	path := os.Getenv("SAMPLE_XML_DIR")
	if path == "" {
		fmt.Println(".env path empty")
		return
	}

	fmt.Println("Database initialization...")
	cfg, err := InitApiConig()
	if err != nil {
		fmt.Printf("InitDB error: %s\n", err.Error())
		return
	}

	if err := cfg.Scanfolder(path); err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println("Placeholder db write...")
	board := database.StoreBoardDataParams{
		Sn:  "TestSN123",
		Pb:  "TESTpb321",
		Rev: "rev_T",
	}
	fmt.Printf("Writing board: sn-%s pb-%s rev-%s\n", board.Sn, board.Pb, board.Rev)
	db_board, err := cfg.db.StoreBoardData(context.Background(), board)
	if err != nil {
		fmt.Printf("DB store fail: %s\n", err.Error())
		return
	}
	fmt.Printf("returning data board: sn-%s pb-%s rev-%s\n", db_board.Sn, db_board.Pb, db_board.Rev)

	board_slice, err := cfg.db.FetchBoards(context.Background())
	if err != nil {
		fmt.Printf("DB fetch fail: %s\n", err.Error())
		return
	}
	for _, b := range board_slice {
		fmt.Printf("fetching board data: sn-%s pb-%s rev-%s\n", b.Sn, b.Pb, b.Rev)
	}
	fmt.Println("Deleting test data...")
	if err := cfg.db.DeleteBoards(context.Background()); err != nil {
		fmt.Printf("DB delete fail: %s\n", err.Error())
		return
	}
	fmt.Println("Test ended without err raised.")
}
