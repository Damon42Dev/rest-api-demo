package utils

type Configuration struct {
	Database DatabaseSetting
	Server   ServerSettings
	App      Application
}

type DatabaseSetting struct {
	Uri         string
	DbName      string
	Collections []string
}

type ServerSettings struct {
	Port string
}

type Application struct {
	Name    string
	Timeout int
}
