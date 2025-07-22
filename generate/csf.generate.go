package main

import (
	"encoding/base64"
	"encoding/json"
	"flag"
	"fmt"
	"io/fs"
	"log"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

type Data struct {
	Paths []string `yaml:"paths"`
}

type EncodedFile struct {
	Path    string // Relativer Pfad
	Content string // base64-kodierter Inhalt
}

// Globale Flag-Variablen als Pointer
var (
	Pkg    *string
	Output *string
	Var    *string
)

func init() {
	// Flags definieren (Pointer speichern)
	Pkg = flag.String("package", "main", "Name des Go-Packages")
	Output = flag.String("output", "checkstaticfiles.data.go", "Zieldatei für die generierte Go-Datei")
	Var = flag.String("variable", "CheckstaticfilesOutputJSONGz", "Place where the Dataarray is saved!")
}

func main() {
	flag.Parse()

	fmt.Println("Package:", *Pkg)
	fmt.Println("Output-Datei:", *Output)
	fmt.Println("Data Variable:", *Var)

	if flag.NArg() > 0 {
		fmt.Println("Zusätzliche Argumente:", flag.Args())
	}

	paths := CheckConfig()
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

func Generate(inputPaths []string) {
	var results []EncodedFile

	for _, path := range inputPaths {
		info, err := os.Stat(path)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Fehler beim Prüfen von %s: %v\n", path, err)
			continue
		}

		if info.IsDir() {
			filepath.WalkDir(path, func(p string, d fs.DirEntry, err error) error {
				if err != nil {
					return err
				}
				if !d.IsDir() {
					ef, err := encodeFile(p)
					if err == nil {
						results = append(results, ef)
					} else {
						fmt.Fprintf(os.Stderr, "Fehler beim Lesen von %s: %v\n", p, err)
					}
				}
				return nil
			})
		} else {
			ef, err := encodeFile(path)
			if err == nil {
				results = append(results, ef)
			} else {
				fmt.Fprintf(os.Stderr, "Fehler beim Lesen von %s: %v\n", path, err)
			}
		}
	}

	jsonData, err := json.MarshalIndent(results, "", "  ")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fehler beim JSON-Encoding: %v\n", err)
		return
	}

	err = os.WriteFile("checkstaticfiles.output.json", jsonData, 0644)
	if err != nil {
		panic(err)
	}

	err = WriteCompressedJSON("checkstaticfiles.output.json.gz", jsonData)
	if err != nil {
		panic(err)
	}

	creategofile()
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
