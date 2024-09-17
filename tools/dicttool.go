package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

type WordEntry struct {
	Word       string   `json:"word"`
	Definition string   `json:"definition"`
	Synonyms   []string `json:"synonyms"`
}

func parseWordNetData(filename string) ([]WordEntry, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var entries []WordEntry
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "  ") || strings.HasPrefix(line, "\t") {
			continue // Skip comments or empty lines
		}
		// 解析行数据，这里需要根据文件的实际格式进行解析
		// 这是一个示例，具体解析逻辑需要根据文件格式调整
		fields := strings.Split(line, "|")
		if len(fields) < 2 {
			continue
		}
		wordData := strings.Fields(fields[0])
		definition := strings.TrimSpace(fields[1])

		entry := WordEntry{
			Word:       wordData[4], // 假设第5个字段是单词
			Definition: definition,
			Synonyms:   []string{}, // 可以添加逻辑解析同义词
		}
		entries = append(entries, entry)
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return entries, nil
}
func writeEntriesToJSON(entries []WordEntry, outputFilename string) error {
	file, err := os.Create(outputFilename)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	return encoder.Encode(entries)
}
func main() {
	inputFilename := "data.noun"        // 输入文件名
	outputFilename := "dictionary.json" // 输出文件名

	entries, err := parseWordNetData(inputFilename)
	if err != nil {
		fmt.Println("Error parsing data:", err)
		return
	}

	err = writeEntriesToJSON(entries, outputFilename)
	if err != nil {
		fmt.Println("Error writing JSON:", err)
		return
	}

	fmt.Println("Dictionary data has been written to", outputFilename)
}
