package routes

import (
	"echo-mysql-default/internal/domain"
	"echo-mysql-default/internal/util"
	"github.com/doug-martin/goqu/v9"
	"github.com/labstack/echo/v4"
)

func ProductsRoutes(e *echo.Echo, db *goqu.Database) {
	binder := echo.DefaultBinder{}

	e.GET("/products/", func(c echo.Context) error {
		var products = make([]domain.Product, 0)

		err := db.
			Select(
				"id",
				"name",
				"description",
				"price",
				"online",
				"created_at",
				"modified_at",
			).
			From("products").
			ScanStructs(&products)

		if err != nil {
			e.Logger.Error(util.GetStackTrace(err))
			return err
		}

		return c.JSON(200, products)
	})

	e.POST("/products/", func(c echo.Context) error {
		var productPayload = new(domain.Product)

		err := binder.BindBody(c, &productPayload)

		if err != nil {
			e.Logger.Error(util.GetStackTrace(err))
			return err
		}

		rows, err := db.Query(
			goqu.L(
				"call `createProduct`(?,?,?)",
			).Literal(),
			productPayload.Name,
			productPayload.Description,
			productPayload.Price,
		)

		if err != nil {
			e.Logger.Error(util.GetStackTrace(err))
			return err
		}

		var product = new(domain.Product)

		for rows.Next() {
			err = rows.Scan(
				&product.Id,
				&product.Name,
				&product.Description,
				&product.Price,
				&product.Online,
				&product.CreatedAt,
				&product.ModifiedAt,
			)
			if err != nil {
				e.Logger.Error(util.GetStackTrace(err))
				return err
			}
		}

		return c.JSON(200, product)
	})

	e.PUT("/products/:id/", func(c echo.Context) error {
		id := c.Param("id")

		var productPayload = new(domain.Product)
		err := binder.BindBody(c, &productPayload)

		if err != nil {
			e.Logger.Error(util.GetStackTrace(err))
			return err
		}

		query, _, err := db.
			Update("products").
			Where(goqu.Ex{
				"id": id,
			}).
			Set(
				goqu.Record{
					"name":        productPayload.Name,
					"description": productPayload.Description,
					"price":       productPayload.Price,
					"online":      productPayload.Online,
					"modified_at": goqu.L("now()"),
				},
			).ToSQL()

		if err != nil {
			e.Logger.Error(util.GetStackTrace(err))
			return err
		}

		result, err := db.Exec(query)

		if err != nil {
			e.Logger.Error(util.GetStackTrace(err))
			return err
		}

		affected, err := result.RowsAffected()

		if err != nil {
			e.Logger.Error(util.GetStackTrace(err))
			return err
		}

		if affected == 1 {
			return c.NoContent(200)
		}

		return c.NoContent(304)
	})

	e.DELETE("/products/:id/", func(c echo.Context) error {
		id := c.Param("id")

		query, _, err := db.Delete("products").Where(goqu.Ex{"id": id}).ToSQL()

		if err != nil {
			e.Logger.Error(util.GetStackTrace(err))
			return err
		}

		result, err := db.Exec(query)

		if err != nil {
			e.Logger.Error(util.GetStackTrace(err))
			return err
		}

		affected, err := result.RowsAffected()

		if err != nil {
			e.Logger.Error(util.GetStackTrace(err))
			return err
		}

		if affected == 1 {
			return c.NoContent(200)
		}

		return c.NoContent(304)
	})

}
