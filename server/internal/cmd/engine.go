package cmd

import (
	"github.com/douyu/jupiter"
	"sync"
)

var Default = DefaultEngine()

type Engine struct {
	jupiter.Application
	initOnce sync.Once
}

func DefaultEngine() *Engine {
	app := &Engine{}
	app.initialize()
	return app
}

// initialize application
func (app *Engine) initialize() {
	app.initOnce.Do(func() {

	})
}
