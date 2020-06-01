package gql

import (
	"context"
	"log"

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

func (s *Subscription) DatasetsChanged(ctx context.Context) <-chan []*datasetResolver {
	c := make(chan []*datasetResolver)
	ic := make(chan interface{})
	s.datasetsBroadcaster.Register(ic)
	go func() {
		for {
			select {
			case d := <-ic:
				{
					log.Println("Send push")
					c <- d.([]*datasetResolver)
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

func (s *Subscription) DatasetsUpdated() {
	datasets := s.datasets.GetAll()
	datasetResolvers := make([]*datasetResolver, 0, len(datasets))
	for _, dataset := range datasets {
		datasetResolvers = append(datasetResolvers, &datasetResolver{dataset: dataset})
	}
	s.datasetsBroadcaster.Submit(datasetResolvers)
}
