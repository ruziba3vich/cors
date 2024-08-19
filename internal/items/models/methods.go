package models

type (
	Methods string
)

const (
	CREATE   = Methods("POST")
	RETRIEVE = Methods("GET")
	REMOVE   = Methods("DELETE")
)
