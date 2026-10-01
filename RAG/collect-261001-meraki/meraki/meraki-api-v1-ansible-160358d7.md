---
id: collect-261001-meraki/meraki/meraki-api-v1-ansible-160358d7
title: "meraki-api-v1-ansible-160358d7"
domain: meraki
role: reference
task: reference
actors: []
dates: []
keywords: ["parameters"]
source: docs/RAG/collect-261001-meraki/meraki-api-v1-ansible-160358d7.md
source_anchor: ""
source_lines: [1, 55]
sha256: 9a4b8e3913464642a3ed97bcf0681246432458c9813a102553f742bc02f60f25
---

# meraki-api-v1-ansible-160358d7

Ansible
Introduction
Ansible is an open-source automation tool sponsored by Red Hat, widely used across IT roles from system administrators to developers. This Ansible Collection is tailored to work with the Cisco Meraki Dashboard API, providing a powerful and simple Infrastructure as Code solution.  
More info: Meraki Ansible Collection - Reference Guide
Ansible Basics
An Ansible Collection is a package format that bundles various Ansible content types, such as playbooks, roles, modules, and plugins. 
A playbook serves as a blueprint for automation tasks. It outlines the steps that Ansible will execute on specified inventories or groups of hosts. A playbook comprises plays, which are ordered groupings of tasks. Each task is executed by an Ansible module that encapsulates the logic and parameters for that task. 
Playbooks can be saved, shared, or reused, which ensures consistent execution of tasks and codifies operational knowledge.
Prerequisites
Installation
- Install Ansible  or on a Mac  More info: Ansible Installation docs.
- Install Python Meraki SDK  or on a Mac  More info: Meraki Python library docs
- Install Ansible Collection
(Alternative install with Virtual Environment)
Create a virtual environment for Ansible and the Meraki API to run in.
More info: Python Virtual Environments 
How to Use
Once you have everything installed, obtain your Meraki API key, set your authentication up and start building your first Playbook.
API Authentication
The easiest way to provide access to your Meraki infrastracture is by setting your API key to an environment variable. Ansible will use the Meraki python library to make the API requests with the provided key details.
- Environment Variable  In your terminal, assign your Meraki API key to an environment variable.
Playbooks
Let's build our first playbook.
In this example, we will be gathering the identity of the administrator associated with this Meraki API key and then print the name and email. We will then return a list of the Meraki organization names this administrator can manage.
- Hosts File  Create a file called hosts and copy the following code into it.
- Develop the Playbook  Create a file called myplaybook.yml and copy the following code into it.
- Execute the Playbook  This command runs the playbook, targeting the hosts defined in the hosts file and performs the tasks specified inmyplaybook.yml .
- Results  The terminal should display our results as each task is completed. If you run into any issues, check your spacing, or refer to the Troubleshooting section in this guide. 
Success! Now you have a working Meraki Ansible collection and can begin configuring your Infrastructure as Code! 
Explore more example playbooks.
Advanced Options
Authentication
There are alternatives to providing your Meraki API key for use with the Ansible collection. 
- Credentials File
- Configuration File 
  - Create or use an exsiting ansible.cfg file defined in the next section.
  - Set meraki_api_key: "Your-API-Key" to your API key
  More info: Ansible Configurations 
Security Alert: This option could store API keys in plain text, which is not recommended.
Ansible Configuration
Ansible supports multiple sources for configuring its behavior, including configuration files, environment variables, command-line options, playbook keywords, and variables. Configuration files are sought in the following order:
- ANSIBLE_CONFIG environment variable, if set
- ansible.cfg in the current directory
- ~/.ansible.cfg in the home directory
- /etc/ansible/ansible.cfg
Ansible will use the first configuration file it finds from this list, ignoring the others.
Below is an example ansible.cfg file with configuration options specific to the Cisco Meraki Ansible collection:
Resources
Troubleshooting
Mac OS
If you encounter ERROR! A worker was found in a dead state or objc_initializeAfterForkError  errors, set the following environment variable:
Contributions and Feedback
For contributions, issues, or enhancements, please open an issue or create a PR.
Release Management
We adhere to Semantic Versioning. Version updates will align with Cisco Meraki product updates, REST API changes, and Python SDK releases.
