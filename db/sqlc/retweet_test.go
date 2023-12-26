package db

import (
	"context"
	"database/sql"
	"testing"

	"github.com/bihany-harsh/mini-twitter/util"
	"github.com/stretchr/testify/require"
)

func TestCreateRetweet(t *testing.T) {
	account1 := createRandomAccount(t)
	tweet, account1 := createRandomTweetByAccount(t, account1)

	n_retweets_old := tweet.NRetweets

	account2 := createRandomAccount(t)
	retweetArg := CreateRetweetParams{
		UserID:  account2.ID,
		TweetID: tweet.ID,
	}
	retweet, err := testQueries.CreateRetweet(context.Background(), retweetArg)
	require.NoError(t, err)
	require.NotEmpty(t, retweet)
	require.Equal(t, retweetArg.UserID, retweet.UserID)

	tweet_r, account2 := createRandomTweetByAccount(t, account2)
	setRTweetArg := SetRetweetIDParams{
		ID: tweet_r.ID,
		RetweetID: sql.NullInt64{
			Int64: tweet.ID,
			Valid: true,
		},
	}
	tweet_r, err = testQueries.SetRetweetID(context.Background(), setRTweetArg)
	require.NoError(t, err)
	require.NotEmpty(t, tweet_r)
	require.Equal(t, setRTweetArg.RetweetID.Int64, tweet_r.RetweetID.Int64)
	require.Equal(t, setRTweetArg.RetweetID.Valid, tweet_r.RetweetID.Valid)

	tweet, err = testQueries.AddNRetweetsByOne(context.Background(), tweet.ID)
	require.NoError(t, err)
	require.NotEmpty(t, tweet)
	require.Equal(t, n_retweets_old+1, tweet.NRetweets)
}

func TestDeleteRetweet(t *testing.T) {
	account1 := createRandomAccount(t)
	tweet, account1 := createRandomTweetByAccount(t, account1)

	n_retweets_old := tweet.NRetweets

	account2 := createRandomAccount(t)
	retweetArg := CreateRetweetParams{
		UserID:  account2.ID,
		TweetID: tweet.ID,
	}
	retweet, err := testQueries.CreateRetweet(context.Background(), retweetArg)
	require.NoError(t, err)
	require.NotEmpty(t, retweet)
	require.Equal(t, retweetArg.UserID, retweet.UserID)

	tweet, err = testQueries.AddNRetweetsByOne(context.Background(), tweet.ID)
	require.NoError(t, err)
	require.NotEmpty(t, tweet)
	require.Equal(t, n_retweets_old+1, tweet.NRetweets)

	deleteRetweetArg := DeleteRetweetParams{
		UserID:  account2.ID,
		TweetID: tweet.ID,
	}
	err = testQueries.DeleteRetweet(context.Background(), deleteRetweetArg)
	require.NoError(t, err)

	tweet, err = testQueries.SubtractNRetweetsByOne(context.Background(), tweet.ID)
	require.NoError(t, err)
	require.NotEmpty(t, tweet)
	require.Equal(t, n_retweets_old, tweet.NRetweets)
}

func TestGetRetweetsByTweetID(t *testing.T) {
	account1 := createRandomAccount(t)
	n := util.RandomInt(1, 10)
	tweet, account1 := createRandomTweetByAccount(t, account1)

	for i := int64(0); i < n; i++ {
		n_retweets_old := tweet.NRetweets
		account2 := createRandomAccount(t)
		retweetArg := CreateRetweetParams{
			UserID:  account2.ID,
			TweetID: tweet.ID,
		}
		retweet, err := testQueries.CreateRetweet(context.Background(), retweetArg)
		require.NoError(t, err)
		require.NotEmpty(t, retweet)
		require.Equal(t, retweetArg.UserID, retweet.UserID)

		tweet, err = testQueries.AddNRetweetsByOne(context.Background(), tweet.ID)
		require.NoError(t, err)
		require.NotEmpty(t, tweet)
		require.Equal(t, n_retweets_old+1, tweet.NRetweets)
	}

	arg := GetRetweetsByTweetIDParams{
		TweetID: tweet.ID,
		Limit:   int32(n),
		Offset:  0,
	}

	retweets, err := testQueries.GetRetweetsByTweetID(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, retweets)
	require.Equal(t, len(retweets), int(arg.Limit))
}

func TestGetRetweetsByUserID(t *testing.T) {
	n := util.RandomInt(1, 10)
	account2 := createRandomAccount(t)

	for i := int64(0); i < n; i++ {
		account1 := createRandomAccount(t)
		tweet, account1 := createRandomTweetByAccount(t, account1)

		n_retweets_old := tweet.NRetweets

		retweetArg := CreateRetweetParams{
			UserID:  account2.ID,
			TweetID: tweet.ID,
		}

		retweet, err := testQueries.CreateRetweet(context.Background(), retweetArg)
		require.NoError(t, err)
		require.NotEmpty(t, retweet)
		require.Equal(t, retweetArg.UserID, retweet.UserID)

		tweet, err = testQueries.AddNRetweetsByOne(context.Background(), tweet.ID)
		require.NoError(t, err)
		require.NotEmpty(t, tweet)
		require.Equal(t, n_retweets_old+1, tweet.NRetweets)
	}

	arg := GetRetweetsByUserIDParams{
		UserID: account2.ID,
		Limit:  int32(n),
		Offset: 0,
	}

	retweets, err := testQueries.GetRetweetsByUserID(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, retweets)
	require.Equal(t, len(retweets), int(arg.Limit))
}
