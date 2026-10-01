---
id: collect-261001-fortinet/fortinet/t5-fortigate-technical-tip-application-of-ansible-on-fortigate-ta-p-258346-e4a6afb8
title: "t5-fortigate-technical-tip-application-of-ansible-on-fortigate-ta-p-258346-e4a6afb8"
domain: fortinet
role: reference
task: reference
actors: ["Apple"]
dates: []
keywords: []
source: docs/RAG/collect-261001-fortinet/t5-fortigate-technical-tip-application-of-ansible-on-fortigate-ta-p-258346-e4a6afb8.md
source_anchor: ""
source_lines: [1, 22]
sha256: e0b374e9c517760da154ddb6ca88c1c740b6f54f748d2beaa491c248e65b9af5
---

# t5-fortigate-technical-tip-application-of-ansible-on-fortigate-ta-p-258346-e4a6afb8

Technical Tip: Application of Ansible on FortiGate
Description
This article describes applications of Ansible on FortiGate.
FortiOS support two ways of authentication. Using user (ansible_user) and password (ansible_password), or an access token (fortios_access_token).
Scope
FortiGate, Ansible.
Solution
Ansible can be applied in multiple FortiGates to automate the provisioning, configuration, and management processes.
Follow each of the steps below to install and apply Ansible on FortiGate for multiple applications:
- Install Ansible on Linux.
Execute Ansible installation with the following terminal commands:
apt-get update
apt-get -y install ansible
Note:
Ensure the current user is the root user when executing the command. Otherwise, add the Linux user as a sudoer.
-  Download FortiGate modules with Ansible by using the command below: ansible-galaxy collection install fortinet.fortios
-  Create a variable text file to provide Ansible with FortiGate device information. nano <file name> or vim <file name> --- > Linux command to create a file After the file is created, fill in the information below: ### FortiGate Host### [fortigates] fgt ansible_host=<FGT IP> ansible_user="<FGT admin username>" ansible_password="<FGT username password>" fortigate ansible_host=<FGT IP> fortios_access_token=<access token> ### FortiGate OS Modules ### [fortigates:vars] ansible_network_os=fortinet.fortios.fortios ### Specify API connection with HTTP (default SSH)### [all:vars] ansible_connection=httpapi ansible_httpapi_validate_certs=no ansible_httpapi_use_ssl=yes Note: If using Linux nano, press Ctrl + X to save. Upon doing so, nano will show a directory path. Press Enter to save the file in the directory path. If using Linux vim, press the Esc key, type :w and hit the Enter key to save the file.
-  Create a YAML file and prepare the script with the following format: hosts: <object created in steps 3> tasks: - name: <tasks name> <Ansible Module> <Configuration on FortiGate> <Objects in the Configuration> The example below demonstrates a Basic Configuration on FortiGate with YAML: hosts: fortigates tasks: - name: Change hostname fortinet.fortios.fortios_system_global: system_global: hostname: “FortiGate_Lab” - name: Create Address fortinet.fortios.fortios_firewall_address: state: “present” firewall_address: name: test_123 subnet: 10.1.1.1 255.255.255.0 Playbook example below uses access token only, vars can also be placed within the playbook or removed if already present in the hosts file, this playbook example considers that the vars were not present on the hosts file: - hosts: fortigates connection: httpapi collections: - fortinet.fortios vars: vdom: "root" ansible_httpapi_use_ssl: true ansible_httpapi_validate_certs: false ansible_httpapi_port: 443 tasks: - name: Configure global attribute fortios_system_global: vdom: "{{ vdom }}" access_token: "{{ fortios_access_token }}" system_global: hostname: 'FortiGateHostName' The next example demonstrates another Basic Configuration on FortiGate with YAML related to tasks only getting invoked by the Playbook. In this case each YAML should contain inside the FortiGate ansible user and password, or access token. tasks: - include_tasks: file1_get_policy.yml - include_tasks: file2_get_address.yml Note: YAML is space sensitive. Each layer is differentiated and separated with spaces. DO NOT use Tab on the keyboard to create the space, as this will cause syntax errors. See the Ansible command for FortiOS here: Fortinet.Fortios The Ansible command for FortiOS MUST be included in the YAML configuration with 'required' highlighted. For example:
-  After the YAML file is created, use the Ansible command below to run the file:
ansible-playbook -i <variable text file name> < YAML file name>
The results should look similar to the following example:
Related Ansible documents:
