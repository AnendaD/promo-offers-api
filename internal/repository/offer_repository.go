package repository

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

	"promo-offers-api/internal/domain"
)

type OfferRepository struct {
	db *sql.DB
}

func NewOfferRepository(db *sql.DB) *OfferRepository {
	return &OfferRepository{db: db}
}

func (r *OfferRepository) CreateOffer(ctx context.Context, offer *domain.Offer) error {
	query := `
        INSERT INTO offers (
            id, merchant_id, type, title, description, discount,
            minimum_cost, max_discount, category, payment_method,
            is_active, starts_at, ends_at, created_at
        ) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14)
    `

	_, err := r.db.ExecContext(
		ctx, query,
		offer.ID, offer.MerchantID, offer.Type, offer.Title,
		offer.Description, offer.Discount, offer.MinimumCost,
		offer.MaxDiscount, offer.Category, offer.PaymentMethod,
		offer.IsActive, offer.StartsAt, offer.EndsAt, offer.CreatedAt,
	)
	return err
}

func (r *OfferRepository) GetOffersByMerchant(ctx context.Context, merchantID string) ([]domain.Offer, error) {
	query := `
        SELECT id, merchant_id, type, title, description, discount,
               minimum_cost, max_discount, category, payment_method,
               is_active, created_at, starts_at, ends_at
        FROM offers
        WHERE merchant_id = $1 AND is_active = true
        ORDER BY created_at DESC
    `

	rows, err := r.db.QueryContext(ctx, query, merchantID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var offers []domain.Offer
	for rows.Next() {
		var o domain.Offer
		err := rows.Scan(
			&o.ID, &o.MerchantID, &o.Type, &o.Title, &o.Description,
			&o.Discount, &o.MinimumCost, &o.MaxDiscount, &o.Category,
			&o.PaymentMethod, &o.IsActive, &o.CreatedAt, &o.StartsAt, &o.EndsAt,
		)
		if err != nil {
			return nil, err
		}
		offers = append(offers, o)
	}

	return offers, nil
}

func (r *OfferRepository) GetIncompatibilities(ctx context.Context, offerIDs []string) (map[string][]string, error) {
	if len(offerIDs) == 0 {
		return nil, nil
	}

	placeholders := make([]string, len(offerIDs))
	args := make([]interface{}, len(offerIDs))
	for i, id := range offerIDs {
		placeholders[i] = fmt.Sprintf("$%d", i+1)
		args[i] = id
	}

	query := fmt.Sprintf(`
        SELECT offer_id_1, offer_id_2 
        FROM offer_incompatibilities 
        WHERE offer_id_1 IN (%s) OR offer_id_2 IN (%s)
    `, strings.Join(placeholders, ","), strings.Join(placeholders, ","))

	rows, err := r.db.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	incompatibilities := make(map[string][]string)
	for rows.Next() {
		var id1, id2 string
		if err := rows.Scan(&id1, &id2); err != nil {
			return nil, err
		}
		incompatibilities[id1] = append(incompatibilities[id1], id2)
		incompatibilities[id2] = append(incompatibilities[id2], id1)
	}

	return incompatibilities, nil
}
