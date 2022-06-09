package datacenter

import (
	"errors"
	tjconfig "github.com/crossplane/terrajet/pkg/config"
)

// Configure configures the null group
func Configure(p *tjconfig.Provider) {
	p.AddResourceConfigurator("vsphere_datacenter", func(r *tjconfig.Resource) {
		r.ShortGroup = "datacenter"
		r.ExternalName.GetExternalNameFn = func(tfstate map[string]interface{}) (string, error) {
			w := tfstate["id"].(string)
			//TODO validate the formatting for a datacenter id
			if len(w) < 1 {
				return "", errors.New("terraform ID should be the datacenter ID (suspiciously short)")
			}
			return w, nil
		}
	})
}
