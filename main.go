package main

import (
	"log"
	"log/slog"
	"os"

	"github.com/natefinch/lumberjack"
)

func ConfigureServiceLogging() {
	log.Println("ConfigureServiceLogging::start")

	log.Println("ConfigureServiceLogging::defineRotator")
	// Configure Lumberjack for log rotation
	logRotator := &lumberjack.Logger{
		Filename:   os.Getenv("LOG_DIR"), // Name of the log file
		MaxSize:    10,                   // Max size in megabytes before rotation
		MaxBackups: 3,                    // Max number of old log files to keep
		MaxAge:     7,                    // Max number of days to retain old log files
		Compress:   true,                 // Compress old log files
	}

	log.Println("ConfigureServiceLogging::defineHandler")

	// Create a slog handler that writes to the Lumberjack rotator
	handler := slog.NewJSONHandler(logRotator, &slog.HandlerOptions{
		AddSource: true, // Add source file and line number to logs
		Level:     slog.LevelInfo,
	})

	log.Println("ConfigureServiceLogging::new_slog")
	logger := slog.New(handler)
	log.Println("ConfigureServiceLogging::setDefault")

	slog.SetDefault(logger)
	log.Println("ConfigureServiceLogging::end")
}

func main() {
	log.Println("Starting Application")
	ConfigureServiceLogging()
	log.Println("Stopping Application")
}
