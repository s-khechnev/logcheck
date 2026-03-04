package testdata

import (
	"context"
	"log/slog"
)

var ctx = context.Background()

func TestSpecialCharsAndEmojis() {
	// Эмодзи и смайлики
	slog.Info("Hello 👋")                             // want "Message contains special char or emoji: Hello 👋"
	slog.Error("Error 😱")                            // want "Message contains special char or emoji: Error 😱"
	slog.Debug("Debug 🐛")                            // want "Message contains special char or emoji: Debug 🐛"
	slog.Warn("Warning ⚠️")                          // want "Message contains special char or emoji: Warning ⚠️"
	slog.LogAttrs(ctx, slog.LevelError, "Message 🔥") // want "Message contains special char or emoji: Message 🔥"

	slog.InfoContext(ctx, "Hello 🌍")   // want "Message contains special char or emoji: Hello 🌍"
	slog.ErrorContext(ctx, "Error 💥")  // want "Message contains special char or emoji: Error 💥"
	slog.DebugContext(ctx, "Debug 🔍")  // want "Message contains special char or emoji: Debug 🔍"
	slog.WarnContext(ctx, "Warning 💀") // want "Message contains special char or emoji: Warning 💀"

	// Комбинации эмодзи
	slog.Info("Hello 👨‍👩‍👧‍👦") // want "Message contains special char or emoji: Hello 👨‍👩‍👧‍👦"
	slog.Info("Hello 🏳️‍🌈")    // want "Message contains special char or emoji: Hello 🏳️‍🌈"
	slog.Info("Hello 👩‍💻")     // want "Message contains special char or emoji: Hello 👩‍💻"

	// Символы валют и специальные символы
	slog.Info("Price: €100")       // want "Message contains special char or emoji: Price: €100"
	slog.Info("Price: £100")       // want "Message contains special char or emoji: Price: £100"
	slog.Info("Price: ¥100")       // want "Message contains special char or emoji: Price: ¥100"
	slog.Info("Temperature: 25°C") // want "Message contains special char or emoji: Temperature: 25°C"
	slog.Info("Progress: 50%")     // want "Message contains special char or emoji: Progress: 50%"
	slog.Info("Copyright © 2024")  // want "Message contains special char or emoji: Copyright © 2024"
	slog.Info("Registered ®")      // want "Message contains special char or emoji: Registered ®"
	slog.Info("Trademark ™")       // want "Message contains special char or emoji: Trademark ™"
	slog.Info("Section §")         // want "Message contains special char or emoji: Section §"
	slog.Info("Paragraph ¶")       // want "Message contains special char or emoji: Paragraph ¶"

	// Математические символы
	slog.Info("2 + 2 = 4")    // want "Message contains special char or emoji: 2 \\+ 2 = 4"
	slog.Info("5 - 3 = 2")    // want "Message contains special char or emoji: 5 - 3 = 2"
	slog.Info("10 * 5 = 50")  // want "Message contains special char or emoji: 10 \\* 5 = 50"
	slog.Info("100 / 5 = 20") // want "Message contains special char or emoji: 100 / 5 = 20"
	slog.Info("2^3 = 8")      // want "Message contains special char or emoji: 2\\^3 = 8"
	slog.Info("√4 = 2")       // want "Message contains special char or emoji: √4 = 2"
	slog.Info("∞ loop")       // want "Message contains special char or emoji: ∞ loop"
	slog.Info("x ≠ y")        // want "Message contains special char or emoji: x ≠ y"
	slog.Info("x ≤ y")        // want "Message contains special char or emoji: x ≤ y"
	slog.Info("x ≥ y")        // want "Message contains special char or emoji: x ≥ y"

	// Пунктуация и типографские символы
	slog.Info("Hello… world") // want "Message contains special char or emoji: Hello… world"
	slog.Info("Hello—world")  // want "Message contains special char or emoji: Hello—world"
	slog.Info("Hello–world")  // want "Message contains special char or emoji: Hello–world"
	slog.Info("«Hello»")      // want "Message contains special char or emoji: «Hello»"
	slog.Info("„Hello“")      // want "Message contains special char or emoji: „Hello“"
	slog.Info("Hello!")       // want "Message contains special char or emoji: Hello!"
	slog.Info("Hello?")       // want "Message contains special char or emoji: Hello?"
	slog.Info("Hello,")       // want "Message contains special char or emoji: Hello,"
	slog.Info("Hello.")       // want "Message contains special char or emoji: Hello."
	slog.Info("Hello:")       // want "Message contains special char or emoji: Hello:"
	slog.Info("Hello;")       // want "Message contains special char or emoji: Hello;"

	// Стрелки и символы направления
	slog.Info("Next →")     // want "Message contains special char or emoji: Next →"
	slog.Info("Previous ←") // want "Message contains special char or emoji: Previous ←"
	slog.Info("Up ↑")       // want "Message contains special char or emoji: Up ↑"
	slog.Info("Down ↓")     // want "Message contains special char or emoji: Down ↓"
	slog.Info("Left ⇐")     // want "Message contains special char or emoji: Left ⇐"
	slog.Info("Right ⇒")    // want "Message contains special char or emoji: Right ⇒"
	slog.Info("Both ↔")     // want "Message contains special char or emoji: Both ↔"

	// Символы в аттрибутах (должны быть ок, проверяем только сообщения)
	slog.Info("Hello", "key", "value 👋")              // it's ok. value can contain emojis
	slog.Info("Hello", "key", "value \t")             // it's ok. value can contain special chars
	slog.Info("Hello", slog.String("key", "value 👋")) // it's ok. value can contain emojis
	slog.Info("Hello", slog.Any("key", "value \n"))   // it's ok. value can contain special chars

	// Конкатенация с эмодзи
	slog.Info("Hello" + "👋")                            // want "Message contains special char or emoji: Hello👋"
	slog.Info("Hello", "key"+"👋", "value")              // want "Message contains special char or emoji: key👋"
	slog.Info("Hello" + " " + "👋" + " world")           // want "Message contains special char or emoji: Hello 👋 world"
	slog.Info("Hello", slog.String("key"+"👋", "value")) // want "Message contains special char or emoji: key👋"

	// Множественные символы в одном сообщении
	slog.Info("Hello 👋 world 🌍 and 🚀") // want "Message contains special char or emoji: Hello 👋 world 🌍 and 🚀"

	// Граничные случаи
	slog.Info("")            // ok - пустое сообщение
	slog.Info("HelloWorld")  // ok - только буквы
	slog.Info("Hello123")    // ok - буквы и цифры
	slog.Info("hello_world") // ok - только буквы и подчеркивание
}
