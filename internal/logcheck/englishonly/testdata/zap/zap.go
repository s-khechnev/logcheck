package testdata

import (
	"time"

	"go.uber.org/zap"
)

func Test() {
	logger, _ := zap.NewDevelopment()
	sugar := logger.Sugar()

	logger.Info("Hello", // want "Message contains non-English letter: ключ" "Message contains non-English letter: привет"
		zap.Dict("Hello",
			zap.Dict("ключ", zap.String("привет", "qwe"))))

	// Проверка сообщений
	logger.Info("Привет")               // want "Message contains non-English letter: Привет"
	logger.Error("Привет")              // want "Message contains non-English letter: Привет"
	logger.Debug("Привет")              // want "Message contains non-English letter: Привет"
	logger.Warn("Привет")               // want "Message contains non-English letter: Привет"
	logger.Log(zap.InfoLevel, "Привет") // want "Message contains non-English letter: Привет"

	// Sugar методы
	sugar.Info("Привет")  // want "Message contains non-English letter: Привет"
	sugar.Error("Привет") // want "Message contains non-English letter: Привет"
	sugar.Debug("Привет") // want "Message contains non-English letter: Привет"
	sugar.Warn("Привет")  // want "Message contains non-English letter: Привет"

	// Контекстные методы (если используется zap with context wrapper)
	// logger.InfoCtx(ctx, "Привет") // может отсутствовать в стандартном zap

	// Сообщение с полями - поле может содержать не английские символы в значении
	logger.Info("Hello", zap.String("key", "Привет")) // it's ok. value can be non english

	// Проверка ключей полей (должны быть на английском)
	logger.Info("Hello", zap.String("Привет", "hello"))     // want "Message contains non-English letter: Привет"
	logger.Info("Hello", zap.Int("Привет", 123))            // want "Message contains non-English letter: Привет"
	logger.Info("Hello", zap.Int64("Привет", 123))          // want "Message contains non-English letter: Привет"
	logger.Info("Hello", zap.Uint64("Привет", 123))         // want "Message contains non-English letter: Привет"
	logger.Info("Hello", zap.Float64("Привет", 123))        // want "Message contains non-English letter: Привет"
	logger.Info("Hello", zap.Bool("Привет", false))         // want "Message contains non-English letter: Привет"
	logger.Info("Hello", zap.Time("Привет", time.Now()))    // want "Message contains non-English letter: Привет"
	logger.Info("Hello", zap.Duration("Привет", time.Hour)) // want "Message contains non-English letter: Привет"
	logger.Info("Hello", zap.Any("Привет", "qwe"))          // want "Message contains non-English letter: Привет"
	logger.Info("Hello", zap.Uint32("Привет", 123))         // want "Message contains non-English letter: Привет"
	logger.Info("Hello", zap.Uint("Привет", 123))           // want "Message contains non-English letter: Привет"
	logger.Info("Hello", zap.Complex64("Привет", 1+2i))     // want "Message contains non-English letter: Привет"
	logger.Info("Hello", zap.Complex128("Привет", 1+2i))    // want "Message contains non-English letter: Привет"
	logger.Info("Hello", zap.Stringp("Привет", nil))        // want "Message contains non-English letter: Привет"
	logger.Info("Hello", zap.Strings("Привет", []string{})) // want "Message contains non-English letter: Привет"
	logger.Info("Hello", zap.Errors("Привет", []error{}))   // want "Message contains non-English letter: Привет"

	// Sugar методы с ключами-значениями
	sugar.Info("Hello", "Привет", "val")                  // want "Message contains non-English letter: Привет"
	sugar.Info("Hello", "key1", "val1", "Привет", "val2") // want "Message contains non-English letter: Привет"

	// Проверка вложенных полей через Namespace (правильный способ создания вложенности в zap)
	logger.Info("Hello",
		zap.Namespace("request"),
		zap.Namespace("eng"),
		zap.String("method", "GET"),
		zap.Int("status", 200),
	)

	// С неанглийским ключом на верхнем уровне
	logger.Info("Hello", // want "Message contains non-English letter: запрос"
		zap.Namespace("запрос"),
		zap.Namespace("eng"),
		zap.String("method", "GET"),
		zap.Int("status", 200),
	)

	// С неанглийским ключом на вложенном уровне
	logger.Info("Hello", // want "Message contains non-English letter: рус"
		zap.Namespace("request"),
		zap.Namespace("рус"),
		zap.String("method", "GET"),
		zap.Int("status", 200),
	)

	// Проверка конкатенации строк
	logger.Info("Hello" + "привет")                            // want "Message contains non-English letter: Helloпривет"
	sugar.Info("Hello" + "привет")                             // want "Message contains non-English letter: Helloпривет"
	logger.Info("Hello", zap.String("Hello"+"привет", "val1")) // want "Message contains non-English letter: Helloпривет"
	logger.Info("Hello" + "hello1" + "привет" + "hello")       // want "Message contains non-English letter: Hellohello1приветhello"

	// Проверка массивов
	logger.Info("Hello", zap.Strings("Привет", []string{"a", "b"})) // want "Message contains non-English letter: Привет"
	logger.Info("Hello", zap.Ints("Привет", []int{1, 2}))           // want "Message contains non-English letter: Привет"

	// Проверка через zap.Reflect
	type request struct {
		Method string
		Status int
	}

	logger.Info("Hello", zap.Reflect("Привет", request{Method: "GET"})) // want "Message contains non-English letter: Привет"

	// Проверка через zap.Binary
	logger.Info("Hello", zap.Binary("Привет", []byte("data"))) // want "Message contains non-English letter: Привет"

	// Проверка через zap.ByteString
	logger.Info("Hello", zap.ByteString("Привет", []byte("data"))) // want "Message contains non-English letter: Привет"

	// Проверка через zap.Stack
	logger.Info("Hello", zap.Stack("Привет")) // want "Message contains non-English letter: Привет"

	// Проверка через zap.Skip
	logger.Info("Hello", zap.Skip())

	// Пример с кастомным ObjectMarshaler (обычно используется редко в простых случаях)
	type MyStruct struct {
		Name string
		Age  int
	}
}
