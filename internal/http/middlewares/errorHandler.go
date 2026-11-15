package middlewares

import (
	"errors"
	"fmt"
	"pokemon-lab-api/pkg/mderrors"

	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

// maxLoggedBodyBytes caps the request body written to the log.
const maxLoggedBodyBytes = 4 * 1024

func MakeErrorHandler(logger *zap.Logger) func(ctx *fiber.Ctx, err error) error {

	return func(ctx *fiber.Ctx, err error) error {
		// Status code defaults to 500
		code := fiber.StatusInternalServerError
		msg := "An unexpected error happened"

		var (
			e   *fiber.Error
			mde *mderrors.MetadataError
		)

		lf := requestFields(ctx)

		switch {
		case errors.As(err, &e):
			code = e.Code
			msg = e.Message
			lf = append(lf,
				zap.String("type", "fiber error"),
				zap.Error(err))

		case errors.As(err, &mde):
			lf = append(lf,
				zap.String("type", "metadata error"),
				zap.Error(mde))
			for _, d := range mde.Data {
				lf = append(lf, zap.Any(d.Key, d.Value))
			}
			for i, s := range mde.Stack {
				lf = append(lf, zap.String(fmt.Sprintf("stack_%d", i), s))
			}
			// TODO: custom error mapping
			code = fiber.ErrConflict.Code

		default:
			lf = append(lf,
				zap.String("type", "unknown error"),
				zap.Error(err))
		}

		lf = append(lf, zap.Int("status_code", code))

		logger.Error(msg, lf...)

		return ctx.Status(code).SendString(msg)
	}
}

// requestFields describes the incoming request so failures can be reproduced.
func requestFields(ctx *fiber.Ctx) []zap.Field {
	body := ctx.Body()
	truncated := false
	if len(body) > maxLoggedBodyBytes {
		body = body[:maxLoggedBodyBytes]
		truncated = true
	}

	return []zap.Field{
		zap.String("method", ctx.Method()),
		zap.String("path", ctx.OriginalURL()),
		zap.String("content_type", string(ctx.Request().Header.ContentType())),
		zap.Int("body_size", len(ctx.Body())),
		// copy: fasthttp reuses the body buffer after the handler returns
		zap.String("body", string(body)),
		zap.Bool("body_truncated", truncated),
	}
}
