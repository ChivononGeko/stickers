package models

type APIResponse struct {
	Response []Product `json:"response"`
}

type Product struct {
	ID        string            `json:"product_id"`
	Name      string            `json:"product_name"`
	Price     map[string]string `json:"price"`
	Workshop  string            `json:"workshop"`
	Modifiers []ModifierGroup   `json:"group_modifications"`
}

type ModifierGroup struct {
	Name          string     `json:"name"`
	Modifications []Modifier `json:"modifications"`
}

type Modifier struct {
	ID   int    `json:"dish_modification_id"`
	Name string `json:"name"`
}
