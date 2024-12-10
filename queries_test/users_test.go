package queries

import (
	"context"
	"database/sql"
	"tangapp-be/queries"
	"tangapp-be/utils"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func GenerateUser(t *testing.T) queries.User {
	// Arrange
	username := utils.RandomUsername()
	arg := queries.AddUserParams{
		Username: sql.NullString{String: username, Valid: true},
		Email:    username + "@gmail.com",
	}

	// Act
	user, err := testQueries.AddUser(context.Background(), arg)

	// Assert
	require.NoError(t, err)
	require.NotEmpty(t, user)
	require.Equal(t, arg.Username.String, user.Username.String)
	require.Equal(t, arg.Email, user.Email)

	return user
}

func TestAddUser(t *testing.T) {
	GenerateUser(t)
}

func TestGetUserByID(t *testing.T) {
	user := GenerateUser(t)

	payload, err := testQueries.GetUserByID(context.Background(), user.ID)

	require.NoError(t, err)
	require.NotEmpty(t, payload)
	require.Equal(t, user, payload) // Haha langsung compare struct haha

}

func TestGetUserByEmail(t *testing.T) {
	user := GenerateUser(t)

	payload, err := testQueries.GetUserByEmail(context.Background(), user.Email)

	require.NoError(t, err)
	require.NotEmpty(t, payload)
	require.Equal(t, user, payload) // Haha langsung compare struct haha

}

func TestUpdateUser(t *testing.T) {
	user := GenerateUser(t)
	newUsername := utils.RandomUsername()
	arg := queries.UpdateUserParams{
		ID:       user.ID,
		Username: sql.NullString{String: newUsername, Valid: true},
	}

	payload, err := testQueries.UpdateUser(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, payload)
	require.Equal(t, user.ID, payload.ID)
	require.Equal(t, sql.NullString{String: newUsername, Valid: true}, payload.Username)
	require.Equal(t, user.Email, payload.Email)
	require.WithinDuration(t, user.CreatedAt, payload.CreatedAt, time.Second)
}

// Generates an username
func RandomUsername() string {
	return utils.RandomString(6)
}
