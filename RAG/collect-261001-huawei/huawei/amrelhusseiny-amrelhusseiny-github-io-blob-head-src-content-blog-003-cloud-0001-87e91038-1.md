---
id: collect-261001-huawei/huawei/amrelhusseiny-amrelhusseiny-github-io-blob-head-src-content-blog-003-cloud-0001-87e91038-1
title: "Switch Configuration"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2022-08-25"]
keywords: ["license", "memory"]
source: docs/RAG/collect-261001-huawei/amrelhusseiny-amrelhusseiny-github-io-blob-head-src-content-blog-003-cloud-0001--87e91038.md
source_anchor: ""
source_lines: [1, 137]
sha256: 4ab4341ec03fb6f9ac6969ea82e8fda9c0d5506bc6f87338953031ada096af64
---

# Switch Configuration

| title | Managing Huawei Cloud_Engine switches using Ansible | 
|---|---|
| description | helooooooooooo | 
| date | 2022-08-25 12:19:49 +0200 | 
| draft | false | 
| tags |  | 
| automation | 
I have been working with the Huawei Cloud Engine series of switches for the Data Center, and one of the great things I was introduced to is the efficient usage of Ansible+Netconf to configure Legacy network devices instead of the trial and error handling way of Python+SSH , which you had to anticipate the delay of the SSH session and to handle a console output using Regex which while very customizable , when it comes to production environments , your supervisors would much prefer a tried and tested technology backed up by the vendor itself .
Here I will share the steps I was able to start my journey in configuring the Cloud Engine switches using Ansible .
1st I would advise you to configure an EVE-NG Community Environment or a GNS3 environment so you are able to test your playbooks (Ansible) before deploying them to the production environment .
For my deployment , I will be using EVE NG , you can find the download page (https://www.eve-ng.net/index.php/download/) , download the Community version , as it won't require a retail license .
You can download it as ISO to install or and OVF to run directly , for me , I am running it on ESXI machine .
After you downloaded and installed EVE-NG , you should be able to reach the eve machine using SSH 1st to configure it , default SSH user is "root" and default password is "eve" .
When you login to SSH , you will follow a series of configuration prompts for DNS , NTP and so on , after which you will be able to login to the EVE-NG Web portal to start using the Emulation software .
On the Web portal , the default user is "admin" and password is "eve" .
In order to work with EVE , you will have to use a preconfigured image , for our device , please follow the following link to download the Huawei CE12800 VM that we will use : https://forum.huawei.com/enterprise/en/run-ce12800-ne40e-in-eve-ng/thread/653457-861
After you followed the above guide to add the Huawei Cloud Engine image to EVE , we will be able to create my lab setup :
I created an ESXi CentOS machine to run the Ansible playbook from , you can install locally on your laptop or machine if you would like , we will be able to reach our Huawei switch from our local Home network .
- 
Add new Lab
- 
Go to left side node , "Add an object -> Node -> search Huawei , you find Cloud Engine 12800" , click on it , you will need 2 GB of RAM/Memory reserved per switch to run it :
- 
Next , to bridge the switch to our local Network , so we can access it from our local Laptop or in my case from my CentOS VM , you need to go to "Add an object -> Network -> Management (Cloud0) " , and connect one of the switches interface to the cloud (ex: GE1/0/0) , and right click on the switch then choose start :
##C) Configuring Huawei Cloud Engine for Netconf :
$ telnet 192.168.1.109 32769
- On device , paste the following configuration to add a user and start the Netconf Service :
# Switch Configuration
system
!
interface GE 1/0/0
undo portswitch
undo shutdown 
ip address 192.168.1.130 24
quit
!
ip route-static 0.0.0.0 0 192.168.1.1 
!
netconf
protocol inbound ssh port 830
!
ssh user client001
!
aaa
local-user client001 password irreversible-cipher SetUesrPasswd@123
local-user client001 service-type ssh
quit
commit
!
ssh server cipher aes128_ctr aes256_ctr aes192_ctr aes128_gcm aes256_gcm
ssh user client001 authentication-type password
ssh user client001 service-type snetconf
!
snetconf server enable
!
commit
!
save
```bash
- To confirm that you are able to reach network connectivity , ping your gateway and your Ansible hosting device : 
```text
<CloudEngine>ping -c 5 192.168.1.1 
  PING 192.168.1.1: 56  data bytes, press CTRL_C to break
    Reply from 192.168.1.1: bytes=56 Sequence=1 ttl=64 time=5 ms
    Reply from 192.168.1.1: bytes=56 Sequence=2 ttl=64 time=3 ms
    Reply from 192.168.1.1: bytes=56 Sequence=3 ttl=64 time=4 ms
    Reply from 192.168.1.1: bytes=56 Sequence=4 ttl=64 time=2 ms
    Reply from 192.168.1.1: bytes=56 Sequence=5 ttl=64 time=3 ms
  --- 192.168.1.1 ping statistics ---
    5 packet(s) transmitted
    5 packet(s) received
    0.00% packet loss
    round-trip min/avg/max = 2/3/5 ms
 
<CloudEngine>ping -c 5 192.168.1.102 
  PING 192.168.1.102: 56  data bytes, press CTRL_C to break
    Reply from 192.168.1.102: bytes=56 Sequence=1 ttl=64 time=2 ms
    Reply from 192.168.1.102: bytes=56 Sequence=2 ttl=64 time=3 ms
    Reply from 192.168.1.102: bytes=56 Sequence=3 ttl=64 time=3 ms
    Reply from 192.168.1.102: bytes=56 Sequence=4 ttl=64 time=4 ms
    Reply from 192.168.1.102: bytes=56 Sequence=5 ttl=64 time=1 ms
  --- 192.168.1.102 ping statistics ---
    5 packet(s) transmitted
    5 packet(s) received
    0.00% packet loss
    round-trip min/avg/max = 1/2/4 ms
```bash
- To test the Netconf before moving forward to ansible , from your laptop or Ansible hosting VM , run the following command , you should get a long XML output , this means that we are ready for ansible :
`$ ssh client001@192.168.1.130 -p 830 netconf`
## 4- Setup of the Ansible machine: 
Following are the steps for a CentOS machine 
### A)Create a virtual Python environment:
Recommended in the time of writing the article to use Python 3.8 and above.
Create a virtual environment , in my case using Python 3.10 :
`$ sudo /usr/local/bin/python3.10 -m venv huawei_venv`
Activate the created virtual environment , you will find the default Python version is the one you created it with .
The Ansible module we will be using is part of the net_commons collection : 
```bash
(huawei_venv)$ python --version
Python 3.10.5
# Install Ansible and needed libraries for the lab
(huawei_venv)$ sudo /usr/local/bin/python3.10  -m pip install ansible ansible-core ncclient jxmlease xmltodict ansible-pylibssh
# Check Ansible version 
(huawei_venv)$ ansible --version
ansible [core 2.13.3]
  config file = None
  configured module search path = ['/home/amroashram/.ansible/plugins/modules', '/usr/share/ansible/plugins/modules']
  ansible python module location = /usr/local/lib/python3.10/site-packages/ansible
  ansible collection location = /home/amroashram/.ansible/collections:/usr/share/ansible/collections
  executable location = /usr/local/bin/ansible
  python version = 3.10.5 (main, Aug  8 2022, 17:01:37) [GCC 8.5.0 20210514 (Red Hat 8.5.0-15)]
  jinja version = 3.1.2
  libyaml = True
# Double confirm that the Netcommons collection is installed 
(huawei_venv)$ ansible-galaxy collection install ansible.netcommon
```bash
### B) lets create a simple Inventory and Playbook to run on our current switch :
Inverntory.yaml
```yaml
# inventory.yaml
all:
  # ----- Variables for the whole inventory 
  vars:
    ansible_network_os: community.network.ce
    ansible_user: client001
    ansible_password: "SetUesrPasswd@123"
    ansible_ssh_common_args: '-o ProxyCommand="ssh -W %h:%p -q bastion01"'
    # log_path: ansible.log 
  # ----- Group names and relations to hosts 
  children:
    network:
      children:
        huawei:
          hosts:
            leaf_1:
              ansible_host: 192.168.1.130
              
