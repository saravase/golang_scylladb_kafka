package productstore

import (
	"time"
)

type ProductStatus int

const (
	ProductStatusUnknown ProductStatus = iota
	ProductStatusPending
	ProductStatusShipped
	ProductStatusCompleted
	ProductStatusCancelled
)

func (s ProductStatus) String() string {
	switch s {
	case ProductStatusPending:
		return "pending"
	case ProductStatusShipped:
		return "shipped"
	case ProductStatusCompleted:
		return "completed"
	case ProductStatusCancelled:
		return "cancelled"
	default:
		return ""
	}
}

func ParseProductStatus(s string) ProductStatus {
	switch s {
	case "pending":
		return ProductStatusPending
	case "shipped":
		return ProductStatusShipped
	case "completed":
		return ProductStatusCompleted
	case "cancelled":
		return ProductStatusCancelled
	default:
		return ProductStatusUnknown
	}
}

type Product struct {
	ProductID   string
	UserID      string
	ResolverID  string
	ReviewerIDs []string
	CreatedAt   time.Time
	UpdateAt    time.Time
	Status      ProductStatus
	Title       string
}

type Message struct {
	ProductID string
	MessageID string
	Body      string
	CreatedAt time.Time
}
