package db

import (
	"context"
	"database/sql"
	"testing"

	"github.com/bihany-harsh/mini-twitter/util"
	"github.com/stretchr/testify/require"
)

func createRandomFollow(t *testing.T) Follow {
	account1 := createRandomAccount(t)
	account2 := createRandomAccount(t)

	arg := CreateFollowParams{
		FollowerID:  account1.ID,
		FollowingID: account2.ID,
	}

	account1_following := account1.NFollowing
	account2_followers := account2.NFollowers

	follow, err := testQueries.CreateFollow(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, follow)
	require.Equal(t, arg.FollowerID, follow.FollowerID)
	require.Equal(t, arg.FollowingID, follow.FollowingID)

	account2, err = testQueries.AddNFollowersByOne(context.Background(), account2.ID)
	require.NoError(t, err)
	account1, err = testQueries.AddNFollowingByOne(context.Background(), account1.ID)
	require.NoError(t, err)

	require.Equal(t, account1_following+1, account1.NFollowing)
	require.Equal(t, account2_followers+1, account2.NFollowers)

	return follow
}

func createRandomFollowerForAccount(t *testing.T, account1 *Account) {
	account2 := createRandomAccount(t)

	arg := CreateFollowParams{
		FollowerID:  account2.ID,
		FollowingID: account1.ID,
	}

	account2_following := account2.NFollowing
	account1_followers := account1.NFollowers

	follow, err := testQueries.CreateFollow(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, follow)
	require.Equal(t, arg.FollowerID, follow.FollowerID)
	require.Equal(t, arg.FollowingID, follow.FollowingID)

	account2, err = testQueries.AddNFollowingByOne(context.Background(), account2.ID)
	require.NoError(t, err)
	*account1, err = testQueries.AddNFollowersByOne(context.Background(), account1.ID)
	require.NoError(t, err)

	require.Equal(t, account2_following+1, account2.NFollowing)
	require.Equal(t, account1_followers+1, account1.NFollowers)
}

func createRandomFollowingByAccount(t *testing.T, account1 *Account) {
	account2 := createRandomAccount(t)

	arg := CreateFollowParams{
		FollowerID:  account1.ID,
		FollowingID: account2.ID,
	}

	account1_following := account1.NFollowing
	account2_followers := account2.NFollowers

	follow, err := testQueries.CreateFollow(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, follow)
	require.Equal(t, arg.FollowerID, follow.FollowerID)
	require.Equal(t, arg.FollowingID, follow.FollowingID)

	*account1, err = testQueries.AddNFollowingByOne(context.Background(), account1.ID)
	require.NoError(t, err)
	account2, err = testQueries.AddNFollowersByOne(context.Background(), account2.ID)
	require.NoError(t, err)

	require.Equal(t, account1_following+1, account1.NFollowing)
	require.Equal(t, account2_followers+1, account2.NFollowers)
}

func TestCreateFollow(t *testing.T) {
	createRandomFollow(t)
}

func TestGetFollow(t *testing.T) {
	follow := createRandomFollow(t)

	followArg := GetFollowParams{
		FollowerID:  follow.FollowerID,
		FollowingID: follow.FollowingID,
	}

	follow2, err := testQueries.GetFollow(context.Background(), followArg)
	require.NoError(t, err)
	require.NotEmpty(t, follow2)
	require.Equal(t, follow.FollowerID, follow2.FollowerID)
	require.Equal(t, follow.FollowingID, follow2.FollowingID)
}

func TestDeleteFollow(t *testing.T) {
	follow := createRandomFollow(t)
	account1, err := testQueries.GetAccountByID(context.Background(), follow.FollowerID)
	require.NoError(t, err)
	account2, err := testQueries.GetAccountByID(context.Background(), follow.FollowingID)
	require.NoError(t, err)

	account1_following := account1.NFollowing
	account2_followers := account2.NFollowers

	arg := DeleteFollowParams{
		FollowerID:  follow.FollowerID,
		FollowingID: follow.FollowingID,
	}

	err = testQueries.DeleteFollow(context.Background(), arg)
	require.NoError(t, err)

	account1, err = testQueries.SubtractNFollowingByOne(context.Background(), account1.ID)
	require.NoError(t, err)
	account2, err = testQueries.SubtractNFollowersByOne(context.Background(), account2.ID)
	require.NoError(t, err)

	followArg := GetFollowParams{
		FollowerID:  follow.FollowerID,
		FollowingID: follow.FollowingID,
	}

	follow2, err := testQueries.GetFollow(context.Background(), followArg)
	require.Error(t, err)
	require.Empty(t, follow2)
	require.EqualError(t, err, sql.ErrNoRows.Error())

	require.Equal(t, account1_following-1, account1.NFollowing)
	require.Equal(t, account2_followers-1, account2.NFollowers)
}

func TestGetFollowers(t *testing.T) {
	account1 := createRandomAccount(t)
	n := util.RandomInt(1, 10)
	for i := int64(0); i < n; i++ {
		createRandomFollowerForAccount(t, &account1)
	}

	arg := GetFollowersParams{
		FollowingID: account1.ID,
		Limit:       int32(n),
		Offset:      0,
	}

	follows, err := testQueries.GetFollowers(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, follows)
	require.Equal(t, len(follows), int(arg.Limit))

	require.Equal(t, account1.NFollowers, int32(n))
}

func TestGetFollowing(t *testing.T) {
	account1 := createRandomAccount(t)
	n := util.RandomInt(1, 10)
	for i := int64(0); i < n; i++ {
		createRandomFollowingByAccount(t, &account1)
	}

	arg := GetFollowingParams{
		FollowerID: account1.ID,
		Limit:      int32(n),
		Offset:     0,
	}

	follows, err := testQueries.GetFollowing(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, follows)
	require.Equal(t, len(follows), int(arg.Limit))

	require.Equal(t, account1.NFollowing, int32(n))
}
