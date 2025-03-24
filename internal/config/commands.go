package config

type CommandConfig struct {
	Header string
	Query  string
}

type CommandsConfig struct {
	HTTPCode      CommandConfig
	HTTPBody      CommandConfig
	HTTPEnvBody   CommandConfig
	HTTPHeaders   CommandConfig
	Time          CommandConfig
	File          CommandConfig
}

func NewDefaultCommandsConfig() CommandsConfig {
	return CommandsConfig{
		HTTPCode: CommandConfig{
			Header: "X-ECHO-CODE",
			Query:  "echo_code",
		},
		HTTPBody: CommandConfig{
			Header: "X-ECHO-BODY",
			Query:  "echo_body",
		},
		HTTPEnvBody: CommandConfig{
			Header: "X-ECHO-ENV",
			Query:  "echo_env",
		},
		HTTPHeaders: CommandConfig{
			Header: "X-ECHO-HEADER",
			Query:  "echo_header",
		},
		Time: CommandConfig{
			Header: "X-ECHO-TIME",
			Query:  "echo_time",
		},
		File: CommandConfig{
			Header: "X-ECHO-FILE",
			Query:  "echo_file",
		},
	}
}
