package domain

import "time"

type Offer struct {
	ID            string    `json:"id" db:"id"`
	MerchantID    string    `json:"merchant_id" db:"merchant_id"`
	Type          string    `json:"type" db:"type"`
	Title         string    `json:"title" db:"title"`
	Description   string    `json:"description" db:"description"`
	Discount      int       `json:"discount" db:"discount"`
	MinimumCost   int       `json:"minimum_cost" db:"minimum_cost"`
	MaxDiscount   int       `json:"max_discount" db:"max_discount"`
	Category      string    `json:"category" db:"category"`
	PaymentMethod string    `json:"payment_method" db:"payment_method"`
	IsActive      bool      `json:"is_active" db:"is_active"`
	CreatedAt     time.Time `json:"created_at" db:"created_at"`
	StartsAt      time.Time `json:"starts_at" db:"starts_at"`
	EndsAt        time.Time `json:"ends_at" db:"ends_at"`
}

type Merchant struct {
	ID        string    `json:"id" db:"id"`
	Name      string    `json:"name" db:"name"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
}

type UserSegment struct {
	ID   string `json:"id" db:"id"`
	Name string `json:"name" db:"name"`
}

type CartItem struct {
	Sku      string `json:"sku"`
	Category string `json:"category"`
	Price    int    `json:"price"`
	Quantity int    `json:"quantity"`
}

type Cart struct {
	Items []CartItem `json:"items"`
}

type CalculationRequest struct {
	MerchantID    string `json:"merchant_id"`
	UserID        string `json:"user_id"`
	PaymentMethod string `json:"payment_method"`
	Cart          Cart   `json:"cart"`
}

type OfferResult struct {
	OfferID     string `json:"offer_id"`
	Title       string `json:"title"`
	Discount    int    `json:"discount"`
	FinalAmount int    `json:"final_amount"`
}

type CalculationResponse struct {
	OriginalAmount   int           `json:"original_amount"`
	ApplicableOffers []OfferResult `json:"applicable_offers"`
	BestOfferID      string        `json:"best_offer_id"`
}

type OfferIncompatibility struct {
	OfferID1 string `json:"offer_id_1" db:"offer_id_1"`
	OfferID2 string `json:"offer_id_2" db:"offer_id_2"`
}
