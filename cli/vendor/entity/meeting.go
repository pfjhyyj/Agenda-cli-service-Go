package entity

import (
	"utils"
)

// Meeting model for one meeting
type Meeting struct {
	Title         string   `json:"tile"`
	Speecher      string   `json:"speecher"`
	Participators []string `json:"participators"`
	StartTime     string   `json:"startTime"`
	EndTime       string   `json:"endTime"`
}

// AddMeeting add a new meeting to database
func AddMeeting(meeting *Meeting, err error) {
	logger.Println("[meetingmodel] try adding new meeting", meeting.Title)
	var code int
	var resBody struct {
		Msg string `json:"msg"`
	}
	if code, err = request("POST", "/api/meetings", meeting, &resBody); err != nil {
		return
	}
	err = utils.HTTPErrorHandler(code, resBody.Msg)
	return
}

// AddParticipatorToMeeting Add a Participator To Meeting
func AddParticipatorToMeeting(title string, participator string, err error) {
	logger.Println("[meetingmodel] try adding a participator to meeting", title)
	var code int
	var reqBody struct {
		Participators string `json:"participators"`
	}
	reqBody.Participators = participator
	var resBody struct {
		Msg string `json:"msg"`
	}

	if code, err = request("POST", "/api/meetings/"+title+"/participators", reqBody, &resBody); err != nil {
		return
	}
	err = utils.HTTPErrorHandler(code, resBody.Msg)
	return
}

// FindMeetingsByTime find meetings by time interval
func FindMeetingsByTime(startTime string, endTime string, err error) []Meeting {
	var code int
	var resBody []Meeting
	if code, err = request("GET", "/api/meetings"+"?startTime="+startTime+"&endTime="+endTime,
		nil, &resBody); err != nil {
		return nil
	}
	err = utils.HTTPErrorHandler(code, "")
	if err != nil {
		return nil
	}
	return resBody
}

// DeleteMeeting delete an existed meeting
func DeleteMeeting(title string, err error) {
	logger.Println("[meetingmodel] try deleting a meeting", title)
	var code int
	var resBody struct {
		Msg string `json:"msg"`
	}

	if code, err = request("DELETE", "/api/meetings/"+title, nil, &resBody); err != nil {
		return
	}
	err = utils.HTTPErrorHandler(code, resBody.Msg)
	return
}

// DeleteParticipatorFromMeeting delete a participator from meeting
func DeleteParticipatorFromMeeting(title string, participator string, err error) {
	logger.Println("[meetingmodel] try deleting a participator from meeting", title)
	var reqBody struct {
		Participators string `json:"participators"`
	}
	reqBody.Participators = participator
	var code int
	var resBody struct {
		Msg string `json:"msg"`
	}

	if code, err = request("DELETE", "/api/meetings/"+title+"/participators", reqBody, &resBody); err != nil {
		return
	}
	err = utils.HTTPErrorHandler(code, resBody.Msg)
	return
}
