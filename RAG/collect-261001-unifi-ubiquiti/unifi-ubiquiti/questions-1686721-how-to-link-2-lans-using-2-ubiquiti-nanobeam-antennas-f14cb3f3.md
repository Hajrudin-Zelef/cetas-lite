---
id: collect-261001-unifi-ubiquiti/unifi-ubiquiti/questions-1686721-how-to-link-2-lans-using-2-ubiquiti-nanobeam-antennas-f14cb3f3
title: "questions-1686721-how-to-link-2-lans-using-2-ubiquiti-nanobeam-antennas-f14cb3f3"
domain: unifi-ubiquiti
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-unifi-ubiquiti/questions-1686721-how-to-link-2-lans-using-2-ubiquiti-nanobeam-antennas-f14cb3f3.md
source_anchor: ""
source_lines: [1, 26]
sha256: 6a11659b3d74bb1869b3ca98b17a96aae235e12f1bae0d7564608419f0146277
---

# questions-1686721-how-to-link-2-lans-using-2-ubiquiti-nanobeam-antennas-f14cb3f3

We have 2 Ubiquiti antennas model "Nanobeam m5-16-5ghz" I want to link my home network to my sister home, where she has an internet provider. The visibility is perfect and the distance is less than one kilometre.
Her LAN network is like this :
192.168.1.1 : router to internet
192.168.1.20 : Nanobeam antenna [1]
192.168.1.30 and up : local DHCP machines
My LAN network is like this :
192.168.1.1 : Nanobeam antenna [2]
192.168.1.2 : local router providing wifi to my home
192.168.1.30 and up : local DHCP machines
Question 1: one antenna must be "AP" and the other must be "station" or not?
Question 2: both antennas must work in "router" network mode so both LANs are separated?
Supposing Q1 and Q2 answers are "yes", we come to the WLAN settings.
I set my antenna [2] to "station" mode and my sisters antenna [1] to "AP" mode.
I set her SSID to her name, my antenna sees this ESSID and they associate correctly (-45 dBm)
My WLAN settings are easy :
Static IP
IP = 10.139.130.113
GW = 10.139.130.97
Her WLAN settings :
Static IP
IP = 10.139.130.97
GW = 192.168.1.1            <<<<<<<<<<<< ISP router IP [point "3"]
From my antena [2] I can ping hers using Ubuquiti tools.
But I have no access to 192.168.1.1, neither 8.8.8.8
It seems to me the default gateway in antena [1] does not do the trick, point "3".
I would be grateful for any ideas, URLs, improvements, comments, etc.
