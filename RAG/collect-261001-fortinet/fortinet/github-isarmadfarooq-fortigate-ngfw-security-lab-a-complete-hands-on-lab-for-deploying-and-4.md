---
id: collect-261001-fortinet/fortinet/github-isarmadfarooq-fortigate-ngfw-security-lab-a-complete-hands-on-lab-for-deploying-and-4
title: "Required"
domain: fortinet
role: reference
task: reference
actors: ["Apple"]
dates: []
keywords: ["cybersecurity"]
source: docs/RAG/collect-261001-fortinet/github-isarmadfarooq-fortigate-ngfw-security-lab-a-complete-hands-on-lab-for-deploying-and-configuri.md
source_anchor: ""
source_lines: [519, 618]
sha256: 959bf5e959c9b85b8f69c30d96c5c8e9d4f0afa02e9c0a54ecda4feeb124987c
---

# Required

```
Try: File → Open → Browse to .ovf file (not .vmx)
If error persists, use VMware OVFTool:
  ovftool FortiGate-VM64.ovf output/FortiGate-VM64.vmx
```
```
fortigate-security-lab/
│
├── README.md                          ← This file
│
├── docs/
│   ├── FortiGate_Security_Lab_Report.docx   ← Full technical report (PDF-quality)
│   └── Lab_Assignment_Brief.pdf             ← Original assignment brief
│
├── screenshots/
│   ├── task1/
│   │   ├── 01_download_portal.png
│   │   ├── 02_extracted_files.png
│   │   ├── 03_vmnet_editor.png
│   │   ├── 04_privilege_escalation.png
│   │   ├── 05_vmnet_config.png
│   │   ├── 06_ovf_import.png
│   │   ├── 07_vm_name_path.png
│   │   ├── 08_import_progress.png
│   │   ├── 09_vm_settings.png
│   │   ├── 10_nic_allocation.png
│   │   ├── 11_cli_login.png
│   │   ├── 12_get_interface.png
│   │   ├── 13_set_mgmt_ip.png
│   │   ├── 14_show_interface.png
│   │   ├── 15_ping_mgmt.png
│   │   ├── 16_gui_login.png
│   │   ├── 17_setup_wizard.png
│   │   ├── 18_dashboard.png
│   │   ├── 19_interfaces_list.png
│   │   ├── 20_wan_config.png
│   │   ├── 21_lan_config.png
│   │   ├── 22_dhcp_enable.png
│   │   ├── 23_static_routes.png
│   │   ├── 24_default_route.png
│   │   └── 25_connectivity_test.png
│   │
│   ├── task2/
│   │   ├── 01_https_verify.png
│   │   ├── 02_admin_access.png
│   │   ├── 03_administrators_menu.png
│   │   ├── 04_admin_list.png
│   │   ├── 05_edit_admin.png
│   │   ├── 06_login_page.png
│   │   └── 07_login_success.png
│   │
│   ├── task3/
│   │   ├── 01_fw_policy_create.png
│   │   ├── 02_fw_policy_config.png
│   │   ├── 03_nat_enable.png
│   │   ├── 04_policy_saved.png
│   │   ├── 05_dos_policy_l3.png
│   │   ├── 06_dos_l3_thresholds.png
│   │   ├── 07_dos_policy_tcp_udp.png
│   │   ├── 08_dos_policy_icmp_sctp.png
│   │   ├── 09_dos_policy_saved.png
│   │   ├── 10_feature_visibility.png
│   │   ├── 11_webfilter_profile.png
│   │   ├── 12_url_filter.png
│   │   ├── 13_facebook_block_entry.png
│   │   ├── 14_webfilter_policy.png
│   │   ├── 15_security_profile_attach.png
│   │   └── 16_policy_order.png
│   │
│   └── task4/
│       ├── 01_facebook_blocked_browser.png
│       ├── 02_block_page.png
│       ├── 03_block_details.png
│       ├── 04_webfilter_log.png
│       ├── 05_hping3_command.png
│       ├── 06_dos_detection_log.png
│       └── 07_dos_event_detail.png
│
└── configs/
    └── fortigate_baseline.conf        ← Sample FortiGate CLI config export
```
| Resource | URL | 
|---|---|
| FortiOS Administration Guide | https://docs.fortinet.com/document/fortigate | 
| FortiGate VM on VMware Guide | https://docs.fortinet.com/document/fortigate-vm | 
| FortiGate DoS Policy Docs | https://docs.fortinet.com/document/fortigate/latest/administration-guide | 
| NIST SP 800-41 Rev. 1 — Firewall Guidelines | https://csrc.nist.gov/publications/detail/sp/800-41/rev-1/final | 
| NIST SP 800-94 — IDPS Guide | https://csrc.nist.gov/publications/detail/sp/800-94/final | 
| hping3 Manual | https://linux.die.net/man/8/hping3 | 
| VMware Workstation Documentation | https://docs.vmware.com/en/VMware-Workstation-Pro | 
| Fortinet Community Forums | https://community.fortinet.com | 

| Field | Detail | 
|---|---|
| **Name** | Sarmad Farooq | 
| **Student ID** | 25I-7722 | 
| **Program** | MS Cybersecurity | 
| **Course** | Advanced Network Security | 
| **Instructor** | Dr. Zafar Iqbal | 
| **Institution** | NUCES — FAST Islamabad |
