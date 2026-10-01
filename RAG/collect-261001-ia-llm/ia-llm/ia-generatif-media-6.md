---
id: collect-261001-ia-llm/ia-llm/ia-generatif-media-6
title: "IA générative : image, vidéo, recherche"
domain: ia-llm
role: reference
task: reference
actors: ["Google", "Huawei", "Microsoft", "OpenAI", "Perplexity"]
dates: ["2026-09-27"]
keywords: ["chatgpt", "gpu", "perplexity", "research"]
source: docs/RAG/collect-261001-ia-llm/ia_generatif_media.md
source_anchor: ""
source_lines: [525, 599]
sha256: 0d020fa1180dde36dfb9ffb8492bf0ee62c08ffb93c6459b1d197419219c67aa
---

# IA générative : image, vidéo, recherche

1. **Veille normative et documentaire** : « évolutions NF C 15-100 / NF C 18-510 en 2026 », « nouveaux onduleurs 60 kVA comparatif » — avec sources, en 2 minutes au lieu de 30.
2. **Recherche de documentation obscure** : manuels de service, firmware, notes d'application (le prolongement naturel de ton pipeline RAG : Perplexity trouve, ton scraper archive).
3. **Préparation d'intervention** : « procédure de remplacement des condensateurs bus DC sur Easy UPS 3S » + citations → tu croises avec la doc constructeur.
4. **Veille automatisée via l'API Sonar** : un script qui interroge chaque semaine (« nouvelles failles CVE onduleurs connectés », « rappels produits Schneider/Eaton ») et t'envoie un digest. Voir le cas pratique section 163.
5. **Spaces partagés** : un Space « Veille énergie » avec tâches planifiées, partagé à l'équipe — chacun y pioche, personne ne refait la recherche.

## 32. Perplexity vs recherche classique : comparatif honnête

| Critère | Perplexity (Pro) | Google / Bing classique |
|---------|------------------|------------------------|
| Vitesse de synthèse | **Excellente** : réponse rédigée en ~10 s | Tu ouvres 5 onglets et synthétises toi-même |
| Citations | Intégrées, cliquables | Tu dois les trouver toi-même |
| Fraîcheur | Web temps réel | Web temps réel |
| Contrôle des sources | Choix du périmètre (web, académique, social) mais pas de liste blanche perso fine | Opérateurs avancés (`site:`, `-`, guillemets), total |
| Coût | 20 $/mois ou API au token | Gratuit |
| Biais de synthèse | **Le modèle peut mal résumer ou sur-interpréter une source** | Tu lis la source toi-même |
| Recherche obscure / long tail | Bonne, mais parfois « répond à côté » | Souvent meilleure avec les bons opérateurs |
| Confidentialité | Ta question part chez Perplexity (+ Microsoft Azure) | Ta question part chez Google |

**Règle d'usage** : Perplexity pour **dégrossir vite** (80 % du travail en 20 % du temps), puis **toujours ouvrir les 2-3 sources clés** avant d'utiliser une valeur dans un dimensionnement, un devis ou une procédure. Une citation n'est pas une preuve : c'est un pointeur. La preuve, c'est la source lue.

## 33. Perplexity : les pièges spécifiques

- **La synthèse confiante** : Perplexity écrit bien, avec des citations, même quand il a compris de travers. Pour les chiffres (tensions, couples, normes), **vérifier la source primaire**.
- **Sources de qualité variable** : un blog SEO peut être cité au même niveau qu'une doc constructeur. Regarder *qui* est cité, pas seulement *que* c'est cité.
- **Le contentieux éditeurs** : Perplexity est en litige avec des éditeurs de presse (News Corp, Nikkei/Asahi — procédures en cours en 2025-2026). Conséquence possible : retrait ou dégradation de certaines sources. Ne pas en dépendre comme source unique d'actualité.
- **L'API n'est pas gratuite** : le compteur tourne à chaque appel (tokens + frais de requête). Un script de veille mal calibré (boucle trop fréquente, `sonar-pro` au lieu de `sonar`) fait grimper la facture. Commencer par `sonar`, monter en gamme si besoin.
- **Données sensibles** : ne pas coller dans Perplexity (ni son API) des infos confidentielles client, des plans internes, des identifiants. C'est un service cloud américain.
- **Deep Research ≠ expertise** : excellent pour cartographier un sujet, insuffisant pour trancher un point de sécurité électrique. L'expert reste toi.

## 34. Perplexity : Space de veille — exemple concret

Un Space « Veille onduleurs & énergie » utile au quotidien :

1. Créer un Space, ajouter comme instructions : « Sources prioritaires : documentations constructeurs (Schneider Electric, Eaton, Vertiv, Riello, Socomec, Huawei), normes (UTE, IEC), presse spécialisée énergie. Signaler la date de chaque source. »
2. Ajouter une **tâche planifiée hebdomadaire** : « Nouveautés de la semaine : onduleurs triphasés 10-120 kVA, batteries Li-ion pour UPS, réglementation batteries, CVE touchant des équipements d'énergie connectés. Format : puces avec lien source et une ligne de résumé. »
3. Partager le Space à l'équipe (Enterprise) ou exporter le digest.

Coût : inclus dans Pro/Max. Via API, le même digest hebdomadaire coûte quelques dizaines de centimes en `sonar` (voir cas pratique section 163).

---

# PARTIE D — PANORAMA IMAGE : les autres acteurs

## 35. Pourquoi un panorama (et pas un seul outil)

Aucun générateur d'image n'est le meilleur partout. Le choix dépend de ce que tu optimises : **esthétique** (Midjourney), **fidélité aux instructions** (GPT Image), **photoréalisme** (Imagen/Nano Banana Pro), **coût + local** (FLUX), **texte dans l'image** (Ideogram, FLUX.2, Nano Banana Pro), **sécurité juridique** (Adobe Firefly). Le tableau comparatif est en section 43.

## 36. Midjourney : la référence esthétique

- **Modèle actuel au 27/09/2026 : V8.2** (défaut depuis le 24 juillet 2026 ; V8.1 par défaut depuis le 11 juin 2026). Images natives **2K** sans upscaling, ~5× plus rapide que V7 selon l'éditeur.
- Accès : **midjourney.com** (interface web ; Discord devenu optionnel). **Pas d'offre gratuite** (l'essai gratuit a été supprimé en mars 2023 et n'est jamais revenu).
- Fonctionnalités 2026 : **Edit Model** (édition par instruction, jusqu'à 4 références, inpainting/outpainting — en test août 2026), **personnalisation** (profil esthétique appris de tes notations), moodboards et style references, **génération vidéo** (courts clips SD/HD selon le plan), mode Draft pour itérer vite, paramètres `--p`, `--tile 2.0`, `--weird`, `--stylize`, `--hd`.
- Le point fort historique, inchangé : le **« look Midjourney »** — rendu pictural/cinématographique léché **sans effort de prompt**. Pour des visuels « beaux » rapidement, imbattable. Revers : ce style signature est reconnaissable, parfois indésirable pour du documentaire technique (préférer FLUX dans ce cas).

## 37. Midjourney : prix (vérifiés le 27/09/2026)

Facturation au **temps GPU**, pas à l'image. ~20 % de remise en annuel.

| Plan | Mensuel | Annuel (équiv./mois) | GPU rapides | Relax (illimité lent) | Stealth (privé) |
|------|---------|----------------------|-------------|----------------------|-----------------|
| Basic | 10 $ | 8 $ | 3,3 h (~200-250 images) | Non | Non |
| Standard | 30 $ | 24 $ | 15 h (~900 images) | Oui | Non |
| Pro | 60 $ | 48 $ | 30 h (~1 800 images) | Oui (+ vidéo SD) | **Oui** |
| Mega | 120 $ | 96 $ | 60 h (~3 600 images) | Oui (+ vidéo SD) | **Oui** |

- Une image standard V8 = ~0,8 min GPU ; une image HD 2048px = ~1,3 min GPU. Heure GPU supplémentaire : **4 $**.
- **Droits commerciaux** : inclus sur tous les plans ; **les entreprises à plus d'1 M$ de revenus annuels doivent être en Pro ou Mega**.
- **Pas d'API publique** (une offre Enterprise sur mesure existerait selon des sources tierces — **à vérifier** ; en pratique, pas d'intégration programmatique standard).
- Images **publiques par défaut** : sans Stealth (Pro/Mega), tes générations sont visibles par la communauté. Pour du confidentiel : à proscrire, ou passer en Pro.

## 38. OpenAI : GPT Image (DALL-E nouvelle génération)

OpenAI a fait basculer sa génération d'image vers la famille **GPT Image** (le modèle derrière le générateur de ChatGPT), DALL-E 3 restant disponible via API :

