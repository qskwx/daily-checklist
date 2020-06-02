package session

import "time"

type Sessions struct {
	sessions map[string]Session
}

func SessionsFabric() Sessions {
	return Sessions{make(map[string]Session)}
}

func (sessions Sessions) Session(username string, currentTime time.Time) (ss Session, err error) {
	ss, ok := sessions.sessions[username]
	if !ok || !ss.IsActual(currentTime) {
		ss, err := SessionFabric(username)
		if err != nil {
			return ss, err
		}
		sessions.sessions[username] = ss
	}
	return sessions.sessions[username], nil
}
