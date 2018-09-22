package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"text/template"

	"github.com/tealeg/xlsx"
)

var (
	xlsFlag, templateFlag, htmlFlag *string
)

func init() {
	xlsFlag = flag.String("xls", "", "Excel input file")
	templateFlag = flag.String("template", "", "HTML input template")
	htmlFlag = flag.String("html", "", "HTML output file")

	flag.Parse()
}

// Abstracts are article items
type Abstracts struct {
	Title   string
	Authors string
}

func main() {
	var (
		err       error
		tmpl      string
		abstracts []Abstracts
	)
	if tmpl, err = readTemplate(*templateFlag); err != nil {
		fmt.Printf("template read error: %s", err)
		os.Exit(2)
	}
	if abstracts, err = readExcel(*xlsFlag); err != nil {
		fmt.Printf("XLS read error: %s", err)
		os.Exit(2)
	}
	if err = writeHTML(*htmlFlag, tmpl, abstracts); err != nil {
		fmt.Printf("HTML write error: %s", err)
		os.Exit(2)
	}
	fmt.Println("HTML saved")
}

func readTemplate(name string) (string, error) {
	var (
		err    error
		result string
		file   *os.File
		c      int
	)
	file, err = os.Open(name)
	if err != nil {
		return "", err
	}
	defer file.Close()
	data := make([]byte, 10240)
	c, err = file.Read(data)
	if err != nil {
		return "", err
	}
	result = string(data[:c])
	return result, err
}

func readExcel(fileName string) (result []Abstracts, err error) {
	xlFile, err := xlsx.OpenFile(fileName)
	if err != nil {
		return result, err
	}
	if len(xlFile.Sheets) < 1 {
		return result, fmt.Errorf("sheet not found %v", xlFile)
	}
	result = []Abstracts{}
	sheet := xlFile.Sheets[0]
	startAbstracts := false
	for _, rows := range sheet.Rows {
		if !startAbstracts && len(rows.Cells) > 0 && isPositiveInt(rows.Cells[0].Value) {
			startAbstracts = true
		}
		if startAbstracts && len(rows.Cells) > 0 && isPositiveInt(rows.Cells[0].Value) {
			cells := rows.Cells
			item := Abstracts{
				Authors: cells[1].Value,
				Title:   cells[2].Value,
			}
			if item.Title == "" {
				break
			}
			result = append(result, item)
		}
	}
	return result, err
}

func isPositiveInt(val string) bool {
	i, err := strconv.Atoi(val)
	return err == nil && i > 0
}

func writeHTML(fileName, entTemplate string, abstracts []Abstracts) (err error) {
	var (
		fh   *os.File
		data = map[string]interface{}{}
	)
	fh, err = os.Create(fileName)
	if err != nil {
		return err
	}

	articles := template.Must(template.New("abstracts").Parse(entTemplate))
	data["Articles"] = abstracts

	err = articles.Execute(fh, data)
	return err
}
