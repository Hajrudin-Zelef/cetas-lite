---
id: collect-261001-automatisation-infra/automatisation-infra/github-baltah666-netbox-pyats-1
title: "Description: This script generates a testbed file based on the Netbox data"
domain: automatisation-infra
role: reference
task: reference
actors: []
dates: ["2024-04"]
keywords: ["ethernet"]
source: docs/RAG/collect-261001-automatisation-infra/github-baltah666-netbox-pyats.md
source_anchor: ""
source_lines: [1, 184]
sha256: 49799ce4d1cba7950b26805dc5220cbb7b7f2b672a6d5eea06a398aa98035ba2
---

# Description: This script generates a testbed file based on the Netbox data

Code to accompany the **Webinar: Getting Started with Network Test Automation: NetBox + pyATS +** hosted by **NetBox Labs** on 23rd April 2024. Click here to watch the webinar on-demand.

For hassle-free access to NetBox you can either use the NetBox Labs demo site, or request a free 14 Day Trial of NetBox Cloud:

1. 
Clone the Git repo and change into the `netbox-pyats-webinar` directory:```
git clone https://github.com/netboxlabs/netbox-learning.git
cd netbox-learning/netbox-pyats-webinar
```
2. 
Create and activate Python 3 virtual environment: ```
python3 -m venv ./venv
source venv/bin/activate
```
3. 
Upgrade pip: ```
python3 -m pip install --upgrade pip
```
4. 
Install PyATS: As per the official documentation, there are a options to perform a minimal installation (option 1) or a full installation (opton 2). **Option 1** Minimal install that includes the Genie library and that allows you to use the interactive testbed creation command to create your testbed files from NetBox:```
pip install pyats[library]
pip install pyats.contrib
```
**Option 2** Full installation that includes all packages and libraries:```
pip install pyats[full]
```
*Note* If you are using Zsh on a Mac then you need to quote the install string (this had me stuck for a long time until I figured it out!)```
pip install "pyats[full]"
```
*Note* if you plan to run the example script`ospf_neighbor_table.py` then you will also need to install the`prettytable` library with`pip install prettytable` or you can simply run the command`pip install -r requirements.txt` to install this along with pyATS, Genie and the`contrib` library.There is also a PyATS Docker Image. This command will pull down the container if you don't have it locally and drop you into a Bash shell: ```
docker run -it ciscotestautomation/pyats:latest /bin/bash
```

Our lab network consists of 2 x Cisco CSR100V routers and they are documented in NetBox under the Site `PyATS Webinar` and are directly connected to each other over port `GigabitEthernet2` on the `192.168.1.0/30` subnet. They are both running OSPF, and you can find the the configuration for this in the initial_device_configs.md file:

**Option 1**
Use the `pyats create testbed netbox` command to build your testbed file. Note that where a value is prefixed with `os.getenv` or `%ENV` then these values are being pulled in from the local environment variables that you need to set with the `export` command eg. `export NETBOX_URL=https://example.cloud.netboxapp.com/`, `export DEF_PYATS_USER=admin`:

```
pyats create testbed netbox \
--output testbed.yaml \
--netbox-url=${NETBOX_URL} \
--user-token=${NETBOX_USER_TOKEN} \
--def_user='%ENV{DEF_PYATS_USER}' \
--def_pass='%ENV{DEF_PYATS_PASS}' \
--url_filter='site=pyats-webinar' \
--topology
```
In this example we are generating a testbed file called `testbed.yaml` and filtering NetBox by the site name `pyats-webinar`. When you hit enter the output will look like this:

```
Begin retrieving data from netbox...
Configuring testbed default credentials.
Retrieving associated data for CSR1...
Retrieving associated data for CSR2...
Testbed file generated: 
testbed.yaml 
```
**Option 2**
Run the `generate_testbed_file.py` Python script. Note that where a value is prefixed with `os.getenv` or `%ENV` then these values are being pulled in from the local environment variables that you need to set with the `export` command eg. `export NETBOX_URL=https://example.cloud.netboxapp.com/`, `export DEF_PYATS_USER=admin`.

In this script we are generating a testbed file called `testbed.yaml` and filtering NetBox by the site name `pyats-webinar`, but you could just as easily filter on other fields as in the examples commented out:

```
# Description: This script generates a testbed file based on the Netbox data
#              using the pyATS framework. It uses the Netbox class from the
#              pyats.contrib.creators.netbox module to create the testbed file.
# Import the necessary libraries
from pyats.contrib.creators.netbox import Netbox
import yaml
import os
# Define Netbox URL, user token, and default credentials
netbox_url = os.getenv('NETBOX_URL')
user_token = os.getenv('NETBOX_USER_TOKEN')
def_user = '%ENV{DEF_PYATS_USER}'
def_pass = '%ENV{DEF_PYATS_PASS}'
url_filter = 'site=pyats-webinar'
# url_filter = 'site_id=68'
# url_filter = 'site=pyats-webinar&os=ios-xe'
# url_filter = 'platform=ios-xe'
# Create testbed object and build data structure
nb_testbed = Netbox(
    netbox_url=netbox_url,
    user_token=user_token,
    def_user=def_user,
    def_pass=def_pass,
    url_filter=url_filter,
    ssl_verify=False,
    topology=True
)
# Generate testbed file
tb = nb_testbed._generate()
tb_yaml = yaml.dump(tb)
with open("testbed.yaml", "w") as f:
    f.write(tb_yaml)
```
The resulting testbed file produced by either option will look something like this, depending on your network. Note that as we included the `--topology` switch the testbed file output includes the interfaces and connections from NetBox also:

```
devices:
  CSR1:
    alias: CSR1
    connections:
      cli:
        ip: 10.90.0.35
        protocol: ssh
    credentials:
      default:
        password: '%ENV{DEF_PYATS_PASS}'
        username: '%ENV{DEF_PYATS_USER}'
    os: iosxe
    platform: iosxe
    type: CSR1000V
  CSR2:
    alias: CSR2
    connections:
      cli:
        ip: 10.90.0.36
        protocol: ssh
    credentials:
      default:
        password: '%ENV{DEF_PYATS_PASS}'
        username: '%ENV{DEF_PYATS_USER}'
    os: iosxe
    platform: iosxe
    type: CSR1000V
testbed:
  credentials:
    default:
      password: '%ENV{DEF_PYATS_PASS}'
      username: '%ENV{DEF_PYATS_USER}'
topology:
  CSR1:
    interfaces:
      GigabitEthernet1:
        alias: CSR1_GigabitEthernet1
        ipv4: 10.90.0.35/27
        type: ethernet
      GigabitEthernet2:
        alias: CSR1_GigabitEthernet2
        ipv4: 192.168.1.1/30
        link: cable_num_34
        type: ethernet
      GigabitEthernet3:
        alias: CSR1_GigabitEthernet3
        type: ethernet
      GigabitEthernet4:
        alias: CSR1_GigabitEthernet4
        type: ethernet
      GigabitEthernet5:
        alias: CSR1_GigabitEthernet5
        type: ethernet
  CSR2:
    interfaces:
      GigabitEthernet1:
        alias: CSR2_GigabitEthernet1
        ipv4: 10.90.0.36/27
        type: ethernet
      GigabitEthernet2:
        alias: CSR2_GigabitEthernet2
        ipv4: 192.168.1.2/30
        link: cable_num_34
        type: ethernet
      GigabitEthernet3:
        alias: CSR2_GigabitEthernet3
        type: ethernet
      GigabitEthernet4:
        alias: CSR2_GigabitEthernet4
        type: ethernet
      GigabitEthernet5:
        alias: CSR2_GigabitEthernet5
        type: ethernet
```
When you run a command at the CLI of a network device, you get unstructured data back as the response, which is just a blob of text:

```
CSR1#sh ip interface brief 
Interface              IP-Address      OK? Method Status                Protocol
GigabitEthernet1       10.0.0.15       YES manual up                    up      
GigabitEthernet2       192.168.1.1     YES manual up                    up      
Loopback0              1.1.1.1         YES manual up                    up  
```
This is great for humans, as we can read this, but a computer cannot understand this data. Also if the next version of the OS you are using makes a change to the way that the output is formatted then you will have a problem and have to re-write your scripts to handle this. This is where the Genie parser comes into play as it will parse the output into structured data. The data is then represented using key/value pairs in JSON format that can be used by a computer:

