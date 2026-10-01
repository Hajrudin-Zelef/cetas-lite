---
id: collect-261001-ia-llm/ia-llm/ia-agents-concepts-4
title: "Concepts : agents IA, agentic, autonomie"
domain: ia-llm
role: reference
task: reference
actors: ["Anthropic", "Microsoft"]
dates: []
keywords: ["agent", "agents", "arr", "attention", "incident", "mcp", "memory", "model context protocol", "sandbox"]
source: docs/RAG/collect-261001-ia-llm/ia_agents_concepts.md
source_anchor: ""
source_lines: [451, 597]
sha256: 0f3b984ebbb295ebce5397da79a708168305df47f599c6bbb85c18b4d48c11bb
---

# Concepts : agents IA, agentic, autonomie

Sans mémoire explicite, chaque session repart de zéro : l'agent refait les mêmes
erreurs, redemande les mêmes infos. Avec trop de mémoire injectée, le contexte
explose et le modèle se noie. L'art est dans la **sélection** : n'injecter que
le pertinent (voir sections 26-28).

## 22. Mémoire court terme : l'historique et ses limites

L'historique brut = tous les messages + appels d'outils + observations.
Problèmes :

- croissance linéaire avec les étapes ;
- les observations d'outils (sorties shell, pages web) sont les plus grosses ;
- au-delà d'un seuil, le modèle « oublie » le début (dilution de l'attention)
  ou l'appel échoue (fenêtre dépassée).

Techniques de gestion :

- **Troncature** : garder les N derniers tours. Simple, brutal, perd l'objectif initial
  → toujours re-injecter l'objectif et les faits clés en tête (prompt système dynamique).
- **Élagage des observations** : ne garder que les N derniers caractères d'une sortie
  d'outil, ou un résumé. Les sorties shell > 4 000 caractères sont presque toujours
  du bruit pour la décision suivante.
- **Résumé glissant** (section 27) : la méthode propre.

## 23. Mémoire long terme : fichiers, bases, vecteurs

Options, par complexité croissante :

1. **Fichier Markdown** (`MEMORY.md`, `AGENTS.md`) : faits stables, runbooks,
   préférences. Lisible par l'humain, versionnable (git), injecté en tête de contexte.
   C'est ce qu'utilisent beaucoup d'assistants perso en 2026. Suffit pour 80 % des cas.
2. **Base clé-valeur / SQLite** : faits structurés (inventaire, états, historiques
   d'incidents). Requêtable par l'agent via un outil `memory_search`.
3. **Base vectorielle** : souvenirs en langage naturel retrouvés par similarité
   (ton RAG peut servir de mémoire long terme à ton agent — synergie directe
   avec ton projet).
4. **Graphe de connaissances** : relations entre entités (serveur X dépend du
   switch Y, l'onduleur Z protège la baie B). Puissant, coûteux à maintenir.

Règle : commence par le fichier Markdown. Passe à SQLite quand tu as > 200 faits.
Passe aux vecteurs quand la recherche par mots-clés ne suffit plus.

## 24. Mémoire : le pattern « résumé + faits »

Le pattern qui marche en pratique pour les sessions longues :

```
À chaque K étapes (ex : K=10) ou quand le contexte dépasse 70 % :
1. le modèle résume la trajectoire en ~20 lignes :
   - objectif, faits établis, actions tentées, échecs et leçons, état courant
2. le résumé REMPLACE l'historique détaillé (gardé en archive hors contexte)
3. les faits durables sont extraits vers la mémoire long terme
4. la boucle repart avec : objectif + résumé + faits pertinents
```

Avantages : contexte borné, sessions « infinies » possibles, apprentissage cumulatif.
Pièges : le résumé peut perdre un détail critique (→ garde les artefacts bruts :
commandes exactes, chemins, codes d'erreur, dans un bloc « faits » non résumé) ;
le résumé lui-même peut être manipulé par injection (→ ne résume jamais une
observation non vérifiée comme un fait).

## 25. Outils : anatomie d'un bon outil

Un outil = nom + description + schéma de paramètres + implémentation + politique.
Checklist d'un bon outil :

- [ ] **nom verbeux et clair** : `restart_systemd_service`, pas `rs` ;
- [ ] **description précise** : ce qu'il fait, ce qu'il ne fait PAS, effets de bord ;
- [ ] **paramètres typés et bornés** : enum quand possible, longueurs max, formats ;
- [ ] **idempotence quand possible** : relancer deux fois = même effet qu'une fois ;
- [ ] **mode dry-run** : `restart_systemd_service(service="nginx", dry_run=true)`
      → « je redémarrerais nginx » sans le faire. Indispensable pour plan-then-execute ;
- [ ] **sortie structurée** : JSON avec `ok`, `code`, `stdout`, `stderr`, pas du texte libre ;
- [ ] **erreurs explicites** : « service inconnu : 'ngnix' (vouliez-vous 'nginx' ?) » ;
- [ ] **journalisation** : chaque appel loggé (qui, quand, quoi, résultat).

Un outil mal conçu (paramètres texte libre, effets de bord cachés, erreurs muettes)
est un incident en puissance, quel que soit le modèle.

## 26. Outils : le principe du moindre privilège

Donne à chaque agent exactement les outils dont il a besoin, et pas un de plus :

- agent de **diagnostic** : lecture seule (logs, métriques, `systemctl status`) ;
- agent de **documentation** : lecture fichiers + écriture dans UN dossier de sortie ;
- agent de **maintenance** : écriture, mais avec approbation humaine par action sensible.

Techniques :

- **scopes par rôle** : le superviseur n'a pas le shell, le worker shell n'a pas le réseau ;
- **listes blanches** : commandes autorisées explicites (voir section 83) ;
- **bac à sable (sandbox)** : l'outil shell s'exécute dans un conteneur/VM dédié,
  pas sur ta machine (voir section 52) ;
- **quotas par outil** : max N appels à `web_search` par session, max M Mo écrits, etc.

## 27. MCP : le protocole, en bref (renvoi)

**MCP (Model Context Protocol)**, initié par Anthropic fin 2024, est devenu en 2025-2026
le standard ouvert pour brancher des outils/des sources de données sur les agents :
un **serveur MCP** expose des outils/ressources via un protocole unique, au lieu d'un
connecteur ad hoc par service. Un agent « MCP-compatible » peut utiliser des milliers
de serveurs MCP existants (fichiers, bases, navigateurs, SaaS…).

Ce qu'il faut en retenir ici (le détail est l'objet d'un autre guide) :

- MCP résout l'interopérabilité des outils : tu écris un serveur une fois,
  tous les agents compatibles l'utilisent.
- MCP **déplace** le risque, il ne le supprime pas : un serveur MCP malveillant ou
  compromis peut injecter des instructions via la **description de ses outils**
  (alerte Microsoft, juin 2026) ; 36,7 % de 7 000+ serveurs MCP analysés présentaient
  des failles SSRF (février 2026, à vérifier dans le détail).
- Règle : traite un serveur MCP comme du code tiers : épingle la version, lis ce qu'il
  fait, limite son périmètre réseau/fichiers, journalise ses appels.

## 28. Sélection du contexte : ne pas tout donner au modèle

Plus tu donnes d'outils, de docs et d'historique au modèle, plus il se disperse.
Techniques de sélection :

- **outils paresseux (lazy)** : ne déclarer que les outils probables ; en charger
  d'autres à la demande (« j'ai besoin d'un outil réseau → charge le pack réseau »).
- **RAG sur les outils** : décrire les outils dans une base et ne déclarer que
  les K plus pertinents pour l'objectif (utile au-delà de ~30 outils).
- **mémoire à la demande** : l'agent appelle `memory_search("onduleur baie B")`
  au lieu de recevoir toute la mémoire en tête.
- **compression des observations** : résumer les grosses sorties avant réinjection.

Objectif chiffré (ordre de grandeur) : garder le contexte de travail sous ~50 %
de la fenêtre du modèle. Au-delà, la qualité de décision chute et le coût explose.

## 29. Le prompt système d'un agent : ce qu'il doit contenir

Un bon prompt système d'agent (5 blocs) :

1. **Rôle et objectif** : « tu es un agent de diagnostic réseau, ton objectif est… ».
2. **Règles d'action** : quand agir seul / quand demander / ce qui est interdit
   (« ne redémarre jamais un service sans validation », « ne supprime rien »).
3. **Format de travail** : ReAct explicite, langue, format de réponse finale.
4. **Budgets** : « tu as 20 étapes max ; si tu bloques 3 étapes, résume et demande ».
5. **Contexte utile** : faits stables (inventaire, conventions), pas l'historique.

Anti-patterns : prompt de 3 000 lignes que personne ne relit ; règles contradictoires
(« sois autonome » + « demande pour tout ») ; instructions de sécurité vagues
(« fais attention » ne veut rien dire — écris des règles testables).

## 30. Arrêt et fin de tâche : définir « c'est fini »

