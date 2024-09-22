package tools

import (
	"go-oapi-test/util"

	"github.com/lestrrat-go/jwx/jwa"
	"github.com/lestrrat-go/jwx/jwt"
)

type Authenticator struct {
	Config util.Config
}

var _ JWSValidator = (*Authenticator)(nil)

// NewJwsAuthenticator creates an authenticator example which uses a hard coded
// ECDSA key to validate JWT's that it has signed itself.
func NewJwsAuthenticator(config util.Config) (*Authenticator, error) {
	return &Authenticator{Config: config}, nil
}

// ValidateJWS ensures that the critical JWT claims needed to ensure that we
// trust the JWT are present and with the correct values.
func (f *Authenticator) ValidateJWS(jwsString string) (jwt.Token, error) {
	return jwt.Parse([]byte(jwsString), jwt.WithVerify(jwa.HS256, []byte(f.Config.JwtSecret)),
		jwt.WithAudience(f.Config.JwtAudience), jwt.WithIssuer(f.Config.JwtIssuer))
}
