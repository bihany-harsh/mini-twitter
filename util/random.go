package util

import (
	"database/sql"
	"math/rand"
	"strings"
	"time"
)

const (
	alphabet = "abcdefghijklmnopqrstuvwxyz"
)

func init() {
	rand.New(rand.NewSource(time.Now().UnixNano()))
}

func RandomInt(min, max int64) int64 {
	return min + rand.Int63n(max-min+1)
}

func RandomString(n int64) string {
	var sb strings.Builder
	k := len(alphabet)
	for i := int64(0); i < n; i++ {
		c := alphabet[rand.Intn(k)]
		sb.WriteByte(c)
	}
	return sb.String()
}

func RandomUsername() string {
	return RandomString(10)
}

func RandomEmail() string {
	return RandomString(6) + "@" + RandomString(4) + ".com"
}

func RandomProfilePictureUrl(p_empty float32, n int64) sql.NullString {
	if rand.Float32() < p_empty {
		return sql.NullString{
			Valid: false,
		}
	}

	url := "http://example.com/" + RandomString(n) + ".png"
	return sql.NullString{
		String: url,
		Valid:  true,
	}
}

func RandomBio(p_empty float32, n int64) sql.NullString {
	if rand.Float32() < p_empty {
		return sql.NullString{
			Valid: false,
		}
	}

	bio := RandomString(n)
	return sql.NullString{
		String: bio,
		Valid:  true,
	}
}

func RandomTime_Nullable(p_empty float32) sql.NullTime {
	if rand.Float32() < p_empty {
		return sql.NullTime{
			Valid: false,
		}
	}

	return sql.NullTime{
		Time:  time.Now().UTC(),
		Valid: true,
	}
}

func RandomBool() bool {
	return rand.Float32() < 0.5
}
