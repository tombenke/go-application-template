package webserver

import (
	"context"
	"sync"
	"time"
)

type sessionContextKey string

const sessionIDKey sessionContextKey = "session-id"

type Session struct {
	ID        string
	Data      map[string]string
	ExpiresAt time.Time
}

type SessionStore struct {
	mu       sync.RWMutex
	ttl      time.Duration
	sessions map[string]*Session
}

func NewSessionStore(ttl time.Duration) *SessionStore {
	return &SessionStore{
		ttl:      ttl,
		sessions: map[string]*Session{},
	}
}

func (s *SessionStore) GetOrCreate(id string) *Session {
	now := time.Now()
	s.mu.Lock()
	defer s.mu.Unlock()

	if session, ok := s.sessions[id]; ok && session.ExpiresAt.After(now) {
		session.ExpiresAt = now.Add(s.ttl)
		return session
	}

	newSession := &Session{
		ID:        id,
		Data:      map[string]string{},
		ExpiresAt: now.Add(s.ttl),
	}
	s.sessions[id] = newSession
	s.cleanupExpiredLocked(now)

	return newSession
}

func (s *SessionStore) cleanupExpiredLocked(now time.Time) {
	for sessionID, session := range s.sessions {
		if session.ExpiresAt.Before(now) {
			delete(s.sessions, sessionID)
		}
	}
}

// func (ws *WebServer) withSession(next http.Handler) http.Handler {
// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		sessionID := ""
// 		cookie, err := r.Cookie("propel-session-id")
// 		if err == nil {
// 			sessionID = cookie.Value
// 		}

// 		if sessionID == "" {
// 			sessionID = uuid.NewString()
// 			http.SetCookie(w, &http.Cookie{
// 				Name:     "propel-session-id",
// 				Value:    sessionID,
// 				Path:     "/",
// 				HttpOnly: true,
// 				SameSite: http.SameSiteLaxMode,
// 			})
// 		}

// 		_ = ws.sessions.GetOrCreate(sessionID)
// 		ctx := context.WithValue(r.Context(), sessionIDKey, sessionID)
// 		next.ServeHTTP(w, r.WithContext(ctx))
// 	})
// }

func SessionIDFromContext(ctx context.Context) string {
	sessionID, ok := ctx.Value(sessionIDKey).(string)
	if !ok {
		return ""
	}
	return sessionID
}
