---
id: collect-261001-ia-llm/ia-llm/ia-agents-concepts-5
title: "Concepts : agents IA, agentic, autonomie"
domain: ia-llm
role: reference
task: reference
actors: []
dates: []
keywords: ["agent", "agents", "arr", "incident"]
source: docs/RAG/collect-261001-ia-llm/ia_agents_concepts.md
source_anchor: ""
source_lines: [598, 760]
sha256: c526e3f7faf0b231fada75aa734bee640e92a2320d2e3494ffb31c39cad7e3a2
---

# Concepts : agents IA, agentic, autonomie

Un agent a besoin d'un critère d'arrêt explicite, sinon il continue (ou s'arrête
trop tôt). Trois mécanismes complémentaires :

1. **Signal du modèle** : réponse finale sans appel d'outil = « j'ai fini ».
   À valider : vérifier que l'objectif est réellement atteint (tests, état).
2. **Vérificateur externe** : un script/check qui dit si l'objectif est atteint
   (ex : le service répond ? le fichier existe avec le bon contenu ?).
   C'est le plus fiable : ne crois pas l'agent sur parole.
3. **Budgets durs** : max étapes / tokens / temps → arrêt forcé avec rapport
   partiel (« voici ce qui a été fait, voici ce qui reste »).

Toujours produire un **rapport de fin** : objectif, actions effectuées, résultats
vérifiés, actions restantes, coût (étapes/tokens/temps). C'est ton audit trail.

## 31. Streaming et UX : voir l'agent travailler

Pour un usage interactif, l'agent doit **streamer** : afficher chaque pensée
(résumée), chaque appel d'outil et chaque observation au fil de l'eau.
Pourquoi c'est important :

- l'humain détecte une dérive en 10 secondes de lecture, contre 10 minutes
  d'analyse de logs après coup ;
- bouton **Stop** visible en permanence (l'équivalent du Ctrl+C) ;
- les actions sensibles s'affichent AVANT exécution avec « approuver / refuser ».

Côté implémentation : la plupart des API LLM supportent le streaming des tokens ;
pour les appels d'outils, affiche au minimum nom + paramètres résumés avant exécution.

## 32. Parallélisme : fan-out contrôlé

Lancer des sous-tâches en parallèle (4 recherches, 3 diagnostics) fait gagner
du temps mural. Règles :

- parallélise les tâches **indépendantes et en lecture seule** ;
- borne le degré de parallélisme (ex : 4 max) : chaque branche consomme du contexte ;
- chaque branche a son **budget local** (sinon une branche qui boucle aspire tout) ;
- la synthèse finale est une étape séparée, avec accès aux résultats des branches.

En Python maison : `concurrent.futures` + un agent par branche + agrégation.
Dans LangGraph : nœuds parallèles + `Send` API (à vérifier selon version).

## 33. Erreurs et retries : l'agent face au réel

Les outils échouent : timeouts, 429, DNS, fichiers absents. Stratégie :

- **classification** : erreur transitoire (retry) vs erreur définitive (abandonner
  cette piste) vs erreur de l'agent (mauvais paramètre → corriger et réessayer une fois) ;
- **retry borné** : 3 essais max, backoff exponentiel, jamais de retry infini ;
- **dégradation** : si l'outil web échoue, bascule sur le cache/doc locale ;
- **journal** : chaque échec est loggé avec l'appel exact (reproductibilité).

Ce que l'agent ne doit JAMAIS faire face à un échec : réessayer la même chose
à l'identique en boucle (piège n°2) ; contourner un refus de sécurité en changeant
de formulation (comportement observé dans l'incident australien de sept. 2026 —
voir section 50) ; inventer un résultat (« le service a l'air OK » sans l'avoir vérifié).

## 34. Idempotence et reprise : l'agent qui peut être interrompu

Un agent prod doit supporter l'interruption (Ctrl+C, crash, timeout) :

- **checkpoints** : état sérialisé après chaque étape (LangGraph le fait nativement) ;
- **actions idempotentes** : « créer le dossier s'il n'existe pas », pas « créer le dossier » ;
- **journal d'actions** : ce qui a été fait est écrit au fur et à mesure (pas seulement
  en mémoire du modèle) ;
- **reprise** : relancer repart du dernier checkpoint, ne rejoue pas les actions
  déjà effectuées (ou les rejoue sans effet grâce à l'idempotence).

Test simple : lance ton agent, tue-le à mi-parcours, relance-le. S'il refait des
bêtises ou recommence à zéro, ton design est à revoir.

## 35. Observabilité : ce qu'il faut logger

Minimum vital pour chaque exécution d'agent :

- objectif initial (verbatim), horodatage début/fin ;
- chaque tour : pensée (résumée), outil appelé + paramètres, résultat (tronqué),
  tokens consommés, durée ;
- décisions sensibles : approbation demandée/obtenue/refusée, par qui, quand ;
- coût total : tokens in/out, appels modèle, durée murale ;
- résultat final + vérification (atteint ? partiel ? échec ? pourquoi).

Sans ces logs, tu ne peux ni déboguer, ni auditer, ni facturer, ni prouver
quoi que ce soit en cas d'incident. Les frameworks 2026 (LangSmith, traçage
OpenTelemetry des SDK agents) fournissent ça ; en Python maison, c'est ~30 lignes
de plus — écris-les dès le début.

## 36. Anti-patterns d'architecture (résumé partie B)

- [ ] multi-agents avant d'avoir un single-agent qui marche ;
- [ ] 50 outils déclarés « au cas où » ;
- [ ] mémoire « on met tout dans le prompt » ;
- [ ] pas de critère d'arrêt vérifiable ;
- [ ] le même agent lit le web ET écrit en prod (sépare les privilèges) ;
- [ ] pas de logs structurés dès le jour 1 ;
- [ ] retry infini « pour être robuste ».

---

# PARTIE C — Agents autonomes : sens, garde-fous, dérives

## 37. « Autonome » : ce que ça veut dire concrètement

« Agent autonome » ne veut pas dire « intelligent et indépendant ». Ça veut dire :
**la boucle plan→action→observation tourne sans intervention humaine entre le
lancement et la fin** (ou l'épuisement du budget).

Concrètement, un agent autonome c'est :

- un objectif formulé une fois au départ ;
- un harness qui boucle sans te demander ton avis à chaque étape ;
- des outils qui modifient réellement des systèmes ;
- un arrêt sur : objectif atteint (vérifié), budget épuisé, ou erreur fatale.

Tout le reste — la « compréhension », la « volonté » — est une projection anthropomorphique.
Un agent autonome est un programme avec une boucle et des API. C'est pour ça que
les garde-fous sont du ressort de l'ingénierie, pas de la philosophie.

## 38. Boucles d'exécution : les 3 formes

1. **Boucle simple** (ReAct) : tant que pas fini et budget OK → un tour.
   Suffit pour 80 % des cas.
2. **Boucle avec vérificateur** : après chaque action « importante », un check
   externe valide l'état avant de continuer. Ralentit, sécurise.
3. **Boucle planifiée** : plan global → exécution → replanification sur échec.
   Pour les tâches longues multi-étapes.

Dans les trois cas, les invariants sont les mêmes : budget dur, journalisation,
arrêt vérifiable, périmètre d'outils fermé.

## 39. Garde-fous indispensables (1/3) : budgets

Trois budgets, toujours, dès le prototype :

```python
BUDGETS = {
    "max_steps": 25,        # tours plan→action→observation
    "max_tokens": 200_000,  # tokens totaux (in + out)
    "timeout_s": 600,       # temps mural max
    "max_tool_calls": 40,   # appels d'outils cumulés
}
```

- `max_steps` : contre les boucles infinies et les dérives lentes.
- `max_tokens` : contre l'explosion de coût (c'est le budget qui protège ton portefeuille).
- `timeout_s` : contre les outils qui pendent (même avec timeout par outil, la somme peut diverger).
- À l'épuisement : arrêt + rapport partiel, jamais de « encore 5 étapes » automatique.

Règle : les budgets sont **durs** (le harness les impose, le modèle ne peut pas
les négocier). Un budget « suggéré dans le prompt » n'est pas un budget.

## 40. Garde-fous indispensables (2/3) : approbations humaines

Matrice de décision (à adapter à ton contexte) :

| Action | Niveau 2 (supervisé) | Niveau 3 (semi-auto) | Niveau 4 (autonome) |
|---|---|---|---|
| Lecture (logs, état) | auto | auto | auto |
| Écriture réversible (fichier tmp, brouillon) | auto | auto | auto |
| Écriture sensible (config prod, DNS, firewall) | approbation | approbation | auto + audit |
| Suppression / redémarrage service | approbation | approbation | interdit ou approbation |
| Exécution réseau sortant | approbation | auto (allowlist) | auto (allowlist) |
| Accès secrets / mots de passe | interdit | interdit | interdit |

