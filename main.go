package main

import (
	"io"
	"log"
	"log/slog"
	"os"

	"github.com/lmittmann/tint" // Import the tint library
	slogmulti "github.com/samber/slog-multi"

	"github.com/natefinch/lumberjack"
)

func ConfigureServiceLogging() {

	// Create a new tint handler with desired options
	tintHandler := tint.NewHandler(os.Stdout, &tint.Options{
		Level:      slog.LevelDebug, // Set the minimum logging level
		AddSource:  true,            // Add source file and line number
		TimeFormat: "15:04:05",      // Customize timestamp format
	})

	// Configure Lumberjack for log rotation
	logRotator := &lumberjack.Logger{
		Filename:   os.Getenv("LOG_DIR") + "/app.log", // Name of the log file
		MaxSize:    10,                                // Max size in megabytes before rotation
		MaxBackups: 3,                                 // Max number of old log files to keep
		MaxAge:     7,                                 // Max number of days to retain old log files
		Compress:   true,                              // Compress old log files
	}

	multiWriter := slog.NewJSONHandler(
		io.MultiWriter(logRotator),
		&slog.HandlerOptions{
			AddSource: true,
			Level:     slog.LevelInfo,
		},
	)

	logger := slog.New(slogmulti.Fanout(
		tintHandler,
		multiWriter,
	))

	slog.SetDefault(logger)
}

func main() {
	log.Println("Starting Application")
	ConfigureServiceLogging()
	log.Println("Stopping Application")
}
