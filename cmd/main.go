package main

import (
	"fmt"

	"github.com/Arthur-Conti/gi_nutri/internal/infra/patient"
)

func main() {
	name := "Gabriela"
	age := 58
	timeDays := 90
	height := 162.0
	weight := 61.0
	usualWeight := 70.0
	p := patient.NewPatient(
		name,
		age,
		timeDays,
		height,
		weight,
		usualWeight,
	)
	p.CalculateIMC()
	p.CalculateAdjustedWeight()
	p.CalculatePercentageWeightAdequacy()
	p.CalculatePercentageWeightChange()

	fmt.Println("IMC Status:", p.FormulasInfo.IMC.Status)
	fmt.Println("IMC Resultado:", p.FormulasInfo.IMC.Result)

	fmt.Println("-----------------------------------------------------------")
	
	fmt.Println("Peso ajustado:", p.FormulasInfo.AdjustedWeight.Result)
	fmt.Println("Peso ideal:", p.FormulasInfo.AdjustedWeight.IdealWeight)
	
	fmt.Println("-----------------------------------------------------------")
	
	fmt.Printf("Porcentagem de adequação de peso: %v%%\n", p.FormulasInfo.PercentageWeightAdequacy.Result)
	fmt.Println("Classificação de adequação de peso:", p.FormulasInfo.PercentageWeightAdequacy.Classification)
	
	fmt.Println("-----------------------------------------------------------")

	fmt.Printf("Porcentagem de mudança de peso: %v%%\n", p.FormulasInfo.PercentageWeightChange.Result)
	fmt.Println("Classificação de mudança de peso:", p.FormulasInfo.PercentageWeightChange.Classification)
}
