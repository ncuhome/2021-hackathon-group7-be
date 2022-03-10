package model

import "time"

type Limiter struct {
	m map[string]int
}

func (s *Limiter) Init(interval time.Duration) {
	s.m = make(map[string]int)
	go func() {
		for {
			time.Sleep(interval)
			s.m = make(map[string]int)
		}
	}()
	return
}

// limit应该大于0，记录次数超过limit返回false
func (s *Limiter) LogAndCheck(key string, limit int) bool {
	v, ok := s.m[key]
	if ok {
		s.m[key]++
		return v < limit
	} else {
		s.m[key] = 1
	}
	return true
}
