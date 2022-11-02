package helper

import (
	"log"
	"strings"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

var helper = InitHelper()

func TestMain(m *testing.M) {
	// before
	log.Print("before")
	m.Run()
	// after
	log.Print("after")
}

func TestJwtEncryptSession(t *testing.T) {
	token, err := helper.JwtEncryptSession(uuid.New().String(), 60)
	require.Nil(t, err)
	require.Equal(t, 3, len(strings.Split(token, ".")))
}

// test if id match
func TestJwtDecryptSession(t *testing.T) {
	id := uuid.New().String()
	tokenStr, err := helper.JwtEncryptSession(id, 60)
	require.Nil(t, err)
	require.Equal(t, 3, len(strings.Split(tokenStr, ".")))
	userIdentity, err := helper.JwtDecryptSession(tokenStr)
	require.Nil(t, err)
	require.Equal(t, id, userIdentity["user_id"])
}
