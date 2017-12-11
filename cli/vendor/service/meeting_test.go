package service_test

import (
	"service"
	"testing"
)

func TestMeetingService(t *testing.T) {
	var err error
	err = service.Register("test1", "test1", "email@email.com", "12345678912")
	if err != nil {
		t.Fatal(err)
	}
	err = service.Register("test2", "test2", "email@email.com", "12345678912")
	if err != nil {
		t.Fatal(err)
	}
	err = service.Register("test3", "test3", "email@email.com", "12345678912")
	if err != nil {
		t.Fatal(err)
	}

	err = service.AddMeeting("meeting", []string{"test1"}, "2017-01-01/12:00:00", "2017-01-02/12:00:00")
	if err != nil {
		t.Fatal(err)
	}

	err = service.AddParticipatorToMeeting("meeting", []string{"test2"})
	if err != nil {
		t.Fatal(err)
	}

	err = service.DeleteParticipatorFromMeeting("meeting", []string{"test2"})
	if err != nil {
		t.Fatal(err)
	}

	_, err = service.FindMeetingByTime("2017-01-01/12:00:00", "2017-01-01/12:00:00")
	if err != nil {
		t.Fatal(err)
	}
}
