# NOTES — corpus `briefing-ia-2026`

## Défauts de la source (conservés verbatim — NE PAS corriger ici)

La découpe est strictement fidèle : les défauts ci-dessous existent dans la source
`docs/RAG/briefing-ia-2026-en.md` et sont **recopiés tels quels**. Ce corpus n'est pas
le lieu pour les corriger ; toute correction doit être faite à la source puis le corpus
régénéré.

1. **Ligne 9957** — liste contradictoire : « disclosure by the victim in August, **never**,
   acknowledgment by the operator in late July » ; le « never » parasite contredit les dates
   établies ailleurs (Hugging Face divulgue le 16 juillet, OpenAI reconnaît le 21 juillet).
2. **Ligne 7035** — tournure probablement fautive : « a chaotically released release publicly
   owned ».
3. **§3 « Master chronology » (4372–4426)** — le texte annonce « eighteen dates » mais le
   tableau compte **20 lignes** de données.
4. **Incohérences ToC ↔ titres** — la table des matières source (fichier
   `00-front-matter/01-table-of-contents.md`) diverge par endroits des titres réels
   (ex. « export controls » vs « export control » ; « permanently » vs « durably » ;
   « In synthesis » vs « In summary »). Les titres des chunks reprennent les **titres réels**.
5. **Numérotation double** — la Section 1 numérote en interne `## 1.`–`## 12.`
   (1–3 = intro/méthode/guide, 4–12 = janvier–septembre), ce qui entre en collision avec la
   numérotation de premier niveau `## 2.`–`## 12.`. Les métadonnées `domain`/`folder` lèvent
   l'ambiguïté.

## Contenu conservé même s'il est redondant

Le dossier narre plusieurs fois les mêmes événements (chronologie §1, profils §2–§6, fond
§7–§11, annexes §12). **Rien n'a été supprimé.** La table « Événements canoniques »
(`INDEX.md`) indique la version de fond (`canonical_for`) ; les autres occurrences restent
présentes (utiles pour le recall).

## Métadonnées

`actors`, `dates` et `keywords` sont **auto-dérivés** par `RAG/_tools/build_rag.py`
(listes contrôlées + extraction regex), à titre d'aide au filtrage — non exhaustifs.
`task` et `canonical_for` proviennent de tables explicites du générateur.

## Régénérer

```bash
python3 RAG/_tools/build_rag.py
python3 RAG/_tools/verify_rag.py
```
