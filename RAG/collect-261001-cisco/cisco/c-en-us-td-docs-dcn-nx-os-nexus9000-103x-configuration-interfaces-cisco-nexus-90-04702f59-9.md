---
id: collect-261001-cisco/cisco/c-en-us-td-docs-dcn-nx-os-nexus9000-103x-configuration-interfaces-cisco-nexus-90-04702f59-9
title: "c-en-us-td-docs-dcn-nx-os-nexus9000-103x-configuration-interfaces-cisco-nexus-90-04702f59"
domain: cisco
role: reference
task: reference
actors: []
dates: []
keywords: ["parameters"]
source: docs/RAG/collect-261001-cisco/c-en-us-td-docs-dcn-nx-os-nexus9000-103x-configuration-interfaces-cisco-nexus-90-04702f59.md
source_anchor: ""
source_lines: [1209, 1251]
sha256: 4f45333d9ea96daf05dcaefa9e658cd88b91f08cf084396e5b6b64f2a7e33a56
---

# c-en-us-td-docs-dcn-nx-os-nexus9000-103x-configuration-interfaces-cisco-nexus-90-04702f59

using the port-channel load-balance command. Only Cisco Nexus 9364C and 9300-EX/FX/FX2 platform switches support this option. Beginning with Cisco NX-OS Release 9.3(5), Cisco Nexus N9K-C9316D-GX, N9K-C93600CD-GX, N9K-C9364C-GX switches support this
option.
The udf option includes the user-defined field for computing ECMP hashing. You can configure the offset base and the length of the
UDF field (in bits). The range for the offset base is from 0 to 127 bytes. The range for the length of the UDF field is from
1 to 32 bits. Enabling or disabling this option also enables or disables it for port-channel load-balancing if Layer 4 parameters
are enabled using the port-channel load-balance command. Only Cisco Nexus 9364C and 9300-EX/FX/FX2 platform switches support this option. Beginning with Cisco NX-OS Release 9.3(5), Cisco Nexus N9K-C9316D-GX, N9K-C93600CD-GX, N9K-C9364C-GX switches support this
option.
The symmetric option enables symmetric hashing globally. To disable ECMP symmetric hashing, use the no keyword in the command. You must execute this command in global configuration mode.
Note
Ensure that the configured universal-id seed value is consistent across the nodes in the path of ECMP symmetric hashing for symmetric hashing to work effectively.
The inner option enables inner header based hashing for GRE traffic globally. To disable inner header based hashing, use the no keyword in the command. You must execute this command in global configuration mode.
all: Configuring this option for GRE encapsulated packets starts using inner headers to hash onto a path in ECMP, which may impact
other encapsulation types as well. This is supported on Cisco Nexus 9364C and 9300-EX/FX/FX2 platform switches; and Cisco
Nexus 9500 platform switches with X9700-EX/FX line cards.
greheader: Configuring this option only for GRE encapsulated packets, starts using inner headers to hash onto a path in ECMP. This
is supported on Cisco Nexus 9364C and 9300-FX/FX2 platform switches; and Cisco Nexus 9500 platform switches with X9700-FX
line cards.
The following options are available for all IP load sharing configurations:
The universal-id option sets the random seed for the hash algorithm and shifts the flow from one link to another.
You do not need to configure the universal ID. Cisco NX-OS chooses the universal ID if you do not configure it. The universal-id range is from 1 to 4294967295.
The rotate option causes the hash algorithm to rotate the link picking selection so that it does not continually choose the same link
across all nodes in the network. It does so by influencing the bit pattern for the hash algorithm. This option shifts the
flow from one link to another and load balances the already load-balanced (polarized) traffic from the first ECMP level across
multiple links.
If you specify a rotate value, the 64-bit stream is interpreted starting from that bit position in a cyclic rotation. The rotate range is from 1 to 63, and the default is 32.
Note
With multi-tier Layer 3 topology, polarization is possible. To avoid polarization, use a different rotate bit at each tier
of the topology.
Note
To configure a rotation value for port channels, use the port-channel load-balance src-dst ip-l4port rotaterotate command.
The concatenation option ties together the hash tag values for ECMP and the hash tag values for port channels in order to use a stronger 64-bit
hash. If you do not use this option, you can control ECMP load-balancing and port-channel load-balancing independently. The
default is disabled.
Step 2
(Optional)
show ip load-sharing
Example:
switch(config)# show ip load-sharing
address source-destination
(Optional)
Displays the ECMP load-sharing algorithm for data traffic.
Verifying the ECMP Resilient Hashing Configuration
To display ECMP Resilient Hashing configuration information, perform one of the following tasks:
