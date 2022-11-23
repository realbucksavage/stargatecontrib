package eureka

import (
	"fmt"
	"strings"

	"github.com/hudl/fargo"
	"github.com/realbucksavage/stargate"
)

type eurekaLister struct {
	conn fargo.EurekaConnection
}

// NewLister returns an implementation of `stargate.ServiceLister` that can query instances from the addresses passed to this function.
// The application name of each registered service is used as a route key, e.g, if an application "SOME-App" is registered with Eureka,
// it will be accessible at /some-app.
func NewLister(address ...string) stargate.ServiceLister {
	return &eurekaLister{conn: fargo.NewConn(address...)}
}

func (e *eurekaLister) List(route string) ([]string, error) {

	route = strings.TrimPrefix(route, "/")
	app, err := e.conn.GetApp(route)
	if err != nil {
		return nil, err
	}

	return e.listInstances(app), nil
}

func (e *eurekaLister) ListAll() (map[string][]string, error) {
	apps, err := e.conn.GetApps()
	if err != nil {
		return nil, err
	}

	routes := make(map[string][]string)
	for _, a := range apps {
		r := fmt.Sprintf("/%s", strings.ToLower(a.Name))
		routes[r] = e.listInstances(a)
	}

	return routes, nil
}

func (e *eurekaLister) listInstances(app *fargo.Application) []string {
	instances := make([]string, 0)

	for _, i := range app.Instances {
		if i.Status == fargo.UP {
			url := fmt.Sprintf(`http://%s:%d`, i.IPAddr, i.Port)
			if i.SecurePortEnabled {
				url = fmt.Sprintf(`https://%s:%d`, i.IPAddr, i.SecurePort)
			}

			instances = append(instances, url)
		}
	}

	return instances
}
