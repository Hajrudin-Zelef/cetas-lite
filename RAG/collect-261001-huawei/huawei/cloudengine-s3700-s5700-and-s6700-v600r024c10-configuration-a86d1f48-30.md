---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-a86d1f48-30
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--a86d1f48"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--a86d1f48.md
source_anchor: ""
source_lines: [3511, 3647]
sha256: 969dd4dab1a3fc8327501cb75a2f899f6cadfee67c8c0a97cfbcc09f12b28acc
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--a86d1f48

Issue 01 (2025-03-03)                 Copyright © Huawei Technologies Co., Ltd.                                       60
Security Configuration
Security Configuration                                                            4 Storm Suppression Configuration

                         storm control { broadcast | multicast | unknown-unicast } min-rate percent min-rate-value-
                         percent max-rate percent max-rate-value-percent

         Step 5 Configure a storm control action.
                  1.     Configure the storm control action as error-down, block, or suppress.
                         storm control action { error-down | block | suppress }

                  2.     (Optional) Enable an interface to automatically go up before the interface
                         enters the error-down state, and set the recovery delay (this step is only
                         applicable to the error-down action).
                         This function should be used if a large number of interfaces may enter the
                         error-down state, as manually recovering so many error-down interfaces is
                         time-consuming and error-prone. To prevent this problem, you can configure
                         interfaces in error-down state to automatically go up, and set the recovery
                         delay.
                               NOTE

                             This method does not take effect on interfaces that are already in error-down state. It
                             takes effect only on interfaces that enter the error-down state after the error-down
                             auto-recovery command is run.
                             You can run the display error-down recovery command to check information about
                             automatic interface recovery.
                         error-down auto-recovery cause storm-control interval interval-value

         Step 6 Configure the storm detection interval.
                  storm control interval interval-value

         Step 7 (Optional) Enable the device to record logs or report traps during storm control.
                  storm control enable { log | trap }

                  ----End

Verifying the Configuration
                  Run the display storm control [ interface interface-type interface-number
                  [ verbose ] ] command to check storm control information on an interface.

Follow-up Procedure
                  After the storm control action is set to error-down on an interface, the interface is
                  shut down when the average rate of received broadcast packets, unknown
                  multicast packets, or unknown unicast packets exceeds the specified upper
                  threshold within the detection interval. If an interface is in error-down state, run
                  the shutdown and undo shutdown commands in the interface view, or run the
                  restart command to restart the interface.

4.5.3 Example for Configuring Storm Control
Networking Requirements
                  On the network shown in Figure 4-2, DeviceA connects a Layer 2 network to a
                  Layer 3 device. Storm control needs to be configured on DeviceA so as to limit the
                  number of broadcast packets, unknown multicast packets, and unknown unicast
                  packets forwarded at Layer 2, preventing broadcast storms.

Issue 01 (2025-03-03)              Copyright © Huawei Technologies Co., Ltd.                                          61
Security Configuration
Security Configuration                                                                       4 Storm Suppression Configuration


                  Figure 4-2 Networking diagram of storm control
                          NOTE

                         In this example, interface 1 represents 10GE1/0/1.




Configuration Roadmap
                  The configuration roadmap is as follows:
                  ●      Configure storm control in the view of 10GE1/0/1 to prevent broadcast storms
                         caused by broadcast packets, unknown multicast packets, and unknown
                         unicast packets forwarded at Layer 2.
                  ●      Enable the device to record logs during storm control to remind a network
                         administrator to take measures to protect the device.

Procedure
         Step 1 Enter the interface view.
                  <HUAWEI> system-view
                  [HUAWEI] sysname DeviceA
                  [DeviceA] interface 10ge 1/0/1
                  [DeviceA-10GE1/0/1] portswitch

         Step 2 Configure the upper and lower thresholds for broadcast packets, unknown
                multicast packets, and unknown unicast packets. When the average rate at which
                an interface receives one of these types of packets is greater than 2000 pps within
                the storm detection interval, storm control is performed on the packets of the
                corresponding type on the interface. When the average rate drops below 1000 pps,
                the interface starts to forward packets of the corresponding type again.
                  [DeviceA-10GE1/0/1] storm control broadcast min-rate 1000 max-rate 2000
                  [DeviceA-10GE1/0/1] storm control multicast min-rate 1000 max-rate 2000
                  [DeviceA-10GE1/0/1] storm control unknown-unicast min-rate 1000 max-rate 2000

         Step 3 Set the storm control action to block.
                  [DeviceA-10GE1/0/1] storm control action block

         Step 4 Set the storm detection interval to 90s.
                  [DeviceA-10GE1/0/1] storm control interval 90

         Step 5 Enable the device to record logs during storm control.
                  [DeviceA-10GE1/0/1] storm control enable log
                  [DeviceA-10GE1/0/1] quit

                  ----End

Verifying the Configuration
                  # Check the storm control configuration.
                  [DeviceA] display storm control interface 10ge 1/0/1
                  --------------------------------------------------------------------------------
                  NOTE:
                  BC = Broadcast; MC = Multicast; UUC = Unknown Unicast
                  Int = Interval value (unit: seconds)


Issue 01 (2025-03-03)               Copyright © Huawei Technologies Co., Ltd.                                              62
Security Configuration
Security Configuration                                                                    4 Storm Suppression Configuration

                  --------------------------------------------------------------------------------
                  PortName        Type MaxRate Mode Action Punish- Trap Log Int Last
                                                    Status                Punish-Time
                  --------------------------------------------------------------------------------
                  10GE1/0/1 BC             2000 Pps Block        Normal Off On 90 --
                  10GE1/0/1 MC              2000 Pps Block        Normal Off On 90 --
                  10GE1/0/1 UUC             2000 Pps Block        Normal Off On 90 --

                  The Punish-Status field displays the packet status on the current interface, and
                  the Last Punish-Time field displays the time when the storm control penalty was
                  last implemented. The preceding output shows that the broadcast packets,
                  unknown multicast packets, and unknown unicast packets on 10GE1/0/1 of
                  DeviceA are normal and that no storm control penalty is performed. This indicates
                  that the average rate of these packets within the detection interval does not
                  exceed the specified value and the network runs properly.

