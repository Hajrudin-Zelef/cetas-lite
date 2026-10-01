---
id: collect-261001-cisco/cisco/github-emomeni-cisco-isr4k-hardening-ansible-playbook-ansible-playbook-that-applies-cis-be
title: "SSH hardening only"
domain: cisco
role: reference
task: reference
actors: ["CISA"]
dates: []
keywords: ["benchmark"]
source: docs/RAG/collect-261001-cisco/github-emomeni-cisco-isr4k-hardening-ansible-playbook-ansible-playbook-that-applies-cis-benchmark-v4.md
source_anchor: ""
source_lines: [1, 158]
sha256: ad362b28cc3eb1f0ec9029500d1e1ecc2cff8dcfb90b3d38b0eab84e23eb8545
---

# SSH hardening only

Ansible playbook that applies CIS Benchmark (v4.x) and NSA/CISA hardening controls to Cisco ISR 4000 series routers running IOS XE.

| Section | Controls | Tag | 
|---|---|---|
| Service hardening | Disable BOOTP, PAD, finger, CDP, LLDP, HTTP, small servers, source routing | `services` | 
| Passwords | Enable secret (scrypt), min-length 14, remove plaintext password | `passwords` | 
| AAA / TACACS+ | AAA new-model, TACACS+ group with local fallback, login blocking | `aaa` | 
| SSH | SSHv2 only, 4096-bit RSA keys, timeout, retries, VTY ACL | `ssh` | 
| Interfaces | Disable proxy-arp, redirects, unreachables, directed-broadcast on all interfaces | `interfaces` | 
| SNMP | Remove v1/v2c communities, configure SNMPv3 (SHA + AES128), restrict with ACL | `snmp` | 
| Banners | Login, MOTD, and EXEC banners | `banners` | 
| NTP | Authenticated NTP with trusted keys | `ntp` | 
| Logging | Remote syslog, buffered logging, source-interface, disable console logging | `logging` | 
| CoPP | Rate-limit ICMP, SSH, SNMP, routing protocols, drop everything else | `copp` | 

- **Ansible** >= 2.14
- **Python** >= 3.10
- **Collections** (install before first run):

`ansible-galaxy collection install cisco.ios ansible.netcommon`
- SSH reachability from the Ansible control node to all target routers
- A local user account on each router (fallback for TACACS+ outages)
- A pre-generated scrypt (`type 9` ) hash for the enable secret

On any IOS XE device:

```
Router(config)# enable algorithm-type scrypt secret YourStrongPasswordHere
Router# show running-config | include enable secret
```
Copy the full `$9$...` hash into `group_vars/all.yml` as the `enable_secret_hash` value.

```
cisco-isr4k-hardening/
  inventory/
    hosts.yml                  # Inventory with connection vars
  group_vars/
    all.yml                    # Variables (vault-encrypt this)
  cisco_isr4k_hardening.yml    # Main playbook
  README.md                    # This file
```
Open `inventory/hosts.yml` and replace the example hostnames and IPs with your routers:

```
isr4k-core-01:
  ansible_host: 10.0.100.1
```
Edit `group_vars/all.yml`. Replace every `CHANGE_ME` placeholder with real values:

- `vault_ansible_user` /`vault_ansible_password` - SSH credentials
- `vault_enable_secret` - enable mode password
- `enable_secret_hash` - scrypt hash (see above)
- `tacacs_server_ip` /`tacacs_key` - your TACACS+ server
- `snmp_auth_pass` /`snmp_priv_pass` - SNMPv3 credentials
- `ntp_key_values` - NTP authentication keys
- `mgmt_acl_permit_subnet` - your management network
- `syslog_server_ip` - your syslog collector
- `management_interface` - the interface used for management traffic

`ansible-vault encrypt group_vars/all.yml`
You'll set a vault password. Remember it - you'll need it every time you run the playbook.

```
ansible-playbook -i inventory/hosts.yml cisco_isr4k_hardening.yml \
  --ask-vault-pass --check --diff
```
```
ansible-playbook -i inventory/hosts.yml cisco_isr4k_hardening.yml \
  --ask-vault-pass
```
```
# SSH hardening only
ansible-playbook -i inventory/hosts.yml cisco_isr4k_hardening.yml \
  --ask-vault-pass --tags ssh
# AAA and TACACS+ only
ansible-playbook -i inventory/hosts.yml cisco_isr4k_hardening.yml \
  --ask-vault-pass --tags aaa
# CoPP only
ansible-playbook -i inventory/hosts.yml cisco_isr4k_hardening.yml \
  --ask-vault-pass --tags copp
# Multiple tags
ansible-playbook -i inventory/hosts.yml cisco_isr4k_hardening.yml \
  --ask-vault-pass --tags "ssh,aaa,logging"
```
```
ansible-playbook -i inventory/hosts.yml cisco_isr4k_hardening.yml \
  --ask-vault-pass --limit isr4k-core-01
```
```
echo 'your-vault-password' > .vault_pass
chmod 600 .vault_pass
ansible-playbook -i inventory/hosts.yml cisco_isr4k_hardening.yml \
  --vault-password-file .vault_pass
```
Add `.vault_pass` to `.gitignore`. Don't commit it.

| Tag | What it covers | 
|---|---|
| `services` | BOOTP, PAD, finger, CDP, LLDP, HTTP, small servers | 
| `passwords` | Enable secret, min-length, remove plaintext | 
| `aaa` | AAA new-model, TACACS+, login blocking | 
| `tacacs` | TACACS+ server configuration only | 
| `ssh` | SSHv2, keys, timeouts, VTY hardening | 
| `vty` | VTY line configuration only | 
| `console` | Console line configuration only | 
| `interfaces` | Proxy-arp, redirects, unreachables, directed-broadcast | 
| `snmp` | Remove v1/v2c, configure SNMPv3 | 
| `banners` | Login, MOTD, EXEC banners | 
| `ntp` | NTP servers and authentication | 
| `logging` | Syslog, buffer, timestamps | 
| `copp` | Control Plane Policing class-maps, policy-map, ACLs | 
| `cis` | All CIS Benchmark controls | 
| `nsa` | All NSA/CISA controls | 
| `verify` | Post-run verification checks only | 
| `acl` | All ACL-related tasks | 

Every task is safe to re-run. The `cisco.ios.ios_config` module compares desired state against the running config and only pushes changes when needed.

The RSA key generation task (`CIS 4.1.4`) uses `changed_when` logic to detect whether keys already exist at the requested bit length.

The SNMPv3 user task uses `ios_command` because IOS XE doesn't store `snmp-server user` in the running config. It will re-apply on each run. This is the standard approach - the router accepts the duplicate command without error.

The playbook runs four verification checks automatically:

1. SSHv2 is active
2. AAA new-model is present in running config
3. CoPP policy is applied to the control plane
4. NTP status (displayed for manual review)

Run verification checks alone:

```
ansible-playbook -i inventory/hosts.yml cisco_isr4k_hardening.yml \
  --ask-vault-pass --tags verify
```
All CoPP rates are in `group_vars/all.yml`. The defaults are conservative starting points:

| Class | Rate (bps) | Burst | 
|---|---|---|
| ICMP | 8000 | 1500 | 
| SSH | 64000 | 8000 | 
| SNMP | 32000 | 4000 | 
| Routing | 128000 | 16000 | 
| Default | 8000 | 1000 | 

Tune based on your traffic patterns. Monitor with `show policy-map control-plane` after deployment.

This playbook doesn't configure routing protocol authentication (OSPF MD5/SHA, BGP TCP-AO, EIGRP key chains) because those are topology-specific. Add them as separate tasks or a dedicated role.

This playbook targets IOS XE specifically. For IOS (classic), NX-OS, or IOS XR, you'll need platform-specific modules and adjusted commands.

**"Target does not appear to be a Cisco IOS XE device"** - The pre-check assertion failed. Verify the device model and image name in the error message. If it's a valid ISR 4000 but the string matching is off, adjust the assert conditions in the pre-tasks.

**SSH connection timeout** - Check that SSH is already enabled on the target. This playbook hardens SSH - it doesn't enable it from scratch. You need basic SSH access before running the playbook.

**TACACS+ lockout** - If TACACS+ is unreachable after applying AAA, the `local` fallback in the authentication line lets you log in with local credentials. Always maintain a working local account.

**SNMPv3 user already exists** - IOS XE may show a warning. The command is still accepted. This is expected behaviour.
