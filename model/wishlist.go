package model

import "fmt"

// Wishlist is structur data for Emarsys
// wishlist
type Wishlist struct {
	Title string  `json:"title"`
	Link  string  `json:"link"`
	Image string  `json:"image"`
	Msrp  float32 `json:"msrp"`
	Price float32 `json:"price"`
}

// AllWishlist : gets all product loved list
func (db *DB) AllWishlist() ([]*Wishlist, error) {
	rows, err := db.Query(`
			SELECT
				xxx
			FROM tableXXX
		`)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}
	defer rows.Close()

	wList := make([]*Wishlist, 0)
	for rows.Next() {
		w := new(Wishlist)
		err := rows.Scan(&w.Title, &w.Link, &w.Image, &w.Msrp, &w.Price)
		if err != nil {
			return nil, err
		}
		wList = append(wList, w)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}
	return wList, nil
}
