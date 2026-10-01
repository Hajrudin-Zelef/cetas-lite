---
id: collect-261001-meraki/meraki/deploying-meraki-vmx-in-a-transit-vpc-with-aws-cloud-wan-tunnel-less-connect-3
title: "deploying-meraki-vmx-in-a-transit-vpc-with-aws-cloud-wan-tunnel-less-connect"
domain: meraki
role: reference
task: reference
actors: ["AWS"]
dates: []
keywords: ["aws"]
source: docs/RAG/collect-261001-meraki/deploying-meraki-vmx-in-a-transit-vpc-with-aws-cloud-wan-tunnel-less-connect.md
source_anchor: ""
source_lines: [204, 294]
sha256: 312fc0ac2cdedc602544c7182fe1d5296dc07d569dad27f6ec9bcd30da9b4dcd
---

# deploying-meraki-vmx-in-a-transit-vpc-with-aws-cloud-wan-tunnel-less-connect

1. 
    Back in your AWS Console in the Cloud WAN attachments list, select the Connect attachment you just created and choose Connect peers → Create Connect peer

1. 
    For the first Connect peer, specify the IP address of the first vMX in the VPC that you wrote down in item 5, as well as the ASN you specified in 6, and select the subnet this vMX belongs to
2. 
    Wait a few seconds for the Connect peer to show up in your connect peers list, and look for the IP addresses assigned to each of the two BGP interfaces

1. 
    Go back to your Meraki Dashboard where you were configuring your vMX BGP settings, and add the IP address of each of the two BGP interfaces for the Connect peer, and set the EBGP multihop count to 2

1. 
    Repeat steps 7-10 for the other vMX in the VPC (create a second connect peer in the same Connect attachment for it)
2. 
    Repeat steps 1-11 for the Transit VPC in the other region

### Step 7) Connect branches to transit vMXs

1. 
    Navigate back to your Meraki Dashboard
2. 
    For every group of branches that needs to connect to your Transit VPCs, navigate to Site-to-site VPN, specify them as Spokes, and specify the two Transit vMXs you have in their respective region as priority 1 and 2

1. 
    Repeat these steps for all regions where you have Transit VPCs, making sure to group the appropriate branches to it (for example, branches in the West Coast use us-west-1, and branches in the East Coast use us-east-1)

### Step 8) Adding a Workload VPC to the Production Segment

1. 
    Navigate to your AWS Console → Cloud WAN → Core network → Attachments, and for every workload VPC in your environment that needs connectivity to your branches, create a new VPC attachment in the appropriate region, allowing the needed subnets and the tag/value pair segment/Production to make sure the attachment is routed to the proper segment

1. 
    Navigate to your VPC Dashboard for each of these Workload VPCs and update the route tables your instances are using to have a default route to the VPC attachment you created for this VPC. You may choose to do more specific routes in your environment, but we’re using a default route for simplicity in this example.

1. 
    Navigate to your EC2 Dashboard, and make sure to update the Security Group of the instances your branches need reachability to, to allow the necessary protocols and IP addresses inbound. In our case, we’re just allowing ICMPv4 from the RFC1918 addresses for testing connectivity, and SSH from our own public IP. This will likely vary for your environment.

### Step 9) Implement segment sharing between Production and SDWAN

1. By default, segments do not share routes with each other. You can confirm this by navigating to Core network → Routes and entering a given segment and region and hit search. You will see that only that own segment’s routes are populated.
2. To implement segment sharing, navigate to Core network → Policy versions and edit your latest active policy.
3. 3. Navigate to Segment actions and under Sharing click Create
4. 4. Select the SDWAN segment, and choose Allow selected (or all if you want to share with all segments), and pick the desired segments from the list. By default, segment sharing is bidirectional, so sharing SDWAN to Production also shares Production to SDWAN. Click Create sharing.
5. Select create policy.
6. Select your new policy (should say LATEST) and click View or apply change set and click Apply change set in the next screen
7. After your policy finishes executing, navigate back to Core network → Routes and input a segment and region and click search. You should now see the other segment’s routes in this routing table as well.
8. Navigate to your Meraki Dashboard, and choose one of the branches connected to your Transit vMXs. Go to Security & SD-WAN → Route table, and you should see the routes from the Production segment in your route table. You can do the same verification from your Transit vMXs.

### Step 10) Testing connectivity for the first region

1. 
    If you wish to test connectivity to any of your vMXs, you need to add an inbound ICMP rule for RFC 1918 prefixes in the security group associated with them. You can navigate to your EC2 instances, pick a vMX and select the Security tab. Then click on the security group to edit it.
2. 
    Navigate to your Meraki Dashboard, select a branch network and go to Security & SD-WAN → Appliance status → Tools, and initiate pings to the private IPs of your transit VPC vMXs
3. 
    From the same page you can add a few pings toward the EC2 instances within your Workload VPCs in the same region, and any cross-region workloads

### Step 11) Testing redundancy

1. 
    You can test redundancy for your branches by bringing down the primary vMX in their corresponding Transit VPC. You can do this by navigating to EC2 in the AWS Console and choosing the primary vMX. Then from the Instance state choose Stop instance. This process can take up to 5 minutes to complete, so you will want to wait. You can also Force stop the instance, by selecting the instance state again.
2. 
    After the vMX goes down, you can attempt a new ping from the branches to your instances and vMXs. BGP may take some time to converge (up to the hold timer set to 240s by default). After convergence, your workloads should respond again, and the primary vMX should be down.

### Step 12) Allowing branch-to-branch connectivity across regions

1. 
    If you have disabled hub-to-hub connectivity as mentioned in the solution architecture section of this document, and you wish to allow branches connecting to different transit VPCs to use Cloud WAN as transit to communicate with each other, you need to do some additional manual work. This is because since the entire Meraki AutoVPN mesh shares a single ASN, vMXs in opposite regions will reject routes from the other region’s branches, as the originating ASN is its own.

Another potential approach to allow cross-regional spoke connectivity with hub-to-hub tunnels disabled would be to simply add the hubs of the opposing region as additional hubs for the spokes in the local region. However, two side effects of this will be:

1. Each hub vMX will advertise a larger amount of prefixes to Cloud WAN, and you could end up exceeding the 1,000 route limit per connect peer in this way
2. Cross-regional spoke communications will bypass the Cloud WAN backbone and use the public Internet instead, which doesn't reap the same performance benefits

2. This means you need to inject a manual route from Cloud WAN into each region to be able to reach the opposite region.

1. 
    Navigate to Core network → Policy versions, and edit your latest active policy.
2. 
    Navigate to Segment actions, and under Routes select Create
3. 
    Specify the SDWAN segment and add a summary route or multiple summary routes for the branch prefixes in one of the regions, pointing to that region’s VPC attachment. For example, in this case to reach the 172.25.0.0/16 prefix associated with the West Branches, we would specify the next hop attachment as the us-west-1 VPC attachment, and for the 172.26.0.0/16 prefix associated with the East Branches, we would specify the next hop attachment as the us-east-1 VPC attachment.
4. 
    Click create segment route and create policy, and activate the policy.

### Step 13) Testing Cross-Regional connectivity

1. 
    Navigate to your Meraki Dashboard, and you should now see the other regions’ MX routes in the branch MX route tables.
2. 
    Start a ping from a branch in one region and ping a branch in the opposite region. This should now succeed.
