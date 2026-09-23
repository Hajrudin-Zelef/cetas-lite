---
id: vague2-vision-ia/vision-ia/openai-re-sout-un-proble-me-du-mille-naire-un-mathe-maticien-l-accuse-de-l-avoir-vole
title: "OpenAI résout un problème du millénaire, un mathématicien l'accuse de l'avoir volé"
domain: vision-ia
role: reference
task: article
actors: ["Alibaba", "Anthropic", "Apple", "China", "Google", "Huawei", "Meta", "OpenAI", "United States"]
dates: ["2026-09-09", "2026-09-23"]
keywords: ["agent", "agents", "alignment", "astra", "chatgpt", "claude", "compute", "gpt-5.6", "latency", "lawsuit", "moratorium", "muse"]
source: docs/RAG/Collect RAG Vague 2/04_vision_ia/openai-re-sout-un-proble-me-du-mille-naire-un-mathe-maticien-l-accuse-de-l-avoir-vole.md
source_anchor: ""
source_lines: [1, 47]
sha256: 0e17c67929f4df5910d3433bb89d2203a932b76010f585f8d3823affbe5692fd
---

# OpenAI résout un problème du millénaire, un mathématicien l'accuse de l'avoir volé

## Metadata

- **Source** : https://vision-ia.beehiiv.com/p/openai-re-sout-un-proble-me-du-mille-naire-un-mathe-maticien-l-accuse-de-l-avoir-vole
- **Site** : Vision-IA (beehiiv)
- **Type** : Newsletter article
- **Language** : fr (original) / summary in English
- **Verification status** : ✅ reachable
- **Collection date** : 2026-09-23

## Full summary

This September 9, 2026 issue leads with OpenAI announcing it proved that the Navier-Stokes equations—describing water and airflow—can indeed break down and predict a physically impossible state, such as a fluid at infinite speed. This would be the second of the seven Clay Institute Millennium Problems to fall since their 2000 selection, after the Poincaré conjecture. But the announcement was eclipsed within hours by a dossier from Tristan Buckmaster, a mathematician at NYU, who accuses the company of having scooped him on a lifetime problem. The proof was obtained with a non-public internal model presented as clearly superior to Astra, released the previous week. Method: about 10,000 agents launched in parallel, for several million dollars of compute. OpenAI says it will not claim the attached $1 million prize. Buckmaster and Levent Alpöge, an Anthropic employee, had worked on it for nearly a year with both houses' public models. On Monday, Buckmaster posted a proof on Mastodon concerning a simplified version, the Euler equations. According to his detailed account, OpenAI left him two options: publish the day before being crushed by the announcement, or co-sign a paper excluding Alpöge because he works for a competitor. He refused both and claims he was then threatened. He especially asks whether OpenAI's agents had access to transcriptions of his work. Sam Altman denies it, defends his researcher Sébastien Bubeck ("he acted with integrity"), and says he only discovered their paper after publication. Buckmaster had deposited all his drafts in Codex; OpenAI assured him the model did not consult user data but did not answer his training question. The episode establishes that the greatest open math problems are now solved with compute only two or three private labs possess.

Next, Jacob Coxon, a 27-year-old pretraining researcher, announced on X that he was leaving Anthropic and the AI sector entirely, accusing his employer and OpenAI of "running straight toward a self-improving superintelligence while playing with our lives." Hours later, Evan Hubinger, who leads alignment research at Anthropic, publicly replied: "Jacob is right, we sincerely believe AI could kill all humans," then quantified it: more than 10% within the coming decade. Hubinger judges current models' risk low; his concern is recursive self-improvement. His heaviest admission: Anthropic "does not yet have a plan to solve superintelligence alignment" and "is not clearly on track." Coxon calls for coordination among US labs, even via a temporary moratorium on capability improvement.

Other items: ChatGPT Images 2.5 deployed September 8 with up to 50% lower generation latency, comment-based editing, Sketch, and templates, with API models GPT-Image-2.5 Flare and Sunburst at unchanged pricing ($5/$8/$30 per M text/image input/image output). DeepMind published the AlphaGenome Atlas: 9 billion precomputed single-point variants, each with ~27,000 predictions, 1 petabyte of data, and a new AVI score; the Broad Institute used it to identify a non-coding variant causing severe epilepsy. Research briefs cover Qwen3.8 27B quantization (4-bit matches BF16, 1-bit collapses), Danijar Hafner's world-model robotics startup, Terence Tao warning that open problems are a non-renewable resource, and GPT-5.6 Sol running quantum experiments at MIT. Industry briefs cover LibreOffice's no-AI record, Meta's Muse agent in WhatsApp, Adobe's generative media in Premiere, Arm's robotics standard, IFA Berlin's robot fashion show, Meta dropping token-count metrics, Chinese micro-drama face licensing, Wispr Notetaker, Inworld TTS-2, Chrome's two-week cycle, Claude token theft, Google's European search changes, ASML/Huawei, an Anthropic lawsuit, the Sam Altman biopic "Artificial," Patagonian data centers, Starship 40, the iPhone Duo, Cognition's $2B raise, and Marvell.

## Key points

- OpenAI claims to have proven Navier-Stokes can break down, its second Millennium Problem, using ~10,000 parallel agents and millions in compute.
- Tristan Buckmaster accuses OpenAI of scooping him and questions whether his Codex drafts were used for training; OpenAI denies it.
- Anthropic's alignment lead Evan Hubinger publicly put >10% probability on AI killing all humans this decade and said Anthropic lacks a superintelligence alignment plan.
- Researcher Jacob Coxon left Anthropic, calling for a possible capability moratorium.
- ChatGPT Images 2.5 launched with faster generation and comment-based editing at unchanged API pricing.
- DeepMind's AlphaGenome Atlas precomputes 9 billion DNA variants with ~27,000 predictions each.

## Technical data / figures

| Item | Value |
| --- | --- |
| Navier-Stokes proof compute | ~10,000 parallel agents, several $M |
| Millennium prize (declined) | $1M |
| Hubinger extinction risk estimate | >10% in the decade |
| ChatGPT Images 2.5 latency reduction | up to 50% |
| GPT-Image-2.5 API pricing | $5 / $8 / $30 per M (text in / image in / image out) |
| AlphaGenome Atlas variants | 9 billion single-point variants |
| AlphaGenome predictions per variant | ~27,000 |
| AlphaGenome data volume | 1 petabyte (30x AlphaFold) |
| AVI score scale | 10 = top 10%, 30 = top 0.1% |
| Cognition valuation | $48B (raised $2B) |
| Cognition annualized revenue | $492M → $900M since May |

## Why this source matters for the RAG

It captures two pivotal controversies: AI's role in solving a Millennium Problem and the credit/data-provenance dispute with Buckmaster, plus Anthropic's alignment lead publicly quantifying catastrophic risk. These are central to AI-safety, research-ethics, and capability-trend queries. The AlphaGenome and image-generation specs add valuable technical reference data.
