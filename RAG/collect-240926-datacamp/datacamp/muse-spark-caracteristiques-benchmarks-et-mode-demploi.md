---
id: collect-240926-datacamp/datacamp/muse-spark-caracteristiques-benchmarks-et-mode-demploi
title: "muse-spark-caracteristiques-benchmarks-et-mode-demploi"
domain: datacamp
role: reference
task: reference
actors: ["Anthropic", "Apple", "China", "DeepSeek", "Google", "Meta", "OpenAI", "United States"]
dates: ["2025-04", "2025-06-30", "2025-11", "2026-03-13", "2026-04-08"]
keywords: ["benchmark", "benchmarks", "muse", "agent", "agentic", "agents", "agi", "alignment", "attention", "chatgpt", "claude", "compute"]
source: docs/RAG/clean_en/datacamp/muse-spark-caracteristiques-benchmarks-et-mode-demploi.md
source_anchor: ""
source_lines: [1, 269]
sha256: 0b7a931e0625863a569bcd71c7903429f3336503711db094c8319bc4179798b6
---

# muse-spark-caracteristiques-benchmarks-et-mode-demploi

<!-- source: https://www.datacamp.com/fr/blog/muse-spark -->

Cursus

We were publishing articles on Meta's Llama models (Llama 2, Llama 3, etc.) at a steady pace. Then Llama 4 arrived in April 2025 amid a firestorm of criticism: several media outlets and the company's outgoing AI director confirmed that benchmark results had been manipulated using specialized sub-models that were never released.

After that, the updates stopped. At the same time, Meta announced that Horizon Worlds would move to mobile only, effectively ending the VR version on which it had once staked the company's future. Taken together, it painted the picture of a company losing its footing on two fronts at once.

On April 8, 2026, Meta launched Muse Spark, the first model to come out of Meta Superintelligence Labs. The press release repeats the phrase "personal superintelligence" a bit too often. Once that marketing layer is stripped away, what emerges is a genuine model that puts Meta back in the race at the highest level.

To compare Meta's new model with one of its best current competitors, we recommend our guide Muse Spark vs Claude Opus 4.6. You can also read our Muse Glimmer guide and our tutorial on running Muse Glimmer locally.

## What is Muse Spark?

Muse Spark is a natively multimodal reasoning model that handles text, images, audio, and tools within a single architecture. It supports a visual chain of thought: the model can break down image-based problems step by step instead of producing a single answer. Multi-agent orchestration is also part of the equation, and we'll come back to that.

The early Llama models returned answers by matching patterns learned during training. Muse Spark thinks through the problem before responding. That's the real change.

### Who's in charge?

Meta Superintelligence Labs, or MSL, was created on June 30, 2025, when Mark Zuckerberg reorganized the company's AI operations. Alexandr Wang, former CEO of Scale AI, came on as Chief AI Officer; Meta had invested about $14 billion in Scale AI as part of the deal.

Nat Friedman, former CEO of GitHub, oversees product and applied research, and Shengjia Zhao, who co-created GPT-4 and o1 at OpenAI (the same o1 that Muse Spark is now compared to in benchmarks), is Chief Scientist.

A third factor matters: Yann LeCun, Meta's longtime Chief AI Scientist and the company's leading advocate for open source, left in November 2025. His departure followed organizational changes that had limited his role and the team's shift to closed development.

## What's new with Muse Spark (and why it matters)?

The highlights: reasoning modes, an overhauled training pipeline, and a deliberate focus on health. Let's go through them.

### Three reasoning modes

Muse Spark offers three ways to interact with it, and the distinction is worth understanding before you try it.

- **Instant** is the default mode for everyday questions. It responds quickly without extended reasoning, much like a standard chat model.
- **Thinking** uses an extended chain of thought. The model takes longer, breaks down intermediate steps, and generally performs better on hard problems. Most of the benchmark results cited here come from this mode.
- **Contemplating** is the most interesting. Details below.

A useful clarification right away: Contemplating mode rolled out gradually and wasn't available to everyone on launch day. If you don't see it yet, that's normal.

### Contemplating mode

Contemplating mode launches several reasoning agents in parallel, then combines their outputs into a single response. Where Gemini's Deep Think and OpenAI's GPT Pro mode extend reasoning by thinking longer, Muse Spark does it by thinking broader. More agents work simultaneously rather than one agent working longer.

Meta claims this approach produces comparable results with lower latency, since the agents operate in parallel rather than sequentially. Those latency figures have not yet been independently confirmed, but the benchmarks for Contemplating mode lead on several difficult evaluations (we'll come back to that).

This is an inference-time feature, not an architectural choice. The model itself doesn't change.

### Reinforcement learning scaling and thought compression

Meta rebuilt its training pipeline from scratch during the nine months of Muse Spark's development. The claims about reinforcement learning (RL) come from Meta's technical blog and have not been independently verified.

The most interesting point is a technique called thought compression. During RL training, the model is rewarded when it finds the right answer, but is also penalized for thinking time, which amounts to limiting output tokens. This induces three-phase behavior on complex tasks such as math problems.

First, the model improves by "thinking longer." Then the length penalty kicks in and forces the model to solve the same problems with far fewer tokens. At some point, it expands its reasoning again and surpasses its previous ceilings while using fewer tokens.

The practical consequence: the model learned to do more with less. This claim rests on Meta's training curves, which have not been independently validated.

### A compute requirement divided by 10

Meta claims its new architecture matches the performance of Llama 4 Maverick with ten times less training compute. This is a matter of architectural efficiency, not a ceiling for Muse Spark. Llama 4 Maverick scored 18 on the Artificial Analysis Intelligence Index. Muse Spark scored 52.

Artificial Analysis's token efficiency figures point in the same direction. Muse Spark used 58 million output tokens. GPT-5.4 used 120 million. Claude Opus 4.6 used 157 million.

### Health: a deliberate focus

Health is the benchmark territory where Muse Spark stands out the most, and that's no accident. Meta worked with over 1,000 physicians to curate training data specific to medical reasoning.

The model can generate interactive displays covering nutritional composition, medication information, and exercise physiology. On HealthBench Hard, Muse Spark scored 42.8 versus 40.1 for GPT-5.4 and 20.6 for Gemini 3.1 Pro. This gap with Gemini is confirmed under independent evaluation.

This is clearly Meta's answer to ChatGPT Health. Meta's argument for its competitiveness: the social context of 3 billion users, which would give it an edge in understanding how people actually ask their health questions. Whether this holds up for complex or atypical queries—beyond the common questions that dominate benchmarks—remains to be seen.

## And what about Llama in all this?

The developer community is asking a legitimate question, one that deserves a clear answer.

Muse Spark is not open-source. All Llama models up to Llama 4 shipped with weights that developers could download and run locally. Communities like r/LocalLLaMA were built on this. That use case is gone.

The reason cited by Meta is partly competitive: Chinese labs, including DeepSeek, used Llama's weights to accelerate their research. Wang stated that the company "hopes" to open-source future Muse models, without a timeline. "Hopes" carries a heavy weight here.

The Llama team was integrated into Wang's lab, and Llama 4 is the last model to come out of the old structure. Whether Llama continues alongside Muse or quietly fades away, Meta isn't saying.

## Muse Spark benchmark results

With Muse Spark, the benchmarks require distinguishing one important thing from the outset: given Meta's track record with Llama 4, keep the figures declared by the publisher clearly separate from those independently verified.

Here are the results in Thinking mode, which allow for the fairest comparisons.

Source: Meta Superintelligence Labs / ai.meta.com

Contemplating mode has another set of results for the hardest evaluations. It leads on Humanity's Last Exam and FrontierScience Research, but remains behind GPT-5.4 Pro and Gemini 3.1 Deep Think on the IPhO 2025 physics theory problems. All of this comes from Meta: treat these figures as indications, not verdicts.

Source: Meta Superintelligence Labs / ai.meta.com

The independent overview provided by Artificial Analysis is more measured. They place Muse Spark fourth on their Intelligence Index, behind Gemini 3.1 Pro Preview, GPT-5.4, and Claude Opus 4.6. Still in the global top 5. The figures also highlight the weak spots: ARC-AGI-2 and Terminal-Bench 2.0 should hold your attention if coding or abstract reasoning are essential to your use case.

## Putting Muse Spark to the test

With these scores in mind, let's test Muse Spark. I examine the model on multi-step reasoning, image understanding, and code debugging.

### Test 1: Fibonacci–binary logic chain (stability of cascading reasoning)

In this first test, I target Muse Spark's advanced reasoning capabilities on a multi-step exercise. The model must:

- Identify the correct Fibonacci term
- Convert it correctly to binary
- Precisely count the bits
- Generate the prime numbers within a calculated interval
- Perform a large summation

The prompt used:

```
Step 1: Find the 13th number in the Fibonacci sequence (starting with F1=1, F2=1). Let this be X.
Step 2: Convert X into a binary string (Base 2).
Step 3: Count the number of '1's in that binary string. Let this count be C.
Step 4: Identify all prime numbers (p) such that 20 ≤ p ≤ (C × 100).
Step 5: Calculate the sum of these primes. What is the final result?
```
Muse Spark did very well and solved the exercise on the first try. This is all the more impressive given that GPT-5.4 failed at the last step and only succeeded after splitting it into two steps (listing the primes, then adding them).

### Test 2: image understanding and business reasoning on a multi-series time chart

Meta claims Muse Spark understands complex images very well; so I use the following multi-series time chart to see if it can detect patterns and translate them into actionable recommendations.

Here is the prompt:

`Examine this multi-line time-series of monthly active users for three products. Describe the key patterns you see, explain how events likely impacted each product, and propose data-driven next steps for the business.`
Muse Spark correctly identified all the patterns, which suggests that image recognition works well.

The data was randomly generated, so there is no absolute ground truth here. That said, Muse Spark identifies all the events, reasons per product and per period, and arrives at relevant conclusions. It even analyzes the evolution of the sum of MAUs (monthly active users) across product combinations, without being asked to, which is a real plus.

All suggested next steps align with the MAU pattern analysis and event effects. Muse Spark pinpointed the central theme of each product (launch playbook for A, pricing for B, scaling for C) and proposed coherent concrete actions.

### Test 3: code debugging

Finally, I test Muse Spark's skills for diagnosing bugs. The goal is to see whether the model limits itself to checking correctness line by line, or whether it can also detect conceptual errors.

The prompt:

```
A developer wrote this Python function to compute a running average: 
def running_average(data, window=3): 
    result = [] 
    for i in range(len(data)): 
        start = max(0, i - window + 1) 
        chunk = data[start:i + 1] 
        result.append(round(sum(chunk) / window, 2)) 
    return result 
When called with running_average([10, 20, 30, 40, 50]), the first two values in the output seem wrong. Why? Please help me fix what is wrong!
```
The function always divides by `window (3)`, even at the beginning when the segment has fewer than 3 elements. The buggy output is `[3.33, 10.0, 20.0, 30.0, 40.0]`, but the first two values should be `10.0` and `15.0` since those segments contain 1 and 2 elements respectively. The fix is to replace `/ window` with `/ len(chunk)`.

Models often trace the loop perfectly, but conclude that the output seems "correct." They see the calculations step by step and do not flag that dividing a single element by 3 makes no sense. Here, you need to preserve the intention (what a moving average should do) while following the execution (what the code does) and detect the discrepancy.

Muse Spark identified the intention (a moving average) and detected the error. It proposed the right fix and explained why it is necessary. It even suggested a variant in case one wanted to ignore partial windows.

In the end, the model passed all three tests and leaves an excellent first impression.

## How to access Muse Spark?

You can access Muse Spark on meta.ai or via the Meta AI app on iOS and Android. Both are free. The rollout begins in the United States, with an expansion announced in the following weeks.

Meta plans a rollout at the same pace on WhatsApp, Instagram, Facebook, Messenger, and its Ray-Ban AI glasses.

There is no public API. A private preview is open to select enterprise partners, with no confirmed date for broader access. On privacy: Meta's policy sets few limits on the use of conversations to improve its models. If you are considering sharing sensitive information, read the terms first.

## Where Muse Spark falls short

Meta said it clearly in its technical post: the model has gaps on agent-driven multi-step tasks and on code workflows.

On SWE-Bench Verified, the gap with Gemini and Opus 4.6 is small. It widens on agentic work: Terminal-Bench 2.0 (59.0 versus 75.1 for GPT-5.4) and GDPval-AA for office automation (1,444 versus 1,672 for GPT-5.4). The gaps are clear.

Abstract visual reasoning follows the same pattern: ARC-AGI-2 shows 42.5 for Muse Spark, versus around 70 and higher for GPT-5.4 and Gemini. The model that excels at reading charts is outpaced on novel visual patterns.

This last point triggered a reaction on launch day. François Chollet, co-founder of the ARC Prize and creator of Keras and ARC-AGI, called the model "over-optimized for public benchmark numbers at the expense of everything else." Wang acknowledged the gap on ARC-AGI-2 and highlighted positive user feedback on code and visual reasoning. It remains to be seen whether this is confirmed at scale.

The absence of a public API, as mentioned above, adds a competitive deficit. Wang acknowledged this at launch: "There are certainly some behavioral rough edges that we will polish over time."

## Muse Spark safety

Meta conducted evaluations within the framework of its Advanced AI Scaling Framework before launch. On BioTIER-refuse, Muse Spark tops the comparison for refusing requests about biological weapons. These figures come from Meta.

Source: Meta Superintelligence Labs / ai.meta.com

The most interesting finding comes from Apollo Research. They observed that Muse Spark had the highest rate of evaluation sensitivity among the models tested: the model frequently identified a safety-test context and adapted its behavior accordingly.

A model that "behaves well" only when it knows it is being observed poses a real problem. Apollo's earlier work documented that this pattern can increase what they call "strategic behavior" in production.

Meta acknowledged this finding at launch, which few labs do. Their follow-up indicates that it affected a limited subset of alignment evaluations, unrelated to dangerous capabilities, and was not blocking. Research is ongoing.

## Muse Spark vs GPT-5.4 vs Opus 4.6 vs Gemini 3.1

Benchmarks show what these models can do. This section helps you choose which one to use in practice.

### At a glance

| **Feature** | **Muse Spark** | **GPT-5.4** | **Opus 4.6** | **Gemini 3.1 Pro** | 
| Release date | Apr. 8, 2026 | Mar. 5, 2026 | Feb. 5, 2026 | Feb. 19, 2026 | 
| Context window | 262 K* | 1.05 M | 1 M since Mar. 13 | 1 M | 
| Input modalities | Text, image, voice | Text, image | Text, image | Text, image, audio, video | 
| API price (per 1 M input/output tokens) | No public API | $2.50 / $15.00 | $5.00 / $25.00 | $2.00 / $12.00 | 
| Consumer access | meta.ai (US first) | ChatGPT | Claude.ai | Gemini app |

**Artificial Analysis indicates a context window of 262K for Muse Spark. Some sources cite 1M. Meta has not published any model card confirming either value.*

### Which one to choose?

Choose Muse Spark if your use cases involve health questions, chart reading, or multimodal consumer applications. There is no public API yet: if you need to integrate into production, you will have to wait.

Choose GPT-5.4 if you need a versatile model you can use today. It leads on code, abstract visual reasoning, and desktop automation, with a public API and a 1M window already available.

Choose Claude Opus 4.6 if you work on long documents or need polished and reliable writing. The 1M window moved to standard pricing on March 13, 2026. It is the most expensive option at $5/$25 per 1M tokens.

Choose Gemini 3.1 Pro if your pipeline processes video. It is the only model here to accept video input, and at $2/$12 per 1M tokens, it is the cheapest cutting-edge option in this group.

## What people are saying about Muse Spark

Early feedback is split as one might expect. Some found very surprising things. Others looked at the benchmark table and drew opposite conclusions.

The phrase "full stack rebuilt from scratch" came up often. Depending on how much trust you place in Meta, these nine months are either impressive or hard to believe.

Pietro Schirano shared a concrete example: he asked Muse Spark to convert a screenshot of an interface into code, and the model extracted the assets from the image instead of treating it as a simple bitmap.

This is not a benchmark; it is the kind of example that circulates because it is truly unexpected.

The most biting analysis is from Aakash Gupta: "It is the model of a data labeling CEO. His fingerprints are all over the results." The benchmarks where Muse Spark leads are all highly sensitive to data quality, where curation sets the ceiling.

Those where it trails (ARC-AGI-2, Terminal-Bench, GDPval) are precisely those where architecture and RL scaling matter more than data. His conclusion: "he designed the best model for what data pipelines solve, and an average model for the rest."

## Conclusion

The jump from 18 for Llama 4 Maverick to 52 for Muse Spark on the Artificial Analysis Intelligence Index is not subtle. For a team that rebuilt everything in nine months, the results in health and multimodal hold up, including under independent testing.

Admittedly, the gaps are glaring. On code and agentic tasks against GPT-5.4, the difference is significant; abstract visual reasoning is a clear weak point, and there is still no public API. If you need a model you can integrate today, Muse Spark is not yet the right choice.

What I keep coming back to is the question of open source. The Llama ecosystem relied on the trust that the weights would be available. Muse Spark breaks that contract. Wang's "hope" regarding the opening of future versions is not a commitment. This is, in my view, the most consequential aspect of this launch, too little discussed relative to the benchmark numbers.

Larger Muse models are in the works. If the architecture scales as advertised, today's numbers will seem modest. That is the bet.

To learn how to get the most out of any large language model, we recommend our Understanding Prompt Engineering course.

## FAQ about Muse Spark

### If I was using Llama locally, does Muse Spark replace that?

**No. Muse Spark is cloud-only. You cannot download it, run it on your own hardware, or fine-tune it. Access is through meta.ai or the Meta AI app, both of which require a Meta account. The open-weight usage around which Llama built its community does not exist here.**

### When should I use Contemplating mode rather than Thinking?

**Contemplating mode is most useful when the problem truly admits several valid solution paths: complex scientific questions, multi-step reasoning with ambiguous inputs, or research work where different angles lead to varied conclusions. For most everyday queries, Thinking mode is faster and the results are comparable. Another point: Contemplating mode is still rolling out gradually; you may not have access to it yet.**

### What does the "10 times less" compute promise concretely mean for me?

**Probably nothing for now. The comparison is against Muse Spark's previous model, not GPT-5.4 or Gemini, and the figure has not been independently verified. The most relevant point is inference efficiency: as noted above, Muse Spark used 58 million output tokens during Artificial Analysis's independent run, compared with 157 million for Claude Opus 4.6. This gap could one day be reflected in pricing, but API rates have not yet been announced.**

### Is it worth switching to Muse Spark from my current tool?

**If you use ChatGPT for general tasks, the day-to-day experience is similar. If your main use cases involve health, science, or chart analysis, Muse Spark is a reasonable upgrade. If you depend on coding assistants or long-document tools, it does not yet replace GPT-5.4 or Opus 4.6. The comparison table above details the specifics.**

### Should I worry about evaluation sensitivity?

**Not in everyday practice, but it is worth understanding. The finding is that Muse Spark adopts more cautious behavior when it detects a safety evaluation context, not because its "values" differ but because it recognizes the context. Meta's monitoring indicates that this affected a narrow subset of alignment tests, unrelated to dangerous capabilities. If you are evaluating models for sensitive deployments, read Apollo's full report before drawing conclusions from safety scores alone.**

I am a data engineer and community builder. I work on data pipelines, cloud, and AI tools, while writing practical and impactful tutorials for DataCamp and emerging developers.

I am a writer and editor in the field of data science. I am particularly interested in linear algebra, statistics, R, etc. I also play a lot of chess!

**Data Science Editor-in-Chief at DataCamp |** **I am passionate about forecasting and development using APIs.**
