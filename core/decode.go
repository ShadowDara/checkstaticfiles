package core

import (
    "encoding/base64"
    "fmt"
    "os"
    "path/filepath"
)

type EncodedFile struct {
    Path    string
    Content string // base64-kodierter Inhalt
}

func Main() {
    // Beispielhafte Daten (normalerweise kommen die aus deinem Array)
    files := []EncodedFile{
        {
            Path:    "output/folder1/hello.txt",
            Content: "SGVsbG8gV29ybGQh", // "Hello World!"
        },
        {
            Path:    "output/folder2/readme.md",
            Content: "IyBSZWFkbWUKVGhpcyBpcyBhIHRlc3Qh", // "# Readme\nThis is a test!"
        },
    }

    for _, f := range files {
        err := decodeAndWriteFile(f)
        if err != nil {
            fmt.Fprintf(os.Stderr, "Fehler beim Schreiben von %s: %v\n", f.Path, err)
        }
    }
}

func decodeAndWriteFile(f EncodedFile) error {
    // Existenz prüfen
    if _, err := os.Stat(f.Path); err == nil {
        fmt.Printf("Überspringe bestehende Datei: %s\n", f.Path)
        return nil
    } else if !os.IsNotExist(err) {
        return fmt.Errorf("Fehler beim Prüfen der Datei: %w", err)
    }

    // base64 dekodieren
    decoded, err := base64.StdEncoding.DecodeString(f.Content)
    if err != nil {
        return fmt.Errorf("Fehler beim Decodieren: %w", err)
    }

    // Ordnerstruktur sicherstellen
    dir := filepath.Dir(f.Path)
    err = os.MkdirAll(dir, 0755)
    if err != nil {
        return fmt.Errorf("Fehler beim Erstellen von Verzeichnissen: %w", err)
    }

    // Datei schreiben
    err = os.WriteFile(f.Path, decoded, 0644)
    if err != nil {
        return fmt.Errorf("Fehler beim Schreiben der Datei: %w", err)
    }

    fmt.Printf("Datei erstellt: %s\n", f.Path)
    return nil
}
