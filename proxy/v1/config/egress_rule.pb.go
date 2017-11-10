// Code generated by protoc-gen-go.
// source: proxy/v1/config/egress_rule.proto
// DO NOT EDIT!

package istio_proxy_v1_config

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// Egress rules describe the properties of a service outside Istio. When transparent proxying
// is used, egress rules signify a white listed set of domains that microserves in the mesh
// are allowed to access. A subset of routing rules and all destination policies can be applied
// on the service targeted by an egress rule. The destination of an egress rule is allowed to
// contain wildcards (e.g., *.foo.com). Currently, only HTTP-based services can be expressed
// through the egress rule. If TLS origination from the sidecar is desired, the protocol
// associated with the service port must be marked as HTTPS, and the service is expected to
// be accessed over HTTP (e.g., http://gmail.com:443). The sidecar will automatically upgrade
// the connection to TLS when initiating a connection with the external service.
//
// For example, the following egress rule describes the set of services hosted under the *.foo.com domain
//
//     kind: EgressRule
//     metadata:
//       name: foo-egress-rule
//     spec:
//       destination:
//         service: *.foo.com
//       ports:
//         - port: 80
//           protocol: http
//         - port: 443
//           protocol: https
//
type EgressRule struct {
	// REQUIRED: Hostname or a wildcard domain name associated with the external service.
	// ONLY the "service" field of destination will be taken into consideration. Name,
	// namespace, domain and labels are ignored. Routing rules and destination policies that
	// refer to these external services must have identical specification for the destination
	// as the corresponding egress rule. Wildcard domain specifications must conform to format
	// allowed by Envoy's Virtual Host specification, such as “*.foo.com” or “*-bar.foo.com”.
	// The character '*' in a domain specification indicates a non-empty string. Hence, a wildcard
	// domain of form “*-bar.foo.com” will match “baz-bar.foo.com” but not “-bar.foo.com”.
	Destination *IstioService `protobuf:"bytes,1,opt,name=destination" json:"destination,omitempty"`
	// REQUIRED: list of ports on which the external service is available.
	Ports []*EgressRule_Port `protobuf:"bytes,2,rep,name=ports" json:"ports,omitempty"`
	// Forward all the external traffic through a dedicated egress proxy. It is used in some scenarios
	// where there is a requirement that all the external traffic goes through special dedicated nodes/pods.
	// These dedicated egress nodes could then be more closely monitored for security vulnerabilities.
	//
	// The default is false, i.e. the sidecar forwards external traffic directly to the external service.
	UseEgressProxy bool `protobuf:"varint,3,opt,name=use_egress_proxy,json=useEgressProxy" json:"use_egress_proxy,omitempty"`
}

func (m *EgressRule) Reset()                    { *m = EgressRule{} }
func (m *EgressRule) String() string            { return proto.CompactTextString(m) }
func (*EgressRule) ProtoMessage()               {}
func (*EgressRule) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{0} }

func (m *EgressRule) GetDestination() *IstioService {
	if m != nil {
		return m.Destination
	}
	return nil
}

func (m *EgressRule) GetPorts() []*EgressRule_Port {
	if m != nil {
		return m.Ports
	}
	return nil
}

func (m *EgressRule) GetUseEgressProxy() bool {
	if m != nil {
		return m.UseEgressProxy
	}
	return false
}

// Port describes the properties of a specific TCP port of an external service.
type EgressRule_Port struct {
	// A valid non-negative integer port number.
	Port int32 `protobuf:"varint,1,opt,name=port" json:"port,omitempty"`
	// The protocol to communicate with the external services.
	// MUST BE one of HTTP|HTTPS|GRPC|HTTP2.
	Protocol string `protobuf:"bytes,2,opt,name=protocol" json:"protocol,omitempty"`
}

func (m *EgressRule_Port) Reset()                    { *m = EgressRule_Port{} }
func (m *EgressRule_Port) String() string            { return proto.CompactTextString(m) }
func (*EgressRule_Port) ProtoMessage()               {}
func (*EgressRule_Port) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{0, 0} }

func (m *EgressRule_Port) GetPort() int32 {
	if m != nil {
		return m.Port
	}
	return 0
}

func (m *EgressRule_Port) GetProtocol() string {
	if m != nil {
		return m.Protocol
	}
	return ""
}

func init() {
	proto.RegisterType((*EgressRule)(nil), "istio.proxy.v1.config.EgressRule")
	proto.RegisterType((*EgressRule_Port)(nil), "istio.proxy.v1.config.EgressRule.Port")
}

func init() { proto.RegisterFile("proxy/v1/config/egress_rule.proto", fileDescriptor1) }

var fileDescriptor1 = []byte{
	// 232 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0xcf, 0xcd, 0x4a, 0x03, 0x31,
	0x10, 0x07, 0x70, 0xd2, 0x0f, 0xa9, 0xb3, 0x20, 0x12, 0x10, 0xc2, 0x9e, 0xa2, 0x82, 0xe4, 0x94,
	0xa5, 0x15, 0x3c, 0x79, 0xed, 0xc1, 0x5b, 0x89, 0x0f, 0x50, 0x74, 0x1d, 0x4b, 0x60, 0xd9, 0x59,
	0x26, 0xc9, 0xa2, 0xcf, 0xed, 0x0b, 0x48, 0xb2, 0xa2, 0x45, 0xf4, 0x36, 0x19, 0xe6, 0xf7, 0xcf,
	0x0c, 0x5c, 0x0e, 0x4c, 0x6f, 0xef, 0xcd, 0xb8, 0x6e, 0x5a, 0xea, 0x5f, 0xfd, 0xa1, 0xc1, 0x03,
	0x63, 0x08, 0x7b, 0x4e, 0x1d, 0xda, 0x81, 0x29, 0x92, 0xbc, 0xf0, 0x21, 0x7a, 0xb2, 0x65, 0xd0,
	0x8e, 0x6b, 0x3b, 0x0d, 0xd6, 0xfa, 0xb7, 0x64, 0x4a, 0x11, 0x8f, 0xe0, 0xd5, 0x87, 0x00, 0xd8,
	0x96, 0x38, 0x97, 0x3a, 0x94, 0x5b, 0xa8, 0x5e, 0x30, 0x44, 0xdf, 0x3f, 0x45, 0x4f, 0xbd, 0x12,
	0x5a, 0x98, 0x6a, 0x73, 0x6d, 0xff, 0x4c, 0xb7, 0x0f, 0xb9, 0xfb, 0x88, 0x3c, 0xfa, 0x16, 0xdd,
	0xb1, 0x93, 0xf7, 0xb0, 0x1c, 0x88, 0x63, 0x50, 0x33, 0x3d, 0x37, 0xd5, 0xe6, 0xe6, 0x9f, 0x80,
	0x9f, 0x8f, 0xed, 0x8e, 0x38, 0xba, 0x09, 0x49, 0x03, 0xe7, 0x29, 0xe0, 0xfe, 0xeb, 0xca, 0x82,
	0xd4, 0x5c, 0x0b, 0xb3, 0x72, 0x67, 0x29, 0xe0, 0x84, 0x76, 0xb9, 0x5b, 0xdf, 0xc1, 0x22, 0x43,
	0x29, 0x61, 0x91, 0x69, 0xd9, 0x77, 0xe9, 0x4a, 0x2d, 0x6b, 0x58, 0x95, 0x13, 0x5b, 0xea, 0xd4,
	0x4c, 0x0b, 0x73, 0xea, 0xbe, 0xdf, 0xcf, 0x27, 0xa5, 0xba, 0xfd, 0x0c, 0x00, 0x00, 0xff, 0xff,
	0xa2, 0xbc, 0x4d, 0x09, 0x5a, 0x01, 0x00, 0x00,
}