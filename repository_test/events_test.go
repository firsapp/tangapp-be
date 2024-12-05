package repository

import (
	"context"
	"database/sql"
	"tangapp-be/repository"
	"tangapp-be/utils"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func GenerateEvent(t *testing.T, user repository.User) repository.Event {

	arg := repository.AddEventParams{
		CreatedBy:   user.ID,
		Title:       sql.NullString{String: utils.RandomString(8), Valid: true},
		Description: sql.NullString{String: utils.RandomString(8), Valid: true},
		Status:      "berhasil",
		TotalAmount: int32(utils.RandomInt(1000, 100000)),
		DateEvent:   sql.NullTime{Time: time.Now(), Valid: true},
		CanEdit:     true,
	}

	payload, err := testQueries.AddEvent(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, payload)

	require.Equal(t, arg.CreatedBy, payload.CreatedBy)
	require.Equal(t, arg.Title, payload.Title)
	require.Equal(t, arg.Description, payload.Description)
	require.Equal(t, arg.Status, payload.Status)
	require.Equal(t, arg.TotalAmount, payload.TotalAmount)
	require.Equal(t, arg.CanEdit, payload.CanEdit)
	require.True(t, payload.IsActive)

	return payload
}

func TestAddEvent(t *testing.T) {
	user := GenerateUser(t)
	GenerateEvent(t, user)

}

func TestListEventByUser(t *testing.T) {
	user := GenerateUser(t)
	for i := 0; i < 5; i++ {

		event := GenerateEvent(t, user)

		payload, err := testQueries.ListEventByUser(context.Background(), user.ID)

		require.NoError(t, err)
		require.NotEmpty(t, payload)

		require.Equal(t, event.CreatedBy, payload[i].CreatedBy)
		require.Equal(t, event.Title, payload[i].Title)
		require.Equal(t, event.Description, payload[i].Description)
		require.Equal(t, event.Status, payload[i].Status)
		require.Equal(t, event.TotalAmount, payload[i].TotalAmount)
		require.Equal(t, event.CanEdit, payload[i].CanEdit)
		require.True(t, payload[i].IsActive)
	}

}

func TestUpdateEvent(t *testing.T) {
	user := GenerateUser(t)
	event := GenerateEvent(t, user)
	arg := repository.UpdateEventParams{
		ID:          event.ID,
		Title:       NullString(utils.RandomString(10)),
		Description: NullString(utils.RandomString(10)),
		Status:      "Gagal",
		TotalAmount: utils.RandomInt(1000, 10000),
		DateEvent:   NullTime(time.Now()),
	}

	payload, err := testQueries.UpdateEvent(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, payload)

	require.Equal(t, arg.Title, payload.Title)
	require.Equal(t, arg.Description, payload.Description)
	require.Equal(t, arg.Status, payload.Status)
	require.Equal(t, arg.TotalAmount, payload.TotalAmount)
}

func TestDeleteEvent(t *testing.T) {
	user := GenerateUser(t)
	event := GenerateEvent(t, user)
	arg := repository.DeleteEventParams{
		ID:       event.ID,
		IsActive: false,
	}

	payload, err := testQueries.DeleteEvent(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, payload)

	require.False(t, arg.IsActive)
}
