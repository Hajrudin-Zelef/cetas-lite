---
id: collect-261001-unifi-ubiquiti/unifi-ubiquiti/home-network-unifi-teleport-ea746d95-1
title: "home-network-unifi-teleport-ea746d95"
domain: unifi-ubiquiti
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-unifi-ubiquiti/home-network-unifi-teleport-ea746d95.md
source_anchor: ""
source_lines: [1, 62]
sha256: 6109e9dc9facac5525febf50d75e2dbe82aaf1937b5504d2255941f112d452da
---

# home-network-unifi-teleport-ea746d95

Teleport was originally released in 2018 for the AmpliFi product line of Ubiquiti. But it’s now also available on all UniFi Cloud Gateways and Next-Gen gateways. It allows you to create a VPN connection with one click from your mobile device or desktop to your home network.
With a traditional VPN, you will need to configure your network, maybe open ports, create a username and password, etc, before you can make a VPN connection. With UniFi Teleport, you only need to create an invitation link in your controller.
In this article
In this article, I will explain what you need for UniFi Teleport and how to use it.
What is UniFi Teleport
UniFi Teleport allows you to make a VPN connection to your own network with one click. It uses the WireGuard VPN protocol, which is commonly used by large VPN providers, like NordVPN or Surfshark.
The difference compared to these VPN providers is that with teleport you create a VPN tunnel to your own network. This is ideal when you are on a public wireless network and want to securely access your bank account or other sensitive information.
With Teleport you can not only safely browse the internet, but you can also access your home network. After you have made the VPN connection you can access all your home network devices just like when you are connected to your wireless network at home.
Connecting to Teleport
There are two ways you can connect a device with UniFi Teleport. We can use an invitation link, that we can generate in UniFi Network, or use the single-sign-on method in the WiFiman app from Ubiquiti.
The invitation link is a unique link that is only valid for 24 hours. The link can only be used by one client device. So you need to create an invitation link for each device that you want to give access to. The VPN tunnel is stored on the device after accepting the link, allowing you to use the VPN connection at any moment that you want through the Wifiman app.
Another option to connect to UniFi Teleport is to sign in with your Ubiquiti account. For this to work you need to be a site-admin and Teleport has to be enabled. If you want to offer Teleport to multiple users, then UniFi Identity might also be a good option.
Requirements for UniFi Teleport
All UniFi Cloud Gateway consoles and Next-Gen Gateways support UniFi Teleport. Other requirements for Teleport are:
- UniFi network 7.1 or later
- Remote Access enabled in UniFi OS
- WiFiman mobile or desktop app
Remote Access
Remote access to the UniFi console must be enabled to use Teleport. You can enable remote access in the Control Plane. If you are using an older version, then you will find remote access in the UniFi OS settings.
- Open the UniFi Network App
- Go to Settings > Control Plane
- Open the Console tab
- Check if Remote Access is enabled.
Enable UniFi Teleport
Enabling Teleport is really easy after you have made sure that everything is up-to-date. All we need to do is enable the feature in the UniFi Network app.
- Open the UniFi Network app
- Goto Settings > VPN
- Enable Teleport
You only need to generate a new invitation link (4) after you have enabled Teleport. Keep in mind that the link expires after 24 hours. Copy the link and send it to your mobile device for example.
Using UniFi Teleport
As mentioned there are two ways to use UniFi Teleport, but for both cases, we will first need to install the WiFiman app. This app is currently available for all operating systems and both mobile and desktop devices.
You can download the WiFiman app for desktop applications here on the UniFi download page. For mobile devices, you can find the app in the app stores, or use the QR code from the invitation link.
Connecting with WiFiman Desktop
After you have downloaded and installed WiFiman Desktop, you will see the option to connect to UniFi Teleport when you open the app. From here you have two options, log in with the site-admin account, or use the invitation link.
To log in, click on the blue link “Log in to the UI Account” or click in the top-left corner on the user icon and choose Log In. This will open the browser and allow you to sign in with your Ubiquiti account.
After you have successfully logged in, you can select your UniFi Cloud Gateway from the list and click on Connect. Once connected, you will see a green connection line between your device and the Cloud Gateway and how much data is going through the connection.
Another option to connect the WiFiman desktop app is to use the invitation link. Simply copy and paste the link in the link field and click on Connect. The connection will be remembered, allowing you to reconnect anytime when needed.
Connecting with the Mobile app
The WiFiman mobile app works similarly to the desktop version. Here we also have the option to sign in with our Ubiquiti account or use the invitation link. For the latter, we have a couple of options.
If you open the link on a desktop device, you will see a QR code that you can scan with your mobile device. The link will either take you to an introduction page, where you can download the WiFiman app, or if you already have the app installed, open the app directly.
In the app, click on Connect to add the UniFi Teleport connection. To use the connection, slide the toggle to On to make the connection with your UniFi Cloud Gateway. The first time you will be asked to install the VPN Configuration, after which the connection will be made.
It will take 5 to 10 sec for the connection to build after which you have a secure connection to the internet through your home network.
You can also use the sign-in method instead of the invitation link. Simply sign in with your site-admin account to connect to your network.
Revoking Access to Teleport
There are two ways to revoke access to the UniFi Teleport. The method depends on the status of the invitation. When the invitation is already accepted, you will need to go to Client Devices in the UniFi network, select the device, and Revoke Access under Settings.
If the invitation has not been accepted yet, then you can Revoke the invitation from the Teleport settings screen.
- Expand the Invitation History (click on Show)
- Hover over an invitation
- Click on Revoke
Wrapping Up
UniFi Teleport is a great way to easily set up and make a VPN connection through/to your home network. Just make sure that your UniFi OS and Network app are up-to-date to use this feature.
If you want to use this option with multiple users, then this method might not be the easiest to manage. A better option then is to use UniFi Identity.
Great article.. any plan to refresh about how to setup with zone base firewall rules and restrict access to other zones behind the same firewall appliance. thanks!
I already have an article about the zone-based firewall: UniFi Zone-Based Firewall – What you need to Know
I am not an NW expert but I plan to enhance my home network and install a Unifi Cloud Gateway (CG) Ultra or Max. But I will not be able to connect the CG directly to the Internet due to the router provided by my provider where reconfiguring it as pure modem is disabled. So, I will end-up with a cascaded router setup with double NAT.
As I understand from several forums this setup should work with the CG.
But I still want to access my NW behind the CG from the Internet by VPN. Would the Teleport setup work with it, or what other option do I have?
It will work just fine. Teleport can be used behind a NAT without the need for port forwarding or other firewall rules.
Second paragraph, last line, “create an invention link” – invention? I assume you mean “invitation” 😉 Thanks for all the good lazyadmin tricks and tips!
Thanks
June ’22 I reported that I couldn’t use Teleport over the cellular network. In the mean time I switched ISP’s from KPN to Odido and from that moment on my Teleport is working well. I didn’t need to enable IPv6 on my UDMP.
So my guess is that the problem might have had a relation with the way the ISP configures their equipment.
