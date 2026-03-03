package testdata

import (
	"context"
	"log/slog"
	"time"
)

var ctx = context.Background()

func Test() {
	slog.Info("Привет")                           // want "Message contains non-English letter: Привет"
	slog.Error("Привет")                          // want "Message contains non-English letter: Привет"
	slog.Debug("Привет")                          // want "Message contains non-English letter: Привет"
	slog.Warn("Привет")                           // want "Message contains non-English letter: Привет"
	slog.LogAttrs(ctx, slog.LevelError, "Привет") // want "Message contains non-English letter: Привет"

	slog.InfoContext(ctx, "Привет")               // want "Message contains non-English letter: Привет"
	slog.ErrorContext(ctx, "Привет")              // want "Message contains non-English letter: Привет"
	slog.DebugContext(ctx, "Привет")              // want "Message contains non-English letter: Привет"
	slog.WarnContext(ctx, "Привет")               // want "Message contains non-English letter: Привет"
	slog.LogAttrs(ctx, slog.LevelError, "Привет") // want "Message contains non-English letter: Привет"

	slog.Info("Hello", "Привет", "val") // want "Message contains non-English letter: Привет"

	slog.Info("Hello", "Hello", "Привет") // it's ok. value can be non english

	slog.Info("Hello", slog.String("Привет", "hello"))     // want "Message contains non-English letter: Привет"
	slog.Info("Hello", slog.Int("Привет", 123))            // want "Message contains non-English letter: Привет"
	slog.Info("Hello", slog.Int64("Привет", 123))          // want "Message contains non-English letter: Привет"
	slog.Info("Hello", slog.Uint64("Привет", 123))         // want "Message contains non-English letter: Привет"
	slog.Info("Hello", slog.Float64("Привет", 123))        // want "Message contains non-English letter: Привет"
	slog.Info("Hello", slog.Bool("Привет", false))         // want "Message contains non-English letter: Привет"
	slog.Info("Hello", slog.Time("Привет", time.Now()))    // want "Message contains non-English letter: Привет"
	slog.Info("Hello", slog.Duration("Привет", time.Hour)) // want "Message contains non-English letter: Привет"
	slog.Info("Hello", slog.Any("Привет", "qwe"))          // want "Message contains non-English letter: Привет"

	slog.Info("Hello", "Hello", "значение_это_ок", slog.String("Привет", "qwe")) // want "Message contains non-English letter: Привет"

	slog.Info("Hello", "Hello", "значение_это_ок", "Привет", "Прив") // want "Message contains non-English letter: Привет"

	slog.Info("Hello", "Hello", "значение_это_ок", "Hello", "значение_это_ок")

	slog.Info("Hello", slog.Group("request",
		slog.Group("eng", slog.String("method", "GET"),
			slog.Int("status", 200)),
		slog.String("qwe1", "val"),
		slog.String("qwe2", "val"),
	))

	slog.Info("Hello", slog.Group("request", // want "Message contains non-English letter: метод"
		slog.Group("eng", slog.String("метод", "GET"),
			slog.Int("status", 200)),
		slog.String("qwe1", "val"),
		slog.String("qwe2", "val"),
	))

	slog.Info("Hello", slog.Group("request", // want "Message contains non-English letter: рус_ключ_вложенной_группы"
		slog.Group("рус_ключ_вложенной_группы", slog.String("method", "GET"),
			slog.Int("status", 200)),
		slog.String("qwe1", "val"),
		slog.String("qwe2", "val"),
	))

	slog.Info("Hello", slog.Group("рус_ключ_группы", // want "Message contains non-English letter: рус_ключ_группы"
		slog.Group("key", slog.String("method", "GET"),
			slog.Int("status", 200)),
		slog.String("qwe1", "val"),
		slog.String("qwe2", "val"),
	))

	slog.Info("Hello", slog.Group("key1", // want "Message contains non-English letter: ключ"
		slog.Group("key", slog.String("method", "GET"),
			slog.Int("status", 200)),
		slog.String("ключ", "val"),
		slog.String("qwe2", "val"),
	))

	slog.Info("Hello", slog.Group("key1", // want "Message contains non-English letter: ключ_аттрибута"
		slog.Group("key", slog.String("method", "GET"),
			slog.Int("ключ_аттрибута", 200)),
		slog.String("qwe1", "val"),
		slog.String("qwe2", "val"),
	))
}
