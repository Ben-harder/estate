package choreManager

import (
	"fmt"
	"log"

	"github.com/Ben-harder/estate/choreManager/chore"
	"github.com/Ben-harder/estate/household"
	"github.com/Ben-harder/estate/household/member"
	"github.com/Ben-harder/estate/schedule"
)

func NewChoreManager(household household.HouseholdInterface) ChoreManagerInterface {
	log.Printf("creating new chore manager for household{%v}", household.String())
	chrManager := &choreManager{
		currentChores: make(map[string]chore.ChoreInterface, 0),
		schedules:     make([]schedule.ScheduleInterface, 0),
		household:     household,
	}
	return chrManager
}

type ChoreManagerInterface interface {
	AddSchedule(schedule schedule.ScheduleInterface, turnList [][]member.MemberInterface, whoStarts int)
	Chores() []chore.ChoreInterface
	Schedules() []string
	DefaultTurnList() [][]member.MemberInterface
	CustomTurnList(listsOfNames ...[]string) ([][]member.MemberInterface, error)
	SetTurnForSchedule(scheduleName string, memberName string) error
	UpdateChoresIfOld()
}

type choreManager struct {
	currentChores map[string]chore.ChoreInterface
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

// DefaultTurnList returns a slice of slices where each slice contains one household member
func (chrManager *choreManager) DefaultTurnList() [][]member.MemberInterface {
	defaultTurnList := make([][]member.MemberInterface, 0)
	for _, m := range chrManager.household.Members() {
		defaultTurnList = append(defaultTurnList, []member.MemberInterface{m})
	}
	return defaultTurnList
}

// CustomTurnList takes sublists of names and returns a 2D slice containing those lists as members
func (chrManager *choreManager) CustomTurnList(listsOfNames ...[]string) ([][]member.MemberInterface, error) {
	customTurnList := make([][]member.MemberInterface, len(listsOfNames))
	for i, listOfNames := range listsOfNames {
		customTurnList[i] = make([]member.MemberInterface, 0)
		for _, name := range listOfNames {
			member, err := chrManager.household.GetHouseholdMember(name)
			if err != nil {
				return nil, err
			}
			customTurnList[i] = append(customTurnList[i], member)
		}
	}
	return customTurnList, nil
}

// Chores returns the current chores as strings
func (chrManager *choreManager) Chores() []chore.ChoreInterface {
	chores := make([]chore.ChoreInterface, 0)
	for _, chore := range chrManager.currentChores {
		chores = append(chores, chore)
	}
	return chores
}

// AddSchedule adds the provided schedule to the chore manager and creates an accompanying chore for the next job
func (chrManager *choreManager) AddSchedule(schedule schedule.ScheduleInterface, turnList [][]member.MemberInterface, whoStarts int) {
	log.Printf("adding new schedule to chore manager. Name: %v", schedule.Name())
	chrManager.schedules = append(chrManager.schedules, schedule)
	responsibilities, date := schedule.NextJob()

	// Create the chore that will correspond to this schedule
	chore := chore.NewChore(schedule.Name(), responsibilities, date, turnList)
	chore.SetTurn(whoStarts)
	log.Printf("initialized chore: schedule{%v}, chore{%v, %v, %v}", schedule.Name(), chore.Responsibilities(), chore.Date(), chore.WhoseTurn())
	chrManager.currentChores[schedule.Name()] = chore
}

// SetTurnForSchedule will take the provided name and set the chore's turn to them or to whichever chore group they're in.
func (chrManager *choreManager) SetTurnForSchedule(scheduleName string, memberName string) error {
	log.Printf("setting turn for schedule, %v, to %v", scheduleName, memberName)
	chore, ok := chrManager.currentChores[scheduleName]
	if !ok {
		return fmt.Errorf("schedule, %v, did not exist", scheduleName)
	}
	member, err := chrManager.household.GetHouseholdMember(memberName)
	if err != nil {
		return err
	}
	err = chore.SetTurnWithMember(member)
	if err != nil {
		return err
	}
	return nil
}

// UpdateChoresIfOld iterate through schedules, and looks for any old jobs. Any old jobs will be removed and turns advanced.
func (chrManager *choreManager) UpdateChoresIfOld() {
	log.Println("checking schedules for chore updates...")
	for _, schedule := range chrManager.schedules {
		chore := chrManager.currentChores[schedule.Name()]
		for schedule.IsNextJobOld() {
			schedule.RemoveNextJob()
			responsibilities, date := schedule.NextJob()
			chore.SetResponsibilities(responsibilities)
			chore.SetDate(date)
			chore.AdvanceToNextTurn()
			log.Printf("chore update: schedule{%v}, chore{%v, %v, %v}", schedule.Name(), chore.Responsibilities(), chore.Date(), chore.WhoseTurn())
		}
	}
}
