package main

import (
    "compress/gzip"
    "os"
)

func WriteCompressedJSON(filename string, jsonData []byte) error {
    // Datei zum Schreiben öffnen
    f, err := os.Create(filename)
    if err != nil {
        return err
    }
    defer f.Close()

    // gzip Writer auf die Datei setzen
    gw := gzip.NewWriter(f)
    defer gw.Close()

    // JSON-Daten in den gzip Writer schreiben (komprimieren)
    _, err = gw.Write(jsonData)
    if err != nil {
        return err
    }

    return nil
}
