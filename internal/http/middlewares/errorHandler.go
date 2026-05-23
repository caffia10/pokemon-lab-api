package middlewares

import (
	"errors"
	"fmt"
	"pokemon-lab-api/pkg/mderrors"

	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

func MakeErrorHandler(logger *zap.Logger) func(ctx *fiber.Ctx, err error) error {

	return func(ctx *fiber.Ctx, err error) error {
		// Status code defaults to 500
		code := fiber.StatusInternalServerError
		msg := "An unexpected error happened"

		// Retrieve the custom status code if it's a *fiber.Error
		var (
			e   *fiber.Error
			mde *mderrors.MetadataError
		)

		var lf []zap.Field

		switch {
		case errors.As(err, &e):
			lf = make([]zap.Field, 0, 3)
			code = e.Code
			msg = e.Message
			lf = append(lf,
				zap.String("type", "fiber error"),
				zap.Error(e))

		case errors.As(err, &mde):
			lf = make([]zap.Field, 0, 3+len(mde.Data)+len(mde.Stack))
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
		}

		lf = append(lf, zap.Int("status_code", code))

		logger.Error(msg, lf...)

		return ctx.Status(code).SendString(msg)

	}
}
