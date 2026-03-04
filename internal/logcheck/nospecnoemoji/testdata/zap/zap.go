package testdata

import (
	"context"

	"go.uber.org/zap"
)

var ctx = context.Background()

func TestSpecialCharsAndEmojisZap() {
	logger := zap.L()

	// Эмодзи и смайлики
	logger.Info("Hello 👋")                  // want "Message contains special char or emoji: Hello 👋"
	logger.Error("Error 😱")                 // want "Message contains special char or emoji: Error 😱"
	logger.Debug("Debug 🐛")                 // want "Message contains special char or emoji: Debug 🐛"
	logger.Warn("Warning ⚠️")               // want "Message contains special char or emoji: Warning ⚠️"
	logger.Log(zap.ErrorLevel, "Message 🔥") // want "Message contains special char or emoji: Message 🔥"

	// Контекстные версии (zap не имеет прямых аналогов Context методов, используем сахар)
	logger.Info("Hello 🌍")   // want "Message contains special char or emoji: Hello 🌍"
	logger.Error("Error 💥")  // want "Message contains special char or emoji: Error 💥"
	logger.Debug("Debug 🔍")  // want "Message contains special char or emoji: Debug 🔍"
	logger.Warn("Warning 💀") // want "Message contains special char or emoji: Warning 💀"

	// Комбинации эмодзи
	logger.Info("Hello 👨‍👩‍👧‍👦") // want "Message contains special char or emoji: Hello 👨‍👩‍👧‍👦"
	logger.Info("Hello 🏳️‍🌈")    // want "Message contains special char or emoji: Hello 🏳️‍🌈"
	logger.Info("Hello 👩‍💻")     // want "Message contains special char or emoji: Hello 👩‍💻"

	// Символы валют и специальные символы
	logger.Info("Price: €100")       // want "Message contains special char or emoji: Price: €100"
	logger.Info("Price: £100")       // want "Message contains special char or emoji: Price: £100"
	logger.Info("Price: ¥100")       // want "Message contains special char or emoji: Price: ¥100"
	logger.Info("Temperature: 25°C") // want "Message contains special char or emoji: Temperature: 25°C"
	logger.Info("Progress: 50%")     // want "Message contains special char or emoji: Progress: 50%"
	logger.Info("Copyright © 2024")  // want "Message contains special char or emoji: Copyright © 2024"
	logger.Info("Registered ®")      // want "Message contains special char or emoji: Registered ®"
	logger.Info("Trademark ™")       // want "Message contains special char or emoji: Trademark ™"
	logger.Info("Section §")         // want "Message contains special char or emoji: Section §"
	logger.Info("Paragraph ¶")       // want "Message contains special char or emoji: Paragraph ¶"

	// Математические символы
	logger.Info("2 + 2 = 4")    // want "Message contains special char or emoji: 2 \\+ 2 = 4"
	logger.Info("5 - 3 = 2")    // want "Message contains special char or emoji: 5 - 3 = 2"
	logger.Info("10 * 5 = 50")  // want "Message contains special char or emoji: 10 \\* 5 = 50"
	logger.Info("100 / 5 = 20") // want "Message contains special char or emoji: 100 / 5 = 20"
	logger.Info("2^3 = 8")      // want "Message contains special char or emoji: 2\\^3 = 8"
	logger.Info("√4 = 2")       // want "Message contains special char or emoji: √4 = 2"
	logger.Info("∞ loop")       // want "Message contains special char or emoji: ∞ loop"
	logger.Info("x ≠ y")        // want "Message contains special char or emoji: x ≠ y"
	logger.Info("x ≤ y")        // want "Message contains special char or emoji: x ≤ y"
	logger.Info("x ≥ y")        // want "Message contains special char or emoji: x ≥ y"

	// Пунктуация и типографские символы
	logger.Info("Hello… world") // want "Message contains special char or emoji: Hello… world"
	logger.Info("Hello—world")  // want "Message contains special char or emoji: Hello—world"
	logger.Info("Hello–world")  // want "Message contains special char or emoji: Hello–world"
	logger.Info("«Hello»")      // want "Message contains special char or emoji: «Hello»"
	logger.Info("„Hello“")      // want "Message contains special char or emoji: „Hello“"
	logger.Info("Hello!")       // want "Message contains special char or emoji: Hello!"
	logger.Info("Hello?")       // want "Message contains special char or emoji: Hello?"
	logger.Info("Hello,")       // want "Message contains special char or emoji: Hello,"
	logger.Info("Hello.")       // want "Message contains special char or emoji: Hello."
	logger.Info("Hello:")       // want "Message contains special char or emoji: Hello:"
	logger.Info("Hello;")       // want "Message contains special char or emoji: Hello;"

	// Стрелки и символы направления
	logger.Info("Next →")     // want "Message contains special char or emoji: Next →"
	logger.Info("Previous ←") // want "Message contains special char or emoji: Previous ←"
	logger.Info("Up ↑")       // want "Message contains special char or emoji: Up ↑"
	logger.Info("Down ↓")     // want "Message contains special char or emoji: Down ↓"
	logger.Info("Left ⇐")     // want "Message contains special char or emoji: Left ⇐"
	logger.Info("Right ⇒")    // want "Message contains special char or emoji: Right ⇒"
	logger.Info("Both ↔")     // want "Message contains special char or emoji: Both ↔"

	// Символы в полях (должны быть ок, проверяем только сообщения)
	logger.Info("Hello", zap.String("key", "value 👋"))  // it's ok. value can contain emojis
	logger.Info("Hello", zap.String("key", "value \t")) // it's ok. value can contain special chars
	logger.Info("Hello", zap.String("key", "value 👋"))  // it's ok. value can contain emojis
	logger.Info("Hello", zap.Any("key", "value \n"))    // it's ok. value can contain special chars

	// Конкатенация с эмодзи
	logger.Info("Hello" + "👋")                           // want "Message contains special char or emoji: Hello👋"
	logger.Info("Hello", zap.String("key"+"👋", "value")) // want "Message contains special char or emoji: key👋"
	logger.Info("Hello" + " " + "👋" + " world")          // want "Message contains special char or emoji: Hello 👋 world"
	logger.Info("Hello", zap.String("key"+"👋", "value")) // want "Message contains special char or emoji: key👋"

	// Множественные символы в одном сообщении
	logger.Info("Hello 👋 world 🌍 and 🚀") // want "Message contains special char or emoji: Hello 👋 world 🌍 and 🚀"

	// Граничные случаи
	logger.Info("")            // ok - пустое сообщение
	logger.Info("HelloWorld")  // ok - только буквы
	logger.Info("Hello123")    // ok - буквы и цифры
	logger.Info("hello_world") // ok - только буквы и подчеркивание
}

func TestSpecialCharsAndEmojisZapSugar() {
	sugar := zap.L().Sugar()

	// Эмодзи и смайлики
	sugar.Info("Hello 👋")                         // want "Message contains special char or emoji: Hello 👋"
	sugar.Error("Error 😱")                        // want "Message contains special char or emoji: Error 😱"
	sugar.Debug("Debug 🐛")                        // want "Message contains special char or emoji: Debug 🐛"
	sugar.Warn("Warning ⚠️")                      // want "Message contains special char or emoji: Warning ⚠️"
	sugar.Log(zap.ErrorLevel, "Message!!!!!!!!!") // want "Message contains special char or emoji: Message!!!!!!!!!"

	// Конкатенация с эмодзи
	sugar.Info("Hello" + "👋")                  // want "Message contains special char or emoji: Hello👋"
	sugar.Info("Hello" + " " + "👋" + " world") // want "Message contains special char or emoji: Hello 👋 world"

	// Множественные символы в одном сообщении
	sugar.Info("Hello 👋 world 🌍 and 🚀") // want "Message contains special char or emoji: Hello 👋 world 🌍 and 🚀"

	// Граничные случаи
	sugar.Info("")            // ok - пустое сообщение
	sugar.Info("HelloWorld")  // ok - только буквы
	sugar.Info("Hello123")    // ok - буквы и цифры
	sugar.Info("hello_world") // ok - только буквы и подчеркивание
}

func TestSpecialCharsAndEmojisZapWithOptions() {
	logger := zap.L()

	// Использование With для создания контекстного логгера
	logger.With(zap.String("trace_id", "123")).Info("Hello 👋") // want "Message contains special char or emoji: Hello 👋"

	// Проверка сообщений с полями, где имена полей содержат спецсимволы
	logger.Info("Hello", zap.String("emoji_key👋", "value")) // want "Message contains special char or emoji: emoji_key👋"

	// Проверка, что поля не триггерят предупреждение на свои значения
	logger.Info("Hello",
		zap.String("key1", "normal value"),
		zap.String("key2", "value with 👋"),
		zap.Int("key3", 123),
	) // it's ok - все спецсимволы только в значениях
}
