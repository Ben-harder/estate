package choreManager

import (
	"log"

	"github.com/Ben-harder/estate/household"
	"github.com/Ben-harder/estate/schedule"
)

func NewChoreManager(household household.HouseholdInterface) ChoreManagerInterface {
	chrManager := &choreManager{
		currentChores: make(map[string]ChoreInterface, 0),
		schedules:     make([]schedule.ScheduleInterface, 0),
		household:     household,
	}
	return chrManager
}

type ChoreManagerInterface interface {
	AddSchedule(schedule schedule.ScheduleInterface)
	Chores() []string
	Schedules() []string
	setChore(scheduleName string, responsibilities string, date string, whoseTurn string)
	updateCurrentChores()
}

type choreManager struct {
	currentChores map[string]ChoreInterface
	schedules     []schedule.ScheduleInterface
	household     household.HouseholdInterface
}

// Schedules returns the names of the chore manager's schedules
func (chrManager *choreManager) Schedules() []string {
	names := make([]string, 0)
	for _, schedule := range chrManager.schedules {
		names = append(names, schedule.Name())
	}
	return names
}

// Chores returns the current chores as strings
func (chrManager *choreManager) Chores() []string {
	choreStrings := make([]string, 0)
	for _, chore := range chrManager.currentChores {
		choreStrings = append(choreStrings, chore.String())
	}
	return choreStrings
}

func (chrManager *choreManager) AddSchedule(schedule schedule.ScheduleInterface) {
	chrManager.schedules = append(chrManager.schedules, schedule)
	chrManager.updateCurrentChores()
}

// BuildChores: iterate through schedules, get next jobs which contain the responsibilities and date, then creates a chore for each by attaching name(s) to it
func (chrManager *choreManager) updateCurrentChores() {
	log.Println("checking schedules for chore updates...")
	for _, schedule := range chrManager.schedules {
		responsibilities, date := schedule.NextJob()
		chrManager.setChore(schedule.Name(), responsibilities, date, "")
	}
}

func (chrManager *choreManager) setChore(scheduleName string, responsibilities string, date string, whoseTurn string) {
	newChore := NewChore(scheduleName, responsibilities, date)
	if whoseTurn == "" {
		newChore.SetTurn([]household.MemberInterface{chrManager.household.First()})
	}
	// TODO: allow other people to take a turn instead of the first person in the house
	chrManager.currentChores[scheduleName] = newChore
}
