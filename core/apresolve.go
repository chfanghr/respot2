package core

const (
	APFallBack="ap.spotify.com:443"
	APResolveEndpoint="http://apresolve.spotify.com/"
)

type APResolve struct{
	ApList []string `json:"ap_list"`
}