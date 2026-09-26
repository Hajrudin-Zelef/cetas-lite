---
id: collect-240926-tomshardware/tomshardware/amd-ryzen-9-9950x3d2-review-more-cache-more-cash-2
title: "amd-ryzen-9-9950x3d2-review-more-cache-more-cash"
domain: tomshardware
role: reference
task: reference
actors: ["AMD", "Intel", "Nvidia"]
dates: []
keywords: ["amd", "attention", "benchmarks", "cost", "gpu", "intel", "nvidia", "voice"]
source: docs/RAG/clean_en/tomshardware/amd-ryzen-9-9950x3d2-review-more-cache-more-cash.md
source_anchor: ""
source_lines: [71, 103]
sha256: 028809970cc659942602df6b2cf11ae7d4fe6dd24c6e42f3e5ec7b6a2c88d5d3
---

# amd-ryzen-9-9950x3d2-review-more-cache-more-cash

- 
Reply
 Very nice and informative article! Can I get a clean pdf version for education use? Thank you!Admin said:The Ryzen 9 9950X3D2 is one of the most unique CPUs we’ve ever reviewed, and although its price feels like a kick in the gut, it offers some interesting, highly specialized improvements in certain workloads based on our testing.
 
 **AMD Ryzen 9 9950X3D2 review: More cache, more cash : Read more**
- 
Reply
It's the same list of benchmarks we use in every review, AMD and Intel.Gururu said:How was it decided to find a bunch of obscure benchmarks that are rarely used in CPU testing? Seems like a little voice whispered in someone's ear...
- 
Reply
Yes, I don't see anything wrong. It's just a little confusing on the SPECWorkstation 4 Benchmarks where if you compare the 270K review to this review, some tests seem different. Maybe they are just listed in a different order.JakeRoach said:It's the same list of benchmarks we use in every review, AMD and Intel.
- 
Reply
 Cache is completely transparent (invisible) to apps, they either have enough data to fill the cache or they don't.Marlin1975 said:Seems like a chip that with the right software will be a beast. I'm assuming since most is not written for this much cache its left spinning its wheels when it has more to go.
You would need to invent an app that causes the problem of needing that much cache for it to use that much cache. (Which is what a lot of the benchmarks do, they use a lot more data (or at least coherent/fixed amount of data that never needs to change) than what a normal real world usage would be)
- 
"A terrible value, but one of the most unique...isn’t worth the money for the vast majority of people, but it was never meant to be. It’s a halo product.....Reply**one of the most unique .... ever reviewed, and although its price feels like a kick in the gut...** ."
 
Sounds like summary statements appropriate to an Nvidia GPU reviews.
- 
Reply
 No need to invent anything. If you just consider an app that at its core does FFT or vector arithmetic then as soon as the size of the data is larger than cache of 9950 but smaller than 9950X3D2 you will see a big difference in speed.TerryLaze said:Cache is completely transparent (invisible) to apps, they either have enough data to fill the cache or they don't.
 You would need to invent an app that causes the problem of needing that much cache for it to use that much cache. (Which is what a lot of the benchmarks do, they use a lot more data (or at least coherent/fixed amount of data that never needs to change) than what a normal real world usage would be)
 
 The reason you don't quite see this in charts of this article is because most apps are in two categories - either they are written without much attention to performance, in which case they spend most time in CPU executing some byte code or inefficient loops and the extra cache does not matter.
 
Or they have been well optimized and part of that optimization was to fit them into the cache of the CPUs they were designed for, and the case of my data is larger than cache was treated as a slow path.
- 
Thanks for the comprehensive benchmarks! I had always wondered how such a product would perform - now I know!Reply
 
 In your intro, I didn't notice a reference to why AMD said they didn't offer this before. I'd have to go searching for it, but they've previously said they didn't think it would be cost-effective. It seems they were right.
 
I'm glad to see very few regressions vs. the 9950X and 9950X3D, however. That means it's a safe buy for someone who wants the top AM5 multithreaded performer, if money is no object.
