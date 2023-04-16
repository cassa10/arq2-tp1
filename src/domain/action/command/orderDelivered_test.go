package command

import (
	"context"
	"github.com/cassa10/arq2-tp1/src/domain/mock"
	"github.com/cassa10/arq2-tp1/src/domain/model"
	"github.com/cassa10/arq2-tp1/src/domain/model/exception"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_GivenAConfirmedOrderAndDeliveredOrderCmd_WhenDo_ThenReturnNoErrorAndOrderIsDelivered(t *testing.T) {
	deliveredOrderCmd, mocks := setUpDeliveredOrderCmd(t)
	ctx := context.Background()
	order := &model.Order{
		Id:    int64(4),
		State: model.ConfirmedOrderState{},
	}

	orderRepo := *order
	orderRepo.Delivered()
	mocks.OrderRepo.EXPECT().Update(ctx, orderRepo).Return(true, nil)

	err := deliveredOrderCmd.Do(ctx, order)

	assert.NoError(t, err)
	assert.Equal(t, model.DeliveredOrderState{}, order.State)
}

func Test_GivenANoConfirmedOrderAndDeliveredOrderCmd_WhenDo_ThenReturnInvalidTransitionStateErrorAndOrderNotMutated(t *testing.T) {
	deliveredOrderCmd, _ := setUpDeliveredOrderCmd(t)
	ctx := context.Background()
	idPendingOrder := int64(4)
	pendingOrder := &model.Order{
		Id:    idPendingOrder,
		State: model.PendingOrderState{},
	}
	idDeliveredOrder := int64(6)
	deliveredOrder := &model.Order{
		Id:    idDeliveredOrder,
		State: model.DeliveredOrderState{},
	}
	copyPendingOrder := *pendingOrder
	copyDeliveredOrder := *deliveredOrder

	err1 := deliveredOrderCmd.Do(ctx, pendingOrder)
	err2 := deliveredOrderCmd.Do(ctx, deliveredOrder)

	assert.ErrorIs(t, err1, exception.OrderInvalidTransitionState{Id: idPendingOrder})
	assert.ErrorIs(t, err2, exception.OrderInvalidTransitionState{Id: idDeliveredOrder})
	assert.Equal(t, &copyPendingOrder, pendingOrder)
	assert.Equal(t, &copyDeliveredOrder, deliveredOrder)
}

func Test_GivenAConfirmedOrderAndDeliveredOrderCmdAndOrderRepoUpdateWithError_WhenDo_ThenReturnThatError(t *testing.T) {
	deliveredOrderCmd, mocks := setUpDeliveredOrderCmd(t)
	ctx := context.Background()
	order := &model.Order{
		Id:    int64(4),
		State: model.ConfirmedOrderState{},
	}

	orderRepo := *order
	orderRepo.Delivered()
	mocks.OrderRepo.EXPECT().Update(ctx, orderRepo).Return(false, exception.OrderCannotUpdate{Id: order.Id})

	err := deliveredOrderCmd.Do(ctx, order)

	assert.ErrorIs(t, err, exception.OrderCannotUpdate{Id: order.Id})
}

func setUpDeliveredOrderCmd(t *testing.T) (*DeliveredOrder, *mock.InterfaceMocks) {
	mocks := mock.NewInterfaceMocks(t)
	return NewDeliveredOrder(mocks.OrderRepo), mocks
}
