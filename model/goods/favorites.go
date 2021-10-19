package goods

type Favorites struct {
	FavoritesID    int64                  `json:"favorites_id"`
	FavoritesTitle string                 `json:"favorites_title"`
	Data           map[string]interface{} `json:"data,omitempty"`
}

type PddCats struct {
	CatName string `json:"cat_name"`
	CatID   int64  `json:"cat_id"`
}
