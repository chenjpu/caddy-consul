package caddyconsul

import (
	"fmt"
	"strings"
	"time"

	"github.com/hashicorp/consul/api"
)

type service struct {
	Name      string
	Instances []*api.CatalogService
}

func (s *caddyfile) WatchServices(reload bool) {

	opts := api.QueryOptions{
		WaitIndex: s.lastService,
		WaitTime:  5 * time.Minute,
	}
	if !reload {
		opts.WaitTime = time.Second
	}
	fmt.Println("Watching for new service with index", s.lastService, "or better")
	services, meta, err := catalog.Services(&opts)
	if err != nil {
		return
	}

	if meta.LastIndex > s.lastService {
		s.lastService = meta.LastIndex
	}

	myservices := make(map[string][]*api.CatalogService)
	for servicename, tags := range services {
		// Get all instances for this service
		if contain(tags, "gateway") {
			instances, _, _ := catalog.Service(servicename, "", nil)

			var domain = strings.TrimRight(servicename, "-gateway")

			//keybits := strings.SplitN(servicename, "/", 3)
			myservices[strings.ToLower(domain)] = instances
		}

	}

	s.services = myservices
	s.buildConfig()
	if reload {
		reloadCaddy()
	}
}

func contain(tags []string, tag string) bool {
	for _, v := range tags {
		if v == tag {
			return true
		}
	}
	return false
}
