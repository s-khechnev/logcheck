package testdata

import (
	"context"
	"log/slog"
)

var ctx = context.Background()

// Переменные с чувствительными данными в разных регистрах
var (
	// Нижний регистр
	password   = "secret123"
	apiKey     = "abc123xyz"
	token      = "jwt.token.here"
	creditCard = "4111-1111-1111-1111"
	ssn        = "123-45-6789"
	email      = "user@example.com"

	// Верхний регистр
	PASSWORD   = "secret123"
	APIKEY     = "abc123xyz"
	TOKEN      = "jwt.token.here"
	CREDITCARD = "4111-1111-1111-1111"
	SSN        = "123-45-6789"
	EMAIL      = "user@example.com"

	// Смешанный регистр
	ApiKey     = "abc123xyz"
	AuthToken  = "jwt.token.here"
	CreditCard = "4111-1111-1111-1111"
	UserEmail  = "user@example.com"
	DbPassword = "secret123"

	// С подчеркиваниями
	db_password = "secret123"
	api_key     = "abc123xyz"
	auth_token  = "jwt.token.here"
	credit_card = "4111-1111-1111-1111"
	user_ssn    = "123-45-6789"
	user_email  = "user@example.com"
)

func TestSensitiveData() {
	// Прямая конкатенация с переменными в нижнем регистре
	slog.Info("user password: " + password) // want "Message contains sensitive data: password"
	slog.Info("api_key=" + apiKey)          // want "Message contains sensitive data: apiKey"
	slog.Info("token: " + token)            // want "Message contains sensitive data: token"
	slog.Info("credit card: " + creditCard) // want "Message contains sensitive data: creditCard"
	slog.Info("email: " + email)            // want "Message contains sensitive data: email"

	// Прямая конкатенация с переменными в верхнем регистре
	slog.Info("user password: " + PASSWORD) // want "Message contains sensitive data: PASSWORD"
	slog.Info("api_key=" + APIKEY)          // want "Message contains sensitive data: APIKEY"
	slog.Info("token: " + TOKEN)            // want "Message contains sensitive data: TOKEN"
	slog.Info("credit card: " + CREDITCARD) // want "Message contains sensitive data: CREDITCARD"
	slog.Info("email: " + EMAIL)            // want "Message contains sensitive data: EMAIL"

	// Прямая конкатенация с переменными в смешанном регистре
	slog.Info("api_key=" + ApiKey)          // want "Message contains sensitive data: ApiKey"
	slog.Info("token: " + AuthToken)        // want "Message contains sensitive data: AuthToken"
	slog.Info("credit card: " + CreditCard) // want "Message contains sensitive data: CreditCard"
	slog.Info("email: " + UserEmail)        // want "Message contains sensitive data: UserEmail"
	slog.Info("db password: " + DbPassword) // want "Message contains sensitive data: DbPassword"

	// Прямая конкатенация с переменными, содержащими подчеркивания
	slog.Info("db password: " + db_password) // want "Message contains sensitive data: db_password"
	slog.Info("api_key=" + api_key)          // want "Message contains sensitive data: api_key"
	slog.Info("auth token: " + auth_token)   // want "Message contains sensitive data: auth_token"
	slog.Info("credit card: " + credit_card) // want "Message contains sensitive data: credit_card"
	slog.Info("user email: " + user_email)   // want "Message contains sensitive data: user_email"

	// Чувствительные данные в формате printf с переменными в разных регистрах
	slog.Info("user password: %s", password)  // want "Message contains sensitive data: password"
	slog.Info("api_key=%s", APIKEY)           // want "Message contains sensitive data: APIKEY"
	slog.Info("token: %v", AuthToken)         // want "Message contains sensitive data: AuthToken"
	slog.Info("credit card: %s", credit_card) // want "Message contains sensitive data: credit_card"

	// Чувствительные данные в контексте с переменными в разных регистрах
	slog.InfoContext(ctx, "user password: "+password)  // want "Message contains sensitive data: password"
	slog.ErrorContext(ctx, "api_key="+APIKEY)          // want "Message contains sensitive data: APIKEY"
	slog.DebugContext(ctx, "token: "+AuthToken)        // want "Message contains sensitive data: AuthToken"
	slog.WarnContext(ctx, "credit card: "+credit_card) // want "Message contains sensitive data: credit_card"

	// Чувствительные данные в LogAttrs с переменными в разных регистрах
	slog.LogAttrs(ctx, slog.LevelInfo, "user password: "+password) // want "Message contains sensitive data: password"
	slog.LogAttrs(ctx, slog.LevelDebug, "api_key="+APIKEY)         // want "Message contains sensitive data: APIKEY"
	slog.LogAttrs(ctx, slog.LevelError, "token: "+AuthToken)       // want "Message contains sensitive data: AuthToken"

	// Разные вариации ключевых слов с переменными в разных регистрах
	slog.Info("pass: " + password)        // want "Message contains sensitive data: password"
	slog.Info("pwd: " + PASSWORD)         // want "Message contains sensitive data: PASSWORD"
	slog.Info("secret: " + DbPassword)    // want "Message contains sensitive data: DbPassword"
	slog.Info("auth_token: " + AuthToken) // want "Message contains sensitive data: AuthToken"
	slog.Info("access_token: " + token)   // want "Message contains sensitive data: token"
	slog.Info("refresh_token: " + TOKEN)  // want "Message contains sensitive data: TOKEN"
	slog.Info("api_key: " + ApiKey)       // want "Message contains sensitive data: ApiKey"
	slog.Info("apikey: " + api_key)       // want "Message contains sensitive data: api_key"

	// Конкатенация с разными разделителями
	slog.Info("password=" + PASSWORD)      // want "Message contains sensitive data: PASSWORD"
	slog.Info("password:" + password)      // want "Message contains sensitive data: password"
	slog.Info("password -> " + DbPassword) // want "Message contains sensitive data: DbPassword"
	slog.Info("[password] " + db_password) // want "Message contains sensitive data: db_password"

	// Чувствительные данные как часть сообщения
	slog.Info("user password is " + password + " and it's secret") // want "Message contains sensitive data: password"
	slog.Info("using api key: " + APIKEY + " for auth")            // want "Message contains sensitive data: APIKEY"
	slog.Info("token expired: " + AuthToken)                       // want "Message contains sensitive data: AuthToken"
	slog.Info("credit card " + credit_card + " was used")          // want "Message contains sensitive data: credit_card"

	// Чувствительные данные в аттрибутах (должны быть ок, проверяем только сообщения)
	slog.Info("user authenticated", "password", password)       // want "Message contains sensitive data: password"
	slog.Info("api called", "apiKey", APIKEY)                   // want "Message contains sensitive data: APIKEY"
	slog.Info("token validated", "token", AuthToken)            // want "Message contains sensitive data: AuthToken"
	slog.Info("user logged in", "email", user_email)            // want "Message contains sensitive data: user_email"
	slog.Info("auth failed", slog.String("password", PASSWORD)) // want "Message contains sensitive data: PASSWORD"
	slog.Info("api request", slog.String("apiKey", api_key))    // want "Message contains sensitive data: api_key"
	slog.Info("token check", slog.Any("token", TOKEN))          // want "Message contains sensitive data: TOKEN"

	// Правильные сообщения (без чувствительных данных)
	slog.Info("user authenticated successfully") // ok
	slog.Debug("api request completed")          // ok
	slog.Info("token validated")                 // ok
	slog.Error("database connection failed")     // ok
	slog.Warn("rate limit exceeded")             // ok
	slog.Info("password reset link sent")        // ok
	slog.Info("token bucket algorithm used")     // ok
}
