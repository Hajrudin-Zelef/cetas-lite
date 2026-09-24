---
id: collect-240926-storagereview/storagereview/fr-review-unifi-protect-superlink-environmental-sensor-review-reliable-leak-and-e9cfea41
title: "fr-review-unifi-protect-superlink-environmental-sensor-review-reliable-leak-and--e9cfea41"
domain: storagereview
role: reference
task: reference
actors: ["EU", "United States"]
dates: []
keywords: ["ethernet", "latency", "parameters"]
source: docs/RAG/clean_en/storagereview/fr-review-unifi-protect-superlink-environmental-sensor-review-reliable-leak-and--e9cfea41.md
source_anchor: ""
source_lines: [1, 57]
sha256: de6152967d0668405eee437243e1fd3a8d687bca995845239b8c89a1d7845733
---

# fr-review-unifi-protect-superlink-environmental-sensor-review-reliable-leak-and--e9cfea41

<!-- source: https://www.storagereview.com/fr/review/unifi-protect-superlink-environmental-sensor-review-reliable-leak-and-climate-alerts-anywhere -->

Ubiquiti recently expanded its UniFi Protect ecosystem with a new generation of smart sensors using its long-range SuperLink wireless protocol. For this test, Ubiquiti provided us with the UP-SuperLink gateway and the new USL environmental sensor. The gateway is priced at $129 and the sensors at $49 each (affiliate link) on the Ubiquiti store.
Although the environmental sensor focuses on water leaks, temperature, humidity, and ambient light, it represents only one element of the vast SuperLink sensor family.
SuperLink is specially designed for long-distance, low-latency communication with UniFi Protect devices. Ubiquiti presents it as a dedicated and extremely reliable alternative to Wi-Fi or Bluetooth for homes and businesses that need sensors installed in basements, attics, garages, detached buildings, or anywhere traditional wireless signals may be weak.
The constantly expanding Protect sensor range offers motion and glass break detection options, entry and environmental monitoring, and even a SuperLink siren, providing users with a wide range of protection and automation solutions within the Protect ecosystem. Battery-powered and equipped with long-range wireless connectivity, these sensors are designed for simple, wireless installation and increased responsiveness.
SuperLink and Environmental Sensor Specifications
| Specifications | SuperLink Gateway | Environmental Sensor | 
|---|---|---|
| Market |  |  | 
| Dimensions | With antenna: 159.3 x 81.9 x 26 mm DIN rail: 159.3 x 81.9 x 38.5 mm Wall mount: 159.3 x 81.9 x 30.5 mm | Device: 53 x 49 x 23.5 mm Mount: 41.3 x 36.6 x 3.1 mm | 
| Weight | 179 g (6.3 oz) | Device: 47.8 g (1.7 oz) With mount: 70 g (2.5 oz) | 
| Enclosure Material | Polycarbonate, aluminum alloy | Polycarbonate | 
| Mounting | Wall, DIN rail | Adhesive, magnet, screw mount | 
| Weatherproofing | Not rated | IPX5 | 
| Connectivity |  |  | 
| Network Interfaces | (1) 10/100 MbE RJ45 | SuperLink wireless | 
| Operating Frequency | N/A | US: 915.6–927.6 MHz EU: 865.1–869.5 MHz | 
| Antenna Gain | Bluetooth: 3 dBi SuperLink: 1 dBi | 0 dBi | 
| Max TX Power | Bluetooth: 10 dBm SuperLink: 27 dBm | 14 dBm | 
| Max Range | Depends on sensor link | 2 km (1.2 miles) | 
| Hardware |  |  | 
| Power Method | PoE or USB Type C (5V, 1A) | CR123A lithium battery | 
| Max Power Consumption | 3.4 W | 34.9 mW | 
| Battery Life | N/A | Up to 6 years | 
| Direction | Bluetooth 5.2 | UniFi Protect | 
| Sensor Features (Environmental Sensor) |  |  | 
| Temperature | N/A | Yes | 
| Humidity | N/A | Yes | 
| Ambient Light | N/A | Yes | 
| Water Leak Detection | N/A | Yes, with optional probe | 
| Environmental |  |  | 
| Operating Temperature | 0 to 40°C | -20 to 40°C | 
| Operating Humidity | 10 to 90% non-condensing | 10 to 90% non-condensing | 
SuperLink and Environmental Sensor Design and Development
SuperLink
The SuperLink enclosure, made of aluminum alloy and polycarbonate, features a clean, compact, and easy-to-install design. Despite its mounting options, its footprint remains small, allowing for wall, cabinet, or DIN rail installation in structured cabling environments. At only 179 grams, it is lightweight while offering a pleasant and sturdy feel.
The front face is clean, with a simple blue and white status LED for quick and discreet visibility. All input/output connectors are located on the bottom edge: a 10/100 Mbps PoE Ethernet port, an optional USB-C power input, and a recessed factory reset button. This bottom port layout allows cables to be hidden and ensures a clean downward cable path once the device is installed.
The removable antenna, located at the top of the device, screws and unscrews easily. This design gives installers great flexibility in positioning: the antenna can be used for standard mounting, or removed if the device is placed in an enclosure where a compact form factor is preferred.
At the rear, the SuperLink features mounting points for wall and DIN rail mounting, with secure snap-in brackets. These options make it suitable for residential spaces and small industrial settings.
Environmental Sensor
The environmental sensor features a compact, durable polycarbonate enclosure that integrates easily and discreetly into tight spaces. Measuring only 53 x 49 x 23.5 mm, it is small enough for desks, shelves, server rooms, or technical cabinets. It comes with a CR123A lithium-ion battery installed and ready to use.
The front face of the sensor includes two buttons and two LEDs to simplify installation and monitoring. The main status LED uses red, blue, and white lighting to indicate device activity, while the signal LED uses blue and red to indicate SuperLink connection strength. The front-facing function button allows interaction with the sensor during configuration, and a small recessed reset button is available for quick re-provisioning if needed.
The sensor features multiple water detection points integrated around its perimeter, allowing it to detect leaks when placed directly on a surface. For extended coverage, a 3.5 mm auxiliary jack allows connection of an external water leak probe. Inside the enclosure, temperature, humidity, and ambient light sensors are strategically positioned to continuously monitor environmental conditions while being protected from dust and accidental contact.
The mounting system offers great installation flexibility. The galvanized steel bracket allows adhesive, magnetic, or screw mounting, providing options for walls, metal surfaces, or network racks. Even with the bracket installed, the device remains very lightweight (only 70 grams).
UniFi Protect Overview
The SuperLink gateway can be connected to any PoE-compatible UniFi switch port, where it will be automatically powered and appear in the UniFi topology. Once plugged in, it integrates seamlessly with any UniFi gateway running UniFi Protect, allowing it to function like any other UniFi device. In the following example, the SuperLink is visible in the UniFi network topology view, confirming its connectivity and system detection.
When you view the SuperLink in the UniFi Protect device list, you can see its current status, connection quality, and connected sensors. To adopt it, remove the battery tab located at the back of the device and wait a few seconds. The sensor will then immediately appear in the Protect interface as a device available for adoption. After selecting it, the adoption process will complete automatically and the sensor will begin transmitting data without any additional configuration.
Once adopted, the sensor appears as a separate device in UniFi Protect. On its configuration panel, wireless signal strength, firmware version, and overall battery status are displayed instantly. Protect also displays real-time data from each integrated sensor, including water detection, ambient light, humidity level, and temperature. This provides a quick overview of the sensor's status and confirms its active communication with the SuperLink gateway.
For simplified management, the sensor includes an intuitive configuration panel in UniFi Protect. You can, if desired, associate the sensor with a nearby camera to record contextual events, create organization labels, and enable/disable the integrated status light. The sensor also allows you to choose which parameters to monitor, such as temperature, humidity, brightness, and leak detection, for both the integrated sensor and an external probe. You can configure security zones for each parameter to trigger alerts when values exceed defined thresholds. The update interval is adjustable to optimize responsiveness and battery life.
In the Sensor Manager view, Protect provides a clear visual timeline of sensor readings. Temperature, humidity, and brightness can be graphed individually, and the display period can be adjusted to include the last day, last week, or last month. This makes it easy to track changes in environmental conditions and quickly identify trends or anomalies in the monitored area.
Conclusion
At $129 for the SuperLink gateway and $49 for the environmental sensor, Ubiquiti offers a powerful and affordable sensor platform for UniFi Protect users. Installation is simple: any PoE port on a switch can power the gateway, which turns on automatically. The sensor begins transmitting temperature, humidity, brightness, and leak data within minutes, and its CR123A battery ensures extended operation with minimal maintenance.
SuperLink adds value through its long-range wireless connection, performing well even in areas where Wi-Fi or Bluetooth are unstable. This makes it ideal for basements, garages, detached buildings, technical rooms, or any other location requiring constant communication with sensors. Protect presents data through a simple interface, supports camera pairing, and offers adjustable update intervals.
Overall, the SuperLink gateway and environmental sensor provide a reliable and user-friendly solution for extending environmental monitoring in homes and small businesses. For users who have already invested in UniFi Protect, this combination constitutes an economical and well-integrated solution that enhances environmental awareness and protection.
Direct link to the product page (affiliate)
Environmental product page (affiliate)
