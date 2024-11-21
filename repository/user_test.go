package repository

import (
	"context"
	"database/sql"
	"tangapp-be/utils"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestCreateUser(t *testing.T) {
	// Arrange
	ctx := context.Background()
	arg := CreateAccountParams{
		Name:      sql.NullString{String: utils.RandomOwner(), Valid: true},
		Title:     sql.NullString{String: "Pejabat", Valid: true},
		CreatedAt: sql.NullTime{Time: time.Now(), Valid: true},
	}

	// Act
	user, err := testQueries.CreateAccount(ctx, arg)

	// Assert
	require.NoError(t, err)
	require.NotEmpty(t, user)
	require.Equal(t, arg.Name.String, user.Name.String)
	require.Equal(t, arg.Title.String, user.Title.String)
}
