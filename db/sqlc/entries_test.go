package db

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func createRandomEntries(t *testing.T) Entry {
	// 1. BUAT DATA INDUK (ACCOUNT) TERLEBIH DAHULU
	account := createRandomAccount(t)

	// 2. Siapkan parameter untuk membuat data anak (entry)
	arg := CreateEntriesParams{
		AccountID: account.ID,
		Amount:    account.Balance,
	}

	// 4. Buat data anak (entry)
	entry, err := testQueries.CreateEntries(context.Background(), arg)

	// 5. Lakukan verifikasi
	require.NoError(t, err)
	require.NotEmpty(t, entry)

	// Pastikan data yang dimasukkan sesuai
	require.Equal(t, arg.AccountID, entry.AccountID)
	require.Equal(t, arg.Amount, entry.Amount)

	// Pastikan ID dan timestamp dibuat oleh database
	require.NotZero(t, entry.ID)
	require.NotZero(t, entry.CreatedAt)

	return entry

}

func TestCreateEntries(t *testing.T) {
	createRandomEntries(t)
}

func TestGetEntries(t *testing.T) {
	entries := createRandomEntries(t)
	entry, err := testQueries.GetEntries(context.Background(), entries.ID)
	require.NoError(t, err)
	require.NotEmpty(t, entry)

	require.Equal(t, entries.ID, entry.ID)
	require.Equal(t, entries.AccountID, entry.AccountID)
	require.Equal(t, entries.Amount, entry.Amount)
	require.WithinDuration(t, entries.CreatedAt.Time, entry.CreatedAt.Time, time.Second)
}

func TestUpdateEntries(t *testing.T) {
	entries := createRandomEntries(t)
	arg := UpdateEntriesParams{
		ID:     entries.ID,
		Amount: 200,
	}
	entry, err := testQueries.UpdateEntries(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, entry)

	require.Equal(t, entries.ID, entry.ID)
}

func TestListEntries(t *testing.T) {
	for i := 0; i < 10; i++ {
		createRandomEntries(t)
	}
	arg := ListEntriesParams{
		Limit:  5,
		Offset: 5,
	}
	entries, err := testQueries.ListEntries(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, entries, 5)

	for _, entry := range entries {
		require.NotEmpty(t, entry)
	}
}
