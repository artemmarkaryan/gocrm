package encryption

import (
	"github.com/artemmarkaryan/gocrm/cmd/app/internal"
	"os"
	"testing"
)

func TestEncrypt(t *testing.T) {
	if err := internal.Setup(); err != nil {
		t.Fatal(err.Error())
	}

	t.Logf("%x", Encrypt("password", os.Getenv("SECRET")))
}
