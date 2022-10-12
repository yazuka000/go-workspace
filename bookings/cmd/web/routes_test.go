package main

import (
	"testing"

	"github.com/go-chi/chi/v5"
	"github.com/yazuka000/bookings/internal/config"
)

func TestRoute(t *testing.T)  {
	var app config.AppConfig

	mux := routes(&app)

	switch v := mux.(type){
	case *chi.Mux:

	default:
		t.Errorf("type is not *chi.Mux, type is %T", v)
	}
}
