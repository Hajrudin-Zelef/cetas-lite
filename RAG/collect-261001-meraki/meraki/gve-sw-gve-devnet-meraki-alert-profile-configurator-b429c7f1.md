---
id: collect-261001-meraki/meraki/gve-sw-gve-devnet-meraki-alert-profile-configurator-b429c7f1
title: "gve-sw-gve-devnet-meraki-alert-profile-configurator-b429c7f1"
domain: meraki
role: reference
task: reference
actors: []
dates: []
keywords: ["license"]
source: docs/RAG/collect-261001-meraki/gve-sw-gve-devnet-meraki-alert-profile-configurator-b429c7f1.md
source_anchor: ""
source_lines: [1, 33]
sha256: 27d95acdecd3c484299ac70b8319ca6e4d9f924ceb755469457fb0a8471b6f3b
---

# gve-sw-gve-devnet-meraki-alert-profile-configurator-b429c7f1

python prototype leveraging the Meraki Dashboard API to manage environmental alert conditions and enables users to configure temperature, humidity, water detection, and door status alerts.
- Jorge Banegas
- Meraki MT
In order to use the Meraki API, you need to enable the API for your organization first. After enabling API access, you can generate an API key. Follow these instructions to enable API access and generate an API key:
- Login to the Meraki dashboard
- In the left-hand menu, navigate to Organization > Settings > Dashboard API access
- Click on Enable access to the Cisco Meraki Dashboard API
- Go to My Profile > API access
- Under API access, click on Generate API key
- Save the API key in a safe place. The API key will only be shown once for security purposes, so it is very important to take note of the key then. In case you lose the key, then you have to revoke the key and a generate a new key. Moreover, there is a limit of only two API keys per profile.
For more information on how to generate an API key, please click here.
Note: You can add your account as Full Organization Admin to your organizations by following the instructions here.
- 
Set up a Python virtual environment. Make sure Python 3 is installed in your environment, and if not, you may download Python here. Once Python 3 is installed in your environment, you can activate the virtual environment with the instructions found here.
- 
Install the requirements with pip3 install -r requirements.txt
- 
Add Meraki API key to the config.py python file, run python scripts/output_org.py if you need to find ORG_ID
API_KEY = ''
ORG_ID = ''
- Fill out datafiles/devices.csv and enter the list of device serials for the sensors to attach to the Alert Profile
Serial
XXXX-XXXX-XXXX
XXXX-XXXX-XXXX
XXXX-XXXX-XXXX
XXXX-XXXX-XXXX
- Edit data_files/alert_settings.json file to configure the Alert Profile settings to your choice, refer to the API Documentation if needed.
To run the program, use the command:
$ python3 scripts/main.py
Provided under Cisco Sample Code License, for details see LICENSE
Our code of conduct is available here
See our contributing guidelines here
Please note: This script is meant for demo purposes only. All tools/ scripts in this repo are released for use "AS IS" without any warranties of any kind, including, but not limited to their installation, use, or performance. Any use of these scripts and tools is at your own risk. There is no guarantee that they have been through thorough testing in a comparable environment and we are not responsible for any damage or data loss incurred with their use. You are responsible for reviewing and testing any scripts you run thoroughly before use in any non-testing environment.
