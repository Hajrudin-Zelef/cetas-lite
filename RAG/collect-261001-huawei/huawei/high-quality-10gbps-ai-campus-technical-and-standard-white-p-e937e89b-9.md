---
id: collect-261001-huawei/huawei/high-quality-10gbps-ai-campus-technical-and-standard-white-p-e937e89b-9
title: "high-quality-10gbps-ai-campus-technical-and-standard-white-p-e937e89b"
domain: huawei
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-huawei/high-quality-10gbps-ai-campus-technical-and-standard-white-p-e937e89b.md
source_anchor: ""
source_lines: [275, 312]
sha256: d1e5acb860b06c480b0665ef8c5fbf9ca9d7362e00c7b72b66b95998314130f7
---

# high-quality-10gbps-ai-campus-technical-and-standard-white-p-e937e89b

The IoT trend and multi-network convergence (5G/         and become equally important as the services it                                       Blurred network                   Lateral movement           Physical isolation
Wi-Fi and cloud-edge synergy) have completely            protects. Therefore, campus networks must build a                                     boundaries:                       attack (attacker           failure: 5G/Wi-Fi hybrid
                                                                                                                                               5G and Wi-Fi 6/7 are              moves freely inside        networking, cloud-edge
broken the traditional "castle-and-moat" security        full-chain, all-scenario security protection system
                                                                                                                                               deeply integrated with            the network)               synergy, and multi-cloud
boundary model that implements physical                  covering terminal access, network transmission,                                       wired networks.                   APT stealthy               interconnection make
isolation, resulting in exponential expansion            and service data, to provide multi-dimensional                                        Data is frequently                propagation (using         the separation between
of the attack surface and ubiquitous security            i n - d e p t h d e fe n s e i n a co n ve rg e d n e t wo r k                        exchanged between cloud           complex network            internal and external
                                                                                                                                               data centers (DCs), edge          topologies to hide         networks disappear.
threats. Network security has evolved from a             environment.
                                                                                                                                               nodes, and terminals              traces)                    Surge in east-west traffic:
support system to a core productivity element,                                                                                                 (east-west traffic).              Cross-domain               Traditional border firewalls
                                                                                                                              Network          The microservice                  ransomware                 struggle to effectively
                                                                                                                            transmission       architecture greatly              infection (from the        monitor the high-volume,
                          Table 1-3 Security service scenarios and typical risks                                                               increases the complexity          IT domain to the OT        complex communication
                                                                                                                                               of internal network               domain or different        traffic between servers,
    Service                                                                                                                                    communication.                    cloud environments)        applications, and
                    Specific Scenario/Risk               Typical Risk               Key Feature/Challenge
   Scenario                                                                                                                                                                      Network sniffing           microservices.
                                                                                                                                                                                 and man-in-the-            Traffic encryption: Deep
                    IoT devices:                    Ransomware                      Massive heterogeneous                                                                        middle (MITM)              inspection is harder due to
                    IoT devices deployed in         infection (causing              devices: Numerous types                                                                      attacks (listening to      the significant volume of
                    public places are prone to      device suspension               and unclear ledgers make                                                                     or tampering with          encrypted traffic (such as
                    unauthorized access and         and data encryption)            management difficult.                                                                        data on transmission       TLS traffic).
                    spoofing.                       DDoS botnet                     Long lifecycle/Difficult                                                                     links)
                    Firmware of cameras and         (devices are                    update: Firmware
                    sensors is not updated in       controlled to launch            update is delayed or
                    a timely manner (high-risk      attacks)                        impossible, leaving                                        Increased data mobility           Data theft                 Ubiquitous data: Data is
                    vulnerabilities exist, CVSS     Key infrastructure              numerous vulnerabilities                                   and exposure:                     through MITM               distributed on terminals,
                    ≥ 9.0).                         suspension                      unaddressed.                                               Sensitive data frequently         attacks (sensitive         edges, and clouds, and
                    Industrial systems such         (industrial control             Physical exposure: Devices                                 flows between terminals,          information                the data flow paths are
                    as PLC and SCADA                devices are                     are often deployed in                                      edge nodes, clouds, and           intercepted during         complex.
                    use default or weak             damaged)                        unattended or open                                         different service systems.        transmission)              Multi-environment
                    passwords.                      Vehicle control                 environments and are                                       Data is stored on hybrid          Phishing for               storage: Data is stored
                    The OTA upgrade                 hijacking                       prone to unauthorized                                      clouds, edge nodes, and           credentials (for           in public clouds, private
                    mechanism of the in-            (endangering                    access and spoofing.                                       even terminals.                   unauthorized access        clouds, local DCs, edge
                    vehicle system has              personal safety)                Lack of security baselines:                                Microservice-based                to data)                   devices, and other
    Terminal        vulnerabilities.                                                The default configurations                                 applications complicate           Internal data leakage      environments.
