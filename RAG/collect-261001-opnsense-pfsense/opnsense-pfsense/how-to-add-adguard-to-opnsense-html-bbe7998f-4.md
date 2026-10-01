---
id: collect-261001-opnsense-pfsense/opnsense-pfsense/how-to-add-adguard-to-opnsense-html-bbe7998f-4
title: "how-to-add-adguard-to-opnsense-html-bbe7998f"
domain: opnsense-pfsense
role: reference
task: reference
actors: []
dates: []
keywords: ["memory"]
source: docs/RAG/collect-261001-opnsense-pfsense/how-to-add-adguard-to-opnsense-html-bbe7998f.md
source_anchor: ""
source_lines: [194, 275]
sha256: 0a77d0a2549c47063b22a2b433126a08f7e671df57fcaa03539afa8b1950213e
---

# how-to-add-adguard-to-opnsense-html-bbe7998f

Step 3: Firewall Rules Configuration
- Allow DNS traffic (UDP port 53) from LAN clients to the AdGuard server IP.
- Set rules so that DNS traffic from clients cannot bypass AdGuard. For example, block outbound DNS requests to external DNS servers directly from clients.
- Implement rules to restrict DNS traffic to prevent leaks, ensuring only AdGuard handles DNS queries.
Step 4: Testing and Validation
After configuration, verify DNS traffic is routing through AdGuard:
- Use tools like nslookup or dig from a client device to confirm queries resolve via your AdGuard instance.
- Check AdGuard logs to see incoming DNS requests from your network.
- Use online DNS leak testers to ensure queries are not escaping your intended DNS resolver.
Following these steps ensures that your DNS traffic flows correctly through AdGuard on OPNsense, enhancing network privacy and ad-blocking effectiveness.
Testing the AdGuard Integration
After successfully integrating AdGuard with your OPNsense setup, it’s crucial to verify that the configuration is functioning correctly. Proper testing ensures that your network benefits from ad blocking and privacy features without disruptions.
Free tools Windows power users keep installed
One-click scans. No signup required.
Verify DNS Resolution
- Open a device connected to your network and browse to a website known for intrusive ads, such as example.com.
- Check if ads are blocked or if the website loads faster. If ads appear, review your AdGuard DNS settings.
- Use command-line tools like nslookup or dig to query the DNS server IP configured for AdGuard. Confirm it returns the expected responses.
Confirm Ad Blocking Functionality
- Visit popular ad-heavy sites—such as social media platforms or news outlets—to observe if advertisements are being blocked.
- Utilize browser extensions or developer tools to inspect network requests. Blocked requests typically show as failed or redirected.
- Test known ad domains, e.g., doubleclick.net, by pinging or querying them directly. Responses should be blocked or null, indicating proper filtering.
Check AdGuard Logs
- Access AdGuard’s dashboard through your web browser. The interface usually runs on the IP and port you specified during setup.
- Review logs for blocked domains, queries, and filtering actions. Confirm that requests from your devices are being processed and appropriately blocked.
- If logs show no activity, revisit your DNS forwarding rules, firewall settings, and AdGuard configuration to ensure all traffic passes through AdGuard correctly.
Troubleshooting Tips
- Ensure DNS settings on your devices point to the OPNsense IP where AdGuard is active.
- Verify that firewall rules permit DNS traffic (UDP port 53) to and from AdGuard.
- Clear browser cache and DNS cache on your test device to eliminate stale data.
Conducting these tests helps confirm that AdGuard is effectively filtering ads and protecting your network. Regular checks ensure ongoing performance and security benefits.
Troubleshooting Common Issues When Adding AdGuard to OPNsense
Integrating AdGuard with OPNsense can enhance your network’s ad-blocking and privacy capabilities. However, you may encounter some common issues during setup. Here’s a straightforward guide to troubleshoot effectively.
1. AdGuard Service Not Starting
- Check Service Status: Verify if the AdGuard service is running in OPNsense’s Services menu.
- Log Files: Review system logs for errors related to AdGuard startup. Use the Diagnostics > System Logs menu.
- Port Conflicts: Ensure that the ports AdGuard uses are not occupied by other services. Change ports if necessary in the AdGuard configuration.
2. DNS Resolution Failures
- DNS Settings: Confirm that OPNsense’s DNS resolver is correctly configured to route DNS queries through AdGuard.
- Firewall Rules: Ensure firewall rules allow DNS traffic (UDP/53) from your clients to AdGuard.
- AdGuard DNS: Make sure AdGuard DNS servers are correctly specified in the configuration.
3. Ad Blocking Not Working
- Filters and Blocklists: Verify that AdGuard’s filter lists are enabled and up to date.
- Client Configuration: Confirm that client devices are using the OPNsense device as their DNS server.
- Refresh Cache: Clear browser cache or restart devices to ensure new DNS settings take effect.
4. Performance Issues
- Resource Allocation: Check CPU and memory usage. Assign more resources if AdGuard impacts system performance.
- Update Software: Ensure both OPNsense and AdGuard are running the latest versions for stability improvements.
By systematically addressing these areas, most common issues with AdGuard on OPNsense can be resolved efficiently. Always consult official documentation or community forums if problems persist.
Additional Tips for Optimal Performance
To ensure AdGuard runs smoothly on your OPNsense setup, follow these best practices. Proper configuration enhances filtering accuracy and maintains system stability.
Quick wins for a faster PC:
Clear out junk files and repair common Windows errorsFree Scan →Fix the driver behind crashes, sound loss and screen glitchesFind Drivers →
Keep Your System Updated
- Regularly update OPNsense to benefit from security patches and new features.
- Ensure the AdGuard package is current. Check for updates through the OPNsense plugin manager or command line interface.
Optimize Resource Allocation
- Assign sufficient CPU and RAM resources to handle DNS filtering without impacting network performance.
- Monitor system resources via the OPNsense dashboard and scale as needed.
- Disable unnecessary services to free up resources for AdGuard.
Configure DNS Settings Carefully
- Set AdGuard as the primary DNS resolver in your network settings to ensure all DNS queries pass through it.
- Avoid conflicting DNS entries; use static IP addresses for devices that require fixed DNS configurations.
- Use DNS over HTTPS (DoH) or DNS over TLS (DoT) if supported, to enhance privacy and security.
Implement Proper Firewall Rules
- Create rules that allow DNS traffic (typically UDP port 53) to and from AdGuard.
- Restrict DNS traffic from untrusted sources to prevent DNS hijacking or poisoning.
- Regularly review firewall logs for anomalies related to DNS queries.
Utilize Logging and Monitoring
- Enable logging in AdGuard to track and troubleshoot filtering issues.
- Use system monitoring tools to observe the impact on network traffic and performance.
- Adjust filtering rules based on observed patterns and network needs.
Maintain Backup Configurations
- Regularly back up your OPNsense configuration, including AdGuard settings.
- Test restore procedures periodically to ensure quick recovery in case of issues.
Adhering to these tips helps optimize your AdGuard deployment on OPNsense, ensuring robust, efficient, and secure DNS filtering tailored to your network’s needs.
Conclusion
Integrating AdGuard with OPNsense enhances your network’s privacy and security by providing effective ad and tracker blocking. Following the steps outlined in this guide ensures a seamless setup, allowing you to enjoy a cleaner, faster browsing experience across all devices connected to your network.
To recap, the process involves installing the necessary packages, configuring DNS settings, and setting up AdGuard as a DNS resolver or forwarding service. This setup helps prevent unwanted content from reaching your devices, reduces bandwidth consumption, and enhances overall security by blocking malicious domains.
It’s important to regularly update AdGuard and your OPNsense system to benefit from the latest features and security patches. Monitoring logs and adjusting filtering rules can optimize performance and ensure you maintain control over what content gets blocked.
Crashes, No Sound, or Screen Glitches?
Random freezes, missing sound and display glitches usually trace back to one bad driver. Find and replace yours safely.Free scan · under a minute
