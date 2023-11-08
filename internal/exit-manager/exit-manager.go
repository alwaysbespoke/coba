package exitmanager

import (
	"context"
	"os"
	"os/signal"
	"syscall"
)

const (
	SUCCESS = iota
	FAILURE
)

// Callback ...
type Callback func() error

// ExitManager ...
type ExitManager struct {
	ctx       context.Context
	cancel    context.CancelFunc
	callbacks map[string]Callback
}

var instance *ExitManager

// New returns a new ExitManager instance. ExitManager is a singleton
func New() (*ExitManager, context.Context) {
	// check if an instance exists
	if instance != nil {
		panic("exit instance already exists")
	}

	// create a cancellable context
	ctx, cancel := context.WithCancel(context.Background())

	// create a new manager instance
	instance = &ExitManager{
		ctx:       ctx,
		cancel:    cancel,
		callbacks: make(map[string]Callback),
	}

	return instance, ctx
}

// Run starts listening for OS signals
func (e *ExitManager) Run() {
	// create a buffered channel for OS signals
	sigs := make(chan os.Signal, 2)

	// listen for OS signals
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	// block until a signal is received
	<-sigs

	// gracefully exit the application
	e.Exit(SUCCESS)
}

// Exit receives an error code and begins the process of exiting
// the application
func (e *ExitManager) Exit(code int) {
	e.executeCallbacks()
	e.cancel()
}

// Wait blocks until the done channel is closed
func (e *ExitManager) Wait() {
	e.ctx.Done()
}

// RegisterCallback registers a callback.
// Callbacks are called prior to the app exiting
func (e *ExitManager) RegisterCallback(key string, callback Callback) {
	e.callbacks[key] = callback
}

func (e *ExitManager) executeCallbacks() {
	for _, callback := range e.callbacks {
		callback()
	}
}
