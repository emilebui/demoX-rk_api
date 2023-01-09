package handlers_test

import (
	"github.com/emilebui/demoX-rk_api/internal/handlers"
	"github.com/emilebui/demoX-rk_api/pkg/unitest"
	"testing"
)

// Init functions

func checkGreeter(b *testing.B, h *handlers.DemoHandler) {
	// Create random input
	ctx, rec := unitest.CreateEchoContext(unitest.FakeContext{
		Method: "GET",
		//Body:   fmt.Sprintf("Message number %d", count),
	})
	_ = h.Greeter(ctx)

	if rec.Code != 200 {
		b.Errorf("Expected 200, got %d", rec.Code)
	}
}

// Test Functions

func BenchmarkDemoHandler_Greeter(b *testing.B) {
	fakeHandler := &handlers.DemoHandler{}

	for i := 0; i < b.N; i++ {
		checkGreeter(b, fakeHandler)
	}
}

func TestNewDemoHandler(t *testing.T) {
	check := handlers.NewDemoHandler()
	if check == nil {
		t.Error("This should not be nil")
	}
}
