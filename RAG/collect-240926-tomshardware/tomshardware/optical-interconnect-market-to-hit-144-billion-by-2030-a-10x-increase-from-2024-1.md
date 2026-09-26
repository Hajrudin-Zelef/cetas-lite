---
id: collect-240926-tomshardware/tomshardware/optical-interconnect-market-to-hit-144-billion-by-2030-a-10x-increase-from-2024-1
title: "optical-interconnect-market-to-hit-144-billion-by-2030-a-10x-increase-from-2024-"
domain: tomshardware
role: reference
task: reference
actors: ["China"]
dates: []
keywords: ["accelerator", "compute", "cpo", "dsp", "foundry", "gpu", "gpus", "ipo", "optics", "revenue", "wafer"]
source: docs/RAG/clean_en/tomshardware/optical-interconnect-market-to-hit-144-billion-by-2030-a-10x-increase-from-2024-.md
source_anchor: ""
source_lines: [1, 30]
sha256: fb8c970442dee804ca70ffc65bea0617f2932a64db7aca6ab0149fe330886b64
---

# optical-interconnect-market-to-hit-144-billion-by-2030-a-10x-increase-from-2024-

<!-- source: https://www.tomshardware.com/tech-industry/photonics/ai-data-center-optical-interconnect-market-to-hit-usd144-billion-by-2030-an-over-ten-fold-increase-from-2024-figures-according-to-new-projections-silicon-photonics-expected-to-account-for-nearly-two-thirds-of-revenue-driven-by-co-packaged-optics -->

The global data center optical interconnect market is expected to reach $144.4 billion by 2030, up from $13.7 billion in 2024 — a 48.1% compound annual growth rate (CAGR) — according to a China Insights Consultancy (CIC) report commissioned by a Chinese laser-chip maker, Yuanjie Semiconductors, as part of its Hong Kong IPO filing. The report draws on data from LightCounting and interviews with industry experts. Of that future market, silicon photonics, the practice of manufacturing photonic chips from the same silicon material and mature CMOS foundry processes used for conventional semiconductors, is projected to account for 63.7% of revenue, its share climbing from 16.6% in 2020 as the industry shifts toward denser, more power-efficient designs like co-packaged optics.

The moves that would turn those projections into reality are already well underway. Over the past year, the AI industry has invested more than $15 billion in co-packaged optics, photonic chips, higher-speed transceiver modules, and fiber, developing new integration techniques, acquiring photonics startups, and forming alliances among the biggest players. More recently, OpenLight and Tower Semiconductor placed OpenLight's photonic design kit inside Cadence's mainstream chip-design software. This step makes the laser-integrated 400G and 1.6T chips at the heart of co-packaged optics easier to design and bring to market.

## The tech behind the numbers

For decades, data centers have relied largely on copper traces and cables to move data across circuit boards, within racks, and across clusters. Copper is cheap, reliable, and easy to integrate. However, its power consumption and signal losses increase sharply with bandwidth and distance. As AI data centers are packed with ever more powerful GPUs, shuttling enormous volumes of data and pushing networks toward higher speeds, copper hit a wall. Past a few hundred gigabits per lane, its usable reach collapses to a meter, or two, before signal loss and power draw become unmanageable.

The solution has been a transition to photonics, moving data as light instead of electrical signals. An optical transceiver converts electrical signals from switches and processors into laser light, sends it down a fiber, and converts it back at the far end, carrying far more bandwidth over greater distances at much higher speeds. Today, pluggable transceivers pack a laser chip, digital signal processor (DSP) chips, and several optical components into one compact module. As GPUs grow more capable and AI workloads swell, both the volume of data and the speed it must travel keep climbing, pushing the industry from 400G links to 800G to 1.6T and beyond

At the same time, the industry is trying to move the optics closer to the compute. In conventional systems, GPU signals travel inches along copper traces across the board to reach the transceiver on the faceplate. At extreme data rates, even that relatively short electrical journey consumes considerable power. Co-packaged optics (CPO) fixes this by pulling the optical engine out of the pluggable module and placing it as a chip — the photonic integrated circuit (PIC) — directly on the switch or accelerator package, shrinking the electrical path to millimeters. The push for faster optical chips, aiming for terabits-per-second speeds, serves both CPO and the pluggable modules that remain the industry mainstay.

The PIC does everything but generate light. Because laser chips are highly sensitive to heat, they can't be folded into the PIC, which, in co-packaged optics, becomes part of the switch or accelerator package that runs extremely hot. Therefore, the laser stays a separate chip. Whether feeding a co-packaged PIC or a pluggable module, those laser chips are always needed, which is exactly what Yuanjie, the company that commissioned the CIC forecast, makes.

The PIC itself is where silicon photonics comes in. Photonic chips were once built entirely from costly III-V materials in specialized fabs; Silicon Photonics (SiPh) instead patterns the optical circuitry onto silicon using the same mature, high-volume CMOS processes as ordinary chips — saving money and time and making PICs mass-producible. On the other hand, silicon cannot lase. As a result, laser chips still rely on the more expensive III-V method, using materials such as indium phosphide. Circumventing exactly that is the aim of the recent OpenLight–Tower platform. Their approach integrates III-V laser material directly with silicon photonics at the wafer level, aiming to bring the laser into the same scalable manufacturing flow as the rest of the PIC.

## The numbers behind the growth

Together, the different photonics technologies solving the AI data transfer bottleneck are driving a huge market in the industry. CIC projects the data center optical interconnect market growing from $13.7 billion in 2024 to $144.4 billion in 2030, a compound annual growth rate of 48.1%, more than a tenfold increase in six years. Growth accelerates over this period, with the steepest gains occurring after 2027. The report breaks down the three main ways: technology, use case, and data rates, each showing where the spending and resulting revenue are concentrated.

By technology, it splits the market between silicon photonics and everything else. SiPho's share climbs from 16.6% in 2020 to 63.7% of total revenue ($91.9 billion) by 2030, with its revenue compounding at 68.5% annually, compared with 32.6% for the rest, and the crossover past half the market landing around 2027. According to CIC, silicon-based optics— primarily PICs — will grow to become the default.

The use-case breakdown shows which parts of data center networking are driving optical demand: scale-up inside the rack, scale-out across a data center, and scale-across between data centers. Scale-up — the short-reach links from servers and chips to the top-of-rack switch — takes the lead with a 561.5% CAGR in revenue. Note that this figure compounds off a near-zero 2024 base, where a tiny absolute gain reads as an absurd percentage. Scale-across follows at 108.5% and scale-out at 42.7%, while the entire non-AI segment, traditional workloads from telecom to enterprise servers, trails at 37.5%. However, in actual revenue and not growth, scale-out is the largest tier at $64.5 billion in 2030, followed by non-AI at $40.2 billion and scale-up, for all its headline growth — at just $32.1 billion, with scale-across last at $7.6 billion.

The data-rate breakdown captures a generational migration of speed. In 2024, the market still ran on the previous two generations, 400G at $5.9 billion and 800G at $4.5 billion, with legacy 200G-and-below links adding another $3.3 billion and the faster tiers barely registering. By 2030, that order is inverted. 1.6T, only entering commercial deployment in 2026, becomes the single biggest segment at $65.6 billion, expanding at an 867.3% CAGR from a 2024 base of essentially zero; 3.2T, a category that didn't exist in 2024 at all, appears only from 2027 and still vaults to $44.5 billion. Together, those two next-generation speeds make up roughly $110 billion of the $144.4 billion total. The rest slides down the ladder: 800G stays healthy at a 34.8% CAGR to $26.7 billion, but 400G flatlines — 1.0% annual growth to $6.3 billion, after leading the prior cycle at 75.3% — and 200 G and below actively contracts, shrinking 15.4% a year.

## The moves behind the numbers

