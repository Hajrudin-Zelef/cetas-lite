---
id: collect-261001-cisco/cisco/ucmf-groups-entpublic-documents-enterprise-country-webasset-hw-330680-pdf-3dbaaa58-26
title: "A line starting with the # sign is comments."
domain: cisco
role: reference
task: reference
actors: ["Huawei"]
dates: ["2013-04-15"]
keywords: ["copyright", "parameters"]
source: docs/RAG/collect-261001-cisco/ucmf-groups-entpublic-documents-enterprise-country-webasset-hw-330680-pdf-3dbaaa58.md
source_anchor: ""
source_lines: [3527, 3740]
sha256: 90f9e34e51bca235bfb1628ea7955376fce4889a64c7d547030e53b247b84302
---

# A line starting with the # sign is comments.

Field Length Description
RSVP Length 16 bits Indicates the total length of the RSVP message, in bytes.
Objects Variable Indicates the object of the RSVP message. Each RSVP
message contains kinds of objects. The carried objects vary
with types of messages.
Length 16 bits Indicates the total length of the object, in bytes. Its value must
be a multiple of 4, and at least 4.
Class_Numb
er
8 bits Identifies an object class. Each object class has a name, such
as SESSION, SENDER_TEMPLATE, and TIME_VALUE.
C-Type 8 bits Indicates the object type, unique within the Class_Number.
The Class-Number and C-Type is used together to define a
unique type for each object.
Object
Content
Variable Indicates contents of objects. The length of this field is
changeable.
 
NOTE
For details of each type of RSVP messages, refer to RFC 3209 and RFC 2205.
Path Message
In RSVP-TE, a Path message is used to create an RSVP session and maintain a path state. The
Path message is sent from the ingress node to the egress node in the direction of data flows. On
each node, the path state block (PSB) is created.
NOTE
The source IP address of a Path message is the LSR ID of the ingress node and the destination IP address
is the LSR ID of the egress node.
Table 3-12 lists some objects carried in the Path message.
Table 3-12 Path message objects
Message
Object
Class_Num
ber
C-Type Object Content
SESSION 1 1 Carries RSVP session information,
including the destination address, tunnel
ID, and extend tunnel ID.
RSVP_HOP 3 1 Identifies the IP address and the handle of
the outgoing interface of the previous hop
that sends the Path message.
TIME_VALU
E
5 1 Carries the refreshing interval.
Enterprise Data Communication Products
Feature Description - MPLS 3 MPLS TE
Issue 04 (2013-04-15) Huawei Proprietary and Confidential
Copyright © Huawei Technologies Co., Ltd.
73

Message
Object
Class_Num
ber
C-Type Object Content
SENDER_TE
MPLATE
11 1 Specifies the sender IP address and LSP
ID.
SENDER_TS
PEC
12 2 Defines traffic characteristics of the data
flow.
LABEL_REQ
UEST
19 1 Indicates LABEL_REQUEST object,
which is carried only in Path messages.
ADSPEC 13 2 Collects actual QoS parameters about the
path, such as estimation of bandwidth of
the path, minimal path delay, and path
MTU.
EXPLICIT_R
OUTE
20 1 ERO, describes information about the
path through which the LSP passes. The
explicit paths can be strict or loose. Path
messages are then forwarded along the
specified ERO, without being restricted
by IGP shortest path.
RECORD_R
OUTE
21 1 RRO, lists the LSRs that the Path message
passes when being transmitted. RRO can
be used to collect path information and
discover route loops. It can also be copied
to the next Path message for implementing
Route Pinning.
SESSION_A
TTRIBUTE
207 l 1:
LSP_TUN
NEL_RA
l 7: LSP
Tunnel
Specifies the setup priority, hold priority,
reservation style, affinity, and other
information.
 
Resv message
After receiving a Path message, the egress node reply with Resv messages. The Resv message,
carrying resource reservation information, is sent to the previous node hop-by-hop. Each passing
node creates and maintains a reserved state block (RSB) and allocates a label. When the Resv
message reaches the ingress node, an LSP is set up successfully.
Table 3-13 describes objects carried in the Resv message.
Enterprise Data Communication Products
Feature Description - MPLS 3 MPLS TE
Issue 04 (2013-04-15) Huawei Proprietary and Confidential
Copyright © Huawei Technologies Co., Ltd.
74

Table 3-13 Resv message object
Message
Object
Class_Num
ber
C-Type Object Content
INTEGRITY 4 1 Carries cryptographic data to authenticate
the originating node and to verify the
contents of this RSVP message.
SESSION 1 1 Carries RSVP session information,
including the destination address, tunnel
ID, and extend tunnel ID.
RSVP_HOP 3 1 Identifies the IP address and the index of
the outgoing interface that sends the Resv
message.
TIME_VALU
E
5 1 Carries the refreshing interval. By default,
the value is 30 seconds.
STYLE 8 1 Indicates the resource reservation style. It
is specified on the ingress node.
FLOW_SPEC 9 l 1:
Reserved
(obsolete)
flowspec
object
l 2: Inv-serv
flowspec
object
Specifies the QoS characteristics of the
data flow.
FILTER_SPE
C
10 1 Specifies the sender IP address and LSP
ID of the node that sends the message.
RECORD_R
OUTE
21 1 RRO, collects the IP address of the
incoming interface, LSR-ID, and the IP
address of the outgoing interface of the
node along the path.
LABEL 16 1 Indicates the assigned label.
RESV_CONF
IRM
15 1 Indicates a confirmation of the resource
reservation is requested when this object
is received. This object carries the IP
address of the node that requests a
confirmation of the resource reservation.
 
Reservation Styles
The treatment style of reserving resources for different senders within the same session is called
a reservation style. The following reservation styles are supported:
Enterprise Data Communication Products
Feature Description - MPLS 3 MPLS TE
Issue 04 (2013-04-15) Huawei Proprietary and Confidential
Copyright © Huawei Technologies Co., Ltd.
75

l Fixed Filter (FF) style: creates a separate reservation for a tunnel from a particular sender.
This sender does not share its resource reservation with other senders. A resource
reservation on the same link is used by a specific CR-LSP.
l Shared Explicit (SE) style: creates a single reservation shared by a set of selected upstream
senders. The same resource reservation on the same link is shared by different CR-LSPs.
3.2.6 Traffic Forwarding
Importing Traffic to an MPLS TE Tunnel
The traffic forwarding function imports traffic to a tunnel and forwards traffic over the tunnel.
Although the information advertisement, path calculation, and path establishment are used to
establish a CR-LSP in an MPLS TE tunnel, a CR-LSP (unlike an LDP LSP) cannot automatically
import traffic. The traffic forwarding component must be used to import traffic to the CR-LSP
before it forwards traffic.
l Static Routes: applies to scenarios with simple network topology or stable network
environment.
l Policy-based Routing: applies to scenarios that require load balancing and security
monitoring.
l Tunnel Policies: applies to scenarios in which TE tunnels need to be established to transmit
VPN services.
l Auto Routes: applies to scenarios with complex network topology or unstable network
environment.
Static Routes
Using static routes is the simplest method to import traffic to an MPLS TE tunnel. A TE static
route works in the same way as a common static route and has a TE tunnel interface as an
outbound interface.
Policy-based Routing
The policy-based routing (PBR) allows the system to select routes based on user-defined
policies, improving security and load balancing traffic. If PBR is enabled on an MPLS network,
IP packets are forwarded over specific CR-LSPs based on PBR rules.
MPLS TE PBR is implemented based on a set of matching rules and behaviors. The rules and
behaviors are defined using an apply clause, in which the outbound interface is a specific tunnel
interface. If packets do not match PBR rules, they are properly forwarded using IP; if they match
PBR rules, they are forwarded over specific CR-LSPs.
Tunnel Policies
Generally, VPN traffic is forwarded through an LSP but not an MPLS TE tunnel. To import
VPN traffic to the MPLS TE tunnel, you need to configure a tunnel policy. Two tunnel policies
are available.
l Tunnel type prioritizing policy: Such a policy specifies the sequence in which different
types of tunnels are selected by the VPN. You can specify the VPN to select the TE tunnel
first.
Enterprise Data Communication Products
Feature Description - MPLS 3 MPLS TE
Issue 04 (2013-04-15) Huawei Proprietary and Confidential
Copyright © Huawei Technologies Co., Ltd.
76

