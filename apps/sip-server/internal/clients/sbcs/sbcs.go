package sbcs

import (
	"context"
	"errors"
	"fmt"
	"math/rand"
	"net"
	"sync"
	"time"

	"go.uber.org/zap"
	"k8s.io/client-go/tools/cache"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	clientsetv1 "github.com/alwaysbespoke/coba/pkg/crds/generated/clientset/versioned/typed/sbc/v1"
	informerv1 "github.com/alwaysbespoke/coba/pkg/crds/generated/informers/externalversions/sbc/v1"
	v1 "github.com/alwaysbespoke/coba/pkg/crds/sbc/v1"
)

var ErrSbcNotFound error = errors.New("SBC not found")

type Config struct {
	Namespace string
}

type Client interface {
	AssignSBC() *SBC
	GetSBC(id string) (*SBC, bool)
}

type client struct {
	*Input
	sbcsCache map[string]*SBC
	lock      *sync.RWMutex
}

type Input struct {
	Ctx           context.Context
	Config        *Config
	Logger        *zap.SugaredLogger
	SbcV1Client   clientsetv1.SbcV1Interface
	SbcV1Informer informerv1.SBCInformer
}

// New returns a new SBC client
func New(input *Input) *client {
	// list the SBCs
	sbcList, err := input.SbcV1Client.SBCs(input.Config.Namespace).List(input.Ctx, metav1.ListOptions{})
	if err != nil {
		panic(fmt.Errorf("failed to list SBCs: %w", err))
	}

	// create the SBCs map
	sbcsCache := map[string]*SBC{}

	// populate the SBCs map
	for _, sbc := range sbcList.Items {
		// resolve the server address
		addr, err := net.ResolveUDPAddr("udp", sbc.Spec.Address)
		if err != nil {
			input.Logger.Errorf("failed to resolve UDP address (%s): %w", sbc.Spec.Address, err)
			continue
		}

		// connect to the server
		conn, err := net.DialUDP("udp", nil, addr)
		if err != nil {
			input.Logger.Errorf("failed to dial UDP server (%s): %w", sbc.Spec.Address, err)
			continue
		}

		sbcsCache[sbc.Name] = &SBC{
			Obj:  &sbc,
			Conn: conn,
		}
	}

	return &client{
		Input:     input,
		sbcsCache: map[string]*SBC{},
		lock:      &sync.RWMutex{},
	}
}

// RunController creates and runs a minimal SBC controller
// The minimal controller does not use a work queue and lister but rather processes
// objects immediately upon arrival
func (c *client) RunController() {
	// add event handlers
	c.Input.SbcV1Informer.Informer().AddEventHandler(cache.ResourceEventHandlerFuncs{
		AddFunc:    c.addFunc,
		UpdateFunc: c.updateFunc,
		DeleteFunc: c.deleteFunc,
	})

	// start the informer
	go c.Input.SbcV1Informer.Informer().Run(c.Ctx.Done())

	// wait for the caches to sync
	if ok := cache.WaitForCacheSync(c.Ctx.Done(), c.Input.SbcV1Informer.Informer().HasSynced); ok {
		panic("failed to wait for caches to sync")
	}
}

func (c *client) addFunc(obj interface{}) {
	sbc := obj.(*v1.SBC)

	// resolve the server address
	addr, err := net.ResolveUDPAddr("udp", sbc.Spec.Address)
	if err != nil {
		c.Logger.Errorf("failed to resolve UDP address (%s): %w", sbc.Spec.Address, err)
		return
	}

	// connect to the server
	conn, err := net.DialUDP("udp", nil, addr)
	if err != nil {
		c.Logger.Errorf("failed to dial UDP server (%s): %w", sbc.Spec.Address, err)
		return
	}

	c.lock.Lock()
	defer c.lock.Unlock()

	c.sbcsCache[sbc.Name] = &SBC{
		Obj:  sbc,
		Conn: conn,
	}
}

func (c *client) updateFunc(oldObj, newObj interface{}) {
	// todo: implement
}

func (c *client) deleteFunc(obj interface{}) {
	sbc := obj.(*v1.SBC)

	c.lock.Lock()
	defer c.lock.Unlock()

	delete(c.sbcsCache, sbc.Name)
}

// AssignSBC randomly returns an SBC instance from the SBC cache
// todo: improve the algorithm
func (c *client) AssignSBC() *SBC {
	seed := rand.NewSource(time.Now().UnixNano())
	r := rand.New(seed)
	randInt := r.Intn(len(c.sbcsCache))

	i := 0

	c.lock.RLock()
	defer c.lock.RUnlock()

	for _, sbc := range c.sbcsCache {
		if i == randInt {
			return sbc
		}
		i++
	}

	return nil
}

// GetSBC returns an SBC object from the SBCs cache and a bool
func (c *client) GetSBC(id string) (*SBC, bool) {
	c.lock.RLock()
	defer c.lock.RUnlock()

	sbc, ok := c.sbcsCache[id]

	return sbc, ok
}
