package main

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"os"
	"path/filepath"
)

func creategofile() {
	inputFile := "checkstaticfiles.output.json"

	jsonData, err := os.ReadFile(inputFile)
	if err != nil {
		panic(fmt.Errorf("Fehler beim Lesen der JSON-Datei: %w", err))
	}

	var buf bytes.Buffer
	gz := gzip.NewWriter(&buf)
	_, err = gz.Write(jsonData)
	if err != nil {
		panic(fmt.Errorf("Fehler beim Komprimieren: %w", err))
	}
	gz.Close()

	compressed := buf.Bytes()

	f, err := os.Create(*Output)
	if err != nil {
		panic(fmt.Errorf("Fehler beim Erstellen von %s: %w", *Output, err))
	}
	defer f.Close()

	fmt.Fprintf(f, "package %s\n\n", *Pkg)
	fmt.Fprintf(f, "// %s enthält gzip-komprimierte JSON-Daten aus %q\n", *Var, filepath.Base(inputFile))
	fmt.Fprintf(f, "var %s = []byte{\n", *Var)

	for i, b := range compressed {
		if i%12 == 0 {
			fmt.Fprintf(f, "    ")
		}
		fmt.Fprintf(f, "0x%02x", b)
		if i < len(compressed)-1 {
			fmt.Fprintf(f, ", ")
		}
		if (i+1)%12 == 0 {
			fmt.Fprintln(f)
		}
	}
	fmt.Fprintln(f, "}")

	fmt.Printf("✔ Go-Datei %s mit Variable %s wurde erstellt (%d bytes, gzip)\n", *Output, *Var, len(compressed))
}
