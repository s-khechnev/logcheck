package testdata

import (
	"context"

	"go.uber.org/zap"
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
	logger, _ := zap.NewProduction()
	sugar := logger.Sugar()
	defer logger.Sync()

	// Прямая конкатенация с переменными в нижнем регистре
	sugar.Info("user password: " + password) // want "Message contains sensitive data: password"
	sugar.Info("api_key=" + apiKey)          // want "Message contains sensitive data: apiKey"
	sugar.Info("token: " + token)            // want "Message contains sensitive data: token"
	sugar.Info("credit card: " + creditCard) // want "Message contains sensitive data: creditCard"
	sugar.Info("email: " + email)            // want "Message contains sensitive data: email"

	// Прямая конкатенация с переменными в верхнем регистре
	sugar.Info("user password: " + PASSWORD) // want "Message contains sensitive data: PASSWORD"
	sugar.Info("api_key=" + APIKEY)          // want "Message contains sensitive data: APIKEY"
	sugar.Info("token: " + TOKEN)            // want "Message contains sensitive data: TOKEN"
	sugar.Info("credit card: " + CREDITCARD) // want "Message contains sensitive data: CREDITCARD"
	sugar.Info("email: " + EMAIL)            // want "Message contains sensitive data: EMAIL"

	// Прямая конкатенация с переменными в смешанном регистре
	sugar.Info("api_key=" + ApiKey)          // want "Message contains sensitive data: ApiKey"
	sugar.Info("token: " + AuthToken)        // want "Message contains sensitive data: AuthToken"
	sugar.Info("credit card: " + CreditCard) // want "Message contains sensitive data: CreditCard"
	sugar.Info("email: " + UserEmail)        // want "Message contains sensitive data: UserEmail"
	sugar.Info("db password: " + DbPassword) // want "Message contains sensitive data: DbPassword"

	// Прямая конкатенация с переменными, содержащими подчеркивания
	sugar.Info("db password: " + db_password) // want "Message contains sensitive data: db_password"
	sugar.Info("api_key=" + api_key)          // want "Message contains sensitive data: api_key"
	sugar.Info("auth token: " + auth_token)   // want "Message contains sensitive data: auth_token"
	sugar.Info("credit card: " + credit_card) // want "Message contains sensitive data: credit_card"
	sugar.Info("user email: " + user_email)   // want "Message contains sensitive data: user_email"

	// Чувствительные данные в формате printf с переменными в разных регистрах
	sugar.Infof("user password: %s", password)  // want "Message contains sensitive data: password"
	sugar.Infof("api_key=%s", APIKEY)           // want "Message contains sensitive data: APIKEY"
	sugar.Infof("token: %v", AuthToken)         // want "Message contains sensitive data: AuthToken"
	sugar.Infof("credit card: %s", credit_card) // want "Message contains sensitive data: credit_card"

	// Чувствительные данные в контексте с переменными в разных регистрах
	sugar.Infow("user password: "+password, "context", ctx)  // want "Message contains sensitive data: password"
	sugar.Errorw("api_key="+APIKEY, "context", ctx)          // want "Message contains sensitive data: APIKEY"
	sugar.Debugw("token: "+AuthToken, "context", ctx)        // want "Message contains sensitive data: AuthToken"
	sugar.Warnw("credit card: "+credit_card, "context", ctx) // want "Message contains sensitive data: credit_card"

	// Чувствительные данные в structured logging с переменными в разных регистрах
	logger.Info("user password: " + password)  // want "Message contains sensitive data: password"
	logger.Debug("api_key=" + APIKEY)          // want "Message contains sensitive data: APIKEY"
	logger.Error("token: " + AuthToken)        // want "Message contains sensitive data: AuthToken"
	logger.Warn("credit card: " + credit_card) // want "Message contains sensitive data: credit_card"

	// Разные вариации ключевых слов с переменными в разных регистрах
	sugar.Info("pass: " + password)        // want "Message contains sensitive data: password"
	sugar.Info("pwd: " + PASSWORD)         // want "Message contains sensitive data: PASSWORD"
	sugar.Info("secret: " + DbPassword)    // want "Message contains sensitive data: DbPassword"
	sugar.Info("auth_token: " + AuthToken) // want "Message contains sensitive data: AuthToken"
	sugar.Info("access_token: " + token)   // want "Message contains sensitive data: token"
	sugar.Info("refresh_token: " + TOKEN)  // want "Message contains sensitive data: TOKEN"
	sugar.Info("api_key: " + ApiKey)       // want "Message contains sensitive data: ApiKey"
	sugar.Info("apikey: " + api_key)       // want "Message contains sensitive data: api_key"

	// Конкатенация с разными разделителями
	sugar.Info("password=" + PASSWORD)      // want "Message contains sensitive data: PASSWORD"
	sugar.Info("password:" + password)      // want "Message contains sensitive data: password"
	sugar.Info("password -> " + DbPassword) // want "Message contains sensitive data: DbPassword"
	sugar.Info("[password] " + db_password) // want "Message contains sensitive data: db_password"

	// Чувствительные данные как часть сообщения
	sugar.Info("user password is " + password + " and it's secret") // want "Message contains sensitive data: password"
	sugar.Info("using api key: " + APIKEY + " for auth")            // want "Message contains sensitive data: APIKEY"
	sugar.Info("token expired: " + AuthToken)                       // want "Message contains sensitive data: AuthToken"
	sugar.Info("credit card " + credit_card + " was used")          // want "Message contains sensitive data: credit_card"

	// Чувствительные данные в полях (должны быть ок, проверяем только сообщения)
	logger.Info("user authenticated", zap.String("password", password)) // want "Message contains sensitive data: password"
	logger.Info("api called", zap.String("apiKey", APIKEY))             // want "Message contains sensitive data: APIKEY"
	logger.Info("token validated", zap.String("token", AuthToken))      // want "Message contains sensitive data: AuthToken"
	logger.Info("user logged in", zap.String("email", user_email))      // want "Message contains sensitive data: user_email"
	logger.Info("auth failed", zap.String("password", PASSWORD))        // want "Message contains sensitive data: PASSWORD"
	logger.Info("api request", zap.String("apiKey", api_key))           // want "Message contains sensitive data: api_key"
	logger.Info("token check", zap.Any("token", TOKEN))                 // want "Message contains sensitive data: TOKEN"

	// Правильные сообщения (без чувствительных данных)
	sugar.Info("user authenticated successfully") // ok
	sugar.Debug("api request completed")          // ok
	sugar.Info("token validated")                 // ok
	sugar.Error("database connection failed")     // ok
	sugar.Warn("rate limit exceeded")             // ok
	sugar.Info("password reset link sent")        // ok
	sugar.Info("token bucket algorithm used")     // ok
}
