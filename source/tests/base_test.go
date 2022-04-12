package tests

import (
	_ "app/source/_testinit/fixture"
	"app/source/configuration"
	"app/source/repository"
	"app/source/routes"
	"app/source/tests/testcontainers"
	"app/source/tests/testutils"
	"log"
	"os"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var router *gin.Engine
var mysqlContainer testcontainers.ContainerResult
var db *gorm.DB

var runTest = testutils.CreateForEach(setUp, tearDown)

func TestMain(m *testing.M) {
	log.Println("Starting test setup")
	start := time.Now()
	BeforeAll()
	log.Printf("Setup took %s seconds\n", time.Since(start))
	exitVal := m.Run()
	os.Exit(exitVal)
}

func BeforeAll() {
	router = routes.InitRouter()
	mysqlContainer = testcontainers.SetupMysqlContainer(
		&testcontainers.Testcontainer{
			Database:     configuration.Config.DBName,
			RootPassword: configuration.Config.DBPass,
		},
	)
	repository.OpenConnectionDb()
}

func setUp() {
	repository.AutoMigrate()
}

func tearDown() {
	repository.DropAll()
}
