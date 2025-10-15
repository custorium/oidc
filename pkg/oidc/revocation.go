package oidc

type RevocationRequest struct {
	Token         string `schema:"token" json:"token"`
	TokenTypeHint string `schema:"token_type_hint" json:"token_type_hint"`
}
