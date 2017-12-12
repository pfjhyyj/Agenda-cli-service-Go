package entities_test

import (
	"entities"
	"go/build"
	"runtime"
	"testing"
)

func init() {
	if runtime.GOOS == "windows" {
		entities.Init(build.Default.GOPATH + "\\tmp\\test_agenda.db")
	} else {
		entities.Init("/tmp/test_agenda.db")
	}
}

func TestSession(t *testing.T) {
	t.Log("[sessiontest] adding session")
	openid := "testtest"
	entities.SessionServ.Add(&entities.Session{
		Openid:   openid,
		Username: "test",
	})
	t.Log("[sessiontest] finding session")
	if entities.SessionServ.FindByOpenid(openid) == nil {
		t.Fatalf("could not find session '%s' that was just added", openid)
	}
	t.Log("[sessiontest] deleting session")
	entities.SessionServ.Delete(openid)
}
