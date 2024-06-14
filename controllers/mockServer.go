package controllers

import (
	"testing"

	"github.com/gin-gonic/gin"
)

func WithRouter(router *gin.Engine) Option {
	return func(c *Container) {
		c.r = router
	}
}

func WithAnalysisController(controller AnalysisController) Option {
	return func(c *Container) {
		if c.controllers == nil {
			c.controllers = &Controllers{}
		}
		c.controllers.AnalysisController = controller
	}
}

func WithAllowedPaths() Option {
	return func(c *Container) {
		c.allowedPaths()
	}
}

type Option func(*Container)

func NewTestServer(t *testing.T, opts ...Option) *Container {
	t.Helper()
	testContainer := &Container{}

	for _, opt := range opts {
		opt(testContainer)
	}

	// Set default AnalysisController if not provided
	if testContainer.controllers == nil || testContainer.controllers.AnalysisController == nil {
		testContainer.controllers = new(Controllers)
		testContainer.controllers.AnalysisController = NewAnalysisController(nil)
	}

	// Set default router if not provided
	if testContainer.r == nil {
		testContainer.r = SetupRouter()
		testContainer.allowedPaths()
	}

	return testContainer
}
