---
id: collect-260926-mikrotik/mikrotik/how-to-configure-mikrotik-wan-failover-for-high-availability-complete-guide-for-network-en-2
title: "Configure interfaces"
domain: mikrotik
role: reference
task: reference
actors: ["Google"]
dates: []
keywords: ["cost", "latency", "license", "memory"]
source: docs/RAG/lot-mikrotik/forum/routing-bgp-ospf/how-to-configure-mikrotik-wan-failover-for-high-availability-complete-guide-for-network-engineers-te.md
source_anchor: ""
source_lines: [249, 509]
sha256: 6abb71990e1f212cb7939f517006e24b946b250321676afe3ab27a9b29f4fc36
---

# Configure interfaces

```
/system script
add name=primary-up source={
    :global currentActiveWAN
    :global primaryGW
    
    :if ($currentActiveWAN = "secondary") do={
        :log info "Primary WAN restored - switching back"
        
        # Enable primary route
        /ip route set [find comment="Primary Route"] disabled=no
        
        # Secondary route will be disabled by distance
        
        # Update global variable
        :set currentActiveWAN "primary"
        
        :log info "Failback to primary WAN completed"
    }
}
```
#### Configure Netwatch with Scripts

```
/tool netwatch
set [find comment="Monitor Primary"] up-script=primary-up down-script=primary-down
```
### Multi-Target Redundancy

Monitor multiple targets to prevent false positives:

#### Advanced Multi-Target Script

```
/system script
add name=advanced-failover source={
    :global primaryTargets {"8.8.8.8";"1.1.1.1";"208.67.222.222"}
    :global failedTargets 0
    :local maxFailures 2
    
    # Check each target
    :foreach target in=$primaryTargets do={
        :local pingResult [/ping $target count=3 interval=1s]
        :if ($pingResult = 0) do={
            :set failedTargets ($failedTargets + 1)
        }
    }
    
    # Trigger failover if threshold exceeded
    :if ($failedTargets >= $maxFailures) do={
        :execute script=primary-down
    }
    
    :set failedTargets 0
}
```
### Connection Tracking and Session Management

Manage existing connections during failover:

```
/system script
add name=connection-cleanup source={
    # Remove connections using failed interface
    :local failedInterface "wan1-primary"
    
    /ip firewall connection
    :foreach connection in=[find] do={
        :local connInterface [get $connection orig-interface]
        :if ($connInterface = $failedInterface) do={
            remove $connection
        }
    }
    
    :log info "Cleaned up connections for failed interface"
}
```
## 5. Method 3: Hardware Redundancy with VRRP

VRRP (Virtual Router Redundancy Protocol) creates virtual routers shared between multiple physical devices. This provides hardware-level redundancy.

### VRRP Fundamentals on MikroTik

#### How VRRP Works

- Multiple routers share a virtual IP address
- One router acts as master, others as backup
- Master election based on priority values
- Automatic failover when master fails

#### VRRP Prerequisites

- Two or more MikroTik routers
- Same subnet for VRRP interfaces
- RouterOS Advanced license or higher
- Synchronized time between devices

### Dual Router VRRP Configuration

#### Master Router Configuration

```
# Configure physical interface
/ip address
add address=192.168.1.10/24 interface=ether3 comment="Master LAN IP"
# Configure VRRP interface
/interface vrrp
add interface=ether3 vrid=1 priority=200 \
    interval=1s fast-leave=yes \
    name=vrrp-lan comment="Master VRRP"
# Assign virtual IP to VRRP interface
/ip address
add address=192.168.1.1/24 interface=vrrp-lan comment="Virtual Gateway"
# Configure WAN failover for master
/ip route
add dst-address=0.0.0.0/0 gateway=203.0.113.1 distance=1 check-gateway=ping
add dst-address=0.0.0.0/0 gateway=198.51.100.1 distance=2
```
#### Backup Router Configuration

```
# Configure physical interface
/ip address
add address=192.168.1.11/24 interface=ether3 comment="Backup LAN IP"
# Configure VRRP interface
/interface vrrp
add interface=ether3 vrid=1 priority=100 \
    interval=1s fast-leave=yes \
    name=vrrp-lan comment="Backup VRRP"
# Assign virtual IP to VRRP interface
/ip address
add address=192.168.1.1/24 interface=vrrp-lan comment="Virtual Gateway"
# Configure WAN failover for backup
/ip route
add dst-address=0.0.0.0/0 gateway=203.0.113.2 distance=1 check-gateway=ping
add dst-address=0.0.0.0/0 gateway=198.51.100.2 distance=2
```
### Combining VRRP with WAN Failover

#### Advanced VRRP Script Integration

```
/system script
add name=vrrp-wan-monitor source={
    :local vrrpStatus [/interface vrrp get [find name=vrrp-lan] running]
    :local primaryWAN [/ip route get [find comment="Primary Route"] active]
    
    # Adjust VRRP priority based on WAN status
    :if ($primaryWAN = false) do={
        # Primary WAN failed, reduce priority
        /interface vrrp set [find name=vrrp-lan] priority=150
        :log warning "Primary WAN failed, reduced VRRP priority"
    } else={
        # Primary WAN active, restore priority
        /interface vrrp set [find name=vrrp-lan] priority=200
        :log info "Primary WAN active, restored VRRP priority"
    }
}
# Schedule script execution
/system scheduler
add name=vrrp-monitoring interval=30s on-event=vrrp-wan-monitor
```
## 6. Load Balancing vs. Failover Considerations

### When to Use Load Balancing

Load balancing distributes traffic across multiple connections. Consider load balancing when:

- Both WAN connections have sufficient bandwidth
- Applications can handle asymmetric routing
- Cost optimization is important
- Bandwidth aggregation is needed

#### Basic Load Balancing Configuration

```
# Configure route-marking
/ip firewall mangle
add chain=input in-interface=wan1-primary action=mark-connection \
    new-connection-mark=wan1-conn passthrough=yes
add chain=input in-interface=wan2-secondary action=mark-connection \
    new-connection-mark=wan2-conn passthrough=yes
# Configure outbound load balancing
/ip firewall mangle
add chain=output connection-mark=wan1-conn action=mark-routing \
    new-routing-mark=to-wan1 passthrough=yes
add chain=output connection-mark=wan2-secondary action=mark-routing \
    new-routing-mark=to-wan2 passthrough=yes
# Create routing tables
/ip route
add dst-address=0.0.0.0/0 gateway=203.0.113.1 routing-mark=to-wan1
add dst-address=0.0.0.0/0 gateway=198.51.100.1 routing-mark=to-wan2
```
### Pure Failover Implementation

Pure failover maintains one active connection. Benefits include:

- Simpler configuration
- Predictable routing paths
- Lower bandwidth costs
- Easier troubleshooting

#### Bandwidth Conservation Setup

```
# Disable secondary interface when primary is active
/system script
add name=bandwidth-conservation source={
    :local primaryActive [/ip route get [find comment="Primary Route"] active]
    
    :if ($primaryActive = true) do={
        # Primary active - disable secondary DHCP client
        /ip dhcp-client set [find interface=wan2-secondary] disabled=yes
        :log info "Primary active - secondary WAN disabled"
    } else={
        # Primary failed - enable secondary DHCP client
        /ip dhcp-client set [find interface=wan2-secondary] disabled=no
        :log info "Primary failed - secondary WAN enabled"
    }
}
```
## 7. Monitoring and Troubleshooting

### Health Check Configuration

#### Optimal Monitoring Settings

- **Ping Interval** : 10-30 seconds (balance between responsiveness and overhead)
- **Timeout** : 2-5 seconds (account for network latency)
- **Failure Threshold** : 3-5 consecutive failures before triggering

#### Multiple Target Selection Strategy

Choose monitoring targets carefully:

- **Primary targets** : ISP DNS servers, gateway addresses
- **Secondary targets** : Public DNS (8.8.8.8, 1.1.1.1)
- **Tertiary targets** : Well-known websites (google.com, cloudflare.com)

```
/tool netwatch
add host=203.0.113.1 interval=15s timeout=3s comment="ISP1 Gateway"
add host=198.51.100.1 interval=15s timeout=3s comment="ISP2 Gateway" 
add host=8.8.8.8 interval=30s timeout=5s comment="Google DNS"
add host=1.1.1.1 interval=30s timeout=5s comment="Cloudflare DNS"
```
### Logging and Alerting

#### Configure System Logging

```
# Enable detailed logging
/system logging
add topics=info,warning,error action=memory prefix="FAILOVER"
add topics=route,script action=memory prefix="ROUTING"
# Configure log rotation
/system logging action
set memory memory-lines=5000 memory-stop-on-full=no
```
#### Email Notification Setup

