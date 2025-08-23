package subscribers

import (
	"github.com/hophuoc14/packlib/pkg/hook"
)

type UserSubscriber struct {
}

func (s UserSubscriber) Created(payload hook.EventPayload) {

}

func (s UserSubscriber) Deleted(payload hook.EventPayload) {

}

func (s UserSubscriber) Updated(payload hook.EventPayload) {

}
