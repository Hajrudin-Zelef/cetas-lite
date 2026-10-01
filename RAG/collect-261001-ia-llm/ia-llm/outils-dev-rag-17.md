---
id: collect-261001-ia-llm/ia-llm/outils-dev-rag-17
title: "Outils dev + ingénierie RAG (chunk & corpus)"
domain: ia-llm
role: reference
task: reference
actors: []
dates: []
keywords: ["fine-tuning", "valuation"]
source: docs/RAG/collect-261001-ia-llm/outils_dev_rag.md
source_anchor: ""
source_lines: [3175, 3350]
sha256: 9f7ea5aaddb2c0d43168dc254cd32ee466fa13d7f6d79708a30a4a4d849c709e
---

# Outils dev + ingénierie RAG (chunk & corpus)

1. **Re-scraper au lieu de re-jouer** : le brut est conservé (§93).
2. **Headers/footers non strippés** : 600 chunks pollués par le titre du PDF (§95).
3. **Doublons inter-passes** : le même article récupéré en passe 2 ET 4 → dédup doc (§97).
4. **URLs canoniques** : `?utm_source=` crée des faux doublons (§97).
5. **Near-dups non traités** : 5 copies du même fait tuent la diversité du top-k (§98).
6. **Near-dup trop agressif** : « Install Linux » vs « Install Windows » fusionnés (§98).
7. **Mélanger les versions de corpus** dans la même table (§103).
8. **Pas de manifeste** : impossible de savoir ce que contient la base dans 6 mois (§100).
9. **Corpus brut dans git** : repo de 40 Go, clone impossible (§15).
10. **Secrets dans le corpus** : clés API copiées depuis des tutos → filtre PII (§105).
11. **Encodage non vérifié** : `�` dans les chunks (§96).
12. **Pas de quality gate** : pages vides/challenges anti-bot indexées comme du contenu.
13. **Collections mélangées** : question CLI → réponse guide approximative (§101).
14. **Seuil unique** pour des collections hétérogènes (§101).
15. **Évaluation absente** : on ne sait pas si ça marche (§102).
16. **Golden set jamais mis à jour** : il mesure le passé, pas le besoin actuel.
17. **Re-embed partiel** : moitié v2, moitié v3 → résultats incohérents (§103).
18. **Métadonnées perdues** en route (script qui ne les propage pas) (§84).
19. **Ingestion sans transaction** : crash à 80 % = base à moitié remplie → `BEGIN/COMMIT` par batch (§113).
20. **Aucune sauvegarde** de la base avant une réingestion massive (`pg_dump` !).

## 107. Sauvegarde : pg_dump avant chaque réingestion

```bash
# Avant toute opération destructive :
pg_dump "$SUPABASE_DB_URL" -t rag_chunks -Fc -f backups/rag_chunks_$(date +%F).dump
# -Fc = format custom (restauration sélective possible)
# Restaurer :
pg_restore -d "$SUPABASE_DB_URL" backups/rag_chunks_2026-09-27.dump
```

## 108. Checklist corpus

- [ ] Brut conservé hors git, format d'échange standardisé
- [ ] Headers/footers strippés (vérifié sur échantillon)
- [ ] Normalisation UTF-8 (NFC)
- [ ] Dédup exacte (doc + chunk)
- [ ] Near-dup (SimHash puis MinHash), scope par section
- [ ] Manifeste versionné, `corpus_version` sur chaque chunk
- [ ] Collections séparées (`source`), seuils calibrés
- [ ] Golden set (30+ questions) + hit_rate@5 mesuré
- [ ] `pg_dump` avant réingestion

## 109. Pense-bête corpus

```bash
./scripts/ingest.sh v3                        # pipeline complet
python scripts/dedup_stream.py < in.jsonl > out.jsonl
pg_dump "$DB_URL" -t rag_chunks -Fc -f backups/$(date +%F).dump
```

---

## 110. Pourquoi le prompt engineering est ton deuxième levier RAG

Le premier levier, c'est le retrieval (bons chunks remontés). Le deuxième,
c'est **ce que le LLM fait de ces chunks**. Un mauvais prompt transforme
5 bons chunks en réponse inventée ; un bon prompt transforme 5 chunks
moyens en réponse honnête et citée.

```text
Question → [retrieval] → chunks → [PROMPT] → LLM → réponse
                                            ▲
                                     c'est ici que tout
                                     se joue en génération
```

Règles vérifiées en 2026 (consensus des guides de bonnes pratiques) :
- **Commence par le plus simple** : un prompt clair sans exemples. N'ajoute
  du few-shot ou du chain-of-thought que si le simple échoue (décision
  vérifiée : zero-shot → few-shot → fine-tuning).
- **Les prompts sont du code** : versionnés, testés, régressés (voir §20
  et §124).
- **Précis plutôt que verbeux** : des contraintes explicites > des
  paragraphes de prose.

## 111. System prompts : anatomie (vérifié 2026)

Le **system prompt** définit l'identité, la mission et les règles du jeu —
il persiste sur toute la session et a la priorité sur les instructions
utilisateur. Anatomie d'un bon system prompt (structure vérifiée) :

```text
1. RÔLE        → « Tu es [rôle] expert en [domaine]. »
2. TÂCHE       → description claire de ce que l'assistant doit faire.
3. DIRECTIVES  → contraintes de comportement (ton, langue, longueur…).
4. RÈGLES      → NEVER / ALWAYS / fallback si incertain.
5. FORMAT      → spécification EXACTE de la sortie attendue.
```

Bonnes pratiques (vérifié) :
- **Rôle d'abord** : « Tu es un assistant documentation réseau » donne un
  cadre stable à toutes les réponses suivantes.
- **Langue explicite** : si ton corpus est français, impose « réponds en
  français » — sinon le modèle peut basculer en anglais sur des chunks FR.
- **Une instruction = une phrase** : les listes à puces sont mieux suivies
  que les paragraphes.
- **Règles négatives explicites** : « N'invente JAMAIS une commande » est
  plus efficace que « sois précis ».
- **Ordre de priorité** : system > instructions développeur > contexte >
  question utilisateur (garde-fou anti-injection, voir §122).

## 112. Exemple : system prompt RAG FR pour ton assistant réseau

```text
Tu es un assistant expert en documentation réseau et systèmes, qui répond
exclusivement à partir des documents fournis dans le CONTEXTE.

## Règles strictes
- Utilise UNIQUEMENT les informations du CONTEXTE. N'utilise jamais tes
  connaissances générales pour des faits techniques (syntaxes, valeurs,
  versions).
- Cite chaque affirmation avec [N], où N est le numéro du passage source.
- Les chiffres (versions, puissances, tensions, ports) sont recopiés à
  l'identique — jamais paraphrasés ni arrondis.
- Si la réponse n'est pas dans le CONTEXTE, réponds exactement :
  « Je n'ai pas trouvé cette information dans les documents fournis. »
- Si la question est ambiguë, précise l'interprétation choisie AVANT de répondre.

## Format de réponse
- Réponds en français, de façon directe et dense.
- Pour une commande : bloc de code + une phrase d'explication.
- Termine par « Sources : [1] titre, [2] titre ».

## Interdits
- N'invente JAMAIS une syntaxe de commande.
- Ne complète pas un tableau avec des valeurs supposées.
- Ne cite pas un passage qui ne contient pas l'information.
```

> Teste ce system prompt contre tes 3 cas de référence (cas facile, cas
> difficile, cas adversarial) à chaque modification (§102).

## 113. Few-shot prompting : quand et comment (vérifié 2026)

**Principe** : montrer 2 à 5 paires entrée/sortie pour que le modèle
**imite le pattern** — trois bons exemples battent un paragraphe
d'explication.

**Quand l'utiliser :**
- Le zero-shot produit un format instable (JSON parfois cassé, citations
  oubliées).
- Tâche métier spécifique (classifier l'intention, extraire un format).
- Cas limites à démontrer (« je ne sais pas »).

**Comment (règles vérifiées) :**
1. **Qualité > quantité** : 2-5 exemples bien choisis > 20 médiocres
   (rendements décroissants au-delà de 5).
2. **Diversité** : couvrir le cas nominal, un cas limite, et le cas
   « réponse impossible ».
3. **Format identique** : les exemples doivent ressembler EXACTEMENT à la
   sortie attendue (mêmes balises, même JSON, mêmes `[N]`).
4. **Simple → complexe** : ordonner du plus simple au plus difficile.
5. **Pas de pollution** : un exemple hors sujet dégrade plus qu'il n'aide.

```text
Exemple few-shot (classification d'intention pour ton RAG) :

Question : « Quelle est la syntaxe de display interface brief ? »
Intention : commande

Question : « Résume-moi la procédure de mise à jour d'un S310. »
Intention : procedure

Question : « Quelle est la capitale de la France ? »
Intention : hors_corpus

Question : « <question réelle> »
Intention :
```

## 114. Few-shot pour le RAG : imposer le format de réponse

Le cas le plus rentable du few-shot dans TON pipeline : **forcer les
citations et l'abstention** en montrant le comportement attendu.

