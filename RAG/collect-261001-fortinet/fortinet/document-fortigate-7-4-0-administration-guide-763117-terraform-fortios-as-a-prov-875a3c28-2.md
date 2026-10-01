---
id: collect-261001-fortinet/fortinet/document-fortigate-7-4-0-administration-guide-763117-terraform-fortios-as-a-prov-875a3c28-2
title: "dst = \"110.2.2.122/32\""
domain: fortinet
role: reference
task: reference
actors: ["Apple"]
dates: []
keywords: []
source: docs/RAG/collect-261001-fortinet/document-fortigate-7-4-0-administration-guide-763117-terraform-fortios-as-a-prov-875a3c28.md
source_anchor: ""
source_lines: [57, 82]
sha256: baff43e7cf01e92efc70638ea4bb33c2ce768fed55bd54e8704d6ebd41c8fd65
---

# dst = "110.2.2.122/32"

  - Edit the configuration file:# Configure the FortiOS Provider
provider "fortios" {
hostname = "10.6.30.5"
token = "17b********************63ck"
}
resource "fortios_system_setting_dns" "test1" {
primary = "96.45.45.45"
secondary = "208.91.112.22"
}
#resource "fortios_networking_route_static" "test1" {
# dst = "110.2.2.122/32"
# gateway = "2.2.2.2"
# blackhole = "disable"
# distance = "22"
# weight = "3"
# priority = "3"
# device = "port2"
# comment = "Terraform test"
#}
  - Entering terraform apply deletes the static route that is commented out of the configuration file, and reverts the DNS address to the old address:root@mail:/home/terraform# terraform apply fortios_system_setting_dns.test1: Refreshing state... (ID: 172.16.95.16) fortios_networking_route_static.test1: Refreshing state... (ID: 2) An execution plan has been generated and is shown below. Resource actions are indicated with the following symbols: ~ update in-place - destroy Terraform will perform the following actions: - fortios_networking_route_static.test1 ~ fortios_system_setting_dns.test1 primary: "172.16.95.16" => "96.45.45.45" secondary: "8.8.8.8" => "208.91.112.22" Plan: 0 to add, 1 to change, 1 to destroy. Do you want to perform these actions? Terraform will perform the actions described above. Only 'yes' will be accepted to approve. Enter a value: yes fortios_networking_route_static.test1: Destroying... (ID: 2) fortios_system_setting_dns.test1: Modifying... (ID: 172.16.95.16) primary: "172.16.95.16" => "96.45.45.45" secondary: "8.8.8.8" => "208.91.112.22" fortios_networking_route_static.test1: Destruction complete after 0s fortios_system_setting_dns.test1: Modifications complete after 0s (ID: 96.45.45.45) Apply complete! Resources: 0 added, 1 changed, 1 destroyed.
- Edit the configuration file:
Troubleshooting
Use the HTTPS daemon debug to begin troubleshooting why a configuration was not accepted:
# diagnose debug enable
# diagnose debug application httpsd -1
|  | The REST API 403 error means that your administrator profile does not have sufficient permissions. The REST API 401 error means that you do not have the correct token or trusted host. |
