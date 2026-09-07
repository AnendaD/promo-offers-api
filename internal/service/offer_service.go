package service

import (
	"context"
	"time"

	"promo-offers-api/internal/domain"
	"promo-offers-api/internal/repository"
)

type OfferService struct {
	offerRepo *repository.OfferRepository
}

func NewOfferService(offerRepo *repository.OfferRepository) *OfferService {
	return &OfferService{offerRepo: offerRepo}
}

func (s *OfferService) CreateOffer(ctx context.Context, offer *domain.Offer) error {
	return s.offerRepo.CreateOffer(ctx, offer)
}

func (s *OfferService) GetOffersByMerchant(ctx context.Context, merchantID string) ([]domain.Offer, error) {
	return s.offerRepo.GetOffersByMerchant(ctx, merchantID)
}

func (s *OfferService) Calculate(ctx context.Context, req domain.CalculationRequest) (*domain.CalculationResponse, error) {
	offers, err := s.offerRepo.GetOffersByMerchant(ctx, req.MerchantID)
	if err != nil {
		return nil, err
	}

	totalAmount := 0
	categoryAmounts := make(map[string]int)
	for _, item := range req.Cart.Items {
		amount := item.Price * item.Quantity
		totalAmount += amount
		categoryAmounts[item.Category] += amount
	}

	applicableOffers := []domain.OfferResult{}
	applicableOfferIDs := []string{}

	for _, offer := range offers {
		now := time.Now()
		if now.Before(offer.StartsAt) || now.After(offer.EndsAt) {
			continue
		}

		if totalAmount < offer.MinimumCost {
			continue
		}

		if offer.PaymentMethod != "" && offer.PaymentMethod != req.PaymentMethod {
			continue
		}

		if offer.Category != "" {
			if _, exists := categoryAmounts[offer.Category]; !exists {
				continue
			}
		}

		discount := 0
		var amountToDiscount int

		if offer.Category != "" {
			amountToDiscount = categoryAmounts[offer.Category]
		} else {
			amountToDiscount = totalAmount
		}

		if offer.Type == "percentage" {
			discount = amountToDiscount * offer.Discount / 100
			if offer.MaxDiscount > 0 && discount > offer.MaxDiscount {
				discount = offer.MaxDiscount
			}
		} else {
			discount = offer.Discount
		}

		finalAmount := totalAmount - discount
		if finalAmount < 0 {
			finalAmount = 0
		}

		applicableOffers = append(applicableOffers, domain.OfferResult{
			OfferID:     offer.ID,
			Title:       offer.Title,
			Discount:    discount,
			FinalAmount: finalAmount,
		})
		applicableOfferIDs = append(applicableOfferIDs, offer.ID)
	}

	if len(applicableOfferIDs) > 1 {
		incompatibilities, err := s.offerRepo.GetIncompatibilities(ctx, applicableOfferIDs)
		if err != nil {
			return nil, err
		}

		filteredOffers := []domain.OfferResult{}
		for i, offer := range applicableOffers {
			compatible := true
			for j, other := range applicableOffers {
				if i != j {
					if incompatible, exists := incompatibilities[offer.OfferID]; exists {
						for _, inc := range incompatible {
							if inc == other.OfferID {
								compatible = false
								break
							}
						}
					}
				}
				if !compatible {
					break
				}
			}
			if compatible {
				filteredOffers = append(filteredOffers, offer)
			}
		}
		applicableOffers = filteredOffers
	}

	bestOfferID := ""
	if len(applicableOffers) > 0 {
		best := applicableOffers[0]
		for _, offer := range applicableOffers {
			if offer.Discount > best.Discount {
				best = offer
			}
		}
		bestOfferID = best.OfferID
	}

	return &domain.CalculationResponse{
		OriginalAmount:   totalAmount,
		ApplicableOffers: applicableOffers,
		BestOfferID:      bestOfferID,
	}, nil
}
