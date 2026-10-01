---
id: collect-261001-ia-llm/ia-llm/claude-mythos-5-fonctionnalites-benchmarks-et-capacites-3
title: "claude-mythos-5-fonctionnalites-benchmarks-et-capacites"
domain: ia-llm
role: reference
task: reference
actors: ["Anthropic", "Glasswing", "Google", "OpenAI"]
dates: []
keywords: ["benchmarks", "claude", "cyber", "fable 5", "gemini", "mythos 5", "opus 4"]
source: docs/RAG/collect-261001-ia-llm/claude-mythos-5-fonctionnalites-benchmarks-et-capacites.md
source_anchor: ""
source_lines: [137, 173]
sha256: 24ada09008309fd7ddb09e3397626a62e0c0e1b3fade917438a6418335b9d356
---

# claude-mythos-5-fonctionnalites-benchmarks-et-capacites

Un programme d’accès de confiance plus large est envisagé pour permettre aux organisations de cybersécurité de candidater de manière plus systématique, en concertation avec le gouvernement américain. Anthropic n’a pas communiqué de calendrier pour une disponibilité générale. Pour la plupart des développeurs, Claude Fable 5 est aujourd’hui l’option la plus pratique, avec le même modèle sous-jacent et un accès via les abonnements et l’API standard.

Un point opérationnel à noter : Anthropic a instauré une politique de rétention des données de 30 jours pour tout le trafic des modèles de classe Mythos. Les données ne sont pas utilisées pour l’entraînement et sont supprimées au bout de 30 jours dans quasiment tous les cas, mais elles sont conservées à des fins de sécurité. Si vous développez avec Mythos 5 sur des données sensibles, consultez la documentation d’assistance d’Anthropic sur cette politique avant le déploiement.

## Conclusion

Claude Mythos 5 est, à ce jour, la preuve la plus claire qu’Anthropic entend déployer une IA de pointe dans des contextes professionnels à forts enjeux, et les résultats le confirment.

L’écart sur SWE-bench Pro (80,3 % vs 69,2 %), sur Terminal-Bench 2.1 (88,0 % vs 82,7 %) et sur ExploitBench (78,0 % vs 40,0 %) pointe vers un modèle qui gère les tâches les plus difficiles de manière plus fiable que toute alternative disponible.

Le modèle d’accès restreint est une approche raisonnable au vu des risques de double usage, et les scores ExploitBench plaident pour que les outils offensifs les plus puissants ne soient pas accessibles au grand public. La vraie question est de savoir si Anthropic pourra étendre suffisamment vite le programme d’accès de confiance pour servir la communauté sécurité et biomédicale plus large, avant que les concurrents ne comblent l’écart.

Pour les organisations éligibles, la mise à niveau depuis Mythos Preview est simple et à moins de la moitié du prix.

## FAQ sur Claude Mythos 5

### Quelle est la différence entre Claude Mythos 5 et Claude Fable 5 ?

Mythos 5 et Fable 5 partagent la même architecture, mais diffèrent par leurs garde-fous de sécurité. Fable 5 redirige, via des classifieurs, les requêtes sensibles en cybersécurité et en biologie vers Claude Opus 4.8, tandis que Mythos 5 lève ces classifieurs pour des partenaires approuvés. La différence de nom reflète les garde-fous, pas les capacités.

### Qui peut accéder à Claude Mythos 5 ?

L’accès est actuellement limité à deux groupes : les partenaires cybersécurité de Project Glasswing, qui peuvent utiliser Mythos 5 avec les garde-fous cyber levés, et un petit groupe de chercheurs biomédicaux évalués, qui peuvent y accéder avec les garde-fous biologie et chimie levés, mais les garde-fous cyber maintenus. Anthropic prévoit d’étendre ces deux programmes, avec un programme d’accès de confiance plus large pour les organisations de cybersécurité, en concertation avec le gouvernement américain.

### Comment Claude Mythos 5 se compare-t-il à GPT-5.5 et Gemini 3.1 Pro ?

Mythos 5 devance les deux concurrents sur tous les benchmarks testés. Les écarts les plus importants se trouvent sur ExploitBench (78,0 % vs 34,0 % pour GPT-5.5), FrontierCode Diamond (29,3 % vs 5,7 %) et Humanity's Last Exam avec outils (64,5 % vs 52,2 %). Gemini 3.1 Pro est encore plus en retrait sur la plupart des benchmarks.

### Claude Mythos 5 est-il sûr à utiliser avec des données sensibles ?

Anthropic a instauré une politique de rétention des données de 30 jours pour tout le trafic des modèles de classe Mythos. Les données ne sont pas utilisées pour l’entraînement et sont supprimées après 30 jours dans presque tous les cas, mais elles sont conservées à des fins de supervision de la sécurité. Les organisations manipulant des données sensibles doivent consulter la documentation d’assistance d’Anthropic sur cette politique avant de déployer.

### Que signifie la politique de rétention des données de 30 jours pour les utilisateurs de Mythos 5 ?

Contrairement à l’usage API standard, où les données ne sont pas conservées, le trafic de la classe Mythos est retenu jusqu’à 30 jours pour supervision de la sécurité avant suppression. Cela s’applique à tous les appels API Mythos 5 et n’est pas utilisé pour l’entraînement. C’est un point opérationnel important pour toute organisation qui déploie le modèle en production avec des données confidentielles ou réglementées.

**Rédacteur en chef Data Science chez DataCamp |** **Je suis passionné par la prévision et le développement à l'aide d'API.**
