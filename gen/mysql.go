package gen

import (
	"fmt"
	"os"
	"strings"

	"github.com/zxzixuanwang/sql2ent/util"

	"github.com/zxzixuanwang/sql2ent/parser"

	ddlParser "github.com/miaogaolin/ddlparser/parser"
)

type MysqlGenerator struct {
	Dir string
}

func NewMysqlGenerator(dir string) *MysqlGenerator {
	return &MysqlGenerator{
		Dir: dir,
	}
}

func (g *MysqlGenerator) FromMysql(content string) error {
	p := ddlParser.NewParser()
	tables, err := p.From([]byte(content))
	if err != nil {
		return err
	}
	tablesContent := make(map[string]string)
	fmt.Println("table", tables)
	for _, v := range tables {
		fmt.Println("name", v.Name, "columns", v.Columns, "constraints", v.Constraints)
	}

	for _, t := range tables {
		sch, err := parser.ParseSchema(t)
		if err != nil {
			return err
		}

		schContent, err := parser.ParseTpl(sch)
		if err != nil {
			return err
		}
		tablesContent[sch.TableName] = schContent
	}
	return g.createFile(tablesContent)
}

func (g *MysqlGenerator) FromFile(fileName string) error {
	bytes, err := os.ReadFile(fileName)
	if err != nil {
		return err
	}

	p := ddlParser.NewParser()
	tables, err := p.From(bytes)
	if err != nil {
		return err
	}

	tablesContent := make(map[string]string)
	for _, t := range tables {
		sch, err := parser.ParseSchema(t)
		if err != nil {
			return err
		}

		schContent, err := parser.ParseTpl(sch)
		if err != nil {
			return err
		}
		tablesContent[sch.TableName] = schContent
	}

	return g.createFile(tablesContent)

}

func (g *MysqlGenerator) createFile(tablesContent map[string]string) error {
	files := make(map[string]string)
	for k, v := range tablesContent {
		name := strings.ToLower(k) + ".go"
		files[name] = v
	}

	return util.CreateFiles(files, g.Dir)
}
