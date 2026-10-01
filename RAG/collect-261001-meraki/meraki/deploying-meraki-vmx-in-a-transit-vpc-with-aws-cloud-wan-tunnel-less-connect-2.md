---
id: collect-261001-meraki/meraki/deploying-meraki-vmx-in-a-transit-vpc-with-aws-cloud-wan-tunnel-less-connect-2
title: "deploying-meraki-vmx-in-a-transit-vpc-with-aws-cloud-wan-tunnel-less-connect"
domain: meraki
role: reference
task: reference
actors: ["AWS"]
dates: []
keywords: ["aws"]
source: docs/RAG/collect-261001-meraki/deploying-meraki-vmx-in-a-transit-vpc-with-aws-cloud-wan-tunnel-less-connect.md
source_anchor: ""
source_lines: [67, 203]
sha256: 9c32734b523918ffd486dfd707baddf054678aeae6c883994e95fa3026002902
---

# deploying-meraki-vmx-in-a-transit-vpc-with-aws-cloud-wan-tunnel-less-connect

1. In AutoVPN, hubs automatically share prefixes they learn with each other
2. Hubs will perform path pre-pending for any spoke prefixes they learn as they advertise them to their EBGP peers, where the primary hub for a spoke prefix will announce its AS_PATH once, the secondary will announce its AS_PATH twice and so forth
3. However, for prefixes a hub learns from a different hub, no path pre-pending is done
4. This means that EBGP peers learning of a prefix belonging to a different hub from the ones they're peered to directly may interchangeably use a any of the peers they're learning this prefix from using ECMP and introducing asymmetric routing in the network, which is often undesirable

Additionally, if you do not disable hub to hub tunnels, traffic traversing two vMX hubs in different AWS regions (e.g. so that a spoke in Region 1 communicates with a spoke in Region 2) might not use the AWS Cloud WAN backbone and use the public Internet instead, which may partially defeat the purpose of using Cloud WAN. For more information on Meraki's BGP implementation see this documentation article.

You can request for this to be disabled for your organization through Meraki support. If you do end up disabling this functionality, steps 12 and 13 of this guide describe the process to allow cross-regional communications when hubs don't form tunnels to each other.

## Deployment Steps

### Step 1) Prepare AWS environment and deploy SD-WAN VPC in your desired regions.

1. Log into your AWS Console and navigate to the __VPC console__

2. Choose your desired base region, for example us-west-1 in our case.

3. Create the required VPC resources for the SD-WAN VPC and the Production VPC as mentioned on the AWS VPC getting started __guide__. Optionally, add tags to your resources for future reference and automation.

4. Deploy the necessary subnets. It is recommended to deploy at least two subnets in the SD-WAN, each in a separate AZ for your vMXs to be highly available. In this case we will be deploying two subnets in two dedicated AZs.

5. Make sure that the VPC route tables allow access to the Internet through an __Internet Gateway__ or __NAT Gateway__.

6. Repeat steps 1-5 for each other region where you will have Transit VPCs. In the case of the guide, us-west-1 and us-east-1.

### Step 2) Deploy vMXs in Dashboard and AWS

1. 
    Deploy two vMXs in each of the VPCs created in Step 1. Make sure to reference the VPCs and subnets created before when deploying.
2. 
    You can deploy vMXs following either the __manual step by step deployment guide__ , or the automated deployment using the__Cloud Integrations functionality.__ This second option requires API Access Keys to AWS with the proper permissions, so if this is not feasible use the manual option.
3. 
    Make sure to choose the appropriate vMX size for your deployment using our __Meraki MX sizing guide__ .
4. 
    Once your vMXs are deployed, navigate to your Meraki Dashboard and to each of the four vMX networks.
5. 
    In each network, navigate to Security & SD-WAN → Addressing and VLANs. Set the mode of operation to Passthrough or VPN concentrator.

1. 
    In each network, navigate to Security & SD-WAN → Site-to-site VPN. Set the VPN mode to Hub.

### Step 3) Set up AWS Cloud WAN

1. 
    Navigate to Cloud WAN in your AWS console

1. 
    Click on Create global network and give your global network a name, and then click Next

1. 
    Specify a name for the core network.

1. 
    Scroll down to core network policy settings, and specify a range of ASNs for your core network. These are used to provision CNEs in your regions, with each CNE requiring a dedicated ASN. Specify the edge locations where you want to deploy these CNEs (us-east-1 and us-west-1 in this case), and create an SDWAN segment to start.

1. 
    Click create global network.

1. 
    Navigate to Core network - attachments and click Create attachment.

1. 
    Deploy an attachment to your first Transit VPCs. Set the attachment type to VPC.

1. 
    Select appliance mode support, and reference your VPC’s ID, as well as the two subnets you deployed in each of the transit VPCs. Optionally, add tags for future reference.

1. 
    Repeat steps 6-8 for your second Transit VPC.

### Step 4) Edit your core network policy

1. Navigate to Core network → Policy versions

2. Select the latest executed policy and click on Edit

3. Create an Inside CIDR blocks and provision a block of addresses. This block of addresses will be used to provision BGP peering points for your connect attachments at each CNE. Each connect attachment peer will use up 2 IP addresses.

4. Edit the edge location for one of the regions

1. 
    Allocate an ASN for this edge location from the pool assigned to the core network, and an Inside CIDR block for it from the large block you configured above. This CIDR block is used to provision BGP peers in Connect attachments for this edge location, so consider future growth (each connect attachment uses at least 2 of these addresses).

6. Edit the other edge location and provision an ASN and Inside CIDR block for it.

7. Select Segments, and choose the Create segment option.

8. Provision a new segment for the Production workloads. If you will have other types of workloads (Production, Development, Testing) make sure to add segments for those.

9. Navigate to attachment policies and choose to Create a new attachment policy.

10. Assign rule number 100, specify a policy to assign any attachment with the tag/value pair segment/SDWAN to the SDWAN segment.

11. Repeat this process for the Production segment

12. Select Create policy

13. From the policy versions screen, select your newly created policy (likely the one labeled LATEST), and choose View or apply change set, and then choose Apply change set.

### Step 5) Edit the route table for the SD-WAN VPC and verify that the vMX Security Group has the appropriate permissions

1. 
    Navigate to the SD-WAN VPC in each of your regions and update the route table so that RFC1918 prefixes (10.0.0.0/8, 172.16.0.0/12, 192.168.0.0/16) point to the newly created core network attachment.

1. 
    Navigate to EC2 in each of your regions, and select Security Groups

1. 
    Create a security group for your Transit VPC vMXs, and for the Inbound rules allow TCP port 179 from the Inbound CIDR you created for Cloud WAN in the previous section (allows BGP peering). For the Outbound rules allow all traffic to 0.0.0.0/0. This ensures your vMX is able to communicate with the Meraki Cloud.

1. 
    Navigate to your instances in each region and apply the new security group to all transit vMX instances.

### Step 6) Deploy Tunnel-less Connect Attachment

1. 
    In your AWS Console, navigate to Core network → Attachments, and select Create attachment
2. 
    Choose Connect attachment as the attachment type, and choose one of the two regions with Transit VPCs

1. 
    Specify Tunnel-less (No Encapsulation) as the Connect protocol, and select the VPC attachment for the Transit VPC as the Transport Attachment ID

1. 
    Make sure to add a Tag at the bottom with Key **SDWAN** and Value**cloudwan-segment** . This tag makes it so the attachment is associated with the SDWAN segment of Cloud WAN.

1. 
    After deployment, verify that your Connect attachment is associated with the SDWAN segment by clicking on it and clicking details.

1. 
    Navigate to your Meraki Dashboard, and select the network of your first vMX in this transit VPC
2. 
    Navigate to Security & SD-WAN → Appliance Status → Uplink, and make note of the private IP address assigned to the vMX uplink

1. 
    Navigate to Security & SD-WAN → Routing, and turn on BGP for this vMX, and specify a private ASN for your AutoVPN mesh (this is shared across the entire mesh) but leave the peers blank for now (don’t save yet)

