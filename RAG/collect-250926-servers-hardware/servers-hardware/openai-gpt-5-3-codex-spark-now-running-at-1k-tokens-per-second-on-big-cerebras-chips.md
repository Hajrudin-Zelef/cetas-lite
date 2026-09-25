---
id: collect-250926-servers-hardware/servers-hardware/openai-gpt-5-3-codex-spark-now-running-at-1k-tokens-per-second-on-big-cerebras-chips
title: "openai-gpt-5-3-codex-spark-now-running-at-1k-tokens-per-second-on-big-cerebras-chips"
domain: servers-hardware
role: reference
task: reference
actors: ["Cerebras", "Intel", "Nvidia", "OpenAI", "TSMC"]
dates: []
keywords: ["agentic", "agents", "compute", "cost", "inference", "intel", "latency", "memory", "nvidia", "packaging", "research", "tokens per second"]
source: docs/RAG/clean4/openai-gpt-5-3-codex-spark-now-running-at-1k-tokens-per-second-on-big-cerebras-chips.md
source_anchor: ""
source_lines: [1, 33]
sha256: b1c1635a6a8b6e624d610a73b5b0e17120ba48fa770260e3151c7dfa0f138f67
---

# openai-gpt-5-3-codex-spark-now-running-at-1k-tokens-per-second-on-big-cerebras-chips

OpenAI announced that it is rolling out a new GPT-5.3-Codex-Spark as a research preview. OpenAI has released many models, but this is designed as a coding assistant that is not just another model. Instead, it is designed to be super-fast with 1,000 tokens per second performance on giant Cerebras chips. This is the first public collaboration between OpenAI and Cerebras, and it is notable.

## OpenAI GPT-5.3-Codex-Spark Now Running at 1K Tokens Per Second on BIG Cerebras Chips

In a quick demo, OpenAI showed a “build a snake game” task issued to GPT-5.3-Codex-Spark and GPT-5.3-Codex at medium. Both completed the task, but the Cerebras-backed Spark ran in 9 seconds, compared with nearly 43 seconds on the non-Spark model. If you want to see the video of the side-by-side, here is a link. The Spark model is said to be higher quality than the GPT-5.1-Codex as well as much faster.

OpenAI says that it is running these on the Cerebras Wafer-Scale Engine 3 (WSE-3) that we highlighted in 2024 on STH. For those who have not been following Cerebras, they take the largest square they can from a wafer and then manage to power and liquid-cool it without cutting it into many smaller chips.

If you want to see a bit more of how these are cooled, we covered this in our SC22 video:

The idea is that with these huge chips even faster inference can happen. For agentic AI, having faster task completion means that workflows run faster.

## Final Words

This is a cool announcement because Cerebras has been one of the few AI companies that I have been saying has a shot. In 2020, just about six years ago, we published Our Interview with Andrew Feldman CEO of Cerebras Systems. Folks that know STH know I am not a fan of doing interviews, but I made an exception for this one. I remember that as I was flying to New Zealand later that evening and got to tour the Los Altos lab that was cooled by partially cutting the side of the building. Not long after that interview, the world went into pandemic mode, but this is one that stuck with me for years. It is great to see the Cerebras team have success.

For all of us that have been setting up workflows with tools like n8n (see our Using the Dell Pro Max with GB10 to Profit within 12 Months piece) or OpenClaw, the reality is that fast inference will be important as we ask systems to do larger tasks. The “build a snake game” prompt leading to a working browser game in 9 seconds certainly made me think of how long it took me to do the same task a quarter of a century ago. When I think about that, I think about the time I spent actually making a (then) Java-based snake game. Perhaps underestimated is that it took some time learning concepts even just to get to that point. Now, it is a simple prompt and nine seconds without coding. It also makes me think of how fast the agents of the future are going to turn ideas into reality.

Hopefully we get to see more Cerebras inference solutions in the near future because this is cool.

So Cerebras is roughly 10x to 20x faster than H100 in ballpark figure?

This reads more as “Patrick has 10x to 20x more money invested in Cerebras than NVIDIA,” in ballpark figures.

I remember going through school and being told that wafer scale designs were impossible. Historically there was evidence for this as the early attempts in the 1980s all failed. The idea would occasionally pop up again, with some further research being done by Sun Microsystems decades ago. This idea was the holy grail of computing as all the complexities of building off chip IO would disappear with everything becoming tightly integrated. Through clever engineering, Cerebus has pulled it off for compute and memory. Each tile in the design is identical which is a key attribute for how they’re able to scale.

With more advanced packaging technologies out there, I can see Cerebus or nVidia leveraging an entire wafer as the silicon interposer where smaller, more advanced chiplets are placed on top of. The advantage here is that the chiplets can be tested prior to packaging so that defective units can be discarded. (Cerebus currently disables various compute tiles as necessary.) The other big feature of going full wafer with chiplets is that the full circular shape can be exploited to include smaller IO chiplets along the parameter. With silicon photonics set to become mainstream soon, it would be straight forward to package networking around the compute cores.

One thing that didn’t happen a decade ago was the transition to 450 cm waters. Intel at the time did not want to fund the investment as they were the leading edge manufacturer at the time (how times have changed!). TSMC was reportedly interested as a technology but had to quickly down development as they did not have the means to fund it on their own. Had 450 mm wafer production moved forward, the early pandemic chip crunch would not have happened. Less clear would be the current supply side difficulties be addressed. As for Cerebus, they would be able to offer a chip with twice as much compute vs. what they’re offering now.

I think the Cerebras Wafer Scale Engine 3 is being sold for $2-3 million per chip. Are these better than buying the smaller, competing accelerators, from cost, efficiency, or latency standpoints?

I remember reading that article way back as a newer ServeTheHome reader–my head exploded at the absurdity of it. No doubt a very bold product strategy!
