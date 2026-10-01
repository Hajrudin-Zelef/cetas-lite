---
id: collect-261001-unifi-ubiquiti/unifi-ubiquiti/hc-en-us-articles-31504711305879-configuring-third-party-sip-server-with-unifi-a-918247c3
title: "hc-en-us-articles-31504711305879-configuring-third-party-sip-server-with-unifi-a-918247c3"
domain: unifi-ubiquiti
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-unifi-ubiquiti/hc-en-us-articles-31504711305879-configuring-third-party-sip-server-with-unifi-a-918247c3.md
source_anchor: ""
source_lines: [1, 105]
sha256: cc6a944423d03b0a14b706ae07663e23fb320818144b6001cbe4931e067cf3cb
---

# hc-en-us-articles-31504711305879-configuring-third-party-sip-server-with-unifi-a-918247c3

Configuring Third-Party SIP Server with UniFi Access
UniFi Access Intercoms and Reader Pros both support integration with third-party SIP (Session Initiation Protocol) servers, enabling doorbell calls to be routed through your existing SIP system or UniFi Talk. This allows compatible SIP devices to receive doorbell notifications, communicate with visitors, and remotely unlock doors.
If you're looking to use the UniFi Access mobile app, UniFi Endpoint mobile app, or phone numbers as doorbell call receivers, click here.
| Update your UniFi OS, Access application, and Access devices to the latest versions for the newest features and optimal performance. | 
Requirements
UniFi device and version support
- Access application 4.0.21 or later
- UniFi OS 4.1.22 or later
- UniFi Access Readers
  - G3 Intercom (UA-G3-Intercom) 1.9.25 or later
  - Intercom (UA-Intercom) 1.9.24 or later
  - G3 Reader Pro (UA-G3-Pro) 1.12.25 or later
  - G2 Reader Pro (UA-G2-Pro) 1.12.24 or later
Compatible SIP servers
Some of the supported SIP servers include:
- UniFi Talk
- RingCentral (Video calling not supported due to lack of early media support)
- FreeSWITCH
- FreePBX
- AVLinkPro
- Control4
Configuring SIP in UniFi Access
SIP servers and recipient numbers are shared across all Intercoms and Reader Pros within the same UniFi Access deployment.
Adding an SIP Server
- Navigate to your SIP server and configure the settings.
- Copy the required information for later use. See the table below for an example.
| UniFi Access | RingCentral | FreeSWITCH | FreePBX | 
| SIP Server Field Names |  |  |  | 
| Domain/Server | SIP Domain | Server IP / Domain | Server IP / Domain | 
| Proxy Domain/Server (Optional) | Outbound Proxy | - | - | 
| SIP Client Account Field Names |  |  |  | 
| User ID/Extension | User Name | User Extension | User Extension | 
| Password/Secret | Password | Secret | Secret | 
| Authentication ID (Optional) | Authentication ID | - | - | 
Configuring SIP in UniFi Access
- Navigate to Access application > Interface Designer > select a reader > Directory or Doorbell Call > Third-Party SIP > Configure Third-Party SIP.
- Enter the SIP server information:
  - Server Name: Create a name for easy identification.
  - Domain/Server
  - Port: Defaults to 5060. It may differ according to the SIP server.
  - Select the UDP (default), TCP, or TLS transport mode.
- Enter the SIP client account information for your reader. This allows the reader to make a doorbell call to other extensions within the SIP system:
  - User ID/Extension: Enter the SIP client account's user name/extension (e.g., 0001) for your reader. You must create it in your SIP server system first.
  - Password/Secret
  - Authentication ID (Optional)
- Click Save.
Adding the SIP Phone as a Doorbell Call Receiver
- Navigate to your SIP server and add a SIP device. Copy its SIP username/extension (e.g., 0005).
- Navigate to Access application > Interface Designer > select the reader where the SIP server was configured > Directory or Doorbell Call > Third-Party SIP > Enter Recipient Number and add the SIP username (e.g., 0005).
- On your SIP phone screen, enter the following information. This may differ according to the SIP phone.
  - SIP username/extension (e.g., 0005)
  - SIP password
  - SIP server host/domain
  - Port
Once done, the SIP phone with the recipient number 0005 will receive a doorbell call when the doorbell or directory is triggered on the corresponding reader.
FAQs
Can I add multiple SIP servers to UniFi Access?
No. Currently, UniFi Access supports only one SIP server at a time.
How do I answer a doorbell call on my SIP phone?
If your SIP phone's recipient number is added as a doorbell call receiver, you will get a phone call when a visitor triggers the doorbell or directory on a reader.
Answer an incoming call and press * on the keypad to unlock the door.
For some SIP phones, such as those using the FreeSWITCH SIP server, you may need to press * twice to unlock the door.
What should I do if there is no live view when receiving a doorbell call on my SIP phone?
If your SIP phone rings but doesn't display live video, the issue is likely related to the early media setting on your SIP server platform. To fix this, ensure that early media is enabled on your SIP server platform.
Note that the exact setting name can vary across different platforms. Check your SIP platform’s documentation or contact their support team for guidance on enabling this option.
How do I edit the SIP account configured for a reader?
If a reader is configured as a SIP client to receive SIP calls, you can edit its SIP account using one of the following methods.
- Method 1: Navigate to Interface Designer > select a reader > Directory or Doorbell Call > Third-Party SIP > Edit SIP Account.
- Method 2: Navigate to Interface Designer > Templates > select a template > Assigned Devices > Manage > hover over the i icon > Configure SIP Account.
How do I add a G3 Talk Phone as a SIP server?
UniFi Talk Phone Requirements
- G3 Touch Enterprise (UTP-G3-Touch-Enterprise)
- G3 Touch Wall (UTP-G3-Touch-Wall)
- G3 Touch Pro (UTP-G3-Touch-Pro)
Adding an SIP Server in UniFi Talk
- Navigate to Talk application > Phones > Add Third-Party Device, enter the Device Name and Assign To a user.
- Click on the device you just created and go to Overview to copy the SIP Server Host, SIP Username, and SIP Password.
Configuring SIP in UniFi Access
- Navigate to Access application > Interface Designer > select a reader > Directory or Doorbell Call > Third-Party SIP > Configure Third-Party SIP.
- Enter the SIP server information:
  - Server Name: Create a name for easy identification.
  - Domain/Server: Enter the SIP Server Host you copied earlier.
  - Port: Defaults to 5060. It may vary depending on the SIP server.
  - Select the UDP transport mode.
- Enter the SIP client account information:
  - User ID/Extension: Enter the SIP Username (e.g., 0002) you copied earlier.
  - Password/Secret: Enter the SIP Password you copied earlier.
- Click Save.
Adding the G3 Talk Phone as a Doorbell Call Receiver
- Navigate to Talk application > Phones > Add Third-Party Device, enter the Device Name and Assign To a user.
- Click on the device you just created and go to Overview to copy the SIP Username (e.g., 0003).
- Navigate to Access application > Interface Designer > select the reader where the SIP server was configured > Directory or Doorbell Call > Third-Party SIP > Enter Recipient Number and add the SIP username (e.g., 0003).
- On your G3 Talk phone's initial setup screen, select Third-Party PBX > Manual Configuration, enter the following information, and click Connect.
  - Username: Enter the SIP Username.
  - Password: Enter the SIP Password.
  - Domain: Enter the SIP Server Host.
  - Port: Enter 5060.
Once done, UniFi Talk Phone assigned extension 0003 will receive a doorbell call when the doorbell or directory is triggered on the corresponding reader. Answer an incoming call and press * on the keypad to unlock the door.
How do I download the UniFi Talk phone support files?
To provide the necessary information for technical support troubleshooting, please prepare the following:
- Download the UniFi Talk application support file: Navigate to Site Manager > UniFi Console > Talk > Settings > Control Plane > Console > Support File > Full > Download.
- Provide the Talk phone serial number: On your Talk phone, navigate to Settings > System > About > Serial Number.
- Provide the Talk phone support file:
  - On your Talk phone, navigate to Settings > System > Upload Support File.
  - Navigate to Settings > Advanced > Feedback, and provide the time of upload along with the description you entered (for example, the serial number or support ticket number).
