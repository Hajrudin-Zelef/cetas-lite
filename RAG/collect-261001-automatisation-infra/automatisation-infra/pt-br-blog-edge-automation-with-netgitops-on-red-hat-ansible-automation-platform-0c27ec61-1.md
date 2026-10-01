---
id: collect-261001-automatisation-infra/automatisation-infra/pt-br-blog-edge-automation-with-netgitops-on-red-hat-ansible-automation-platform-0c27ec61-1
title: "pt-br-blog-edge-automation-with-netgitops-on-red-hat-ansible-automation-platform-0c27ec61"
domain: automatisation-infra
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-automatisation-infra/pt-br-blog-edge-automation-with-netgitops-on-red-hat-ansible-automation-platform-0c27ec61.md
source_anchor: ""
source_lines: [1, 83]
sha256: 9aac5bd278c6949eee7b64058a38a1ccb3bc208f8a9a9457cf51bdcc0a556fc0
---

# pt-br-blog-edge-automation-with-netgitops-on-red-hat-ansible-automation-platform-0c27ec61

Network edge automation challenges
As organizations grow and expand geographi cally, they start extending their IT infrastructure into the distributed and far edge layers through opening new branch offices.
Restaurants, retail stores, and other customer-centric businesses provide differentiated wireless access for their employees, contractors and customers to interconnect within their designated areas.
Configuring and managing multiple wireless settings via Red Hat Ansible Automation Platform simplifies the deployments at scale.
Network administrators can use GitOps practices to automate wireless infrastructure as a code (IaC).
This case covers a sample use case for a company that uses an SDN (software-defined network) controller with a large network infrastructure, including access points, switches, and firewalls/routers to provide connectivity for thousands of branches across multiple countries. We will show you step by step how to automate wireless network access point settings at scale through a SD-WAN controller, which will be Cisco Meraki for purposes of this demo.
Considerations about using a source of control. Why not scripts?
Typically an SDN controller has an API. Having access to an SDN API is an advantage, since we have a single point of contact with the controller, and we can operate the whole network using a single entity and we can interact with it programmatically, leaving behind human intervention.
Does it mean that I have to write a Python script or a program to use it?
Should I learn a programming language in order to use the API?
Short answer: No!
Although learning a programming language is a highly valuable skill, it’s not the fastest or more efficient way to consume the SDN API or solve this problem at enterprise level.
Going through the development path means you’ll end up with a short-term solution that has to be developed, maintained, tested, and documented, with users that must learn how to use it. And frequently, it’s not the best way to use the organization's resources.
Instead, the goal should be to operate the network infrastructure and treat it like code. Describe network elements, its properties and being able to use this description as the source of truth.
What does this mean for our use case (WiFi)?
It means that we describe any required wireless properties, such as credentials and encryption type, and where they should be available (the APs or campuses, device labels).
Therefore we now have two components to consider: First the desired state of the network, and second the logic needed to apply it. We are going to describe how to manage both.
Automating the WiFi network through a SD-WAN controller
We propose using Ansible Automation Platform to handle all the automation process and YAML files to describe (not program!) the tasks needed to achieve the desired state.
The scenario is described by the following diagram:
But how do I trigger the execution of the Ansible Playbooks when I change the desired state? How do I know who made a change? When was it launched?
This is when Ansible Automation Platform comes to help.
Let’s start from the beginning. We need to:
- Describe the wireless network state
- Put this state in a Git repository
- Create a playbook that reads this repository and apply it if needed
Network state
The state is described in YAML below. The specific file for these values will be network_vars.yml.
SSID:
  - name: MyCompany
    number: 2
    enabled: yes
    auth_mode: psk
    encryption_mode: wpa
    psk: Ansible123456!
  - name: MyCompany_employees
    number: 3
    enabled: yes
    auth_mode: psk
    encryption_mode: wpa
    psk: Ansible123456!
  - name: MyCompany_contractors
    number: 4
    enabled: yes
    auth_mode: psk
    encryption_mode: wpa
    psk: Ansible123456!
  - name: MyCompany_customers
    number: 5
    state: absent
    enabled: yes
    auth_mode: psk
    encryption_mode: wpa
    psk: Ansible123456!
WiFi SSIDs playbook:
How do I apply the settings? What will the playbook look like?
Actually, it is quite straightforward.
First, let’s create a playbook file called config_ssids.yml.
Since we’re interacting with the SDN API, in Ansible this means we’re working from localhost.
We also need to read the YAML with values in Ansible to populate a few variables.
So far the playbook looks like this:
---
- hosts: localhost
  vars_files:
    - network_vars.yml
Then we need to use them in the module (cisco.meraki.meraki_mr_ssid). 
    - name: Add SSIDs
      cisco.meraki.meraki_mr_ssid:
        auth_key: "{{ meraki_key }}"
        org_name: "{{ Org_Name }}"
        net_name: "{{ Net_Name }}"
        state: "{{ item.state | default('present') }}"
        name: "{{ item.name }}"
        number: "{{ item.number }}"
        encryption_mode: "{{ item.encryption_mode }}"
        auth_mode: "{{ item.auth_mode }}"
        psk: "{{ item.psk }}" 
      delegate_to: localhost
      loop: "{{ SSID }}"
The final playbook looks like this:
---
- hosts: localhost
  
