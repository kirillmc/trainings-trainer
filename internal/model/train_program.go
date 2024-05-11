package model

import (
	"github.com/kirillmc/platform_common/pkg/nillable"
)

type TrainProgram struct {
	Id          int64
	ProgramName string
	UserId      int64
	Description string
	Status      Status
}

type TrainProgramToCreate struct {
	ProgramName string
	UserId      int64
	Description string
	Status      Status
}

type TrainProgramToUpdate struct {
	Id          int64
	ProgramName nillable.NilString
	Description nillable.NilString
	Status      Status
}

type Status int32

const (
	Unknown  = 0
	Local    = 1
	Modering = 2
	Public   = 3
)
