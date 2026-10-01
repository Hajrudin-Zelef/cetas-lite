---
id: collect-261001-cisco/cisco/enterprise-en-understaing-stack-in-pratice-thread-667248013413859328-66721385295-1ea706e3
title: "enterprise-en-understaing-stack-in-pratice-thread-667248013413859328-66721385295-1ea706e3"
domain: cisco
role: reference
task: reference
actors: ["Huawei"]
dates: []
keywords: []
source: docs/RAG/collect-261001-cisco/enterprise-en-understaing-stack-in-pratice-thread-667248013413859328-66721385295-1ea706e3.md
source_anchor: ""
source_lines: [1, 27]
sha256: 7b7ec718141bd04d31f19c9c5357ac117855bf18599df1f03744a35b5c7057bf
---

# enterprise-en-understaing-stack-in-pratice-thread-667248013413859328-66721385295-1ea706e3

Hello everyone! 
Today, I want to bring a post about stacking on huawei switches, I see a lot of questions on the forum on how to do it, and I would like to introduce a stack configuration assistance tool.
The purpose of this post is to introduce Huawei's iStack setup wizard and bring a practical example on real equipment.
Stack & SVF Assistant provides iStack, CSS and SVF deployment guide, including connection modes and configuration procedures.
You can find it at the following link:
https://info.support.huawei.com/network/virtual/index?lang=en
So that the assistant can help you, you must fill in the fields according to the availability of equipment for your project or task.
In the practical example I bring, I will use the switch model S5735-L24T4X-A and dedicated stack cable.
Run the system-view command to enter the system view.
Run the interface stack-port member-id/port-id command to create a logical stack port and enter the logical stack port view.
Run the port interface { interface-type interface-number1 [ to interface-type interface-number2 ] } &<1-10> enable to configure service ports as physical member ports and add the ports to the logical stack port.
Run the quit command to return to the system view.
(Optional) Run the stack slot slot-id renumber new-slot-id command to configure a stack ID for the member switch.
(Optional) Run the stack slot slot-id priority priority command to configure a stack priority for the member switch.
The following figure shows the appearance of one switch model among a switch sub-series, which may be different from the selected product model. The figure shows the locations of ports that can be used as stack ports on the same switch sub-series and illustrates how to connect these ports in different stack connection modes. Port locations, appearance, and stack connection modes of different switch models are similar.
The bend radius of optical fibers or cables must be larger than the minimum bend radius. The minimum bend radius of SFP+ cables is 25 mm; the minimum bend radius of AOC cables is 30 mm; the bend radius of optical fibers is generally greater than or equal to 40 mm.
Stack member ports of a logical stack port (stack-port n/1) on one switch must be connected to stack member ports of a logical stack port (stack-port m/2) on another switch.
Check whether the stack has been established successfully.
Run the display device command to check whether the number of member switches in the stack is the same as the number of switches in the networking.
Check whether the stack topology is the same as the actual hardware connections.
Run the display stack command to check whether the stack topology is the same as the actual hardware connections.
Run the display stack peers command to check whether neighboring information about the stack is the same as the actual hardware connections.
Run the display stack channel all command to check the stack link connections and status.
If they are the same, the stack has been established successfully.
If they are different, reconnect the member switches.
If the fault persists, rectify the fault according to Rectifying Common Stack Faults.
- END -
