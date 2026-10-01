---
id: collect-261001-unifi-ubiquiti/unifi-ubiquiti/blog-how-to-manage-multiple-unifi-controllers-on-a-single-ip-897e6bbd
title: "blog-how-to-manage-multiple-unifi-controllers-on-a-single-ip-897e6bbd"
domain: unifi-ubiquiti
role: reference
task: reference
actors: []
dates: []
keywords: ["cost", "distribution"]
source: docs/RAG/collect-261001-unifi-ubiquiti/blog-how-to-manage-multiple-unifi-controllers-on-a-single-ip-897e6bbd.md
source_anchor: ""
source_lines: [1, 37]
sha256: 4ae07da521b2ebb08e330f5e87ff45db2e793e776f200c626b4c22327afd30aa
---

# blog-how-to-manage-multiple-unifi-controllers-on-a-single-ip-897e6bbd

Manage multiple UniFi controllers on a single IP: A step by step guide
Running multiple UniFi Controllers on a single IP addres is an efficient way to run your networks. Quite a few MSPs and larger organizations use this method (because its also cost-effective).
Here is how to set it up:
Before we get to juicy part, just a quick note: If you manage client networks on a self-hosted UniFi controller. Please stop. Sooner or later this will cause issues! It's fine for home users, but definitely not recommended for businesses. We've built a secure and reliable UniFi hosting solution that takes the hassle out of managing controllers. You can try it for free.
Port Management
- Each UniFi Controller must listen on different ports to avoid conflicts. For instance, you could set one controller to use the default ports (HTTP: 8080, HTTPS: 8443) and configure another to use alternative ports (HTTP: 8081, HTTPS: 8444).
- Remember to update your firewall rules to allow traffic on these ports.
DNS Configuration
- Use subdomains to direct to different controllers, e.g., controller1.domain.com ,controller2.domain.com .
- These subdomains can be mapped to the same IP but different ports, directing users to the correct controller interface.
SSL Certificates
- For HTTPS, consider using wildcard SSL certificates or individual SSL certificates for each subdomain.
- This ensures secure and encrypted connections to each controller.
Network Segmentation
- Although the controllers share an IP, they should be logically separated within the network. This enhances security and prevents potential interference between controllers.
- VLANs can be useful in this context to segregate traffic.
Resource Allocation
- Ensure the server hosting the controllers has sufficient resources (CPU, RAM, Disk Space) to handle multiple instances.
- Monitor performance regularly to avoid resource contention.
Backup and Recovery
- Independently back up the configuration of each controller.
- This is crucial for disaster recovery and maintaining operational integrity.
Controller Version Consistency:
- Keep all controllers updated to the same UniFi software version for consistency and compatibility.
Access Management:
- Implement strong access controls and authentication methods for each controller to prevent unauthorized access.
Advanced Considerations
- Load Balancing: For high availability and load distribution, consider implementing load balancers that can direct traffic to the appropriate controller based on load, ensuring optimal performance.
- Automated Deployment: Utilize scripting or automation tools for deploying new controllers or making bulk changes, ensuring efficient and error-free operations.
Final Thoughts
Managing multiple UniFi Controllers on a single IP is a viable solution for MSPs and large organizations looking to streamline their network management. By ensuring proper port management, secure DNS configurations, adequate resource allocation, and robust security practices, you can effectively run multiple controllers in a single, consolidated environment.
At UniHosted, we understand the complexities involved in such setups. Our cloud-hosted UniFi Controller solutions offer the flexibility and scalability needed for managing multiple controllers efficiently.
Related guides
Keep reading
- How to configure UniFi SSL certificate: A step by step guideStep-by-step guide on how to configure UniFi SSL certificates to improve your network Read guide
- UniFi Controller ExplainedWhat does the UniFi Controller actually do? Plain breakdown of self-hosted, cloud, and managed hosting setups with the pros and cons of each. Read guide
- UniFi DNS settings: How to configure themSet DNS settings in your UniFi Controller and gateway to improve speed, stability, and device adoption across your network. Full step-by-step included Read guide
