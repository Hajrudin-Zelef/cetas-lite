---
id: collect-261001-meraki/meraki/meraki-api-v1-api-index-21d30cd4-9
title: "meraki-api-v1-api-index-21d30cd4"
domain: meraki
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-meraki/meraki-api-v1-api-index-21d30cd4.md
source_anchor: ""
source_lines: [187, 211]
sha256: ac49109b9756e91c580bd1c4463f7f86b57cb1e595562399d5b7e69330d28ce6
---

# meraki-api-v1-api-index-21d30cd4

|  | PUT /networks/{networkId}/appliance/firewall/l3FirewallRules Update the L3 firewall rules of an MX network  >  updateNetworkApplianceFirewallL3FirewallRules | networkId | comment, destCidr, destPort, policy, protocol, rules, srcCidr, srcPort, syslogDefaultRule, syslogEnabled | comment, destCidr, destPort, policy, protocol, rules, srcCidr, srcPort, syslogEnabled | sdwan:config:write | 
|  | GET /networks/{networkId}/appliance/firewall/l7FirewallRules List the MX L7 firewall rules for an MX network  >  getNetworkApplianceFirewallL7FirewallRules | networkId | `` | policy, rules, type, value | sdwan:config:read | 
|  | PUT /networks/{networkId}/appliance/firewall/l7FirewallRules Update the MX L7 firewall rules for an MX network  >  updateNetworkApplianceFirewallL7FirewallRules | networkId | policy, rules, type, value | policy, rules, type, value | sdwan:config:write | 
|  | GET /networks/{networkId}/appliance/firewall/l7FirewallRules/applicationCategories Return the L7 firewall application categories and their associated applications for an MX network  >  getNetworkApplianceFirewallL7FirewallRulesApplicationCategories | networkId | `` | applicationCategories, applications, id, name | sdwan:config:read | 
|  | PUT /networks/{networkId}/appliance/firewall/multicastForwarding Update static multicast forward rules for a network  >  updateNetworkApplianceFirewallMulticastForwarding | networkId | address, description, rules, vlanIds | address, description, id, name, network, rules, vlanIds | sdwan:config:write | 
|  | GET /networks/{networkId}/appliance/firewall/oneToManyNatRules Return the 1:Many NAT mapping rules for an MX network  >  getNetworkApplianceFirewallOneToManyNatRules | networkId | `` | allowedIps, localIp, localPort, name, portRules, protocol, publicIp, publicPort, rules, uplink | sdwan:config:read | 
|  | PUT /networks/{networkId}/appliance/firewall/oneToManyNatRules Set the 1:Many NAT mapping rules for an MX network  >  updateNetworkApplianceFirewallOneToManyNatRules | networkId | allowedIps, localIp, localPort, name, portRules, protocol, publicIp, publicPort, rules, uplink | allowedIps, localIp, localPort, name, portRules, protocol, publicIp, publicPort, rules, uplink | sdwan:config:write | 
|  | GET /networks/{networkId}/appliance/firewall/oneToOneNatRules Return the 1:1 NAT mapping rules for an MX network  >  getNetworkApplianceFirewallOneToOneNatRules | networkId | `` | allowedInbound, allowedIps, destinationPorts, lanIp, name, protocol, publicIp, rules, uplink | sdwan:config:read | 
|  | PUT /networks/{networkId}/appliance/firewall/oneToOneNatRules Set the 1:1 NAT mapping rules for an MX network  >  updateNetworkApplianceFirewallOneToOneNatRules | networkId | allowedInbound, allowedIps, destinationPorts, lanIp, name, protocol, publicIp, rules, uplink | allowedInbound, allowedIps, destinationPorts, lanIp, name, protocol, publicIp, rules, uplink | sdwan:config:write | 
|  | GET /networks/{networkId}/appliance/firewall/portForwardingRules Return the port forwarding rules for an MX network  >  getNetworkApplianceFirewallPortForwardingRules | networkId | `` | allowedIps, lanIp, localPort, name, protocol, publicPort, rules, uplink | sdwan:config:read | 
|  | PUT /networks/{networkId}/appliance/firewall/portForwardingRules Update the port forwarding rules for an MX network  >  updateNetworkApplianceFirewallPortForwardingRules | networkId | allowedIps, lanIp, localPort, name, protocol, publicPort, rules, uplink | allowedIps, lanIp, localPort, name, protocol, publicPort, rules, uplink | sdwan:config:write | 
|  | GET /networks/{networkId}/appliance/firewall/settings Return the firewall settings for this network  >  getNetworkApplianceFirewallSettings | networkId | `` | ipSourceGuard, mode, spoofingProtection | sdwan:config:read | 
|  | PUT /networks/{networkId}/appliance/firewall/settings Update the firewall settings for this network  >  updateNetworkApplianceFirewallSettings | networkId | ipSourceGuard, mode, spoofingProtection | ipSourceGuard, mode, spoofingProtection | sdwan:config:write | 
|  | POST /networks/{networkId}/appliance/interfaces/l3 Create wired L3 interface  >  createNetworkApplianceInterfacesL3 | networkId | address, id, interface, ipv4, name, number, port, slot, subnet, subslot, vrf | address, id, interface, interfaceId, ipv4, name, number, port, slot, subnet, subslot, vrf | `` | 
|  | PUT /networks/{networkId}/appliance/interfaces/l3/{interfaceId} Update wired L3 interface  >  updateNetworkApplianceInterfacesL3 | networkId, interfaceId | address, id, interface, ipv4, name, number, port, slot, subnet, subslot, vrf | address, id, interface, interfaceId, ipv4, name, number, port, slot, subnet, subslot, vrf | `` | 
|  | DELETE /networks/{networkId}/appliance/interfaces/l3/{interfaceId} Delete wired L3 interface  >  deleteNetworkApplianceInterfacesL3 | networkId, interfaceId | `` | `` | `` | 
|  | GET /networks/{networkId}/appliance/ports List per-port VLAN settings for all ports of a secure router or security appliance.  >  getNetworkAppliancePorts | networkId | `` | accessPolicy, adaptivePolicyGroupId, allowedVlans, dropUntaggedTraffic, enabled, id, number, peerSgtCapable, sgt, type, vlan | sdwan:config:read | 
|  | POST /networks/{networkId}/appliance/ports/radius/servers Create a shared MX port RADIUS server for a network (BETA)  >  createNetworkAppliancePortsRadiusServer | networkId | host, port, secret, value | host, id, name, network, port, serverId | sdwan:config:write | 
|  | PUT /networks/{networkId}/appliance/ports/radius/servers/{serverId} Update a shared MX port RADIUS server for a network (BETA)  >  updateNetworkAppliancePortsRadiusServer | networkId, serverId | host, port, secret, value | host, id, name, network, port, serverId | sdwan:config:write | 
|  | DELETE /networks/{networkId}/appliance/ports/radius/servers/{serverId} Delete a network-owned shared MX port RADIUS server (BETA)  >  deleteNetworkAppliancePortsRadiusServer | networkId, serverId | `` | `` | sdwan:config:write | 
|  | GET /networks/{networkId}/appliance/ports/{portId} Return per-port VLAN settings for a single secure router or security appliance port.  >  getNetworkAppliancePort | networkId, portId | `` | accessPolicy, adaptivePolicyGroupId, allowedVlans, dropUntaggedTraffic, enabled, id, number, peerSgtCapable, sgt, type, vlan | sdwan:config:read | 
|  | PUT /networks/{networkId}/appliance/ports/{portId} Update the per-port VLAN settings for a single secure router or security appliance port.  >  updateNetworkAppliancePort | networkId, portId | accessPolicy, adaptivePolicyGroupId, allowedVlans, dropUntaggedTraffic, enabled, id, peerSgtCapable, sgt, type, vlan | accessPolicy, adaptivePolicyGroupId, allowedVlans, dropUntaggedTraffic, enabled, id, number, peerSgtCapable, sgt, type, vlan | sdwan:config:write | 
|  | GET /networks/{networkId}/appliance/prefixes/delegated/statics List static delegated prefixes for a network  >  getNetworkAppliancePrefixesDelegatedStatics | networkId | `` | createdAt, description, interfaces, origin, prefix, staticDelegatedPrefixId, type, updatedAt | sdwan:config:read | 
|  | POST /networks/{networkId}/appliance/prefixes/delegated/statics Add a static delegated prefix from a network  >  createNetworkAppliancePrefixesDelegatedStatic | networkId | description, interfaces, origin, prefix, type | createdAt, description, interfaces, origin, prefix, staticDelegatedPrefixId, type, updatedAt | sdwan:config:write | 
|  | GET /networks/{networkId}/appliance/prefixes/delegated/statics/{staticDelegatedPrefixId} Return a static delegated prefix from a network  >  getNetworkAppliancePrefixesDelegatedStatic | networkId, staticDelegatedPrefixId | `` | createdAt, description, interfaces, origin, prefix, staticDelegatedPrefixId, type, updatedAt | sdwan:config:read | 
