---
id: collect-260926-mikrotik/mikrotik/arborescence-des-pages-69
title: "Basic Setup"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: ["preemption"]
source: docs/RAG/lot-mikrotik/RouterOS/arborescence-des-pages-69.md
source_anchor: ""
source_lines: [1, 59]
sha256: 1fb0bc1930c9533ba0374494c69247f8846d89cbf5cf9e887a603ff0d4c03b94
---

# Basic Setup

This is the basic VRRP configuration example.

Note

It is recommended to use the same version of RouterOS for all devices with the same VRID used to implement VRRP.

According to this configuration, as long as the master, R1, is functional, all traffic destined to the external network gets directed to R1. But as soon as R1 fails, R2 takes over as the master and starts handling packets forwarded to the interface associated with IP(R1). In this setup router "R2" is completely idle during the Backup period.

## Configuration

R1 configuration:

R2 configuration:

## Testing

First of all, check if both routers have correct flags at VRRP interfaces. On router R1 it should look like this

and on router R2:

As you can see VRRP interface MAC addresses are identical on both routers. Now to check if VRRP is working correctly, try to ping the virtual address from a client and check ARP entries:

Now unplug the ether1 cable on router R1. R2 will become VRRP master, and the ARP table on a client will not change but traffic will start to flow over the R2 router.

In case VRRP is used with Reverse Path Filtering, then it is recommended that `rp-filter` is set to `loose`, otherwise, the VRRP interface might not be reachable.

# Load sharing

In the basic configuration example, R2 is completely idle during the Backup state. This behavior may be considered a waste of valuable resources. In such circumstances, the R2 router can be set as the gateway for some clients. 

The obvious advantage of this configuration is the establishment of a load-sharing scheme. But by doing so R2 router is not protected by the current VRRP setup.

To make this setup work we need two virtual routers.

Configuration for V1 virtual router will be identical to a configuration in basic example - R1 is the Master and R2 is the Backup router. In V2 Master is R2 and Backup is R1.

With this configuration, we establish load-sharing between R1 and R2; moreover, we create a protection setup by having two routers acting as backups for each other.

## Configuration

R1 configuration:

R2 configuration:

# VRRP without Preemption

Each time when the router with a higher priority becomes available it becomes the Master router. Sometimes this is not the desired behavior and can be turned off by setting `preemption-mode=no` in VRRP configuration.

## Configuration

We will be using the same setup as in the basic example. The only difference is during configuration set preemption-mode=no. It can be done easily by modifying the existing configuration:

## Testing

Try turning off the R1 router, R2 will become the Master router because it has the highest priority among available routers.

Now turn the R1 router on and you will see that the R2 router continues to be the Master even if R1 has the higher priority.
