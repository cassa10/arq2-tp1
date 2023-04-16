package usecase

import (
	"context"
	"github.com/cassa10/arq2-tp1/src/domain/action/command"
	"github.com/cassa10/arq2-tp1/src/domain/action/query"
	"github.com/cassa10/arq2-tp1/src/domain/mock"
	"github.com/cassa10/arq2-tp1/src/domain/model"
	"github.com/cassa10/arq2-tp1/src/domain/model/exception"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_GivenAConfirmedOrderAndDeliveredOrderUseCase_WhenDo_ThenReturnNoErrorAndOrderIsDelivereded(t *testing.T) {
	deliveredOrderUseCase, mocks := setUpDeliveredOrderUseCase(t)
	ctx := context.Background()
	orderId := int64(9)
	order := &model.Order{
		Id:    orderId,
		State: model.ConfirmedOrderState{},
	}

	orderRepo := *order
	orderRepo.Delivered()
	mocks.OrderRepo.EXPECT().FindById(ctx, orderId).Return(order, nil)
	mocks.OrderRepo.EXPECT().Update(ctx, orderRepo).Return(true, nil)

	err := deliveredOrderUseCase.Do(ctx, orderId)

	assert.NoError(t, err)
	assert.Equal(t, model.DeliveredOrderState{}, order.State)
}

func Test_GivenANoConfirmedOrderAndDeliveredOrderUseCase_WhenDo_ThenReturnErrorInvalidTransitionStateAndNoMutateOrder(t *testing.T) {
	deliveredOrderUseCase, mocks := setUpDeliveredOrderUseCase(t)
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

	copyDeliverededOrder := *pendingOrder
	copyDeliveredOrder := *deliveredOrder
	mocks.OrderRepo.EXPECT().FindById(ctx, idPendingOrder).Return(pendingOrder, nil)
	mocks.OrderRepo.EXPECT().FindById(ctx, idDeliveredOrder).Return(deliveredOrder, nil)

	err1 := deliveredOrderUseCase.Do(ctx, idPendingOrder)
	err2 := deliveredOrderUseCase.Do(ctx, idDeliveredOrder)

	assert.ErrorIs(t, err1, exception.OrderInvalidTransitionState{Id: idPendingOrder})
	assert.ErrorIs(t, err2, exception.OrderInvalidTransitionState{Id: idDeliveredOrder})
	assert.Equal(t, &copyDeliverededOrder, pendingOrder)
	assert.Equal(t, &copyDeliveredOrder, deliveredOrder)
}

func Test_GivenDeliveredOrderUseCaseAndAConfirmedOrderAndOrderRepoFindByIdError_WhenDo_ThenReturnThatError(t *testing.T) {
	deliveredOrderUseCase, mocks := setUpDeliveredOrderUseCase(t)
	ctx := context.Background()
	orderId := int64(9)

	mocks.OrderRepo.EXPECT().FindById(ctx, orderId).Return(nil, exception.OrderNotFound{Id: orderId})

	err := deliveredOrderUseCase.Do(ctx, orderId)

	assert.ErrorIs(t, err, exception.OrderNotFound{Id: orderId})
}

func Test_GivenDeliveredOrderUseCaseAndAConfirmedOrderAndOrderRepoUpdateError_WhenDo_ThenReturnThatError(t *testing.T) {
	deliveredOrderUseCase, mocks := setUpDeliveredOrderUseCase(t)
	ctx := context.Background()
	orderId := int64(9)
	order := &model.Order{
		Id:    orderId,
		State: model.ConfirmedOrderState{},
	}

	orderRepo := *order
	orderRepo.Delivered()
	mocks.OrderRepo.EXPECT().FindById(ctx, orderId).Return(order, nil)
	mocks.OrderRepo.EXPECT().Update(ctx, orderRepo).Return(false, exception.OrderCannotUpdate{Id: orderId})

	err := deliveredOrderUseCase.Do(ctx, orderId)

	assert.ErrorIs(t, err, exception.OrderCannotUpdate{Id: orderId})
}

func setUpDeliveredOrderUseCase(t *testing.T) (*DeliveredOrder, *mock.InterfaceMocks) {
	mocks := mock.NewInterfaceMocks(t)
	deliveredOrderCmd := *command.NewDeliveredOrder(mocks.OrderRepo)
	findOrderByIdQuery := *query.NewFindOrderById(mocks.OrderRepo)
	return NewDeliveredOrder(mocks.Logger, deliveredOrderCmd, findOrderByIdQuery), mocks
}
