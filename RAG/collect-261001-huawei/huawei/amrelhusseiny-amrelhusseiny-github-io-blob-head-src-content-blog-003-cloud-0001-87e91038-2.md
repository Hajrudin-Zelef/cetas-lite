---
id: collect-261001-huawei/huawei/amrelhusseiny-amrelhusseiny-github-io-blob-head-src-content-blog-003-cloud-0001-87e91038-2
title: "Switch Configuration"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: []
keywords: []
source: docs/RAG/collect-261001-huawei/amrelhusseiny-amrelhusseiny-github-io-blob-head-src-content-blog-003-cloud-0001--87e91038.md
source_anchor: ""
source_lines: [138, 225]
sha256: 73ff2373a461cc86d7d2caacbed65a424560da0eff04430cd8a14d184526493b
---

# Switch Configuration

```bash
Playbook.yaml : will be adding a vlan (20) to the switch 
```yaml
# Playbook.yaml
- name: test
  hosts: leaf_1
  connection: ansible.netcommon.netconf
  gather_facts: no
  tasks:
  - name: add_VLAN
    community.network.ce_vlan:
      vlan_id: 120
      name: WEB_Vlan
      description: "hello"
```bash
To run the playbook on the inventory :
`$ ansible-playbook -i huawei_cloud_engine_inventory.yaml huawei_cloud_engine_playbook.yaml `
Output confirming change has been done as follows : 
PLAY [test] *********************************************************************************************************************************************** skipping: no hosts matched
PLAY RECAP ************************************************************************************************************************************************
(huawei_venv) [amroashram@centos_proxy huawei_cloud_engine]$ ansible-playbook -i huawei_cloud_engine_inventory.yaml huawei_cloud_engine_playbook.yaml
PLAY [test] ***********************************************************************************************************************************************
TASK [add_VLAN] ******************************************************************************************************************************************* changed: [leaf_1]
PLAY RECAP ************************************************************************************************************************************************ leaf_1 : ok=1 changed=1 unreachable=0 failed=0 skipped=0 rescued=0 ignored=0
1         UT:GE1/0/1(D)      GE1/0/2(D)      GE1/0/3(D)      GE1/0/4(D)
GE1/0/5(D)      GE1/0/6(D)      GE1/0/7(D)      GE1/0/8(D)
GE1/0/9(D)
120
1 common   enable  default   enable  disable FWD FWD FWD VLAN 0001
120 common   enable  default   enable  disable FWD FWD FWD hello
#Further Ansible capabilities 
To check further usage of the netcommon collection with Cloud Engine , you have the following modules available to do changes on the switches : 
- Handy command , if you do not know what a module does or how to use it in the Playbook , you can always use :
`$ ansible-doc community.network.ce_vlan`
- Modules : refer back to link (https://docs.ansible.com/ansible/latest/network/user_guide/platform_ce.html) for up to date modules .
community.network.ce_aaa_server
community.network.ce_aaa_server_host
community.network.ce_acl
community.network.ce_acl_advance
community.network.ce_bfd_global
community.network.ce_bfd_session
community.network.ce_bfd_view
community.network.ce_bgp
community.network.ce_bgp_af
community.network.ce_bgp_neighbor
community.network.ce_bgp_neighbor_af
community.network.ce_dldp
community.network.ce_dldp_interface
community.network.ce_eth_trunk
community.network.ce_evpn_bd_vni
community.network.ce_file_copy
community.network.ce_info_center_debug
community.network.ce_info_center_global
community.network.ce_info_center_log
community.network.ce_info_center_trap
community.network.ce_interface
community.network.ce_interface_ospf
community.network.ce_ip_interface
community.network.ce_lacp
community.network.ce_link_status
community.network.ce_lldp
community.network.ce_lldp_interface
community.network.ce_mlag_config
community.network.ce_netconf
community.network.ce_ntp
community.network.ce_ospf
community.network.ce_ospf_vrf
community.network.ce_reboot
community.network.ce_sflow
community.network.ce_snmp_community
community.network.ce_snmp_target_host
community.network.ce_snmp_user
community.network.ce_static_route
community.network.ce_static_route_bfd
community.network.ce_switchport
community.network.ce_vlan
community.network.ce_vrf
community.network.ce_vrf_af
community.network.ce_vrf_interface
community.network.ce_vrrp
community.network.ce_vxlan_tunnel
community.network.ce_vxlan_vap
# References : 
- EVE_NG Installation : https://www.eve-ng.net/index.php/documentation/installation/virtual-machine-install/
- Huawei Cloud Engine EVE Image : https://forum.huawei.com/enterprise/en/run-ce12800-ne40e-in-eve-ng/thread/653457-861
- Huawei Cloud Engine Netconf Configuration : https://support.huawei.com/enterprise/en/doc/EDOC1100198823/4b902677/example-for-establishing-communication-between-the-nms-and-a-device-using-netconf
- Cloud Engine  Netcommons Collection Doc : https://docs.ansible.com/ansible/latest/network/user_guide/platform_ce.html
- Python 3.10 veneer : https://docs.python.org/3/library/venv.html
