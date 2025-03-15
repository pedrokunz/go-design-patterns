package event

type Observer interface {
	Update(event Event)
}

type Notifier interface {
	Attach(observer Observer)
	Notify(event Event)
}

type Subject struct {
	observers []Observer
}

func (subject *Subject) Attach(observer Observer) {
	subject.observers = append(subject.observers, observer)
}

func (subject *Subject) Notify(event Event) {
	for _, observer := range subject.observers {
		observer.Update(event)
	}
}
