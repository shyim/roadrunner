package container

import (
	cache "github.com/darkweak/souin/plugins/roadrunner"
	"github.com/roadrunner-server/amqp/v3"
	"github.com/roadrunner-server/beanstalk/v3"
	"github.com/roadrunner-server/gzip/v3"
	httpPlugin "github.com/roadrunner-server/http/v3"
	"github.com/roadrunner-server/informer/v3"
	"github.com/roadrunner-server/jobs/v3"
	"github.com/roadrunner-server/kafka/v3"
	"github.com/roadrunner-server/logger/v3"
	"github.com/roadrunner-server/metrics/v3"
	"github.com/roadrunner-server/nats/v3"
	"github.com/roadrunner-server/reload/v3"
	"github.com/roadrunner-server/resetter/v3"
	rpcPlugin "github.com/roadrunner-server/rpc/v3"
	"github.com/roadrunner-server/server/v3"
	"github.com/roadrunner-server/service/v3"
	"github.com/roadrunner-server/sqs/v3"
	rrt "github.com/temporalio/roadrunner-temporal/v2"
)

// Plugins returns active plugins for the endure container. Feel free to add or remove any plugins.
func Plugins() []any { //nolint:funlen
	return []any{
		// bundled
		// informer plugin (./rr workers, ./rr workers -i)
		&informer.Plugin{},
		// resetter plugin (./rr reset)
		&resetter.Plugin{},
		//
		// logger plugin
		&logger.Plugin{},
		// metrics plugin
		&metrics.Plugin{},
		// reload plugin
		&reload.Plugin{},
		// rpc plugin (workers, reset)
		&rpcPlugin.Plugin{},
		// server plugin (NewWorker, NewWorkerPool)
		&server.Plugin{},
		// service plugin
		&service.Plugin{},
		//
		// ========= JOBS bundle
		&jobs.Plugin{},
		&amqp.Plugin{},
		&sqs.Plugin{},
		&nats.Plugin{},
		&beanstalk.Plugin{},
		// new in 2.11
		&kafka.Plugin{},
		//// =========
		//
		// http server plugin with middleware
		&httpPlugin.Plugin{},
		//&static.Plugin{},
		//&headers.Plugin{},
		//&status.Plugin{},
		&gzip.Plugin{},
		//&prometheus.Plugin{},
		// third-party--
		&cache.Plugin{},
		// --
		//&send.Plugin{},
		//&proxyIP.Plugin{},
		//&rrOtel.Plugin{},
		// --
		//&fileserver.Plugin{},
		// ===================

		//&grpcPlugin.Plugin{},
		// kv + ws + jobs plugin
		//&memory.Plugin{},
		// KV + Jobs
		//&boltdb.Plugin{},
		// broadcast via memory or redis
		// used in conjunction with Websockets, memory and redis plugins
		//&broadcast.Plugin{},
		// ======== websockets broadcast bundle
		//&websockets.Plugin{},
		//&redis.Plugin{},
		// =========

		// ============== KV
		//&kv.Plugin{},
		//&memcached.Plugin{},
		// ==============

		// raw TCP connections handling
		//&tcp.Plugin{},

		// temporal plugins
		&rrt.Plugin{},
	}
}
