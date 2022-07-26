package srv

import (
	"github.com/zeromicro/go-zero/zrpc"
	"miya/api/internal/svc"
)

type hub struct {
	User zrpc.Client
}

var (
	_hub hub
)

func BuildAndSetupSrvHub(ctx *svc.ServiceContext) {
	_hub = hub{
		User: zrpc.MustNewClient(zrpc.NewEtcdClientConf(
			ctx.Config.UserRpc.Etcd.Hosts, "user.rpc", "", "")),
	}
}

// Must call BuildAndSetupSrvClients before
func Hub() hub {
	return _hub
}
