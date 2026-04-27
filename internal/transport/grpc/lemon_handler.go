package grpchandler

import (
	"context"

	lemonv1 "github.com/fruty-gw/lemon/gen/go/api/lemon/v1"
)

type lemonService interface {
	GetLemon()
	SqueezeLemon()
	AddLemon()
	WashLemons()
}

type LemonHandler struct {
	lemonv1.UnimplementedLemonServiceServer
	svc lemonService
}

func NewLemonHandler(svc lemonService) *LemonHandler {
	return &LemonHandler{svc: svc}
}

func (h *LemonHandler) GetLemon(
	ctx context.Context,
	in *lemonv1.GetLemonRequest,
) (*lemonv1.GetLemonResponse, error) {
	panic("implement me")
}

func (h *LemonHandler) SqueezeLemon(
	ctx context.Context,
	in *lemonv1.SqueezeLemonRequest,
) (*lemonv1.SqueezeLemonResponse, error) {
	panic("implement me")
}

func (h *LemonHandler) AddLemon(
	ctx context.Context,
	in *lemonv1.AddLemonRequest,
) (*lemonv1.AddLemonResponse, error) {
	panic("implement me")
}

func (h *LemonHandler) ListLemons(
	ctx context.Context,
	in *lemonv1.ListLemonsRequest,
) (*lemonv1.ListLemonsResponse, error) {
	panic("implement me")
}
