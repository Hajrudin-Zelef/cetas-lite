---
id: collect-240926-tomshardware/tomshardware/amd-ryzen-9-9950x3d2-review-more-cache-more-cash-1
title: "amd-ryzen-9-9950x3d2-review-more-cache-more-cash"
domain: tomshardware
role: reference
task: reference
actors: ["AMD", "Intel"]
dates: []
keywords: ["amd", "benchmark", "benchmarks", "consumer", "gpu", "intel", "latency", "packaging", "pricing"]
source: docs/RAG/clean_en/tomshardware/amd-ryzen-9-9950x3d2-review-more-cache-more-cash.md
source_anchor: ""
source_lines: [1, 70]
sha256: 2d232798cac71917236f426b0b5172c029e550bb3245bbcf8ce0dad7b1c7a12c
---

# amd-ryzen-9-9950x3d2-review-more-cache-more-cash

<!-- source: https://www.tomshardware.com/pc-components/cpus/amd-ryzen-9-9950x3d2-review -->

### Tom's Hardware Verdict

The Ryzen 9 9950X3D2 isn’t worth the money for the vast majority of people, but it was never meant to be. It’s a halo product with some surprising upsides in a few niche workloads, and it stands as AMD’s first-ever dual-3D V-Cache CPU.

#### Pros

- +Matches the Ryzen 7 9800X3D in games
- +Around 4% faster in multithreaded performance compared to 9950X3D
- +Double-digit improvements in some specialized workloads
- +Unlocked multiplier for overclocking

#### Cons

- -Very expensive
- -Slight regressions in single-threaded workloads
- -Higher power consumption

Given the prevalence of AMD’s dual-CCD 3D V-Cache-equipped X3D CPUs among the __best CPUs for gaming__ and our __CPU benchmark hierarchy__, a natural question has formed in the minds of PC enthusiasts since the first model launched: What if you stacked cache on both CCDs? AMD has restricted its wildly popular (and marketable) 3D V-Cache to a single CCD on its Ryzen 9 models, leaving a natural slot for a CPU that stacked cache on both CCDs. The Ryzen 9 9950X3D2 Dual Edition finally brings that idea to the company’s long-lived AM5 platform.

It’s a fascinating CPU, and not just because it’s the first of its kind. For starters, it’s the most expensive Ryzen CPU AMD has ever released — Intel messed around with $1,000+ consumer CPUs years ago — with a recommended retail price of $899. Anything above that price goes into the Threadripper range. It’s also a halo product launching in a PC market that’s plagued by inflated prices, from __ongoing RAM and SSD price increases__ to the ever-present climb of GPU prices.

Instead of just offering a binned version of the 9950X3D, similar to what we saw with the __Ryzen 7 9850X3D__, AMD is bridging the gap between its consumer and HEDT ranges with the 9950X3D2. It’s firmly a workstation-class processor, with little to no advantage in games, and even slight regressions in some titles. There are some minor advantages in application performance, particularly in heavily-threaded workloads like rendering and encoding, though not enough to justify the $899 price tag.

Even the packaging here hints that this is a highly-specialized CPU, sporting a neutral black and gray color scheme in stark contrast to the vibrant oranges and reds we typically see with AMD chips.

The performance improvements show up mainly in specialized workloads in fields like data science, where massive instructions can show the latency benefits of having L3 access easily available on both CCDs. Looking at overall performance, the Ryzen 7 9850X3D offers a better gaming experience while the 9950X3D delivers a greater value in applications. But for these specialized workloads, AMD is offering performance increases as large as 25%.

Our full benchmarks provide a look into those advantages, spanning gaming, productivity, and power testing. The Ryzen 9 9950X3D2 was never meant to be a mass market CPU, and our testing goes further to prove that it’s a niche product for a small audience. Take 100 PC builders, and we’d recommend another chip to 99 of them. For that small market, however, the 9950X3D2 delivers.

## AMD Ryzen 9 9950X3D2 specifications and pricing

| **CPU / (MSRP)** | **Street Price** | **Architecture** | **Cache (L2 + L3)** | **Cores/Threads (P+E)** | **Base/Boost Clock (GHz)** | **TDP / Maximum Power** | 
| **Ryzen 9 9950X3D 2 Dual Edition ($900)** | **$900** | **Zen 5 X3D** | **208 MB** | **16 / 32** | **4.3 / 5.6** | **200W / 270W** | 
| Ryzen 9 9950X3D ($700) | $676 | Zen 5 X3D | 144 MB | 16 / 32 | 4.3 / 5.7 | 170W / 230W | 
| Ryzen 9 9950X ($650) | $520 | Zen 5 | 80 MB | 16 / 32 | 4.7 / 5/7 | 170W / 230W | 
| Ryzen 9 9900X3D ($600) | $530 | Zen 5 X3D | 140 MB | 12 / 24 | 4.4 / 5.5 | 120W / 230W | 
| Ryzen 9 9900X ($500) | $439 | Zen 5 | 76 MB | 12 / 24 | 4.4 / 5.6 | 120W / 162W | 
| Ryzen 7 9850X3D ($500) | $499 | Zen 5 X3D | 104 MB | 8 / 16 | 4.7 / 5.6 | 120W / 162W | 
| Ryzen 7 9800X3D ($480) | $464 | Zen 5 X3D | 104 MB | 8 / 16 | 4.7 / 5.2 | 120W / 162W | 
| Ryzen 7 9700X ($360) | $305 | Zen 5 | 40 MB | 8 / 16 | 3.8 / 5.5 | 65W / 88W | 
| Ryzen 5 9600X ($280) | $188 | Zen 5 | 38 MB | 6 / 12 | 3.9 / 5.4 | 65W / 88W | 

AMD’s 12- and 16-core X3D chips only use the stacked cache on one of their two CCDs, effectively giving only eight of their cores instant access to the large pool of L3. The Ryzen 9 9950X3D2 changes things up by putting the SRAM chunk on both CCDs.

Get Tom's Hardware's best news and in-depth reviews, straight to your inbox.

It’s not double the cache of the 9950X3D, however, just double the amount of stacked cache. Each Zen 5 CCD has 32 MB of L3 cache, so the 9950X, for example, has 64 MB of L3 total. The 9950X3D boosts that to 128 MB by stacking an additional 64 MB under one CCD. Here, we have the extra 64 MB under both CCDs for a total of 192 MB of L3.

Particularly in heavily-threaded tasks, the goal it seems is to keep data close to the cores. For the other eight cores that otherwise wouldn’t have access to the large, shared pool of L3, they now don’t need to cross over to other CCD to get the data they need. That’s the idea, at least.

Adding extra cache comes with downsides, particularly in thermal and power demands. The Ryzen 9 9950X3D2 peaks slightly lower than its single-cache counterpart at 5.6 GHz, and it demands a 200W TDP, with a peak platform power of 270W; the highest of any consumer Zen 5 CPU. You still get the same 16 Zen 5 cores and 32 threads as the Ryzen 9 9950X3D and Ryzen 9 9950X, however.

The haircut in peak clock speed is telling here, suggesting that a dual cache X3D part wouldn’t have been possible with the previous Zen 4 design. AMD previously stacked the cache on top of the CCD, insulating the cores from direct access to the IHS. Now, the extra cache is under the CCD. That not only allows the Ryzen 9 9950X3D2 to maintain a high-end thermal design, but also to offer full overclocking support, alongside AMD’s Precision Boost Overdrive (PBO).

## AMD Ryzen 9 9950X3D2 cache latency testing

The Ryzen 9 9950X3D2 is the first 3D V-Cache CPU to have stacked L3 cache on both CCDs. The natural question: Does that actually matter? More cache is more cache, but the additional stack here seems like it has less to do with capacity and more to do with latency. As mentioned, if you can access data from L3 on CCD 2, that means you don’t need to go over to CCD 1, which creates additional latency.

First, let’s establish that’s actually the case. Above, you can see a core-to-core heatmap for the Ryzen 9 9950X3D2. This is all about core-to-core communication, and it’s just an illustration to establish that, yes, there’s an additional latency penalty when a core has to cross over to the other CCD. So, if core 16 on CCD 2 needs to get data from the L3 stored on CCD 1, there’s an additional latency hit.

Now, we can look at our latency testing, which measures the time in nanoseconds when allocating different region sizes. Both the Ryzen 9 9950X3D and the non-X3D model see a big increase in latency past about 32 MB, which is the amount of L3 that’s actually on the CCD (the extra 64 MB is stacked for the Ryzen 9 9950X3D). The Ryzen 9 9950X3D2 extends that further, to around 64 MB.

This doesn’t translate directly into better performance, mind you. Lower latency is better performance, but we’re already looking at a truncated version of this chart, starting with a 64 KB region size. The benefits show up when you’re pushing past that normal 32 MB L3 chunk on a Zen 5 CCD.

- **MORE:** __**CPU Benchmark Hierarchy**__
- **MORE:** __**AMD vs. Intel**__
- **MORE:** __**AMD Ryzen 7 9800X3D review**__

