package goods

type Favorites struct {
	FavoritesID    int64                  `json:"favorites_id"`
	FavoritesTitle string                 `json:"favorites_title"`
	Data           map[string]interface{} `json:"data,omitempty"`
}
