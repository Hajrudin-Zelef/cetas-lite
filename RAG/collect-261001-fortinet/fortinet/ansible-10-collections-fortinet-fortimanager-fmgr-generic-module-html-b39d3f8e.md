---
id: collect-261001-fortinet/fortinet/ansible-10-collections-fortinet-fortimanager-fmgr-generic-module-html-b39d3f8e
title: "ansible-10-collections-fortinet-fortimanager-fmgr-generic-module-html-b39d3f8e"
domain: fortinet
role: reference
task: reference
actors: []
dates: []
keywords: ["parameters"]
source: docs/RAG/collect-261001-fortinet/ansible-10-collections-fortinet-fortimanager-fmgr-generic-module-html-b39d3f8e.md
source_anchor: ""
source_lines: [1, 73]
sha256: c7fc5ad6b46bb0a56168bc621cccca9abc506dee1d0c530d835be12a074c4478
---

# ansible-10-collections-fortinet-fortimanager-fmgr-generic-module-html-b39d3f8e

fortinet.fortimanager.fmgr_generic module – Build and send generic FortiManager API request.
Note
This module is part of the fortinet.fortimanager collection (version 2.8.2).
You might already have this collection installed if you are using the ansible package.
It is not included in ansible-core.
To check whether it is installed, run ansible-galaxy collection list.
To install it, use: ansible-galaxy collection install fortinet.fortimanager.
To use it in a playbook, specify: fortinet.fortimanager.fmgr_generic.
New in fortinet.fortimanager 2.0.0
Synopsis
- This module is for generic fortimanager requests. it receives raw json-rpc data, and sends it to fortimanager, finally returns the response to users.
- This module also rely on fortimanager httpapi plugin as the transport.
- the payload doesn’t include session, the httpapi plugin will automatically fill the session later.
- the username and password is not managed by the module, but by the plugin.
Parameters
| Parameter | Comments | 
|---|---|
|  | The token to access FortiManager without using username and password. | 
|  | Enable/Disable logging for task Choices:  | 
|  | authenticate Ansible client with forticloud API access token | 
|  | the raw json-formatted payload to send to fortimanager | 
|  | the method of the json-rpc it must be in [get, add, set, update, delete, move, clone, exec] | 
|  | the parameter collection. | 
|  | the rc codes list with which the conditions to fail will be overriden | 
|  | the rc codes list with which the conditions to succeed will be overriden | 
|  | no description | 
|  | no description Default: 300 | 
Notes
Note
- two parameters schemes are supported, either in raw json format or in ansible recognnizable top-level parameters format.
- json is defined as string, user is response for make it json-formatted
- method and params should be specified by users if ‘json’ is not present
- if all three parameters are provided, the ‘json’ is preferred.
Examples
- name: Generic module
  hosts: fortimanagers
  connection: httpapi
  vars:
    adom: "root"
    ansible_httpapi_use_ssl: true
    ansible_httpapi_validate_certs: false
    ansible_httpapi_port: 443
  tasks:
    - name: Login a user
      fortinet.fortimanager.fmgr_generic:
        method: "exec"
        params:
          - url: "sys/login/user"
            data:
              - user: "APIUser"
                passwd: "Fortinet1!e"
    - name: Login another user
      fortinet.fortimanager.fmgr_generic:
        json: |
          {
           "method":"exec",
           "params":[
            {
                 "url":"sys/login/user",
                 "data":[
                    {
                       "user":"APIUser",
                       "passwd":"Fortinet1!"
                    }
                  ]
             }
            ]
          }
Return Values
Common return values are documented here, the following are the fields unique to this module:
| Key | Description | 
|---|---|
|  | full API response, includes status code and message Returned: always |
