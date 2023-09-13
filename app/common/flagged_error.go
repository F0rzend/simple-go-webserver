package common

import "errors"

type Flag string

// I think the statuses suggested in GRPC are very versatile and easy to use.
// Based on them, you can give a fairly clear error message
// regardless of the type of interface, be it http or grpc.
// Therefore, these flags should correspond to grpc error status codes.
// This solution will make it easier to switch to grpc if necessary.
//
// See https://github.com/grpc/grpc/blob/master/doc/statuscodes.md
const (
	FlagInvalidArgument Flag = "INVALID_ARGUMENT"
	FlagNotFound        Flag = "NOT_FOUND"
)

type FlaggedError interface {
	error
	Flag() Flag
}

func NewFlaggedError(msg string, flag Flag) error {
	return fault{error: errors.New(msg), flag: flag}
}

func FlagError(err error, flag Flag) error {
	return fault{error: err, flag: flag}
}

func IsFlaggedError(err error, flag Flag) bool {
	var flagged FlaggedError
	if errors.As(err, &flagged) {
		return flagged.Flag() == flag
	}

	return false
}

type fault struct {
	error
	flag Flag
}

func (e fault) Unwrap() error {
	return e.error
}

func (e fault) Flag() Flag {
	return e.flag
}
