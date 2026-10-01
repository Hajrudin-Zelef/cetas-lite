---
id: collect-261001-ia-llm/ia-llm/claude-computes-a-nine-loop-amplitude-in-n-4-super-yang-mills-3
title: "claude-computes-a-nine-loop-amplitude-in-n-4-super-yang-mills"
domain: ia-llm
role: reference
task: reference
actors: ["Anthropic", "OpenAI", "United States", "Z.ai"]
dates: []
keywords: ["claude", "compute", "cost", "cyber", "glm", "gpt-6", "safeguards"]
source: docs/RAG/collect-261001-ia-llm/claude-computes-a-nine-loop-amplitude-in-n-4-super-yang-mills.md
source_anchor: ""
source_lines: [93, 117]
sha256: 5ca1140252095a2232f30d28e16d34669cc0e3a592048d1c0bb886e26d5003e0
---

# claude-computes-a-nine-loop-amplitude-in-n-4-super-yang-mills

Most theoretical physicists I know recognize that the current era of large language models is going to completely transform the way we think about physics. The question was just: when was it going to really hit home? For me, it happened on September 1, when Liam Fitzpatrick and Siddharth Mishra-Sharma at Anthropic told me that Claude had computed the nine-loop MHV six-particle amplitude in planar N=4 super Yang-Mills, and asked me to validate its result.

I'm not going to explain all the technical terms in that last sentence; Matt has covered the background above. I do need to mention that there are really two related objects, the "amplitude" and something we call the “form factor.” Each has an associated number of loops: one, two, three, and so on. Every loop order is harder than the previous one, computationally, even after finding lots of tricks to make things easier. Also, the form factor is easier than the amplitude at the same loop order. In 2023 Andy Liu and I showed how to use the form factor and a weird symmetry we call antipodal duality to get the amplitude at eight loops.

Since 2023, my collaborators and I have eyed getting to nine loops, first for the form factor and then for the amplitude, using our 2023 idea. I thought it would be too hard to do the amplitude directly. So I was really quite impressed that Claude could do it directly. Not so much because it was a big computational task, but because the whole setup is very fragile: if you make any mistake at all in the computational recipe, it all crashes down like a failed soufflé, and you are left to wonder why (and debug). Also, there are so many details of the construction that are too boring to document fully in a publication. So Claude had to develop all that code from scratch.

From the nine-loop amplitude it is relatively easy to go back to the form factor, and it was easier for me to validate the result mostly that way. That meant that for the last two weeks I've been validating a result, the nine-loop form factor, that our team had been working toward for a couple of years. And a machine had solved a problem that I thought was too hard to do directly. Does that bother me personally? Is it soul-crushing?

No, for two reasons. One is that our team already had a campaign to use custom transformer models to predict higher loops, and part of our slogan was: “We have all the tools to validate any candidate solution a machine would provide us.” Claude is a different kind of transformer model, probably over a million times bigger than our custom one. But sure, we said we could validate any result an AI model would give us, so we can and should do it. The second reason is that, if you look at how Claude solved the problem, it used all the methods my collaborators and I developed over the years, and it presented the solution (maybe as a favor to us) in the same format we had already set up. So while I'm validating Claude's result, Claude is validating all of our previous work. In fact, I would assert that Claude understands our 2019 and 2023 papers better than any human, aside from my co-authors.

After I wrote this, Song He told me that his group had also computed the piece of the nine-loop amplitude called the symbol. (People just seem to like to tell me about their nine-loop successes, for whatever reason.) Song's group used AI (GPT-6) to help them compute some of the constraints, but not for the overall framework. So now I've been scooped by both a machine and by humans plus a machine, within two weeks.

Going back to the Claude computation: it's quite a triumph, in my opinion, for a large language model to execute all of the steps in the complicated recipe we laid out, and to organize the computational horsepower. But the more soul-searching moments will come when large language models start to come up with new physical principles and insights before humans.

Anthropic invited Matt von Hippel to write this post and compensated him for his time. Anthropic staff gave feedback on drafts; the content and opinions are his own. Lance Dixon validated the result independently and received Claude usage credits.

We built an index of how well today’s robots can perform US job tasks. Robots can already do three-quarters of physical tasks, mostly in limited settings, but are cost-competitive for just 0.3% of them.

Read more
We’re launching a new study using Anthropic Interviewer to learn from your experiences with AI, and we invite you to participate.

Read more
Like Claude Mythos Preview, GLM-5.3 has strong capabilities for autonomously building end-to-end cyber exploits. But GLM-5.3 is unlike other frontier models in that it has been released without meaningful safeguards to limit misuse.

Read more
