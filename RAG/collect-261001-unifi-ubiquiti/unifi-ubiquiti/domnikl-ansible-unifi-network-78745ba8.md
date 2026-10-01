---
id: collect-261001-unifi-ubiquiti/unifi-ubiquiti/domnikl-ansible-unifi-network-78745ba8
title: "Create collections directory if it doesn't exist"
domain: unifi-ubiquiti
role: reference
task: reference
actors: []
dates: []
keywords: ["license", "parameters"]
source: docs/RAG/collect-261001-unifi-ubiquiti/domnikl-ansible-unifi-network-78745ba8.md
source_anchor: ""
source_lines: [1, 163]
sha256: c2eb1661e83441c2365698bed5d77ed5545fcb93bc0bcd53de6d4cb3240d1c49
---

# Create collections directory if it doesn't exist

A comprehensive Ansible collection for managing UniFi Network resources through the UniFi Network Integration API.
This collection provides Ansible modules for automating UniFi Network management tasks, including DNS policy management. It communicates directly with the UniFi Network controller through its Integration API, enabling Infrastructure as Code approaches for network configurations.
- Ansible: >= 2.16.0
- Python: >= 3.9
- UniFi Network Controller: Access to a UniFi Network controller with Integration API enabled
- API Key: Valid API key for the UniFi Network controller
The collection requires the following Python packages:
requests >= 2.25.0
certifi
Clone the repository directly to your Ansible collections path:
# Create collections directory if it doesn't exist
mkdir -p ~/.ansible/collections/ansible_collections/domnikl
# Clone the repository
git clone https://github.com/domnikl/ansible-unifi-network.git ~/.ansible/collections/ansible_collections/domnikl/unifi_network
- Clone the repository:
git clone https://github.com/domnikl/ansible-unifi-network.git
cd ansible-unifi-network
- Build and install the collection:
ansible-galaxy collection build .
ansible-galaxy collection install domnikl-unifi_network-*.tar.gz
After installing the collection, ensure the Python dependencies are available:
# System-wide installation
pip install requests certifi
# Or in a virtual environment
python -m venv ansible-env
source ansible-env/bin/activate  # On Windows: ansible-env\Scripts\activate
pip install requests certifi
- UniFi Network Controller: Ensure you have access to a UniFi Network controller
- API Key: Generate an Integration API key in your UniFi Network controller:
  - Navigate to Settings → System → Integrations
  - Create a new Integration API key
  - Note the API key for use in your playbooks
Create a simple playbook to manage DNS policies:
---
- name: Manage UniFi Network DNS Policies
  hosts: localhost
  gather_facts: false
  vars:
    unifi_controller: "your-controller.example.com"
    unifi_api_key: "{{ vault_unifi_api_key }}"  # Store securely with ansible-vault
    
  tasks:
    - name: Create A record for server
      domnikl.unifi_network.dns_policy:
        host: "{{ unifi_controller }}"
        api_key: "{{ unifi_api_key }}"
        site_name: "default"
        type: "A_RECORD"
        domain: "server.internal.example.com"
        ipv4_address: "192.168.1.100"
        ttl_seconds: 3600
        enabled: true
        state: present
        
    - name: Create CNAME record
      domnikl.unifi_network.dns_policy:
        host: "{{ unifi_controller }}"
        api_key: "{{ unifi_api_key }}"
        site_name: "default" 
        type: "CNAME_RECORD"
        domain: "app.internal.example.com"
        ipv4_address: "server.internal.example.com"
        state: present
        
    - name: Remove DNS policy
      domnikl.unifi_network.dns_policy:
        host: "{{ unifi_controller }}"
        api_key: "{{ unifi_api_key }}"
        site_name: "default"
        domain: "old-server.internal.example.com" 
        state: absent
Always store API keys securely using ansible-vault:
# Create encrypted variable file
ansible-vault create group_vars/all/vault.yml
# Add your API key
vault_unifi_api_key: "your-actual-api-key-here"
Manages DNS policies on UniFi Network controllers.
Parameters:
| Parameter | Type | Required | Default | Description | 
|---|---|---|---|---|
| host | string | yes |  | UniFi Network controller hostname/IP | 
| api_key | string | yes |  | Integration API key | 
| site_name | string | no | "default" | Site name where policy should be managed | 
| verify_ssl | boolean | no | true | Whether to verify SSL certificates | 
| type | string | yes |  | DNS record type ( A_RECORD ,CNAME_RECORD ) | 
| domain | string | yes |  | Domain name for the DNS policy | 
| ipv4_address | string | no |  | IPv4 address or target domain | 
| ttl_seconds | integer | no | 3600 | TTL in seconds | 
| enabled | boolean | no | true | Whether the policy is enabled | 
| state | string | no | "present" | Whether policy should exist ( present ,absent ) | 
Examples:
# Create A record
- domnikl.unifi_network.dns_policy:
    host: controller.example.com
    api_key: "{{ unifi_api_key }}"
    type: A_RECORD
    domain: web.local
    ipv4_address: 192.168.1.10
    
# Create CNAME record
- domnikl.unifi_network.dns_policy:
    host: controller.example.com  
    api_key: "{{ unifi_api_key }}"
    type: CNAME_RECORD
    domain: www.local
    ipv4_address: web.local
    
# Delete DNS policy
- domnikl.unifi_network.dns_policy:
    host: controller.example.com
    api_key: "{{ unifi_api_key }}" 
    domain: old.local
    state: absent
- Clone the repository:
git clone https://github.com/domnikl/ansible-unifi-network.git
cd ansible-unifi-network
- Create and activate virtual environment:
python -m venv venv
source venv/bin/activate  # On Windows: venv\Scripts\activate
- Install development dependencies:
pip install -r requirements-dev.txt
- Run tests:
# Syntax validation
ansible-test sanity --docker
# Unit tests (if available)
ansible-test units --docker
Create a test playbook and ensure you have:
- Access to a UniFi Network controller
- Valid API key
- Network connectivity to the controller
ModuleNotFoundError: No module named 'requests'
# Install requests in your Python environment
pip install requests
SSL Certificate Verification Errors
# Disable SSL verification (not recommended for production)
- domnikl.unifi_network.dns_policy:
    verify_ssl: false
    # ... other parameters
API Authentication Errors
- Verify your API key is correct
- Ensure the API key has sufficient permissions
- Check that the Integration API is enabled on your controller
Connection Timeout
- Verify controller hostname/IP address
- Check network connectivity to the controller
- Ensure the controller is accessible on the network
Enable verbose output for troubleshooting:
ansible-playbook -vvv your-playbook.yml
Contributions are welcome! Please feel free to submit a Pull Request. For major changes, please open an issue first to discuss what you would like to change.
- Code Style: Follow Python PEP 8 and Ansible collection best practices
- Testing: Add tests for new modules and features
- Documentation: Update documentation for new features
- Commit Messages: Use clear, descriptive commit messages
- Fork the repository
- Create a feature branch (git checkout -b feature/amazing-feature )
- Commit your changes (git commit -m 'Add amazing feature' )
- Push to the branch (git push origin feature/amazing-feature )
- Open a Pull Request
This project is licensed under the GPL-3.0-or-later License - see the LICENSE file for details.
- Issues: GitHub Issues
- Documentation: Collection Documentation
- Repository: GitHub Repository
Author: Dominik Liebler (@domnikl)
