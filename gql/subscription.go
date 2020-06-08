package gql

import (
	"context"
	"log"

	"github.com/graph-gophers/graphql-go"

	"github.com/dustin/go-broadcast"
)

type Subscription struct {
	datasets            *datasetRegistry
	datasetsBroadcaster broadcast.Broadcaster
}

func newSubscription(registry *datasetRegistry) Subscription {
	b := broadcast.NewBroadcaster(10)
	s := Subscription{datasets: registry, datasetsBroadcaster: b}
	return s
}

func (s *Subscription) DatasetUpdated(ctx context.Context) <-chan *datasetEventResolver {
	c := make(chan *datasetEventResolver)
	ic := make(chan interface{})
	s.datasetsBroadcaster.Register(ic)
	go func() {
		for {
			select {
			case d := <-ic:
				{
					log.Println("Send push")
					c <- d.(*datasetEventResolver)
				}
			case <-ctx.Done():
				{
					s.datasetsBroadcaster.Unregister(ic)
					return
				}
			}
		}
	}()
	return c
}

func (s *Subscription) NotifyDatasetAdded(r *datasetResolver) {
	s.datasetsBroadcaster.Submit(&datasetEventResolver{
		eventType: "added",
		dataset:   r,
	})
}

func (s *Subscription) NotifyDatasetRemoved(id graphql.ID) {
	s.datasetsBroadcaster.Submit(&datasetEventResolver{
		eventType:  "removed",
		idToRemove: &id,
	})
}

type datasetEventResolver struct {
	eventType  string
	idToRemove *graphql.ID
	dataset    *datasetResolver
}

func (r datasetEventResolver) EventType() string {
	return r.eventType
}

func (r datasetEventResolver) IdToRemove() *graphql.ID {
	return r.idToRemove
}

func (r datasetEventResolver) Dataset() *datasetResolver {
	return r.dataset
}
