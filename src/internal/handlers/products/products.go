package products

import (
	"net/http"
	"reflect"
	"strings"

	"github.com/benjaspet/generate-s25-software-challenge/src/internal/utils"
	"github.com/gofiber/fiber/v2"
)

func (s *Service) GetProducts(c *fiber.Ctx) error {

	sort := c.Query("sort", "name")
	order := c.Query("order", "asc")
	categories := c.Query("categories")
	offset := c.QueryInt("offset", 0)
	limit := c.QueryInt("limit", 3)
	price_min := c.QueryInt("price_min", 0)
	price_max := c.QueryInt("price_max", 4294967295)
	star_min := c.QueryInt("star_min", 0)
	star_max := c.QueryInt("star_max", 500)

	var result, err = utils.GetMyProducts()

	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	if sort != "name" && sort != "price" && sort != "stars" {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid sort field"})
	}

	if order != "asc" && order != "desc" {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid order field"})
	}

	result = utils.SortProductsByKey(result, sort)

	shouldBeReversed := order == "desc"

	if shouldBeReversed {
		result = utils.ReverseProducts(result)
	}

	VALID_CATEGORIES := []string{"electronics", "apparel", "home goods", "sports", "beauty", "grocery", "office supplies", "outdoor", "toys", "health", "automotive", "luxury", "books"}

	if categories != "" {
		categoriesArr := strings.Split(strings.ToLower(categories), ",")
		if !utils.ContainsAll(VALID_CATEGORIES, categoriesArr) {
			return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid category"})
		}
		result = utils.FilterProductsByCategories(result, categoriesArr)
	}

	if price_max < price_min {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid price range"})
	} else if reflect.TypeOf(price_min).Kind() != reflect.Int || reflect.TypeOf(price_max).Kind() != reflect.Int {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid price type"})
	}

	result = utils.FilterProductsByPriceRange(result, price_min, price_max)

	if star_max < star_min {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid star range"})
	} else if reflect.TypeOf(star_min).Kind() != reflect.Int || reflect.TypeOf(star_max).Kind() != reflect.Int {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid star type"})
	}

	result = utils.FilterProductsByStarRange(result, star_min, star_max)

	result = utils.ApplyOffsetAndLimit(result, offset, limit)

	return c.Status(http.StatusOK).JSON(result)
}