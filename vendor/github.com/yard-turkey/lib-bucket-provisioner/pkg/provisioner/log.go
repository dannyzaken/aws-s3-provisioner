package provisioner

import (
	"github.com/go-logr/logr"
	"k8s.io/klog/klogr"

	"github.com/yard-turkey/lib-bucket-provisioner/pkg/provisioner/api"
)

var (
	// log is a short lived variable that is overwritten each Reconcile() iteration with the current Request injected
	// into it.  This is for convenience of identifying which log lines were generated by which request without having
	// to pass the request to every method call.
	log logr.Logger

	// logD is a short lived variable that is overwritten each Reconcile() iteraion with the current Request injected
	// into it.  This is for convenience of identifying which log lines were generated by which request without having
	// to pass the request to every log method call.
	logD logr.InfoLogger
)

func init() {
	log = klogr.New().WithName(api.Domain + "/claim-reconciler")
	logD = log.V(1)
}

// setLoggerWith request overwrites log and logD with a new logger.  The passed in request is injected into the loggers.
func setLoggersWithRequest(key string) {
	log = klogr.New().WithValues("key", key)
	logD = klogr.New().WithValues("key", key)
}
