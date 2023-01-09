package unitest

import (
	"github.com/emilebui/demoX-rk_api/internal/services"
	"github.com/stretchr/testify/mock"
)

type FakeRepo struct {
	mock.Mock
}

func (r *FakeRepo) PushMessage(msg string) error {
	args := r.Called(msg)
	return args.Error(0)
}

func CreateFakeService() *services.Service {
	logger := CreateFakeLogger()
	repo := new(FakeRepo)
	_ = repo.On("PushMessage", mock.Anything).Return(nil)
	return services.NewService(repo, logger)
}
