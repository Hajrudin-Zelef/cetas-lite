---
id: collect-240926-tomshardware/tomshardware/amd-ryzen-7-9800x3d-review-devastating-gaming-performance-2
title: "amd-ryzen-7-9800x3d-review-devastating-gaming-performance"
domain: tomshardware
role: reference
task: reference
actors: ["AMD", "Intel"]
dates: []
keywords: ["amd", "benchmark", "benchmarks", "chiplet", "compute", "intel", "latency", "memory"]
source: docs/RAG/clean_en/tomshardware/amd-ryzen-7-9800x3d-review-devastating-gaming-performance.md
source_anchor: ""
source_lines: [67, 150]
sha256: 1b1e95c44446156b5e0f2c937376b94fca0f86871f268c065e361c6697d619f7
---

# amd-ryzen-7-9800x3d-review-devastating-gaming-performance

The 9800X3D's 5.2 GHz boost is impressive, but that lags the direct Zen 5 comparable, the eight-core Ryzen 7 9700X, by 300 MHz. However, AMD has compensated with a 900 MHz higher base clock, allowing the 9800X3D to outpace the 9700X in some threaded workloads, as you'll see in our benchmarks on the following pages.


As with the prior-gen models, the Ryzen 7 9800X3D doesn't come with a cooler. AMD recommends at least a 240-280mm liquid (or equivalent) cooler. The processor has the same 95C maximum temperature (TjMax), but the new L3 cache chiplet design allows the chip to operate at higher clock rates for longer periods of time (enhanced boost residency), which equates to stronger performance gains within the same TDP envelope.


The Ryzen 7 9800X3D is fully overclockable — you can tune the CPU cores, fabrics, and memory to your liking. However, while multiplier-based overclocking is available, as we've seen with other Zen 5 processors, most users with conventional cooling will be best served using the auto-overclocking Precision Boost Overdrive (PBO) feature. We have extensive testing with this feature enabled on the following pages.


AMD also bumped memory support up to DDR5-5600 from the DDR5-5200 found with the previous-gen 7800X3D.

## AMD Ryzen 7 9800X3D architecture

In the illustrations above, you can see AMD's original approach with its 3D V-Cache tech. With previous designs that leveraged the Zen 3 and Zen 4 architectures, AMD stacked an additional L3 SRAM chiplet directly in the center of the compute die (CCD) chiplet to isolate it from the heat-generating cores. We’ve covered the details of the first generation of this technology here.


AMD used hybrid bonding technology to join the cache chiplet to the underlying compute die. This boosted cache capacity to 96MB to accelerate performance in latency-sensitive applications, like gaming.


AMD placed the chiplet and several pieces of structural silicon on top of the compute die, but as we demonstrated with thermal throttling tests of the Ryzen 9 7950X3D processor, the cache chiplet and structural silicon shims served as a thermal blanket that hampered effective thermal transfer, essentially trapping waste heat. As a result, AMD limited the effective voltage and frequency range of the 3D V-Cache-equipped die to rein in heat generation. Simply put, the chip ran hot and thus had to run slower. 

AMD has now placed the cache chiplet underneath the die to alleviate the thermal challenges, placing the compute die closer to the thermal interface material, IHS, and eventually the CPU cooler. This improves the amount of compute performance AMD can extract from the 120W TDP envelope, and it also improves boost residency/duration. However, it comes with a new set of challenges — this design requires AMD to route the power and signal TSVs for the entire chip through the underlying L3 cache chiplet.


The new L3 cache chiplet is based on the same 7nm SRAM-optimized process node that AMD used in the previous two generations of 3D V-Cache technology in a technique that AMD now refers to as its 'Second-Gen 3D V-Cache' technology (this is an odd branding choice — this is actually the third refinement of the technology).


As before, AMD thins both the L3 cache chiplet and the 4nm (N4P) compute die to adhere to the standard Z-Height requirements for the processor package. AMD uses the same hybrid copper-to-copper bonding with a 9-micron bump pitch to join the two die. 

Get Tom's Hardware's best news and in-depth reviews, straight to your inbox.

The chiplet now spans the entire bottom of the compute die, thus eliminating the need for the structural silicon shims. AMD isn't sharing the size of the L3 cache chiplet or the transistor density, but the compute die measures 70.6 mm^2, so it's safe to bet that's also the size of the cache chiplet. That's much larger than the 36 mm^2 L3 cache chiplet used for the Zen 4 X3D models — this is despite the use of the same process node and the same 64MB of L3 capacity. AMD likely has a lot of 'empty' silicon on the cache chiplet that's simply there for structural support, without the need to add separate shims.

| 3D V-Cache Technology |  |  |  |  |  |  | 
|---|---|---|---|---|---|---|
| Row 0 - Cell 0 | "New" 2nd-Gen 7nm 3D V-Cache Die | 2nd-Gen 7nm 3D V-Cache Die | First-Gen 7nm 3D V-Cache Die | 4nm Zen 5 Core Complex Die (CCD) | 5nm Zen 4 Core Complex Die (CCD) | 7nm Zen 3 Core Complex Die (CCD) | 
| Size | 36mm^2 | 36mm^2 | 41mm^2 | 70.6 mm^2 | 66.3 mm^2 | 80.7mm^2 | 
| Transistor Count | ? | ~4.7 Billion | 4.7 Billion | 8.6 billion | 6.57 Billion | 4.15 Billion | 
| MTr/mm^2 (Transistor Density) | ? | ~130.6 Million | ~114.6 Million | 121.81 MTr/mm^2 | ~99 Million | ~51.4 Million | 

Naturally, AMD has to feed more power and signal TSVs through the L3 cache chiplet, which now serves as an interposer of sorts, up to the compute die. AMD says that power TSVs are distributed throughout the SRAM die to feed the compute die that resides above. Extra power TSVs are distributed in areas that are unused for other functionality (presumably as more cache storage), which should at least partially account for the larger size.

| AIDA L3 Cache Latency Measurements — Ryzen 7 9800X3D |  |  |  | 
|---|---|---|---|
| Memory Latency - Tom's Hardware | DDR5 | CUDIMM DDR5 | L3 Latency | 
| Ryzen 7 9800X3D | 78.5 ns (DDR5-6000) | n/a | 13.2ns | 
| Ryzen 7 7800X3D | 73.4 ns (DDR5-6000) | n/a | 12.8 ns | 
| Ryzen 7 9700X | 69.4 ns (DDR5-6000) | n/a | 11.3 ns | 
| Core Ultra 9 285K | 94.1 ns (DDR5-5600) | 91.9 ns | 16.6 / 15.8 ns | 
| Core i9-14900K | 79.1 ns (DDR5-5600) | N/A | 21.8 ns | 

AMD says accesses to the L3 cache chiplet incur the same four-clock penalty as with the prior generation. We measured a sub-ns increase in L3 latency compared to the prior-gen 3D V-Cache-equipped Ryzen 7 7800X3D, which falls within the expected variance. AMD says the bandwidth between the compute die and L3 chiplet is similar to the previous-gen at 2.5 TB/s, but bandwidth varies with clock speeds.


Accesses to both main memory and I/O now have to travel through the L3 cache chiplet. As you can see in the measurements above, we see a marked increase in DDR5 latency over the standard Ryzen 7 9700X and the Ryzen 7 7800X3D. AMD says there is no increase in I/O latency due to traveling through the cache chiplet, though. We're following up for more details.


Naturally, any increase in memory latency is undesirable. However, we expect this will have minimal impact on overall performance due to the larger slab of L3 cache that minimizes accesses to main memory in some scenarios. Let's see what performance looks like in gaming on the next page. 

- **MORE:** **Best CPUs for Gaming**
- **MORE:** **CPU Benchmark** **Hierarchy**
- **MORE:** **AMD vs Intel**

- 
Thanks a lot for the review, Paul.Reply
 
 Calling this a (gaming) bloodbath is being mild. Holy cow...
 
 "How screwed is Intel after this?
 Yes"
 
 EDIT: Thanks Jarred for the YT side as well! I was thinking if you were going to do it again :D
 
Regards.
- 
christ, AMD really was understating how good a gaming cpu this was.Reply
 
what a beast. +30% fps gain going AMD x3D over intel. It's like 2012 all over again, only instead of intel styling on amd it's amd styling on intel.
- 
Reply
I think a 30% fps increase team red over blue would make it so that you'd need some serious productivity reasons to even consider intel at this point. furthermore the 9950x exists... and that will out perform intel in gaming as well (not by nearly as much but the productivity will be on par)YSCCC said:Now this is seriously tempting for the gaming only PC for power usage and performance. and it isn't bad if occasionally do production usage either
- 
Have we ever seen such a gaming performance uplift from one gen to the next in the past? I don't recall myself. This is one impressive CPU.Reply
 
