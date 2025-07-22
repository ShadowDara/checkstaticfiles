package main

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"os"
	"log"

	"encoding/base64"
	"encoding/json"
	"io/fs"
	"path/filepath"
)

type Data struct {
	Paths []string `yaml:"paths"`
}

func main() {
	var paths = CheckConfig()
	Generate(paths)
}

func CheckConfig() []string {
	wd, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	log.Println("Executing folder:", wd)

	data, err := os.ReadFile("checkstaticfiles.config.yaml")
	if err != nil {
		panic(err)
	}

	var p Data
	err = yaml.Unmarshal(data, &p)
	if err != nil {
		panic(err)
	}

	return p.Paths
}

type EncodedFile struct {
	Path    string // Relativer Pfad
	Content string // base64-kodierter Inhalt
}

func Generate(inputPaths []string) {
	var results []EncodedFile

	for _, path := range inputPaths {
		info, err := os.Stat(path)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Fehler beim Prüfen von %s: %v\n", path, err)
			continue
		}

		if info.IsDir() {
			// Ordner rekursiv durchlaufen
			filepath.WalkDir(path, func(p string, d fs.DirEntry, err error) error {
				if err != nil {
					return err
				}
				if !d.IsDir() {
					ef, err := encodeFile(p)
					if err == nil {
						results = append(results, ef)
					}
				}
				return nil
			})
		} else {
			// Einzelne Datei
			ef, err := encodeFile(path)
			if err == nil {
				results = append(results, ef)
			}
		}
	}

	jsonData, err := json.MarshalIndent(results, "", "  ")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fehler beim JSON-Encoding: %v\n", err)
		return
	}

	log.Println("JSON Data: ")
	fmt.Println(string(jsonData))
}

func encodeFile(path string) (EncodedFile, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return EncodedFile{}, err
	}

	encoded := base64.StdEncoding.EncodeToString(data)

	return EncodedFile{
		Path:    path,
		Content: encoded,
	}, nil
}
