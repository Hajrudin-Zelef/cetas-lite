---
id: collect-260926-mikrotik/mikrotik/arborescence-des-pages-76-1
title: "Overview"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: ["license", "parameters"]
source: docs/RAG/lot-mikrotik/RouterOS/arborescence-des-pages-76.md
source_anchor: ""
source_lines: [1, 12]
sha256: 807ce33048c9302e778a043ee57b40192a8b93fcc72bb2cfd02d2eaea10da6cb
---

# Overview

User Manager is RADIUS server implementation in RouterOS which provides centralized user authentication and authorization to a certain service. Having a central user database allows better tracking of system users and customers. As a separate package, User Manager is available on all architectures except SMIPS, however, care must be taken due to limited free space available. It supports many different authentication methods including PAP, CHAP, MS-CHAP, MS-CHAPv2, EAP-TLS, EAP-TTLS, and EAP-PEAP. In RouterOS, DHCP, Dot1x, Hotspot, IPsec, PPP, and Wireless are features that benefit from User Manager the most. Each user can see their account statistics and manage available profiles using the WEB interface. Additionally, users can buy their own data plans (profiles) using the most popular payment gateway - PayPal making it a great system for service providers. Customized reports can be generated to ease processing by the billing department. User Manager works according to RADIUS standards defined in RFC2865 and RFC3579.

User Manager is one of the RouterOS features, that is limited by the RouterOS license level. Depending on the License level, number of active sessions will be limited, including multiple connections per user (not unique accounts).

# Attributes

**Sub-menu:** `/user-manager attribute`

RADIUS attributes are defined authorization, information, and configuration parameters that are passed between the RADIUS server and the client. User Manager allows sending customized attributes defined in the "attributes" menu. RouterOS has a set of predefined attributes already present, but it is also possible to add additional attributes if necessary. Predefined attributes:

