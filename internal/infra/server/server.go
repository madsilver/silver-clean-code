package server

type Server interface {
	Start(manager *Manager)
}
