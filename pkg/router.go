package pkg

import (
	"cmp"
	"fmt"
	"slices"

	"github.com/gofiber/fiber/v2"
)

func ListRouters(app *fiber.App) {
	routes := app.GetRoutes(true)

	slices.SortFunc(routes, func(a, b fiber.Route) int {
		return cmp.Or(
			cmp.Compare(a.Method, b.Method),
			cmp.Compare(a.Path, b.Path),
		)
	})

	for _, router := range routes {
		if router.Method != "HEAD" {
			fmt.Println(router.Method, " ", router.Path)
		}
	}

}
