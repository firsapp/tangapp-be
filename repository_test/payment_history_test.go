package repository

import (
	"context"
	"fmt"
	"tangapp-be/repository"
	"tangapp-be/utils"
	"testing"

	"github.com/stretchr/testify/require"
)

func GeneratePaymentHistory(t *testing.T, fromUser repository.User, toUser repository.User, event repository.Event) repository.PaymentHistory {
	eventMember := GenerateMemberDetail(t, fromUser, event)
	arg := repository.AddPaymentHistoryParams{
		EventMemberDetailsID: eventMember.ID,
		FromUserID:           fromUser.ID,
		ToUserID:             toUser.ID,
		Nominal:              utils.RandomMoney(),
		Description:          NullString(utils.RandomString(10)),
	}

	payload, err := testQueries.AddPaymentHistory(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, payload)

	require.Equal(t, arg.EventMemberDetailsID, payload.EventMemberDetailsID)
	require.Equal(t, arg.FromUserID, payload.FromUserID)
	require.Equal(t, arg.ToUserID, payload.ToUserID)
	require.Equal(t, arg.Nominal, payload.Nominal)
	require.Equal(t, arg.Description, payload.Description)

	return payload
}
func TestAddPaymentHistory(t *testing.T) {
	fromUser := GenerateUser(t)
	toUser := GenerateUser(t)
	event := GenerateEvent(t, toUser)

	GeneratePaymentHistory(t, fromUser, toUser, event)

}

func TestListPaymentHistoryByEvent(t *testing.T) {
	toUser := GenerateUser(t)
	event := GenerateEvent(t, toUser)

	for i := 0; i < 5; i++ {
		fromUser := GenerateUser(t)
		ph := GeneratePaymentHistory(t, fromUser, toUser, event)

		payload, err := testQueries.ListPaymentHistoryByEvent(context.Background(), event.ID)

		require.NoError(t, err)
		require.NotEmpty(t, payload)

		require.Equal(t, ph.FromUserID, payload[i].FromUserID)
		require.Equal(t, ph.ToUserID, payload[i].ToUserID)
		require.Equal(t, ph.Nominal, payload[i].Nominal)
		require.Equal(t, ph.Description, payload[i].Description)

	}

}

func TestListPaymentHistoryByUser(t *testing.T) {
	toUser := GenerateUser(t)
	event := GenerateEvent(t, toUser)

	for i := 0; i < 5; i++ {
		fromUser := GenerateUser(t)
		ph := GeneratePaymentHistory(t, fromUser, toUser, event)
		arg := repository.ListPaymentHistoryByUserParams{
			ToUserID: toUser.ID,
		}

		payload, err := testQueries.ListPaymentHistoryByUser(context.Background(), arg)

		require.NoError(t, err)
		require.NotEmpty(t, payload)

		require.Equal(t, ph.FromUserID, payload[i].FromUserID)
		require.Equal(t, ph.ToUserID, payload[i].ToUserID)
		require.Equal(t, ph.Nominal, payload[i].Nominal)
		require.Equal(t, ph.Description, payload[i].Description)
		fmt.Println(payload[i])
	}

}

func TestListPaymentHistoryByUser2(t *testing.T) {
	fromUser := GenerateUser(t)

	for i := 0; i < 5; i++ {
		toUser := GenerateUser(t)
		event := GenerateEvent(t, toUser)
		ph := GeneratePaymentHistory(t, fromUser, toUser, event)
		arg := repository.ListPaymentHistoryByUserParams{
			FromUserID: fromUser.ID,
		}

		payload, err := testQueries.ListPaymentHistoryByUser(context.Background(), arg)

		require.NoError(t, err)
		require.NotEmpty(t, payload)

		require.Equal(t, ph.FromUserID, payload[i].FromUserID)
		require.Equal(t, ph.ToUserID, payload[i].ToUserID)
		require.Equal(t, ph.Nominal, payload[i].Nominal)
		require.Equal(t, ph.Description, payload[i].Description)
		fmt.Println(payload[i])
	}

}

func TestUpdatePaymentHistory(t *testing.T) {
	fromUser := GenerateUser(t)
	toUser := GenerateUser(t)
	toUser2 := GenerateUser(t)
	event := GenerateEvent(t, toUser)
	ph := GeneratePaymentHistory(t, fromUser, toUser, event)
	arg := repository.UpdatePaymentHistoryParams{
		ID:                   ph.ID,
		EventMemberDetailsID: ph.EventMemberDetailsID,
		ToUserID:             toUser2.ID,
		Nominal:              utils.RandomMoney(),
		Description:          NullString(utils.RandomString(10)),
	}

	payload, err := testQueries.UpdatePaymentHistory(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, payload)

	require.Equal(t, arg.ID, payload.ID)
	require.Equal(t, arg.EventMemberDetailsID, payload.EventMemberDetailsID)
	require.Equal(t, arg.ToUserID, payload.ToUserID)
	require.Equal(t, arg.Nominal, payload.Nominal)
	require.Equal(t, arg.Description, payload.Description)

}
