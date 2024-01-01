package routes

type IServer interface {
	Get(route string)
}

func SetupRoutes(server IServer) {

	server.Get("/get")
}
