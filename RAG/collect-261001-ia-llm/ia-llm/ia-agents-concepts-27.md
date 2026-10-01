---
id: collect-261001-ia-llm/ia-llm/ia-agents-concepts-27
title: "Concepts : agents IA, agentic, autonomie"
domain: ia-llm
role: reference
task: reference
actors: []
dates: []
keywords: ["rlhf"]
source: docs/RAG/collect-261001-ia-llm/ia_agents_concepts.md
source_anchor: ""
source_lines: [3714, 3736]
sha256: c3a57dac37d2c9f17ec55a785d968e581deb96928cf07496436a0f1797c2928a
---

# Concepts : agents IA, agentic, autonomie

- Avant (tout-LLM) :
```
  événement → LLM (raisonne + génère + décide, 3–329 s, $$$)
              → parse le JSON → action
```
  Chaque micro-décision paie le prix fort d'une génération complète, avec
  le risque de JSON malformé et de surconfiance polie (RLHF).
- Après (Jev en couche de décision) :
```
  événement → Jev (state + questions typées, 70–500 ms, ~0 $)
              ├─ confidence ≥ seuil haut → action directe
              ├─ zone grise           → LLM (System Two : raisonne, génère)
              └─ risque high / conf basse → humain dans la boucle
```
- **Le LLM devient l'exception, pas la règle** : il ne traite que l'ambigu
  et le nouveau. Jev absorbe le volume (routage, scoring, garde-fous) à
  coût marginal quasi nul. Les seuils de `confidence` sont **tes** règles
  métier, versionnées dans ton code — pas enfouies dans un prompt.
- **Règle d'or** : épingle la version du modèle (`jev-1.13.0`, pas
  `jev-latest`) une fois tes seuils calibrés — un alias mouvant + des
  seuils fixes = dérive silencieuse.

---
