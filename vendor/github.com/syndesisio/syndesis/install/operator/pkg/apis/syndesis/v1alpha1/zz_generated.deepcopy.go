// +build !ignore_autogenerated

// Code generated by deepcopy-gen. DO NOT EDIT.

package v1alpha1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in AddonsSpec) DeepCopyInto(out *AddonsSpec) {
	{
		in := &in
		*out = make(AddonsSpec, len(*in))
		for key, val := range *in {
			var outVal map[string]string
			if val == nil {
				(*out)[key] = nil
			} else {
				in, out := &val, &outVal
				*out = make(Parameters, len(*in))
				for key, val := range *in {
					(*out)[key] = val
				}
			}
			(*out)[key] = outVal
		}
		return
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AddonsSpec.
func (in AddonsSpec) DeepCopy() AddonsSpec {
	if in == nil {
		return nil
	}
	out := new(AddonsSpec)
	in.DeepCopyInto(out)
	return *out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ComponentsSpec) DeepCopyInto(out *ComponentsSpec) {
	*out = *in
	in.Server.DeepCopyInto(&out.Server)
	in.Meta.DeepCopyInto(&out.Meta)
	out.UI = in.UI
	out.S2I = in.S2I
	in.Db.DeepCopyInto(&out.Db)
	in.Oauth.DeepCopyInto(&out.Oauth)
	out.PostgresExporter = in.PostgresExporter
	in.Prometheus.DeepCopyInto(&out.Prometheus)
	in.Grafana.DeepCopyInto(&out.Grafana)
	in.Komodo.DeepCopyInto(&out.Komodo)
	out.Upgrade = in.Upgrade
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ComponentsSpec.
func (in *ComponentsSpec) DeepCopy() *ComponentsSpec {
	if in == nil {
		return nil
	}
	out := new(ComponentsSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DbConfiguration) DeepCopyInto(out *DbConfiguration) {
	*out = *in
	in.Resources.DeepCopyInto(&out.Resources)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DbConfiguration.
func (in *DbConfiguration) DeepCopy() *DbConfiguration {
	if in == nil {
		return nil
	}
	out := new(DbConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GrafanaConfiguration) DeepCopyInto(out *GrafanaConfiguration) {
	*out = *in
	in.Resources.DeepCopyInto(&out.Resources)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GrafanaConfiguration.
func (in *GrafanaConfiguration) DeepCopy() *GrafanaConfiguration {
	if in == nil {
		return nil
	}
	out := new(GrafanaConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IntegrationSpec) DeepCopyInto(out *IntegrationSpec) {
	*out = *in
	if in.Limit != nil {
		in, out := &in.Limit, &out.Limit
		*out = new(int)
		**out = **in
	}
	if in.StateCheckInterval != nil {
		in, out := &in.StateCheckInterval, &out.StateCheckInterval
		*out = new(int)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IntegrationSpec.
func (in *IntegrationSpec) DeepCopy() *IntegrationSpec {
	if in == nil {
		return nil
	}
	out := new(IntegrationSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KomodoConfiguration) DeepCopyInto(out *KomodoConfiguration) {
	*out = *in
	in.Resources.DeepCopyInto(&out.Resources)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KomodoConfiguration.
func (in *KomodoConfiguration) DeepCopy() *KomodoConfiguration {
	if in == nil {
		return nil
	}
	out := new(KomodoConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MetaConfiguration) DeepCopyInto(out *MetaConfiguration) {
	*out = *in
	in.Resources.DeepCopyInto(&out.Resources)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MetaConfiguration.
func (in *MetaConfiguration) DeepCopy() *MetaConfiguration {
	if in == nil {
		return nil
	}
	out := new(MetaConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OauthConfiguration) DeepCopyInto(out *OauthConfiguration) {
	*out = *in
	if in.DisableSarCheck != nil {
		in, out := &in.DisableSarCheck, &out.DisableSarCheck
		*out = new(bool)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OauthConfiguration.
func (in *OauthConfiguration) DeepCopy() *OauthConfiguration {
	if in == nil {
		return nil
	}
	out := new(OauthConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in Parameters) DeepCopyInto(out *Parameters) {
	{
		in := &in
		*out = make(Parameters, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
		return
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Parameters.
func (in Parameters) DeepCopy() Parameters {
	if in == nil {
		return nil
	}
	out := new(Parameters)
	in.DeepCopyInto(out)
	return *out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PostgresExporterConfiguration) DeepCopyInto(out *PostgresExporterConfiguration) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PostgresExporterConfiguration.
func (in *PostgresExporterConfiguration) DeepCopy() *PostgresExporterConfiguration {
	if in == nil {
		return nil
	}
	out := new(PostgresExporterConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PrometheusConfiguration) DeepCopyInto(out *PrometheusConfiguration) {
	*out = *in
	in.Resources.DeepCopyInto(&out.Resources)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PrometheusConfiguration.
func (in *PrometheusConfiguration) DeepCopy() *PrometheusConfiguration {
	if in == nil {
		return nil
	}
	out := new(PrometheusConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Resources) DeepCopyInto(out *Resources) {
	*out = *in
	in.ResourceRequirements.DeepCopyInto(&out.ResourceRequirements)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Resources.
func (in *Resources) DeepCopy() *Resources {
	if in == nil {
		return nil
	}
	out := new(Resources)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResourcesWithVolume) DeepCopyInto(out *ResourcesWithVolume) {
	*out = *in
	in.Resources.DeepCopyInto(&out.Resources)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResourcesWithVolume.
func (in *ResourcesWithVolume) DeepCopy() *ResourcesWithVolume {
	if in == nil {
		return nil
	}
	out := new(ResourcesWithVolume)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *S2IConfiguration) DeepCopyInto(out *S2IConfiguration) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new S2IConfiguration.
func (in *S2IConfiguration) DeepCopy() *S2IConfiguration {
	if in == nil {
		return nil
	}
	out := new(S2IConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServerConfiguration) DeepCopyInto(out *ServerConfiguration) {
	*out = *in
	in.Resources.DeepCopyInto(&out.Resources)
	out.Features = in.Features
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServerConfiguration.
func (in *ServerConfiguration) DeepCopy() *ServerConfiguration {
	if in == nil {
		return nil
	}
	out := new(ServerConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServerFeatures) DeepCopyInto(out *ServerFeatures) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServerFeatures.
func (in *ServerFeatures) DeepCopy() *ServerFeatures {
	if in == nil {
		return nil
	}
	out := new(ServerFeatures)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Syndesis) DeepCopyInto(out *Syndesis) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Syndesis.
func (in *Syndesis) DeepCopy() *Syndesis {
	if in == nil {
		return nil
	}
	out := new(Syndesis)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Syndesis) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SyndesisList) DeepCopyInto(out *SyndesisList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Syndesis, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SyndesisList.
func (in *SyndesisList) DeepCopy() *SyndesisList {
	if in == nil {
		return nil
	}
	out := new(SyndesisList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SyndesisList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SyndesisSpec) DeepCopyInto(out *SyndesisSpec) {
	*out = *in
	if in.DemoData != nil {
		in, out := &in.DemoData, &out.DemoData
		*out = new(bool)
		**out = **in
	}
	if in.DeployIntegrations != nil {
		in, out := &in.DeployIntegrations, &out.DeployIntegrations
		*out = new(bool)
		**out = **in
	}
	if in.TestSupport != nil {
		in, out := &in.TestSupport, &out.TestSupport
		*out = new(bool)
		**out = **in
	}
	in.Integration.DeepCopyInto(&out.Integration)
	in.Components.DeepCopyInto(&out.Components)
	if in.Addons != nil {
		in, out := &in.Addons, &out.Addons
		*out = make(AddonsSpec, len(*in))
		for key, val := range *in {
			var outVal map[string]string
			if val == nil {
				(*out)[key] = nil
			} else {
				in, out := &val, &outVal
				*out = make(Parameters, len(*in))
				for key, val := range *in {
					(*out)[key] = val
				}
			}
			(*out)[key] = outVal
		}
	}
	if in.MavenRepositories != nil {
		in, out := &in.MavenRepositories, &out.MavenRepositories
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SyndesisSpec.
func (in *SyndesisSpec) DeepCopy() *SyndesisSpec {
	if in == nil {
		return nil
	}
	out := new(SyndesisSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SyndesisStatus) DeepCopyInto(out *SyndesisStatus) {
	*out = *in
	if in.LastUpgradeFailure != nil {
		in, out := &in.LastUpgradeFailure, &out.LastUpgradeFailure
		*out = (*in).DeepCopy()
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SyndesisStatus.
func (in *SyndesisStatus) DeepCopy() *SyndesisStatus {
	if in == nil {
		return nil
	}
	out := new(SyndesisStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *UIConfiguration) DeepCopyInto(out *UIConfiguration) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new UIConfiguration.
func (in *UIConfiguration) DeepCopy() *UIConfiguration {
	if in == nil {
		return nil
	}
	out := new(UIConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *UpgradeConfiguration) DeepCopyInto(out *UpgradeConfiguration) {
	*out = *in
	out.Resources = in.Resources
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new UpgradeConfiguration.
func (in *UpgradeConfiguration) DeepCopy() *UpgradeConfiguration {
	if in == nil {
		return nil
	}
	out := new(UpgradeConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VolumeOnlyResources) DeepCopyInto(out *VolumeOnlyResources) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VolumeOnlyResources.
func (in *VolumeOnlyResources) DeepCopy() *VolumeOnlyResources {
	if in == nil {
		return nil
	}
	out := new(VolumeOnlyResources)
	in.DeepCopyInto(out)
	return out
}
