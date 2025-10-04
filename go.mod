module slog_armv8deathbug

go 1.25.1

godebug fips140=on

require github.com/natefinch/lumberjack v2.0.0+incompatible

require (
	github.com/BurntSushi/toml v1.5.0 // indirect
	gopkg.in/natefinch/lumberjack.v2 v2.2.1 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
)
