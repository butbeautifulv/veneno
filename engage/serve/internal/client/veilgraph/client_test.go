package veilgraph

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetJSON_withoutOAuth_noAuthorizationHeader(t *testing.T) {
	var gotAuth string
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		gotAuth = r.Header.Get("Authorization")
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"ok":true}`))
	}))
	defer srv.Close()

	c := New(Config{BaseURL: srv.URL})
	raw, err := c.GetJSON(context.Background(), "/v1/categories")
	if err != nil {
		t.Fatal(err)
	}
	if string(raw) != `{"ok":true}` {
		t.Fatalf("body %s", raw)
	}
	if gotAuth != "" {
		t.Fatalf("expected no Authorization, got %q", gotAuth)
	}
}

func TestGetJSON_withAuthBroker(t *testing.T) {
	broker := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/v1/token" {
			http.NotFound(w, r)
			return
		}
		_ = json.NewEncoder(w).Encode(map[string]any{
			"access_token": "broker-tok",
			"expires_in":   300,
			"token_type":   "Bearer",
		})
	}))
	defer broker.Close()

	var gotAuth string
	api := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		gotAuth = r.Header.Get("Authorization")
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"ok":true}`))
	}))
	defer api.Close()

	c := New(Config{
		BaseURL:                api.URL,
		UseAuthBroker:          true,
		AuthBrokerURL:          broker.URL,
		AuthBrokerServiceToken: "svc",
		AuthBrokerServiceID:    "veneno-engage",
		AuthBrokerAudience:     "veil-api",
	})
	if _, err := c.GetJSON(context.Background(), "/v1/categories"); err != nil {
		t.Fatal(err)
	}
	if gotAuth != "Bearer broker-tok" {
		t.Fatalf("auth=%q", gotAuth)
	}
}
