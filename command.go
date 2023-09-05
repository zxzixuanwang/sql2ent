package main

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"strings"

	"github.com/zxzixuanwang/sql2ent/util"

	_ "github.com/go-sql-driver/mysql"
	"github.com/urfave/cli"
	"github.com/zxzixuanwang/sql2ent/gen"
)

const (
	flagSrc      = "src"
	flagDir      = "dir"
	mysql_target = "mysql_target"
)

func MysqlDDL(cli *cli.Context) error {
	src := cli.String(flagSrc)
	dir := cli.String(flagDir)
	target := cli.String(mysql_target)
	src = strings.TrimSpace(src)

	if len(target) > 0 {
		db, err := sql.Open("mysql", target)
		if err != nil {
			log.Fatal("Failed to connect to MySQL database:", err)
		}
		defer db.Close()

		// 获取所有表名
		tables, err := getTables(db)
		if err != nil {
			log.Fatal("Failed to get tables:", err)
			return err
		}

		// 获取每个表的结构
		content := ""
		for _, tableName := range tables {
			tableStructure, err := getTableStructure(db, tableName)
			if err != nil {
				log.Fatal("Failed to get table structure:", err)
				return err
			}

			content += fmt.Sprintf("%s\n\n", tableStructure)
		}

		g := gen.NewMysqlGenerator(dir)
		err = g.FromMysql(content)
		if err != nil {
			return err
		}
	} else {
		if len(src) == 0 {
			return errors.New("expected path or path globbing patterns, but nothing found")
		}

		files, err := util.MatchFiles(src)
		if err != nil {
			return err
		}

		if len(files) == 0 {
			return errors.New("sql not matched")
		}

		g := gen.NewMysqlGenerator(dir)

		for _, f := range files {
			err := g.FromFile(f)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

// 获取所有表名
func getTables(db *sql.DB) ([]string, error) {
	rows, err := db.Query("SHOW TABLES")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tables []string
	for rows.Next() {
		var tableName string
		if err := rows.Scan(&tableName); err != nil {
			return nil, err
		}
		tables = append(tables, tableName)
	}

	return tables, rows.Err()
}

// 获取表结构
func getTableStructure(db *sql.DB, tableName string) (string, error) {
	rows, err := db.Query("SHOW CREATE TABLE " + tableName)
	if err != nil {
		return "", err
	}
	defer rows.Close()

	if rows.Next() {
		var tableName string
		var tableStructure string
		if err := rows.Scan(&tableName, &tableStructure); err != nil {
			return "", err
		}
		return tableStructure, nil
	}

	return "", fmt.Errorf("Table not found: %s", tableName)
}
