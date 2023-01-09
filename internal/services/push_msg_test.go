package services_test

import (
	"github.com/emilebui/demoX-rk_api/internal/services"
	"github.com/emilebui/demoX-rk_api/pkg/unitest"
	"testing"
)

func TestNewService(t *testing.T) {
	repo := new(unitest.FakeRepo)
	logger := unitest.CreateFakeLogger()
	service := services.NewService(repo, logger)
	if service == nil {
		t.Error("This should not be nil!")
	}
}

func TestService_Process(t *testing.T) {
	service := unitest.CreateFakeService()
	err := service.Process("blah")
	if err != nil {
		t.Error("This should not cause an error!")
	}
}
