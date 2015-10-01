package graph

import (
	"encoding/csv"
	"io"
	"os"

	"fmt"
)

/*
	Exported
*/


/*
	Reads in csv file as (undirected for now) edge list.
	E.g. 1,2,5 means from node 1 to node 2 with weight 5
*/
func ReadEdgeList(filepath string) {
	readCsv(filepath)
}



/*
	Internal
*/

func readCsv(filepath string) error {
	csvfile, err := os.Open(filepath)
	if err != nil { return err }
	defer csvfile.Close()

	reader := csv.NewReader(csvfile)

	reader.FieldsPerRecord = -1 // see the Reader struct information below
	reader.TrimLeadingSpace = true


	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil { return err }

		fmt.Println(record)
	}

	return nil
}