package handlers

import (
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/proxy"
	"github.com/google/uuid"
)

func ProxyTo(target string) fiber.Handler {
	return func(c *fiber.Ctx) error {
		// ❌ jangan forward JWT
		c.Request().Header.Del("Authorization")

		// ✅ inject identity (hasil verify JWT)
		if role := c.Locals("role"); role != nil {
			c.Request().Header.Set("X-User-Role", role.(string))
		}

		if scope := c.Locals("scope"); scope != nil {
			c.Request().Header.Set("X-User-Scope", strings.Join(scope.([]string), ","))
		}

		// trace id
		c.Request().Header.Set("X-Request-Id", uuid.NewString())

		return proxy.Do(c, target)
	}
}
