package gateway

import (
	"context"
	"log"

	"net/url"
	"net/http"
	"net/http/httputil"
	"sync"

	"kraken/persistence"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

type Gateway struct {
	Queries *persistence.Queries
}

type Service struct {
	Backend string
	Routes []Route
	Proxy *httputil.ReverseProxy
}

type Route struct {
	Path string
}

func (gateway Gateway) getServices(ctx context.Context) ([]Service, error) {
	var err error

	rawServices, err := gateway.Queries.ListServices(ctx)
	services := make([]Service, len(rawServices))

	
	if err != nil {
		return nil, err
	}
	
	for index, rawService := range rawServices {
		services[index] = Service{ Backend: rawService.Backend }

		rawRoutes, err := gateway.Queries.ListRoutesForService(ctx, rawService.ID)
		if err != nil {
			return nil, err
		}

		routes := make([]Route, len(rawRoutes))
		for index, rawRoute := range rawRoutes {
			routes[index] = Route{ Path: rawRoute.Path }
		}

		services[index].Routes = routes
	}

	return services, nil
}

func registerRoutesForService(router *mux.Router, service Service) {
	proxyFunc := func(route Route) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			service.Proxy.ServeHTTP(w, r)
		}
	}

	for _, route := range service.Routes {
		log.Printf(route.Path + " -> " + service.Backend + "\n")
		router.HandleFunc(route.Path, proxyFunc(route))
	}
}

func createProxyForService(service Service) *httputil.ReverseProxy {
	// fmt.Println("Creating proxy for service: ", service.Backend)
		
	target, _ := url.Parse(service.Backend)
	proxy := httputil.NewSingleHostReverseProxy(target)
	
	return proxy
}

func (gateway Gateway) New(ctx context.Context, wg *sync.WaitGroup, bind string) (*http.Server, error) {
	router := mux.NewRouter()

	srv := &http.Server{
		Addr:		bind,
		Handler:	router,
	}

	services, err := gateway.getServices(ctx)
	if err != nil {
		return srv, err
	}

	for _, service := range services {
		service.Proxy = createProxyForService(service)
		registerRoutesForService(router, service)
	}

	go func() {
		defer wg.Done()

		if err := srv.ListenAndServe(); err != http.ErrServerClosed {
			log.Fatalf("ListenAndServe(): %s", err)
		}
	}()
	
	return srv, nil
}