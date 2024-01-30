package usecase_test

import (
	"context"
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/lp9087/go_otello_dashboard_api/internal/entity"
	"github.com/lp9087/go_otello_dashboard_api/internal/usecase"
	"github.com/stretchr/testify/require"
	"testing"
)

var errInternalServErr = errors.New("internal server error")

type test struct {
	name string
	mock func()
	res  interface{}
	err  error
}

func loyalHotelsUseCase(t *testing.T) (*usecase.MostLoyalUseCase, *MockMostLoyalHotelsRepo) {
	t.Helper()

	mockCtl := gomock.NewController(t)
	defer mockCtl.Finish()

	repo := NewMockMostLoyalHotelsRepo(mockCtl)

	loyalHotelsUseCase := usecase.NewMostLoyalHotelsUseCase(repo)

	return loyalHotelsUseCase, repo
}

func TestMostLoyalHotels(t *testing.T) {
	t.Parallel()

	loyalHotelsUseCase, repo := loyalHotelsUseCase(t)
	mockRes := entity.MostLoyalHotels{ID: "1", HotelName: "TestHotel", ExternalID: "1010", RoomTypes: 5, Rates: 6, TotalAmount: 255}

	tests := []test{
		{
			name: "empty result",
			mock: func() {
				repo.EXPECT().Store(context.Background()).Return(nil, nil)
			},
			res: []entity.MostLoyalHotels(nil),
			err: nil,
		},
		{
			name: "result with error",
			mock: func() {
				repo.EXPECT().Store(context.Background()).Return(nil, errInternalServErr)
			},
			res: []entity.MostLoyalHotels{},
			err: errInternalServErr,
		},
		{
			name: "result with one hotel",
			mock: func() {
				repo.EXPECT().Store(context.Background()).Return([]entity.MostLoyalHotels{mockRes}, nil)
			},
			res: []entity.MostLoyalHotels{mockRes},
			err: nil,
		},
	}

	for _, tc := range tests {
		tc := tc

		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			tc.mock()

			res, err := loyalHotelsUseCase.Get(context.Background())

			require.Equal(t, res, tc.res)
			require.ErrorIs(t, err, tc.err)
		})
	}
}
