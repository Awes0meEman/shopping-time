package models

type Game struct {
	Id			int		`json:"id"`
	Title 		string 	`json:"title"`
	Description string 	`json:"description"`
	Currencies	[]Currency
	Shops 		[]Shop
}

type Currency struct {
	Id 			int 	`json:"id"`
	GameId		int 	`json:"gameid"`
	Name 		string 	`json:"name"`
	Description string 	`json:"description"`
}

type Shop struct {
	Id			int		`json:"id"`
	GameId		int 	`json:"gameid"`
	Title 		string 	`json:"title"`
	Description string 	`json:"description"`
	Items		[]ItemStock
}

type ItemStock struct {
	Id 			int 	`json:"id"`
	ItemId		int 	`json:"itemid"`
	StockCount	int		`json:"stockcount"`
}

type Item struct {
	Id			int	 	`json:"id"`
	ShopId 		int 	`json:"shopid"`
	CurrencyId 	int 	`json:"currencyid"`
	Name 		string 	`json:"name"`
	Description string 	`json:"description"`
	BaseCost 	int 	`json:"cost"`
	StockLimit	bool	`json:"stocklimit"`
	IsRurual	bool	`json:"isrural"`
	IsUrban 	bool	`json:"isurban"`
	IsPremium 	bool	`json:"ispremium"`
	StockCount	int 	`json:"stockcount"`
}
