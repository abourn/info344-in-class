package models

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
)

// capitalized to export to other packages
type Zip struct {
	Code  string
	City  string
	State string
}

type ZipSlice []*Zip // a slice of pointers to Zip structs

type ZipIndex map[string]ZipSlice // a map from strings to a slice of pointers to Zip structs

func LoadZips(fileName string) (ZipSlice, error) {
	f, err := os.Open(fileName)
	if err != nil { // super common in Go (because there is no try/catch in Go)
		return nil, fmt.Errorf("error opening file: %v", err)
	}
	reader := csv.NewReader(f)
	// okay, time to ignore the header row
	_, err = reader.Read() // ignore the value (the header) it give us back with the _

	if err != nil {
		return nil, fmt.Errorf("error reading header row: %v", err)
	}

	zips := make(ZipSlice, 0, 43000) // we don't want to keep reallocating underlying array so lets use make func
	for {                            // iterate until we decide to return out of it
		fields, err := reader.Read() // one line at a time -- slice of strings, one string for each column
		if err == io.EOF {           // above other err check so we don't return it as other error
			return zips, nil // end of file so lets return the slice we made
		}

		if err != nil {
			return nil, fmt.Errorf("error reading record: %v", err)
		}
		z := &Zip{ // with the & we create the new zip struct and then set z equal to the pointer (because ZipSlice is a slice of pointers to Zip structs)
			Code:  fields[0],
			City:  fields[3],
			State: fields[6],
		}
		zips = append(zips, z)
	}
}
