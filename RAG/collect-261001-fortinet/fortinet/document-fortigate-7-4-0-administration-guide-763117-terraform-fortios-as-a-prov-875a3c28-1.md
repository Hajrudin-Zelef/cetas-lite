---
id: collect-261001-fortinet/fortinet/document-fortigate-7-4-0-administration-guide-763117-terraform-fortios-as-a-prov-875a3c28-1
title: "dst = \"110.2.2.122/32\""
domain: fortinet
role: reference
task: reference
actors: ["Apple"]
dates: []
keywords: ["memory"]
source: docs/RAG/collect-261001-fortinet/document-fortigate-7-4-0-administration-guide-763117-terraform-fortios-as-a-prov-875a3c28.md
source_anchor: ""
source_lines: [1, 56]
sha256: 43a598297e73fce31cc87ccda2980b6a285eda5b2a99099743c2f809774c2980
---

# dst = "110.2.2.122/32"

Terraform: FortiOS as a provider
Terraform: FortiOS as a provider
Fortinet's Terraform support provides customers with more ways to efficiently deploy, manage, and automate security across physical FortiGate appliances and virtual environments. You can use Terraform to automate various IT infrastructure needs, thereby diminishing mistakes from repetitive manual configurations.
For example, if Fortinet is releasing a new FortiOS version, your organization may require you to test a new functionality to determine how it may impact the environment before globally deploying the new version. In this case, the ability to rapidly stand up environments and test these functions prior to production environment integration provides a resource-efficient and fault-tolerant approach.
The following example demonstrates how to use the Terraform FortiOS provider to perform simple configuration changes on a FortiGate unit. It requires the following:
- FortiOS 6.0 or later
- FortiOS Provider: This example uses terraform-provider-fortios 1.0.0.
- Terraform: This example uses Terraform 0.11.14.
- REST API administrator created on the FortiGate with the API key
For more information, see the Terraform FortiOS Provider at https://www.terraform.io/docs/providers/fortios/index.html.
To create a REST API administrator:
- On the FortiGate, go to System > Administrators and click Create New > REST API Admin.
- Enter the Username and, optionally, enter Comments.
- Select an Administrator Profile.
- We recommend that you create a new profile with minimal privileges for this terraform script:
                
  - In the Administrator Profile drop down click Create New.
  - Enter a name for the profile.
  - Configure the Access Permissions:
                
    - None: The REST API is not permitted access to the resource.
    - Read: The REST API can send read requests (HTTP GET ) to the resource.
    - Read/Write: The REST API can send read and write requests (HTTP GET /POST /PUT /DELETE ) to the resource.
  - Click OK.
- Enter Trusted Hosts to specify the devices that are allowed to access this FortiGate.
- Click OK. An API key is displayed. This key is only shown once, so you must copy and store it securely.
To configure FortiGate with Terraform Provider module support:
- Download the terraform-provider-fortios file to a directory on the management computer.
- Create a new file with the .tf extension for configuring your FortiGate:
			root@mail:/home/terraform# ls terraform-provider-fortios_v1.0.0_x4 test.tf
- Edit the test.tf Terraform configuration file:In this example, the FortiGate's IP address is 10.6.30.5, and the API user token is 17b********************63ck. Your provider information must also be changed. # Configure the FortiOS Provider
provider "fortios" {
hostname = "10.6.30.5"
token = "17b********************63ck"
}
- Create the resources for configuring your DNS object and adding a static route:resource "fortios_system_setting_dns" "test1" {
primary = "172.16.95.16"
secondary = "8.8.8.8"
}
resource "fortios_networking_route_static" "test1" {
dst = "110.2.2.122/32"
gateway = "2.2.2.2"
blackhole = "disable"
distance = "22"
weight = "3"
priority = "3"
device = "port2"
comment = "Terraform test"
}
- Save your Terraform configuration file.
- In the terminal, enter terraform init to initialize the working directory.It reads the provider if the name follows the convention terraform-provider-[name] :root@mail:/home/terraform# terraform init Initializing the backend... Terraform has been successfully initialized! You may now begin working with Terraform. Try running "terraform plan" to see any changes that are required for your infrastructure. All Terraform commands should now work. If you ever set or change modules or backend configuration for Terraform, rerun this command to reinitialize your working directory. If you forget, other commands will detect it and remind you to do so if necessary.
- Run terraform -v to verify the version of loaded provider module:root@mail:/home/terraform# terraform -v Terraform v0. 11.14 + provider.fortios v1.0.0
- Enter terraform plan to parse the configuration file and read from the FortiGate configuration to see what Terraform changes:This example create a static route and updates the DNS address. You can see that Terraform reads the DNS addresses from the FortiGate and then lists them. root@mail:/home/terraform# terraform plan Refreshing Terraform state in-memory prior to plan... The refreshed state will be used to calculate this plan, but will not be persisted to local or remote state storage. fortios_networking_route_static.test1: Refreshing state... (ID: 2) fortios_system_setting_dns.test1: Refreshing state... (ID: 96.45.45.45) ------------------------------------------------------------------------ An execution plan has been generated and is shown below. Resource actions are indicated with the following symbols: + create ~ update in-place Terraform will perform the following actions: + fortios_networking_route_static.test1 id: <computed> blackhole: "disable" comment: "Terraform test" device: "port2" distance: "22" dst: "110.2.2.122/32" gateway: "2.2.2.2" priority: "3" weight: "3" ~ fortios_system_setting_dns.test1 primary: "96.45.45.45" => "172.16.95.16" secondary: "208.91.112.22" => "8.8.8.8" Plan: 1 to add, 1 to change, 0 to destroy. ------------------------------------------------------------------------ Note: You didn't specify an "-out" parameter to save this plan, so Terraform can't guarantee that exactly these actions will be performed if "terraform apply" is subsequently run. If you are running terraform-provider-fortios 1.1.0, you may see the following error: Error: Error getting CA Bundle, CA Bundle should be set when insecure is false. In this case, add the following line to the FortiOS provider configuration in the test.tf file: insecure = "true"
- Enter terraform apply to continue the configuration:root@mail:/home/terraform# terraform apply fortios_system_setting_dns.test1: Refreshing state... (ID: 96.45.45.45) fortios_networking_route_static.test1: Refreshing state... (ID: 2) An execution plan has been generated and is shown below. Resource actions are indicated with the following symbols: + create ~ update in-place Terraform will perform the following actions: + fortios_networking_route_static.test1 id: <computed> blackhole: "disable" comment: "Terraform test" device: "port2" distance: "22" dst: "110.2.2.122/32" gateway: "2.2.2.2" priority: "3" weight: "3" ~ fortios_system_setting_dns.test1 primary: "96.45.45.45" => "172.16.95.16" secondary: "208.91.112.22" => "8.8.8.8" Plan: 1 to add, 1 to change, 0 to destroy. Do you want to perform these actions? Terraform will perform the actions described above. Only 'yes' will be accepted to approve. Enter a value: yes fortios_networking_route_static.test1: Creating... blackhole: "" => "disable" comment: "" => "Terraform test" device: "" => "port2" distance: "" => "22" dst: "" => "110.2.2.122/32" gateway: "" => "2.2.2.2" priority: "" => "3" weight: "" => "3" fortios_system_setting_dns.test1: Modifying... (ID: 96.45.45.45) primary: "96.45.45.45" => "172.16.95.16" secondary: "208.91.112.22" => "8.8.8.8" fortios_networking_route_static.test1: Creation complete after 0s (ID: 2) fortios_system_setting_dns.test1: Modifications complete after 0s (ID: 172.16.95.16) Apply complete! Resources: 1 added, 1 changed, 0 destroyed. The FortiGate is now configured according to the configuration file.
- To change or delete something in the future, edit the configuration file and then apply it again. In supported cases, it deletes, adds, or updates new entries as configured. For instance, in this example you can remove the static route and revert the DNS address to its original configuration by changing the .tf file:
        
