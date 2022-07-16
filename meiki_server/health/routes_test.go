package health_test

import (
	"context"
	"net/http"
	"testing"

	"github.com/RainingComputers/Meiki/health"
	"github.com/RainingComputers/Meiki/log"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/suite"

	"github.com/RainingComputers/Meiki/testhelpers"
)

type HealthRoutesTestSuite struct {
	suite.Suite
	cancel context.CancelFunc
	router *gin.Engine
}

func (s *HealthRoutesTestSuite) SetupTest() {
	log.Initialize()

	s.router = gin.Default()
	healthRouter := s.router.Group("/")
	health.CreateRoutes(healthRouter)
}

func TestHealthRoutesTestSuite(t *testing.T) {
	suite.Run(t, new(HealthRoutesTestSuite))
}
func (s *HealthRoutesTestSuite) TestHealthRoutes() {
	gin.SetMode(gin.ReleaseMode)

	req, _ := http.NewRequest("GET", "/", nil)
	testhelpers.AssertResponseString(s.T(), s.router, req, 200, "Healthy")
}
