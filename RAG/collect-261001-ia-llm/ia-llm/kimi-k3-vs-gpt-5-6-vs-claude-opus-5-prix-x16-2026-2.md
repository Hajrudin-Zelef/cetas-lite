---
id: collect-261001-ia-llm/ia-llm/kimi-k3-vs-gpt-5-6-vs-claude-opus-5-prix-x16-2026-2
title: "kimi-k3-vs-gpt-5-6-vs-claude-opus-5-prix-x16-2026"
domain: ia-llm
role: reference
task: reference
actors: ["Anthropic", "Hugging Face", "Moonshot", "OpenAI"]
dates: []
keywords: ["claude", "kimi", "agi", "attention", "benchmarks", "chatgpt", "fable 5", "gpt-5.6", "mai", "open source", "opus 4", "opus 5"]
source: docs/RAG/collect-261001-ia-llm/kimi-k3-vs-gpt-5-6-vs-claude-opus-5-prix-x16-2026.md
source_anchor: ""
source_lines: [33, 73]
sha256: 61317c910d586121172da0e637e38ca66bec9dc41350d0617ad269752a9ebaca
---

# kimi-k3-vs-gpt-5-6-vs-claude-opus-5-prix-x16-2026

Claude Opus 5 a été lancé le 24 juillet 2026, positionné par Anthropic comme son modèle de référence pour le codage agentique et l’utilisation d’ordinateur en autonomie. La différence majeure avec la génération précédente, Opus 4.8, tient dans le comportement par défaut : sur Opus 4.8, la réflexion ne s’activait que si le développeur définissait explicitement `thinking: {"type": "adaptive"}` dans sa requête. Sur Opus 5, cette réflexion adaptative tourne par défaut sur chaque appel, avec un niveau d’effort `high` appliqué automatiquement si rien n’est précisé.

Le paramètre `effort` accepte cinq valeurs, `low`, `medium`, `high`, `xhigh` et `max`, et gouverne l’ensemble des tokens produits par une réponse, pas seulement la partie visible de la réflexion. Un détail technique mérite l’attention des équipes qui cherchent à couper les coûts au maximum : désactiver la réflexion via `thinking: {"type": "disabled"}` n’est autorisé qu’aux niveaux d’effort `high` ou inférieur. Si un développeur tente de désactiver la réflexion tout en gardant un effort `xhigh` ou `max`, l’API renvoie une erreur 400. Ce garde-fou empêche de demander un raisonnement maximal sans en payer le prix.

Anthropic a maintenu la tarification d’Opus 5 strictement identique à celle d’Opus 4.8 : 5 dollars par million de tokens en entrée, 25 dollars par million de tokens en sortie, sans surcoût annoncé pour les niveaux d’effort élevés au-delà de la facturation classique des tokens de sortie. La fenêtre de contexte atteint 1 million de tokens, avec une sortie maximale de 128 000 tokens et une date de coupure des connaissances fixée à mai 2026. Cette stabilité tarifaire malgré l’activation par défaut du raisonnement constitue l’argument central d’Anthropic : le prix par token ne change pas, mais l’usage réel, lui, augmente mécaniquement pour les tâches complexes puisque le modèle réfléchit désormais systématiquement.

## Kimi K3 : le raisonnement permanent en poids ouverts

Moonshot AI, start-up pékinoise, a dévoilé Kimi K3 le 16 juillet 2026 via son API, avant de publier l’intégralité des poids du modèle sur Hugging Face onze jours plus tard, le 27 juillet. Avec 2,8 billions de paramètres au total, dont 104 milliards actifs par token grâce à une architecture Mixture-of-Experts à 896 experts, Moonshot AI revendique le titre de premier modèle ouvert de classe 3 000 milliards de paramètres et le plus grand système à poids ouverts jamais publié. L’architecture repose sur deux innovations techniques, l’attention Kimi Delta et les résidus d’attention, associées à une compréhension visuelle native qui place K3 dans la catégorie des modèles multimodaux agentiques.

Sur le plan du raisonnement, Kimi K3 se distingue radicalement des deux modèles américains : sa fiche technique officielle le décrit comme fonctionnant en mode de réflexion permanent, sans interrupteur pour le désactiver. Cette réflexion continue s’accompagne d’une fenêtre de contexte de 1 048 576 tokens, une taille identique quel que soit le volume de la requête puisque Moonshot AI applique une tarification plate sur l’ensemble de cette fenêtre, sans palier “long contexte” comme chez OpenAI.

Le tarif officiel de l’API Kimi K3 s’établit à 3,00 dollars par million de tokens en entrée pour un cache manqué, 0,30 dollar par million de tokens pour un cache retrouvé, soit une réduction de 90 %, et 15,00 dollars par million de tokens en sortie. Reuters et CNBC ont couvert le lancement en soulignant que Moonshot AI revendique des performances approchant celles du modèle Fable d’Anthropic sur le raisonnement avancé et le codage, un point que Digital Applied, un cabinet d’analyse indépendant, nuance en précisant que les résultats de K3 se situent globalement à un ou deux points des scores de Claude Fable 5 et de GPT-5.6 Sol sur plusieurs suites de tests agentiques, tout en dépassant Claude Opus 4.8 et GPT-5.5 sur la majorité de ces mêmes suites. Le modèle est distribué sous une licence de poids ouverts avec des restrictions d’usage spécifiques, ce qui le distingue d’un logiciel véritablement open source au sens strict.

## Tableau comparatif : specs techniques des trois modèles

Le tableau ci-dessous rassemble les caractéristiques techniques déterminantes pour un choix d’architecture, à partir des documentations officielles d’OpenAI, d’Anthropic et de Moonshot AI publiées entre juillet et septembre 2026.

| Critère | GPT-5.6 Sol (OpenAI) | Claude Opus 5 (Anthropic) | Kimi K3 (Moonshot AI) | 
|---|---|---|---|
| Date de lancement | Juillet 2026 (curseur ajouté le 6 août) | 24 juillet 2026 | 16 juillet 2026 (poids le 27 juillet) | 
| Contrôle du raisonnement | Curseur produit à 5 niveaux (ChatGPT) | Paramètre API thinking_effort | Aucun, toujours actif | 
| Niveaux disponibles | 5 (continu, interface uniquement) | low / medium / high / xhigh / max | Sans objet | 
| Raisonnement par défaut | Adaptatif selon le curseur | Actif, effort high par défaut | Toujours actif | 
| Fenêtre de contexte | ~1,05 million de tokens | 1 million de tokens | 1 048 576 tokens | 
| Sortie maximale | 128 000 tokens | 128 000 tokens | Non communiqué publiquement | 
| Paramètres totaux | Non divulgué | Non divulgué | 2,8 billions (104 Md actifs par token) | 
| Architecture | Non divulguée | Non divulguée | Mixture-of-Experts, 896 experts, 16 actifs | 
| Poids ouverts | Non | Non | Oui, licence modifiée | 
| Prix entrée standard | 4,00 $ / 1M tokens | 5,00 $ / 1M tokens | 3,00 $ / 1M tokens (cache manqué) | 
| Prix sortie standard | 20,00 $ / 1M tokens | 25,00 $ / 1M tokens | 15,00 $ / 1M tokens | 
| Prix entrée en cache | 0,40 $ / 1M tokens | Non communiqué publiquement | 0,30 $ / 1M tokens | 
| Surcoût long contexte | Oui, 2x au-delà de 272K tokens | Non annoncé jusqu’à 1M | Non, tarif plat sur toute la fenêtre | 
| Écosystème produit | ChatGPT, API, Codex | Claude.ai, API, Claude Code | Kimi.com, Kimi Work, Kimi Code, API | 

## Benchmarks : ce que les scores publics révèlent, et ce qu’ils cachent

Contrairement aux générations précédentes de modèles, où chaque lancement s’accompagnait d’un tableau détaillé de scores GPQA Diamond, MMLU ou SWE-bench, les trois fournisseurs comparés ici publient beaucoup moins de chiffres officiels pour leurs modèles de raisonnement les plus récents. Un guide indépendant sur les benchmarks de raisonnement IA, mis à jour début août 2026, note explicitement que pour la catégorie “frontier tier”, qui regroupe Claude Fable 5 et GPT-5.6 Sol, aucun score numérique n’est publié sur GPQA Diamond ou sur ARC-AGI-2, les fournisseurs se contentant d’indiquer que les meilleurs modèles se situent désormais dans les 90 % et plus, ce qui traduit une saturation progressive de ces benchmarks historiques.

Les chiffres qui existent proviennent en grande partie de la génération précédente ou d’évaluations tierces indépendantes, à prendre avec la prudence qui s’impose face à des méthodologies non standardisées.

