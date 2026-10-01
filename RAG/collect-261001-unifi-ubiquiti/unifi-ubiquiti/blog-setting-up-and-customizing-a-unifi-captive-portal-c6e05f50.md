---
id: collect-261001-unifi-ubiquiti/unifi-ubiquiti/blog-setting-up-and-customizing-a-unifi-captive-portal-c6e05f50
title: "blog-setting-up-and-customizing-a-unifi-captive-portal-c6e05f50"
domain: unifi-ubiquiti
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-unifi-ubiquiti/blog-setting-up-and-customizing-a-unifi-captive-portal-c6e05f50.md
source_anchor: ""
source_lines: [1, 87]
sha256: 9dcf81e130aa22d3e4a015e65c6843b0119beaae912330463ca94af48f33e4dd
---

# blog-setting-up-and-customizing-a-unifi-captive-portal-c6e05f50

Set up and customize a UniFi Captive Portal
Want to control how guests access your network? You've come to the right place. A Captive Portal is a great way to do it. With the UniFi Captive Portal, you can create a branded login page and manage guest access. It's also a great way to help keeping your network secure.
Let me show you how to set this up and customize step-by-step.
Let's go!
Setting up a UniFi Captive Portal
I feel that I don't need to explain anymore what a Captive Portal is so let's get right to it. To set up a Captive Portal on your UniFi network, follow these steps:
Access the UniFi Controller
Log in to your UniFi Controller via your web browser. The Controller’s dashboard will provide an overview of your network and devices.
Create a guest network
- Navigate to the "Settings" section and select "Wi-Fi".
- Click "Create New Wi-Fi Network".
- Enter a name for your guest network (e.g., "Guest Wi-Fi").
- Under "Network", select "Guest" as the purpose.
- Configure other settings as needed (e.g., security options, VLAN ID).
- Click "Apply Changes".
Enable the Captive Portal
- Go to "Settings" > "Guest Control".
- Toggle "Enable Guest Portal" to ON.
- 
Under "Authentication", select the method you prefer:
  - No authentication: Users will only see the splash page.
  - Simple password: Users must enter a password.
  - Hotspot: Users must enter a voucher code or login credentials.
- Customize the "Redirect URL" if you want users to be redirected to a specific webpage after logging in.
Customize the splash page
- In the "Guest Control" section, click on "Customize Portal".
- You can use the built-in editor to modify the splash page or upload your own HTML file for full customization.
- Add your company logo, welcome message, and terms of service to give the portal a branded look.
- Save your changes.
Adding some advanced customization
For more advanced customization, you can integrate third-party services or use the UniFi API (which isn't always great to be honest).
This allows you to create dynamic portals that can interact with external databases, provide personalized content, or integrate with marketing tools.
Example: Customizing with HTML and CSS
If you choose to upload your own HTML file, you can create a fully customized splash page. Here’s a basic example:
<!DOCTYPE html>
<html>
<head>
    <title>Welcome to Our Guest WiFi</title>
    <style>
        body {
            font-family: Arial, sans-serif;
            text-align: center;
            background-color: #f4f4f4;
            padding: 50px;
        }
        .container {
            max-width: 600px;
            margin: auto;
            background: white;
            padding: 20px;
            box-shadow: 0 0 10px rgba(0, 0, 0, 0.1);
        }
        h1 {
            color: #333;
        }
        p {
            color: #666;
        }
    </style>
</head>
<body>
    <div class="container">
        <h1>Welcome to Our Guest WiFi</h1>
        <p>Please agree to the terms of service to connect.</p>
        <form action="/guest/s/default/">
            <input type="hidden" name="cmd" value="authenticate">
            <button type="submit">Connect</button>
        </form>
    </div>
</body>
</html>
Upload this HTML file to your UniFi Controller, and it will be used as the splash page for your guest network.
Monitoring and managing guest access
Once your Captive Portal is set up, you can monitor and manage guest access through the UniFi Controller:
- View connected clients: Go to the "Clients" section to see a list of all devices connected to your guest network.
- Generate reports: Use the "Insights" and "Statistics" sections to generate reports on guest network usage.
- Manage sessions: You can disconnect clients, extend session times, or generate vouchers (if using the Hotspot method).
Final Thoughts
Setting up a UniFi Captive Portal is a straightforward process. It's a great way to manage guest access.
Whether you opt for a simple setup or a fully customized solution, a Captive Portal improves network security and provides valuable insights into network usage.
To design the guest splash page itself, try our captive portal generator.
If you'd like a hand setting this up, we can help. At UniHosted we specialize in UniFi network management. Let me know if you have any questions!
Related guides
Keep reading
- UniFi repeater: how to set up an AP as a WiFi extenderUniFi has no repeater mode. It uses wireless meshing. How to set up a UniFi AP as a WiFi extender, what you need, and what it costs in speed. Read guide
- How to set up UniFi Cloud Key for multi-site managementIn this guide, we’ll walk you through the steps to set up your UniFi Cloud Key for multi-site management. Read guide
- How to set up alerts for your UniFi Controller: A step by step guideStep-by-step guide on how to set push notifications for your UniFi Controller Read guide
