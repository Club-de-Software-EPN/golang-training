package main

import(
	"encoding/csv" // a file will have csv format
	"fmt" // format package
	"os" // create files, close files
)

func saveDataOnCSVFile(data [][]string, fileName string){
	// "fileName" param only refers to the name of the file but not to its extension

	// Create a file
	membersFile, err := os.Create(fmt.Sprintf("%s.csv", fileName))
	// Close file using defer
	defer membersFile.Close()

	if err != nil {
		fmt.Println("An error has ocurred")
	}

	// create a csv writer
	writerCsvFile := csv.NewWriter(membersFile)
	// buffer --> file "".csv
	defer writerCsvFile.Flush()

	// save data as csv format rows
	for _, rowMember := range data {
		_ = writerCsvFile.Write(rowMember)
	}
}

func main(){

	// data
	data := [][]string{
		{"MemberName", "TeamName", "Email"},
		{"Alejandro Llanganate", "Ciberseguridad", "alejo@gmail.com"},
		{"Juan Perez", "AI", "juanito@gmail.com"},
		{"Victoria Endara", "Web", "vic@gmail.com"},
	}

	domainsData := [][]string{
		{"Domain", "IP"},
		{"www.epn.edu.ec", "192.168.4.3"},
	}

	saveDataOnCSVFile(data, "clubData")
	saveDataOnCSVFile(domainsData, "epnData")
}