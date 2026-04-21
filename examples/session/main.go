package main

import (
	"context"
	"fmt"

	"go.naturallyfunny.dev/adk/session"
	adksession "google.golang.org/adk/session"
)

type ctxKey string
const ResponseFormatKey ctxKey = "format"

func main() {
	// Assume baseSvc is an existing session implementation (like Zep)
	var baseSvc adksession.Service = &mockService{}

	// Enhance the session service with specific middleware and options
	svc := session.Wrap(baseSvc,
		// Disable storing user messages in the database (Privacy focused)
		session.WithoutUserMessagePersistence(),
		
		// Enforce a specific system-level policy
		session.WithPolicy("Only provide answers based on medical facts."),
		
		// Dynamically control output format via context
		session.EnableDynamicResponseFormat(ResponseFormatKey),
	)

	// Context carrying a specific format instruction
	ctx := context.WithValue(context.Background(), ResponseFormatKey, "JSON")

	resp, _ := svc.Get(ctx, &adksession.GetRequest{
		SessionID: "session-789",
		UserID:    "user-abc",
	})

	fmt.Printf("Session is active with middleware: %s\n", resp.Session.ID())
}

// Mock service for demonstration purposes
type mockService struct{ adksession.Service }
func (m *mockService) Get(_ context.Context, _ *adksession.GetRequest) (*adksession.GetResponse, error) {
	return &adksession.GetResponse{Session: &mockSession{}}, nil
}
type mockSession struct{ adksession.Session }
func (m *mockSession) ID() string { return "mock-id" }
func (m *mockSession) State() adksession.State { return nil }
func (m *mockSession) Events() adksession.Events { return nil }
func (m *mockSession) LastUpdateTime() (t java_time_Time) { return }
