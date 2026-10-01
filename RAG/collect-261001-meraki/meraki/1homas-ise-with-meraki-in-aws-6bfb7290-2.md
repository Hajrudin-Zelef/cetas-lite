---
id: collect-261001-meraki/meraki/1homas-ise-with-meraki-in-aws-6bfb7290-2
title: "1homas-ise-with-meraki-in-aws-6bfb7290"
domain: meraki
role: reference
task: reference
actors: ["AWS"]
dates: []
keywords: ["aws", "parameters"]
source: docs/RAG/collect-261001-meraki/1homas-ise-with-meraki-in-aws-6bfb7290.md
source_anchor: ""
source_lines: [99, 285]
sha256: 8974f97bce0e42d6d83010c0ffeb39c744580a2f57546699689cc503153c1750
---

# 1homas-ise-with-meraki-in-aws-6bfb7290

1. From the left menu, select `Virtual Private Cloud > Route Tables`
2. Click `Create Route Table`
3. Create the *Public* Route Table
  - Name: `Public-RT`
  - VPC: `ISEinAWS` Click`Create Route Table`
4. Name: 
5. From the left menu, select `Virtual Private Cloud > Route Tables`
6. Create the *Private* Route Table
  - Name: `Private-RT`
  - VPC: `ISEinAWS` Click`Create Route Table`
7. Name: 
8. Associate the respective route tables and subnets:
  1. Check ✅ `Public-RT` and select the`Subnet Assocations` tab below
  2. Click `Edit Subnet Associations` , check ✅`Public-Subnet` and click`Save Associations`
  3. Check ✅ `Private-RT` and select the`Subnet Assocations` tab below
  4. Click `Edit Subnet Associations` , check ✅`Private-Subnet` and click`Save Associations`
9. Check ✅ 

1. Check ✅ `Public-RT` , select the`Routes` tab below, and click`Edit Routes`
2. Click `Add Route` and add a`0.0.0.0/0` default route to the target type`Internet Gateway` >`igw-*` then click`Save Changes`
3. Your Public Route Table should now look like this:
Destination Target `172.31.0.0/16`local `0.0.0.0/0`igw-0ed2b71a84a9588ac

1. Login to the Meraki Dashboard
2. From the network menu, choose `Create a New Network`  - Network name: `ISEinAWS`
  - Network type: `Security Appliance`
  - Network Configuration: `Default Meraki configuration`
  - Check ✅ the **vMX Serial Number** you want to use
Click`Create Network`
3. Network name: 
4. Scroll down and select `Generate Authentication Token` and copy the text for use in the`User Data` of the vMX instance in the next section.

1. In the AWS Search Box at the top of the page, search for "Meraki vMX" and open the link to Cisco Meraki vMX
2. Click `Continue to Subscribe`
3. Click `Continue to Configuration`
4. Choose your region (`us-west` ) and click`Continue to Launch`  1. ami-09db17cd0ae68ce37
5. For **Choose Action** , choose`Launch Through EC2` then click`Launch`
6. Choose the appropriate Instance Type, `c5.large` and click`Next: Configure Instance Details`
7. For Configure Instance Details use the settings:
  - Network: `vpc-* | ISEinAWS`
  - Subnet: `subnet-* | Public-Subnet`
  - Auto-assign Public IP : `Enable`
  - User data: `As Text` ,*paste the vMX Authentication Token from the Meraki Dashboard*
8. Network: 
9. Click `Next: Add Storage`
10. Click `Next: Add Tags`
11. Click `Add Tag` and use a key:value of`Name` :`vMX`
12. Click `Next: Configure Security Group`
13. For the Security Group, use the settings:
  - Assign a security group : `⦿ Create a new security group`
  - Security group name: `vMX-SG`
  - Description: `vMX-SG`
  - Change the existing rule to be:
    - Type : `All Traffic`
    - Description: `Allow All`
  - Type : 
14. Assign a security group : 
15. Click `Review and Launch`
16. Click `Launch`
17. Select your Key Pair `ISEinAWS.pem` , acknowledge that you have the private key file and click`Launch Instances`
18. After launching, select `View Instances` and you should see your instance running!
19. Check ✅ your vMX, choose `Actions > Networking > Change Source / destination check` , ✅`Stop` Source / destination checking and click`Save`

1. 
In the AWS Console, go to `Services > EC2`
2. 
Click `Launch Instances`
3. 
Find a free tier Linux AMI such as the `Amazon Linux 2 AMI (HVM)` and click it's`Select` button
4. 
Choose the `t2.micro` instance type and click`Next: Configure Instance Details`
5. 
For Configure Instance Details use the settings: 
  - Network: `vpc-* | ISEinAWS`
  - Subnet: `subnet-* | Private-Subnet`
  - Auto-assign Public IP : `Enable)`
6. Network: 
7. 
Click `Next: Add Storage`
8. 
Click `Next: Add Tags`
9. 
Click `Add Tag` and use a key:value of`Name` :`Ping`
10. 
Click `Next: Configure Security Group`
11. 
For the Security Group, use the settings: 
  - Assign a security group : `⦿ Create a new security group`
  - Security group name: `SSH+Ping`
  - Description: `SSH+Ping`
  - Click `Add Rule` and choose :
    - Type: `All ICMP-IPv4`
    - Source: `0.0.0.0/0`
  - Type: 
12. Assign a security group : 
13. 
Click `Review and Launch`
14. 
Click `Launch`
15. 
Select your Key Pair `ISEinAWS.pem` , acknowledge that you have the private key file and click`Launch Instances`
16. 
Click `View Instances` and you should see your instance running!
17. 
Login to your Linux VM is you need to do any troubleshooting or add software: ssh -i "~/.ssh/ISEinAWS.pem" ec2-user@{ hostname | IP }

1. In the Meraki Dashboard, view your `ISEinAWS` network
2. Choose `Security & SD-WAN > Configure > Site-to-Site VPN`
3. For the Site-to-Site VPN settings, use:
  - Type: `Spoke`
  - Local network:
Network VPN mode Subnet Private-Subnet Enabled `172.31.2.0/24`
  - NAT traversal : `⦿ Automatic`
4. Type: 
5. Click `Save Changes`
6. Choose `Security & SD-WAN > Appliance Status` and you should now see the vMX public WAN address and it should be**Active** !

Now you will connect your other MX in the mesh to the vMX

1. Choose your Meraki `Lab` network for your physical, on-premise MX
2. Choose `Security & SD-WAN > Configure > Site to Site VPN`
3. For the Site-to-Site VPN settings, use:
  - Type: `Hub (Mesh)`
  - Local Networks :
💡 you may use a Single LAN or VLANs - I used a Single LAN to keep it simple | Network | VPN mode | Subnet | |-------------|----------|--------| | Main Subnet | Enabled | `192.168.101.0/24`
  - NAT traversal : `⦿ Automatic`
  - If you have alreaady configured a Meraki MX has a hub, you should see `ISEinAWS` under**Remote VPN participants** !
4. Type: 
5. Click `Save Changes`

You will need to update the `Private-RT` to the `Lab` MX

1. In the AWS Console, go to `Services > VPC` and select the`Route Tables`
2. Check ✅ `Private-RT` and select the`Routes` tab below
3. Click `Edit Routes`
4. Select `Add Route`  - Destination: `192.168.0.0/16`
  - Target: `Instance` >`i-* | vMX` Click`Save Changes`
5. Destination: 
6. Select `Add Route`  - Destination: `0.0.0.0/0`
  - Target: `Internet Gateway` >`igw-*` Click`Save Changes`
7. Destination: 
8. Your **Private Route Table** should now look like this:Destination Target `172.31.0.0/16`local `192.168.0.0/16`i-* / vMX `0.0.0.0/0`igw-*
9. You can now try to ping through your site-to-site VPN to the Linux instance

1. 
In the AWS Console, go to `Services > EC2` and
2. 
Select the `Ping` Linux VM instance and locate it's`Private IPv4 addresses`
3. 
SSH to the instance: `ssh -i ~/.ssh/ISEinAWS.pem ec2-user@172.31.2.35`

1. In the AWS Search Box at the top of the page, search for "Cisco ISE" and open the link to Cisco Identity Services Engine (ISE)
2. Click `Continue to Subscribe`
3. Click `Continue to Configuration`
4. For Configure This Software, choose:
  - Delivery Method: `Cloud Formation Template` then choose the normal or GovCloud version of`Cisco Identity Services Engine (ISE)`
  - Software Version: `3.1`
  - Region: `us-west` Click`Continue to Launch`
5. Delivery Method: 
6. For Launch this Software, choose: `Launch CloudFormation` then click`Launch`
7. Choose `Upload a template file` , choose your file (`files/ISE-3-1-518.CFT.yaml` ) and click`Next`
8. Verify and complete the ISE CloudFormation Parameters then click `Next`  - Stack Name: `ISE-3-1-518`
  - AMIid: `ami-00a1a68f5519aa150` # ISE in us-west-1
  - Hostname: `ise`
  - KeyName / Instance Key Pair: `ISEinAWS`
  - Management Security Group: `ISE`
  - Management Network: `Private-Subnet`
  - TimeZone: `America/Los_Angeles`
  - ISEInstanceType: `c5.4xlarge`
  - EBS Encryption : `false`
  - Volume Size / StorageSize: `300`
  - DNSDomain: `aws.local`
  - NameServer: `208.67.222.222`
  - NTPServer: `time.nist.gov`
  - ERS: `yes`
  - OpenAPI: `yes`
  - pxGrid: `no`
  - pxGridCloud: `no`
  - ERS: `yes`
  - password: `C1sco12345`
  - ConfirmPwd: `C1sco12345` Click`Next`
9. Stack Name: 
10. Review the Stack Options then click `Next`
11. Do a final review of the parameters then click `Creat Stack`

To create another instance:

