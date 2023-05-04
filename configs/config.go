package configs

type RemoteApiProfile struct {
	Url string `envconfig:"REMOTE_API_PROFILE_URL"`
}

type RemotePageProfile struct {
	Url string `envconfig:"REMOTE_PROGILE_PROFILE_URL"`
}

type ServerConfig struct {
	GrpcServer struct {
		Port string `envconfig:"GRPC_SERVER_PORT"`
	}
	RestServer struct {
		Port string `envconfig:"REST_SERVER_PORT"`
	}

	RemoteApiProfile  RemoteApiProfile
	RemotePageProfile RemotePageProfile
}
