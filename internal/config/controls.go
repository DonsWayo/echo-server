package config

type TimesControl struct {
	Min int
	Max int
}

type ControlsConfig struct {
	Times TimesControl
}

func NewDefaultControlsConfig() ControlsConfig {
	return ControlsConfig{
		Times: TimesControl{
			Min: getEnvInt("CONTROLS__TIMES__MIN", 0),
			Max: getEnvInt("CONTROLS__TIMES__MAX", 60000),
		},
	}
}
