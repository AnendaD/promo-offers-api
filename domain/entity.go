package domain

import "time"

type Offer struct {
	ID          string
	Type        string
	Title       string
	Discount    int
	MinimumCost int
	Created_at  time.Time
	StartsAt    time.Time
	EndsAt      time.Time
}

type Merchant struct {
	ID     string
	Offers []Offer
}

type UserSegment struct {
	ID string
}

type Cart struct {
	CartItems []CartItem
}

type CartItem struct {
	Sku      string
	Category string
	Price    int
	Quantity int
}

type CalculationRequest struct {
	MerchantID    string
	UserID        string
	PaymentMethod string
	Cart          Cart
}

type CalculationResponse struct {
	OriginalAmount   int
	ApplicableOffers []struct {
		OfferID     string
		Title       string
		Discount    int
		FinalAmount int
		BestOfferID string
	}
}
