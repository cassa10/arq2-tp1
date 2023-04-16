package command

import (
	"context"
	"github.com/cassa10/arq2-tp1/src/domain/mock"
	"github.com/cassa10/arq2-tp1/src/domain/model"
	"github.com/cassa10/arq2-tp1/src/domain/model/exception"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_GivenAPendingOrderAndConfirmOrderCmd_WhenDo_ThenReturnNoErrorAndOrderIsConfirmed(t *testing.T) {
	confirmOrderCmd, mocks := setUpConfirmOrderCmd(t)
	ctx := context.Background()
	order := &model.Order{
		Id:    int64(4),
		State: model.PendingOrderState{},
	}

	orderRepo := *order
	orderRepo.Confirm()
	mocks.OrderRepo.EXPECT().Update(ctx, orderRepo).Return(true, nil)

	err := confirmOrderCmd.Do(ctx, order)

	assert.NoError(t, err)
	assert.Equal(t, model.ConfirmedOrderState{}, order.State)
}

func Test_GivenANoPendingOrderAndConfirmOrderCmd_WhenDo_ThenReturnInvalidTransitionStateErrorAndOrderNotMutated(t *testing.T) {
	confirmOrderCmd, _ := setUpConfirmOrderCmd(t)
	ctx := context.Background()
	idConfirmedOrder := int64(4)
	confirmedOrder := &model.Order{
		Id:    idConfirmedOrder,
		State: model.ConfirmedOrderState{},
	}
	idDeliveredOrder := int64(6)
	deliveredOrder := &model.Order{
		Id:    idDeliveredOrder,
		State: model.DeliveredOrderState{},
	}
	copyConfirmedOrder := *confirmedOrder
	copyDeliveredOrder := *deliveredOrder

	err1 := confirmOrderCmd.Do(ctx, confirmedOrder)
	err2 := confirmOrderCmd.Do(ctx, deliveredOrder)

	assert.ErrorIs(t, err1, exception.OrderInvalidTransitionState{Id: idConfirmedOrder})
	assert.ErrorIs(t, err2, exception.OrderInvalidTransitionState{Id: idDeliveredOrder})
	assert.Equal(t, &copyConfirmedOrder, confirmedOrder)
	assert.Equal(t, &copyDeliveredOrder, deliveredOrder)
}

func Test_GivenAPendingOrderAndConfirmOrderCmdAndOrderRepoUpdateWithError_WhenDo_ThenReturnThatError(t *testing.T) {
	confirmOrderCmd, mocks := setUpConfirmOrderCmd(t)
	ctx := context.Background()
	order := &model.Order{
		Id:    int64(4),
		State: model.PendingOrderState{},
	}

	orderRepo := *order
	orderRepo.Confirm()
	mocks.OrderRepo.EXPECT().Update(ctx, orderRepo).Return(false, exception.OrderCannotUpdate{Id: order.Id})

	err := confirmOrderCmd.Do(ctx, order)

	assert.ErrorIs(t, err, exception.OrderCannotUpdate{Id: order.Id})
}

func setUpConfirmOrderCmd(t *testing.T) (*ConfirmOrder, *mock.InterfaceMocks) {
	mocks := mock.NewInterfaceMocks(t)
	return NewConfirmOrder(mocks.OrderRepo), mocks
}
