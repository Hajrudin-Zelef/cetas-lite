---
id: collect-240926-storagereview/storagereview/fr-review-dynatron-2u-aio-cpu-closed-liquid-loop-review-9b3982fe
title: "fr-review-dynatron-2u-aio-cpu-closed-liquid-loop-review-9b3982fe"
domain: storagereview
role: reference
task: reference
actors: ["AMD", "Intel"]
dates: []
keywords: ["amd", "consumer", "energy", "gpus", "intel", "liquid cooling"]
source: docs/RAG/clean_en/storagereview/fr-review-dynatron-2u-aio-cpu-closed-liquid-loop-review-9b3982fe.md
source_anchor: ""
source_lines: [1, 32]
sha256: f1fedd478378295bd947c4ebd00b428f78baf4607d6eac0522741a87ce067f9c
---

# fr-review-dynatron-2u-aio-cpu-closed-liquid-loop-review-9b3982fe

<!-- source: https://www.storagereview.com/fr/review/dynatron-2u-aio-cpu-closed-liquid-loop-review -->

When it comes to cooling servers and workstations, Dynatron is a reference. In addition to its air coolers, Dynatron also offers many liquid cooling solutions. You may have already seen their name, or even unbranded coolers, in machines like the 45Homelab HL15 or the TYAN Transport HX FT65T-B8050 that we tested. Dynatron manufactures and supplies these coolers to manufacturers (OEMs) for their systems.
We received a custom Dynatron all-in-one liquid cooling system to test in the lab. To do this, we installed it in the Tyan HX FT65T-B8050 chassis that we tested a few weeks ago. This system replaces the original Dynatron J10 air cooler. In our configuration, we removed the fans from the radiator and positioned the latter against them, in the center of the chassis.
This AIO system consists of Dynatron's L35 radiator and their SP5 cold plate from their L32, assembled in a custom configuration to fit the beast that is the EPYC 9684X.
Dynatron AIO Specifications
It is difficult to obtain the exact specifications of this unit because it is a custom configuration and there is no formal datasheet for this configuration. The specifications below pertain to the L35 radiator and fan configuration.
| Categories | Information | 
|---|---|
| Server applications | 2U and above servers, tower servers | 
| Fan dimensions | 80 x 80 x 38 mm (3.15 x 3.15 x 1.5 inches) | 
| Radiator assembly dimensions | 323.2 x 44 (82 with fans) x 85.15 mm (12.72 x 1.73 (3.23 with fans) x 3.35 inches) | 
| Fan speed |  | 
| Power consumption |  | 
| Noise level |  | 
| Airflow |  | 
| Air pressure |  | 
Closed-loop liquid applications
Closed-loop liquid cooling systems can be an attractive option for servers, as they offer reduced noise levels and greater footprint flexibility than other solutions. OSS took advantage of this flexibility in its Gen 5 SDS configuration presented at SC23, which we discussed last November. The use of a Dynatron SP5 liquid cooling system allowed OSS to position the power supplies in the location originally intended for an air cooler, thereby freeing up space for GPUs while respecting their form factor.
Regarding noise, we encountered issues on the Tyan where the small fan of the J10 air cooler screamed under load. However, once the liquid cooler was installed, we noticed a much lower volume under load. The CPU fan, which was once tied to CPU load in terms of noise, completely disappeared, as we used the existing chassis fans for radiator airflow. Thus, in our system, noise levels dropped dramatically.
Dynatron Loop Performance
For performance advantages, this is where things are biased between air coolers and liquid coolers. Good air coolers work very well, and so do liquid coolers. When both are designed to follow the same specifications in terms of the amount of thermal energy they can move from a given area, the results will be fairly similar between the two.
In our case, we performed tests on the same hardware, swapping only the cooler, and we found minimal change. We loaded the CPU with an intensive HPC workload and compared the average core clock speed between the two coolers.
One could say that it may have changed things by a very small percentage, but it remains within the margin of error of these tests. The advantage of liquid cooling is not always better performance, but it gives the customer greater flexibility in terms of managing heat, noise, and power consumption in different environments.
Other Dynatron Offerings
As mentioned earlier, Dynatron covers far more than CPU air and liquid coolers; they have their name all over the cooling market. Dynatron offers other products such as passive heat sinks, custom liquid loop parts, RGB parts, blower fans, and chassis fans. These different options allow OEM manufacturers to obtain all their cooling needs from a single supplier, as well as the ability to have custom cooling solutions.
The CPU cooler options offered by Dynatron cover multiple form factors and shapes on the same socket to offer you great flexibility in your configurations. Dynatron also offers options that cover tons of sockets not only for server processors but also for desktop processors. AMD socket options are FM1, FM2(+), AM2(+), AM3(+), AM4, AM5, C32, G34, Opteron 6000 and 6100, socket F, SP3, SP5, SP6, sWRX8, sTRX4, TR4, and TR5. For Intel processors, Dynatron covers sockets LGA1200, 115x, 1356, 1366, 1700, 1851, 2011 narrow and square, 2066 narrow and square, 3647 narrow and square, 4677, 4710 and 7529, PGA479 and PGA988.
For liquid cooling systems, Dynatron offers products suitable for consumer sockets like AM5, as well as server sockets like SP5 for AMD processors. Dynatron also offers a wide range of solutions for Intel processors. The AM5 socket has several cooling options, from the single 120 mm fan radiator of the L5 to the L25-u, equipped with five 40 mm fans for a 1U rackable chassis.
Currently, the SP5 socket is only offered with the L32 cooler which is designed to support up to 500 W of TDP in only 1U, but we expect to see more options in the future. Dynatron will likely offer different form factors such as 2U radiators, and perhaps up to 4U or larger tower sizes for this socket in the future, similar to what they have done with other sockets.
All these different Dynatron options give you coverage almost everywhere you need it in the cooling market.
Conclusion
Overall, we may not have seen huge thermal differences with this liquid cooler, but we did see a reduction in noise. Using a water cooler also allows you to move the cooler to other areas of the chassis to free up space at the top. Another advantage of the liquid cooler is that you have to worry less about airflow routing since the radiator can cover a larger area. Thanks to the variety of cooling options offered by Dynatron, they cover many applications on the market. If you have server CPU cooling needs, Dynatron has probably crossed your radar with at least one of its products.
