---
id: collect-261001-huawei/huawei/docs-ansible-2-10-collections-community-network-ce-netconf-module-0551d5c2
title: "docs-ansible-2-10-collections-community-network-ce-netconf-module-0551d5c2"
domain: huawei
role: reference
task: reference
actors: []
dates: []
keywords: ["license", "parameters"]
source: docs/RAG/collect-261001-huawei/docs-ansible-2-10-collections-community-network-ce-netconf-module-0551d5c2.md
source_anchor: ""
source_lines: [1, 86]
sha256: 406b3c1c928e7f87fc4a5b66108eb5bca8c7e178418122f14f8da1a4eed71b98
---

# docs-ansible-2-10-collections-community-network-ce-netconf-module-0551d5c2

community.network.ce_netconf – Run an arbitrary netconf command on HUAWEI CloudEngine switches.
Note
This plugin is part of the community.network collection (version 1.3.0).
To install it use: ansible-galaxy collection install community.network.
To use it in a playbook, specify: community.network.ce_netconf.
Synopsis
- Sends an arbitrary netconf command on HUAWEI CloudEngine switches.
Parameters
| Parameter | Choices/Defaults | Comments | 
|---|---|---|
| cfg_xml  string / required  |  | The config xml string. | 
| rpc  string / required  |  | The type of rpc. | 
Notes
Note
- This module requires the netconf system service be enabled on the remote device being managed.
- Recommended connection is netconf .
- This module also works with local connections for legacy playbooks.
Examples
- name: CloudEngine netconf test
  hosts: cloudengine
  connection: local
  gather_facts: no
  vars:
    cli:
      host: "{{ inventory_hostname }}"
      port: "{{ ansible_ssh_port }}"
      username: "{{ username }}"
      password: "{{ password }}"
      transport: cli
  tasks:
  - name: "Netconf get operation"
    community.network.ce_netconf:
      rpc: get
      cfg_xml: '<filter type="subtree">
                  <vlan xmlns="http://www.huawei.com/netconf/vrp" content-version="1.0" format-version="1.0">
                    <vlans>
                      <vlan>
                        <vlanId>10</vlanId>
                        <vlanif>
                          <ifName></ifName>
                          <cfgBand></cfgBand>
                          <dampTime></dampTime>
                        </vlanif>
                      </vlan>
                    </vlans>
                  </vlan>
                </filter>'
      provider: "{{ cli }}"
  - name: "Netconf edit-config operation"
    community.network.ce_netconf:
      rpc: edit-config
      cfg_xml: '<config>
                    <aaa xmlns="http://www.huawei.com/netconf/vrp" content-version="1.0" format-version="1.0">
                      <authenticationSchemes>
                        <authenticationScheme operation="create">
                          <authenSchemeName>default_wdz</authenSchemeName>
                          <firstAuthenMode>local</firstAuthenMode>
                          <secondAuthenMode>invalid</secondAuthenMode>
                        </authenticationScheme>
                      </authenticationSchemes>
                    </aaa>
                   </config>'
      provider: "{{ cli }}"
  - name: "Netconf execute-action operation"
    community.network.ce_netconf:
      rpc: execute-action
      cfg_xml: '<action>
                     <l2mc xmlns="http://www.huawei.com/netconf/vrp" content-version="1.0" format-version="1.0">
                       <l2McResetAllVlanStatis>
                         <addrFamily>ipv4unicast</addrFamily>
                       </l2McResetAllVlanStatis>
                     </l2mc>
                   </action>'
      provider: "{{ cli }}"
   Return Values
Common return values are documented here, the following are the fields unique to this module:
| Key | Returned | Description | 
|---|---|---|
| changed  boolean  | always | check to see if a change was made on the device Sample: True | 
| end_state  dictionary  | always | k/v pairs of aaa params after module execution Sample: {'result': ['ok']} | 
Authors
- wangdezhuang (@QijunPan)
    © 2012–2018 Michael DeHaan
© 2018–2019 Red Hat, Inc.
Licensed under the GNU General Public License version 3.
    https://docs.ansible.com/ansible/2.10/collections/community/network/ce_netconf_module.html
