package controller

import (
	"github.com/adambkaplan/plaster/pkg/controller/plasterproject"
)

func init() {
	// AddToManagerFuncs is a list of functions to create controllers and add them to a manager.
	AddToManagerFuncs = append(AddToManagerFuncs, plasterproject.Add)
}
