package rancher2

import (
	"github.com/hashicorp/terraform/helper/schema"
	managementClient "github.com/rancher/types/client/management/v3"
)

const (
	cloudProviderVsphereName = "vsphere"
)

//Schemas

func vsphereDiskCloudProviderFields() map[string]*schema.Schema {
	s := map[string]*schema.Schema{
		"scsi_controller_type": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
	}
	return s
}

func vsphereGlobalCloudProviderFields() map[string]*schema.Schema {
	s := map[string]*schema.Schema{
		"datacenters": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"insecure_flag": {
			Type:     schema.TypeBool,
			Optional: true,
			Computed: true,
		},
		"password": {
			Type:      schema.TypeString,
			Optional:  true,
			Computed:  true,
			Sensitive: true,
		},
		"user": {
			Type:      schema.TypeString,
			Optional:  true,
			Computed:  true,
			Sensitive: true,
		},
		"port": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"soap_roundtrip_count": {
			Type:     schema.TypeInt,
			Optional: true,
			Computed: true,
		},
	}
	return s
}

func vsphereNetworkCloudProviderFields() map[string]*schema.Schema {
	s := map[string]*schema.Schema{
		"public_network": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
	}
	return s
}

func vsphereVirtualCenterCloudProviderFields() map[string]*schema.Schema {
	s := map[string]*schema.Schema{
		"datacenters": {
			Type:     schema.TypeString,
			Required: true,
		},
		"name": {
			Type:     schema.TypeString,
			Required: true,
		},
		"password": {
			Type:      schema.TypeString,
			Required:  true,
			Sensitive: true,
		},
		"user": {
			Type:      schema.TypeString,
			Required:  true,
			Sensitive: true,
		},
		"port": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"soap_roundtrip_count": {
			Type:     schema.TypeInt,
			Optional: true,
			Computed: true,
		},
	}
	return s
}

func vsphereWorkspaceCloudProviderFields() map[string]*schema.Schema {
	s := map[string]*schema.Schema{
		"datacenter": {
			Type:     schema.TypeString,
			Required: true,
		},
		"folder": {
			Type:     schema.TypeString,
			Required: true,
		},
		"server": {
			Type:     schema.TypeString,
			Required: true,
		},
		"default_datastore": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"resourcepool_path": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
	}
	return s
}

func vsphereCloudProviderFields() map[string]*schema.Schema {
	s := map[string]*schema.Schema{
		"disk": {
			Type:     schema.TypeList,
			MaxItems: 1,
			Optional: true,
			Computed: true,
			Elem: &schema.Resource{
				Schema: vsphereDiskCloudProviderFields(),
			},
		},
		"global": {
			Type:     schema.TypeList,
			MaxItems: 1,
			Optional: true,
			Computed: true,
			Elem: &schema.Resource{
				Schema: vsphereGlobalCloudProviderFields(),
			},
		},
		"network": {
			Type:     schema.TypeList,
			MaxItems: 1,
			Optional: true,
			Computed: true,
			Elem: &schema.Resource{
				Schema: vsphereNetworkCloudProviderFields(),
			},
		},
		"virtual_center": {
			Type:     schema.TypeList,
			Required: true,
			Elem: &schema.Resource{
				Schema: vsphereVirtualCenterCloudProviderFields(),
			},
		},
		"workspace": {
			Type:     schema.TypeList,
			MaxItems: 1,
			Required: true,
			Elem: &schema.Resource{
				Schema: vsphereWorkspaceCloudProviderFields(),
			},
		},
	}
	return s
}

// Flatteners

func flattenVsphereDiskCloudProvider(in *managementClient.DiskVsphereOpts) ([]interface{}, error) {
	obj := make(map[string]interface{})
	if in == nil {
		return []interface{}{}, nil
	}

	if len(in.SCSIControllerType) > 0 {
		obj["scsi_controller_type"] = in.SCSIControllerType
	}

	return []interface{}{obj}, nil
}

func flattenVsphereGlobalCloudProvider(in *managementClient.GlobalVsphereOpts) ([]interface{}, error) {
	obj := make(map[string]interface{})
	if in == nil {
		return []interface{}{}, nil
	}

	if len(in.Datacenters) > 0 {
		obj["datacenters"] = in.Datacenters
	}

	obj["insecure_flag"] = in.InsecureFlag

	if len(in.Password) > 0 {
		obj["password"] = in.Password
	}

	if len(in.VCenterPort) > 0 {
		obj["port"] = in.VCenterPort
	}

	if len(in.User) > 0 {
		obj["user"] = in.User
	}

	if in.RoundTripperCount > 0 {
		obj["soap_roundtrip_count"] = int(in.RoundTripperCount)
	}

	return []interface{}{obj}, nil
}

func flattenVsphereNetworkCloudProvider(in *managementClient.NetworkVshpereOpts) ([]interface{}, error) {
	obj := make(map[string]interface{})
	if in == nil {
		return []interface{}{}, nil
	}

	if len(in.PublicNetwork) > 0 {
		obj["public_network"] = in.PublicNetwork
	}

	return []interface{}{obj}, nil
}

func flattenVsphereVirtualCenterCloudProvider(in map[string]managementClient.VirtualCenterConfig) ([]interface{}, error) {
	if len(in) == 0 {
		return []interface{}{}, nil
	}

	out := make([]interface{}, len(in))
	i := 0
	for key := range in {
		obj := make(map[string]interface{})
		obj["name"] = key
		if len(in[key].Datacenters) > 0 {
			obj["datacenters"] = in[key].Datacenters
		}

		if len(in[key].Password) > 0 {
			obj["password"] = in[key].Password
		}

		if len(in[key].VCenterPort) > 0 {
			obj["port"] = in[key].VCenterPort
		}

		if len(in[key].User) > 0 {
			obj["user"] = in[key].User
		}

		if in[key].RoundTripperCount > 0 {
			obj["soap_roundtrip_count"] = int(in[key].RoundTripperCount)
		}
		out[i] = obj
		i++
	}

	return out, nil
}

func flattenVsphereWorkspaceCloudProvider(in *managementClient.WorkspaceVsphereOpts) ([]interface{}, error) {
	obj := make(map[string]interface{})
	if in == nil {
		return []interface{}{}, nil
	}

	if len(in.Datacenter) > 0 {
		obj["datacenter"] = in.Datacenter
	}

	if len(in.Folder) > 0 {
		obj["folder"] = in.Folder
	}

	if len(in.VCenterIP) > 0 {
		obj["server"] = in.VCenterIP
	}

	if len(in.DefaultDatastore) > 0 {
		obj["default_datastore"] = in.DefaultDatastore
	}

	if len(in.ResourcePoolPath) > 0 {
		obj["resourcepool_path"] = in.ResourcePoolPath
	}

	return []interface{}{obj}, nil
}

func flattenVsphereCloudProvider(in *managementClient.VsphereCloudProvider) ([]interface{}, error) {
	obj := make(map[string]interface{})
	if in == nil {
		return []interface{}{}, nil
	}

	if in.Disk != nil {
		disk, err := flattenVsphereDiskCloudProvider(in.Disk)
		if err != nil {
			return []interface{}{obj}, err
		}
		obj["disk"] = disk
	}

	if in.Global != nil {
		global, err := flattenVsphereGlobalCloudProvider(in.Global)
		if err != nil {
			return []interface{}{obj}, err
		}
		obj["global"] = global
	}

	if in.Network != nil {
		network, err := flattenVsphereNetworkCloudProvider(in.Network)
		if err != nil {
			return []interface{}{obj}, err
		}
		obj["network"] = network
	}

	if in.VirtualCenter != nil {
		vc, err := flattenVsphereVirtualCenterCloudProvider(in.VirtualCenter)
		if err != nil {
			return []interface{}{obj}, err
		}
		obj["virtual_center"] = vc
	}

	if in.Workspace != nil {
		workspace, err := flattenVsphereWorkspaceCloudProvider(in.Workspace)
		if err != nil {
			return []interface{}{obj}, err
		}
		obj["workspace"] = workspace
	}

	return []interface{}{obj}, nil
}

// Expanders

func expandVsphereDiskCloudProvider(p []interface{}) (*managementClient.DiskVsphereOpts, error) {
	obj := &managementClient.DiskVsphereOpts{}
	if len(p) == 0 || p[0] == nil {
		return obj, nil
	}
	in := p[0].(map[string]interface{})

	if v, ok := in["scsi_controller_type"].(string); ok && len(v) > 0 {
		obj.SCSIControllerType = v
	}

	return obj, nil
}

func expandVsphereGlobalCloudProvider(p []interface{}) (*managementClient.GlobalVsphereOpts, error) {
	obj := &managementClient.GlobalVsphereOpts{}
	if len(p) == 0 || p[0] == nil {
		return obj, nil
	}
	in := p[0].(map[string]interface{})

	if v, ok := in["datacenters"].(string); ok && len(v) > 0 {
		obj.Datacenters = v
	}

	if v, ok := in["insecure_flag"].(bool); ok {
		obj.InsecureFlag = v
	}

	if v, ok := in["password"].(string); ok && len(v) > 0 {
		obj.Password = v
	}

	if v, ok := in["port"].(string); ok && len(v) > 0 {
		obj.VCenterPort = v
	}

	if v, ok := in["user"].(string); ok && len(v) > 0 {
		obj.User = v
	}

	if v, ok := in["soap_roundtrip_count"].(int); ok && v > 0 {
		obj.RoundTripperCount = int64(v)
	}

	return obj, nil
}

func expandVsphereNetworkCloudProvider(p []interface{}) (*managementClient.NetworkVshpereOpts, error) {
	obj := &managementClient.NetworkVshpereOpts{}
	if len(p) == 0 || p[0] == nil {
		return obj, nil
	}
	in := p[0].(map[string]interface{})

	if v, ok := in["public_network"].(string); ok && len(v) > 0 {
		obj.PublicNetwork = v
	}

	return obj, nil
}

func expandVsphereVirtualCenterCloudProvider(p []interface{}) (map[string]managementClient.VirtualCenterConfig, error) {
	if len(p) == 0 || p[0] == nil {
		return map[string]managementClient.VirtualCenterConfig{}, nil
	}

	obj := make(map[string]managementClient.VirtualCenterConfig)

	for i := range p {
		in := p[i].(map[string]interface{})
		aux := managementClient.VirtualCenterConfig{}
		key := in["name"].(string)

		if v, ok := in["datacenters"].(string); ok && len(v) > 0 {
			aux.Datacenters = v
		}

		if v, ok := in["password"].(string); ok && len(v) > 0 {
			aux.Password = v
		}

		if v, ok := in["port"].(string); ok && len(v) > 0 {
			aux.VCenterPort = v
		}

		if v, ok := in["user"].(string); ok && len(v) > 0 {
			aux.User = v
		}

		if v, ok := in["soap_roundtrip_count"].(int); ok && v > 0 {
			aux.RoundTripperCount = int64(v)
		}

		obj[key] = aux
	}

	return obj, nil
}

func expandVsphereWorkspaceCloudProvider(p []interface{}) (*managementClient.WorkspaceVsphereOpts, error) {
	obj := &managementClient.WorkspaceVsphereOpts{}
	if len(p) == 0 || p[0] == nil {
		return obj, nil
	}
	in := p[0].(map[string]interface{})

	if v, ok := in["datacenter"].(string); ok && len(v) > 0 {
		obj.Datacenter = v
	}

	if v, ok := in["folder"].(string); ok && len(v) > 0 {
		obj.Folder = v
	}

	if v, ok := in["server"].(string); ok && len(v) > 0 {
		obj.VCenterIP = v
	}

	if v, ok := in["default_datastore"].(string); ok && len(v) > 0 {
		obj.DefaultDatastore = v
	}

	if v, ok := in["resourcepool_path"].(string); ok && len(v) > 0 {
		obj.ResourcePoolPath = v
	}

	return obj, nil
}

func expandVsphereCloudProvider(p []interface{}) (*managementClient.VsphereCloudProvider, error) {
	obj := &managementClient.VsphereCloudProvider{}
	if len(p) == 0 || p[0] == nil {
		return obj, nil
	}
	in := p[0].(map[string]interface{})

	if v, ok := in["disk"].([]interface{}); ok && len(v) > 0 {
		disk, err := expandVsphereDiskCloudProvider(v)
		if err != nil {
			return obj, err
		}
		obj.Disk = disk
	}

	if v, ok := in["global"].([]interface{}); ok && len(v) > 0 {
		global, err := expandVsphereGlobalCloudProvider(v)
		if err != nil {
			return obj, err
		}
		obj.Global = global
	}

	if v, ok := in["network"].([]interface{}); ok && len(v) > 0 {
		network, err := expandVsphereNetworkCloudProvider(v)
		if err != nil {
			return obj, err
		}
		obj.Network = network
	}

	if v, ok := in["virtual_center"].([]interface{}); ok && len(v) > 0 {
		vc, err := expandVsphereVirtualCenterCloudProvider(v)
		if err != nil {
			return obj, err
		}
		obj.VirtualCenter = vc
	}

	if v, ok := in["workspace"].([]interface{}); ok && len(v) > 0 {
		workspace, err := expandVsphereWorkspaceCloudProvider(v)
		if err != nil {
			return obj, err
		}
		obj.Workspace = workspace
	}

	return obj, nil
}
