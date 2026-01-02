package container

import (
	portsrepositories "github.com/Arthur-Conti/gi_nutri/internal/application/ports/repositories"
	portsservices "github.com/Arthur-Conti/gi_nutri/internal/application/ports/services"
	patientservice "github.com/Arthur-Conti/gi_nutri/internal/application/services/patient"
	resultsservice "github.com/Arthur-Conti/gi_nutri/internal/application/services/results"
	"github.com/Arthur-Conti/gi_nutri/internal/infra/configs"
	patientrepository "github.com/Arthur-Conti/gi_nutri/internal/infra/repositories/patient"
	resultsrepository "github.com/Arthur-Conti/gi_nutri/internal/infra/repositories/results"
)

var BaseContainer *Container

type Container struct {
	Configs     Configs
	Repos       Repos
	Services    Services
}

type Configs struct {
	db           *configs.MongoDB
	logger       *configs.Logger
	errorHandler *configs.AppError
}

type Repos struct {
	PatientRepo portsrepositories.PatientRepository
	ResultsRepo portsrepositories.ResultsRepository
}

type Services struct {
	PatientService portsservices.PatientService
	ResultService  portsservices.ResultsService
}

func NewContainer() *Container {
	return &Container{}
}

func (c *Container) Start() {
	c.startConfigs()
	c.startRepos()
	c.startServices()
}

func (c *Container) startConfigs() {
	c.Configs.db = configs.NewMongoDB()
	if err := c.Configs.db.Connect(); err != nil {
		panic(err)
	}
	c.Configs.logger = configs.GetLogger()
}

func (c *Container) startRepos() {
	c.Repos.PatientRepo = patientrepository.NewRepository(c.Configs.db, c.Configs.logger)
	c.Repos.ResultsRepo = resultsrepository.NewResultsRepository(c.Configs.db, configs.GetLogger())
}

func (c *Container) startServices() {
	c.Services.PatientService = patientservice.NewPatientService(
		c.Repos.PatientRepo,
		c.Repos.ResultsRepo,
	)
	c.Services.ResultService = resultsservice.NewResultsService(
		c.Repos.PatientRepo,
		c.Repos.ResultsRepo,
	)
}
