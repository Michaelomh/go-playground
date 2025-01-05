package utils

import (
	"calc-api/types"
	"log/slog"
	"net/http"
)

func LogCalcCompleted(values types.CalculatorRequest, result types.CalculatorResponse) {
	slog.Info("calculation completed",
		slog.Int("number1", values.First),
		slog.Int("number2", values.Second),
		slog.Int("result", result.Result),
		slog.Int("status", http.StatusOK),
	)
}

func LogCalcReceived(r *http.Request) {
	slog.Info("request",
		slog.String("method", r.Method),
		slog.String("ip address", r.Header.Get("X-Forwarded-For")),
		slog.String("url", r.URL.String()),
		slog.Int("status", http.StatusOK),
	)
}
