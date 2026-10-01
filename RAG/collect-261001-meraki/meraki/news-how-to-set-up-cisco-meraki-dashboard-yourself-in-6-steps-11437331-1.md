---
id: collect-261001-meraki/meraki/news-how-to-set-up-cisco-meraki-dashboard-yourself-in-6-steps-11437331-1
title: "news-how-to-set-up-cisco-meraki-dashboard-yourself-in-6-steps-11437331"
domain: meraki
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-meraki/news-how-to-set-up-cisco-meraki-dashboard-yourself-in-6-steps-11437331.md
source_anchor: ""
source_lines: [1, 102]
sha256: ec71f9af2bf493e0d993c4a26995465730629ca5dc18688f225bd044c3f971a8
---

# news-how-to-set-up-cisco-meraki-dashboard-yourself-in-6-steps-11437331

Cisco Meraki’s dashboard for managing your entire end-to-end network infrastructure in one platform (from anywhere) is the solution you’ve been seeking. Meraki creates a variety of incredible products for companies of all sizes. Whether you’re with a global enterprise looking for a solution that can manage a massive network or a small business that needs a scalable option, Meraki devices are ideal.
Cisco Meraki’s configuration process is user-friendly but can be intimidating for new users. So, should you do it yourself? Or should you call in a professional? We’ll review, step-by-step, how to proceed with your dashboard configuration. Then you can decide whether you’re going to DIY the task or let Stratus Information Services take care of the heavy lifting.
What Is Cisco Meraki Used For?
Before we get into the Cisco Meraki configuration guide, let’s briefly discuss what Cisco Meraki products are used for. As a tech professional with a busy schedule, you recognize that the decision to invest in Meraki products comes with substantial benefits. These are top-of-the-line devices that can transform the way your company handles:
- 
IT issues
- 
Network security
- 
Network visibility
- 
Mobile device management
Cisco Meraki offers several main device types, including MX security appliances, MS switches, and MR wireless access points. Meraki networks can include multiple device types, allowing for flexible and scalable network deployments.
Essentially, you can think of Meraki products as the unification of your network management system. They can control thousands of mobile and desktop devices through one ultra-convenient, secure dashboard. Meraki networks act as logical containers for multiple devices, making management and permission settings more streamlined and efficient.
Cisco Meraki provides a centralized cloud management platform for all Meraki devices and services. Its solutions enhance network security through cloud-managed firewalls and access points.
Can I Configure Cisco Meraki Products By Myself?
After receiving their shiny new hardware, most people start wondering how to configure their Cisco Meraki products. A common question that comes up is whether or not people can fly solo without expert help.
The short answer is: Yes, you can, but you don’t have to.
For IT professionals, the process might be a breeze. For those that feel less equipped to take this project on, start by following our six-step guide, and if you get stuck, reach out to us here at Stratus Information Systems. Our team is on standby, ready to jump in and help you untangle the cords.
Setting Up the Meraki Cloud
Setting up the Meraki cloud is the foundation for managing your Cisco Meraki devices from anywhere in the world. The Meraki cloud leverages a centralized, web-based Meraki dashboard, giving users the power to configure, monitor, and manage all their Meraki devices through a single dashboard account. Once you’ve created your Meraki dashboard account and added your devices, you’ll have access to a comprehensive suite of configuration settings and network details, all in one place.
The Meraki dashboard is designed for simplicity and efficiency. From this interface, users can easily configure device settings, monitor network health, and access advanced features like mobile device management and endpoint management through the systems manager network. Whether you’re managing a single site or a global network, the Meraki cloud ensures you can access and control your network infrastructure securely and remotely. With features like real-time alerts, device inventory, and detailed reporting, the Meraki dashboard empowers users to manage their networks proactively and efficiently—no matter where they are.
6 Steps for Setting Up Your Cisco Meraki Dashboard
As Cisco itself boasts, “There is a tremendous amount of flexibility with the initial setup for a Meraki deployment.” This is good news because “you can configure everything before you even have your devices, thanks to the Meraki cloud.” For multi-site deployments, Configuration Templates can be used to push consistent settings across networks.
Before you begin working your way through the setup, here are other topics you should thoroughly understand first:
- 
Upstream Firewall Rules for Cloud Connectivity
- 
Meraki Cloud Architecture
- 
Meraki Dashboard Organizational Structure
- 
Combined Dashboard Networks
- 
Systems Manager FAQ
- 
MS Warm Sparew)
- 
Cisco Meraki Licensing Guidelines and Limitations
- 
Meraki documentation: Refer to official Meraki documentation for authoritative configuration steps and best practices.
Here are the six steps to set up your Cisco Meraki dashboard correctly.
1. Gather Your Information for Pre-Setup
This is what you’ll need on hand for setting up your dashboard account, network, and devices:
- 
An order number for your Meraki purchase OR the serial #s of your Meraki devices
- 
A plan for how you will group your devices (as well as what you are going to do with them)
- 
An internet uplink for your devices 
Note: The only prerequisite to set up a Meraki device is an uplink connection on the device itself.
- 
A valid firewall with configured rules
When deploying Cisco Meraki devices, follow the recommended deployment hierarchy: install the Router (MX) first, then the Switch (MS), and finally the Access Point (MR).
When creating a network within your organization, be sure to select the appropriate device types—MX for routers, MS for switches, and MR for access points—to match your deployment needs.
Physical installation requires mounting the hardware and connecting it to the internet.
To create a dashboard account, register at dashboard.meraki.com.
2. Create a Meraki Dashboard Account
If you don’t already have an account, it’s time to create one. Here’s how:
- 
Navigate to Meraki’s Dashboard Login.
- 
Click “Create an Account.”
- 
You can sign in with your Cisco SSO or create a free account to access the Meraki dashboard.
- 
Choose your Meraki dashboard and organization region. This can’t be changed, so make sure it’s correct!
- 
Enter in your:
- 
Email address for login and communication
- 
Full name that will be displayed on your account
- 
Password that’s 8+ characters long and a combination of upper- and lower-case letters, numbers, and special characters
- 
Company or organization name
- 
Address for default network locations and maps (optional)
- 
Click “Create account.”
- 
Verify your account by checking your email and following the confirmation link.
Once you’re in and your dashboard is set up, you’ll have access to immense visibility and a wealth of features. This intuitive, interactive web interface lets you see reports that keep you updated on your devices’ health, network usage, client information, and more.
Looking for an overview of your hardware and network features? How about the next steps in the setup process? Fire up your Meraki dashboard to access this information and much more.
3. Create a Network
A device network should be created for each physical location. In Cisco Meraki, a network is a logical container for multiple devices, allowing for streamlined management and configuration. Organizations are collections of networks, and user roles with specific permissions are required to manage them. To manage organizations, users need the Meraki Organization (View) permission, and to manage networks, they need the Meraki Networks (View and Manage) permission.
- 
Log into your Meraki dashboard account.
- 
Click “Register Meraki devices” and “Next.”
- 
Input the following information: 
  - 
Name to be used to identify the network
  - 
Network type (select the appropriate device types for your deployment, such as MX for routers, MS for switches, and MR for access points)
  - 
Devices (optional)
- 
- 
Click “Create network.”
