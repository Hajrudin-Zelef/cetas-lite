---
id: collect-261001-fortinet/fortinet/momentumnn-fortigate-fortios7-4x-cis-hardening-5d5d74b0
title: "Audit all devices"
domain: fortinet
role: reference
task: reference
actors: []
dates: []
keywords: ["parameters"]
source: docs/RAG/collect-261001-fortinet/momentumnn-fortigate-fortios7-4x-cis-hardening-5d5d74b0.md
source_anchor: ""
source_lines: [1, 59]
sha256: d2575827566970ac2948435da46b4c389ce7db910aaa04879971b4c58ef741c1
---

# Audit all devices

Automated CIS Hardening solution for FortiGate firewalls using Ansible.
- ✅ Automated configuration backups
- ✅ Device information documentation
- ✅ Audit logging
pip3 install ansible fortiosapi
ansible-galaxy collection install -r requirements.yamlexport FORTIGATE_USER="admin"
export FORTIGATE_PASSWORD="YourPassword"
Edit hosts.yaml and update IP addresses for your FortiGate devices.
# Audit all devices
ansible-playbook playbooks/forti_audit.yaml
# Audit specific group
ansible-playbook playbooks/forti_audit.yaml --limit production_fortigates
# Audit single device
ansible-playbook playbooks/forti_audit.yaml --limit fw-prod-01# Harden all devices
ansible-playbook playbooks/forti_general_settings.yaml
# Harden specific group
ansible-playbook playbooks/forti_general_settings.yaml --limit production_fortigates
# Harden single device
ansible-playbook playbooks/forti_general_settings.yaml --limit fw-prod-01fortigate-backup/
├── ansible.cfg
├── backup_fortigate.yaml
├── hosts.yaml
├── requirements.yaml
├── group_vars/
│   ├── all.yaml
│   └── fortigates/
│       ├── vars.yaml
│       └── vault.yaml
├── host_vars/
│   ├── fw-prod-01.yaml
│   ├── fw-prod-02.yaml
│   ├── fw-branch-01.yaml
│   ├── fw-branch-02.yaml
│   └── fw-dmz-01.yaml
├── playbooks/
│   ├── backups/
│   ├── forti_add_trust_host.yaml
│   ├── forti_audit.yaml
│   ├── forti_backup.yaml
│   ├── forti_general_settings.yaml
│   ├── forti_logging.yaml
│   ├── forti_security_profiles.yaml
│   └── get.yaml
- backup_dir : Backup storage location
- retention_days : Default retention period
- backup_timestamp : Timestamp format
- Connection parameters (HTTPS, SSL, timeouts)
- Authentication configuration
- FortiGate-specific settings
- Device metadata (site, location, role)
- Custom retention periods
- Device-specific overrides
- Use Ansible Vault for credential storage
- Set appropriate file permissions
- Secure backup directory
- Use dedicated API user with read-only access
ansible fortigates -m fortinet.fortios.fortios_monitor_fact -a "selector=system_status"ansible-inventory --host fw-prod-01ansible-playbook forti_xxx.yaml -vvv
MIT
Ehsan Momeni Bashusqeh, Network Automation Engineer
