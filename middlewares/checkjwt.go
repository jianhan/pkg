package middlewares

import (
	"errors"
	"net/http"
	"strings"

	auth0 "github.com/auth0-community/go-auth0"
	h "github.com/jianhan/pkg/http"
	"github.com/urfave/negroni"
	jose "gopkg.in/square/go-jose.v2"
	"gopkg.in/square/go-jose.v2/jwt"
)

// JwtRequestValidatorScopeChecker contains methods for JWT validation.
type JwtRequestValidatorScopeChecker interface {
	ValidateRequest(r *http.Request) error
	CheckScope(r *http.Request) error
}

type auth0ValidatorScopeChecker struct {
	domain       string
	clientID     string
	clientSecret string
	jwtValidator *auth0.JWTValidator
	token        *jwt.JSONWebToken
}

// NewJWTRequestValidatorScopeChecker creates instance of JwtRequestValidatorScopeChecker.
func NewJWTRequestValidatorScopeChecker(domain, clientID, clientSecret string, audiences []string) (JwtRequestValidatorScopeChecker, error) {
	// start validation for constructor
	if strings.TrimSpace(domain) == "" {
		return nil, errors.New("Domain can not be empty")
	}
	if strings.TrimSpace(clientID) == "" {
		return nil, errors.New("Client ID can not be empty")
	}
	if strings.TrimSpace(clientSecret) == "" {
		return nil, errors.New("Client secret can not be empty")
	}
	if len(audiences) == 0 {
		return nil, errors.New("Audiences can not be empty")
	}
	// start build struct
	jwkSURI := "https://" + domain + "/.well-known/jwks.json"
	client := auth0.NewJWKClient(auth0.JWKClientOptions{URI: jwkSURI})
	apiIssuer := "https://" + domain + "/"
	validator := auth0.NewValidator(auth0.NewConfiguration(client, audiences, apiIssuer, jose.RS256))
	return &auth0ValidatorScopeChecker{
		domain:       domain,
		clientID:     clientID,
		clientSecret: clientSecret,
		jwtValidator: validator,
	}, nil
}

// ValidateRequest performs validation for request.
func (a *auth0ValidatorScopeChecker) ValidateRequest(r *http.Request) error {
	token, err := a.jwtValidator.ValidateRequest(r)
	if err != nil {
		return err
	}
	a.token = token
	return nil
}

// CheckScope performs scope checking for JWT token.
func (a *auth0ValidatorScopeChecker) CheckScope(r *http.Request) error {
	claims := map[string]interface{}{}
	if a.jwtValidator == nil {
		return errors.New("jwtValidator is nil")
	}
	if a.token == nil {
		return errors.New("token is nil, please validate request first, then check scope")
	}
	err := a.jwtValidator.Claims(r, a.token, &claims)
	if err != nil {
		return err
	}
	// TODO: scope not setup just yet
	// if claims["scope"] != nil && strings.Contains(claims["scope"].(string), "read:messages") {
	// 	return nil
	// }
	if claims["scope"] != nil {
		return nil
	}
	return errors.New("Invalid scope")
}

// CheckJWTMiddleware is the middleware which will perform jwt validation
func CheckJWTMiddleware(jvs JwtRequestValidatorScopeChecker) negroni.HandlerFunc {
	return negroni.HandlerFunc(func(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
		err := jvs.ValidateRequest(r)
		if err != nil {
			h.SendJSONResponse(rw, http.StatusUnauthorized, h.NewResponseData(http.StatusUnauthorized, err.Error(), nil))
			return
		}
		err = jvs.CheckScope(r)
		if err != nil {
			h.SendJSONResponse(rw, http.StatusUnauthorized, h.NewResponseData(http.StatusUnauthorized, err.Error(), nil))
			return
		}
		next(rw, r)
	})
}
