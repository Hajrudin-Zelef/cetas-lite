---
id: collect-240926-tomshardware/tomshardware/alibaba-claims-new-qwen-image-2-1-ai-model-beats-google-nano-banana-2-0-with-min
title: "alibaba-claims-new-qwen-image-2-1-ai-model-beats-google-nano-banana-2-0-with-min"
domain: tomshardware
role: reference
task: reference
actors: ["Alibaba", "Google", "LongCat", "Meituan", "Meta", "Nvidia", "OpenAI"]
dates: []
keywords: ["qwen", "apache", "benchmark", "benchmarks", "consumer", "diffusion", "lean", "license", "memory", "muse", "nvidia", "open-weight"]
source: docs/RAG/clean_en/tomshardware/alibaba-claims-new-qwen-image-2-1-ai-model-beats-google-nano-banana-2-0-with-min.md
source_anchor: ""
source_lines: [1, 51]
sha256: 1fa0bb66c6340eb99af3f7c89b60ec9639fea3892be88e1a817f82027af66dd0
---

# alibaba-claims-new-qwen-image-2-1-ai-model-beats-google-nano-banana-2-0-with-min

<!-- source: https://www.tomshardware.com/tech-industry/artificial-intelligence/alibaba-claims-new-qwen-image-2-1-ai-model-beats-google-nano-banana-2-0-with-minuscule-7b-parameter-model-benchmarks-show-open-weight-contender-is-competitive-with-openai-and-meta-image-models -->

Alibaba Cloud has released a new lightweight image generation AI model, called Qwen Image 2.1. Sporting just 7 billion parameters, it's an extremely lean open-weight model, able to run on even older consumer graphics cards like the RTX 3090. Despite its lightweight design, its developers claim it is more capable than a range of closed-weight models, including Google's Nano Banana 2.0.

That is on its own internal benchmark, so we'd like to see some additional testing before making any concrete claims, but early reports suggest it's a very capable image model, with native transparency support and the ability to create unified images from a range of reference images.

One area that has raised eyebrows, though, is the change to the Qwen Image 2.1 licensing agreement. Unlike the previous version, this one explicitly forbids commercial resale of the model, requiring anyone who wants to use it for that to obtain a separate license directly from the developer.

## Competing with the best?

Qwen Image 2.1 introduces a number of new features that improve its utility and help it better compete with established alternatives. It supports native transparency, so you can have it generate images with transparent backgrounds, which can make it particularly useful for artists wanting to use AI as part of something else, or for print-on-demand products like stickers.

It also improves image editing, with the ability to use up to 10 reference images. Cited examples include taking an existing image of a person, feeding the model images of clothing items, and then it can composite an image of the person wearing those clothes. The Qwen team promises consistency across images, preserving people and products effectively.

The standout feature of Qwen Image 2.1, though, is its compact size. Its visual generation component has just 7B parameters, making it one of the leanest models of its kind. In first-party comparison testing, the only model with fewer parameters was LongCat-Image developed by Meituan, coming in at 6 billion parameters. Most other models are several times that size, or closed-weight entirely.

It's those closed-weight models that Qwen Image 2.1's developers claim it can compete directly with, though. On that same internal benchmark, its model scored 60.2 on the Qwen Image Benchmark. In comparison, OpenAI'S GPT Image 2.5 Sunburst scored 67, Muse Image 62.34, and Google's Nano Banana 2.0 59.82.

Preliminary testing on AI Arena suggests it's not quite as strong as that in less favorable benchmark conditions, but still very close. There, it achieved a score of 1228, while Nano Banana 2 managed 1260. Musw Image is ranked a few steps higher again, with a score of 1276, while GPT 65 Sunburst sits at the top with a score of 1423.

On the image editing front, AI Arena claims it's now the best of the open-weight options:

These are still early results, and more testing will be needed to fully confirm or refute the Qwen team's claims, but so far they seem to track pretty closely to reality. Although we can't know how many parameters the proprietary models lean on to achieve their higher scores, Qwen Image 2.1 appears to be nipping at their heels with its modest weights.

## It works well on local hardware

Reports from individual early adopters of Qwen Image 2.1 claim it runs just fine on their consumer graphics cards. Hacker News forum member Vunderba, developer of the GenAI Showdown site, claims they were able to convert a 1MP image using Qwen Image 2.1 in around five seconds on an RTX 4090.

On the Stable Diffusion subreddit, user cgs019283 praised its capabilities, stating that it could generate 1MP images in around 25 seconds on modern Nvidia 50-series graphics cards like the RTX 5070 and 5080. They did, however, state that it can take a lot longer if you use lots of reference images, and that editing an image was the slowest function of the model in their testing.

Others have the model working on older and much lighter hardware. One user claims to be using it to generate 2K resolution images in around 50 seconds using an Nvidia RTX 3060 and 64GB of memory.

One user is lucky enough to have an RTX 6000 Pro to play around with. There, the powerful hardware is able to put out 1024 x 1024 images in just a few seconds.

Although the setup and use of local AI is still a far cry from the accessibility of cloud-based platforms like Google's Nano Banana and OpenAI's GPT Image models, the early results appear, at least on the surface, to be impressive. If the feature set holds up under more rigorous testing, those wanting the kind of rapid, quality image generation but running locally with full control and no need for cloud accounts or subscriptions could make it a fierce competitor.

## The licensing conundrum

Of all the commentary from Qwen Image 2.1 early adopters, a discussion that keeps coming up is how the developers are handling licensing. The wording comes across as ambiguous:

"You are granted a non-exclusive, worldwide, non-transferable, and royalty-free limited license under our intellectual property or other rights owned by us embodied in the Materials to use, reproduce, distribute, copy, create derivative works of, and make modifications to the Materials FOR NON-COMMERCIAL PURPOSES ONLY. "

It then says that anyone wanting to use those "Materials" for commercial purposes would need to acquire a license from the Qwen team specifically for that use case. It's caused enough concern among the community that the Qwen team released a statement on Twitter/X:

That mostly seems to clear things up. Whatever you generate with the model shouldn't be classed as a licensed material, so it shouldn't be covered by this clause.

What this does mean, though, is that the model itself cannot be re-sold without a license from Alibaba. That's quite different from the Apache model the original Qwen Image was based on, and potentially stretches the definition of open-weight. It's certainly not as open as it could be.

Mais pour la plupart des gens, ce n’est pas un problème. Vous pouvez exécuter Qwen Image 2.1 sur votre matériel local pour créer des images générées plutôt efficaces et les utiliser comme bon vous semble. Avec une conception aussi légère offrant des capacités aussi impressionnantes, la réponse des développeurs de modèles à poids fermé sera intéressante à observer.

Jon Martindale est un rédacteur contributeur pour Tom's Hardware. Depuis 20 ans, il écrit sur les composants PC, les technologies émergentes et les dernières avancées logicielles. Son expérience journalistique profonde et étendue lui donne un regard unique sur les tendances technologiques les plus passionnantes d’aujourd’hui et de demain.
