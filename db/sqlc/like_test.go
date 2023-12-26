package db

import (
	"context"
	"database/sql"
	"testing"

	"github.com/bihany-harsh/mini-twitter/util"
	"github.com/stretchr/testify/require"
)

func createRandomLikeToTweet(t *testing.T, tweet *Tweet) Like {
	account2 := createRandomAccount(t)
	n_likes_old := tweet.NLikes

	arg := CreateLikeParams{
		UserID:  account2.ID,
		TweetID: tweet.ID,
	}

	like, err := testQueries.CreateLike(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, like)
	require.Equal(t, arg.UserID, like.UserID)
	require.Equal(t, arg.TweetID, like.TweetID)
	require.NotZero(t, like.CreatedAt)

	*tweet, err = testQueries.AddNLikesByOne(context.Background(), tweet.ID)
	require.NoError(t, err)
	require.NotEmpty(t, tweet)
	require.Equal(t, n_likes_old+1, tweet.NLikes)

	return like
}

func createRandomLikeByAccount(t *testing.T, account2 Account) Like {
	account1 := createRandomAccount(t)
	tweet, account1 := createRandomTweetByAccount(t, account1)
	n_likes_old := tweet.NLikes

	arg := CreateLikeParams{
		UserID:  account2.ID,
		TweetID: tweet.ID,
	}

	like, err := testQueries.CreateLike(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, like)
	require.Equal(t, arg.UserID, like.UserID)
	require.Equal(t, arg.TweetID, like.TweetID)
	require.NotZero(t, like.CreatedAt)

	tweet, err = testQueries.AddNLikesByOne(context.Background(), tweet.ID)
	require.NoError(t, err)
	require.NotEmpty(t, tweet)
	require.Equal(t, n_likes_old+1, tweet.NLikes)

	return like
}

func TestCreateLike(t *testing.T) {
	account1 := createRandomAccount(t)
	tweet, account1 := createRandomTweetByAccount(t, account1)

	createRandomLikeToTweet(t, &tweet)
}

func TestGetLike(t *testing.T) {
	account1 := createRandomAccount(t)
	tweet, account1 := createRandomTweetByAccount(t, account1)

	like := createRandomLikeToTweet(t, &tweet)

	arg := GetLikeParams{
		UserID:  like.UserID,
		TweetID: like.TweetID,
	}

	like, err := testQueries.GetLike(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, like)
	require.Equal(t, arg.UserID, like.UserID)
	require.Equal(t, arg.TweetID, like.TweetID)
	require.NotZero(t, like.CreatedAt)
}

func TestDeleteLike(t *testing.T) {
	account1 := createRandomAccount(t)
	tweet, account1 := createRandomTweetByAccount(t, account1)

	like := createRandomLikeToTweet(t, &tweet)

	n_likes_old := tweet.NLikes

	arg := DeleteLikeParams{
		UserID:  like.UserID,
		TweetID: like.TweetID,
	}
	err := testQueries.DeleteLike(context.Background(), arg)
	require.NoError(t, err)

	tweet, err = testQueries.SubtractNLikesByOne(context.Background(), tweet.ID)
	require.NoError(t, err)
	require.NotEmpty(t, tweet)
	require.Equal(t, n_likes_old-1, tweet.NLikes)

	getLikeArg := GetLikeParams{
		UserID:  like.UserID,
		TweetID: like.TweetID,
	}
	like, err = testQueries.GetLike(context.Background(), getLikeArg)
	require.Error(t, err)
	require.Empty(t, like)
	require.EqualError(t, err, sql.ErrNoRows.Error())
}

func TestGetLikesByTweetID(t *testing.T) {
	account1 := createRandomAccount(t)
	tweet, account1 := createRandomTweetByAccount(t, account1)
	n := util.RandomInt(1, 10)
	for i := int64(0); i < n; i++ {
		createRandomLikeToTweet(t, &tweet)
	}

	arg := GetLikesByTweetIDParams{
		TweetID: tweet.ID,
		Limit:   int32(n),
		Offset:  0,
	}

	likes, err := testQueries.GetLikesByTweetID(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, likes)
	require.Len(t, likes, int(n))
}

func TestGetLikesByUserID(t *testing.T) {
	account2 := createRandomAccount(t)
	n := util.RandomInt(1, 10)
	for i := int64(0); i < n; i++ {
		createRandomLikeByAccount(t, account2)
	}

	arg := GetLikesByUserIDParams{
		UserID: account2.ID,
		Limit:  int32(n),
		Offset: 0,
	}

	likes, err := testQueries.GetLikesByUserID(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, likes)
	require.Len(t, likes, int(n))
}
