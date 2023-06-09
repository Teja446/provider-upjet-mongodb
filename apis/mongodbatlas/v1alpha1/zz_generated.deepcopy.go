//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright 2022 Upbound Inc.
*/

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *APIKeysObservation) DeepCopyInto(out *APIKeysObservation) {
	*out = *in
	if in.APIKeyID != nil {
		in, out := &in.APIKeyID, &out.APIKeyID
		*out = new(string)
		**out = **in
	}
	if in.RoleNames != nil {
		in, out := &in.RoleNames, &out.RoleNames
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new APIKeysObservation.
func (in *APIKeysObservation) DeepCopy() *APIKeysObservation {
	if in == nil {
		return nil
	}
	out := new(APIKeysObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *APIKeysParameters) DeepCopyInto(out *APIKeysParameters) {
	*out = *in
	if in.APIKeyID != nil {
		in, out := &in.APIKeyID, &out.APIKeyID
		*out = new(string)
		**out = **in
	}
	if in.RoleNames != nil {
		in, out := &in.RoleNames, &out.RoleNames
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new APIKeysParameters.
func (in *APIKeysParameters) DeepCopy() *APIKeysParameters {
	if in == nil {
		return nil
	}
	out := new(APIKeysParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Project) DeepCopyInto(out *Project) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Project.
func (in *Project) DeepCopy() *Project {
	if in == nil {
		return nil
	}
	out := new(Project)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Project) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ProjectList) DeepCopyInto(out *ProjectList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Project, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ProjectList.
func (in *ProjectList) DeepCopy() *ProjectList {
	if in == nil {
		return nil
	}
	out := new(ProjectList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ProjectList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ProjectObservation) DeepCopyInto(out *ProjectObservation) {
	*out = *in
	if in.APIKeys != nil {
		in, out := &in.APIKeys, &out.APIKeys
		*out = make([]APIKeysObservation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.ClusterCount != nil {
		in, out := &in.ClusterCount, &out.ClusterCount
		*out = new(float64)
		**out = **in
	}
	if in.Created != nil {
		in, out := &in.Created, &out.Created
		*out = new(string)
		**out = **in
	}
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.IsCollectDatabaseSpecificsStatisticsEnabled != nil {
		in, out := &in.IsCollectDatabaseSpecificsStatisticsEnabled, &out.IsCollectDatabaseSpecificsStatisticsEnabled
		*out = new(bool)
		**out = **in
	}
	if in.IsDataExplorerEnabled != nil {
		in, out := &in.IsDataExplorerEnabled, &out.IsDataExplorerEnabled
		*out = new(bool)
		**out = **in
	}
	if in.IsPerformanceAdvisorEnabled != nil {
		in, out := &in.IsPerformanceAdvisorEnabled, &out.IsPerformanceAdvisorEnabled
		*out = new(bool)
		**out = **in
	}
	if in.IsRealtimePerformancePanelEnabled != nil {
		in, out := &in.IsRealtimePerformancePanelEnabled, &out.IsRealtimePerformancePanelEnabled
		*out = new(bool)
		**out = **in
	}
	if in.IsSchemaAdvisorEnabled != nil {
		in, out := &in.IsSchemaAdvisorEnabled, &out.IsSchemaAdvisorEnabled
		*out = new(bool)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.OrgID != nil {
		in, out := &in.OrgID, &out.OrgID
		*out = new(string)
		**out = **in
	}
	if in.ProjectOwnerID != nil {
		in, out := &in.ProjectOwnerID, &out.ProjectOwnerID
		*out = new(string)
		**out = **in
	}
	if in.RegionUsageRestrictions != nil {
		in, out := &in.RegionUsageRestrictions, &out.RegionUsageRestrictions
		*out = new(string)
		**out = **in
	}
	if in.Teams != nil {
		in, out := &in.Teams, &out.Teams
		*out = make([]TeamsObservation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.WithDefaultAlertsSettings != nil {
		in, out := &in.WithDefaultAlertsSettings, &out.WithDefaultAlertsSettings
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ProjectObservation.
func (in *ProjectObservation) DeepCopy() *ProjectObservation {
	if in == nil {
		return nil
	}
	out := new(ProjectObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ProjectParameters) DeepCopyInto(out *ProjectParameters) {
	*out = *in
	if in.APIKeys != nil {
		in, out := &in.APIKeys, &out.APIKeys
		*out = make([]APIKeysParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.IsCollectDatabaseSpecificsStatisticsEnabled != nil {
		in, out := &in.IsCollectDatabaseSpecificsStatisticsEnabled, &out.IsCollectDatabaseSpecificsStatisticsEnabled
		*out = new(bool)
		**out = **in
	}
	if in.IsDataExplorerEnabled != nil {
		in, out := &in.IsDataExplorerEnabled, &out.IsDataExplorerEnabled
		*out = new(bool)
		**out = **in
	}
	if in.IsPerformanceAdvisorEnabled != nil {
		in, out := &in.IsPerformanceAdvisorEnabled, &out.IsPerformanceAdvisorEnabled
		*out = new(bool)
		**out = **in
	}
	if in.IsRealtimePerformancePanelEnabled != nil {
		in, out := &in.IsRealtimePerformancePanelEnabled, &out.IsRealtimePerformancePanelEnabled
		*out = new(bool)
		**out = **in
	}
	if in.IsSchemaAdvisorEnabled != nil {
		in, out := &in.IsSchemaAdvisorEnabled, &out.IsSchemaAdvisorEnabled
		*out = new(bool)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.OrgID != nil {
		in, out := &in.OrgID, &out.OrgID
		*out = new(string)
		**out = **in
	}
	if in.ProjectOwnerID != nil {
		in, out := &in.ProjectOwnerID, &out.ProjectOwnerID
		*out = new(string)
		**out = **in
	}
	if in.RegionUsageRestrictions != nil {
		in, out := &in.RegionUsageRestrictions, &out.RegionUsageRestrictions
		*out = new(string)
		**out = **in
	}
	if in.Teams != nil {
		in, out := &in.Teams, &out.Teams
		*out = make([]TeamsParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.WithDefaultAlertsSettings != nil {
		in, out := &in.WithDefaultAlertsSettings, &out.WithDefaultAlertsSettings
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ProjectParameters.
func (in *ProjectParameters) DeepCopy() *ProjectParameters {
	if in == nil {
		return nil
	}
	out := new(ProjectParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ProjectSpec) DeepCopyInto(out *ProjectSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ProjectSpec.
func (in *ProjectSpec) DeepCopy() *ProjectSpec {
	if in == nil {
		return nil
	}
	out := new(ProjectSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ProjectStatus) DeepCopyInto(out *ProjectStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ProjectStatus.
func (in *ProjectStatus) DeepCopy() *ProjectStatus {
	if in == nil {
		return nil
	}
	out := new(ProjectStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TeamsObservation) DeepCopyInto(out *TeamsObservation) {
	*out = *in
	if in.RoleNames != nil {
		in, out := &in.RoleNames, &out.RoleNames
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.TeamID != nil {
		in, out := &in.TeamID, &out.TeamID
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TeamsObservation.
func (in *TeamsObservation) DeepCopy() *TeamsObservation {
	if in == nil {
		return nil
	}
	out := new(TeamsObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TeamsParameters) DeepCopyInto(out *TeamsParameters) {
	*out = *in
	if in.RoleNames != nil {
		in, out := &in.RoleNames, &out.RoleNames
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.TeamID != nil {
		in, out := &in.TeamID, &out.TeamID
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TeamsParameters.
func (in *TeamsParameters) DeepCopy() *TeamsParameters {
	if in == nil {
		return nil
	}
	out := new(TeamsParameters)
	in.DeepCopyInto(out)
	return out
}
