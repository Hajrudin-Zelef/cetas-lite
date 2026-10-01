---
id: collect-261001-ia-llm/ia-llm/gpt-6-astra-vs-claude-fable-5-1-performances-et-tarifs-5
title: "gpt-6-astra-vs-claude-fable-5-1-performances-et-tarifs"
domain: ia-llm
role: reference
task: reference
actors: ["AWS", "Anthropic", "OpenAI"]
dates: []
keywords: ["astra", "claude", "gpt-6", "bedrock", "fable 5"]
source: docs/RAG/collect-261001-ia-llm/gpt-6-astra-vs-claude-fable-5-1-performances-et-tarifs.md
source_anchor: ""
source_lines: [300, 302]
sha256: 278015a38112b76f4bd0456e313b0d85c778bc4d1f53d247b44c3f39821dd877
---

# gpt-6-astra-vs-claude-fable-5-1-performances-et-tarifs

Les IDs de modèle sont `gpt-6-astra` pour OpenAI et `claude-fable-5-1` pour Anthropic. Sur Amazon Bedrock, Claude Fable 5.1 est `anthropic.claude-fable-5-1`. Notez trois changements API entre Fable 5 et 5.1 : la sélection forcée d’outil renvoie un 400, les anciens modèles ne lisent pas ses blocs de pensée, et ces blocs sont liés exactement à l’historique précédent. Les refus arrivent en HTTP 200 avec `stop_reason: \"refusal\"`.

**Rédacteur en chef Data Science chez DataCamp |** **Je suis passionné par la prévision et le développement à l'aide d'API.**
