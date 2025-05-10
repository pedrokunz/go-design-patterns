package observer

import (
	"errors"
	"fmt"
	"github.com/pedrokunz/go-design-patterns/internal/domain/event"
)

type Notifier interface {
	Attach(observer Observer) error
	Notify(event event.Event) error
}

type notifier struct {
	observers []Observer
}

func NewNotifier() Notifier {
	return &notifier{
		observers: make([]Observer, 0),
	}
}

func (subject *notifier) Attach(observer Observer) error {
	if observer == nil {
		return errors.New("observer cannot be nil")
	}

	subject.observers = append(subject.observers, observer)

	return nil
}

func (subject *notifier) Notify(event event.Event) (err error) {
	if event == nil {
		return errors.New("event cannot be nil")
	}

	for _, observer := range subject.observers {
		onErr := observer.On(event)
		if onErr != nil {
			// We don't want to stop notifying other observers if one fails
			if err == nil {
				err = onErr
				continue
			}

			err = fmt.Errorf("%w; %w", err, onErr)
		}
	}

	return err
}
