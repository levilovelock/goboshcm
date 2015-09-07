package sessions

import (
	"errors"
	"math/rand"
	"strings"
	"time"

	"github.com/levilovelock/goboshcm/common"
	"github.com/nu7hatch/gouuid"
)

type Session struct {
	Authenticated bool
}

var (
	sessions map[string]*Session
)

func InitialiseSessionsPool() error {
	rand.Seed(int64(time.Now().Nanosecond()))
	sessions = make(map[string]*Session)
	return nil
}

func CreateNewSession() (string, error) {
	sid, err := generateNewSid()
	if err != nil {
		return "", err
	}
	sessions[sid] = new(Session)

	return sid, nil
}

func SessionExists(sid string) bool {
	return sessions[sid] != nil
}

func generateNewSid() (string, error) {
	// Simply appending two uuids together, stripping the dashes
	// and trimming to 40 chars
	sid := ""
	u4, err := uuid.NewV4()
	if err != nil {
		return "", err
	}
	sid += u4.String()
	u4, err = uuid.NewV4()
	if err != nil {
		return "", err
	}
	sid += u4.String()
	sid = strings.Replace(sid, "-", "", -1)

	return sid[0:40], nil
}

func GenerateSessionCreationResponse(p *common.Payload) (string, error) {
	if p.SID == "" {
		return "", errors.New("Error creating session response - no SID for session")
	}
	// Here we manually build up a xml response, because the Marshal function of encoding/xml doesn't give you the option to NOT marshal empty attributes
	return `<body xmlns="http://jabber.org/protocol/httpbind"></body>`, nil
}
