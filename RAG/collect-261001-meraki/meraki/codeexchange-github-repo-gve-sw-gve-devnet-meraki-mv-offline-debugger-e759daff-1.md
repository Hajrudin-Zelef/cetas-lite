---
id: collect-261001-meraki/meraki/codeexchange-github-repo-gve-sw-gve-devnet-meraki-mv-offline-debugger-e759daff-1
title: "config.py"
domain: meraki
role: reference
task: reference
actors: []
dates: []
keywords: ["license", "parameters"]
source: docs/RAG/collect-261001-meraki/codeexchange-github-repo-gve-sw-gve-devnet-meraki-mv-offline-debugger-e759daff.md
source_anchor: ""
source_lines: [1, 69]
sha256: a60c3d811ea945e9f84b9895b3e2416f777c3db25a8e4c1a7e2c30c9678441fe
---

# config.py

This flask app serves as an automatic troubleshooting and ticketing platform for Meraki MV Cameras connected to MS Switches. The app primarily listens for Meraki Camera Offline Webhook Alerts, executes the MV troubleshooting algorithm, and generates a ServiceNow ticket if the troubleshooting algorithm is unsuccessful in bringing up the MV camera.
The troubleshooting algorithm consists of the steps:
Delay 5 minutes -> Check Status
The Delay is configurable and set at 5 minutes by default. To modify the Delay parameter adjust the global variable at the top of app.py
DELAY_TIME = 5  # time to wait between polls in minutes
This app also listens for additional Webhooks:
If an alert is generated for an MX or MS device which affects an MV Camera, the app intelligently creates a ticket for the MX or MS with a list of affected cameras. This reduces the number of tickets and introduces network level intelligence.
Note:
This code works best with the topology MX -> MS -> MV. If the environment doesn't have this topology, there may be issues creating tickets. The intelligent ticket creation depends on the order of alerts and this topology. For ex: an MX alert should come in first before the affected MS and MV devices. The order is usually right, but it can be unpredictable. To guarantee correct behavior, see the Meraki Alerts section.
In order to use the Meraki API, you need to enable the API for your organization first. After enabling API access, you can generate an API key. Follow these instructions to enable API access and generate an API key:
Organization > Settings > Dashboard API accessEnable access to the Cisco Meraki Dashboard APIMy Profile > API accessGenerate API key
For more information on how to generate an API key, please click here.
Note: You can add your account as Full Organization Admin to your organizations by following the instructions here.
The webhook receivers for Meraki Alerts require https URLs. Since this code runs on http://localhost:80 by default, it requires a forwarding address with https in order to receive the webhooks. To emulate this, ngrok is used.
Follow these instructions to set up ngrok:
Getting Started > Your Authtoken on the menu on the left-hand side. Copy the token on this page.ngrok authtoken [yourtoken]
replacing [yourtoken] with the authtoken you copied in Step 2.
ngrok http 80
Note: It's recommend to use port 443 for secure communication (if not using Ngrok). Refer to this guide to see how to enable SSL on Nginx (in docker) using Let's Encrypt. Change the relevant sections of docker-compose.yml and the nginx/conf/default.conf based on the guide.
In order to generate webhooks for the Meraki Alerts this code is written to create tickets for, follow these steps:
Network-wide >  Configure > Alerts.A camera goes offline for [X] minutes. Under Security Appliance, select A security appliance goes offline for [X] minutes. Under Switch, select A switch goes offline for [X] minutes. To guarantee that the alerts are sent in the right order, you can modify the amount of time that should pass before an alert is generated. For instance, keep the security appliance at 5 minutes, but increase the time of the switch to 10 minutes and the camera to 15 minutes. Thus, if a router goes down, then after 5 minutes, a ticket will be created for it once the alert is generated but it will take 5 more minutes before the switch connected to the router will generate an alert.Webhooks section, add an HTTP receiver. Give it a name and copy the https URL for the flask app as the receiver URL followed by /alerts. Use ngrok webhook receiver if running locally followed by Alerts Settings section, add the webhook receiver that you just created to the Default recipients list. Add  shared secret as well, note this for later.Save to save these changes.
For more information about Meraki Alerts and Notifications, please click here.
To support scalability, this app uses several technologies (Celery, Redis, etc.) that are easily tied together with Docker. Docker is the recommended deployment method. Install Docker here.
Gather the url and username/password for a ServiceNow Admin associated with the target instance. A free developer instance can be created with the Developer Program
Default Ticket Parameters:
Default description, impact, urgencies, etc. are set in the method create_service_now_ticket in app.py. Please refer to this method to make changes to default values.
Ticket Clean Up Feature:
This feature marks a ticket as 'Resolved' if the device is active 'X' hours after the creation of a ticket that's still unresolved.
The default checking period is 1 hour, but this can be configured with the global parameter TICKET_REMOVAL_TIME at the top of app.py
TICKET_REMOVAL_TIME = 1
Duplicate Ticket Prevention Feature:
This feature stops duplicate tickets from being created for an MV Camera with an existing ticket. This is useful to avoid multiple ServiceNow tickets for multiple webhooks alerts on a single camera (ex: Critical Hardware Failure and Offline Alert)
git clone [repository name].env_sample file to .env. This file holds sensitive environment variables which will be passed to the docker container securely. Rename config_sample.py to config.py (located at: flask_app/config_sample.py).flask_app/config.py) to environment variables. The CSV File is a backup of ticket information for manual ServiceNow ticket creation.
# config.py
TICKET_CSV_PATH = ""
# .env
MERAKI_API_KEY="<replace>"
SHARED_SECRET="<replace>"
# Example: MERAKI_API_KEY="1234567890"
# config.py
SERVICE_NOW_ENABLED = True
# .env
SERVICENOW_INSTANCE="<replace>"
SERVICENOW_USERNAME="<replace>"
SERVICENOW_PASSWORD="<replace>"
# Example: SERVICENOW_INSTANCE="https://example.com"
TICKET_CLEANUP = False
DUPLICATE_TICKETS = True
pip3 install -r flask_app/requirements.txt
To run the program, use the docker command:
$ docker-compose up -d --build
This will build the containers in the background and launch them. If you have docker desktop, you should see (assuming no errors):
The docker-compose file runs 3 main python scripts:
db.py
This script creates the database that will hold the connection and status information of the routers, switches, and cameras in the network (critical for intelligent ticketing).
Once this code completes, it will print out three empty lists representing the items in the routers, switches, and cameras tables. It will also create the database file sqlite.db.
populate.py
The next step is to populate the database with the devices in your Meraki network, which is done with this script.
Once the code completes, it will print out the contents of the routers, switches, and camera tables in the database representing (Source Serial, Connected Device Serial, Source Status)
Note: this code is written with the assumption that the switches are connected to routers and the access points are connected to switches. If this is not the case in your environment, it will likely generate an error. To remedy this, add more conditions in the file. For example, there is a condition to check if the device types connected to one another are an 'Camera' and 'switch'. Suppose you have switches connected to switches in your environment, you will need to add a condition to the code that checks if the device types are both 'switch' and then write the code to handle this condition.
app.py (flask run)
The main web app will start and begin listening on default port 5000.
Once a webhook is received, the contents will be printed out:
A down camera will trigger the troubleshooting automation. The results of the automation are written to a log file in flask_app/logs/[alerting serial].log:
A Ticket is created in ServiceNow if the troubleshooting process is unable to revive the down camera.
The ticket content is written to a CSV File specified in config.py as well.
Provided under Cisco Sample Code License, for details see LICENSE
Our code of conduct is available here
See our contributing guidelines here
