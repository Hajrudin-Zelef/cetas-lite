---
id: collect-261001-fortinet/fortinet/project-hfortix-fortios-0-5-42-d2e55311-2
title: "Connect to FortiGate"
domain: fortinet
role: reference
task: reference
actors: ["Apple"]
dates: []
keywords: ["license", "parameters"]
source: docs/RAG/collect-261001-fortinet/project-hfortix-fortios-0-5-42-d2e55311.md
source_anchor: ""
source_lines: [166, 401]
sha256: 2bdc41822c817c4c3339c8ff6dacb73e1c61ea78ebdac1da6683802498be611a
---

# Connect to FortiGate

fgt.firewall.service_custom.create(
    name="custom-app",
    tcp_portrange="8080-8090",
    comment="My application"
)
# Schedules
fgt.firewall.schedule_recurring.create(
    name="business-hours",
    day=["monday", "tuesday", "wednesday", "thursday", "friday"],
    start="08:00",
    end="17:00"
)
# Traffic Shaping
fgt.firewall.traffic_shaper.create(
    name="critical-apps",
    guaranteed_bandwidth=50000,
    maximum_bandwidth=100000,
    bandwidth_unit="kbps"
)
# IP/MAC Binding
fgt.firewall.ipmacbinding_table.create(
    ip="10.0.1.100",
    mac="00:11:22:33:44:55",
    name="Server-01"
)
Available Wrappers:
- Service Management: service_custom ,service_category ,service_group
- Schedules: schedule_onetime ,schedule_recurring ,schedule_group
- Traffic Shaping: traffic_shaper ,shaper_per_ip
- IP/MAC Binding: ipmacbinding_table ,ipmacbinding_setting
- SSH/SSL Proxy: ssh_host_key ,ssh_local_ca ,ssh_local_key ,ssh_setting ,ssl_setting (⚠️ with API limitations)
- Firewall Policies: policy with 150+ parameters
Note: Some wrappers have FortiOS API limitations (e.g., SSH CA deletion requires CLI/GUI). See documentation for details.
⚡ Advanced Features
Async/Await Support:
import asyncio
async def main():
    async with FortiOS(host="...", token="...", mode="async") as fgt:
        # All methods support await
        addresses = await fgt.api.cmdb.firewall.address.list()
        # Concurrent operations
        addr, pol, svc = await asyncio.gather(
            fgt.api.cmdb.firewall.address.list(),
            fgt.api.cmdb.firewall.policy.list(),
            fgt.api.cmdb.firewall.service.custom.list()
        )
asyncio.run(main())
Error Handling:
from hfortix_core import (
    APIError,
    ResourceNotFoundError,
    DuplicateEntryError
)
try:
    fgt.api.cmdb.firewall.address.create(name="test", subnet="10.0.0.1/32")
except DuplicateEntryError:
    print("Address already exists")
except ResourceNotFoundError:
    print("Resource not found")
except APIError as e:
    print(f"API Error: {e.message} (code: {e.error_code})")
Read-Only Mode & Operation Tracking:
# Safe testing - block all write operations
fgt = FortiOS(host="...", token="...", read_only=True)
# Audit logging - track all API calls
fgt = FortiOS(host="...", token="...", track_operations=True)
operations = fgt.get_operations()
Performance Testing:
# Test your device and get optimal settings
results = fgt.api.utils.performance_test()
print(f"Recommended settings: {results['recommendations']}")
🔧 Enterprise Features
- Audit Logging: Built-in compliance logging with SIEM integration (SOC 2, HIPAA, PCI-DSS)
- Observability: Structured logging, distributed tracing with trace_id , user context tracking
- HTTP/2 Support: Connection multiplexing for better performance
- Automatic Retry: Handles transient failures (429, 500, 502, 503, 504) with exponential/linear/fibonacci backoff
- Circuit Breaker: Prevents cascade failures with automatic recovery
- Request Tracking: Correlation IDs for distributed tracing
- Validation Framework: 832 auto-generated validators
🔍 Debugging & Monitoring (v0.4.0)
Quick Debug Mode:
# Enable debug logging with simple boolean
fgt = FortiOS(host="...", token="...", debug=True)
Connection Pool Monitoring:
# Real-time connection statistics
stats = fgt.connection_stats
print(f"Active: {stats['active_requests']}/{stats['max_connections']}")
print(f"Total requests: {stats['total_requests']}")
print(f"Pool exhaustion: {stats['pool_exhaustion_count']}")
Request Inspection:
# Debug slow or failed requests
result = fgt.api.cmdb.firewall.address.list()
info = fgt.last_request
print(f"Endpoint: {info['endpoint']}")
print(f"Response time: {info['response_time_ms']}ms")
print(f"Status: {info['status_code']}")
Debug Session:
from hfortix_fortios import DebugSession
# Comprehensive session monitoring
with DebugSession(fgt) as session:
    # Make API calls
    fgt.api.cmdb.firewall.address.list()
    fgt.api.cmdb.firewall.policy.list()
    # Auto-prints summary on exit:
    # - Duration, total requests, success/failure counts
    # - Avg/min/max response times
    # - Connection pool deltas
Performance Profiling:
from hfortix_fortios import debug_timer
# Time individual operations
with debug_timer("Fetch all addresses") as timing:
    result = fgt.api.cmdb.firewall.address.list()
print(f"Took {timing['duration_ms']:.1f}ms")
Enhanced Logging:
from hfortix_fortios import configure_logging
# JSON logging for ELK/Splunk
configure_logging(
    level="INFO",
    format="json",
    include_trace=True,  # Add request_id to all logs
    output_file="/var/log/fortios.log"  # Log to file
)
# Text logging with colors for development
configure_logging(
    level="DEBUG",
    format="text",
    use_color=True
)
Type Hints & IDE Support:
# Full type hints for better autocomplete
from hfortix_fortios import FortiOS
from hfortix_core import APIResponse, ListResponse
fgt: FortiOS = FortiOS(host="...", token="...")
response: APIResponse = fgt.api.cmdb.firewall.address.get(name="test")
See docs/fortios/DEBUGGING.md for complete debugging guide.
- Type Safety: Full type hints with IDE autocomplete
- Structured Logging: Machine-readable JSON logs for ELK/Splunk/CloudWatch
Import Patterns
Recommended (New)
from hfortix_fortios import FortiOS
Legacy (Still Supported)
from hfortix import FortiOS
from hfortix.FortiOS import FortiOS
API Structure
# Configuration Management (CMDB)
fgt.api.cmdb.firewall.policy.*
fgt.api.cmdb.firewall.address.*
fgt.api.cmdb.system.interface.*
fgt.api.cmdb.router.static.*
fgt.api.cmdb.vpn.ipsec.*
# Monitoring
fgt.api.monitor.system.status()
fgt.api.monitor.firewall.session.*
fgt.api.monitor.system.resource.*
# Logging
fgt.api.log.disk.traffic.*
fgt.api.log.disk.event.*
fgt.api.log.disk.virus.*
# Convenience Wrappers
fgt.firewall.policy.*
fgt.firewall.service_custom.*
fgt.firewall.schedule_recurring.*
fgt.firewall.traffic_shaper.*
Documentation
Main Guides:
- Quick Start - Getting started guide
- Async Guide - Async/await patterns
- API Reference - Complete method reference
Convenience Wrappers:
- Overview Guide - All wrappers
- Service Wrappers - Service management
- Schedule Wrappers - Schedule management
- Shaper Wrappers - Traffic shaping
Advanced Features:
- Validation Guide - Using validators
- Filtering Guide - FortiOS filtering
- Performance Testing - Optimization
Full Documentation:
- Complete Changelog - Version history
- Main Repository - Complete docs
Requirements
- Python 3.10+
- FortiOS 7.0+ (tested with 7.6.5)
- hfortix-core >= 0.4.0-dev1
Development Status
Beta - All APIs are functional and tested against live FortiGate devices. The package remains in beta status until version 1.0.0 with comprehensive unit test coverage.
Current Test Coverage:
- 226 test files (145 CMDB, 81 Monitor)
- 75%+ pass rate
- ~50% of endpoints have dedicated tests
- All implementations validated against FortiOS 7.6.5
Examples
Firewall Policies
# Create policy
fgt.firewall.policy.create(
    name="Allow-Web",
    srcintf=["port1"],
    dstintf=["port2"],
    srcaddr=["all"],
    dstaddr=["web-servers"],
    action="accept",
    schedule="always",
    service=["HTTP", "HTTPS"],
    logtraffic="all"
)
# Check if exists
if fgt.firewall.policy.exists(policy_id=10):
    fgt.firewall.policy.update(policy_id=10, status="disable")
Address Management
# Create address
fgt.api.cmdb.firewall.address.create(
    name="web-server",
    subnet="192.168.1.100 255.255.255.255",
    comment="Production web server"
)
# Create address group
fgt.api.cmdb.firewall.addrgrp.create(
    name="internal-networks",
    member=["subnet1", "subnet2", "subnet3"],
    comment="All internal networks"
)
VPN Configuration
# Create IPsec Phase 1
fgt.api.cmdb.vpn.ipsec.phase1_interface.create(
    name="site-to-site",
    type="static",
    interface="wan1",
    ike_version=2,
    peertype="any",
    proposal="aes256-sha256",
    remote_gw="203.0.113.10"
)
License
Proprietary - See LICENSE file
Support
Author
