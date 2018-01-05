// Code generated by protoc-gen-go. DO NOT EDIT.
// source: routing/v1alpha2/egress_routes.proto

package istio_routing_v1alpha2

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// Egress routes describe the properties of a service outside Istio. When
// transparent proxying is used, egress routes describe a white-listed set
// of domains that microserves in the mesh are allowed to access. A subset
// of features from routing rules and all features of destination rules can
// be applied on the service described by an egress route.
//
// Services using HTTPS will be treated as opaque TCP services. If features
// provided by routing rules (e.g., URL rewrites) or HTTP-level metrics
// generated by Envoy are required, the service must be accessed over
// HTTP. The sidecar will perform TLS origination on behalf of the
// application.
//
// Note that services described by an egress route will be visible at all
// sidecars in the mesh.
//
// For example, the following EgressRoute describes the set of services
// hosted under the *.foo.com domain
//
//     apiVersion: config.istio.io/v1alpha2
//     kind: EgressRoute
//     metadata:
//       name: foo-egress-rule
//     spec:
//       servers:
//       - port:
//           number: 443
//           name: http
//         domains:
//         - *.foo.com
//         tls:
//           mode: ORIGINATE
//
// Route rules can be applied to services described by an egress route. The
// following sample route rule rewrites /foocatalog to /barcatalog before
// forwarding the call to the intended destination.
//
//     apiVersion: config.istio.io/v1alpha2
//     kind: RouteRule
//     metadata:
//       name: foo-rule
//     spec:
//       hosts:
//       - *.foo.com
//       http:
//       - match:
//         - uri:
//             prefix: /foocatalog
//         rewrite:
//           uri: /barcatalog
//
type EgressRoute struct {
	// REQUIRED: A list of server specifications.
	Servers []*Server `protobuf:"bytes,1,rep,name=servers" json:"servers,omitempty"`
}

func (m *EgressRoute) Reset()                    { *m = EgressRoute{} }
func (m *EgressRoute) String() string            { return proto.CompactTextString(m) }
func (*EgressRoute) ProtoMessage()               {}
func (*EgressRoute) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{0} }

func (m *EgressRoute) GetServers() []*Server {
	if m != nil {
		return m.Servers
	}
	return nil
}

func init() {
	proto.RegisterType((*EgressRoute)(nil), "istio.routing.v1alpha2.EgressRoute")
}

func init() { proto.RegisterFile("routing/v1alpha2/egress_routes.proto", fileDescriptor1) }

var fileDescriptor1 = []byte{
	// 131 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0x29, 0xca, 0x2f, 0x2d,
	0xc9, 0xcc, 0x4b, 0xd7, 0x2f, 0x33, 0x4c, 0xcc, 0x29, 0xc8, 0x48, 0x34, 0xd2, 0x4f, 0x4d, 0x2f,
	0x4a, 0x2d, 0x2e, 0x8e, 0x07, 0x89, 0xa7, 0x16, 0xeb, 0x15, 0x14, 0xe5, 0x97, 0xe4, 0x0b, 0x89,
	0x65, 0x16, 0x97, 0x64, 0xe6, 0xeb, 0x41, 0xd5, 0xea, 0xc1, 0xd4, 0x4a, 0xc9, 0x61, 0xe8, 0x4e,
	0x4f, 0x2c, 0x49, 0x2d, 0x4f, 0xac, 0x84, 0xe8, 0x53, 0x72, 0xe7, 0xe2, 0x76, 0x05, 0x1b, 0x17,
	0x04, 0x32, 0x4d, 0xc8, 0x82, 0x8b, 0xbd, 0x38, 0xb5, 0xa8, 0x2c, 0xb5, 0xa8, 0x58, 0x82, 0x51,
	0x81, 0x59, 0x83, 0xdb, 0x48, 0x4e, 0x0f, 0xbb, 0xc1, 0x7a, 0xc1, 0x60, 0x65, 0x41, 0x30, 0xe5,
	0x49, 0x6c, 0x60, 0xf3, 0x8c, 0x01, 0x01, 0x00, 0x00, 0xff, 0xff, 0x15, 0x3b, 0xc0, 0xad, 0xaf,
	0x00, 0x00, 0x00,
}
