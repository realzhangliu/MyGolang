package Misc

import (
	"time"
)

const sourceUrl = "https://talks.golang.org/2013/advconc.slide#2"

type Item struct {
	Title, Channel, GUID string
}
type Fetcher interface {
	Fetch() (items []Item, next time.Time, err error)
}

func Fetch(domain string) Fetcher {

}

type Subscription interface {
	Updates() <-chan Item
	Close() error
}

func Subscribe(fetcher Fetcher) Subscription {
	s := &sub{
		fetcher: fetcher,
		updates: make(chan Item),
	}
	s.loop()
	return s
}
func Merge(subs ...Subscription) Subscription {

}

type sub struct {
	fetcher Fetcher
	updates chan Item
}
func (s *sub) Updates() <-chan Item {
	return s.updates
}
func (s *sub) Close() error {
	return err
}
func (s *sub) loop() {
	for{
		if s.closed
	}
}

func Run() {
	merged := Merge()
}
