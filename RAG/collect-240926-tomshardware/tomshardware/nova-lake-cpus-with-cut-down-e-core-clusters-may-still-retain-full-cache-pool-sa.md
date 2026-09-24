---
id: collect-240926-tomshardware/tomshardware/nova-lake-cpus-with-cut-down-e-core-clusters-may-still-retain-full-cache-pool-sa
title: "nova-lake-cpus-with-cut-down-e-core-clusters-may-still-retain-full-cache-pool-sa"
domain: tomshardware
role: reference
task: reference
actors: ["AMD", "Google", "Intel"]
dates: []
keywords: ["amd", "consumer", "intel"]
source: docs/RAG/clean_en/tomshardware/nova-lake-cpus-with-cut-down-e-core-clusters-may-still-retain-full-cache-pool-sa.md
source_anchor: ""
source_lines: [1, 30]
sha256: 1469c49b11db50263b5fe537ec994f92ad639a415a516183341e74af651958db
---

# nova-lake-cpus-with-cut-down-e-core-clusters-may-still-retain-full-cache-pool-sa

<!-- source: https://www.tomshardware.com/pc-components/cpus/nova-lake-cpus-with-cut-down-e-core-clusters-may-still-retain-full-cache-pool-says-new-leak-8p-12e-config-predictions-revised-from-33mb-to-36mb-4p-4e-config-from-15mb-to-18mb -->

Intel's upcoming Nova Lake CPUs have been a part of the rumor mill for months as excitement builds up for a potential CES 2027 announcement. The latest leak comes from reliable tipster Jaykihn, who's actually updating an older report that said SKUs with cut-down E-cores would also have reduced cache. Now, his new leak claims Intel will keep the L3 cache unaltered, even on chips with partially disabled E-core clusters.

Previously, it was rumored that some Nova Lake configs, such as the 8+12+4 silicon, would only have 33MB of L3 cache since the 12 E-cores are cut down from the 16 we expect on the full-fat variant of this SKU. Now, Jaykihn reports that this config will now have 36MB of L3 cache instead, just like the uncut 8+16+4 config. Similarly, the 4+4+4 config is said to retain 18MB of L3 cache instead of 15MB, same as the fully enabled 4+8+4 SKUs.

This does not apply to every SKU of Nova Lake, however. For instance, the leaker claims that the 6+12+4 config is still limited to 27MB and not 30MB, had the disabled E-core cluster retained its cache. If true, it seems there's no linear scaling at play here; rather, Intel just decides on its own which SKUs get to keep all of the L3 cache regardless of their disabled E-cores, and which ones are still relegated to lower amounts.

Nova Lake desktop CPUs are not the only ones reportedly affected by this change; Nova Lake-HX, the mobile lineup, is also said to follow the same methodology at this point. Therefore, we can expect some of the midrange SKUs to have more cache than previously expected. Keep in mind that all of this is preliminary, unofficial information and subject to change between now and the actual launch, as this very development proves.

Intel's big trick for Nova Lake is expected to be the introduction of bLLC in consumer CPUs, directly meant to challenge AMD's 3D V-Cache. Only mid- to high-end SKUs are expected to get it, including the 8+12+4 and 6+12+4 configs we mentioned earlier. They're said to feature 132MB and 108MB of bLLC, respectively, which is separate from the native L3 cache we've been talking about in this story, so those numbers won't change.

| Nova Lake-S Rumored SKUs* |  |  |  |  | 
|---|---|---|---|---|
| SKU | Core Config (P+E+LP-E) | bLLC | L3 Cache (updated) | L3 Cache (Previous) | 
|---|---|---|---|---|
| 52 Cores (dual-tile) | (8+16)+(8+16)+4 | 288MB | 72MB? | 72MB? | 
| 28 Cores | 8+16+4 | 144MB | 36MB | 36MB | 
| 24 Cores | 8+12+4 | 132MB | 36MB | 33MB | 
| 22 Cores | 6+12+4 | 108MB | 27MB | 27MB | 
| 16 Cores | 4+8+4 | - | 18MB | 18MB | 
| 12 Cores | 4+4+4 | - | 18MB | 15MB | 

**non-bLLC variants of all the single-tile SKUs are also rumored. All specifications rumored, not confirmed by Intel.*

As usual, there might be some level of internal segmentation that we aren't seeing here. Although leaked specs give us a glimpse at what Nova Lake could offer, it's always possible that Intel is testing various configurations, even if those chips won't end up in the main lineup.

Get Tom's Hardware's best news and in-depth reviews, straight to your inbox.

*Follow* *Tom's Hardware on Google News**, or* *add us as a preferred source**, to get our latest news, analysis, & reviews in your feeds.*
