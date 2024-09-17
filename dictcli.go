package main

import (
	_ "embed"
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

//go:embed tools/dictionary.json
var dictionaryData []byte

// WordEntry 表示单词词典中的一个条目
type WordEntry struct {
	Word       string   `json:"word"`
	Definition string   `json:"definition"`
	Synonyms   []string `json:"synonyms"`
}

// Dictionary 是一个映射，键是单词，值是 WordEntry 结构
type Dictionary map[string]WordEntry

// loadDictionary 从指定的 JSON 文件加载词典并返回一个 Dictionary
func loadDictionary(filename string) (Dictionary, error) {
	file, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	var entries []WordEntry
	err = json.Unmarshal(file, &entries)
	if err != nil {
		return nil, err
	}

	dictionary := make(Dictionary)
	for _, entry := range entries {
		dictionary[strings.ToLower(entry.Word)] = entry
	}

	return dictionary, nil
}

func loadDictionaryFromEmbeddedData() (Dictionary, error) {
	var entries []WordEntry
	err := json.Unmarshal(dictionaryData, &entries)
	if err != nil {
		return nil, err
	}

	dictionary := make(Dictionary)
	for _, entry := range entries {
		dictionary[strings.ToLower(entry.Word)] = entry
	}

	return dictionary, nil
}

// queryWord 查询给定单词的定义并返回结果
func queryWord(dict Dictionary, word string) {
	word = strings.ToLower(word)
	if entry, exists := dict[word]; exists {
		fmt.Printf("Word: %s\nDefinition: %s\n", entry.Word, entry.Definition)
		if len(entry.Synonyms) > 0 {
			fmt.Printf("Synonyms: %s\n", strings.Join(entry.Synonyms, ", "))
		}
	} else {
		fmt.Printf("Word '%s' not found in dictionary.\n", word)
	}
}
func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: dicttool <word_to_search>")
		return
	}

	// 获取命令行参数
	word := os.Args[1]

	// 加载词典
	dictionary, err := loadDictionaryFromEmbeddedData()
	if err != nil {
		fmt.Println("Error loading dictionary:", err)
		return
	}

	// 查询单词定义
	queryWord(dictionary, word)
}
