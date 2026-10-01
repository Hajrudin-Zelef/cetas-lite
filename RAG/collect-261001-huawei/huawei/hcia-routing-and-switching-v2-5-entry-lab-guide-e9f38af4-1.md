---
id: collect-261001-huawei/huawei/hcia-routing-and-switching-v2-5-entry-lab-guide-e9f38af4-1
title: "hcia-routing-and-switching-v2-5-entry-lab-guide-e9f38af4"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: []
keywords: ["copyright", "training"]
source: docs/RAG/collect-261001-huawei/hcia-routing-and-switching-v2-5-entry-lab-guide-e9f38af4.md
source_anchor: ""
source_lines: [1, 263]
sha256: fb6b22745b28b706f5805230a07aae17e946c762021875e4389a6b75340cf3c5
---

# hcia-routing-and-switching-v2-5-entry-lab-guide-e9f38af4

             Recommendations
       Huawei Learning Website
            http://learning.huawei.com/en

       Huawei e-Learning
            https://ilearningx.huawei.com/portal/#/portal/ebg/51



       Huawei Certification
            http://support.huawei.com/learning/NavigationAction!createNavi?navId=_31
             &lang=en

       Find Training
            http://support.huawei.com/learning/NavigationAction!createNavi?navId=_trai
             ningsearch&lang=en


             More Information
   Huawei learning APP




版权所有© 2019 华为技术有限公司
            Huawei Certification




               HCIA-
  Routing&Switching
              ENTRY

Huawei Networking Technology and Device
                 Lab Guide




           Huawei Technologies Co.,Ltd.
Copyright © Huawei Technologies Co., Ltd. 2019.



 All rights reserved.
 Huawei owns all copyrights, except for references to other parties. No part of this
 document may be reproduced or transmitted in any form or by any means without
 prior written consent of Huawei Technologies Co., Ltd.


 Trademarks and Permissions

        and other Huawei trademarks are trademarks of Huawei Technologies Co., Ltd.

All other trademarks and trade names mentioned in this document are the property of
their respective holders.



 Notice


The information in this manual is subject to change without notice. Every effort has
been made in the preparation of this manual to ensure accuracy of the contents, but all
statements, information, and recommendations in this manual do not constitute the
warranty of any kind, express or implied.




                            Huawei Certification

HCIA-Routing&SwitchingHuawei Networking Technology

                                   and Device

                                Entry Lab Guide

                                   Version 2.5
                     Huawei Certification System

Relying on its strong technical and professional training and certification system
and in accordance with customers of different ICT technology levels, Huawei
certification is committed to providing customers with authentic, professional
certification, and addresses the need for the development of quality engineers that
are capable of supporting Enterprise networks in the face of an ever changing ICT
industry. The Huawei certification portfolio for routing and switching (R&S) is
comprised of three levels to support and validate the growth and value of customer
skills and knowledge in routing and switching technologies.


The Huawei Certified Network Associate (HCIA) certification level validates the skills
and knowledge of IP network engineers to implement and support small to
medium-sized enterprise networks. The HCIA certification provides a rich
foundation of skills and knowledge for the establishment of such enterprise
networks, along with the capability to implement services and features within
existing enterprise networks, to effectively support true industry operations.


HCIA certification covers fundamentals skills for TCP/IP, routing, switching and
related IP network technologies, together with Huawei data communications
products,   and skills for versatile routing platform (VRP)          operation and
management.


The Huawei Certified Network Professional (HCIP-R&S) certification is aimed at
enterprise network engineers involved in design and maintenance, as well as
professionals who wish to develop an in depth knowledge of routing, switching,
network efficiency and optimization technologies. HCIP-R&S consists of three units
including Implementing Enterprise Routing and Switching Network (IERS),
Improving Enterprise Network Performance (IENP), and Implementing Enterprise
Network Engineering Project (IEEP), which includes advanced IPv4 routing and
switching technology principles, network security, high availability and QoS, as well
as application of the covered technologies in Huawei products.


The Huawei Certified Internet Expert (HCIE-R&S) certification is designed to imbue
engineers with a variety of IP network technologies and proficiency in maintenance,
for the diagnosis and troubleshooting of Huawei products, to equip engineers with
in-depth competency in the planning, design and optimization of large-scale IP
networks.
Reference Icons
                                                       CONTENTS

MODULE 1 ESTABLISHING BASIC NETWORKS WITH ENSP ......................................... 1


  LAB 1-1 BUILDING BASIC IP NETWORKS ..................................................................................................1


MODULE 2 BASIC DEVICE NAVIGATION AND CONFIGURATION .............................14


  LAB 2-1 BASIC DEVICE NAVIGATION AND CONFIGURATION ............................................................... 14


MODULE 3 STP AND RSTP..............................................................................................25


  LAB 3-1 CONFIGURING STP ................................................................................................................... 25


  LAB 3-2 CONFIGURING RSTP ................................................................................................................ 42


MODULE 4 ROUTING CONFIGURATION .......................................................................51


  LAB 4-1 CONFIGURING STATIC ROUTES AND DEFAULT ROUTES ........................................................ 51


  LAB 4-2 OSPF SINGLE-AREA CONFIGURATION ................................................................................... 68


MODULE 5 FTP AND DHCP ............................................................................................84


  LAB 5-1 CONFIGURING FTP SERVICES .................................................................................................. 84


  LAB 5-2 IMPLEMENTING DHCP ............................................................................................................. 92
       Module 1 Establishing Basic Networks with eNSP

Lab 1-1 Building Basic IP Networks


Learning Objectives


As a result of this lab section, you should achieve the following tasks:

      Set up and navigate the eNSP simulator application.

      Establish a simple peer-to-peer network in eNSP.

      Perform capture of IP packets using Wireshark within eNSP.




                                      HUAWEI TECHNOLOGIES                  Page1
The fundamental network behavior can be understood through the application of
packet capture tools to the network. The use of Huawei’s simulator platform eNSP
is capable of supporting both the implementation of technologies and the capture
of packets within the network to provide a comprehensive knowledge of IP
networks.


Tasks


Step 1 Install eNSP


1. Login website of eNSP：

   https://support.huawei.com/enterprise/en/tool/ensp-TL1000000015/23917110




2. Download the latest version of eNSP




3. Please refer to the software installation guide below to install eNSP in local PC.




   Then engineer can practice lab with AR, Router, S57, S37, USG5500, AC, AP .



                                      HUAWEI TECHNOLOGIES                    Page2
  If the engineer want to practice lab with USG6000V, CE, NE40, NE5000E, NE9000,

  CX, please follow Step4




4. Enable USG6000V, CE, NE40, NE5000E, NE9000, CX devices in eNSP:

  1） For example, if you want to enable USG6000V in eNSP, you should download

       the corresponding mirror file.




  2） Select USG6000V into new project of eNSP, then right click “start”of

       USG6000V :




  3） The dialog box of “import package” will show up:



                                    HUAWEI TECHNOLOGIES               Page3
   4） Click "Browse" - and import the downloaded mirror files, then engineer can

       practice lab with USG6000V.




   5） If the engineer want to practice CE, NE40, NE5000E, NE9000, CX, please

