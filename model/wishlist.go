package model

import "log"

// Wishlist is data structure used for Emarsys payload
type Wishlist struct {
	Title string  `json:"title"`
	Link  string  `json:"link"`
	Image string  `json:"image"`
	Msrp  float32 `json:"msrp"`
	Price float32 `json:"price"`
}

// AllWishlist : gets all product loved list
func (db *DB) AllWishlist() ([]*Wishlist, error) {
	// AllWishlistQuery() -> it's a function return select query syntax
	rows, err := db.Query(AllWishlistQuery())
	if err != nil {
		log.Println(err.Error())
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
