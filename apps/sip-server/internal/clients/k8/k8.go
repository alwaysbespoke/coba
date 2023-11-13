package k8

import (
	"fmt"
	"time"

	"go.uber.org/zap"
	"k8s.io/client-go/tools/clientcmd"

	clientset "github.com/alwaysbespoke/coba/pkg/crds/generated/clientset/versioned"
	sbcclientsetv1 "github.com/alwaysbespoke/coba/pkg/crds/generated/clientset/versioned/typed/sbc/v1"
	sipclientsetv1 "github.com/alwaysbespoke/coba/pkg/crds/generated/clientset/versioned/typed/sip/v1"
	informers "github.com/alwaysbespoke/coba/pkg/crds/generated/informers/externalversions"
	sbcinformerv1 "github.com/alwaysbespoke/coba/pkg/crds/generated/informers/externalversions/sbc/v1"
)

type Config struct {
	MasterURL     string `envconfig:"master_url"`
	Kubeconfig    string
	ResyncSeconds int `default:"30"`
}

type Clients struct {
	config      *Config
	logger      *zap.SugaredLogger
	SipV1Client sipclientsetv1.SipV1Interface
	SbcV1Client sbcclientsetv1.SbcV1Interface
	SbcInformer sbcinformerv1.SBCInformer
}

func New(config *Config, logger *zap.SugaredLogger) *Clients {
	cfg, err := clientcmd.BuildConfigFromFlags(config.MasterURL, config.Kubeconfig)
	if err != nil {
		panic(fmt.Errorf("failed to build config form flags: %w", err))
	}

	cs, err := clientset.NewForConfig(cfg)
	if err != nil {
		panic(fmt.Errorf("failed to create sip clientset: %w", err))
	}

	informerFactory := informers.NewSharedInformerFactory(cs, time.Second*time.Duration(config.ResyncSeconds))

	sbcInformer := informerFactory.Sbc().V1().SBCs()

	return &Clients{
		config:      config,
		logger:      logger,
		SipV1Client: cs.SipV1(),
		SbcV1Client: cs.SbcV1(),
		SbcInformer: sbcInformer,
	}
}
