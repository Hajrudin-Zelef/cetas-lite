---
id: collect-261001-unifi-ubiquiti/unifi-ubiquiti/hc-en-us-articles-27061316960663-recovering-and-troubleshooting-unifi-talk-insta-6e8d3620
title: "hc-en-us-articles-27061316960663-recovering-and-troubleshooting-unifi-talk-insta-6e8d3620"
domain: unifi-ubiquiti
role: reference
task: reference
actors: ["United States"]
dates: []
keywords: []
source: docs/RAG/collect-261001-unifi-ubiquiti/hc-en-us-articles-27061316960663-recovering-and-troubleshooting-unifi-talk-insta-6e8d3620.md
source_anchor: ""
source_lines: [1, 40]
sha256: bcf6308115731213a2163f22311b5fbd71d0152b5645dfda99eadaf41d0457d0
---

# hc-en-us-articles-27061316960663-recovering-and-troubleshooting-unifi-talk-insta-6e8d3620

Recovering and Troubleshooting UniFi Talk Installations
Whether you're factory resetting, replacing a UniFi Console, or migrating to a new one, UniFi Talk provides a seamless recovery option to restore your subscriptions, phone numbers, and settings. This guide will walk you through the process of recovering your previous installation, managing phone number assignments, and resolving common issues such as call failures or device connection errors.
Recovering a UniFi Talk Installation
If you need to recover your UniFi Talk setup on a new or reset UniFi Console, follow these steps:
- Log in to your Ubiquiti account on the Site Manager and select the UniFi Console you want to recover or migrate your Talk subscriptions to.
- Launch the UniFi Talk Setup Wizard.
  - If you have multiple UniFi Talk deployments, a list will appear. Hover over the information icon to view phone numbers linked to each deployment.
  - Select the deployment with the phone numbers you wish to recover or migrate.
- Click Next to proceed.
- On the Setup Device(s) page, assign the recovered or migrated phone numbers to users and devices. These will be available in the Number/Area Code drop-down menu. Make your selections, then click Next.
- Complete the setup to finalize the recovery or migration of your UniFi Talk installation.
Note: UniFi Talk installations can only be active on a single UniFi Console. Recovering or migrating an installation will transfer it from any previous Console.
Troubleshooting Common Issues
Unable to Receive Incoming Calls
If you're experiencing issues with receiving calls, enabling Advanced Call Routing may resolve the problem:
- In the UniFi Talk application, navigate to Settings > System.
- Enable Advanced Call Routing.
Note: After enabling, restart the Talk application for the changes to take effect.
Unable to Make Outgoing Calls
For outgoing call failures, enable Advanced Call Routing by following the same steps outlined above. If this does not resolve the issue, try the following:
- Disable SIP ALG (Session Initiation Protocol Application Layer Gateway) on the router upstream from your UniFi Console. This setting is often enabled by default on ISP-provided routers and can interfere with telephony.
Intermittent Call Failures After Network Outages
If you were previously able to make or receive calls but are encountering issues after a network event, restarting the UniFi Talk application can often resolve this:
- Go to Settings > Control Plane and select the Talk application.
- Click Stop to halt the Talk application.
- After the application stops, click Start Talk to restart it.
If the issue persists, contact UniFi Technical Support.
Outgoing Calls Labeled as Spam or Blocked
To avoid nuisance labels like "Spam" or "Scam Likely" and prevent call blocking, follow these steps:
- Set a Caller ID Name (CNAM): For business customers, add your business name (up to 15 characters) from Settings > Organization in the Talk application.
- Enroll in STIR/SHAKEN Trusted Calling: Available to verified businesses, this certification enhances call legitimacy. Initiate enrollment from Settings > Organization.
- Register Your Phone Number: Use the free registry service to submit affected phone numbers to analytics providers, supporting major US carriers.
Talk Device Showing a Connection Error
If your Talk device displays a connection error, it is unable to communicate with the Talk application. Follow these troubleshooting steps:
- Confirm that the Talk application is running:
  - Go to unifi.ui.com, select your UniFi Console, and navigate to Settings > Updates.
  - Locate the Talk application tile. If it is stopped, click the three-dot menu and select Start.
- Restart the Talk application by following the steps in the previous section.
- Restart the UniFi Console:
  - Go to Settings > Advanced and click Restart Console under the Console Controls header.
