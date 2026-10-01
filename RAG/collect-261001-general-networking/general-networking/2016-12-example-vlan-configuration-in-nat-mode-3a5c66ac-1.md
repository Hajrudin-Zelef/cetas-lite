---
id: collect-261001-general-networking/general-networking/2016-12-example-vlan-configuration-in-nat-mode-3a5c66ac-1
title: "2016-12-example-vlan-configuration-in-nat-mode-3a5c66ac"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-general-networking/2016-12-example-vlan-configuration-in-nat-mode-3a5c66ac.md
source_anchor: ""
source_lines: [1, 180]
sha256: de6de02efa748162f58c088eca9888b98cb57ff1d1bd4fb2e51e861dd1f955f7
---

# 2016-12-example-vlan-configuration-in-nat-mode-3a5c66ac

**E****xa****m****p****l****e VLAN configuration in NAT mode**

In this example two different internal VLAN networks share one interface on the FortiGate unit, and share the connection to the Internet. This example shows that two networks can have separate traffic streams while sharing a single interface. This configuration could apply to two departments in a single company, or to different companies.

There are two different internal network VLANs in this example. VLAN_100 is on the 10.1.1.0/255.255.255.0 subnet, and VLAN_200 is on the 10.1.2.0/255.255.255.0 subnet. These VLANs are connected to the VLAN switch, such as a Cisco 2950 Catalyst switch.

The FortiGate internal interface connects to the VLAN switch through an 802.1Q trunk. The internal interface has an IP address of 192.168.110.126 and is configured with two VLAN subinterfaces (VLAN_100 and VLAN_200). The external interface has an IP address of 172.16.21.2 and connects to the Internet. The external interface has no VLAN subinterfaces.


**Fo****r****t****i****G****a****t****e unit with VLANs in NAT mode**

When the VLAN switch receives packets from VLAN_100 and VLAN_200, it applies VLAN ID tags and forwards the packets of each VLAN both to local ports and to the FortiGate unit across the trunk link. The FortiGate unit has policies that allow traffic to flow between the VLANs, and from the VLANs to the external network.

This section describes how to configure a FortiGate unit and a Cisco Catalyst 2950 switch for this example network topology. The Cisco configuration commands used in this section are IOS commands.

It is assumed that both the FortiGate unit and the Cisco 2950 switch are installed and connected and that basic configuration has been completed. On the switch, you will need to be able to access the CLI to enter commands.

Refer to the manual for your FortiGate model as well as the manual for the switch you select for more information.

It is also assumed that no VDOMs are enabled.


**G****e****n****e****r****a****l configuration steps**

The following steps provide an overview of configuring and testing the hardware used in this example. For best results in this configuration, follow the procedures in the order given. Also, note that if you perform any additional actions between procedures, your configuration may have different results.

**1****.** Configure the FortiGate unit

- Configure the external interface
- Add two VLAN subinterfaces to the internal network interface
- Add firewall addresses and address ranges for the internal and external networks
- Add security policies to allow:
- the VLAN networks to access each other
- the VLAN networks to access the external network.

**2****.** Configure the VLAN switch


**C****on****f****i****gu****r****e the FortiGate unit**

Configuring the FortiGate unit includes:


**C****on****f****i****gu****r****e the external interface**

The FortiGate unit’s external interface will provide access to the Internet for all internal networks, including the two VLANs.


**T****o configure the external interface – web-based manager**

**1****.** Go to **S****ys****t****e****m > Network > Interface**.

**2****.** Select **E****d****i****t** for the external interface.

**3****.** Enter the following information and select **O****K**:

**A****dd****r****ess****i****n****g mode**                     Manual

**I****P****/****N****e****t****w****o****r****k Mask**                       172.16.21.2/255.255.255.0


**T****o configure the external interface – CLI**

config system interface edit external

set mode static

set ip 172.16.21.2 255.255.255.0

end


**A****d****d VLAN subinterfaces**

This step creates the VLANs on the FortiGate unit internal physical interface. The IP address of the internal interface does not matter to us, as long as it does not overlap with the subnets of the VLAN subinterfaces we are configuring on it.

The rest of this example shows how to configure the VLAN behavior on the FortiGate unit, configure the switches to direct VLAN traffic the same as the FortiGate unit, and test that the configuration is correct.

Adding VLAN subinterfaces can be completed through the web-based manager, or the CLI.


**T****o add VLAN subinterfaces – web-based manager**

**1****.** Go to **S****ys****t****e****m > Network > Interface**.

**2****.** Select **C****r****ea****t****e New**.

**3****.** Enter the following information and select **O****K**:

**N****a****m****e**                                           VLAN_100

**I****n****t****e****r****f****ac****e**                                     internal

**V****L****A****N ID**                                      100

**A****dd****r****ess****i****n****g mode**                     Manual

**I****P****/****N****e****t****w****o****r****k Mask**                       10.1.1.1/255.255.255.0

**A****d****m****i****n****i****s****t****r****a****t****i****v****e Access**             HTTPS, PING, TELNET

**4****.** Select **C****r****ea****t****e New**.

**5****.** Enter the following information and select **O****K**:

**N****a****m****e**                                           VLAN_200

**I****n****t****e****r****f****ac****e**                                     internal

**V****L****A****N ID**                                      200

**A****dd****r****ess****i****n****g mode**                     Manual

**I****P****/****N****e****t****w****o****r****k Mask**                       10.1.2.1/255.255.255.0

**A****d****m****i****n****i****s****t****r****a****t****i****v****e Access**             HTTPS, PING, TELNET


**T****o add VLAN subinterfaces – CLI**

config system interface edit VLAN_100

set vdom root

set interface internal set type vlan

set vlanid 100 set mode static

set ip 10.1.1.1 255.255.255.0

set allowaccess https ping telnet next

edit VLAN_200

set vdom root

set interface internal

end

set type vlan set vlanid 200 set mode static

set ip 10.1.2.1 255.255.255.0

set allowaccess https ping telnet



**A****d****d the firewall addresses**

You need to define the addresses of the VLAN subnets for use in security policies. The FortiGate unit provides one default address, “all”, that you can use when a security policy applies to all addresses as a source or destination of a packet. However, using “all” is less secure and should be avoided when possible.

In this example, the “_Net” part of the address name indicates a range of addresses instead of a unique address. When choosing firewall address names, use informative and unique names.


**T****o add the firewall addresses – web-based manager**

**1****.** Go to **F****i****r****e****w****a****l****l Objects > Address > Addresses**.

**2****.** Select **C****r****ea****t****e New**.

**3****.** Enter the following information and select **O****K**:

**N****a****m****e**                                          VLAN_100_Net

**T****y****p****e**                                            Subnet

**S****ubn****e****t / IP Range**                     10.1.1.0/255.255.255.0

**4****.** Select **C****r****ea****t****e New**.

**5****.** Enter the following information and select **O****K**:

**N****a****m****e**                                          VLAN_200_Net

**T****y****p****e**                                            Subnet

**S****ubn****e****t / IP Range**                     10.1.2.0/255.255.255.0


