---
id: collect-261001-ia-llm/ia-llm/ia-agents-concepts-16
title: "Concepts : agents IA, agentic, autonomie"
domain: ia-llm
role: reference
task: reference
actors: ["OpenAI"]
dates: []
keywords: ["agent", "agents", "incident", "memory"]
source: docs/RAG/collect-261001-ia-llm/ia_agents_concepts.md
source_anchor: ""
source_lines: [2376, 2528]
sha256: 4f0d20603f299ba52bbda0cb6cb76113c312b03f9c8d95c3931b081b17a30a46
---

# Concepts : agents IA, agentic, autonomie

**Quoi** : mots de passe, clés API ou tokens injectés dans le prompt / l'historique
→ ils partent chez le provider du modèle, restent dans les logs, et peuvent
ressortir dans une réponse.
**Exemple** : l'agent lit un fichier de conf contenant `DB_PASSWORD=...` pour
« comprendre la config » : le secret est maintenant dans l'historique, donc
facturé, loggé, et potentiellement régurgité.
**Parade** :
- les secrets vivent dans l'**environnement du harness**, jamais dans le prompt ;
- les outils qui lisent des configs **expurgent** les secrets avant de renvoyer
  (`password=***`) ;
- logs : masque les motifs `password|secret|token|key` (regex simple, à adapter) ;
- avec un modèle cloud : ne jamais traiter de données classifiées sans validation
  de la politique de données.

## 113. Piège n°9 : l'agent écrit en prod depuis une source non fiable

**Quoi** : l'agent lit le web (ou un ticket, ou un mail) puis applique ce qu'il
y trouve sur tes systèmes : commande copiée d'un forum, config « recommandée »
par une page.
**Exemple** : « pour optimiser PostgreSQL, mettez shared_buffers à 80 % de la RAM »
(lu sur un blog) → l'agent l'applique → OOM killer.
**Parade** :
- séparation stricte : l'agent qui **lit le web** n'est pas celui qui **écrit en prod**
  (moindre privilège, section 26) ;
- toute config appliquée vient d'une source **approuvée** (ton RAG, tes runbooks)
  ou est validée par un humain ;
- dry-run + diff affiché avant application (« voici ce que je changerais »).

## 114. Piège n°10 : non-reproductibilité (ça marche une fois sur trois)

**Quoi** : l'agent réussit parfois, échoue parfois, sans changement apparent.
Cause : température > 0, ordre des résultats web, timing, état initial différent.
**Exemple** : le diagnostic réseau réussit le matin (cache DNS chaud) et échoue
le soir.
**Parade** :
- `temperature=0`, graines fixées quand l'API le permet ;
- mesure en `pass^k` (section 101) : exige 3/3 avant de parler de fiabilité ;
- fige l'état initial des tests (snapshots, fixtures) ;
- log tout (section 55) pour comparer deux exécutions divergentes.

## 115. Piège n°11 : l'agent « réussit » sans avoir rien vérifié

**Quoi** : réponse finale confiante (« c'est réparé ! ») alors que l'objectif
n'est pas atteint : l'agent n'a vérifié aucun effet.
**Exemple** : « j'ai redémarré nginx » (l'outil a retourné ok) mais le service
est en failed — l'agent n'a jamais fait `systemctl is-active nginx`.
**Parade** :
- vérificateur externe obligatoire sur les objectifs d'action (section 30) ;
- règle système : « ne déclare un succès que sur la base d'une observation
  qui le prouve » ;
- dans l'eval set : des correcteurs qui vérifient l'état réel, pas le discours
  (section 103).

## 116. Piège n°12 : délégation aveugle en multi-agents

**Quoi** : le superviseur fait confiance aux workers sans vérifier : un worker
hallucine un résultat, le superviseur le recopie dans le rapport final.
**Exemple** : worker « inventaire » renvoie 42 serveurs (il en a inventé 6) ;
le rapport part au chef avec de fausses données.
**Parade** :
- chaque résultat de worker est accompagné de ses **preuves** (commandes, extraits) ;
- un nœud vérificateur contrôle les affirmations critiques (re-exécute un
  échantillon) ;
- les workers n'ont pas le droit d'inventer : « si tu ne trouves pas, dis-le ».

## 117. Piège n°13 : le contexte empoisonné entre sessions

**Quoi** : la mémoire long terme accumule une info fausse ou malveillante
(ex : un runbook « appris » contenant une mauvaise commande), puis la ressert
à chaque session.
**Exemple** : lors d'une session, l'agent note « pour ce switch, la commande est
X » (X était une erreur de frappe) → toutes les sessions suivantes réutilisent X.
**Parade** :
- la mémoire long terme est **versionnée et relue** (git sur MEMORY.md) ;
- distinction faits vérifiés / hypothèses dans la mémoire ;
- purge périodique : ce qui n'a pas servi en 3 mois sort (ou est archivé) ;
- jamais d'écriture automatique en mémoire long terme depuis une observation
  non vérifiée (section 24).

## 118. Piège n°14 : dépendance au fournisseur (API down, prix, modèle retiré)

**Quoi** : ton agent est câblé sur un modèle/une API qui change : hausse de prix,
modèle déprécié, panne, rate limits.
**Exemple** : le provider retire le modèle exact que tu utilises → ton agent
tombe en panne un lundi à 8h.
**Parade** :
- couche d'abstraction : `AGENT_MODEL` en variable d'env, pas en dur ;
- utilise une interface standard (Chat Completions / OpenAI-compatible) pour
  pouvoir changer de provider ;
- teste périodiquement un **modèle de repli** (moins cher / local) sur ton eval set ;
- surveille les annonces de dépréciation de ton provider.

## 119. Piège n°15 : tests insuffisants (l'agent n'a jamais été attaqué)

**Quoi** : on teste le chemin nominal, jamais l'adversarial. Le premier test
d'injection arrive en prod, par un vrai attaquant ou un vrai fichier piégé.
**Exemple** : l'eval set ne contient que des tâches « gentilles » → l'agent
obéit au premier ticket contenant une instruction cachée.
**Parade** :
- 20 % de ton eval set = tâches **piégées** (section 103, étape 5) ;
- rejoue les attaques connues : injection dans fichier, dans page web, dans
  description d'outil, contournement de refus ;
- refais passer l'eval de sécurité à chaque changement de modèle ou de prompt
  système (un nouveau modèle peut être plus crédule… ou plus rusé).

## 120. Piège n°16 : l'observabilité arrive « plus tard »

**Quoi** : « on ajoutera les logs quand ce sera en prod ». Résultat : le premier
incident est inexplicable, le premier dépassement de budget est une surprise.
**Exemple** : facture ×10 un mois, impossible de dire quel agent, quelle tâche,
quelle boucle.
**Parade** :
- logs JSONL dès le premier prototype (section 84) — c'est 10 lignes ;
- tableau de bord minimal : exécutions/jour, tokens/jour, taux de succès,
  garde-fous déclenchés ;
- alerte sur : budget dépassé, approbation refusée, boucle détectée, coût/jour
  > seuil.

## 121. Piège n°17 : périmètre flou (« débrouille-toi »)

**Quoi** : objectif vague + outils larges = l'agent définit lui-même son périmètre,
souvent trop large. « Remets d'aplomb le serveur » peut inclure la suppression
de « fichiers inutiles » selon lui.
**Exemple** : objectif « fais de la place sur /var » → l'agent supprime des logs
dont tu avais besoin pour un audit en cours.
**Parade** :
- objectif **spécifique et borné** : quoi, où, jusqu'où, et surtout **ce qui est
  interdit** (« ne supprime rien, propose seulement ») ;
- outils déjà restreints (liste blanche) : même avec un objectif flou, l'agent
  ne peut pas sortir du bac ;
- en cas d'ambiguïté : l'agent **demande** au lieu de deviner (règle système).

## 122. Piège n°18 : croire que le framework sécurise

**Quoi** : « on utilise [framework à la mode], donc c'est safe ». Le framework
structure la boucle ; il ne connaît ni tes données, ni tes systèmes, ni tes risques.
**Exemple** : un graphe LangGraph impeccable dont l'outil shell est
`subprocess.run(cmd, shell=True)` branché sur l'input du modèle.
**Parade** :
- relis la checklist de la section 98 à chaque migration de framework ;
- audite les **outils**, pas le graphe : c'est là que vit le risque ;
- garde ton eval set de sécurité : il doit passer **quel que soit** le framework.

---

# PARTIE K — 4 cas pratiques complets commentés

> Chaque cas : objectif, architecture, outils, garde-fous, budgets, métriques,
> et les commentaires « ce qui peut mal tourner ». Tous sont au niveau 2-3
> d'autonomie (jamais 4 sans la doctrine de la partie C).

## 123. Cas n°1 : agent de tri de tickets

