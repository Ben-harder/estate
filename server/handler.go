package server

import "github.com/Ben-harder/estate/household"

func newHandler(household household.HouseholdInterface) handlerInterface {
	return &handler{household: household}
}

type handlerInterface interface {
	handle()
}

type handler struct {
	household household.HouseholdInterface
}

func (hlr *handler) handle() {

}
