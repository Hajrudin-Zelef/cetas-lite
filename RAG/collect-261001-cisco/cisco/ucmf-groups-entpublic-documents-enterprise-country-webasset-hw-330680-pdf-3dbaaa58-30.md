---
id: collect-261001-cisco/cisco/ucmf-groups-entpublic-documents-enterprise-country-webasset-hw-330680-pdf-3dbaaa58-30
title: "A line starting with the # sign is comments."
domain: cisco
role: reference
task: reference
actors: ["Huawei"]
dates: ["2013-04-15"]
keywords: ["copyright"]
source: docs/RAG/collect-261001-cisco/ucmf-groups-entpublic-documents-enterprise-country-webasset-hw-330680-pdf-3dbaaa58.md
source_anchor: ""
source_lines: [4144, 4367]
sha256: 0b0781b56392207632fcfa3aa26eba4bb528f32e0cc1837acf6ee015cec5463c
---

# A line starting with the # sign is comments.

NOTE
On conditions when LSRA and LSRB are on the same CR-LSP:
l If GR is disabled and TE FRR is enabled, TE FRR switches traffic to a bypass CR-LSP after the
Hello extension detects that the RSVP neighbor relationship is lost to ensure proper traffic
transmission.
l If GR is enabled, the GR process is performed.
Deployment Scenarios
The RSVP Hello extension applies to networks enabled with both RSVP GR and TE FRR.
3.2.9.4 CR-LSP Backup
CR-LSP backup techniques protect E2E MPLS TE tunnels. If the ingress detects that the primary
CR-LSP is unavailable, the ingress switches traffic to a backup CR-LSP. After the primary CR-
LSP recovers, traffic switches back.
Related Concepts
CR-LSP backup functions include hot standby, ordinary backup, and the best-effort path
function.
l Hot-standby: The backup CR-LSP is set up immediately after the primary CR-LSP is set
up. If the primary CR-LSP fails, the backup CR-LSP takes over traffic from the primary
CR-LSP.
l Ordinary backup: When the primary CR-LSP fails, the backup CR-LSP is set up and takes
over traffic from the primary CR-LSP.
l Best-effort path: The failures of both primary and backup CR-LSPs can trigger the setup
of a temporary CR-LSP, which is called the best-effort path. Traffic is then switched to the
best-effort path.
For example, the primary CR-LSP is established over the path PE1 → P1 → P2 → PE2,
and the backup CR-LSP is established over the path PE1 → P3 → PE2 shown in Figure
3-21. If both CR-LSPs fail, PE1 establishes a best-effort path PE1 → P4 → PE2 to take
over traffic.
Figure 3-21 Best-effort path
Primary 
CR-LSP
Backup CR-LSP
Best-effort 
path
PE1 PE2
P3
P1 P2
P4
Enterprise Data Communication Products
Feature Description - MPLS 3 MPLS TE
Issue 04 (2013-04-15) Huawei Proprietary and Confidential
Copyright © Huawei Technologies Co., Ltd.
85

NOTE
A best-effort path has no bandwidth reserved for traffic, but has an affinity and a hop limit configured
as needed.
Implementation
The procedure of CR-LSP backup is as follows:
1. Planning is implemented.
Plan the paths, bandwidth values, and deployment modes. Table 3-15 lists CR-LSP backup
deployment items.
Table 3-15 CR-LSP backup deployment
It
e
m
Hot Standby Ordinary Backup Best-Effort
Path
Pa
th
Determine whether the primary
and hot-standby CR-LSPs
partially overlap. A hot-standby
CR-LSP can be established over an
explicit path.
The hot-standby CR-LSP supports
the following constraints:
l Explicit path
l Affinity
l Hop limit
l Overlapping Path for a Hot-
standby CR-LSP
The path of the backup
CR-LSP partially
overlaps the path of the
primary CR-LSP
regardless of whether the
backup CR-LSP is set up
along an explicit path.
The ordinary backup CR-
LSP supports the
following constraints:
l Explicit path
l Affinity
l Hop limit
Automatically
calculated by
the ingress.
The best-effort
path supports
the following
constraints:
l Affinity
l Hop limit
Ba
nd
wi
dt
h
A hot-standby CR-LSP and a
primary CR-LSP have the same
bandwidth by default. Dynamic
Bandwidth Protection for Hot-
standby CR-LSPs is supported
and ensures that a hot-standby CR-
LSP does not use additional
bandwidth when transmitting
traffic.
An ordinary backup CR-
LSP and a primary CR-
LSP have the same
bandwidth.
A best-effort
path is only a
protection path
that does not
have reserved
bandwidth.
De
pl
oy
m
en
t
m
od
e
Can be established without
attribute templates.
Can be established
without attribute
templates.
Can be
established
without
attribute
templates.
Enterprise Data Communication Products
Feature Description - MPLS 3 MPLS TE
Issue 04 (2013-04-15) Huawei Proprietary and Confidential
Copyright © Huawei Technologies Co., Ltd.
86

It
e
m
Hot Standby Ordinary Backup Best-Effort
Path
Can be established using attribute
templates.
Can be established using
attribute templates.
Automatically
established and
does not support
attribute
templates.
Co
nfi
gu
rat
io
n
co
m
bi
na
tio
n
l If established without an
attribute template, a hot-
standby CR-LSP can be used
together with a best-effort path.
l If established using an attribute
template, a hot-standby CR-
LSP can be used together with
both an ordinary backup CR-
LSP and a best-effort path.
l If established without
an attribute template,
an ordinary CR-LSP
can only be used alone.
l If established using an
attribute template, an
ordinary backup CR-
LSP can be used
together with a hot-
standby backup CR-
LSP and a best-effort
path.
-
 
2. CR-LSPs are established in sequence.
You can establish backup CR-LSPs using different modes on the same tunnel. To quickly
establish a CR-LSP for service transmission, the system attempts to establish a backup CR-
LSP using different modes in sequence until the backup CR-LSP is successfully established.
The rules for establishing a CR-LSP are as follows:
a. If new tunnel configuration is committed or a tunnel goes Down, the ingress first
attempts to establish a primary CR-LSP. If the attempt fails, the ingress attempts to
establish a hot-standby CR-LSP. If establishing the hot-standby CR-LSP fails, the
ingress then attempts to establish an ordinary backup CR-LSP. If this attempt also
fails, the ingress establishes a best-effort path.
b. A maximum of three CR-LSP attribute templates can be configured for hot-standby
CR-LSPs and three for ordinary backup CR-LSPs. These templates are prioritized.
The ingress uses each in descending order by priority until a CR-LSP is successfully
established.
c. If a CR-LSP has been established using a lower-priority attribute template and the
CR-LSP status changes, the ingress will attempt to establish a CR-LSP using a higher-
priority attribute template. The Make-Before-Break mechanism ensures that traffic is
uninterrupted when a new CR-LSP is being established.
d. If a stable CR-LSP has been established using any of the attribute templates, you can
lock the used backup CR-LSP attribute template. After the attribute template is locked,
the ingress will not attempt to use a higher-priority attribute template to establish a
CR-LSP. This locking function prevents unnecessary traffic switchovers and lowers
system costs.
3. Backup CR-LSP attributes are modified.
When the constraints for backup CR-LSPs are modified, the ingress triggers re-
establishment of a backup CR-LSP. The system uses the Make-Before-Break mechanism
to re-establish a backup CR-LSP. After that backup CR-LSP has been successfully
Enterprise Data Communication Products
Feature Description - MPLS 3 MPLS TE
Issue 04 (2013-04-15) Huawei Proprietary and Confidential
Copyright © Huawei Technologies Co., Ltd.
87

