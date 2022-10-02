package awesomeProject1

import "fmt"

type jobSite struct {
	subscribers []Observer
	vacancies   []string
}

func newJobSite() *jobSite {
	j := new(jobSite) // because i cant add j in 8th line
	j.subscribers = make([]Observer, 0, 10)
	j.vacancies = make([]string, 0, 10)
	return j
}
func (j *jobSite) addVacancy(vacancy string) {
	j.vacancies = append(j.vacancies, vacancy)
	j.sendAll()
}
func (j *jobSite) removeVacancy(vacancy string) {
	for i := range j.vacancies {
		if j.vacancies[i] == vacancy {
			j.vacancies, _ = removeFromslice(j.vacancies, i)
			break
		}
	}
	j.sendAll()
}
func (j *jobSite) subscribe(o Observer) { // from observable
	j.subscribers = append(j.subscribers, o) // the same as add vacancy but with subscribers
}
func (j *jobSite) unsubscribe(o Observer) { // from observable
	for i := range j.subscribers {
		if j.subscribers[i] == o {
			j.subscribers, _ = removeFromslice(j.subscribers, i)
			break
		}
	}
}
func (j *jobSite) sendAll() { // from observable
	for _, subscribers := range j.subscribers {
		subscribers.handleEvent(j.vacancies)
	}
}
func removeFromslice[Observer any](s []Observer, i int) ([]Observer, error) {
	if i >= len(s) || i < 0 {
		return nil, fmt.Errorf("Something goes wrong with index %d with length %d ", i, len(s))
	}
	s[i] = s[len(s)-1] //remove from slice
	return s[:len(s)-1], nil
}
