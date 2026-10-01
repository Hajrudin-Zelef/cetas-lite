---
id: collect-261001-meraki/meraki/1homas-ise-with-meraki-in-aws-6bfb7290-1
title: "1homas-ise-with-meraki-in-aws-6bfb7290"
domain: meraki
role: reference
task: reference
actors: ["AWS"]
dates: ["2021-10-05"]
keywords: ["aws", "compute", "license"]
source: docs/RAG/collect-261001-meraki/1homas-ise-with-meraki-in-aws-6bfb7290.md
source_anchor: ""
source_lines: [1, 98]
sha256: 63cfa71d0138ed7a9a626ae6f8fd9aeea44cb285c5d85e39b26e580154da6395
---

# 1homas-ise-with-meraki-in-aws-6bfb7290

This demo environment was created for use with the Cisco ISE with Meraki Webinar on October 5, 2021. Watch the recording of the Cisco ISE with Meraki Webinar in the Cisco ISE YouTube Channel:

You will need at least one additional Meraki MX or Z network to act as a VPN hub to terminate the other side of the VPN connection. You should have your hub MX working *before* you run the Ansible playbook because it will attempt to connect the vMX to an existing `Lab` VPN hub. You may edit the `vars/main.yaml` file to customize your VPN hub network name. I chose to have the vMX network in AWS be a VPN *spoke* because I do not want AWS to charge me for all of my network traffic flowing through AWS to the Internet!

Running `ansible-playbook ise_in_aws.yaml` will create :

- AWS VPC, subnets, route tables, internet gateway. For the gory details, see **Manual Configuration in AWS Console and Meraki Dashboard** below.
- ISE 3.1 or later instance
- Meraki vMX instance to secure RADIUS traffic from network devices to ISE
- Linux VM instance to ping while ISE boots so you will feel confident that the VPN is going to work! You could also turn this into a web server for testing URL redirections or as an internal site that you block.

1. 
Clone this repository: ```
git clone https://github.com/1homas/ISE_with_Meraki_in_AWS.git
cd ISE_with_Meraki_in_AWS
```
2. 
Create your Python environment and install Ansible with other Python packages for AWS and ISE : pip install --upgrade pip pip install pipenv pipenv install --python 3.11 pipenv install ansible boto3 botocore ciscoisesdk jmespath paramiko pipenv shell If you have any problems installing Python or Ansible, see Installing Ansible.
3. 
Export your various keys, tokens, and credentials for ISE, Meraki, and AWS APIs into your shell environment. You can store these in one or more `*.env` then load them with the`source` command.# AWS IAM API Keys export AWS_REGION='us-west-1' export AWS_ACCESS_KEY='AKIAIOSF/EXAMPLE+KEY' export AWS_SECRET_KEY='wJalrXUtnFEMI/K7MDENG/bPxRfi/EXAMPLE+KEY' # Meraki API Authentication Key export MERAKI_KEY='EXAMPLE_KEYc320e12ee407159487a4cabc41abb' # ISE REST API Credentials export ISE_USERNAME=admin export ISE_PASSWORD=ISEisC00L export ISE_VERIFY=false export ISE_DEBUG=false export ISE_INIT_PASSWORD=C1sco12345 # Secrets export ISE_RADIUS_SECRET=ISEisC00L export ISE_TACACS_SECRET=ISEisC00L export ISE_SNMP_SECRET=ISEisC00L # Repository & Backups export ISE_REPOSITORY=ftp.trust0.net export ISE_REPOSITORY_PROTOCOL=FTP export ISE_REPOSITORY_PATH=/ export ISE_REPOSITORY_USERNAME=ise export ISE_REPOSITORY_PASSWORD=ISEisC00L export ISE_BACKUP_ENCRYPTION_KEY=ISEisC00L Alternatively, keep your environment variables in files in a `.secrets` or similar folder in your home directory and use`source {filename}` to load environment variables from the files:source ~/.secrets/aws.sh source ~/.secrets/ise.sh source ~/.secrets/ise_repo.sh source ~/.secrets/meraki.sh
4. 
Review the `vars/*.yaml` configuration files and un/comment or edit them to suit your environment. You*must* edit the`vars/main.yaml` and change the`meraki_org_name` to your respective Meraki org. You will want to review the other settings and change them to match your environment:
  - your desired `project_name`
  - the AMI identifiers for your respective AWS region if not `us-west-1`
  - your desired network CIDR ranges
  - your desired VM instance types
  - your Meraki vMX instance type and license (S/M/L)
  - your default password(s) or pre-shared keys
5. your desired 
6. 
Run the Ansible playbook: ansible-playbook ise_in_aws.yaml
7. 
Due to a Meraki VPN API error, you will need to manually add the vMX Local Network definition in the Meraki Dashboard to advertise the VPC subnet: 
  1. In the Meraki Dashboard, view your `ISE_Meraki_AWS` network
  2. Choose `Security & SD-WAN > Configure > Site-to-Site VPN` and for the Local Networks,**Add a Local Network** :Network VPN mode Subnet ISE_Meraki_AWS Enabled `172.31.0.0/16`
 ⚠ If you cannot ping or SSH to the `Ping` Linux VM this is probably the reason!
8. In the Meraki Dashboard, view your 
9. 
When ISE is up, you may configure it using the additional playbook : ansible-playbook ise.configuration.yaml
10. 
When you are done, terminate the instances and delete all resources to prevent surprise AWS bills: ansible-playbook ise_in_aws.terminate.yaml

In case you wondered exactly what these Ansible playbooks are doing ... here is how to do it the hard way! If you want to do it the old-fashioned way or just understand what the time spent on automation is saving you from!

Your AWS instance(s) will have a public IP address so anyone can - and will - eventually find it and try to login and use it. For this reason, AWS does not allow the use of normal passwords. Instead, they use a private/public cryptographic key pair which is *much* stronger than a password.

1. 
Login to the Amazon Web Services (AWS) Console as a root user (not IAM user) of your account
2. 
Verify or choose your **Region** in the drop-down menu next to your account name
3. 
Open the **Services** menu and choose**Compute > EC2**
4. 
From the left menu, choose **Network & Security > Key Pairs**
5. 
Click **Create Key Pair** , fill in the attributes below, and click**Create Key Pair**
  - Name: `ISEinAWS`
  - Key pair type: **RSA**
  - Private key file format: **.pem**
6. Name: 
7. 
When prompted, save the `ISEinAWS.pem` private key file to your home directory in a folder named`.ssh` (`~/.ssh/ISEinAWS.pem` )
8. 
If you are using macOS, Linux, or WSL, change the file permissions so it cannot be viewed by others or accidentally overwritten or deleted by you: `chmod 400 ~/.ssh/ISEinAWS.pem`🛑 Do not lose this private key file! You will not be able to login to your AWS EC2 instances configured with the corresponding public key!

When you create instances in AWS, you may choose to put the matching public key into your VMs to authorize your SSH login. To use your key with AWS EC2 instances, you will connect using SSH and authenticate with the `-i` *identity file* option which is your `ISEinAWS.pem` private key :

`ssh -i ~/.ssh/ISEinAWS.pem admin@{hostname | IP}`
1. In the AWS Console, go to **Services > Networking & Content Delivery > VPC**
2. Choose your region: `us-west1`
3. Select `Create VPC`  - Name: `ISEinAWS`
  - IPv4 CIDR: `172.31.0.0/16`
  - Tenancy: `Default`
  - Add Tag: `Project : ISEinAWS` Click`Create VPC`
4. Name: 

1. From the left menu, select `Virtual Private Cloud > Subnets`
2. Click `Create Subnet`
3. Create your Public subnet for the Meraki vMX:
  1. VPC ID: `ISEinAWS`
  2. Subnet name: `Public-Subnet`
  3. Availability Zone: `No preference` or*choose your desired AZ*
  4. CIDR: `172.31.1.0/24`
4. VPC ID: 
5. Select `Add a New Subnet` for the Private subnet with ISE
  1. VPC ID: `ISEinAWS`
  2. Subnet name: `Private-Subnet`
  3. Availability Zone: `No preference` or*choose your desired AZ*
  4. CIDR: `172.31.2.0/24`
6. VPC ID: 
7. Select `Create Subnet`
8. Check ✅ `Public-Subnet` , choose`Actions > Modify auto-assign IP settings` , check ✅`Enable auto-assign public IPv4 address` and click`Save`

1. From the left menu, select `Virtual Private Cloud > Internet Gateways`
2. Click `Create Internet Gateway`  - Name: `vMX-IG` Click`Create Internet Gateway`
3. Name: 
4. Associate the Internet Gateway to the VPC by selecting `Actions > Attach to VPC` >`ISEinAWS` and click`Attach Internet Gateway`

