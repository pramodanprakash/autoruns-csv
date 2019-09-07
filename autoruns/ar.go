package autoruns

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/smartystreets/scanners/csv"
)

// Time	Entry Location	Entry	Enabled	Category	Profile	Description	Company	Image Path	Version	Launch String	MD5	SHA-1	PESHA-1	PESHA-256	SHA-256	IMP

// AutoRuns struct defines a row in the autoruns sysinternals CSV output
type AutoRuns struct {
	Time          string `csv:"Time" json:"Time"`
	EntryLocation string `csv:"EntryLocation" json:"EntryLocation"`
	Entry         string `csv:"Entry" json:"Entry"`
	Enabled       string `csv:"Enabled" json:"Enabled"`
	Category      string `csv:"Category" json:"Category"`
	Profile       string `csv:"Description" json:"Description"`
	Company       string `csv:"Company" json:"Company"`
	ImagePath     string `csv:"Image Path" json:"Image Path"`
	Version       string `csv:"Version" json:"Version"`
	LaunchString  string `csv:"Launch String" json:"Launch String"`
	MD5           string `csv:"MD5" json:"MD5"`
	SHA1          string `csv:"SHA-1" json:"SHA-1"`
	PESHA1        string `csv:"PESHA-1" json:"PESHA-1"`
	PESHA256      string `csv:"PESHA-256" json:"PESHA-256"`
	SHA256        string `csv:"SHA-256" json:"SHA-256"`
	IMP           string `csv:"IMP" json:"IMP"`
}

func NewAutoRuns(path string) []AutoRuns {
	var data []AutoRuns
	lines, err := readLines(path)
	if err != nil {
		log.Panic(err)
	}
	csvData := strings.NewReader(strings.Join(lines, "\n"))
	scanner, err := csv.NewStructScanner(csvData)
	if err != nil {
		log.Panic(err)
	}
	for scanner.Scan() {
		var a AutoRuns
		if err := scanner.Populate(&a); err != nil {
			log.Panic(err)
		}
		fmt.Printf("%#v\n", a)
		jSlice, err := json.Marshal(a)
		if err != nil {
			log.Panicln(err)
		}
		fmt.Println(string(jSlice))
		data = append(data, a)
	}
	return data
}

// readLines reads a whole file into memory
// and returns a slice of its lines.
func readLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}
