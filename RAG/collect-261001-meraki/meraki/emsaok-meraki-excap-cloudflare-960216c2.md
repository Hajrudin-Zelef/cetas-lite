---
id: collect-261001-meraki/meraki/emsaok-meraki-excap-cloudflare-960216c2
title: "emsaok-meraki-excap-cloudflare-960216c2"
domain: meraki
role: reference
task: reference
actors: []
dates: []
keywords: ["apache", "license"]
source: docs/RAG/collect-261001-meraki/emsaok-meraki-excap-cloudflare-960216c2.md
source_anchor: ""
source_lines: [1, 47]
sha256: 0807a3cf1e5f8a43d3eddbf1275aa6fade3b3d90b3ab93fc3d8840704104a95d
---

# emsaok-meraki-excap-cloudflare-960216c2

Cloudflare: WiFi Captive Portal
An Open-Source Meraki Captive Portal using Cloudflare Pages.
The Meraki Captive Portal with Cloudflare is an open-source and freely accessible solution for creating customized Captive Portals tailored to Cisco Meraki wireless networks.
Utilizing Cloudflare Pages and the open-source framework of this project, you can effortlessly redirect users to your customized Splash Page, interact with them, and manage internet access as necessary.
Thanks to Cloudflare's reliable usage models, there are no associated fees, subscription costs, or ongoing maintenance requirements, making this solution equally accessible to individuals, small businesses, and organizations.
Check out what it looks like for end users! Live Demo: https://meraki-excap-cloudflare.pages.dev
Want to try it out in your environment? Configure the following link as your WiFi Splash-Page for an access point or ssid, and you are good to go! https://meraki-excap-cloudflare.pages.dev.
Prerequisite: Create a Cloudflare account www.cloudflare.com.
Prerequisite: Create a GitHub account www.github.com.
- Fork this repository to your GitHub.
- Rename the forked repo for your desired live page URL.
- Log in to Cloudflare and navigate to "Workers & Pages" on the sidebar.
- Click "Create Application," then select the "Pages" tab.
- Click "Connect to Git" and choose your forked repository.
- Click "Begin setup" and then "Save and Deploy."
- Wait about 5 minutes for the initial build.
- Success! Check it out at https://<forked-repo-name>.pages.dev .
Prerequisite: Have access to your orgqnizations Meraki Dashboard (https://meraki.com/dashboard).
- Go to Meraki Dashboard.
- Navigate to "Wireless" > "Configure" > "Splash Page"
- Under "Custom splash URL" enter the splash page url.
- Navigate to "Wireless" > "Configure" > "Access Control"
- Under "Security" select "Open (no encryption)"
- Under "Splash page" select "Click-through"
- Under "Advanced splash settings" select:
  - Captive portal strength: Block all access until sign-on is complete
  - Walled garden: Enabled
- Under "Walled garden ranges" enter:
  - *.wifi.direct
  - *.google-analytics.com
  - forked-repo-name.pages.dev
  - fonts.googleapis.com
Setting up a custom domain requires the domain to be added to Cloudflare under the "Websites" tab, as this process will add a new CNAME attribute to the domain's DNS configurations.
Make sure to update the Walled Garden ranges if previosuly configured! Both the cloudflare pages url and the custom domain name will need to be added!
- Select the "Workers & Pages" tab on the sidebar.
- Select the pages project you want to set up a custom domain for.
- Select the "Custom domains" tab in the main section.
- Click the "Setup a custom domain" button.
- Enter the domain you want to configure as the custom domain name.
- Click the "Activate Domain" button at the bottom of the page.
- Success! Check out your new domain in a few minutes at https://<custom-domain> .
Updating your deployment is as easy as it comes!
Simply commit your changes to the repository used to create the page, and Cloudflare will automatically build and deploy the next version for you! This can be altered, though, to only automatically deploy to a preview branch rather than production.
Adam Spera & Brey Rivera
Find us at the links below!
https://adamspera.dev & https://breyr.dev
Apache 2.0 (see LICENSE).
