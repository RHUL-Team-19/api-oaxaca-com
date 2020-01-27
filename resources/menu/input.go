package menu

// type createInput contains fields sent via a HTTP POST request to /menu
// which are used to create a new meal record in the database.
type createInput struct {
  Name string `json:"name"`
  Price float64 `json:"price"`
  Description string `json:"description"`
  IsVegan bool `json:"is_vegan"`
  IsVegetarian bool `json:"is_vegetarian"`
  DoesContainEgg bool `json:"does_contain_egg"`
  DoesContainSoy bool `json:"does_contain_soy"`
  DoesContainFish bool `json:"does_contain_fish"`
  DoesContainLactose bool `json:"does_contain_lactose"`
  DoesContainWheat bool `json:"does_contain_wheat"`
  DoesContainNuts bool `json:"does_contain_nuts"`
  DoesContainGluten bool `json:"does_contain_gluten"`
  DoesContainDairy bool `json:"does_contain_dairy"`
  ImageURL string `json:"image_url"`
}

// IsValid ensures the fields provided in createInput match a set of rules. If
// validation fails, a reason is returned - which should be sent to the client.
func (i *createInput) IsValid() (bool, string) {

  // ...

  return true, ""
}

// type updateInput contains fields sent via a HTTP PUT/PATCH request to
// /menu/{id} which are used to fully update or partially update a menu record
// in the database.
type updateInput struct {
  Name *string `json:"name"`
  Price *float64 `json:"price"`
  Description *string `json:"description"`
  IsVegan *bool `json:"is_vegan"`
  IsVegetarian *bool `json:"is_vegetarian"`
  DoesContainEgg *bool `json:"does_contain_egg"`
  DoesContainSoy *bool `json:"does_contain_soy"`
  DoesContainFish *bool `json:"does_contain_fish"`
  DoesContainLactose *bool `json:"does_contain_lactose"`
  DoesContainWheat *bool `json:"does_contain_wheat"`
  DoesContainNuts *bool `json:"does_contain_nuts"`
  DoesContainGluten *bool `json:"does_contain_gluten"`
  DoesContainDairy *bool `json:"does_contain_dairy"`
  ImageURL *string `json:"image_url"`
}

// IsValid ensures the fields provided in updateInput match a set of rules. If
// validation fails, a reason is returned - which should be sent to the client.
func (i *updateInput) IsValid() (bool, string) {

  // ...

  return true, ""
}
