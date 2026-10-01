---
id: collect-261001-cisco/cisco/c-en-us-td-docs-security-asa-asa-cli-reference-a-h-asa-command-ref-a-h-ca-cld-co-fc223d78-7
title: "c-en-us-td-docs-security-asa-asa-cli-reference-a-h-asa-command-ref-a-h-ca-cld-co-fc223d78"
domain: cisco
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-cisco/c-en-us-td-docs-security-asa-asa-cli-reference-a-h-asa-command-ref-a-h-ca-cld-co-fc223d78.md
source_anchor: ""
source_lines: [822, 952]
sha256: a14a656b2b90c6cafdeb012c45b548e0e001013405293c57048be6eb8c6a271b
---

# c-en-us-td-docs-security-asa-asa-cli-reference-a-h-asa-command-ref-a-h-ca-cld-co-fc223d78

                              Because all the packets that are forwarded or dropped by the ASA hits the two front-end network processors, the packet capture
                              feature is implemented in these network processors. So all the packets that hit the ASA can be captured by these front end
                              processors, if an appropriate capture is configured for those traffic interfaces. On the ingress side, the packets are captured
                              the moment the packet hits the ASA interfaces, and on the egress side the packets are captured just before they are sent out
                              on the wire. 
                           
                           
                           
                              
                                 |  Note |  Enabling WebVPN capture affects the performance of the ASA. Be sure to disable the capture after you generate the capture                                           files that you need for troubleshooting.                                          | 
                           
                                 Save the Capture
                           
                           
                              The contents of any active capture on ASA are saved when the box crashes. 
                           
                           
                           
                              When you activate captures as part of the troubleshooting process, you must note the following points:
                           
                           
                           
                              
                              - 
                                 
                                    The size of capture buffer to use and if there is enough space on flash/disk.
                                 
- 
                                 
                                    The capture buffer should be marked as circular for all the use cases, so that captured packets are the most recent before
                                    crash.
                                 
                              The name of the file for saving contents of an active capture is in the format of:
                           
                           
                           
                              [<context_name>.]<capture_name>.pcap
                              
                           
                           
                           
                              The context_name
                              indicates the name of the user context in which capture is activated in the multi-context mode. For the single context mode,
                              the context_name
                              is not applicable.
                           
                           
                           
                              The capture_name
                              indicates the name of the capture that is activated.
                           
                           
                           
                              The capture save happens before the console or crash dump. This increases the crash downtime by about 5 seconds for a 33 MB
                              capture buffer. The risk of a nested crash is minimal because copying the captured contents to a file is a simple process.
                           
                           
                           
                                 View the Capture
                           
                           
                              
                              - 
                                 
                                 To view the packet capture at the CLI, use the show
                                       capture 
                                    name command.
                                 
- 
                                 
                                 To save the capture to a file, use the copy
                                       capture  command. 
                                 
- 
                                 
                                 To see the packet capture information with a web browser, use the
                                    https://ASA-ip-address/admin/capture/capture_name[/pcap] 
                                    command. 
                                  You are prompted for a username and password. See the username 
                                    command to add a username to the local database.
                                  If you specify the pcap  keyword, then a libpcap-format file is
                                    downloaded to the web browser and can be saved using the web browser. (A libcap file can be
                                    viewed with TCPDUMP or Ethereal.)
                                 
If you copy the buffer contents to a TFTP server in ASCII format, you will see only the
                              headers, not the details and hexadecimal dump of the packets. To see the details and hexadecimal
                              dump, you need to transfer the buffer in PCAP format and read it with TCPDUMP or Ethereal.
                           
                           
                           
                                 Stop and Start the Capture
                           
                           
                              The packets can be stopped from being captured without removing them from the buffer. The stopped status of the capture is
                              displayed. The captured packet is retained in the buffer.
                           
                           
                           
                              Use the following command to manually stop packet capture:
                           
                           
                           capture  name
                              stop 
                           
                           
                              Use the following command to start capturing packets:
                           
                           
                           no capture  name
                              stop 
                           
                           
                                 Delete the Capture
                           
                           
                              Entering no capture  without any keywords deletes the capture. To preserve the capture, specify the access-list  or interface  keyword; the capture is detached from the specified ACL or interface and the capture is preserved.
                           
                           
                           
                                 Real Time Operations
                           
                           
                              You cannot perform any operations on a capture while the real-time display is in progress. Using the real-time  keyword with a slow console connection may result in an excessive number of non-displayed packets because of performance
                              considerations. The fixed limit of the buffer is 1000 packets. If the buffer fills up, a counter is maintained of the captured
                              packets. If you open another session, you can disable the real-time display be entering the no capture real-time  command.
                           
                           
                           
                                 Clustering
                           
                           
