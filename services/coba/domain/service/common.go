package service

import "context"

// SvcTx template for common transaction pattern
type SvcTx interface {
	BeginTx(ctx context.Context, operation func(ctx context.Context, svc *Wrapper) error) error
}
