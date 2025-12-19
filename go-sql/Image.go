package gosql

import (
	_ "github.com/mattn/go-sqlite3"
)

type Image struct {
	ID        int
	ImageName string
	ImageData []byte
	CreatedAt string
}

func (data *Data) AddImage(img Image) int {
	res, err := data.DB.Exec("INSERT INTO images (image_name, image_data) VALUES (?, ?)",
		img.ImageName, img.ImageData)
	if err != nil {
		return 0
	}
	id, err := res.LastInsertId()
	if err == nil {
		return int(id)
	}
	return 0
}

func (data *Data) GetAllImages() ([]Image, error) {
	rows, err := data.DB.Query("SELECT id, image_name, image_data, created_at FROM images")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var images []Image
	for rows.Next() {
		var img Image
		err := rows.Scan(&img.ID, &img.ImageName, &img.ImageData, &img.CreatedAt)
		if err != nil {
			return nil, err
		}
		images = append(images, img)
	}
	return images, nil
}

func (data *Data) GetImageByID(id int) (*Image, error) {
	row := data.DB.QueryRow("SELECT id, image_name, image_data, created_at FROM images WHERE id = ?", id)
	var img Image
	err := row.Scan(&img.ID, &img.ImageName, &img.ImageData, &img.CreatedAt)
	if err != nil {
		return nil, err
	}
	return &img, nil
}

func (data *Data) UpdateImage(img Image) error {
	_, err := data.DB.Exec("UPDATE images SET image_name = ?, image_data = ? WHERE id = ?",
		img.ImageName, img.ImageData, img.ID)
	return err
}

func (data *Data) DeleteImage(id int) error {
	_, err := data.DB.Exec("DELETE FROM images WHERE id = ?", id)
	return err
}
