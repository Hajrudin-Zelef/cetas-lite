---
id: collect-240926-datacamp/datacamp/muse-spark-caracteristiques-benchmarks-et-mode-demploi-2
title: "muse-spark-caracteristiques-benchmarks-et-mode-demploi"
domain: datacamp
role: reference
task: reference
actors: ["Anthropic", "Apple", "Google", "Meta", "OpenAI", "United States"]
dates: []
keywords: ["benchmark", "benchmarks", "muse", "agent", "agentic", "agi", "alignment", "attention", "claude", "compute", "gemini", "muse spark"]
source: docs/RAG/clean_en/datacamp/muse-spark-caracteristiques-benchmarks-et-mode-demploi.md
source_anchor: ""
source_lines: [91, 197]
sha256: b66e39c1f1a97ff532bc5ef383eef0c54d5486cd0c019675bf8dd2c13ede0770
---

# muse-spark-caracteristiques-benchmarks-et-mode-demploi

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

