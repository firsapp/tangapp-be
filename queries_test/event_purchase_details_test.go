package queries

import (
	"context"
	"tangapp-be/queries"
	"tangapp-be/utils"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func GeneratePurchaseDetail(t *testing.T, event queries.Event) queries.EventPurchaseDetail {
	arg := queries.AddPurchaseDetailParams{
		EventID:    event.ID,
		Name:       utils.RandomUsername(),
		Qty:        utils.RandomInt(1, 10),
		EachPrice:  utils.RandomMoney(),
		TotalPrice: utils.RandomMoney(),
	}

	epd, err := testQueries.AddPurchaseDetail(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, epd)

	require.Equal(t, arg.EventID, epd.EventID)
	require.Equal(t, arg.Name, epd.Name)
	require.Equal(t, arg.Qty, epd.Qty)
	require.Equal(t, arg.EachPrice, epd.EachPrice)
	require.Equal(t, arg.TotalPrice, epd.TotalPrice)

	return epd
}

func TestAddPurchaseDetail(t *testing.T) {
	user := GenerateUser(t)
	event := GenerateEvent(t, user)
	GeneratePurchaseDetail(t, event)
}

func TestGetPurchaseDetailByEventID(t *testing.T) {
	user := GenerateUser(t)
	event := GenerateEvent(t, user)
	for i := 0; i < 3; i++ {

		epd := GeneratePurchaseDetail(t, event)

		payload, err := testQueries.GetPurchaseDetailByEventID(context.Background(), event.ID)
		require.NoError(t, err)
		require.NotEmpty(t, payload)

		require.Equal(t, epd.ID, payload[i].ID)
		require.Equal(t, epd.EventID, payload[i].EventID)
		require.Equal(t, epd.Name, payload[i].Name)
		require.Equal(t, epd.Qty, payload[i].Qty)
		require.Equal(t, epd.EachPrice, payload[i].EachPrice)
		require.Equal(t, epd.TotalPrice, payload[i].TotalPrice)
	}
}

func TestUpdatePurchaseDetail(t *testing.T) {
	user := GenerateUser(t)
	event := GenerateEvent(t, user)
	epd := GeneratePurchaseDetail(t, event)
	arg := queries.UpdatePurchaseDetailParams{
		ID:         epd.ID,
		Name:       RandomUsername(),
		Qty:        utils.RandomInt(10, 20),
		EachPrice:  utils.RandomMoney(),
		TotalPrice: utils.RandomMoney(),
		UpdatedAt:  utils.ToNullTime(time.Now()),
	}

	payload, err := testQueries.UpdatePurchaseDetail(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, payload)

	require.Equal(t, arg.ID, payload.ID)
	require.Equal(t, arg.Name, payload.Name)
	require.Equal(t, arg.Qty, payload.Qty)
	require.Equal(t, arg.EachPrice, payload.EachPrice)
	require.Equal(t, arg.TotalPrice, payload.TotalPrice)
}
