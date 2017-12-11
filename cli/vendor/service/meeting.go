package service

import (
	"entity"
	"fmt"
	"time"
)

const timeFormat = "2006-01-02/15:04:05"

func validateNewMeeting(meeting *entity.Meeting) error {
	if len(meeting.Title) == 0 {
		return fmt.Errorf("Title should not be empty")
	}
	if len(meeting.Participators) == 0 {
		return fmt.Errorf("Meeting must have participator")
	}

	return nil
}

func validateNewTimeInterval(startTime string, endTime string) error {
	if len(startTime) == 0 {
		return fmt.Errorf("Empty start time")
	}
	// Start Time format must be XXXX-XX-XX/XX:XX
	_, err := time.Parse(timeFormat, startTime)

	if err != nil {
		return fmt.Errorf("Illegal StartTime format")
	}

	if len(endTime) == 0 {
		return fmt.Errorf("Empty end time")
	}
	// End Time format must be XXXX-XX-XX/XX:XX
	_, err = time.Parse(timeFormat, endTime)

	if err != nil {
		return fmt.Errorf("Illegal EndTime format")
	}

	if startTime > endTime {
		return fmt.Errorf("StartTime is after the EndTime")
	}
	return nil
}

// AddMeeting add a meeting
func AddMeeting(title string, participatorName []string, startTime string, endTime string) (err error) {
	// Login first
	if err = checkIfLoggedin(); err != nil {
		return err
	}
	speecherName, err := entity.GetCurUsername()
	if err != nil {
		return
	}
	newMeeting := &entity.Meeting{
		Title:         title,
		Speecher:      speecherName,
		Participators: participatorName,
		StartTime:     startTime,
		EndTime:       endTime,
	}
	// validate if format is ok
	err = validateNewMeeting(newMeeting)
	if err != nil {
		return
	}
	entity.AddMeeting(newMeeting, err)
	return
}

// FindMeetingByTime find the meeting according to the time interval
func FindMeetingByTime(startTime string, endTime string) ([]entity.Meeting, error) {
	if err := checkIfLoggedin(); err != nil {
		return nil, err
	}
	err := validateNewTimeInterval(startTime, endTime)
	if err != nil {
		return nil, err
	}
	results := entity.FindMeetingsByTime(startTime, endTime, err)
	return results, nil
}

// CancelMeeting cancel a meeting
func CancelMeeting(title string) (err error) {
	if err = checkIfLoggedin(); err != nil {
		return err
	}

	// delete the whole meeting
	entity.DeleteMeeting(title, err)
	if err != nil {
		return
	}
	return nil
}

// QuitMeeting quit a meeting
func QuitMeeting(title string) (err error) {
	if err = checkIfLoggedin(); err != nil {
		return err
	}

	username, err := entity.GetCurUsername()
	if err != nil {
		return
	}

	// delete a participator from the meeting
	entity.DeleteParticipatorFromMeeting(title, username, err)
	if err != nil {
		return
	}
	return nil
}

// AddParticipatorToMeeting add participator to a meeting
func AddParticipatorToMeeting(title string, participatorNames []string) (err error) {
	if err = checkIfLoggedin(); err != nil {
		return err
	}

	for _, participatorName := range participatorNames {
		entity.AddParticipatorToMeeting(title, participatorName, err)
		if err != nil {
			return
		}
	}

	return nil
}

// DeleteParticipatorFromMeeting delete a participator from the meeting
func DeleteParticipatorFromMeeting(title string, participatorNames []string) (err error) {
	if err = checkIfLoggedin(); err != nil {
		return err
	}

	for _, participatorName := range participatorNames {
		entity.DeleteParticipatorFromMeeting(title, participatorName, err)
		if err != nil {
			return
		}
	}

	return nil
}
