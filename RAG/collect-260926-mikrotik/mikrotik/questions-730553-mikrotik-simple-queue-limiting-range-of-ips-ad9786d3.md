---
id: collect-260926-mikrotik/mikrotik/questions-730553-mikrotik-simple-queue-limiting-range-of-ips-ad9786d3
title: "questions-730553-mikrotik-simple-queue-limiting-range-of-ips-ad9786d3"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: ["parameters"]
source: docs/RAG/lot-mikrotik/forum/qos/questions-730553-mikrotik-simple-queue-limiting-range-of-ips-ad9786d3.md
source_anchor: ""
source_lines: [1, 41]
sha256: 9a36127bc61c39e59eb45e7951940eacc64e4391dc5b0ee010842912baa3fa86
---

# questions-730553-mikrotik-simple-queue-limiting-range-of-ips-ad9786d3

Unless you are using PCQ queue type, then the limits are applied to all IPs in the defined range.
With PCQ (Per Connection Queue) you can apply the desired limits on each IP based on the criteria you define on the PCQ (dst/src address, dst/src port, or any combination of those).
From MikroTik official documentation:
  PCQ was introduced to optimize massive QoS systems, where most of the
  queues are exactly the same for different sub-streams. For example a
  sub-stream can be download or upload for one particular client (IP) or
  connection to server.
  
  PCQ algorithm is very simple - at first it uses selected classifiers
  to distinguish one sub-stream from another, then applies individual
  FIFO queue size and limitation on every sub-stream, then groups all
  sub-streams together and applies global queue size and limitation.
  
  PCQ parameters:
pcq-classifier (dst-address | dst-port | src-address | src-port; default: "")  : selection of sub-stream identifiers
pcq-rate (number) : maximal available data rate of each sub-steam
pcq-limit (number) : queue size of single sub-stream (in KB)
pcq-total-limit (number) : maximum amount of queued data in all sub-streams (in KB)
  
  So instead of having 100 queues with 1000kbps limitation for download we can have one PCQ queue with 100 sub-streams 
There is also an example available at MikroTik Wiki
Wiki excerpt:
  There are two ways how to make this: using mangle and queue trees, or,
  using simple queues.
  
  
  - Mark all packets with packet-marks upload/download: (lets constider that ether1-LAN is public interface to the Internet and ether2-LAN is
  local interface where clients are connected /ip firewall mangle add chain=prerouting action=mark-packet in-interface=ether1-LAN new-packet-mark=client_upload
/ip firewall mangle add chain=prerouting action=mark-packet in-interface=ether2-WAN new-packet-mark=client_download
- Setup two PCQ queue types - one for download and one for upload. dst-address is classifier for user's download traffic, src-address for
  upload traffic: /queue type add name="PCQ_download" kind=pcq pcq-rate=64000 pcq-classifier=dst-address  
/queue type add name="PCQ_upload" kind=pcq pcq-rate=32000 pcq-classifier=src-address
- Finally, two queue rules are required, one for download and one for upload: /queue tree add parent=global-in queue=PCQ_download packet-mark=client_download
/queue tree add parent=global-out queue=PCQ_upload packet-mark=client_upload
If you don't like using mangle and queue trees, you can skip step 1,
  do step 2, and step 3 would be to create one simple queue as shown
  here:
/queue simple add target-addresses=192.168.0.0/24 queue=PCQ_upload/PCQ_download
     
    
/queue export# oct/21/2015 15:25:55 by RouterOS 6.23# software id = U3SW-9LU3#/queue simpleadd dst=ether1 max-limit=5M/10M name=Klasat target=172.16.2.0/25add dst=ether1 max-limit=5M/10M name=Administratatarget=172.16.2.128/26add dst=ether1 max-limit=1M/1M name=DVR target=172.16.2.192/27add dst=ether1 name=Sallat target=172.16.2.224/28add dst=ether1 name=Unlimited target=172.16.2.240/28default-smallQueue Type which means that the limit is applied to all IPs. You need to use PCQ type to apply the limits on per-IP basis. Check my answer for more details.
