package core

import (
    "encoding/base64"
    "encoding/json"
    "fmt"
    "os"
    "log"
    "path/filepath"
    "bytes"
    "compress/gzip"
    "io"
)

type EncodedFile struct {
    Path    string
    Content string // base64-kodierter Inhalt
}

func Main(data []byte) {
    // gzip-Daten entpacken
    r, err := gzip.NewReader(bytes.NewReader(data))
    if err != nil {
        panic(fmt.Errorf("Fehler beim Erstellen des gzip Readers: %w", err))
    }
    defer r.Close()

    // Entpackte Daten lesen
    inhalt, err := io.ReadAll(r)
    if err != nil {
        panic(fmt.Errorf("Fehler beim Lesen der entpackten Daten: %w", err))
    }

    var files []EncodedFile
    err = json.Unmarshal(inhalt, &files)
    if err != nil {
        fmt.Println("Fehler beim Parsen der JSON-Datei:", err)
        return
    }

    for _, f := range files {
        err := decodeAndWriteFile(f)
        if err != nil {
            fmt.Fprintf(os.Stderr, "Fehler beim Schreiben von %s: %v\n", f.Path, err)
        }
    }

    log.Println("Finished file checking and creating!")
}

func decodeAndWriteFile(f EncodedFile) error {
    // Existenz prüfen
    if _, err := os.Stat(f.Path); err == nil {
        log.Printf("Skipped existing file: %s\n", f.Path)
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

    log.Printf("File created: %s\n", f.Path)
    return nil
}
