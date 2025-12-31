package main

import (
	"fmt"

	"github.com/Arthur-Conti/gi_nutri/internal/domain/entities/patient"
)

func main() {
	patientOpts := patient.PatientOpts{
		Name:             "Arthur",
		Age:              22,
		Sex:              patient.PatientSexMale,
		Height:           176.0,
		Weight:           65.0,
		UsualWeight:      70.0,
		TimeDays:         30,
		PhysicalActivity: patient.PhysicalActivitySedentary,
	}
	p := patient.NewPatient(patientOpts)

	if err := p.CalculateIMC(); err != nil {
		fmt.Printf("Erro ao calcular IMC: %v\n", err)
		return
	}

	if err := p.CalculateAdjustedWeight(); err != nil {
		fmt.Printf("Aviso ao calcular AdjustedWeight: %v\n", err)
	}

	if err := p.CalculatePercentageWeightAdequacy(); err != nil {
		fmt.Printf("Erro ao calcular PercentageWeightAdequacy: %v\n", err)
	}

	if err := p.CalculatePercentageWeightChange(); err != nil {
		fmt.Printf("Erro ao calcular PercentageWeightChange: %v\n", err)
	}

	if err := p.CalculateEER(); err != nil {
		fmt.Printf("Erro ao calcular EER: %v\n", err)
		return
	}

	if err := p.CalculateTMB(false, false, false, true, 25.0); err != nil {
		fmt.Printf("Erro ao calcular TMB: %v\n", err)
		return
	}

	fmt.Println("IMC Status:", p.GetFormulas().IMC.Status)
	fmt.Println("IMC Resultado:", p.GetFormulas().IMC.Result)

	fmt.Println("-----------------------------------------------------------")

	if p.GetFormulas().EER != nil {
		fmt.Println("EER:", p.GetFormulas().EER.Result)
	}

	fmt.Println("-----------------------------------------------------------")

	if p.GetFormulas().TMB != nil {
		fmt.Println("TMB:", p.GetFormulas().TMB.Result)
	}
}
