package container

import (
	patientservice "github.com/Arthur-Conti/gi_nutri/internal/application/services/patient"
	"github.com/Arthur-Conti/gi_nutri/internal/infra/configs"
	patientcontroller "github.com/Arthur-Conti/gi_nutri/internal/infra/http/controllers/patient"
	patientrepository "github.com/Arthur-Conti/gi_nutri/internal/infra/repositories/patient"
	resultsrepository "github.com/Arthur-Conti/gi_nutri/internal/infra/repositories/results"
)

var BaseContainer *Container

type Container struct {
	Configs Configs
	Repos Repos
	Services Services
	Controllers Controllers
}

type Configs struct {
	db *configs.MongoDB
}

type Repos struct {
	PatientRepo *patientrepository.PatientRepository
	ResultsRepo *resultsrepository.ResultsRepository
}

type Services struct {
	PatientService *patientservice.PatientService
}

type Controllers struct {
	PatientController *patientcontroller.PatientController
}

func NewContainer() *Container {
	return &Container{}
}

func (c *Container) Start() {
	c.startConfigs()
	c.startRepos()
	c.startServices()
	c.startControllers()
}

func (c *Container) startConfigs() {
	c.Configs.db = configs.NewMongoDB()
	if err := c.Configs.db.Connect(); err != nil {
		panic(err)
	}
}

func (c *Container) startRepos() {
	c.Repos.PatientRepo = patientrepository.NewRepository(c.Configs.db)
	c.Repos.ResultsRepo = resultsrepository.NewResultsRepository(c.Configs.db)
}

func (c *Container) startServices() {
	c.Services.PatientService = patientservice.NewPatientService(
		c.Repos.PatientRepo,
		c.Repos.ResultsRepo,
	)
}

func (c *Container) startControllers() {
	c.Controllers.PatientController = patientcontroller.NewPatientController(c.Services.PatientService)
}