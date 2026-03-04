package testdata

import (
	"errors"
	"testing"

	"go.uber.org/zap"
	"go.uber.org/zap/zaptest"
)

func TestZapAllLevels(t *testing.T) {
	logger := zap.NewNop()
	sugar := logger.Sugar()

	// Базовые уровни логирования через Sugar
	sugar.Info("Aboba")  // want "Message starts with capital letter: Aboba"
	sugar.Error("Aboba") // want "Message starts with capital letter: Aboba"
	sugar.Debug("Aboba") // want "Message starts with capital letter: Aboba"
	sugar.Warn("Aboba")  // want "Message starts with capital letter: Aboba"
	sugar.Fatal("Aboba") // want "Message starts with capital letter: Aboba"
	sugar.Panic("Aboba") // want "Message starts with capital letter: Aboba"

	// С дополнительными аргументами через Sugar
	sugar.Infow("Aboba", "key", "value")        // want "Message starts with capital letter: Aboba"
	sugar.Errorw("Aboba", "error", "qwe")       // want "Message starts with capital letter: Aboba"
	sugar.Debugw("Aboba", "count", 42)          // want "Message starts with capital letter: Aboba"
	sugar.Warnw("Aboba", "user", "alice")       // want "Message starts with capital letter: Aboba"
	sugar.Fatalw("Aboba", "reason", "critical") // want "Message starts with capital letter: Aboba"
	sugar.Panicw("Aboba", "cause", "unknown")   // want "Message starts with capital letter: Aboba"

	// С форматированием через Sugar
	sugar.Infof("Aboba %s", "param")  // want "Message starts with capital letter: Aboba"
	sugar.Errorf("Aboba %d", 42)      // want "Message starts with capital letter: Aboba"
	sugar.Debugf("Aboba %v", "value") // want "Message starts with capital letter: Aboba"
	sugar.Warnf("Aboba %t", true)     // want "Message starts with capital letter: Aboba"
	sugar.Fatalf("Aboba %x", 255)     // want "Message starts with capital letter: Aboba"
	sugar.Panicf("Aboba %s", "panic") // want "Message starts with capital letter: Aboba"

	// Через стандартный Logger с полями
	logger.Info("Aboba", zap.String("key", "value"))        // want "Message starts with capital letter: Aboba"
	logger.Error("Aboba", zap.String("error", "qwe"))       // want "Message starts with capital letter: Aboba"
	logger.Debug("Aboba", zap.Int("count", 42))             // want "Message starts with capital letter: Aboba"
	logger.Warn("Aboba", zap.String("user", "alice"))       // want "Message starts with capital letter: Aboba"
	logger.Fatal("Aboba", zap.String("reason", "critical")) // want "Message starts with capital letter: Aboba"
	logger.Panic("Aboba", zap.String("cause", "unknown"))   // want "Message starts with capital letter: Aboba"

	// С группой полей через Namespace
	logger.Info("Aboba", // want "Message starts with capital letter: Aboba"
		zap.Namespace("request"),
		zap.String("method", "GET"),
		zap.Int("status", 200),
	)
}

func TestZapWithLogger(t *testing.T) {
	logger := zaptest.NewLogger(t)
	sugar := logger.Sugar()

	// Вызовы через экземпляр логгера (Sugar)
	sugar.Info("Aboba")  // want "Message starts with capital letter: Aboba"
	sugar.Error("Aboba") // want "Message starts with capital letter: Aboba"
	sugar.Debug("Aboba") // want "Message starts with capital letter: Aboba"
	sugar.Warn("Aboba")  // want "Message starts with capital letter: Aboba"

	// С атрибутами через Sugar
	sugar.Infow("Aboba", "key", "value")  // want "Message starts with capital letter: Aboba"
	sugar.Errorw("Aboba", "error", "qwe") // want "Message starts with capital letter: Aboba"

	// Через стандартный Logger
	logger.Info("Aboba", zap.String("key", "value"))  // want "Message starts with capital letter: Aboba"
	logger.Error("Aboba", zap.String("error", "qwe")) // want "Message starts with capital letter: Aboba"

	// Через With
	logger.With(zap.String("module", "test")).Info("Aboba") // want "Message starts with capital letter: Aboba"
	sugar.With("module", "test").Info("Aboba")              // want "Message starts with capital letter: Aboba"

	// Через Named
	logger.Named("component").Info("Aboba") // want "Message starts with capital letter: Aboba"
	sugar.Named("component").Info("Aboba")  // want "Message starts with capital letter: Aboba"
}

func TestZapEdgeCases(t *testing.T) {
	logger := zap.NewNop()
	sugar := logger.Sugar()

	// Пустое сообщение (не должно триггерить правило)
	sugar.Info("")
	logger.Info("")

	// Сообщение с маленькой буквы (не должно триггерить правило)
	sugar.Info("aboba")
	logger.Info("aboba")

	// Сообщение с цифрой в начале
	sugar.Info("1Aboba")
	logger.Info("2Aboba")

	// Сообщение со специальными символами
	sugar.Info("!Aboba")
	logger.Info("@Aboba")

	// Сообщение на других языках
	sugar.Info("Привет")  // want "Message starts with capital letter: Привет"
	logger.Info("Привет") // want "Message starts with capital letter: Привет"

	// Сообщение с пробелом в начале
	sugar.Info(" Aboba")
	logger.Info(" Aboba")

	// Многословные сообщения
	sugar.Info("Very Long Message With Multiple Words That Starts With Capital Letter")  // want "Message starts with capital letter: Very Long Message With Multiple Words That Starts With Capital Letter"
	logger.Info("Very Long Message With Multiple Words That Starts With Capital Letter") // want "Message starts with capital letter: Very Long Message With Multiple Words That Starts With Capital Letter"
}

func TestZapDifferentLoggers(t *testing.T) {
	// Production конфигурация
	prodLogger, _ := zap.NewProduction()
	defer prodLogger.Sync()
	prodSugar := prodLogger.Sugar()
	prodSugar.Info("Aboba")  // want "Message starts with capital letter: Aboba"
	prodSugar.Error("Aboba") // want "Message starts with capital letter: Aboba"

	// Development конфигурация
	devLogger, _ := zap.NewDevelopment()
	defer devLogger.Sync()
	devSugar := devLogger.Sugar()
	devSugar.Info("Aboba")  // want "Message starts with capital letter: Aboba"
	devSugar.Error("Aboba") // want "Message starts with capital letter: Aboba"

	// Кастомная конфигурация
	config := zap.NewProductionConfig()
	config.Level = zap.NewAtomicLevelAt(zap.DebugLevel)
	customLogger, _ := config.Build()
	defer customLogger.Sync()
	customSugar := customLogger.Sugar()
	customSugar.Debug("Aboba") // want "Message starts with capital letter: Aboba"
	customSugar.Warn("Aboba")  // want "Message starts with capital letter: Aboba"
}

func TestZapNested(t *testing.T) {
	logger := zap.NewNop()
	sugar := logger.Sugar()

	// Вложенные вызовы
	if true {
		sugar.Info("Aboba") // want "Message starts with capital letter: Aboba"

		for i := 0; i < 1; i++ {
			logger.Info("Aboba") // want "Message starts with capital letter: Aboba"

			func() {
				sugar.Debug("Aboba") // want "Message starts with capital letter: Aboba"
			}()
		}
	}

	// В горутине
	go func() {
		sugar.Warn("Aboba") // want "Message starts with capital letter: Aboba"
	}()
}

func TestZapWithError(t *testing.T) {
	logger := zap.NewNop()
	sugar := logger.Sugar()
	err := errors.New("some error")

	// Оборачивание ошибок через Logger
	logger.Error("Failed to process request", zap.Error(err)) // want "Message starts with capital letter: Failed to process request"
	logger.Error("Database connection lost", zap.Error(err))  // want "Message starts with capital letter: Database connection lost"

	// Через Sugar
	sugar.Errorw("Failed to process request", "error", err) // want "Message starts with capital letter: Failed to process request"
	sugar.Errorf("Failed to process request: %v", err)      // want "Message starts with capital letter: Failed to process request"

	// С дополнительным контекстом
	logger.Error("Operation failed", // want "Message starts with capital letter: Operation failed"
		zap.Error(err),
		zap.String("operation", "delete"),
		zap.Int("attempt", 3),
	)
}
