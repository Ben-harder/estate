package choreManager

import (
	"github.com/Ben-harder/estate/household"
	"github.com/Ben-harder/estate/schedule"
)

func NewEmptyChoreManager(household household.HouseholdInterface) ChoreManagerInterface {
	chrManager := &choreManager{
		currentChores: make([]ChoreInterface, 0),
		schedules:     make([]schedule.ScheduleInterface, 0),
		household:     household,
	}
	return chrManager
}

type ChoreManagerInterface interface {
	AddSchedule(schedule schedule.ScheduleInterface)
	addChore(responsibilities string, date string)
}

type choreManager struct {
	currentChores []ChoreInterface
	schedules     []schedule.ScheduleInterface
	household     household.HouseholdInterface
}

func (chrManager *choreManager) AddSchedule(schedule schedule.ScheduleInterface) {
	chrManager.schedules = append(chrManager.schedules, schedule)
}

// BuildChores: iterate through schedules, get next jobs which contain the responsibilities and date, then creates a chore for each by attaching name(s) to it
func (chrManager *choreManager) BuildCurrentChores() {
	for _, schedule := range chrManager.schedules {
		responsibilities, date, _ := schedule.NextJob()
		chrManager.addChore(responsibilities, date)
	}
}

func (chrManager *choreManager) addChore(responsibilities string, date string) {
	chrManager.currentChores = append(chrManager.currentChores, NewChore(responsibilities, date))
}
