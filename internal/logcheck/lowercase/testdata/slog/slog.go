package testdata

import (
	"context"
	"errors"
	"log/slog"
	"os"
	"testing"
)

func TestSlogAllLevels(t *testing.T) {
	// Базовые уровни логирования
	slog.Info("Aboba")  // want "Message starts with capital letter: Aboba"
	slog.Error("Aboba") // want "Message starts with capital letter: Aboba"
	slog.Debug("Aboba") // want "Message starts with capital letter: Aboba"
	slog.Warn("Aboba")  // want "Message starts with capital letter: Aboba"

	// С дополнительными атрибутами
	slog.Info("Aboba", "key", "value")               // want "Message starts with capital letter: Aboba"
	slog.Error("Aboba", "error", "qwe")              // want "Message starts with capital letter: Aboba"
	slog.Debug("Aboba", slog.Int("count", 42))       // want "Message starts with capital letter: Aboba"
	slog.Warn("Aboba", slog.String("user", "alice")) // want "Message starts with capital letter: Aboba"

	// С группой атрибутов
	slog.Info("Aboba", slog.Group("request", // want "Message starts with capital letter: Aboba"
		slog.String("method", "GET"),
		slog.Int("status", 200),
	))

	// Контекстные версии
	ctx := context.Background()
	slog.InfoContext(ctx, "Aboba")  // want "Message starts with capital letter: Aboba"
	slog.ErrorContext(ctx, "Aboba") // want "Message starts with capital letter: Aboba"
	slog.DebugContext(ctx, "Aboba") // want "Message starts with capital letter: Aboba"
	slog.WarnContext(ctx, "Aboba")  // want "Message starts with capital letter: Aboba"

	// С контекстом и атрибутами
	slog.InfoContext(ctx, "Aboba", "key", "value")               // want "Message starts with capital letter: Aboba"
	slog.ErrorContext(ctx, "Aboba", "error", "qwe")              // want "Message starts with capital letter: Aboba"
	slog.DebugContext(ctx, "Aboba", slog.Int("count", 42))       // want "Message starts with capital letter: Aboba"
	slog.WarnContext(ctx, "Aboba", slog.String("user", "alice")) // want "Message starts with capital letter: Aboba"
}

func TestSlogWithLogger(t *testing.T) {
	logger := slog.Default()

	// Вызовы через экземпляр логгера
	logger.Info("Aboba")  // want "Message starts with capital letter: Aboba"
	logger.Error("Aboba") // want "Message starts with capital letter: Aboba"
	logger.Debug("Aboba") // want "Message starts with capital letter: Aboba"
	logger.Warn("Aboba")  // want "Message starts with capital letter: Aboba"

	// С атрибутами
	logger.Info("Aboba", "key", "value")  // want "Message starts with capital letter: Aboba"
	logger.Error("Aboba", "error", "qwe") // want "Message starts with capital letter: Aboba"

	// Контекстные через логгер
	ctx := context.Background()
	logger.InfoContext(ctx, "Aboba")  // want "Message starts with capital letter: Aboba"
	logger.ErrorContext(ctx, "Aboba") // want "Message starts with capital letter: Aboba"
	logger.DebugContext(ctx, "Aboba") // want "Message starts with capital letter: Aboba"
	logger.WarnContext(ctx, "Aboba")  // want "Message starts with capital letter: Aboba"

	// Через With
	logger.With("module", "test").Info("Aboba") // want "Message starts with capital letter: Aboba"
	logger.WithGroup("request").Error("Aboba")  // want "Message starts with capital letter: Aboba"
}

func TestSlogEdgeCases(t *testing.T) {
	// Пустое сообщение (не должно триггерить правило)
	slog.Info("")
	slog.Error("")

	// Сообщение с маленькой буквы (не должно триггерить правило)
	slog.Info("aboba")
	slog.Error("aboba")

	// Сообщение с цифрой в начале
	slog.Info("1Aboba")
	slog.Error("2Aboba")

	// Сообщение со специальными символами
	slog.Info("!Aboba")
	slog.Error("@Aboba")

	// Сообщение на других языках
	slog.Info("Привет")  // want "Message starts with capital letter: Привет"
	slog.Error("Привет") // want "Message starts with capital letter: Привет"

	// Сообщение с пробелом в начале
	slog.Info(" Aboba")
	slog.Error(" Aboba")

	// Многословные сообщения
	slog.Info("Very Long Message With Multiple Words That Starts With Capital Letter") // want "Message starts with capital letter: Very Long Message With Multiple Words That Starts With Capital Letter"
}

func TestSlogDifferentLoggers(t *testing.T) {
	// Кастомный логгер
	customLogger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	customLogger.Info("Aboba")  // want "Message starts with capital letter: Aboba"
	customLogger.Error("Aboba") // want "Message starts with capital letter: Aboba"

	// JSON логгер
	jsonLogger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	jsonLogger.Info("Aboba")  // want "Message starts with capital letter: Aboba"
	jsonLogger.Error("Aboba") // want "Message starts with capital letter: Aboba"

	// Логгер с уровнем
	levelLogger := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelDebug,
	}))
	levelLogger.Debug("Aboba") // want "Message starts with capital letter: Aboba"
	levelLogger.Warn("Aboba")  // want "Message starts with capital letter: Aboba"
}

func TestSlogNested(t *testing.T) {
	logger := slog.Default()

	// Вложенные вызовы
	if true {
		slog.Info("Aboba") // want "Message starts with capital letter: Aboba"

		for i := 0; i < 1; i++ {
			logger.Error("Aboba") // want "Message starts with capital letter: Aboba"

			func() {
				slog.Debug("Aboba") // want "Message starts with capital letter: Aboba"
			}()
		}
	}

	// В горутине
	go func() {
		slog.Warn("Aboba") // want "Message starts with capital letter: Aboba"
	}()
}

func TestSlogLogAttrs(t *testing.T) {
	// LogAttrs метод
	logger := slog.Default()
	logger.LogAttrs(context.Background(), slog.LevelInfo, "Aboba", // want "Message starts with capital letter: Aboba"
		slog.String("key", "value"))

	logger.LogAttrs(context.Background(), slog.LevelError, "Aboba", // want "Message starts with capital letter: Aboba"
		slog.Int("code", 500))

	// Прямой вызов Log
	logger.Log(context.Background(), slog.LevelInfo, "Aboba") // want "Message starts with capital letter: Aboba"
}

func TestSlogWithError(t *testing.T) {
	err := errors.New("some error")

	// Оборачивание ошибок
	slog.Error("Failed to process request", "error", err) // want "Message starts with capital letter: Failed to process request"
	slog.Error("Database connection lost", err)           // want "Message starts with capital letter: Database connection lost"

	// ErrorContext
	slog.ErrorContext(context.Background(), "Operation failed", err) // want "Message starts with capital letter: Operation failed"
}
