package gosql

import (
	_ "github.com/mattn/go-sqlite3"
)

type Config struct {
	ID        int
	API       string
	VaiTro    string
	HuongDan  string
	NhiemVu   string
	NoiDung   string
	CreatedAt string
}

func (data *Data) AddConfig(c Config) int {
	res, err := data.DB.Exec("INSERT INTO config (api, vai_tro, huong_dan, nhiem_vu, noi_dung) VALUES (?, ?, ?, ?, ?)",
		c.API, c.VaiTro, c.HuongDan, c.NhiemVu, c.NoiDung)
	if err != nil {
		return 0
	}
	id, err := res.LastInsertId()
	if err == nil {
		return int(id)
	}
	return 0
}
func (data *Data) GetAllConfigs() ([]Config, error) {
	rows, err := data.DB.Query("SELECT id, api, vai_tro, huong_dan, nhiem_vu, noi_dung, created_at FROM config")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var configs []Config
	for rows.Next() {
		var c Config
		err := rows.Scan(&c.ID, &c.API, &c.VaiTro, &c.HuongDan, &c.NhiemVu, &c.NoiDung, &c.CreatedAt)
		if err != nil {
			return nil, err
		}
		configs = append(configs, c)
	}
	return configs, nil
}

func (data *Data) GetConfigByID(id int) (*Config, error) {
	row := data.DB.QueryRow("SELECT id, api, vai_tro, huong_dan, nhiem_vu, noi_dung, created_at FROM config WHERE id = ?", id)
	var c Config
	err := row.Scan(&c.ID, &c.API, &c.VaiTro, &c.HuongDan, &c.NhiemVu, &c.NoiDung, &c.CreatedAt)
	if err != nil {
		return nil, err
	}
	return &c, nil
}
func (data *Data) UpdateConfig(c Config) error {
	_, err := data.DB.Exec("UPDATE config SET api = ?, vai_tro = ?, huong_dan = ?, nhiem_vu = ?, noi_dung = ? WHERE id = ?",
		c.API, c.VaiTro, c.HuongDan, c.NhiemVu, c.NoiDung, c.ID)
	return err
}
func (data *Data) DeleteConfig(id int) error {
	_, err := data.DB.Exec("DELETE FROM config WHERE id = ?", id)
	return err
}
