---
id: collect-261001-unifi-ubiquiti/unifi-ubiquiti/devices-unifi-captive-portal-service-guest-6a0707c5-1
title: "devices-unifi-captive-portal-service-guest-6a0707c5"
domain: unifi-ubiquiti
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-unifi-ubiquiti/devices-unifi-captive-portal-service-guest-6a0707c5.md
source_anchor: ""
source_lines: [1, 102]
sha256: 37de2479d40177562be1dc907a336d027035d75cfa68b301d2177851c44925c8
---

# devices-unifi-captive-portal-service-guest-6a0707c5

UniFi captive portal service to capture guest Emails
Capture emails and customer data through your UniFi Guest Portal.
MyPlace integrates directly with the UniFi Network Application and lets you launch a captive portal in minutes. No extra hardware or configuration needed.
Trusted by more than 500+ companies across the globe
No Extra Hardware
Integrate directly with your UniFi platform using the native API. No need for extra hardware or flashing of access points.
Quick and Easy Install
Activate your UniFi Guest Portal in minutes by connecting with your UniFi controller. No site visits required.
Grow with UniFi
Grow and scale your business with dedicated UniFi software. Get the most from your UniFi network.
How UniFi captive portal service works
Controller Hostname
Identify controller or network application hostname. For UDM’s and cloud keys, MyPlace will generate a hostname automatically.
Create User
Create a new admin user for your UniFi controller. Limit user access to the site you want to connect to MyPlace.
Connect to MyPlace
Connect to MyPlace using your controller hostname and user credentials. After connecting, MyPlace will configure your guest network via the UniFi API.
UniFi captive portal features
UDM & Cloud Key
Instant integration with UniFi dream machines cloud keys. Free hostname service included.
Multi Site
Connect any number of network applications including combinations of UDM/cloud key and self hosted controllers.
Segmentation
Segment user data by SSID and access point. Multiple venues on single UniFi site.
UniFi Hosting
Free hosting for small deployments and affordable hosting with a hosting partner for larger deployments.
Secure
SSL encryption and data protection. Password secrets manager for additional security.
Access Limits
Define guest WiFi access limits by time and bandwidth. Take control of your UniFi guest network.
Data Compliance
Compliant with all major data legislations. Customizable terms and conditions.
Easy Setup
Connect your UniFi network application quickly and easily. Launch UniFi guest portal service in minutes.
UniFi captive portal marketing service
The UniFi guest portal integration is built to power our WiFi Hotspot CRM. Capturing customer data and providing visitor analytics.
UniFi captive portal for all networks
MyPlace UniFi captive portal service works on all UniFi network setups. Hostname and SSL certs generated if required for a seamless guest WiFi experience.
UniFi Dream Machine
UniFi Captive Portal Service works with UDM and UDM Pro consoles. Automated hostname generation by MyPlace provides for API access.
Self Hosted Controller
UniFi controllers that are self hosted can be easily integrated. Launch a captive portal in minutes with self hosted UniFi controllers.
UniFi Access Point
Connect any number of network applications including combinations of UDM/cloud key and self hosted controllers.
UniFi Cloud Key
Capture email on the UniFi Cloud Key. Hostname and SSL certs are generated automatically by MyPlace for API access.
UniFi Cloud Controller
The new UniFi cloud controller service is ready to go with MyPlace. Get all the benefits of the UniFi captive portal on the UniFi cloud service.
Official UniFi API
Connect your UniFi network to MyPlace with the new official UniFi API. Seamless connection for cloud accounts
UniFi captive portal setup
Setting up the UniFi Captive Portal is simple.
You just need the UniFi Network Application (previously called the controller) to be online.
If the application does not have a hostname with an SSL certificate, you will need to generate one.
Once you have a valid FQDN hostname and local user credentials, you can set up the UniFi Captive Portal to start capturing guest emails.
Have questions? Schedule a demo with one of our UniFi experts.
No UniFi access point?
No UniFi or WiFi Access Point? No worries. Sign up for our special offer and get a free UniFi Access Point.
How to setup the UniFi captive portal for your guests
Looking to offer hassle-free guest WiFi? The UniFi Captive Portal is a smart choice.
Instead of adding another SSID with a basic password, use the guest portal to keep your main network secure. It helps you isolate guest traffic and set time limits for access.
Configuration through the UniFi Network Application is straightforward. You can enable guest isolation to block access to internal resources without needing VLANs.
You can also use the captive portal as a marketing tool by displaying promotions, offers, or brand messaging during login.
Create a guest WiFi SSID
If you don’t have one already you will need to create an SSID for guests. You can do this by adding a new WiFi network as per below:
- Go to Settings > WiFi
- Create New
- Give the SSID a name, that will be recognisable as a guest WiFi service
- Leave password blank to encourage users to join
- Select Manual in advanced section
- Check Hotspot Portal
- Security Protocol should be set to open
- Select Add WiFi Network
Configuring the UniFi captive portal
Once the guest SSID is set up correctly, the next step is to configure the UniFi captive portal, as per the instructions below:
- Go to Hotspot manager
- Select Landing Page tab
- If prompted, enable landing page
- Got to authentication tab on right hand side
- In One Way Methods section, select external portal server
- Enter IP address of server and save
- Go to Settings tab
- Check Show Landing Page
- Check domain and enter hostname of the redirect url
- Adjust any pre or post authorization exceptions
- Save
The above instructions are for configuring the external portal server as the UniFi Captive Portal. To use other authentication adjust in the authentication tab.
Advanced UniFi captive portal healthchecks
The UniFi Guest Portal service works really well as an away to capture guest email and more for busy customer facing venues. However like any business critical service, you need to run regular health checks and status updates so that the WiFi users have a perfect experience every time. That is why we developed our advanced UniFi Guest Portal health check system, so that we can be sure that your integration is set up correctly.
List of Health Check Tests
The UniFi Guest Portal uses the UniFi API which uses the url string for authentication. We need to verify that the site id embedded in the url is unique and not a duplicate or the UniFi “default” ID.
MyPlace verifies that there is an SSID available and that the SSID has the appropriate guest policy applied
In addition to UniFi controller ports, MyPlace also runs tests to verify. that the ports are not restricted on the local firewall
Myplace detects if an active vlan is in place. Where vlans are in place then additional testing is required to make sure that the appropriate firewall rules are in place. In almost all cases vlan tagged networks work perfectly fine without additional amendments
We run a login test to verify that the controller credentials are correct and have not been changed
Where a hostname is used, MyPlace runs a background FQDN test
MyPlace run a series of tests to verify that access via a series of ports are open to the controller. Alerts are issued when port access is lost or restricted
Sometimes the UniFi controller is accessible via a FQDN hostname. When this is the case, MyPlace run some tests to verify that the hostname has a valid SSL cert applied
Frequently asked questions
Can’t find what you’re looking for? Contact us we’re ready to help you.
You need to create a new site specific controller user. MyPlace then uses these credentials to communicate with the controller with the UniFi API.
UniFi’s “One Way Method” only redirects guests to your external captive portal; your system collects data and then calls UniFi’s API to grant internet access. UniFi doesn’t validate credentials or show its own login page.
