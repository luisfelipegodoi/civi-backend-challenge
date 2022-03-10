package models

import "errors"

var (
	XParamInvalidOrNullError        = errors.New("The x param is required for the request. Please verify if exists and is integer.")
	YParamInvalidOrNullError        = errors.New("The y param is required for the request. Please verify if exists and is integer.")
	DistanceParamInvalidOrNullError = errors.New("The x param is required for the request. Please verify if exists and is integer.")
)

type Error struct {
	StatusCode int    `json:"StatusCode"`
	Message    string `json:"Message"`
	Details    string `json:"Details"`
}
