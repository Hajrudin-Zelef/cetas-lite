---
id: collect-260926-mikrotik/mikrotik/how-to-configure-mikrotik-wan-failover-for-high-availability-complete-guide-for-network-en-4
title: "Configure interfaces"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: ["ethernet", "latency", "memory", "revenue", "throughput", "training"]
source: docs/RAG/lot-mikrotik/forum/routing-bgp-ospf/how-to-configure-mikrotik-wan-failover-for-high-availability-complete-guide-for-network-engineers-te.md
source_anchor: ""
source_lines: [768, 979]
sha256: 8da79900ffc728758445a7e1912bf32898b5a98f7fafd9df700d62b5c6a11362
---

# Configure interfaces

```
# Create separate chains for each WAN interface
/ip firewall filter
add chain=forward in-interface=wan1-primary action=jump jump-target=wan1-security
add chain=forward in-interface=wan2-secondary action=jump jump-target=wan2-security
# Configure WAN1 security rules
add chain=wan1-security protocol=tcp dst-port=22,80,443 action=accept
add chain=wan1-security connection-state=established,related action=accept
add chain=wan1-security action=drop
# Configure WAN2 security rules (may have different requirements)
add chain=wan2-security protocol=tcp dst-port=22,80,443 action=accept
add chain=wan2-security connection-state=established,related action=accept
add chain=wan2-security action=drop
```
#### Source NAT Consistency

```
# Ensure consistent NAT behavior
/ip firewall nat
add chain=srcnat out-interface=wan1-primary src-address=192.168.1.0/24 \
    action=masquerade comment="Primary WAN NAT"
add chain=srcnat out-interface=wan2-secondary src-address=192.168.1.0/24 \
    action=masquerade comment="Secondary WAN NAT"
# Log NAT events for security monitoring
add chain=srcnat out-interface=wan1-primary action=log log-prefix="WAN1-OUT"
add chain=srcnat out-interface=wan2-secondary action=log log-prefix="WAN2-OUT"
```
### Logging and Auditing

#### Security Event Correlation

```
# Enhanced security logging
/system logging
add topics=firewall,critical,error action=remote remote=192.168.1.100:514
add topics=account,info action=memory prefix="AUTH"
# Create security monitoring script
/system script
add name=security-monitor source={
    :local suspiciousEvents 0
    
    # Check for unusual connection patterns
    :local connCount [/ip firewall connection print count-only]
    :if ($connCount > 1000) do={
        :log warning "High connection count detected: $connCount"
        :set suspiciousEvents ($suspiciousEvents + 1)
    }
    
    # Monitor failed login attempts
    :local failedLogins [:len [/log find message~"login failed"]]
    :if ($failedLogins > 5) do={
        :log warning "Multiple failed login attempts detected"
        :set suspiciousEvents ($suspiciousEvents + 1)
    }
    
    # Alert if multiple suspicious events
    :if ($suspiciousEvents > 1) do={
        /tool e-mail send to=security@company.com \
            subject="Security Alert" \
            body="Multiple suspicious events detected on router"
    }
}
```
## 11. Testing and Validation

### Failover Testing Procedures

#### Systematic Testing Methodology

1. **Pre-test Preparation**  - Document current network state
  - Identify test applications and services
  - Set up monitoring tools
  - Notify users of testing period
2. **Primary Failure Testing**  - Disconnect primary WAN cable
  - Monitor failover timing
  - Test application connectivity
  - Verify logging events
3. **Restoration Testing**  - Reconnect primary WAN
  - Verify failback occurs
  - Check for connection disruptions
  - Validate routing table updates

#### Automated Testing Scripts

```
# Create comprehensive test script
/system script
add name=failover-test source={
    :log info "Starting automated failover test"
    
    # Record baseline metrics
    :local startTime [/system clock get time]
    :local initialRoutes [/ip route print count-only where active=yes]
    
    # Test primary connection
    :local primaryTest [/ping 8.8.8.8 count=3 interval=1s]
    :log info "Primary connection test: $primaryTest packets received"
    
    # Simulate failure (disable primary route temporarily)
    /ip route set [find comment="Primary Route"] disabled=yes
    :delay 10s
    
    # Test secondary connection  
    :local secondaryTest [/ping 8.8.8.8 count=3 interval=1s]
    :log info "Secondary connection test: $secondaryTest packets received"
    
    # Restore primary connection
    /ip route set [find comment="Primary Route"] disabled=no
    :delay 15s
    
    # Final verification
    :local finalRoutes [/ip route print count-only where active=yes]
    
    :if ($finalRoutes = $initialRoutes) do={
        :log info "Failover test PASSED - routing restored"
    } else={
        :log error "Failover test FAILED - routing inconsistent"
    }
    
    :local endTime [/system clock get time]
    :log info "Test completed in $([:totime ($endTime - $startTime)])"
}
```
### Performance Measurement During Failover

#### Latency and Throughput Testing

```
# Create performance monitoring script
/system script
add name=performance-test source={
    :log info "Performance test starting"
    
    # Test latency to multiple targets
    :local targets {"8.8.8.8";"1.1.1.1";"208.67.222.222"}
    :foreach target in=$targets do={
        :local avgLatency [/ping $target count=10 interval=1s]
        :log info "Average latency to $target: $avgLatency ms"
    }
    
    # Monitor interface statistics
    :foreach interface in=[/interface find type=ethernet] do={
        :local intName [/interface get $interface name]
        :local rxBytes [/interface get $interface rx-byte]
        :local txBytes [/interface get $interface tx-byte]
        :log info "Interface $intName - RX: $rxBytes bytes, TX: $txBytes bytes"
    }
}
```
## 12. Real-World Case Studies

### Small Business Implementation

#### Scenario: Retail Store Chain

**Requirements**:

- 25 locations with point-of-sale systems
- Primary fiber connections with cable backup
- Maximum 60-second failover requirement
- Budget constraints on hardware

**Solution Implemented**:

```
# Basic distance-based failover for retail locations
/interface ethernet
set [find default-name=ether1] name=wan-fiber
set [find default-name=ether2] name=wan-cable
set [find default-name=ether3] name=lan-pos
# Configure automatic IP assignment
/ip dhcp-client  
add interface=wan-fiber disabled=no comment="Store Fiber"
add interface=wan-cable disabled=no comment="Store Cable"
# Simple LAN setup for POS systems
/ip address
add address=10.0.1.1/24 interface=lan-pos
# Configure DHCP for POS devices
/ip pool
add name=pos-pool ranges=10.0.1.100-10.0.1.150
/ip dhcp-server
add address-pool=pos-pool interface=lan-pos name=pos-dhcp
# Basic failover routing
/ip route
add dst-address=0.0.0.0/0 gateway=[find interface=wan-fiber] distance=1 check-gateway=ping
add dst-address=0.0.0.0/0 gateway=[find interface=wan-cable] distance=2
# Simple NAT configuration
/ip firewall nat
add chain=srcnat out-interface=wan-fiber action=masquerade
add chain=srcnat out-interface=wan-cable action=masquerade
```
**Results**:

- Average failover time: 45 seconds
- 99.8% uptime achieved across all locations
- Reduced revenue loss from outages by 95%
- Simple management with minimal training required

### Enterprise Deployment

#### Scenario: Financial Services Company

**Requirements**:

- Headquarters with 500+ users
- Multiple data centers requiring connectivity
- Sub-10-second failover requirement
- Regulatory compliance for logging
- VPN connectivity to branch offices

**Solution Implemented**:

