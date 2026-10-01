---
id: collect-261001-automatisation-infra/automatisation-infra/jmvigueras-ansible-foobar-dac126da
title: "jmvigueras-ansible-foobar-dac126da"
domain: automatisation-infra
role: reference
task: reference
actors: ["Apple"]
dates: []
keywords: ["license"]
source: docs/RAG/collect-261001-automatisation-infra/jmvigueras-ansible-foobar-dac126da.md
source_anchor: ""
source_lines: [1, 83]
sha256: 19cbbaa1e3a6f984a5d394135b94ebe158d51045e3bc6e8b037acee6155c70b9
---

# jmvigueras-ansible-foobar-dac126da

A simple Ansible automation repository for configuring Virtual IPs (VIPs) and firewall policies on FortiGate devices using the FortiOS collection.
This repository provides an easy-to-use Ansible playbook that automates the configuration of:
- Virtual IP (VIP) objects for port forwarding/NAT
- Firewall policies to allow inbound traffic to the VIPs
- Support for multiple applications (APP1 and APP2)
.
├── 00-playbook.yml          # Main Ansible playbook
├── hosts                    # Inventory file with FortiGate connection details
├── roles/
│   └── vips/
│       ├── tasks/
│       │   └── main.yml     # VIP and policy configuration tasks
│       └── vars/
│           └── main.yml     # Variables for VIP configuration
├── .gitignore
└── LICENSE
- Ansible installed on your control machine
- FortiOS Collection for Ansible:
ansible-galaxy collection install fortinet.fortios
- FortiGate device accessible via HTTPS API
- API access token for FortiGate authentication
Edit the hosts file with your FortiGate details:
[fortigate]
fortigate-1
[fortigate:vars]
ansible_host="YOUR_FORTIGATE_IP"
ansible_network_os="fortinet.fortios.fortios"
ansible_httpapi_use_ssl="yes"
ansible_httpapi_validate_certs="no"
ansible_httpapi_port="8443"
vdom="root"
fortios_access_token="YOUR_API_TOKEN"
fgt_ext_ip="YOUR_EXTERNAL_IP"
mapped_ip="YOUR_INTERNAL_IP"
Key Variables:
- ansible_host : FortiGate management IP address
- fortios_access_token : API access token for authentication
- fgt_ext_ip : External IP address for VIP configuration
- mapped_ip : Internal IP address to map traffic to
Modify roles/vips/vars/main.yml to suit your requirements:
app_1_vip_name: "vip-app1-31000"    # Name for first VIP
app_2_vip_name: "vip-app2-31001"    # Name for second VIP
fgt_src_intf: "port1"               # Source interface (usually WAN)
fgt_dst_intf: "port2"               # Destination interface (usually LAN)
app_1_port: "31000"                 # Port for first application
app_2_port: "31001"                 # Port for second application
Execute the main playbook to configure VIPs and policies:
ansible-playbook -i hosts 00-playbook.yml
- Creates VIP Objects: Configures static NAT VIPs for each application
- Creates Firewall Policies: Allows inbound traffic from any source to the VIPs
- Enables Port Forwarding: Maps external ports to internal application ports
- Configures Logging: Enables traffic logging for monitoring
The playbook will create:
- VIP object vip-app1-31000 mapping external port 31000 to internal IP
- VIP object vip-app2-31001 mapping external port 31001 to internal IP
- Firewall policies (ID 100, 101) allowing traffic to these VIPs
- Login to FortiGate web interface
- Go to System > Administrators
- Create or edit an administrator account
- Enable REST API Access
- Generate or copy the API Key
- Use this key as fortios_access_token in your inventory
To add additional VIPs:
- Define new variables in roles/vips/vars/main.yml
- Add corresponding tasks in roles/vips/tasks/main.yml
- Increment policy IDs to avoid conflicts
Edit the firewall policy sections in roles/vips/tasks/main.yml to:
- Restrict source addresses
- Limit specific services instead of "ALL"
- Adjust logging settings
- Modify scheduling
- Authentication Failed: Verify API token and FortiGate accessibility
- Port Conflicts: Ensure policy IDs don't conflict with existing policies
- Interface Names: Confirm interface names match your FortiGate configuration
- VDOM Access: Ensure the API token has access to the specified VDOM
Run with verbose output for troubleshooting:
ansible-playbook -i hosts 00-playbook.yml -vvv
- Store API tokens securely (consider using Ansible Vault)
- Use least-privilege access for API tokens
- Review firewall policies for security compliance
- Regularly rotate API access tokens
This project is licensed under the terms specified in the LICENSE file.
Feel free to submit issues, fork the repository, and create pull requests for improvements.
