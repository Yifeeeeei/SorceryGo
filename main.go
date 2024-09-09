package main

import (
	"encoding/json"
	"fmt"
	"os"
	"time"
)

func main() {

	args := os.Args
	if len(args) == 2 && args[1] == "gen" {
		// generate json
		templatePath := "template_mass_producer_params_xlsx.json"
		templateParams := MassProducerParams{}
		// read the template
		templateFile, err := os.Open(templatePath)
		if err != nil {
			panic(err)
		}
		defer templateFile.Close()
		decoder := json.NewDecoder(templateFile)
		err = decoder.Decode(&templateParams)
		if err != nil {
			panic(err)
		}
		// write the template
		outputPath := "mass_producer_params_xlsx.json"
		outputFile, err := os.Create(outputPath)
		if err != nil {
			panic(err)
		}
		defer outputFile.Close()
		// marshal with indent
		encoder := json.NewEncoder(outputFile)
		encoder.SetIndent("", "    ")
		err = encoder.Encode(templateParams)

		if err != nil {
			panic(err)
		}
		return

	}
	m, err := NewMassProducer("mass_producer_params_xlsx.json")
	if err != nil {
		panic(err)
	}
	timeNow := time.Now()
	m.Produce()
	timeElapsed := time.Since(timeNow)
	fmt.Println("Time elapsed: ", timeElapsed)
}
