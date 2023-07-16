package server

import "silver-clean-code/internal/adapter"

type Controller interface {
	FindByID(ctx adapter.ContextServer) error
	Create(ctx adapter.ContextServer) error
}
