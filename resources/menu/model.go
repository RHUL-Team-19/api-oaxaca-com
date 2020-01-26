package menu

// type menuMeal represents a meal record in the database, and is used by the
// ORM to create the menu meals table.
type menuMeal struct {
  MealID int64 `pg:",pk" json:"meal_id"`
}
