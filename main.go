package awesomeProject1

func main() {
	hhKz := newJobSite()
	bob := newPerson("Bob")
	hhKz.addVacancy("Senior HTML Developer")
	hhKz.subscribe(bob)
	hhKz.addVacancy("Java Junior Developer")
	hhKz.removeVacancy("Senior HTML Developer")
}
