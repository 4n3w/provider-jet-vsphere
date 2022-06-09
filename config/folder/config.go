package folder

import (
	tjconfig "github.com/crossplane/terrajet/pkg/config"
)

func Configure(p *tjconfig.Provider) {
	p.AddResourceConfigurator("vsphere_folder", func(r *tjconfig.Resource) {
		r.ShortGroup = "folder"
		r.ExternalName = tjconfig.IdentifierFromProvider
		//r.References["datacenter_id"] = tjconfig.Reference{
		//	Type: "github.com/crossplane-contrib/provider-jet-vsphere/apis/folder/v1alpha1.Datacenter",
		//}
	})
}
