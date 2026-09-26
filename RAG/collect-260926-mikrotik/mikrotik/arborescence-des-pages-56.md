---
id: collect-260926-mikrotik/mikrotik/arborescence-des-pages-56
title: "Introduction"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/lot-mikrotik/RouterOS/arborescence-des-pages-56.md
source_anchor: ""
source_lines: [1, 32]
sha256: 503a139a96747721a231e758ff575c19d8daf7fa5567078f61d5bbc2a7c86b07
---

# Introduction

RouterOS uses data from the TZ database, Most of the time zones from this database are included, and have the same names. Because local time on the router is used mostly for timestamping and time-dependent configuration, and not for historical date calculations, time zone information about past years is not included. Currently, only information starting from 2005 is included.

Following settings are available in the **/system clock** console path and in the "Time" tab of the "System > Clock" WinBox window.

Startup date and time is **jan/02/1970 00:00:00** [+|-]gmt-offset. 

# Properties

| Property | Description | 
|---|---|
| **time** (*HH:MM:SS);* | where *HH* - hour 00..24,*MM* - minutes 00..59,*SS* - seconds 00..59). | 
| **date** (*mmm/DD/YYYY);* | where *mmm* - month, one of*jan* ,*feb* ,*mar* ,*apr* ,*may* ,*jun* ,*jul* ,*aug* ,*sep* ,*oct* ,*nov* ,*dec* ,*DD* - date, 00..31,*YYYY* - year, 1970..2037):**date** and**time** show current local time on the router. These values can be adjusted using the**set** command. Local time cannot, however, be exported, and is not stored with the rest of the configuration. | 
| **time-zone-name** (*manual* , or name of time zone; default value:*manual* ); | Name of the time zone. As most of the text values in RouterOS, this value is case sensitive. Special value *manual* applies manually configured GMT offset, which by default is*00:00* with no daylight saving time. | 
| **time-zone-autodetect** (*yes* or*no* ; default: yes); | Feature available from v6.27. If enabled, the time zone will be set automatically. | 

Time-zone-autodetect by default is enabled on new RouterOS installation and after configuration reset. The time zone is detected depending on the router's public IP address and our Cloud servers database. Since RouterOS v6.43 your device will use cloud2.mikrotik.com to communicate with MikroTik's Cloud server. Older versions will use cloud.mikrotik.com to communicate with the MikroTik's Cloud server.

Configuration

## Active time zone information

- **dst-active** (*yes* or*no* >; read-only property): This property has the value*yes* while daylight saving time of the current time zone is active.
- **gmt-offset** ([*+* |*-* ]*HH:MM* - offset in hours and minutes; read-only property): This is the current value of GMT offset used by the system, after applying base time zone offset and active daylight saving time offset.

## Manual time zone configuration

These settings are available in **/system clock manual** console path and in the "Manual Time Zone" tab of the "System > Clock" WinBox window. These settings have an effect only when **time-zone-name**=*manual*. It is only possible to manually configure single daylight saving time period.

- **time-zone** ,**dst-delta** ([*+* |*-* ]*HH:MM* - time offset in hours and minutes, leading plus sign is optional; default value:*+00:00* ) : While DST is not active use GMT offset**time-zone** . While DST is active use GMT offset**time-zone** +**dst-delta** .
- **dst-start** ,**dst-end** (*mmm/DD/YYYY HH:MM: SS* - date and time, either date or time can be omitted in the**set** command; default value:*jan/01/1970 00:00:00* ): Local time when DST starts and ends.
