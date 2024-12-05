package repository

import (
	"context"
	"tangapp-be/repository"
	"tangapp-be/utils"
	"testing"

	"github.com/stretchr/testify/require"
)

func GenerateMemberDetail(t *testing.T, user repository.User, event repository.Event) repository.EventMemberDetail {
	arg := repository.AddMemberDetailParams{
		EventID:      event.ID,
		UserID:       user.ID,
		Bill:         NullInt32(utils.RandomInt(1000, 10000)),
		Paid:         NullInt32(utils.RandomInt(1000, 10000)),
		Compensation: NullInt32(utils.RandomInt(1000, 10000)),
		Notes:        NullString(utils.RandomString(10)),
		Done:         false,
	}

	payload, err := testQueries.AddMemberDetail(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, payload)

	require.Equal(t, arg.EventID, payload.EventID)
	require.Equal(t, arg.UserID, payload.UserID)
	require.Equal(t, arg.Bill, payload.Bill)
	require.Equal(t, arg.Paid, payload.Paid)
	require.Equal(t, arg.Compensation, payload.Compensation)
	require.Equal(t, arg.Notes, payload.Notes)

	return payload
}

func TestAddMemberDetail(t *testing.T) {
	user := GenerateUser(t)
	event := GenerateEvent(t, user)

	GenerateMemberDetail(t, user, event)

}

func TestGetMemberDetail(t *testing.T) {
	user := GenerateUser(t)
	event := GenerateEvent(t, user)
	memberDetail := GenerateMemberDetail(t, user, event)

	payload, err := testQueries.GetMemberDetail(context.Background(), memberDetail.ID)

	require.NoError(t, err)
	require.NotEmpty(t, payload)

	require.Equal(t, memberDetail, payload)
}

func TestListMemberDetails(t *testing.T) {
	user := GenerateUser(t)
	event := GenerateEvent(t, user)

	for i := 0; i < 5; i++ {
		memberDetail := GenerateMemberDetail(t, user, event)

		payload, err := testQueries.ListMemberDetail(context.Background(), event.ID)

		require.NoError(t, err)
		require.NotEmpty(t, payload)

		require.Equal(t, memberDetail.EventID, payload[i].EventID)
		require.Equal(t, memberDetail.UserID, payload[i].UserID)
		require.Equal(t, memberDetail.Bill, payload[i].Bill)
		require.Equal(t, memberDetail.Paid, payload[i].Paid)
		require.Equal(t, memberDetail.Compensation, payload[i].Compensation)
		require.Equal(t, memberDetail.Notes, payload[i].Notes)
		require.Equal(t, memberDetail.Done, payload[i].Done)
	}

}

func TestUpdateMemberDetail(t *testing.T) {
	user := GenerateUser(t)
	event := GenerateEvent(t, user)
	memberDetail := GenerateMemberDetail(t, user, event)
	arg := repository.UpdateMemberDetailParams{
		ID:           memberDetail.ID,
		Bill:         utils.ToNullInt32(utils.RandomInt(1000, 10000)),
		Paid:         utils.ToNullInt32(utils.RandomInt(1000, 10000)),
		Compensation: utils.ToNullInt32(utils.RandomInt(1000, 10000)),
		Notes:        utils.ToNullString(utils.RandomString(10)),
		Done:         true,
	}

	payload, err := testQueries.UpdateMemberDetail(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, payload)

	require.Equal(t, arg.ID, payload.ID)
	require.Equal(t, arg.Bill, payload.Bill)
	require.Equal(t, arg.Paid, payload.Paid)
	require.Equal(t, arg.Compensation, payload.Compensation)
	require.Equal(t, arg.Notes, payload.Notes)
	require.Equal(t, arg.Done, payload.Done)
}
