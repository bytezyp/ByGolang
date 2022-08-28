package main

import (
	"context"
	"fmt"
	"github.com/ClickHouse/clickhouse-go/v2"
	"log"
	"time"
)

func example() error {
	conn := clickhouse.OpenDB(&clickhouse.Options{
		Addr: []string{"172.27.58.246:9000"},
		Auth: clickhouse.Auth{
			Database: "ads_prod",
			Username: "adsprod",
			Password: "adsprod1122",
		},
		Settings: clickhouse.Settings{
			"max_execution_time": 60,
		},
		DialTimeout: 5 * time.Second,
		Compression: &clickhouse.Compression{
			Method: clickhouse.CompressionLZ4,
		},
		//Debug: true,
	})
	conn.SetMaxIdleConns(5)
	conn.SetMaxOpenConns(10)
	conn.SetConnMaxLifetime(time.Hour)
	ctx := clickhouse.Context(context.Background(), clickhouse.WithSettings(clickhouse.Settings{
		"max_block_size": 10,
	}), clickhouse.WithProgress(func(p *clickhouse.Progress) {
		fmt.Println("progress: ", p)
	}))
	if err := conn.PingContext(ctx); err != nil {
		if exception, ok := err.(*clickhouse.Exception); ok {
			fmt.Printf("Catch exception [%d] %s \n%s\n", exception.Code, exception.Message, exception.StackTrace)
		}
		return err
	}

	rows, err := conn.QueryContext(ctx, "SELECT dt,hour,funnel_name FROM dws_ads_midas_funnel_log_sum_prod limit 10", 0, "xxx", time.Now())
	if err != nil {
		return err
	}
	for rows.Next() {
		var (
			col1 string
			col2 string
			col3 string
		)
		if err := rows.Scan(&col1, &col2, &col3); err != nil {
			return err
		}
		fmt.Printf("row: col1=%d, col2=%s, col3=%s\n", col1, col2, col3)
	}
	rows.Close()
	return rows.Err()
}

func main() {
	if err := example(); err != nil {
		log.Fatal(err)
	}
}
