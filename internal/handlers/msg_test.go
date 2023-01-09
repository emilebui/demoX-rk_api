package handlers

import (
	"fmt"
	"github.com/emilebui/demoX-rk_api/pkg/unitest"
	"testing"
)

// Init Functions

// Mocking repository

func createHandler() *MessageHandler {
	// Init fake services
	fakeService := unitest.CreateFakeService()
	return &MessageHandler{
		s: fakeService,
	}
}

func testPushMessage(b *testing.B, h *MessageHandler, count int) {

	// create fake context
	ctx, rec := unitest.CreateEchoContext(unitest.FakeContext{
		Method: "GET",
		Body:   fmt.Sprintf("Message number %d", count),
	})
	_ = h.PushMessage(ctx)

	if rec.Code != 200 {
		b.Errorf("Expected 200, got %d", rec.Code)
	}

}

func BenchmarkMessageHandler_PushMessage(b *testing.B) {
	handler := createHandler()

	for i := 0; i < b.N; i++ {
		testPushMessage(b, handler, i)
	}
}

func TestNewMessageHandler(t *testing.T) {
	service := unitest.CreateFakeService()
	check := NewMessageHandler(service)
	if check == nil {
		t.Error("This should not be nil")
	}
}
