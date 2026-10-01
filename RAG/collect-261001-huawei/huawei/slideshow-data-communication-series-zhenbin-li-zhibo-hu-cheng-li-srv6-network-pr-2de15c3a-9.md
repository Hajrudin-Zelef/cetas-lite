---
id: collect-261001-huawei/huawei/slideshow-data-communication-series-zhenbin-li-zhibo-hu-cheng-li-srv6-network-pr-2de15c3a-9
title: "slideshow-data-communication-series-zhenbin-li-zhibo-hu-cheng-li-srv6-network-pr-2de15c3a"
domain: huawei
role: reference
task: reference
actors: []
dates: ["2008-03-14"]
keywords: ["cost", "ethernet", "memory"]
source: docs/RAG/collect-261001-huawei/slideshow-data-communication-series-zhenbin-li-zhibo-hu-cheng-li-srv6-network-pr-2de15c3a.md
source_anchor: ""
source_lines: [1194, 1326]
sha256: b252746985be664ffdc254c5a27f2fc7454dbb87550474b3f0f5c58b9fa8927f
---

# slideshow-data-communication-series-zhenbin-li-zhibo-hu-cheng-li-srv6-network-pr-2de15c3a

• Te data and control planes of a device are tightly coupled and bun-
dled for sale. As such, these two planes rely on each other as far as evo-
lution is concerned, and the control plane on a device provided by one
vendor cannot control the data plane on a device provided by another.
60.
SRv6 Background ◾13
Te traditional network devices in the All IP 1.0 era are similar to the
IBM mainframes of the 1960s. To put it more exactly, the network
device hardware, Operating System (OS), and network applications are
tightly coupled and rely on each other. Because of this, if we aim to
conduct innovation or evolution on one part, we must also upgrade the
other parts correspondingly, and this architecture impedes network
innovation.[20]
In contrast, personal computers have a vastly diferent develop-
ment model, which involves them using a universal processor, based on
which they implement sofware-defned functions. Terefore, the com-
puter has a more fexible programming capability, resulting in explosive
growth in sofware applications. Furthermore, the open-source model
of computer sofware breeds a large amount of open-source sofware,
accelerates the sofware development process, and promotes the rapid
development of the entire computer industry. We can look at the Linux
open-source OS as the best example of this.
Drawing on universal hardware, sofware-defned functions, and the
open-source model in the computer feld, Professor Nick McKeown’s team
proposed a new network architecture — SDN.[4]
In the SDN architecture, the control plane of a network is separated
from its data plane: Te data plane becomes more generalized, similar
to the universal hardware of a computer. It no longer needs to specif-
cally implement the control logics of various network protocols; instead, it
only needs to execute the operation instructions received from the control
plane. Te control logic of a network device is defned by the SDN con-
troller and applications, and as such, network functions are defned by
sofware. With the emergence of open-source SDN controllers and open-
source SDN open interfaces, the network architecture includes three ele-
ments: universal hardware, sofware-defned support, and open-source
model. Figure 1.6 depicts the evolution from the traditional network
architecture to the SDN architecture.
SDN has the following three characteristics[19]:
1. Open network programmability: SDN provides a new network
abstraction model with a complete set of universal APIs for users,
who can then program on the controller to confgure, control, and
manage networks.
61.
14 ◾ SRv6Network Programming
FIGURE 1.6 Evolution from the traditional network architecture to the SDN
architecture.
2. Separation of the control and data planes: Separation refers to the
decoupled control and data planes, which can evolve independently
and communicate using a set of open APIs.
3. Logical centralized control: Tis refers to the centralized control of
distributed network states. Logical centralized control architecture
is the basis for SDN, propelling automated network control into the
realm of possibilities.
A network that has the preceding three characteristics can be called an
SDN network. Among the three characteristics, separation of the control
and data planes creates the necessary conditions for logical centralized
control, which in turn provides the architectural basis for open network
programmability. Note that open network programmability is the core
characteristics of SDN.
Tat being said, it is worth noting that SDN is only a network architec-
ture, and multiple technologies have been proposed to implement it, such
as OpenFlow, Protocol Oblivious Forwarding (POF),[21] Programming
Protocol-independent Packet Processors (P4),[22] and SR.[7]
1.4.1 OpenFlow
On March 14, 2008, Professor Nick McKeown and others proposed
OpenFlow, which is a protocol used between the SDN control plane and
data plane.
62.
SRv6 Background ◾15
FIGURE 1.7 OpenFlow 1.0 architecture.
In the OpenFlow protocol architecture shown in Figure 1.7, an
OpenFlow channel is established between an OpenFlow switch and
OpenFlow controller to exchange information. Te controller can deliver
fow table entries to the OpenFlow switch using the OpenFlow protocol.
Each fow table entry defnes a type of fow and corresponding forwarding
actions. Tat is, if the matching for a specifc type of fow succeeds, cor-
responding actions will be executed to process and forward packets.
Essentially, the forwarding behavior on a device can be abstracted into
“matching and action.” For example, Layer 2 switches and Layer 3 rout-
ers forward packets by searching for a destination Media Access Control
(MAC) address and destination IP address, respectively. Te design prin-
ciple of OpenFlow involves abstracting matching and forwarding actions
into specifc operations and using the controller to deliver fow table
entries to switches to guide packet forwarding. In short, OpenFlow is an
SDN technology that abstracts and generalizes network processing rules
and supports centralized programming.
In OpenFlow, a packet can be forwarded to a specifc interface or dis-
carded by matching protocol felds such as Ethernet, IPv4, or IPv6. Te
controller can program a matching+action rule via a fow table entry to
implement network programming. For example, the controller delivers a
fow table entry to switch A, instructing the switch to forward a packet
whose destination IP address is 192.168.1.20 to outbound interface 1.
OpenFlow’s advantage lies in the fexible programming of forwarding
rules, but that’s not to say it does not have apparent problems.
1. Te fow table entries are stored in the expensive Ternary Content
Addressable Memory (TCAM), meaning that only 1k–10k entries
can generally be supported in OpenFlow switches. Te limited
OpenFlow fow table specifcations result in OpenFlow switches
ofering insufcient scalability. Currently, OpenFlow is mainly
deployed in data centers for simple data switching and cannot be
deployed at a location requiring numerous fow table entries.
63.
16 ◾ SRv6Network Programming
2. Te advantage of an OpenFlow switch (used as a Layer 3 switch or
router) is that it can forward packets according to the fow table entry
generated by the central controller without requiring distributed
routing protocols, such as an Interior Gateway Protocol (IGP).
However, no carrier would like to abandon distributed routing pro-
tocols on their live networks, causing the distributed routing proto-
cols like IGPs to continue running on OpenFlow switches. In this
case, the basic shortest path forwarding can be performed by the IGP
while OpenFlow can only be used to optimize trafc steering, so the
benefts of OpenFlow are limited comparing to the cost brought by
it. Taking this into account, the OpenFlow switch does not simplify
the protocols, but on the contrary, it introduces additional complex-
ity brought by OpenFlow.
3. OpenFlow supports packet forwarding by adding corresponding
fow table entries based on the existing forwarding logic. It cannot
program the forwarding logic of a switch. For this reason, when
new features are added to OpenFlow, the protocol stack implemen-
tation on the controller and switch has to be updated. Sometimes,
even the switch’s chip and other hardware need to be redesigned.
Terefore, the cost of supporting new features is signifcant in
OpenFlow.
4. OpenFlow does not have sufcient capability to support stateful net-
work processing in the data plane that covers a full range of L4–L7
services. Terefore, the limited expressivity impacts the forward-
ing plane programmability. Te network states must be maintained
on the controller and synchronized up with the OpenFlow Switch.
Overdependence on the controller signifcantly burdens it, as well as
brings problems in extensibility and performance.
Due to the preceding limitations, OpenFlow has not been widely deployed.
1.4.2 POF
