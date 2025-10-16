package albums

type AlbumService struct{}

type AlbumController struct {
	service *AlbumService
}

type Album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}