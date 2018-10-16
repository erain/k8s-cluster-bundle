// Code generated by protoc-gen-go. DO NOT EDIT.
// source: pkg/apis/bundle/v1alpha1/cluster_bundle.proto

package v1alpha1 // import "github.com/GoogleCloudPlatform/k8s-cluster-bundle/pkg/apis/bundle/v1alpha1"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// The ClusterBundle is a packaging format for Kubernetes Components.
type ClusterBundle struct {
	// Required. Kubernetes API Version for the Bundle.
	// Must be gke.io/k8s-cluster-bundle/v1alpha1.
	ApiVersion string `protobuf:"bytes,1,opt,name=api_version,json=apiVersion,proto3" json:"api_version,omitempty"`
	// Required. The Kubernetes `kind` for this object. Must be 'ClusterBundle'.
	Kind string `protobuf:"bytes,2,opt,name=kind,proto3" json:"kind,omitempty"`
	// Required. Kubernetes ObjectMeta proto. The Metadata.name field must be
	// filled out for each Bundle.
	Metadata *ObjectMeta `protobuf:"bytes,3,opt,name=metadata,proto3" json:"metadata,omitempty"`
	// Spec for the ClusterBundle, which specifies the intended Kubernetes cluster
	// configuration.
	Spec                 *ClusterBundleSpec `protobuf:"bytes,4,opt,name=spec,proto3" json:"spec,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *ClusterBundle) Reset()         { *m = ClusterBundle{} }
func (m *ClusterBundle) String() string { return proto.CompactTextString(m) }
func (*ClusterBundle) ProtoMessage()    {}
func (*ClusterBundle) Descriptor() ([]byte, []int) {
	return fileDescriptor_6b30f666a9d42204, []int{0}
}
func (m *ClusterBundle) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ClusterBundle.Unmarshal(m, b)
}
func (m *ClusterBundle) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ClusterBundle.Marshal(b, m, deterministic)
}
func (m *ClusterBundle) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ClusterBundle.Merge(m, src)
}
func (m *ClusterBundle) XXX_Size() int {
	return xxx_messageInfo_ClusterBundle.Size(m)
}
func (m *ClusterBundle) XXX_DiscardUnknown() {
	xxx_messageInfo_ClusterBundle.DiscardUnknown(m)
}

var xxx_messageInfo_ClusterBundle proto.InternalMessageInfo

func (m *ClusterBundle) GetApiVersion() string {
	if m != nil {
		return m.ApiVersion
	}
	return ""
}

func (m *ClusterBundle) GetKind() string {
	if m != nil {
		return m.Kind
	}
	return ""
}

func (m *ClusterBundle) GetMetadata() *ObjectMeta {
	if m != nil {
		return m.Metadata
	}
	return nil
}

func (m *ClusterBundle) GetSpec() *ClusterBundleSpec {
	if m != nil {
		return m.Spec
	}
	return nil
}

// A spec object that wraps the Cluster Bundle.
type ClusterBundleSpec struct {
	// Version-string for this bundle. The version should be a SemVer string (see
	// https://semver.org/) of the form X.Y.Z (Major.Minor.Patch).  Generally
	// speaking, a major-version (changes should indicate breaking changes,
	// minor-versions should indicate backwards compatible features, and patch
	// changes should indicate backwords compatible. If there are any changes to
	// the bundle, then the version string must be incremented.
	//
	// If a bundle is versioned, then all its components must be versioned.
	Version string `protobuf:"bytes,1,opt,name=version,proto3" json:"version,omitempty"`
	// Kubernetes objects grouped into component packages and versioned together.
	// These could be applications or they could be some sort of supporting
	// object collection.
	Components []*ComponentPackage `protobuf:"bytes,2,rep,name=components,proto3" json:"components,omitempty"`
	// Cluster components that are specified externally as Files. The process of inlining
	// for a bundle reads component files into components, and so after
	// inlining, this list will be empty.
	ComponentFiles       []*File  `protobuf:"bytes,3,rep,name=component_files,json=componentFiles,proto3" json:"component_files,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ClusterBundleSpec) Reset()         { *m = ClusterBundleSpec{} }
func (m *ClusterBundleSpec) String() string { return proto.CompactTextString(m) }
func (*ClusterBundleSpec) ProtoMessage()    {}
func (*ClusterBundleSpec) Descriptor() ([]byte, []int) {
	return fileDescriptor_6b30f666a9d42204, []int{1}
}
func (m *ClusterBundleSpec) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ClusterBundleSpec.Unmarshal(m, b)
}
func (m *ClusterBundleSpec) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ClusterBundleSpec.Marshal(b, m, deterministic)
}
func (m *ClusterBundleSpec) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ClusterBundleSpec.Merge(m, src)
}
func (m *ClusterBundleSpec) XXX_Size() int {
	return xxx_messageInfo_ClusterBundleSpec.Size(m)
}
func (m *ClusterBundleSpec) XXX_DiscardUnknown() {
	xxx_messageInfo_ClusterBundleSpec.DiscardUnknown(m)
}

var xxx_messageInfo_ClusterBundleSpec proto.InternalMessageInfo

func (m *ClusterBundleSpec) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

func (m *ClusterBundleSpec) GetComponents() []*ComponentPackage {
	if m != nil {
		return m.Components
	}
	return nil
}

func (m *ClusterBundleSpec) GetComponentFiles() []*File {
	if m != nil {
		return m.ComponentFiles
	}
	return nil
}

func init() {
	proto.RegisterType((*ClusterBundle)(nil), "pkg.apis.bundle.v1alpha1.ClusterBundle")
	proto.RegisterType((*ClusterBundleSpec)(nil), "pkg.apis.bundle.v1alpha1.ClusterBundleSpec")
}

func init() {
	proto.RegisterFile("pkg/apis/bundle/v1alpha1/cluster_bundle.proto", fileDescriptor_6b30f666a9d42204)
}

var fileDescriptor_6b30f666a9d42204 = []byte{
	// 351 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x52, 0x4f, 0x4f, 0xfa, 0x30,
	0x18, 0xce, 0x80, 0xfc, 0x7e, 0x5a, 0xa2, 0xc6, 0x9e, 0x1a, 0x0e, 0x4a, 0x88, 0x07, 0x82, 0xb2,
	0x09, 0x5e, 0xbc, 0x69, 0x20, 0x91, 0x84, 0x68, 0x24, 0x33, 0xf1, 0xe0, 0x65, 0xe9, 0xba, 0x32,
	0xea, 0xba, 0xb5, 0x59, 0x3b, 0x3e, 0xa0, 0x17, 0xbf, 0x96, 0xa1, 0x1d, 0x28, 0x31, 0xf5, 0xd6,
	0xf7, 0x7d, 0x9f, 0xe7, 0x79, 0xff, 0x3c, 0x05, 0x43, 0x99, 0xa5, 0x01, 0x96, 0x4c, 0x05, 0x71,
	0x55, 0x24, 0x9c, 0x06, 0xeb, 0x11, 0xe6, 0x72, 0x85, 0x47, 0x01, 0xe1, 0x95, 0xd2, 0xb4, 0x8c,
	0x6c, 0xde, 0x97, 0xa5, 0xd0, 0x02, 0x22, 0x99, 0xa5, 0xfe, 0x06, 0xee, 0xd7, 0xe9, 0x2d, 0xbc,
	0x73, 0xe5, 0x14, 0xb2, 0x71, 0x44, 0x44, 0x9e, 0x8b, 0xc2, 0xea, 0x74, 0xae, 0xdd, 0x6d, 0x45,
	0x2e, 0x45, 0x41, 0x0b, 0x1d, 0x49, 0x4c, 0x32, 0x9c, 0xd6, 0x9d, 0x3b, 0x03, 0x27, 0x43, 0xc4,
	0xef, 0x94, 0xe8, 0x28, 0xa7, 0x1a, 0x5b, 0x6c, 0xef, 0xd3, 0x03, 0x47, 0x53, 0x3b, 0xfe, 0xc4,
	0x80, 0xe1, 0x39, 0x68, 0x63, 0xc9, 0xa2, 0x35, 0x2d, 0x15, 0x13, 0x05, 0xf2, 0xba, 0x5e, 0xff,
	0x30, 0x04, 0x58, 0xb2, 0x57, 0x9b, 0x81, 0x10, 0xb4, 0x32, 0x56, 0x24, 0xa8, 0x61, 0x2a, 0xe6,
	0x0d, 0xef, 0xc1, 0xc1, 0x46, 0x34, 0xc1, 0x1a, 0xa3, 0x66, 0xd7, 0xeb, 0xb7, 0xc7, 0x17, 0xbe,
	0x6b, 0x7f, 0xff, 0xd9, 0x4c, 0xf1, 0x44, 0x35, 0x0e, 0x77, 0x2c, 0x78, 0x07, 0x5a, 0x4a, 0x52,
	0x82, 0x5a, 0x86, 0x7d, 0xe9, 0x66, 0xef, 0x4d, 0xfb, 0x22, 0x29, 0x09, 0x0d, 0xb1, 0xf7, 0xe1,
	0x81, 0xd3, 0x5f, 0x35, 0x88, 0xc0, 0xff, 0xfd, 0x4d, 0xb6, 0x21, 0x9c, 0x03, 0xb0, 0x3b, 0xa0,
	0x42, 0x8d, 0x6e, 0xb3, 0xdf, 0x1e, 0x0f, 0xfe, 0x68, 0xbb, 0xc5, 0x2e, 0xec, 0xad, 0xc3, 0x1f,
	0x6c, 0x38, 0x03, 0x27, 0xdf, 0x66, 0x2c, 0x19, 0xa7, 0x0a, 0x35, 0x8d, 0xe0, 0x99, 0x5b, 0xf0,
	0x81, 0x71, 0x1a, 0x1e, 0xef, 0x68, 0x9b, 0x50, 0x4d, 0x1e, 0xdf, 0xe6, 0x29, 0xd3, 0xab, 0x2a,
	0xf6, 0x89, 0xc8, 0x83, 0x99, 0x10, 0x29, 0xa7, 0x53, 0x2e, 0xaa, 0x64, 0xc1, 0xb1, 0x5e, 0x8a,
	0x32, 0x0f, 0xb2, 0x5b, 0x35, 0xac, 0xff, 0xdb, 0xb0, 0xb6, 0xd7, 0x65, 0x77, 0xfc, 0xcf, 0x78,
	0x7c, 0xf3, 0x15, 0x00, 0x00, 0xff, 0xff, 0xdd, 0xd7, 0x12, 0x8e, 0xba, 0x02, 0x00, 0x00,
}
