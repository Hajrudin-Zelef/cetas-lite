---
id: collect-261001-automatisation-infra/automatisation-infra/pt-br-blog-edge-automation-with-netgitops-on-red-hat-ansible-automation-platform-0c27ec61-2
title: "pt-br-blog-edge-automation-with-netgitops-on-red-hat-ansible-automation-platform-0c27ec61"
domain: automatisation-infra
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-automatisation-infra/pt-br-blog-edge-automation-with-netgitops-on-red-hat-ansible-automation-platform-0c27ec61.md
source_anchor: ""
source_lines: [84, 165]
sha256: ddbfc2eec080d1d8358385625ab9d608b5d2299d201de888d010ab33c3b3148f
---

# pt-br-blog-edge-automation-with-netgitops-on-red-hat-ansible-automation-platform-0c27ec61

  vars_files:
    - network_vars.yml
  tasks:
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
Configuring GitHub
Let’s configure GitHub to trigger events from certain Git actions, such as commits and pushes:
First, we initialize our local repository
git  init
Then, we create a GitHub repository (I guess you already have your GitHub account, right?)
We’ll tell our local repo to use this newly created GitHub repo to be the remote repository.
These steps are “suggested” when a new repository is created in GitHub. Anyway, the steps are:
- Remote (github.com):
- Locally, usually from your laptop:
ls -l
total 24
-rw-r--r--  1 pauloseguel  staff    86B Jan 14 16:14 README.md
-rw-r--r--  1 pauloseguel  staff   766B Jan 14 16:14 config_ssids.yml
-rw-------  1 pauloseguel  staff   625B Jan 14 16:14 network_vars.yml
git init
git add README.md network_vars.yml config_ssids.yml
git commit -m "Config SSIDs"
git branch -M main
git remote add origin git@github.com:pseguel-redhat/gitops_test.git
git push -u origin main
Now we need to actually configure how to trigger events from GitHub. The full instructions in Working with Webhooks, and the short version is:
- Configure webhook personal access token (PAT) in GitHub
  - Click on your profile icon, then “Settings”
  - Click on “Developer Settings”, “Personal Access Tokens”
  - “Generate New Token”
  - Add a good description in the “Note” field
  - Automation controller only needs repo scope access, with the exception of invites. Select them.
  - Copy the token before leaving the page, since you won’t be able to access it later.
- Configure from automation controller; the steps are also demonstrated in this video 
 Configuring GitHub PAT Credentials in Ansible Automation Platform, and go as follow: 
  - Create a GitHub personal access token credential type and use the token from above.
- Create a custom credential type for SDN controller:
  - In this case, our playbook is using a single variable (meraki_key ) to store Meraki Dashboard API Token. We’ll use the integrated vault in automation controller to store this value in a secure fashion.
  - Name: Meraki
  - Input configuration:
- In this case, our playbook is using a single variable (
fields:
  - id: MERAKI_KEY
    type: string
    label: Meraki Dashboard API key
    secret: true
- Injector configuration:
extra_vars:
  meraki_key: '{{ MERAKI_KEY }}'
- Create a Meraki credential:
  - Name: Meraki
  - Credential Type: Meraki
  - Meraki Dashboard API key: Token from Meraki Dashboard API
Next, Create a project in automation controller:
- Click on the projects menu item on the left navigation menu.
- Click the blue “Add” button.
- Fill out the following fields:
  - Name: NetGitOps project
  - Source Control Type: Git
  - Source Control URL: https://github.com/pseguel-redhat/gitops_test
    - It should be the one where you host your source of truth
| Important: Enable the option “Update Revision on job launch” | 
- Create a job template
  - Name: Config SSIDs
  - Credentials: Meraki
  - Project: NetGitOps
  - Playbook: config_ssids.yml
  - Webhook: Select the created GitHub PAT credential
  - These steps will create a job template with a webhook key and a webhook URL. They’ll be used in the following GitHub steps.
- Now you have to create a GitHub webhook to integrate with Ansible Automation Platform, detailed steps can be followed also here: 
 
