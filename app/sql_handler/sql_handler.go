package sql_handler

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"strings"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type SQLHandler struct {
	db *sql.DB
}

func NewHandler(dataSource string) (*SQLHandler, error) {
	db, err := sql.Open("mysql", dataSource+"?parseTime=true")
	if err != nil {
		return nil, fmt.Errorf("init db connection: %w", err)
	}

	ctx := context.Background()
	// TODO タイムアウト時間を定数化
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()
	err = db.PingContext(ctx)
	if err != nil {
		return nil, fmt.Errorf("init db connection: %w", err)
	}

	return &SQLHandler{db: db}, nil
}

func (h *SQLHandler) QueryContext(ctx context.Context, query string, args ...interface{}) (Rows, error) {
	log.Printf("[sql handler] QueryContext, query: %s, args: %v", strings.ReplaceAll(query, "\n", " "), args)
	ctx, cancel := context.WithTimeout(ctx, 2*time.Second)
	defer cancel()
	return h.db.QueryContext(ctx, query, args...)
}
