---
id: collect-260926-mikrotik/mikrotik/how-to-configure-mikrotik-wan-failover-for-high-availability-complete-guide-for-network-en-5
title: "Configure interfaces"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: ["cost"]
source: docs/RAG/lot-mikrotik/forum/routing-bgp-ospf/how-to-configure-mikrotik-wan-failover-for-high-availability-complete-guide-for-network-engineers-te.md
source_anchor: ""
source_lines: [980, 1076]
sha256: 8455a102a10f7f88aae7a50b71971ad245ce45f046a79bb43974845f21df9dbb
---

# Configure interfaces

```
# Enterprise-grade configuration with VRRP
# Master router configuration
/interface bridge
add name=bridge-lan protocol-mode=rstp
/interface bridge port
add bridge=bridge-lan interface=ether3
add bridge=bridge-lan interface=ether4
# VRRP configuration for hardware redundancy
/interface vrrp
add interface=bridge-lan vrid=10 priority=200 name=vrrp-main
/ip address
add address=10.10.10.1/24 interface=vrrp-main
# Advanced routing with multiple ISPs
/ip route
add dst-address=0.0.0.0/0 gateway=203.0.113.1 distance=1 \
    routing-mark=isp1-route check-gateway=ping
add dst-address=0.0.0.0/0 gateway=198.51.100.1 distance=1 \
    routing-mark=isp2-route check-gateway=ping
# Load balancing with failover capability
/ip firewall mangle
add chain=prerouting dst-address-type=!local per-connection-classifier=both-addresses:2/0 \
    action=mark-connection new-connection-mark=isp1-conn passthrough=yes
add chain=prerouting dst-address-type=!local per-connection-classifier=both-addresses:2/1 \
    action=mark-connection new-connection-mark=isp2-conn passthrough=yes
# Comprehensive logging for compliance
/system logging
add topics=info,warning,error,critical action=remote remote=10.10.10.100:514
add topics=firewall action=disk file-name=firewall-log file-lines=10000
```
**Results**:

- Average failover time: 8 seconds
- 99.99% uptime achieved
- Successful regulatory audits
- Scalable architecture supporting growth

## 13. Conclusion and Next Steps

### Key Takeaways

Implementing MikroTik WAN failover provides significant business value through improved network reliability. The three methods covered offer solutions for different scenarios:

- **Distance-based routing** : Best for simple, cost-effective implementations
- **Netwatch with scripts** : Optimal for customized enterprise requirements
- **VRRP** : Essential for mission-critical applications requiring hardware redundancy

#### Critical Configuration Points

1. Choose monitoring targets carefully to avoid false positives
2. Set appropriate timeouts balancing responsiveness and stability
3. Implement proper logging for troubleshooting and compliance
4. Test failover procedures regularly to ensure functionality
5. Document configurations and maintain current backups

#### Common Pitfalls to Avoid

- Using overly aggressive monitoring intervals that cause system overhead
- Forgetting to configure NAT rules for backup connections
- Ignoring asymmetric routing issues in complex topologies
- Inadequate testing of failover and failback procedures
- Poor choice of monitoring targets leading to unnecessary failovers

### Further Learning Resources

#### MikroTik Certification Paths

- **MTCNA (MikroTik Certified Network Associate)** : Foundation certification covering basic RouterOS concepts
- **MTCRE (MikroTik Certified Routing Engineer)** : Advanced routing and failover techniques
- **MTCWE (MikroTik Certified Wireless Engineer)** : Wireless failover implementations

#### Advanced RouterOS Features

- BGP routing for enterprise networks
- MPLS implementation for service providers
- Advanced scripting techniques
- Network monitoring and management tools

#### Community Resources

- **MikroTik Wiki** : Comprehensive documentation and examples
- **MikroTik Forum** : Community support and troubleshooting
- **RouterOS Scripting** : Advanced automation techniques
- **Third-party tools** : PRTG, Zabbix, and other monitoring solutions

Regular practice with different scenarios and staying updated with RouterOS releases ensures optimal network reliability. Consider implementing monitoring dashboards to track failover events and performance metrics for continuous improvement.

#### Next Implementation Steps

1. Assess current network requirements and constraints
2. Choose appropriate failover method based on analysis
3. Create detailed implementation plan with rollback procedures
4. Set up test environment to validate configuration
5. Implement during planned maintenance window
6. Monitor and optimize based on real-world performance

Check our list of MikroTik guides.
