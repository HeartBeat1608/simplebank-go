package db

import (
	"context"
	"testing"
	"time"

	"github.com/Heartbeat1608/simplebank/db/util"
	"github.com/stretchr/testify/require"
)

func createRandomTransfer(t *testing.T) Transfer {
	ac1 := createRandomAccount(t)
	ac2 := createRandomAccount(t)

	arg := CreateTransferParams{
		FromAccountID: ac1.ID,
		ToAccountID:   ac2.ID,
		Amount:        util.RandomMoney(),
	}

	trf, err := testQueries.CreateTransfer(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, trf)

	return trf
}

func TestCreateTransfer(t *testing.T) {
	createRandomTransfer(t)
}

func TestUpdateTransfer(t *testing.T) {
	trf := createRandomTransfer(t)

	arg := UpdateTransferParams{
		ID:     trf.ID,
		Amount: util.RandomMoney(),
	}

	trf2, err := testQueries.UpdateTransfer(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, trf2)

	require.Equal(t, trf.ID, trf2.ID)
	require.Equal(t, trf.FromAccountID, trf2.FromAccountID)
	require.Equal(t, trf.ToAccountID, trf2.ToAccountID)
	require.Equal(t, arg.Amount, trf2.Amount)

	require.WithinDuration(t, trf.CreatedAt, trf2.CreatedAt, time.Second)
}

func TestGetTransfer(t *testing.T) {
	trf := createRandomTransfer(t)

	trf2, err := testQueries.GetTransfer(context.Background(), trf.ID)
	require.NoError(t, err)
	require.NotEmpty(t, trf2)

	require.Equal(t, trf.ID, trf2.ID)
	require.Equal(t, trf.FromAccountID, trf2.FromAccountID)
	require.Equal(t, trf.ToAccountID, trf2.ToAccountID)
	require.Equal(t, trf.Amount, trf2.Amount)

	require.WithinDuration(t, trf.CreatedAt, trf2.CreatedAt, time.Second)
}

func TestDeleteTransfer(t *testing.T) {
	trf := createRandomTransfer(t)

	trf2, err := testQueries.DeleteTransfer(context.Background(), trf.ID)
	require.NoError(t, err)
	require.NotEmpty(t, trf2)

	require.Equal(t, trf.ID, trf2.ID)
	require.Equal(t, trf.FromAccountID, trf2.FromAccountID)
	require.Equal(t, trf.ToAccountID, trf2.ToAccountID)
	require.Equal(t, trf.Amount, trf2.Amount)

	require.WithinDuration(t, trf.CreatedAt, trf2.CreatedAt, time.Second)
}

func TestListTransfers(t *testing.T) {
	for i := 0; i < 10; i++ {
		createRandomTransfer(t)
	}

	arg := ListTransfersParams{
		Limit:  10,
		Offset: 0,
	}

	trfs, err := testQueries.ListTransfers(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, trfs, 10)

	for _, trf := range trfs {
		require.NotEmpty(t, trf)
	}
}
