---
id: collect-260926-mikrotik/mikrotik/github-oallart-mikrotik-routeros7-vlan-documentation-and-code-to-manage-vlans-in-routeros--2
title: "github-oallart-mikrotik-routeros7-vlan-documentation-and-code-to-manage-vlans-in-routeros-7-via-manu"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: ["aws", "decode", "lean"]
source: docs/RAG/lot-mikrotik/RouterOS/github-oallart-mikrotik-routeros7-vlan-documentation-and-code-to-manage-vlans-in-routeros-7-via-manu.md
source_anchor: ""
source_lines: [146, 235]
sha256: 7d1657c3ea37ef23977e804e78708598988e202272ea1f7acafc785a13f83d96
---

# github-oallart-mikrotik-routeros7-vlan-documentation-and-code-to-manage-vlans-in-routeros-7-via-manu

*Note 2:* Do **NOT** mix manual and automated configuration. If something is set manually, keep it that way. If something is configured via terraform, don't mess with it manually. You have been warned.

The best tool for automating ROS in my view is terraform. Its provider is much more advanced than the ansible one. It also can be converted to use with pulumi if you want a different language than hcl. Still, the provider only does so much and a lot of the logic needs to be computed in hcl before being passed to the API. As such, there is a chicken and egg issue:

- configuring API for automation
- automating API configuration

I have taken the approach of a script that sets up the device all the way to the automation part. That meant preconfiguring the bridge and a management port on a given VLAN. It also configures the API requirements like certificates.

For my use, I leverage flat configurations in yaml file and tofu worskspaces to distinguish the yaml files based on device name. The state is stored in a minio s3 type object storage, although it could just be a flat local directory.

The chicken and egg problem: I want to automate configuration, but to automate configuration I need some basic configuration. My solution: A script (initial_script) that sets up the bare minimum so we can access the device through the https REST API. It sets up:

- hostname
- self-signed certificate (required for https)
- api access using the self-sgned cert
- api user (prompts for password)
- common bridge with vlan-filtering enabled
- management VLAN
- management port for that management VLAN
- management IP address, route and DNS (nameserver is only useful for updates)
- SNMP default community removal

These are all single operations that don't change in the lifecycle of the device. More can be added as needed. The rest can be managed via APi, but now we have our entry point

** Reminder: start from a VERY LEAN AND CLEAN BASE **

- adapt it to your needs with the proper IP and settings
- log in, navigate to folder, copy paste, run ...

Terraform has the best *provider* but it's still incomplete and the features I really need are missing, therefore I do some heavy lifting in terraform logic. This document makes use of the `terraform-routeros/routeros` provider

```
terraform {
  required_providers {
    routeros = {
      source = "terraform-routeros/routeros"
    }
  }
```
The workspaces feature of terraform allows the same code to be used for different routeros devices. Each device has its own workspace. Leveraging yaml (see further), we can give each device their own yaml configuration file based on the workspace.

After defining a s3 bucket (I use a minio container) in the code, init tofu

```
  backend "s3" {
    endpoints = {
      s3 = "http://your.server:9000"
    }
    shared_credentials_files    = [".aws/credentials"]
    profile                     = "minio"
    bucket                      = "mikrotik-tfstate"
    key                         = "terraform.tfstate"
    region                      = "main"
    skip_credentials_validation = true
    skip_metadata_api_check     = true
    skip_region_validation      = true
    use_path_style              = true
  }
```
```
tofu init -reconfigure
tofu workspace new mk-sw1.yours.lan
```
There are many ways to go about it but my approach uses the following syntax:

```
---
mk-sw1.yours.lan:
  hardware: RB5009UG+S+
  release: 7.20
  interfaces:
    sfp-sfpplus1: { role: ~ , bond: 1, pvid: 1,  tagged: [100, 110, 999] }
    ether1:       { role: ~ , bond: ~, pvid: 30, tagged: [] }
    ether2:       { role: ~ , bond: ~, pvid: 30, tagged: [] }
    ether3:       { role: ~ , bond: ~, pvid: 30, tagged: [] }
```
Which is loaded to be processed as such:

```
locals {
  # Load yaml by workspace name
  yaml_file = "${terraform.workspace}/config.yaml"
  # Decode the inventory YAML file
  config = yamldecode(file(local.yaml_file))
}
```
This is the first "layer": Bonded interfaces (Link Aggregation) sit on top of the physical interface where needed. The new virtual interface can then be used as a normal interface to create VLANs etc.

*WIP* not functional yet
