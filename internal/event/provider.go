package event

import (
	"github.com/go-redis/redis"
	"github.com/mylxsw/asteria/log"
	"github.com/mylxsw/glacier/event"
	"github.com/mylxsw/glacier/infra"
	"github.com/mylxsw/redis-event-store/store"
)

type Provider struct{}

func (s Provider) Aggregates() []infra.Provider {
	return []infra.Provider{
		event.Provider(s.eventHandler, event.SetStoreOption(func(cc infra.Resolver) event.Store {
			evtStore := store.NewEventStore(cc.MustGet((*redis.Client)(nil)).(*redis.Client), "default", log.Module("event"))
			//return event.NewMemoryEventStore(true, 100)
			evtStore.Register(SystemUpDownEvent{})
			return evtStore
		})),
	}
}

func (s Provider) Register(cc infra.Binder) {}
func (s Provider) Boot(cc infra.Resolver)   {}

func (s Provider) eventHandler(cc infra.Resolver, listener event.Listener) {
	listener.Listen(func(evt SystemUpDownEvent) {
		log.Debugf("new event received: %v", evt)
	})
}