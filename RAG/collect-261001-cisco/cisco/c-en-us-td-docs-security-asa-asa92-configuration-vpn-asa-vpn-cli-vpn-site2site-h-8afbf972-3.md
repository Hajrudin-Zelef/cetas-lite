---
id: collect-261001-cisco/cisco/c-en-us-td-docs-security-asa-asa92-configuration-vpn-asa-vpn-cli-vpn-site2site-h-8afbf972-3
title: "c-en-us-td-docs-security-asa-asa92-configuration-vpn-asa-vpn-cli-vpn-site2site-h-8afbf972"
domain: cisco
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-cisco/c-en-us-td-docs-security-asa-asa92-configuration-vpn-asa-vpn-cli-vpn-site2site-h-8afbf972.md
source_anchor: ""
source_lines: [118, 149]
sha256: 329ed58f7c275079efb2897e5023897d0557db094a2a5831069473e8cd2ab8b8
---

# c-en-us-td-docs-security-asa-asa92-configuration-vpn-asa-vpn-cli-vpn-site2site-h-8afbf972

- The crypto map entries must have at least one transform set in common.
If you create more than one crypto map entry for a given interface, use the sequence number (seq-num) of each entry to rank it: the lower the seq-num, the higher the priority. At the interface that has the crypto map set, the ASA evaluates traffic against the entries of higher priority maps first.
Create multiple crypto map entries for a given interface if either of the following conditions exist:
- Different peers handle different data flows.
- You want to apply different IPsec security to different types of traffic (to the same or separate peers), for example, if you want traffic between one set of subnets to be authenticated, and traffic between another set of subnets to be both authenticated and encrypted. In this case, define the different types of traffic in two separate ACLs, and create a separate crypto map entry for each crypto ACL.
To create a crypto map and apply it to the outside interface in global configuration mode, perform the following steps in either single or multiple context mode:
Step 1 To assign an ACL to a crypto map entry, enter the crypto map match address command.
The syntax is crypto map map-name seq-num match address aclname . In the following example the map name is abcmap, the sequence number is 1, and the ACL name is l2l_list .
Step 2 To identify the peer (s) for the IPsec connection, enter the crypto map set peer command.
The syntax is crypto map map-name seq-num set peer { ip_address1 | hostname1 }[... ip_address10 | hostname10 ]. In the following example the peer name is 10.10.4.108.
Step 3 To specify an IKEv1 transform set for a crypto map entry, enter the crypto map ikev1 set transform-set command.
The syntax is 
crypto map 
map-name seq-num 
ikev1 set transform-set 
transform-set-name
. 
In the following example the transform set name is FirstSet.
Step 4 To specify an IKEv2 proposal for a crypto map entry, enter the crypto map ikev2 set ipsec-proposal command:
The syntax is 
crypto map 
map-name seq-num set 
ikev2 ipsec-proposal proposal-name
. 
In the following example the proposal name is secure.
With the crypto map command, you can specify multiple IPsec proposals for a single map index. In that case, multiple proposals are transmitted to the IKEv2 peer as part of the negotiation, and the order of the proposals is determined by the administrator upon the ordering of the crypto map entry.
Note If combined mode (AES-GCM/GMAC) and normal mode (all others) algorithms exist in the IPsec proposal, then you cannot send a single proposal to the peer. You must have at least two proposals in this case, one for combined mode and one for normal mode algorithms.
Applying Crypto Maps to Interfaces
You must apply a crypto map set to each interface through which IPsec traffic travels. The ASA supports IPsec on all interfaces. Applying the crypto map set to an interface instructs the ASA to evaluate all interface traffic against the crypto map set and to use the specified policy during connection or security association negotiations.
Binding a crypto map to an interface also initializes the runtime data structures, such as the security association database and the security policy database. When you later modify a crypto map in any way, the ASA automatically applies the changes to the running configuration. It drops any existing connections and reestablishes them after applying the new crypto map.
To apply the configured crypto map to the outside interface, perform the following steps:
Step 1 Enter the crypto map interface command. The syntax is crypto map map-name interface interface-name.
