---
id: collect-261001-fortinet/fortinet/project-hfortix-fortios-0-5-42-d2e55311-1
title: "Connect to FortiGate"
domain: fortinet
role: reference
task: reference
actors: ["Apple"]
dates: ["2026-01", "2026-01-08"]
keywords: ["inference", "memory", "parameters"]
source: docs/RAG/collect-261001-fortinet/project-hfortix-fortios-0-5-42-d2e55311.md
source_anchor: ""
source_lines: [1, 165]
sha256: fb4088e5533e5db250013843e15df1554fe74fd73524df7f6d5cba8eb8a7a0ad
---

# Connect to FortiGate

HFortix FortiOS
Python SDK for FortiGate/FortiOS API - Complete, type-safe, production-ready.
⚠️ BETA STATUS - Version 0.5.34 (January 8, 2026)
Breaking Changes: See v0.5.33 and v0.5.32 for important return type changes in dict/object mode. Status: Production-ready but in beta until v1.0 with comprehensive unit tests. What's New: See below for all recent improvements, fixes, and new features!
Version: 0.5.34 Status: Beta (100% auto-generated, production-ready, optimized for performance)
🚀 What's New in v0.5.34 (January 2026)
Major Improvements and Breaking Changes (v0.5.32–v0.5.34)
- 
Enhanced type stub overloads for better IDE autocomplete 
  - All 1,000+ endpoints now have specific overloads for default mode (when response_mode is not specified)
  - Pylance now correctly infers types for queries by mkey without explicit type annotations
  - No longer need to annotate: rule: RuleResponse = ... — just userule = ...
  - Improved overload ordering for better type inference
- All 1,000+ endpoints now have specific overloads for default mode (when 
- 
Dict/Object mode query by name returns single item 
  - When querying by name/mkey with response_mode="dict" orresponse_mode="object" , now returns a single dict/object instead of a list
  - Example: group = fgt.api.cmdb.firewall.service.group.get(name="test") returns adict orFortiObject , not a list
  - Breaking Change: Tests and code expecting a list for single-item queries must be updated
- When querying by name/mkey with 
- 
Nested typed classes for table field children 
  - Table fields now have their own typed classes (e.g., GroupMemberObject )
  - Enables full IDE autocomplete for nested table attributes like .name ,.id , etc.
- Table fields now have their own typed classes (e.g., 
- 
Keyword argument support for mkey parameters 
  - Both get("name") andget(name="name") infer the correct return type
- Both 
- 
IDE autocomplete for table field members in object mode 
  - Table fields now return typed objects instead of generic FortiObject
  - Full attribute autocomplete for nested objects
- Table fields now return typed objects instead of generic 
- 
Universal table field normalization with schema awareness 
  - Handles custom mkeys: interface-name ,id ,index ,seq-num ,priority , etc.
  - All string values are automatically stripped of whitespace
- Handles custom mkeys: 
- 
Enhanced parameter documentation 
  - All POST/PUT method parameters now show field descriptions from the FortiOS schema in IDE tooltips
- 
Type annotations for FortiOS client attributes 
  - Improved type inference and autocomplete for fgt.api ,fgt.api.cmdb , etc.
- Improved type inference and autocomplete for 
- 
Other Notable Fixes and Improvements: 
  - Fixed stub generator comment truncation (no more broken comments in stubs)
  - Enhanced error messages for duplicate name/unique field conflicts
  - set() now accepts all field parameters (not just payload_dict)
  - Singleton endpoints now return a single object, not a list
  - FortiObject now supports both attribute and dictionary-style access
  - Filter parameter now accepts both string and list
  - Interactive help system for all API endpoints: endpoint.help()
  - Formatting utilities: to_json() ,to_csv() ,to_dict() , etc.
See the complete changelog for all details and previous versions.
Overview
Complete Python client for FortiOS 7.6.5 REST API with 100% endpoint coverage (1,219 endpoints), full type safety, and enterprise features. All code is auto-generated from FortiOS API schemas.
Installation
pip install hfortix-fortios
This automatically installs:
- hfortix-core - Core utilities and HTTP client
- hfortix-fortios-stubs - Type stubs for optimal IDE/type checker performance
For minimal installation (without stubs, smaller size):
pip install --no-deps hfortix-fortios
pip install hfortix-core  # Then install only runtime dependencies
For everything (includes future products):
pip install hfortix[all]
Quick Start
from hfortix_fortios import FortiOS
# Connect to FortiGate
fgt = FortiOS(
    host="192.168.1.99",
    token="your-api-token",
    verify=False
)
# Get system status
status = fgt.monitor.system.status()
print(f"Hostname: {status['hostname']}")
print(f"Version: {status['version']}")
# Manage firewall addresses
fgt.api.cmdb.firewall.address.create(
    name="web-server",
    subnet="192.168.1.100 255.255.255.255"
)
# 🎯 NEW! IDE autocomplete with Literal types (v0.5.4+)
fgt.api.cmdb.firewall.policy.create(
    name="allow-web",
    action="accept",      # 💡 IDE suggests: 'accept', 'deny', 'ipsec'
    status="enable",      # 💡 IDE suggests: 'enable', 'disable'
    logtraffic="all"      # 💡 IDE suggests: 'all', 'utm', 'disable'
)
API Coverage
FortiOS 7.6.5 - 100% Coverage (1,219 Endpoints):
- CMDB API: 886 endpoints - Full configuration management (firewall, system, VPN, routing, etc.)
- Monitor API: 295 endpoints - Real-time monitoring (sessions, stats, resources, etc.)
- Log API: 38 endpoints - Log queries (disk, memory, FortiAnalyzer, FortiCloud, search)
All endpoints are 100% auto-generated with:
- Complete .pyi type stub files
- Schema-based parameter validation
- Auto-generated basic tests
- Comprehensive error handling
Key Features
🎯 IDE Autocomplete with Literal Types (NEW in v0.5.4!)
15,000+ parameters with intelligent IDE autocomplete! Every enum parameter provides instant suggestions:
# ✨ Autocomplete for ALL enum fields
fgt.api.cmdb.firewall.policy.create(
    action='accept',      # 💡 IDE: 'accept', 'deny', 'ipsec'
    status='enable',      # 💡 IDE: 'enable', 'disable'
    nat='enable',         # 💡 IDE: 'enable', 'disable'
    logtraffic='all'      # 💡 IDE: 'all', 'utm', 'disable'
)
# 🛡️ Type safety catches errors at development time
fgt.api.cmdb.system.interface.create(
    mode='static',        # 💡 IDE: 'static', 'dhcp', 'pppoe'
    type='physical',      # 💡 IDE: 'physical', 'vlan', 'tunnel', ...
    role='lan'            # 💡 IDE: 'lan', 'wan', 'dmz', 'undefined'
)
Benefits: ⚡ Instant autocomplete • 🛡️ Type safety • 📚 Self-documenting • ✅ 100% backward compatible
🎯 Complete API Coverage
Access every FortiOS endpoint with clean, Pythonic syntax:
# CMDB (Configuration)
fgt.api.cmdb.firewall.policy.get()
fgt.api.cmdb.system.interface.get(name="port1")
fgt.api.cmdb.router.static.create(...)
# Monitor (Real-time data)
sessions = fgt.api.monitor.firewall.session.get()
resources = fgt.api.monitor.system.resource.usage.get()
# Log (Query logs)
vpn_logs = fgt.api.log.disk.event.vpn.get(rows=50)
traffic = fgt.api.log.memory.traffic.forward.get(rows=100)
🎨 Pretty Printing with FortiObject (NEW in v0.5.19!)
Clean, readable output for FortiOS data using response_mode="object":
# Enable object mode for pretty methods
fgt = FortiOS(
    host="192.168.1.99",
    token="your-token",
    response_mode="object"  # Returns FortiObject instead of dict
)
# Get policies and print cleanly
policies = fgt.api.cmdb.firewall.policy.get()
for policy in policies:
    print(f"\nPolicy {policy.policyid}: {policy.name}")
    print(f"  {policy.join('srcintf')} → {policy.join('dstintf')}")
    print(f"  {policy.join('srcaddr')} → {policy.join('dstaddr')}")
    print(f"  Service: {policy.join('service')} [{policy.action.upper()}]")
# Output:
# Policy 11: allow-web
#   port3 → port4
#   login.windows.net → gmail.com
#   Service: SAMBA [DENY]
FortiObject Methods:
- obj.join('field') - Join list values into comma-separated string
- obj.join('field', ' | ') - Custom separator
- obj.pretty('field') - Alias for join() with default separator
- Auto-flattens member_table fields: ['port1'] instead of[{'name': 'port1'}]
Benefits:
- 📊 Clean console output
- 🎯 No manual list comprehension needed
- ✨ Works with all FortiOS list fields
- 🔄 Original data always accessible via .to_dict()
🎨 Direct API Access
All 1,219 endpoints are accessed directly - no wrappers needed:
# Service Management
