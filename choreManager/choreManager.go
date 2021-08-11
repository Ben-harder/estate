package choreManager

import (
	"github.com/Ben-harder/estate/household"
	"github.com/Ben-harder/estate/schedule"
)

func NewEmptyChoreManager(household household.HouseholdInterface) ChoreManagerInterface {
	chrManager := &choreManager{
		currentChores: make(map[string]ChoreInterface, 0),
		schedules:     make([]schedule.ScheduleInterface, 0),
		household:     household,
	}
	return chrManager
}

type ChoreManagerInterface interface {
	AddSchedule(schedule schedule.ScheduleInterface)
	Schedules() []string
	setChore(scheduleName string, responsibilities string, date string)
	updateCurrentChores()
}

type choreManager struct {
	currentChores map[string]ChoreInterface
	schedules     []schedule.ScheduleInterface
	household     household.HouseholdInterface
}

// Schedules returns the names of the chore manager's schedules
func (chrManager *choreManager) Schedules() []string {
	names := make([]string, len(chrManager.schedules))
	for _, schedule := range chrManager.schedules {
		names = append(names, schedule.Name())
	}
	return names
}

func (chrManager *choreManager) AddSchedule(schedule schedule.ScheduleInterface) {
	chrManager.schedules = append(chrManager.schedules, schedule)
	chrManager.updateCurrentChores()
}

// BuildChores: iterate through schedules, get next jobs which contain the responsibilities and date, then creates a chore for each by attaching name(s) to it
func (chrManager *choreManager) updateCurrentChores() {
	for _, schedule := range chrManager.schedules {
		responsibilities, date := schedule.NextJob()
		chrManager.setChore(schedule.Name(), responsibilities, date)
	}
}

func (chrManager *choreManager) setChore(scheduleName string, responsibilities string, date string) {
	chrManager.currentChores[scheduleName] = NewChore(scheduleName, responsibilities, date)
}
