---
id: collect-260926-mikrotik/mikrotik/mikrotik-capsman-the-complete-guide-to-centralized-wireless-network-management-tech-layer--5
title: "Floor 1 CAPs (identified by identity prefix)"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: ["cost", "license", "memory", "training"]
source: docs/RAG/lot-mikrotik/forum/wifi-capsman/mikrotik-capsman-the-complete-guide-to-centralized-wireless-network-management-tech-layer-x-com.md
source_anchor: ""
source_lines: [1107, 1275]
sha256: faae885541e1e012cabbd9117403a8b37777df9a126a7506d5bb6dcdc01e061d
---

# Floor 1 CAPs (identified by identity prefix)

```
# High-density channel configuration
/caps-man channel
add name=ch-5ghz-dense \
    band=5ghz-n/ac \
    control-channel-width=20mhz \
    extension-channel=Ce \
    tx-power=17
# Datapath with client isolation
/caps-man datapath
add name=dp-guest \
    local-forwarding=no \
    bridge=bridge-hotspot \
    client-to-client-forwarding=no
# Configuration with rate limiting
/caps-man configuration
add name=cfg-guest \
    ssid="Hotel-WiFi" \
    channel=ch-5ghz-dense \
    datapath=dp-guest \
    security.authentication-types="" \
    rate.rx-rate=10M \
    rate.tx-rate=5M \
    max-sta-count=100
# Hotspot setup on controller
/ip pool
add name=hotspot-pool ranges=10.0.20.10-10.0.20.250
/ip dhcp-server
add name=dhcp-hotspot interface=bridge-hotspot address-pool=hotspot-pool
/ip hotspot
add name=hs-guest interface=bridge-hotspot address-pool=hotspot-pool
/ip hotspot profile
set default login-by=http-chap,https \
    html-directory=flash/hotspot \
    rate-limit="5M/10M"
```
## 11. CAPsMAN vs. Alternative Solutions

### Feature Comparison Table

| Feature | MikroTik CAPsMAN | Ubiquiti UniFi | Cisco WLC | Aruba Central | 
|---|---|---|---|---|
| License Cost | Free | Free | Per-AP licensing | Subscription | 
| Controller Hardware | Any RouterOS device | Dedicated/Cloud | Dedicated appliance | Cloud-only | 
| Max APs (typical) | 100-150 | 50-2000 (by model) | 500-6000 | 10000+ | 
| Configuration Complexity | Medium-High | Low | High | Medium | 
| CLI Access | Full | Limited | Full | Limited | 
| RADIUS Support | Yes | Yes | Yes | Yes | 
| WPA3 | Yes (ROS7) | Yes | Yes | Yes | 
| Seamless Roaming | Layer 2 | Layer 2/3 | Layer 2/3 | Layer 2/3 | 
| API/Automation | Yes (API) | Yes (API) | Yes (API) | Yes (API) | 
| WiFi 6/6E Support | Limited | Yes | Yes | Yes | 

### When to Choose CAPsMAN

CAPsMAN is the right choice when:

- Budget is a primary concern
- Existing MikroTik infrastructure is in place
- Team has MikroTik expertise
- Full CLI control is required
- Deployment size is under 100 APs
- Integration with MikroTik routing features is needed

### CAPsMAN Limitations

Consider alternatives when:

- Deployments exceed 100-150 APs
- Advanced RF optimization is required
- Dedicated support contracts are mandatory
- WiFi 6E is a requirement (limited MikroTik support currently)
- GUI simplicity is a priority for operations team
- AI/ML-based wireless optimization is desired

### Cost Comparison (50 AP Deployment)

| Solution | Controller Cost | AP Cost (×50) | Annual License | 5-Year TCO | 
|---|---|---|---|---|
| MikroTik CAPsMAN | $300 (RB4011) | $3,500 ($70/AP) | $0 | ~$3,800 | 
| Ubiquiti UniFi | $200 (Cloud Key) | $7,500 ($150/AP) | $0 | ~$7,700 | 
| Cisco (Meraki) | Cloud-hosted | $15,000 ($300/AP) | $7,500 | ~$52,500 | 
| Aruba Central | Cloud-hosted | $12,500 ($250/AP) | $5,000 | ~$37,500 | 

*Note: Prices are approximate and vary by region and reseller.*

## 12. Future of MikroTik CAPsMAN

### RouterOS 7 CAPsMAN Improvements

RouterOS 7 introduces several enhancements:

- **WPA3 Support:** Full WPA3-Personal and WPA3-Enterprise
- **Improved 802.11ax:** WiFi 6 support on compatible hardware
- **Enhanced PMF:** Better protected management frames implementation
- **Performance Improvements:** Optimized control plane processing
- **New Wireless Package:** Rewritten wireless driver architecture

### WiFi 6 and WiFi 6E Considerations

Current MikroTik WiFi 6 capable devices:

- Audience (WiFi 6 tri-band)
- hAP ax² (WiFi 6)
- hAP ax³ (WiFi 6)

WiFi 6E (6GHz) support: Limited availability. Check MikroTik announcements for updates.

### Recommended RouterOS Version

| Deployment Type | Recommended Version | 
|---|---|
| Production (stability focus) | RouterOS 6.49.x (long-term) | 
| New deployments | RouterOS 7.x stable | 
| WiFi 6 required | RouterOS 7.x stable | 
| WPA3 required | RouterOS 7.x stable | 

### Community Resources

- **MikroTik Forum:** forum.mikrotik.com
- **MikroTik Wiki:** wiki.mikrotik.com
- **Reddit r/mikrotik:** Community discussions and troubleshooting
- **MikroTik Training:** Official MTCNA, MTCWE courses

## 13. Conclusion

MikroTik CAPsMAN provides enterprise-grade centralized wireless management without licensing costs. The platform suits organizations with existing MikroTik infrastructure and teams comfortable with RouterOS configuration.

### Key Takeaways

- CAPsMAN eliminates repetitive AP configuration through centralized management
- Local forwarding mode minimizes controller dependency for data traffic
- Provisioning rules enable zero-touch AP deployment
- Certificate-based authentication secures CAP-to-controller communication
- Access lists and signal thresholds improve roaming behavior
- Regular backups and documentation are essential for disaster recovery

### Recommended Next Steps

1. Build a lab environment with one controller and 2-3 CAPs
2. Test basic provisioning with single SSID
3. Add VLANs and multiple SSIDs
4. Implement RADIUS authentication
5. Practice troubleshooting common issues
6. Document your production configuration
7. Deploy in phases, starting with non-critical areas

### Quick Reference Commands

# Enable CAPsMAN
/caps-man manager set enabled=yes
# View connected CAPs
/caps-man remote-cap print
# View connected clients
/caps-man registration-table print
# View provisioned interfaces
/caps-man interface print
# Enable CAP mode
/interface wireless cap set enabled=yes interfaces=wlan1,wlan2
# Debug logging
/system logging add topics=caps,wireless action=memory

## Additional Resources

### Official Documentation

### Related Topics

Check our list of MikroTik guides
