package db

import (
	"context"
	"database/sql"
	"testing"

	"github.com/bihany-harsh/mini-twitter/util"
	"github.com/stretchr/testify/require"
)

func checkEqualAccounts(t *testing.T, account1, account2 Account) {
	require.Equal(t, account1.Username, account2.Username)
	require.Equal(t, account1.Email, account2.Email)
	require.Equal(t, account1.ProfilePictureUrl, account2.ProfilePictureUrl)
	require.Equal(t, account1.Bio, account2.Bio)
	require.Equal(t, account1.LastLogin, account2.LastLogin)
	require.Equal(t, account1.IsAdmin, account2.IsAdmin)
	require.Equal(t, account1.IsActive, account2.IsActive)
	require.Equal(t, account1.LastDeactivatedAt, account2.LastDeactivatedAt)
	require.Equal(t, account1.NFollowers, account2.NFollowers)
	require.Equal(t, account1.NFollowing, account2.NFollowing)
	require.Equal(t, account1.NTweets, account2.NTweets)
	require.Equal(t, account1.ID, account2.ID)
	require.Equal(t, account1.CreatedAt, account2.CreatedAt)
}

func createRandomAccount(t *testing.T) Account {
	arg := CreateAccountParams{
		Username:          util.RandomUsername(),
		Email:             util.RandomEmail(),
		ProfilePictureUrl: util.RandomProfilePictureUrl(p, N),
		Bio:               util.RandomBio(p, N),
		LastLogin:         util.RandomTime_Nullable(p),
		IsActive:          true,
	}

	account, err := testQueries.CreateAccount(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, account)
	require.Equal(t, arg.Username, account.Username)
	require.Equal(t, arg.Email, account.Email)
	require.Equal(t, arg.ProfilePictureUrl, account.ProfilePictureUrl)
	require.Equal(t, arg.Bio, account.Bio)
	require.Equal(t, arg.LastLogin, account.LastLogin)
	require.Equal(t, arg.IsAdmin, account.IsAdmin)
	require.Equal(t, arg.IsActive, account.IsActive)
	require.Equal(t, arg.LastDeactivatedAt, account.LastDeactivatedAt)
	require.Equal(t, arg.NFollowers, account.NFollowers)
	require.Equal(t, arg.NFollowing, account.NFollowing)
	require.Equal(t, arg.NTweets, account.NTweets)
	require.NotZero(t, account.ID)
	require.NotZero(t, account.CreatedAt)

	return account
}

func TestCreateAccount(t *testing.T) {
	createRandomAccount(t)
}

func TestGetAccountByEmail(t *testing.T) {
	account1 := createRandomAccount(t)
	account2, err := testQueries.GetAccountByEmail(context.Background(), account1.Email)

	require.NoError(t, err)
	require.NotEmpty(t, account2)
	checkEqualAccounts(t, account1, account2)
}

func TestGetAccountByUsername(t *testing.T) {
	account1 := createRandomAccount(t)
	account2, err := testQueries.GetAccountByUsername(context.Background(), account1.Username)

	require.NoError(t, err)
	require.NotEmpty(t, account2)
	checkEqualAccounts(t, account1, account2)
}

func TestGetAccountByID(t *testing.T) {
	account1 := createRandomAccount(t)
	account2, err := testQueries.GetAccountByID(context.Background(), account1.ID)

	require.NoError(t, err)
	require.NotEmpty(t, account2)
	checkEqualAccounts(t, account1, account2)
}

func TestGetAccountByUsernameOrEmail(t *testing.T) {
	account1 := createRandomAccount(t)
	account_by_username, err := testQueries.GetAccountByUsernameOrEmail(context.Background(), GetAccountByUsernameOrEmailParams{
		Username: account1.Username,
	})
	require.NoError(t, err)
	require.NotEmpty(t, account_by_username)
	checkEqualAccounts(t, account1, account_by_username)
	account_by_email, err := testQueries.GetAccountByUsernameOrEmail(context.Background(), GetAccountByUsernameOrEmailParams{
		Email: account1.Email,
	})
	require.NoError(t, err)
	require.NotEmpty(t, account_by_email)
	checkEqualAccounts(t, account1, account_by_email)
}

func TestDeleteAccountByEmail(t *testing.T) {
	account1 := createRandomAccount(t)
	err := testQueries.DeleteAccountByEmail(context.Background(), account1.Email)

	require.NoError(t, err)
	account2, err := testQueries.GetAccountByEmail(context.Background(), account1.Email)

	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, account2)
}

func TestDeleteAccountByUsername(t *testing.T) {
	account1 := createRandomAccount(t)
	err := testQueries.DeleteAccountByUsername(context.Background(), account1.Username)

	require.NoError(t, err)
	account2, err := testQueries.GetAccountByUsername(context.Background(), account1.Username)

	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, account2)
}

func TestDeleteAccountByID(t *testing.T) {
	account1 := createRandomAccount(t)
	err := testQueries.DeleteAccountByID(context.Background(), account1.ID)

	require.NoError(t, err)
	account2, err := testQueries.GetAccountByID(context.Background(), account1.ID)

	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, account2)
}

func TestListAccounts(t *testing.T) {
	for i := 0; i < 10; i++ {
		createRandomAccount(t)
	}

	arg := ListAccountsParams{
		Limit:  5,
		Offset: 5,
	}

	accounts, err := testQueries.ListAccounts(context.Background(), arg)

	require.NoError(t, err)
	require.Len(t, accounts, 5)

	for _, account := range accounts {
		require.NotEmpty(t, account)
	}
}

func TestUpdateAccountByEmail(t *testing.T) {
	account := createRandomAccount(t)
	arg := UpdateAccountByEmailParams{
		Email:             account.Email,
		Username:          util.RandomUsername(),
		ProfilePictureUrl: util.RandomProfilePictureUrl(p, N),
		Bio:               util.RandomBio(p, N),
		LastLogin:         util.RandomTime_Nullable(p),
		IsActive:          util.RandomBool(),
		IsAdmin:           util.RandomBool(),
		LastDeactivatedAt: util.RandomTime_Nullable(p),
	}

	account2, err := testQueries.UpdateAccountByEmail(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, account2)
	require.Equal(t, arg.Username, account2.Username)
	require.Equal(t, arg.Email, account2.Email)
	require.Equal(t, arg.ProfilePictureUrl, account2.ProfilePictureUrl)
	require.Equal(t, arg.Bio, account2.Bio)
	require.Equal(t, arg.LastLogin, account2.LastLogin)
	require.Equal(t, arg.IsAdmin, account2.IsAdmin)
	require.Equal(t, arg.IsActive, account2.IsActive)
	require.Equal(t, arg.LastDeactivatedAt, account2.LastDeactivatedAt)
	require.Equal(t, account.NFollowers, account2.NFollowers)
	require.Equal(t, account.NFollowing, account2.NFollowing)
	require.Equal(t, account.NTweets, account2.NTweets)
	require.Equal(t, account.ID, account2.ID)
	require.Equal(t, account.CreatedAt, account2.CreatedAt)
}

func TestUpdateAccountByUsername(t *testing.T) {
	account := createRandomAccount(t)
	arg := UpdateAccountByUsernameParams{
		Username:          account.Username,
		Email:             util.RandomEmail(),
		ProfilePictureUrl: util.RandomProfilePictureUrl(p, N),
		Bio:               util.RandomBio(p, N),
		LastLogin:         util.RandomTime_Nullable(p),
		IsActive:          util.RandomBool(),
		IsAdmin:           util.RandomBool(),
		LastDeactivatedAt: util.RandomTime_Nullable(p),
	}

	account2, err := testQueries.UpdateAccountByUsername(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, account2)
	require.Equal(t, arg.Username, account2.Username)
	require.Equal(t, arg.Email, account2.Email)
	require.Equal(t, arg.ProfilePictureUrl, account2.ProfilePictureUrl)
	require.Equal(t, arg.Bio, account2.Bio)
	require.Equal(t, arg.LastLogin, account2.LastLogin)
	require.Equal(t, arg.IsAdmin, account2.IsAdmin)
	require.Equal(t, arg.IsActive, account2.IsActive)
	require.Equal(t, arg.LastDeactivatedAt, account2.LastDeactivatedAt)
	require.Equal(t, account.NFollowers, account2.NFollowers)
	require.Equal(t, account.NFollowing, account2.NFollowing)
	require.Equal(t, account.NTweets, account2.NTweets)
	require.Equal(t, account.ID, account2.ID)
	require.Equal(t, account.CreatedAt, account2.CreatedAt)
}

func TestUpdateAccountByID(t *testing.T) {
	account := createRandomAccount(t)
	arg := UpdateAccountByIDParams{
		ID:                account.ID,
		Username:          util.RandomUsername(),
		Email:             util.RandomEmail(),
		ProfilePictureUrl: util.RandomProfilePictureUrl(p, N),
		Bio:               util.RandomBio(p, N),
		LastLogin:         util.RandomTime_Nullable(p),
		IsActive:          util.RandomBool(),
		IsAdmin:           util.RandomBool(),
		LastDeactivatedAt: util.RandomTime_Nullable(p),
	}

	account2, err := testQueries.UpdateAccountByID(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, account2)
	require.Equal(t, arg.Username, account2.Username)
	require.Equal(t, arg.Email, account2.Email)
	require.Equal(t, arg.ProfilePictureUrl, account2.ProfilePictureUrl)
	require.Equal(t, arg.Bio, account2.Bio)
	require.Equal(t, arg.LastLogin, account2.LastLogin)
	require.Equal(t, arg.IsAdmin, account2.IsAdmin)
	require.Equal(t, arg.IsActive, account2.IsActive)
	require.Equal(t, arg.LastDeactivatedAt, account2.LastDeactivatedAt)
	require.Equal(t, account.NFollowers, account2.NFollowers)
	require.Equal(t, account.NFollowing, account2.NFollowing)
	require.Equal(t, account.NTweets, account2.NTweets)
	require.Equal(t, account.ID, account2.ID)
	require.Equal(t, account.CreatedAt, account2.CreatedAt)
}
