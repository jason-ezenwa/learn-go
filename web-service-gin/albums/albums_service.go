package albums

func NewAlbumService() *AlbumService{
	return  &AlbumService{}
}

var albumRecords = []Album {
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Calloway", Artist: "Sarah Vaughan", Price: 39.99},
}

func (service *AlbumService) GetAlbums() []Album {
	return albumRecords
}

func (service *AlbumService) GetAlbumByID(id string) *Album {
	for _, a := range albumRecords {
		if a.ID == id {
			return &a
		}
	}

	return nil
}

func (service *AlbumService) PostAlbum(album Album) Album {
	albumRecords = append(albumRecords, album)
	return album
}