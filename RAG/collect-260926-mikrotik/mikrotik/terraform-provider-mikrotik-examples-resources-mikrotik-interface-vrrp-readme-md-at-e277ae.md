---
id: collect-260926-mikrotik/mikrotik/terraform-provider-mikrotik-examples-resources-mikrotik-interface-vrrp-readme-md-at-e277ae
title: "Edit terraform.tfvars with your passwords"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: ["preemption"]
source: docs/RAG/lot-mikrotik/forum/dhcp-dns/terraform-provider-mikrotik-examples-resources-mikrotik-interface-vrrp-readme-md-at-e277ae1131cde9bb.md
source_anchor: ""
source_lines: [1, 75]
sha256: 0e36bd0103984d355fb6bfdea9d7211d80023e729eb7e924ed4b6ce5d8ae4f9c
---

# Edit terraform.tfvars with your passwords

This example demonstrates how to configure VRRP (Virtual Router Redundancy Protocol) for high availability using two MikroTik routers.

- **Master Router** : 192.168.1.10 (priority 254)
- **Backup Router** : 192.168.1.11 (priority 100)
- **Virtual IP** : 192.168.1.1 (gateway for clients)
- **VRID** : 10

- ✅ VRRP interface configuration
- ✅ Master/Backup router setup
- ✅ Authentication (simple password)
- ✅ Virtual IP assignment
- ✅ State change scripts (optional)
- ✅ Multiple VRRP groups for load balancing

- 2x MikroTik RouterOS 7.x devices
- Network connectivity between routers
- Terraform 1.0+ installed

1. 
**Configure variables** :```
cp terraform.tfvars.example terraform.tfvars
# Edit terraform.tfvars with your passwords
```
2. 
**Initialize Terraform** :terraform init
3. 
**Review plan** :terraform plan
4. 
**Apply configuration** :terraform apply

```
/interface vrrp print detail
/interface vrrp monitor [find name=vrrp-gateway]
```
```
name="vrrp-gateway" running=yes master-router=<local>
```
```
name="vrrp-gateway" running=yes master-router=192.168.1.10
```
1. 
Disable master VRRP interface: `/interface vrrp set vrrp-gateway disabled=yes`
2. 
Check backup becomes master: /interface vrrp monitor vrrp-gateway
3. 
Re-enable master (with preemption): `/interface vrrp set vrrp-gateway disabled=no`

This example also shows how to use multiple VRRP groups for load balancing:

- **VRRP 10** : Router A is master, Router B is backup
- **VRRP 20** : Router B is master, Router A is backup

This distributes traffic across both routers while maintaining redundancy.

- **Same VRID** : Both routers must use the same VRID (10 in this example)
- **Authentication** : Use strong passwords in production
- **Preemption** : Enabled by default - higher priority router will become master
- **Firewall** : Don't block VRRP (IP protocol 112)
- **Timing** : Default interval is 1 second

- Check network connectivity
- Verify VRRP packets are not blocked by firewall
- Check for network loops

- Verify priority settings (higher = master)
- Check authentication passwords match
- Ensure physical interface is up
- Verify VRID matches on both routers

`/log print where topics~"vrrp"`
1. ✅ Always use authentication in production
2. ✅ Use strong passwords (min 16 characters)
3. ✅ Consider VRRPv3 with IPsec for encryption
4. ✅ Monitor VRRP state changes
5. ✅ Use scripts for alerting on state transitions
