package core

import "github.com/disgoorg/disgo/bot"

var listeners []func(ctx *Context) bot.EventListener

func AddListener[E bot.Event](fn func(ctx *Context, e E)) {
	listeners = append(listeners, func(ctx *Context) bot.EventListener {
		return bot.NewListenerFunc(func(e E) {
			fn(ctx, e)
		})
	})
}

func BuildListeners(ctx *Context) []bot.EventListener {
	eventListeners := make([]bot.EventListener, 0, len(listeners))
	for _, build := range listeners {
		eventListeners = append(eventListeners, build(ctx))
	}
	return eventListeners
}
