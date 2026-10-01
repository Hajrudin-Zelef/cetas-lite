---
id: collect-261001-meraki/meraki/meraki-cloud-services-azure-aws-gcp-vpn-security-gateway
title: "meraki-cloud-services-azure-aws-gcp-vpn-security-gateway"
domain: meraki
role: reference
task: reference
actors: ["AWS", "Google"]
dates: []
keywords: ["aws"]
source: docs/RAG/collect-261001-meraki/meraki-cloud-services-azure-aws-gcp-vpn-security-gateway.md
source_anchor: ""
source_lines: [1, 39]
sha256: 255af407014127ed12b7e02e42e539d20d92cf23d7133865fe998fc2f201e3e4
---

# meraki-cloud-services-azure-aws-gcp-vpn-security-gateway

## **Configuring Meraki To Azure Site-To-Site-VPN Tunnels**

## **Create Azure VPN Gateway:**

## **Create VPN connection:**

## **Verify the VPN connection:**

In the Azure portal, you can view the connection status of a Resource Manager VPN Gateway by navigating to the connection.

In the Meraki portal, you can view the VPN status of a Meraki by navigating to the Non-Meraki peer.

## **CONFIGURING MERAKI TO AWS SITE-TO-SITE-VPN TUNNELS**

## **Allocate a subnet:**

## **Configure the VPN connection on Meraki's side:**

E. Note: while making a request to a host on the other side of the Site-to-Site VPN, it will take a few attempts for the request to complete while the tunnel is initialized. The more traffic sent across the tunnel the less likely this lag is to occur as the tunnel will stay up. This often leads to people writing quick ping scripts that send a ping every couple second to keep the tunnel up.

## **Configuring Meraki To GCP Site-To-Site VPN**

## **Google Cloud Setup:**

## **Additional VPC Configuration:**

The virtual MX appliance will allow for site-to-site VPN connectivity using Auto VPN between GCP and other remote MXs. In order to have proper bidirectional communication between remote subnets that are terminating into GCP via the vMX and hosts within GCP, the VPC routing table must be updated for the remote Auto VPN-connected subnets.

## **Firmware Version:**

In order for the vMX to function on GCP it must be running 16.8+ firmware.

## **Confirming Cloud Reachability:**

By default, HTTP traffic inbound to the vMX is disabled for security purposes. You can enable inbound HTTP traffic to the vMX (for accessing the local status page) by performing the following:

## **No "Add vMX" Button:**

When navigating to Security & SD-WAN > Appliance Status, if there is no "Add vMX" button, please ensure the following two conditions are met:
