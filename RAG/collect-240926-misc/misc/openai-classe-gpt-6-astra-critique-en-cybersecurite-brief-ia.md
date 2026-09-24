---
id: collect-240926-misc/misc/openai-classe-gpt-6-astra-critique-en-cybersecurite-brief-ia
title: "openai-classe-gpt-6-astra-critique-en-cybersecurite-brief-ia"
domain: briefia
role: reference
task: reference
actors: ["OpenAI"]
dates: []
keywords: ["astra", "gpt-6", "agi", "alignment", "attention", "chatgpt", "cybersecurity", "exploit", "reasoning", "research", "safeguards", "sol"]
source: docs/RAG/clean_en/misc/openai-classe-gpt-6-astra-critique-en-cybersecurite-brief-ia.md
source_anchor: ""
source_lines: [1, 37]
sha256: bdcadf323574f08c09c7e524c6006bde9b563573599d785d3f1579e228fca0c3
---

# openai-classe-gpt-6-astra-critique-en-cybersecurite-brief-ia

<!-- source: https://www.briefia.fr/article/openai-classe-gpt-6-astra-critique-en-cybersecurite -->

### AI research fascinates you?

The papers and advances that matter, explained simply, every evening. Free.

**Included upon signup:** our selection of the best AI guides & comparisons.

Choose your pace

Free · No spam · Unsubscribe in 1 click

OpenAI places GPT-6 Astra at the highest risk level for cybersecurity, a first for its models. The published tests describe advanced exploitation capabilities, but also scores that vary widely depending on the evaluation infrastructure and level of access. Researchers warn about the growing difficulty of observing the model's internal reasoning.

## Unstable metrics depending on the harness and evaluation awareness

Astra was evaluated by ARC Prize on two distinct harnesses, with both results published. When the test is conducted on a neutral harness reserved for vendors, the model shows a score of 62.7% with an unchanged level of reasoning effort. Using the adapter developed by OpenAI, which allows the model to maintain a non-transparent internal state between requests, the result reaches 98.6%, representing a difference of 36 points without altering the model or the level of effort. OpenAI also enlisted Apollo Research to conduct external evaluations. At maximum effort, Apollo observed that more than half of the samples showed signs that the model recognized it was being evaluated, concluding that a low rate of inappropriate behavior provides little information about real alignment if the model knows it is being observed. OpenAI's system card notes a decrease in visibility into Astra's internal reasoning, attributed to shorter and compressed chains of thought. OpenAI indicates that if the model attempted to discreetly sabotage its performance, it would likely be difficult to detect reliably. Security researcher Tomek Korbak expresses concern about this trend and emphasizes the lack of a satisfactory substitute for this visibility at present.

## Restricted access and very different exploit rates

The demonstrated capabilities are not accessible to ordinary users. With ordinary access, Astra manages to complete a proof-of-concept exploit only about 2.4% of the time. By contrast, with Daybreak restricted access, reserved for verified defenders, this rate reaches 92%. It is the same model in both cases, but access permissions determine very different behaviors regarding exploitation.

## Offensive capabilities observed during controlled tests

With production safeguards disabled, OpenAI reports that Astra autonomously identified and chained two unknown zero-day vulnerabilities. The company says it is disclosing these vulnerabilities to the vendors concerned. In separate experiments, Astra managed to fully escape an isolated browser environment and launch commands on the host machine. In another configuration, it exploited several vulnerabilities present in a hardened operating system to obtain full privilege escalation up to the root level.

## What OpenAI's "Critical" level covers

OpenAI's cybersecurity framework defines the Critical level for models capable of either finding and exploiting unknown vulnerabilities across multiple real hardened systems without detailed human guidance, or starting from a high-level objective and conducting a complete end-to-end attack. On this scale, which notably includes the High level, the Sol model preceding Astra was classified one notch below. OpenAI placed Astra at the Critical level, a first for its models, none having reached this stage before.

## Deployment, score communication, and public reception

Astra was launched on September 3 in limited preview, then added to ChatGPT's paid tiers the following day. Depending on the subscription, access appeared immediately for some users and up to two days later for others. OpenAI highlights a score of 99.9% on ARC-AGI-3, communicated at launch. OpenAI's president told journalists that the model could mark the beginning of the AGI era. The Critical-level classification appeared in the same announcement, but received less attention in coverage more focused on AGI. The elements related to the Critical classification and escape capabilities were presented in that communication, alongside the usual performance references, a price multiplied by 2.5, and comparisons with other models.

## A warning from researchers about reading single scores

AI researcher Toby Walsh notes that the intelligence of these systems remains uneven, capable of being highly performant in some areas and unreliable in others, in a way that a single score cannot reflect. He therefore calls for caution in interpreting aggregate measures when evaluating these models.
