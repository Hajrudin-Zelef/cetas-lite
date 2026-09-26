---
id: collect-260926-mikrotik/mikrotik/how-to-configure-mikrotik-wan-failover-for-high-availability-complete-guide-for-network-en-1
title: "Configure interfaces"
domain: mikrotik
role: reference
task: reference
actors: ["Google"]
dates: []
keywords: ["cost", "ethernet", "parameters"]
source: docs/RAG/lot-mikrotik/forum/routing-bgp-ospf/how-to-configure-mikrotik-wan-failover-for-high-availability-complete-guide-for-network-engineers-te.md
source_anchor: ""
source_lines: [1, 248]
sha256: 88d7fefae0817d06f03f7b54cf0245ef8bbea0f2515ecc5c862d8bdf0bb5c8f4
---

# Configure interfaces

Failover is must in networks, read how to configure MikroTik WAN failover for high availability of your network infrastructure.

## Table of Contents

## 1. Introduction

Network downtime costs businesses an average of $5,600 per minute. Single internet connections create dangerous points of failure. MikroTik RouterOS provides multiple failover methods to ensure continuous network availability.

This guide shows network engineers and systems administrators how to configure WAN failover on MikroTik devices. You will learn three proven methods:

- Distance-based routing failover
- Netwatch-based monitoring and automation
- VRRP hardware redundancy

Each method includes step-by-step configuration examples, troubleshooting tips, and performance optimization techniques.

## 2. Understanding MikroTik WAN Failover Fundamentals

### What is WAN Failover?

WAN failover automatically switches network traffic from a failed primary internet connection to a backup connection. This process happens without manual intervention and maintains business continuity.

Key benefits include:

- Reduced network downtime
- Automatic recovery processes
- Improved service reliability
- Cost savings from avoided outages

### MikroTik Failover Methods Overview

MikroTik RouterOS supports three primary failover approaches:

#### Distance-Based Routing Failover

- **Method** : Uses route distance values to prioritize connections
- **Pros** : Simple setup, built-in gateway monitoring
- **Cons** : Limited customization options
- **Best for** : Small to medium networks with basic requirements

#### Netwatch-Based Monitoring

- **Method** : Uses scripts triggered by network monitoring
- **Pros** : Full customization, advanced logic
- **Cons** : Complex setup, requires scripting knowledge
- **Best for** : Enterprise networks with specific requirements

#### VRRP Hardware Redundancy

- **Method** : Multiple routers share virtual IP addresses
- **Pros** : Hardware-level redundancy, fast failover
- **Cons** : Requires multiple devices, complex configuration
- **Best for** : Critical infrastructure with high availability needs

### Prerequisites and Planning

Before starting configuration, verify these requirements:

#### Hardware Requirements

- MikroTik RouterBoard with RouterOS v6.45 or newer
- Minimum two WAN interfaces
- Sufficient RAM for monitoring scripts (128MB recommended)

#### Network Requirements

- Two or more ISP connections
- Different physical paths for redundancy
- Adequate bandwidth on backup connections

#### Planning Checklist

- Document current network topology
- Identify critical services requiring failover
- Define acceptable downtime windows
- Plan IP addressing schemes
- Choose monitoring targets

## 3. Method 1: Distance-Based WAN Failover Configuration

Distance-based failover uses route priorities to control traffic flow. Routes with lower distance values take precedence over higher values.

### Basic Dual WAN Setup

#### Step 1: Configure WAN Interfaces

Set up primary WAN interface:

```
/ip dhcp-client
add interface=ether1 disabled=no comment="Primary WAN"
```
Set up secondary WAN interface:

```
/ip dhcp-client
add interface=ether2 disabled=no comment="Secondary WAN"
```
For static IP configuration:

```
/ip address
add address=203.0.113.10/24 interface=ether1 comment="Primary WAN Static"
add address=198.51.100.20/24 interface=ether2 comment="Secondary WAN Static"
```
#### Step 2: Configure NAT Rules

Create masquerade rules for both WAN interfaces:

```
/ip firewall nat
add chain=srcnat out-interface=ether1 action=masquerade comment="Primary WAN NAT"
add chain=srcnat out-interface=ether2 action=masquerade comment="Secondary WAN NAT"
```
#### Step 3: Set Up Default Routes

Configure primary route with gateway monitoring:

```
/ip route
add dst-address=0.0.0.0/0 gateway=203.0.113.1 distance=1 check-gateway=ping comment="Primary Route"
```
Configure backup route with higher distance:

```
/ip route
add dst-address=0.0.0.0/0 gateway=198.51.100.1 distance=2 comment="Backup Route"
```
### Gateway Monitoring Configuration

The check-gateway parameter monitors gateway reachability. Configure monitoring settings:

#### Advanced Gateway Check Options

```
/ip route
add dst-address=0.0.0.0/0 gateway=203.0.113.1 distance=1 \
    check-gateway=ping timeout=3s comment="Primary with timeout"
```
#### Multiple Target Monitoring

Monitor specific hosts instead of gateways:

```
/ip route
add dst-address=8.8.8.8/32 gateway=203.0.113.1 scope=10 comment="Google DNS Primary"
add dst-address=8.8.4.4/32 gateway=198.51.100.1 scope=10 comment="Google DNS Secondary"
add dst-address=0.0.0.0/0 gateway=8.8.8.8 distance=1 check-gateway=ping scope=11 comment="Primary via DNS"
add dst-address=0.0.0.0/0 gateway=8.8.4.4 distance=2 scope=11 comment="Secondary via DNS"
```
### Complete Basic Configuration Example

Here is a complete working configuration for basic dual WAN failover:

```
# Configure interfaces
/interface ethernet
set [find default-name=ether1] name=wan1-primary
set [find default-name=ether2] name=wan2-secondary
set [find default-name=ether3] name=lan1
# Set up DHCP clients for dynamic IP
/ip dhcp-client
add interface=wan1-primary disabled=no comment="Primary ISP"
add interface=wan2-secondary disabled=no comment="Secondary ISP"
# Configure LAN interface
/ip address
add address=192.168.1.1/24 interface=lan1 comment="LAN Gateway"
# Enable DHCP server for LAN
/ip pool
add name=lan-pool ranges=192.168.1.100-192.168.1.200
/ip dhcp-server
add address-pool=lan-pool interface=lan1 name=lan-dhcp
/ip dhcp-server network
add address=192.168.1.0/24 gateway=192.168.1.1 dns-server=8.8.8.8,8.8.4.4
# Configure NAT
/ip firewall nat
add chain=srcnat out-interface=wan1-primary action=masquerade
add chain=srcnat out-interface=wan2-secondary action=masquerade
# Configure routes with failover
/ip route
add dst-address=0.0.0.0/0 gateway=[/ip dhcp-client get [find interface=wan1-primary] gateway] \
    distance=1 check-gateway=ping comment="Primary Route"
add dst-address=0.0.0.0/0 gateway=[/ip dhcp-client get [find interface=wan2-secondary] gateway] \
    distance=2 comment="Secondary Route"
```
## 4. Method 2: Advanced Failover with Netwatch

Netwatch provides advanced monitoring capabilities with custom script execution. This method offers maximum flexibility for complex failover scenarios.

### Netwatch Configuration for Proactive Monitoring

#### Basic Netwatch Setup

Create netwatch entries for both WAN connections:

```
/tool netwatch
add host=8.8.8.8 interval=10s timeout=2s up-script="" down-script="" comment="Monitor Primary"
add host=1.1.1.1 interval=10s timeout=2s up-script="" down-script="" comment="Monitor Secondary"
```
#### Advanced Monitoring Parameters

- **Interval** : Time between ping attempts (recommended: 10-30 seconds)
- **Timeout** : Maximum wait time for response (recommended: 2-5 seconds)
- **Startup-delay** : Delay before starting monitoring (useful during boot)

### Script-Based Failover Automation

#### Basic Failover Script Structure

Create global variables for route management:

```
/system script
add name=global-vars source={
    :global primaryGW "203.0.113.1"
    :global secondaryGW "198.51.100.1"
    :global primaryRoute ""
    :global secondaryRoute ""
    :global currentActiveWAN "primary"
}
```
#### Primary WAN Down Script

```
/system script
add name=primary-down source={
    :global currentActiveWAN
    :global secondaryGW
    
    :if ($currentActiveWAN = "primary") do={
        :log info "Primary WAN failed - switching to secondary"
        
        # Disable primary route
        /ip route set [find comment="Primary Route"] disabled=yes
        
        # Enable secondary route if disabled
        /ip route set [find comment="Secondary Route"] disabled=no
        
        # Update global variable
        :set currentActiveWAN "secondary"
        
        :log info "Failover to secondary WAN completed"
    }
}
```
#### Primary WAN Up Script

