---
id: collect-261001-ia-llm/ia-llm/ia-agents-concepts-3
title: "Concepts : agents IA, agentic, autonomie"
domain: ia-llm
role: reference
task: reference
actors: ["Google"]
dates: []
keywords: ["agent", "agentic", "agents", "benchmarks", "gemini", "incident", "memory"]
source: docs/RAG/collect-261001-ia-llm/ia_agents_concepts.md
source_anchor: ""
source_lines: [301, 450]
sha256: 1f0c8f4277588c96f5808c9a44435a44d94782f941717c5147b4081c3d298964
---

# Concepts : agents IA, agentic, autonomie

- **Latence** : 30 s à 10 min par tâche, pas 3 s. Incompatible avec du temps réel
  interactif sauf à streamer et à paralléliser.
- **Tokens** : l'historique complet est renvoyé à chaque tour. 20 tours × 4 000 tokens
  de contexte = 80 000 tokens d'entrée facturés (même si le contenu est identique).
  D'où l'importance du résumé de contexte (section 27) et de l'élagage.
- **Non-déterminisme** : deux exécutions du même objectif peuvent diverger
  (température, ordre des outils, contenu web changeant). Pour la prod : température 0,
  graines fixées quand possible, et tests en `pass^k` (voir partie I).
- **Coût de l'échec** : un agent qui tourne en rond 40 étapes avant d'abandonner
  coûte 40× le prix d'un succès rapide. Le budget d'étapes n'est pas une option.

Ordres de grandeur (à vérifier selon ton fournisseur et tes modèles, sept 2026) :
une tâche sysadmin simple (bilan disque + proposition) ≈ quelques milliers de tokens ;
un diagnostic multi-sources ≈ dizaines de milliers ; une boucle qui déraille ≈ centaines
de milliers. **Mets toujours un plafond de coût par exécution.**

## 14. Les 3 erreurs de débutant avec l'agentic

1. **Donner le shell sans garde-fou** : `run_shell(commande_quelconque)` + modèle +
   internet = incident en attente. Commence par une liste blanche de commandes
   (voir l'outil `run_shell` sécurisé, section 83).
2. **Pas de budget** : « tant qu'il n'a pas fini ». C'est comme un `while true`
   sans condition de sortie. `max_steps`, `max_tokens`, `timeout` global : les trois,
   dès le premier prototype.
3. **Faire confiance à l'observation** : la sortie d'un outil n'est pas une vérité,
   c'est une donnée. Un fichier peut contenir une instruction, une page web peut
   mentir, une commande peut échouer silencieusement (cf. incident Gemini CLI,
   section 47). Vérifie les effets, pas les dires : après « dossier créé », liste
   le dossier ; après « service redémarré », interroge son statut.

## 15. Carte mentale : où on va dans ce guide

```
COMPRENDRE (A)          → qu'est-ce qu'un agent, quand l'utiliser
ARCHITECTURES (B)       → single vs multi, orchestration, mémoire, outils
AUTONOMIE + RISQUES (C) → ce que « autonome » veut dire, garde-fous, dérives réelles
OUTILS DU MARCHÉ (D,E,F)→ Open Interpreter (vérifié), « cowork » (ambigu), ZCode (identifié)
CONSTRUIRE (G,H)        → ton agent en Python (~150 lignes), puis version LangGraph
ÉVALUER (I)             → benchmarks publics + tests maison
PIÈGES (J)              → 18 pièges documentés avec parades
PRATIQUE (K)            → 4 scénarios complets commentés
CONSOLIDER (L,M,N)      → quiz, glossaire, pour aller plus loin
```

## 16. Comment lire ce guide

- Tu veux juste comprendre : lis A + B + le glossaire (M).
- Tu veux utiliser un outil existant : D (Open Interpreter), F (ZCode).
- Tu veux construire : G puis H, avec J (pièges) en parallèle.
- Tu veux mettre en prod : C + I + J + K, dans cet ordre.
- Les sections marquées « à vérifier » contiennent des valeurs qui bougent vite
  (versions, prix, scores) : revérifie avant d'agir.

---

# PARTIE B — Architectures d'agents

## 17. Single-agent : un cerveau, des outils

L'architecture la plus simple : **un** modèle, **une** boucle, **N** outils.
Tout passe par le même contexte.

Avantages :
- simple à construire, à déboguer, à auditer ;
- pas de problème de coordination ;
- coût prévisible (un seul historique).

Limites :
- le contexte est un goulot : trop d'outils = descriptions d'outils qui noient
  les instructions ; trop d'étapes = explosion de contexte ;
- un seul « rôle » : difficile de séparer l'explorateur du rédacteur du vérificateur ;
- plafonne sur les tâches longues ou multi-domaines.

Verdict : commence toujours par un single-agent. Passe au multi-agents quand tu as
la preuve (mesures, pas intuition) que le single-agent échoue par manque de contexte
ou par mélange des rôles.

## 18. Multi-agents : quand et pourquoi

Le multi-agents se justifie dans 3 cas :

1. **Séparation des rôles** : un agent cherche, un autre vérifie, un troisième rédige.
   Chacun a son prompt système et ses outils : moins de confusion, meilleure qualité.
2. **Parallélisme** : 4 sous-agents explorent 4 pistes en même temps (diagnostic,
   veille). Gain de temps mural, coût en tokens multiplié.
3. **Spécialisation des modèles** : un gros modèle pour planifier, un petit pas cher
   pour exécuter des sous-tâches mécaniques.

Coûts du multi-agents :
- coordination (qui fait quoi ? qui a raison en cas de désaccord ?) ;
- contexte dupliqué (chaque agent recharge une partie du contexte) ;
- débogage exponentiel (reproduire une exécution à 5 agents est un sport).

Règle : **le multi-agents est une optimisation, pas un point de départ.**
Si ton single-agent réussit 80 % du temps, n'ajoute pas 4 agents pour les 20 %
restants avant d'avoir essayé : meilleurs outils, meilleur prompt, mémoire.

## 19. Orchestration : le superviseur (hub-and-spoke)

Pattern le plus courant en 2026 :

```
              ┌──────────────┐
              │  SUPERVISEUR │
              │ (planifie,   │
              │  délègue,    │
              │  synthétise) │
              └──────┬───────┘
        ┌───────────┼────────────┐
        ▼           ▼            ▼
   ┌─────────┐ ┌─────────┐ ┌──────────┐
   │chercheur│ │exécutant│ │vérificateur│
   │ (web)   │ │(shell)  │ │(relecture) │
   └─────────┘ └─────────┘ └──────────┘
```

- Le superviseur détient l'objectif et le plan ; il ne touche pas aux outils
  dangereux (séparation des privilèges : bonne pratique de sécurité).
- Les workers n'ont que les outils de leur rôle (principe du moindre privilège).
- Le superviseur synthétise ; lui seul parle à l'utilisateur.

Variante « graphe » (LangGraph, section 93+) : au lieu d'un superviseur central,
les agents sont des nœuds d'un graphe d'états avec des arêtes conditionnelles.
Plus explicite, plus auditable, meilleur pour la prod (checkpoints, reprise).

## 20. Orchestration : patterns courants (au-delà du superviseur)

| Pattern | Schéma | Quand l'utiliser |
|---|---|---|
| Séquentiel (pipeline) | A → B → C | étapes fixes mais contenu variable (recherche → synthèse → rapport) |
| Parallèle (fan-out/fan-in) | S → {A,B,C} → synthèse | exploration multi-pistes, veille |
| Hiérarchique | superviseur → chefs d'équipe → workers | tâches complexes multi-domaines |
| Débat / critique | A propose, B critique, A corrige | qualité rédactionnelle, revue de plan |
| Marché / enchères (rare) | agents « enchérissent » sur sous-tâches | recherche, cas exotiques |

En pratique sysadmin, 90 % des besoins = séquentiel + parallèle occasionnel.
Le hiérarchique à 3 niveaux est presque toujours du sur-engineering pour un usage interne.

## 21. Mémoire : pourquoi l'agent oublie (et comment l'aider)

Le modèle ne « retient » rien entre deux appels : tout ce qu'il sait vient du contexte
qu'on lui envoie. D'où trois mémoires à construire explicitement :

1. **Court terme (working memory)** : l'historique de la session en cours
   (messages, appels d'outils, observations). Limitée par la fenêtre de contexte.
2. **Long terme (mémoire persistante)** : ce qui survit aux sessions — préférences,
   faits établis, runbooks appris. Stockée hors modèle (fichier, base, vecteurs).
3. **Épisodique / procédurale** : « la dernière fois que ce playbook a planté,
   c'était à cause de X » ; « pour redémarrer ce service, l'ordre c'est A puis B ».

