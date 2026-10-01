---
id: collect-261001-meraki/meraki/gve-sw-gve-devnet-meraki-mv-offline-debugger-9e079a76-1
title: ".env"
domain: meraki
role: reference
task: reference
actors: []
dates: []
keywords: ["parameters"]
source: docs/RAG/collect-261001-meraki/gve-sw-gve-devnet-meraki-mv-offline-debugger-9e079a76.md
source_anchor: ""
source_lines: [1, 86]
sha256: a759eb9c2e02b34f5a6f9d63a2d761566c4d65e0c793c8c7c1b1288a1158eb51
---

# .env

This flask app serves as an automatic troubleshooting and ticketing platform for Meraki MV Cameras connected to MS Switches. The app primarily listens for Meraki Camera Offline Webhook Alerts, executes the MV troubleshooting algorithm, and generates a ServiceNow ticket if the troubleshooting algorithm is unsuccessful in bringing up the MV camera.
The troubleshooting algorithm consists of the steps:
- Receive Camera Down Alert
- Delay 5 minutes -> Check Status
- Cycle Connected Switch port
- Delay 5 minutes -> Check Status
- Generate Ticket
The Delay is configurable and set at 5 minutes by default. To modify the Delay parameter adjust the global variable at the top of app.py
DELAY_TIME = 5  # time to wait between polls in minutes
This app also listens for additional Webhooks:
- MS Switch Offline
- MX Offline
- MV Critical Hardware Alert
If an alert is generated for an MX or MS device which affects an MV Camera, the app intelligently creates a ticket for the MX or MS with a list of affected cameras. This reduces the number of tickets and introduces network level intelligence.
Note:
This code works best with the topology MX -> MS -> MV. If the environment doesn't have this topology, there may be issues creating tickets. The intelligent ticket creation depends on the order of alerts and this topology. For ex: an MX alert should come in first before the affected MS and MV devices. The order is usually right, but it can be unpredictable. To guarantee correct behavior, see the Meraki Alerts section.
The code supports the MX, MS, and/or MV being in separate networks as well as separate organizations.
- Trevor Maco
- Python 3.x
- Ngrok
- Meraki MV
- Meraki Webhooks, APIs
- Docker
- Flask
- SQLite
- Redis
- Celery
- Nginx
- ServiceNow
In order to use the Meraki API, you need to enable the API for your organization first. After enabling API access, you can generate an API key. Follow these instructions to enable API access and generate an API key:
- Login to the Meraki dashboard
- In the left-hand menu, navigate to Organization > Settings > Dashboard API access
- Click on Enable access to the Cisco Meraki Dashboard API
- Go to My Profile > API access
- Under API access, click on Generate API key
- Save the API key in a safe place. The API key will only be shown once for security purposes, so it is very important to take note of the key then. In case you lose the key, then you have to revoke the key and a generate a new key. Moreover, there is a limit of only two API keys per profile.
For more information on how to generate an API key, please click here.
Note: You can add your account as Full Organization Admin to your organizations by following the instructions here.
The webhook receivers for Meraki Alerts require https URLs. Since this code runs on http://localhost:80 by default, it requires a forwarding address with https in order to receive the webhooks. To emulate this, ngrok is used. Follow these instructions to set up ngrok:
- Create a free account or login to Ngrok.
- Retrieve your auth token by navigating to Getting Started >Your Authtoken on the menu on the left-hand side. Copy the token on this page.
- Then install the client library depending on your OS here.
- Once you have ngrok installed, update the ngrok configuration file with your auth token by running the following command on the terminal/command prompt:
ngrok authtoken [yourtoken]
replacing [yourtoken] with the authtoken you copied in Step 2.
- Start the ngrok tunnel for port 80 with the command:
ngrok http 80
Note: It's recommend to use port 443 for secure communication (if not using Ngrok). Refer to this guide to see how to enable SSL on Nginx (in docker) using Let's Encrypt. Change the relevant sections of docker-compose.yml and the nginx/conf/default.conf based on the guide.
In order to generate webhooks for the Meraki Alerts this code is written to create tickets for, follow these steps:
- Login to the Meraki dashboard.
- In the left-hand menu, navigate to Network-wide >Configure >Alerts .
- Here, select the appropriate checkboxes for which events should generate alerts. Under Cameras, select A camera goes offline for [X] minutes . Under Security Appliance, selectA security appliance goes offline for [X] minutes . Under Switch, selectA switch goes offline for [X] minutes . To guarantee that the alerts are sent in the right order, you can modify the amount of time that should pass before an alert is generated. For instance, keep the security appliance at 5 minutes, but increase the time of the switch to 10 minutes and the camera to 15 minutes. Thus, if a router goes down, then after 5 minutes, a ticket will be created for it once the alert is generated but it will take 5 more minutes before the switch connected to the router will generate an alert.
- Next, you need to add the webhook receiver to send the alert webhooks to. Under the Webhooks section, add an HTTP receiver. Give it a name and copy the https URL for the flask app as the receiver URL followed by/alerts . Usengrok webhook receiver if running locally followed by/alerts .
- Then, add the webhook receiver as a default recipient of alerts. Scroll to the top of the page to the Alerts Settings section, add the webhook receiver that you just created to the Default recipients list. Addshared secret as well, note this for later.
- Scroll to the bottom of the page and click Save to save these changes.
For more information about Meraki Alerts and Notifications, please click here.
To support scalability, this app uses several technologies (Celery, Redis, etc.) that are easily tied together with Docker. Docker is the recommended deployment method. Install Docker here.
Gather the url and username/password for a ServiceNow Admin associated with the target instance. A free developer instance can be created with the Developer Program
- Default Ticket Parameters :
Default description, impact, urgencies, etc. are set in the method create_service_now_ticket in app.py. Please refer to this method to make changes to default values.
- Ticket Clean Up Feature :
This feature marks a ticket as 'Resolved' if the device is active 'X' hours after the creation of a ticket that's still unresolved.
The default checking period is 1 hour, but this can be configured with the global parameter TICKET_REMOVAL_TIME at the top of app.py
TICKET_REMOVAL_TIME = 1
- Duplicate Ticket Prevention Feature :
This feature stops duplicate tickets from being created for an MV Camera with an existing ticket. This is useful to avoid multiple ServiceNow tickets for multiple webhooks alerts on a single camera (ex: Critical Hardware Failure and Offline Alert)
- Clone this repository with git clone [repository name]
- Rename the .env_sample file to.env . This file holds sensitive environment variables which will be passed to the docker container securely. Renameconfig_sample.py toconfig.py (located at:flask_app/config_sample.py ).
- Add Meraki API key and Webhook Shared Secret  (.env ).
# .env
MERAKI_API_KEY="<replace>"
SHARED_SECRET="<replace>"
# Example: MERAKI_API_KEY="1234567890"
- Enable ServiceNow (Default = True) (flask_app/config.py ), add ServiceNow instance variables (.env ) obtained from the prerequisites section.
# config.py
SERVICE_NOW_ENABLED = True# .env
SERVICENOW_INSTANCE="<replace>"
SERVICENOW_USERNAME="<replace>"
SERVICENOW_PASSWORD="<replace>"
# Example: SERVICENOW_INSTANCE="https://example.com"
- Enable or Disable the ServiceNow 'Ticket Cleanup' and/or the 'Duplicate Ticket Prevention' feature (flask_app/config.py ).
TICKET_CLEANUP = False
DUPLICATE_TICKETS = True
- Set up a Python virtual environment. Make sure Python 3 is installed in your environment, and if not, you may download Python here. Once Python 3 is installed in your environment, you can activate the virtual environment with the instructions found here.
- Install the requirements with pip3 install -r flask_app/requirements.txt
To run the program, use the docker command:
