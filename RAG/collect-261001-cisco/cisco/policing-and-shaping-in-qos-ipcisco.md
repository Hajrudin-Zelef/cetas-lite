---
id: collect-261001-cisco/cisco/policing-and-shaping-in-qos-ipcisco
title: "policing-and-shaping-in-qos-ipcisco"
domain: cisco
role: reference
task: reference
actors: ["Huawei"]
dates: []
keywords: []
source: docs/RAG/collect-261001-cisco/policing-and-shaping-in-qos-ipcisco.md
source_anchor: ""
source_lines: [1, 77]
sha256: 3acc473303a202ec46bf54f836fdacb343d3fea71bf375fe191b90bc9b0c28b7
---

# policing-and-shaping-in-qos-ipcisco

In networks, the traffic increses for a variety of reasons. Some **Quality of Service Adjustments** are done to handle this increases. During these increases, if the capacity of the link is exceeded, then the traffic on this link is dropped. Without any arrangement, this drop can cause unexpected and undesired important data drops. To avoid such problems, **QoS Traffic Policing and Shaping** mechanisms are used. This mechanisms are another important mechanisms beside **QoS Classification and QoS Marking**.


Basically, **Traffic QoS Policing** is using a predefined **Traffic Policy** to manage the network traffic. With these configured Traffic Policies, during a bandwidth exceed, the required and desired actions are ordered according to the traffic variety. With this policies, traffics are remarked or they are dropped with a predefined actions.


On the other hand, **Traffic Shaping** is not as strict as Traffic QoS Policing. With Traffic Shaping, the traffic is controlled and if the traffic close to the traffic exceed, it uses queues and made some of the traffic wait (buffered the traffic) without any drop action.



There are also some important terms used with Traffic Policing and Traffic ShapingThese terms are given below:

**CIR :** Committed Information Rate

**EIR :** Exceed Information Rate

**PIR :** Peak Information Rate


Traffic is controlled according to these values. Now, let’s briefly check what are these value. You can also reach the detailed explanation of these values in another article.


**CIR (Committed Information Rate)** is the traffic rate that is guaranteed by the provider. The traffic rate under this value is delivered certainly. After this value, to the PIR value, the traffic is delivered but not guaranteed. During busy times, these traffic can be dropped.


Beside, **PIR (Peak Information Rate)** is the top traffic rate that the traffic will delivered in the normal traffic flows. Any traffic that exceed this value, certainly policied or shaped.


**EIR (Exceed Information Rate)** is the value, between CIR and PIR. This is the amouth of unguaranteed but delivered traffic.


Now, let’s talk about **Traffic QoS Policing** and **QoS Traffic Shaping** detailly.


Table of Contents

**Traffic QoS Policing** is the strict method of controlling bandwidth exceed according to high traffic. For Quality of Service, some predefined actions are determined by the network administrators. This predefined actions build a Traffic Policy. With Traffic Policy, network knows what to do even if a traffic violation occurs.


For example, in the **Traffic Policy**, if a bandwidth exceed occurs, you can define to drop a video streaming traffic. At that moment, users can not watch videos because of the limitation of Traffic Policy. This Traffic limitation provides extra bandwidth in the network and network bandwith exceed is overcomed. This limitation is also done with remarking the traffic.



Here, we can say that there are two key techniques used in **Traffic Policing**. These techniques are:


**Dropper** is the mechanims that drops the traffic if a traffic violation occurs as its name implies.

Marker is the mechsnism that remarks the traffic, with a new priority.


**Traffic Policies** are configured and can be used both inbound and outbound directions. So, both in coming and outgoing traffic can be controlled with Traffic Policies.


Generally **Traffic QoS Policing** causes more TCP retransmissions. This provide more resource utilization.


**QoS Traffic Shaping** is the mechanism that is used during a **Traffic Exceed**. The Exceed Traffic is **buffered** and with a little delay, it is sent again. With this buffer and queue mechanism, **Traffic Shaping** is a soft medhod if we compare with Traffic QoS Policing.


There can be **Traffic Exceeds (traffic bursts)** in the network. With **Traffic Shaping**, this burst become smooth. Think about this graphic. The top of the graphic seems as it is cut with a pair of scissors. The traffic continue to be at top border, but the exceed traffic buffered.

**QoS Traffic Shaping** do not support Marking or Remarking again. This is done only with Traffic Policing. **QoS Traffic Shaping** is applied only through outgoing interfaces. The buffer mechanims used by **Traffic Shaping** minimize TCP retransmissions. This provide less resource utilization.


In **Cisco CCNA Exam** and **CCNP ENCOR Certification Exams**, you can see questions about Traffic Shaping and Policing.


Gokhan Kosem is a Network Engineer, Instructor and the Founder of IPCisco.com with 15+ years of experience in Cisco, Nokia, Huawei, Juniper, Linux, Service Provider Networks, Routing and Switching technologies.

He has worked on the backbone networks of major service providers and network vendors including Nortel, Alcatel-Lucent (Nokia) and has extensive hands-on experience with Cisco, Huawei, Juniper and Nokia networking technologies.

He has trained thousands of networking students worldwide through IPCisco.com, Udemy, books, labs, quizzes, and educational content across multiple social media platforms.

IPCisco.com | Best Route to Your Dreams

## Leave a Reply
