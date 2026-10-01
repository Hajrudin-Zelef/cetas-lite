---
id: collect-261001-ia-llm/ia-llm/claude-mythos-5-fonctionnalites-benchmarks-et-capacites-2
title: "claude-mythos-5-fonctionnalites-benchmarks-et-capacites"
domain: ia-llm
role: reference
task: reference
actors: ["Anthropic", "Glasswing", "Google", "OpenAI"]
dates: []
keywords: ["benchmark", "benchmarks", "claude", "agent", "cyber", "fable 5", "gemini", "jailbreak", "mythos 5", "opus 4"]
source: docs/RAG/collect-261001-ia-llm/claude-mythos-5-fonctionnalites-benchmarks-et-capacites.md
source_anchor: ""
source_lines: [58, 136]
sha256: 144fd042242e198907a222f8b491e7e561dcc1e0b2f7764f58e4315d7fa8bcee
---

# claude-mythos-5-fonctionnalites-benchmarks-et-capacites

| Catégorie | Benchmark | Claude Mythos 5 / Fable 5 | Claude Mythos Preview | Claude Opus 4.8 | GPT 5.5 | Gemini 3.1 Pro | 
|---|---|---|---|---|---|---|
| Code agentique | SWE-Bench Pro | 80,3 % | 77,8 % | 69,2 % | 58,6 % | 54,2 % | 
| Code agentique | FrontierCode (Diamond) | 29,3 % (xhigh) | — | 13,4 % (xhigh) | 5,7 % (xhigh) | — | 
| Travail de connaissance | GDPval-AA | 1932 | — | 1890 | 1769 | 1314 | 
| Travail de connaissance vision | GDP.pdf | 29,8 % (sans outils) | — | 22,5 % (sans outils) | 24,9 % (sans outils) | 16,7 % (sans outils) | 
| Raisonnement spatial | Blueprint-Bench 2 | 38,6 % | — | 14,5 % | 36,2 % | 26,5 % | 
| Utilisation d’outils | AutomationBench | 17,4 % | — | 15,5 % | 12,9 % | 9,6 % | 
| Utilisation de l’ordinateur | OSWorld-Verified | 85,0 % | 85,4 % | 83,4 % | 78,7 % | 76,2 % | 
| Légal | Legal Agent Benchmark | 13,3 % | — | 10,4 % | 2,1 % | 0,0 % | 
| Raisonnement pluridisciplinaire | Humanity's Last Exam (sans outils) | 59,0 %* | 56,8 % | 49,8 % | 41,4 % | 44,4 % | 
| Raisonnement pluridisciplinaire | Humanity's Last Exam (avec outils) | 64,5 %* | 64,7 % | 57,9 % | 52,2 % | 51,4 % | 
| Biologie | BioMysteryBench (difficile) | 46,1 %* | 29,6 % | 40,0 % | — | — | 
| Biologie | BioMysteryBench (résolu par l’humain) | 83,9 %* | 82,6 % | 80,4 % | — | — | 
| Code agentique | Terminal-Bench 2.1 | 88,0 %* | — | 82,7 % | 83,4 % (Codex CLI) | 70,7 % (Gemini CLI) | 
| Cybersécurité | ExploitBench (Cap%) | 78,0 %* | 69,0 % | 40,0 % | 34,0 % | — | 
| Santé | HealthBench Professional | 66,0 %* | 64,7 % | 56,9 % | 51,8 % | — | 

Anthropic annonce des scores communs pour Mythos 5 et Fable 5, en précisant qu’ils varient généralement de 1 à 3 points. Les benchmarks marqués d’un astérisque (*) affichent un écart plus important, car les classifieurs de sécurité de Fable 5 redirigent les requêtes sensibles vers Opus 4.8 ; sur ces benchmarks, Fable 5 se rapproche des performances de la classe Opus.

### Code agentique : SWE-Bench Pro, FrontierCode et Terminal-Bench 2.1

Sur SWE-Bench Pro, Mythos 5 atteint 80,3 %, contre 77,8 % pour Mythos Preview, 69,2 % pour Opus 4.8, 58,6 % pour GPT 5.5 et 54,2 % pour Gemini 3.1 Pro. L’écart de 11 points avec Opus 4.8 est significatif sur un benchmark conçu pour éviter les fuites de vérité terrain.

Pour la qualité et la maintenabilité du code agentique, mesurées par FrontierCode (Diamond), l’écart est encore plus marqué. Mythos 5 obtient 29,3 % au niveau d’effort xhigh, contre 13,4 % pour Opus 4.8 et 5,7 % pour GPT 5.5.

Pour le travail en terminal, Mythos 5 redonne l’avantage à Anthropic face à OpenAI : Mythos 5 obtient 88,0 %* sur Terminal-Bench 2.1, contre 82,7 % pour Opus 4.8, 83,4 % pour GPT 5.5 (Codex CLI) et 70,7 % pour Gemini 3.1 Pro (avec Gemini CLI).

### Travail de connaissance : GDPval-AA et GDPpdf

GDPval-AA évalue le travail de connaissance sur une échelle numérique. Mythos 5 obtient 1932, contre 1890 pour Opus 4.8, 1769 pour GPT 5.5 et 1314 pour Gemini 3.1 Pro.

L’écart s’accentue pour le travail de connaissance sur PDF sans accès aux outils. Mythos 5 atteint 29,8 % sur GDPpdf, contre 22,5 % pour Opus 4.8, 24,9 % pour GPT 5.5 et 16,7 % pour Gemini 3.1 Pro.

### Raisonnement pluridisciplinaire : Humanity's Last Exam

Humanity's Last Exam (HLE) teste un raisonnement de niveau master en sciences, mathématiques et sciences humaines. Mythos 5 obtient 59,0 %* sans outils et 64,5 %* avec outils. Mythos Preview atteint respectivement 56,8 % et 64,7 % — pratiquement à égalité avec outils mais à 2 points derrière sans outils. L’écart avec Opus 4.8 est déjà marqué (49,8 % sans, 57,9 % avec), et encore plus avec les modèles phares concurrents (GPT 5.5 : 41,4 % et 52,2 %, Gemini 3.1 Pro : 44,4 % et 51,4 %).

L’avantage de Mythos 5 apparaît le plus nettement sans outils, où il devance Opus 4.8 de plus de 9 points. Ces scores sont étoilés, ce qui signifie que Fable 5 performe un peu moins en raison de ses classifieurs de sécurité.

### Utilisation de l’ordinateur, outils et raisonnement spatial

Sur OSWorld-Verified, qui évalue la capacité du modèle à réaliser des tâches sur une véritable interface ordinateur, Mythos 5 atteint 85,0 %. Mythos Preview le devance légèrement à 85,4 %, ce qui en fait le seul benchmark où Mythos Preview est en tête. Opus 4.8 suit de près (83,4 %), tandis que les concurrents décrochent : GPT 5.5 à 78,7 % et Gemini 3.1 Pro à 76,2 %.

AutomationBench mesure l’aptitude à utiliser des outils. Mythos 5 atteint 17,4 %, contre 15,5 % pour Opus 4.8, 12,9 % pour GPT 5.5 et 9,6 % pour Gemini 3.1 Pro. Les valeurs absolues faibles suggèrent que l’orchestration d’outils demeure un défi pour tous les modèles de pointe.

Le raisonnement spatial est un domaine où l’avance de Mythos 5 est la plus nette. Il obtient 38,6 % sur Blueprint-Bench 2, soit plus du double des 14,5 % d’Opus 4.8. GPT 5.5 est plus proche à 36,2 %, et Gemini 3.1 Pro atteint 26,5 %.

### Cybersécurité et biologie

Ce sont sans doute les deux domaines les plus mis en avant dans les notes de version, et les résultats l’expliquent.

ExploitBench mesure la fraction d’exploits que le modèle peut reproduire avec succès (Cap%). Mythos 5 atteint 78,0 %*, une progression notable par rapport à Mythos Preview (69,0 %) et un bond spectaculaire face à Opus 4.8 (40,0 %) et GPT 5.5 (34,0 %).

L’écart de 38 points avec Opus 4.8 est le plus important du tableau comparatif, et il justifie l’existence de garde-fous cyber pour Fable 5. Les tests d’intrusion externes d’Anthropic n’ont trouvé aucun jailbreak universel sur des tâches agentiques longues, même si l’UK AISI a progressé vers l’un d’eux lors d’une première fenêtre de test.

BioMysteryBench évalue le raisonnement biologique à deux niveaux de difficulté. Sur le sous-ensemble difficile, Mythos 5 obtient 46,1 %*, contre 29,6 % pour Mythos Preview et 40,0 % pour Opus 4.8. Sur le sous-ensemble « résolu par l’humain », Mythos 5 atteint 83,9 %*, Mythos Preview 82,6 % et Opus 4.8 80,4 %. GPT 5.5 et Gemini 3.1 Pro n’affichent pas de scores rapportés sur ces sous-ensembles.

Comme pour ExploitBench, les scores de Fable 5 se rapprochent d’Opus 4.8 en raison de ses classifieurs de sécurité liés à la biologie.

### Santé et juridique

Claude Mythos 5 montre une solidité notable dans deux domaines professionnels à forts enjeux, où la précision et la qualité du raisonnement ont des conséquences concrètes : la médecine et le droit.

Sur HealthBench Professional, Mythos 5 obtient 66,0 %*, un peu au-dessus des 64,7 % de Mythos Preview. Opus 4.8 atteint 56,9 % et GPT 5.5 51,8 %.

Sur le Legal Agent Benchmark, Mythos 5 atteint 13,3 %, contre 10,4 % pour Opus 4.8 et seulement 2,1 % pour GPT 5.5. Les scores absolus restent faibles, mais l’écart avec GPT 5.5 et Gemini est marqué. Le raisonnement juridique demeure un défi pour tous les modèles.

## Tarification et disponibilité de Claude Mythos 5

Claude Mythos 5 est facturé 10 $ par million de tokens en entrée et 50 $ par million de tokens en sortie. C’est moins de la moitié du prix de Claude Mythos Preview (25 $/125 $), ce qui facilite la mise à niveau pour les partenaires Glasswing existants. Les développeurs peuvent accéder au modèle via l’API Claude avec l’identifiant `claude-mythos-5`.

L’accès est actuellement limité à deux groupes :

- Tous les utilisateurs qui avaient accès à Claude Mythos Preview via Project Glasswing peuvent passer à Mythos 5 avec les garde-fous cyber levés
- Un petit groupe de chercheurs biomédicaux qui peuvent accéder à Mythos 5 avec les garde-fous biologie et chimie levés, mais les garde-fous cyber maintenus.

Anthropic prévoit d’étendre ces deux programmes au fil du temps.

