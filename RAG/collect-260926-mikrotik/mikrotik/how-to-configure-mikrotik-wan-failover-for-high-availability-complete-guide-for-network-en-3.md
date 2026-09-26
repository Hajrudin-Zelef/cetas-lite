---
id: collect-260926-mikrotik/mikrotik/how-to-configure-mikrotik-wan-failover-for-high-availability-complete-guide-for-network-en-3
title: "Configure interfaces"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: ["cost", "ethernet", "memory"]
source: docs/RAG/lot-mikrotik/forum/routing-bgp-ospf/how-to-configure-mikrotik-wan-failover-for-high-availability-complete-guide-for-network-engineers-te.md
source_anchor: ""
source_lines: [510, 767]
sha256: a14b2a3865ff7c30c40f2bfd291b9608842ec3c273dd37b5bccc617c3aa12f1a
---

# Configure interfaces

```
# Configure email settings
/tool e-mail
set server=smtp.company.com port=587 \
    start-tls=yes user=alerts@company.com password=secret123
# Create notification script
/system script
add name=send-failover-alert source={
    :local message "WAN Failover Event: Primary connection failed at $[/system clock get date] $[/system clock get time]"
    /tool e-mail send to=netadmin@company.com subject="Network Alert" body=$message
}
```
### Common Issues and Solutions

#### Asymmetric Routing Problems

**Symptoms**:

- Intermittent connection issues
- Some applications work, others don’t
- Connection timeouts

**Solution**:

```
# Mark connections by input interface
/ip firewall mangle
add chain=prerouting in-interface=wan1-primary action=mark-connection \
    new-connection-mark=wan1-conn passthrough=yes
add chain=prerouting in-interface=wan2-secondary action=mark-connection \
    new-connection-mark=wan2-conn passthrough=yes
# Route responses back through same interface
/ip firewall mangle
add chain=output connection-mark=wan1-conn action=mark-routing \
    new-routing-mark=to-wan1 passthrough=yes
add chain=output connection-mark=wan2-conn action=mark-routing \
    new-routing-mark=to-wan2 passthrough=yes
```
#### DNS Resolution During Failover

**Problem**: DNS queries fail during connection switching

**Solution**:

```
# Configure DNS servers for both connections
/ip dns
set servers=8.8.8.8,8.8.4.4,1.1.1.1,1.0.0.1 \
    allow-remote-requests=yes cache-size=4096KiB
```
#### Troubleshooting Commands

```
# Check route status
/ip route print detail where active=yes
# Monitor gateway reachability
/tool netwatch print detail
# View connection tracking
/ip firewall connection print where connection-mark!=""
# Check interface statistics
/interface print stats
# Monitor system resources
/system resource print
# View recent logs
/log print where topics~"route|script"
```
## 8. Advanced Scenarios and Best Practices

### Multi-WAN Failover (3+ Links)

#### Three-WAN Configuration

```
# Configure interfaces
/interface ethernet
set [find default-name=ether1] name=wan1-fiber
set [find default-name=ether2] name=wan2-cable  
set [find default-name=ether3] name=wan3-lte
set [find default-name=ether4] name=lan1
# Set up DHCP clients
/ip dhcp-client
add interface=wan1-fiber disabled=no comment="Fiber Primary"
add interface=wan2-cable disabled=no comment="Cable Secondary"
# Configure LTE as tertiary
/interface lte
set [find name=lte1] allow-roaming=no
# Configure routes with distance prioritization
/ip route
add dst-address=0.0.0.0/0 gateway=[get wan1-fiber gateway] distance=1 check-gateway=ping
add dst-address=0.0.0.0/0 gateway=[get wan2-cable gateway] distance=2 check-gateway=ping
add dst-address=0.0.0.0/0 gateway=[get lte1 gateway] distance=3 check-gateway=ping
```
#### Advanced Priority Management

```
/system script
add name=multi-wan-manager source={
    :local wanStatus {"fiber"=false; "cable"=false; "lte"=false}
    
    # Check each WAN connection
    :set ($wanStatus->"fiber") [/ip route get [find comment="Fiber Primary"] active]
    :set ($wanStatus->"cable") [/ip route get [find comment="Cable Secondary"] active]  
    :set ($wanStatus->"lte") [/ip route get [find comment="LTE Tertiary"] active]
    
    # Log current status
    :log info ("WAN Status - Fiber: " . ($wanStatus->"fiber") . \
              " Cable: " . ($wanStatus->"cable") . \
              " LTE: " . ($wanStatus->"lte"))
    
    # Implement cost-based routing for LTE
    :if (($wanStatus->"fiber") = false and ($wanStatus->"cable") = false) do={
        :log warning "Using expensive LTE connection - consider bandwidth limits"
    }
}
```
### VPN Integration

#### Site-to-Site VPN Failover

```
# Create IPSec policies for both WAN interfaces
/ip ipsec policy
add dst-address=10.0.2.0/24 src-address=10.0.1.0/24 \
    out-interface=wan1-primary template=yes
add dst-address=10.0.2.0/24 src-address=10.0.1.0/24 \
    out-interface=wan2-secondary template=yes disabled=yes
# Script to manage VPN failover
/system script
add name=vpn-failover source={
    :local primaryActive [/ip route get [find comment="Primary Route"] active]
    
    :if ($primaryActive = false) do={
        # Primary failed - switch VPN to secondary
        /ip ipsec policy set [find out-interface=wan1-primary] disabled=yes
        /ip ipsec policy set [find out-interface=wan2-secondary] disabled=no
        :log info "VPN switched to secondary WAN"
    } else={
        # Primary restored - switch VPN back
        /ip ipsec policy set [find out-interface=wan1-primary] disabled=no
        /ip ipsec policy set [find out-interface=wan2-secondary] disabled=yes  
        :log info "VPN switched back to primary WAN"
    }
}
```
### Enterprise Implementation Guidelines

#### Change Management Procedures

1. **Pre-implementation Testing**  - Test configuration in lab environment
  - Validate failover timing requirements
  - Document expected behavior
  - Create rollback procedures
2. **Implementation Planning**  - Schedule during maintenance windows
  - Notify stakeholders of potential brief outages
  - Prepare monitoring tools
  - Have technical support available
3. **Post-implementation Validation**  - Verify failover functionality
  - Test application connectivity
  - Monitor system performance
  - Update documentation

#### Documentation Requirements

- Network topology diagrams
- Configuration backups with version control
- Monitoring target lists and rationale
- Emergency contact procedures
- Troubleshooting runbooks

## 9. Performance Optimization

### Failover Time Minimization

#### Optimal Monitoring Intervals

Balance between quick detection and system overhead:

- **Critical networks** : 5-10 second intervals
- **Standard networks** : 15-30 second intervals
- **Low priority** : 60+ second intervals

```
# High-performance monitoring configuration
/tool netwatch
add host=8.8.8.8 interval=5s timeout=2s startup-delay=30s \
    up-script="primary-restore" down-script="primary-failure"
```
#### Route Convergence Optimization

```
# Optimize routing table processing
/ip route
add dst-address=0.0.0.0/0 gateway=203.0.113.1 distance=1 \
    check-gateway=ping timeout=1s comment="Optimized Primary"
    
# Reduce ARP timeout for faster detection
/ip arp
set timeout=00:01:00
```
### Resource Management

#### CPU and Memory Considerations

Monitor system resources during failover operations:

```
# Create resource monitoring script
/system script
add name=resource-monitor source={
    :local cpuLoad [/system resource get cpu-load]
    :local freeMemory [/system resource get free-memory]
    :local totalMemory [/system resource get total-memory]
    :local memoryUsage (100 - (($freeMemory * 100) / $totalMemory))
    
    :if ($cpuLoad > 80) do={
        :log warning "High CPU load detected: $cpuLoad%"
    }
    
    :if ($memoryUsage > 80) do={
        :log warning "High memory usage detected: $memoryUsage%"
    }
}
# Schedule resource monitoring
/system scheduler
add name=resource-check interval=5m on-event=resource-monitor
```
#### Script Execution Optimization

```
# Optimized failover script with error handling
/system script
add name=optimized-failover source={
    :local startTime [/system clock get time]
    
    :do {
        # Disable primary route
        /ip route set [find comment="Primary Route"] disabled=yes
        
        # Clear connection tracking for faster convergence  
        /ip firewall connection remove [find orig-interface=wan1-primary]
        
        # Force routing table update
        /ip route check-active
        
        :local endTime [/system clock get time]
        :log info "Failover completed in $([:totime ($endTime - $startTime)])"
        
    } on-error={
        :log error "Failover script failed - manual intervention required"
    }
}
```
## 10. Security Considerations

### Firewall Rule Management

#### Interface-Specific Security Rules

