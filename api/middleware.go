package api

type MiddlewareHanlder interface {
	VerifyJwt(userCredentials any)
}
