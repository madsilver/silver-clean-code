package server

type Server interface {
	Start(port string, manager *Manager)
}
