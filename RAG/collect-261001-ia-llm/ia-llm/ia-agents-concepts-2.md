---
id: collect-261001-ia-llm/ia-llm/ia-agents-concepts-2
title: "Concepts : agents IA, agentic, autonomie"
domain: ia-llm
role: reference
task: reference
actors: ["Microsoft"]
dates: []
keywords: ["agent", "benchmark", "incident", "leaderboard", "mcp", "reasoning"]
source: docs/RAG/collect-261001-ia-llm/ia_agents_concepts.md
source_anchor: ""
source_lines: [145, 300]
sha256: e785889762857ceb4c9361667e99914a458915ae772a1537e9b76527fde860a8
---

# Concepts : agents IA, agentic, autonomie

1. Tu déclares les outils au modèle : nom, description, schéma JSON des paramètres.
2. Le modèle, au lieu de répondre en texte, émet un bloc structuré :
   `{"name": "run_shell", "arguments": {"command": "df -h"}}`.
3. **C'est ton code (le harness) qui exécute vraiment la fonction.** Le modèle n'exécute rien.
4. Tu renvoies la sortie au modèle comme un message « tool ».
5. Le modèle continue : nouveau raisonnement, nouvel appel, ou réponse finale.

Conséquences pratiques :

- Le modèle ne peut appeler que ce que tu as déclaré. Pas d'outil déclaré = pas d'action possible.
- La description de l'outil fait partie du prompt système : un attaquant qui contrôle
  la description (ex : description d'un outil MCP distant, alerte Microsoft juin 2026)
  contrôle partiellement le comportement de l'agent.
- Le typage des paramètres est ta première ligne de défense : paramètres stricts,
  valeurs énumérées quand c'est possible, longueurs bornées.
- Le function calling n'est pas fiable à 100 % : le modèle peut halluciner des paramètres,
  appeler le mauvais outil, ou inventer un outil inexistant. Ton harness doit valider
  chaque appel avant exécution (nom connu ? paramètres conformes au schéma ?).

Le benchmark de référence pour la qualité du function calling est **BFCL** (Berkeley
Function-Calling Leaderboard) : appels simples, multi-tours, multi-étapes.

## 7. ReAct : le pattern fondateur

ReAct (Reasoning + Acting, Yao et al., 2022 — papier « ReAct: Synergizing Reasoning
and Acting in Language Models ») est le pattern historique : le modèle alterne
explicitement des traces de **raisonnement** (Thought) et des **actions** (Action → Observation).

Format canonique :

```
Thought: je dois vérifier l'espace disque avant de proposer un nettoyage.
Action: run_shell
Action Input: {"command": "df -h"}
Observation: /dev/sda1  91%  /var
Thought: /var est presque plein, je regarde ce qui prend de la place.
Action: run_shell
Action Input: {"command": "du -sh /var/log/* | sort -rh | head -5"}
Observation: 4,2G /var/log/syslog.1 ...
Thought: les logs tournent, je propose une rotation + compression.
Final Answer: ...
```

Pourquoi c'est important pour toi :

- C'est le pattern le plus simple à implémenter soi-même (voir partie G).
- Il rend la boucle **auditable** : chaque décision est tracée en langage naturel.
- Ses limites sont connues : verbeux (tokens), pas de vraie planification long terme,
  sensible à la qualité du prompt.
- La plupart des frameworks modernes sont des ReAct améliorés (réflexion, planification,
  mémoire) — comprendre ReAct, c'est comprendre 80 % des architectures.

## 8. Réflexion : l'agent qui se relit

La réflexion (terme popularisé par le papier « Reflexion », Shinn et al., 2023) ajoute
une étape : après un échec ou un résultat médiocre, l'agent génère une **critique
verbale** de sa propre trajectoire et la conserve en mémoire pour l'essai suivant.

```
Essai 1 : échec (mauvais outil, mauvais paramètre)
Réflexion : « j'ai appelé delete_old_logs sans vérifier la date ; la prochaine
fois je liste d'abord avec --dry-run »
Essai 2 : utilise la leçon → succès
```

En pratique :

- Ça marche surtout sur des tâches répétables avec un signal d'échec clair
  (tests qui passent/échouent, commande qui retourne un code d'erreur).
- Ça coûte cher : chaque réflexion = un appel modèle supplémentaire + du contexte.
- Pour un sysadmin, l'équivalent manuel c'est ton runbook : « la dernière fois que
  ce playbook a planté, c'était parce que… ». La réflexion automatise la mise à jour
  du runbook, avec les mêmes risques d'apprentissage de travers.
- Variante utile : la réflexion **bornée** — une seule passe de critique, pas de boucle
  de réflexion infinie (qui est un piège classique, section 106).

## 9. Planification : plan-then-execute vs entrelacé

Deux écoles pour organiser le travail :

**A. Plan-then-execute (planifier puis exécuter)**
1. Le modèle produit un plan complet : étapes 1 à N.
2. Un exécuteur déroule les étapes, avec replanification si une étape échoue.
3. Avantages : plan lisible, validable par un humain **avant** exécution, budget prévisible.
4. Inconvénients : plan fragile si l'environnement change ; replanification coûteuse.

**B. Entrelacé (interleaved, style ReAct)**
1. Le modèle décide de l'étape suivante à chaque tour, en fonction de l'observation.
2. Avantages : s'adapte au réel, pas de plan à jeter.
3. Inconvénients : peut dériver, difficile à auditer a priori, budget imprévisible.

Recommandation pratique (terrain) :

- Tâche dangereuse ou coûteuse (suppression, modification réseau, redémarrage service)
  → **plan-then-execute + validation humaine du plan**.
- Tâche exploratoire en lecture seule (diagnostic, inventaire, recherche doc)
  → **entrelacé**, avec budget strict.
- Hybride courant en 2026 : plan global validé une fois, puis exécution entrelacée
  par sous-tâche avec budgets locaux. C'est ce que font les bons harnesses.

## 10. Quand un simple prompt suffit

N'utilise pas d'agent quand :

- [ ] la tâche tient en un appel modèle (résumer, classer, extraire, traduire) ;
- [ ] les étapes sont fixes et connues d'avance (alors c'est un script, pas un agent) ;
- [ ] tu as besoin de déterminisme total (facturation, conformité, sécurité) ;
- [ ] la latence compte (un agent = N appels modèle en série) ;
- [ ] le coût doit être prévisible au centime près.

Règle simple : **si tu peux écrire la séquence d'étapes à l'avance, écris un script.**
L'agent sert quand la séquence dépend de ce qu'on découvre en route.

Exemples « prompt suffit » côté sysadmin :
résumer un log, générer une config à partir d'un template, expliquer une erreur,
convertir un fichier, rédiger un compte-rendu à partir de notes.

## 11. Quand il faut un agent (checklist de décision)

Passe à l'agent quand tu coches au moins 3 de ces cases :

- [ ] l'ordre des étapes dépend des résultats intermédiaires ;
- [ ] il faut combiner plusieurs sources (logs + métriques + doc + commandes) ;
- [ ] la tâche demande des allers-retours (essai → erreur → correction) ;
- [ ] le périmètre est flou au départ (« débrouille-toi pour que ça remarche ») ;
- [ ] un humain ferait naturellement des boucles d'exploration.

Exemples « agent justifié » :
diagnostic d'incident multi-sources, tri de tickets avec investigation,
migration assistée (lire l'ancien, proposer le nouveau, tester, corriger),
veille + synthèse multi-sources pour ton RAG.

Et même dans ces cas : commence par la version la plus bête possible
(boucle ReAct + 2 outils, voir partie G), mesure, puis complexifie.

## 12. Le spectre d'autonomie (niveaux 0 à 4)

L'autonomie n'est pas un interrupteur. Pense en niveaux :

| Niveau | Nom | Qui décide quoi | Exemple |
|---|---|---|---|
| 0 | Manuel | tout l'humain | tu tapes tes commandes |
| 1 | Assisté | l'IA propose, l'humain exécute | chatbot qui donne la commande `df` |
| 2 | Supervisé | l'IA agit en lecture seule, propose les écritures | agent diag qui propose un plan de nettoyage |
| 3 | Semi-autonome | l'IA agit, l'humain valide les actions sensibles | agent qui nettoie après ton « OK » |
| 4 | Autonome | l'IA agit seule dans un périmètre, l'humain audite après | cron nocturne de maintenance (à tes risques) |

Règle d'or : **monte d'un niveau à la fois, et seulement après avoir mesuré
le niveau précédent** (taux de succès, coût, incidents). La plupart des usages
prod sérieux restent aux niveaux 2-3. Le niveau 4 exige : périmètre verrouillé,
budgets durs, journalisation complète, rollback testé, et quelqu'un d'astreinte.

## 13. Le coût caché de l'agentique : latence, tokens, non-déterminisme

Un appel agentique = 5 à 50 appels modèle. Conséquences :

