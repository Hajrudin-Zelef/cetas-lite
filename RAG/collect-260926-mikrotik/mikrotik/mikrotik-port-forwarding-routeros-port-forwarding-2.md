---
id: collect-260926-mikrotik/mikrotik/mikrotik-port-forwarding-routeros-port-forwarding-2
title: "Mikrotik Port Forwarding - RouterOS Port Forwarding"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: ["attention"]
source: docs/RAG/lot-mikrotik/RouterOS/mikrotik-port-forwarding-routeros-port-forwarding.md
source_anchor: ""
source_lines: [113, 169]
sha256: dbe1fd7b0b7362700059749c00dd4005f642093b25dcbf86e6168fba40660501
---

# Mikrotik Port Forwarding - RouterOS Port Forwarding

### Surveillance Systems (NVR/DVR)

Security is a top concern for many homeowners and businesses. Network Video Recorders (NVRs) and Digital Video Recorders (DVRs) used for video surveillance can benefit from port forwarding. By forwarding specific ports, you can remotely access your surveillance system, view live camera feeds, and review recorded footage from anywhere with an internet connection.

## **Network Diagram**

In this section, we'll provide a visual representation of the network setup, which is essential for you to understand the context of the port forwarding configuration. This diagram illustrates how the MikroTik router is connected to both the WAN and LAN and how port forwarding plays a vital role in the overall network architecture.

The network setup consists of the MikroTik router with its ether1 interface connected to the WAN, having an IP address of 120.50.–.198. Meanwhile, the ether2 interface is connected to a LAN switch with an IP address of 193.168.20. Within the LAN, there are three servers: a web server, an FTP server, and an SSH server. These servers are only accessible to users within the local network. Port forwarding will enable users outside the local network to access these servers via the Internet.

## Alternative Methods for Mikrotik Port Forwarding

While port forwarding is a widely used method for network configuration, there are alternative approaches and tools that can achieve similar results:

### Virtual Private Networks (VPNs)

VPNs provide a secure and versatile way to access resources on your network remotely. By setting up a VPN server on your MikroTik router, you can create a secure tunnel for remote access to your local network, bypassing the need for port forwarding. VPNs are particularly useful for remote desktop access and ensuring data encryption.

### Reverse Proxies

For web hosting purposes, reverse proxies can be an alternative to port forwarding. A reverse proxy server can receive incoming web requests and route them to the appropriate backend server, effectively hiding the internal server's IP address. This method can enhance security and load balancing for web applications.

### Cloud Services

Consider leveraging cloud-based solutions for certain applications. Cloud-hosted services, such as cloud-based surveillance systems or cloud-based web hosting platforms, eliminate the need for port forwarding by providing remote access through dedicated cloud infrastructure.

## Notes and Recommendations

In this section, we present important considerations and recommendations to optimize your MikroTik port forwarding setup and safeguard against common pitfalls. By adhering to these guidelines, you can ensure a seamless and secure port forwarding configuration:

### Keep RouterOS Updated

Before embarking on your port forwarding journey, it is imperative to verify that your MikroTik router is running the latest version of RouterOS. Regularly updating your router's operating system ensures that you benefit from the latest features, bug fixes, and security enhancements. Visit the MikroTik website to download and install the most recent RouterOS release.

### Firewall Configuration

To enable successful port forwarding, pay close attention to your server's firewall settings. In addition to configuring port forwarding rules on your MikroTik router, you must permit incoming connections to the designated ports on the target server. Failing to adjust the server's firewall may result in inaccessible services. Ensure that any firewall software or hardware on the server is properly configured to allow traffic on the forwarded ports.

### Protocol and Port Accuracy

Precision is key when specifying the protocol and port numbers for your port forwarding rules. Be certain that you accurately identify the protocol type (e.g., TCP, UDP) and the port number associated with the service you intend to forward. Incorrect settings can lead to connectivity issues or security vulnerabilities. Consult the documentation of the service or application you are forwarding to verify the correct protocol and port details.

### Logging and Monitoring

Implement logging and monitoring practices to keep track of port forwarding activities. MikroTik routers offer comprehensive logging capabilities that can assist in troubleshooting and security monitoring. Regularly review logs to detect any unusual or unauthorized access attempts. Consider setting up alerts or notifications for specific events, ensuring timely response to potential threats.

### Security Best Practices

While port forwarding is a powerful tool for remote access, it should be approached with security in mind. Implement strong authentication mechanisms for services exposed through port forwarding. Utilize secure, complex passwords and, where applicable, consider implementing additional security layers such as VPNs or two-factor authentication (2FA). Regularly audit and update access credentials to maintain a high level of security.

### Regular Testing

Periodically test the functionality of your port forwarding rules to verify that they are working as intended. Attempt remote access from outside your local network to confirm that the forwarded services are accessible. Regular testing helps identify and address any configuration issues or changes in network conditions.

## Conclusion

We hope that with the help of this article, you now have a better understanding of ***Mikrotik port forwarding*** and can set up the port forwarding configuration on MikroTik without any issues. If you run into an issue or have any questions, you can post them in the comment section below or contact us via live chat or e-mail.
