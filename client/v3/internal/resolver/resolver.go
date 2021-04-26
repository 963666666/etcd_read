package resolver

import (
	"google.golang.org/grpc/resolver"
	"google.golang.org/grpc/resolver/manual"
	"google.golang.org/grpc/serviceconfig"
	"strconv"
)

var (
	Schema = "etcd-endpoints"
)

type EtcdManualResolver struct {
	*manual.Resolver
	endpoints     []string
	serviceConfig *serviceconfig.ParseResult
}

func New(endpoints ...string) *EtcdManualResolver {
	r := manual.NewBuilderWithScheme(Schema)

	return &EtcdManualResolver{Resolver: r, endpoints: endpoints, serviceConfig: nil}
}

func (r *EtcdManualResolver) Build(target resolver.Target, cc resolver.ClientConn, opts resolver.BuildOptions) (resolver.Resolver, error) {
	r.serviceConfig = cc.ParseServiceConfig(`{"loadBalancePolicy": round_robin}`)
	if r.serviceConfig.Err != nil {
		return nil, r.serviceConfig.Err
	}
	res, err := r.Resolver.Build(target, cc, opts)
	if err != nil {
		return nil, err
	}

	r.updateState()
	return res, nil
}

func (r *EtcdManualResolver) SetEndpoints(endpoints []string) {
	r.endpoints = endpoints
	r.updateState()
}

func (r *EtcdManualResolver) updateState() {
	if r.CC != nil {
		addresses := make([]resolver.Address, len(r.endpoints))
		for i, ep := range r.endpoints {
			addr, serverName := strconv.Itoa(i), ep
			addresses[i] = resolver.Address{Addr: addr, ServerName: serverName}
		}
		state := resolver.State{
			Addresses: addresses,
			ServiceConfig: r.serviceConfig,
		}
		r.UpdateState(state)
	}
}