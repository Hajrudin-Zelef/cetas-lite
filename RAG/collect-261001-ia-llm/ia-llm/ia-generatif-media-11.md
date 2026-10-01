---
id: collect-261001-ia-llm/ia-llm/ia-generatif-media-11
title: "IA générative : image, vidéo, recherche"
domain: ia-llm
role: reference
task: reference
actors: ["EU", "Google", "MiniMax", "OpenAI", "Perplexity"]
dates: ["2026-09-27"]
keywords: ["apache", "diffusion", "incident", "open weights", "perplexity"]
source: docs/RAG/collect-261001-ia-llm/ia_generatif_media.md
source_anchor: ""
source_lines: [986, 1096]
sha256: bdaab1379bf18cb0ce720515a24055d0f331a7a776d856ac00fd04444e7cf002
---

# IA générative : image, vidéo, recherche

**Règle** : usage ponctuel → **facturation à l'usage** (fal.ai, Replicate) ou petit pack Trial. Usage régulier et prévisible → abonnement/API officielle. Jamais de gros pack prépayé « au cas où ».

## 72. Budget annuel type « service systèmes & énergies »

Proposition réaliste pour un service qui produit doc + vidéos + veille :

| Poste | Choix | Coût annuel |
|-------|-------|-------------|
| Images doc (volume) | FLUX local (klein 4B/dev) ou API BFL | 0-50 $ |
| Images premium ponctuelles | fal.ai (Nano Banana Pro / Imagen) à l'usage | ~100 $ |
| Vidéo | Kling Pro web (25,99 $/mois, annuel -34 %) | ~205 $ |
| Vidéo premium ponctuelle | Veo 3.1 Fast via API à l'usage | ~100 $ |
| Recherche interactive | Perplexity Pro (200 $/an) | 200 $ |
| Veille automatisée | API Sonar (`sonar`) | ~10 $ |
| Contrôle qualité / routage | Jev (API) | ~5 $ |
| **Total** | | **~670 $/an** |

Soit **moins de 60 $/mois** pour un outillage complet. Le poste qui dérape toujours : la vidéo (multiplier les essais, l'audio, la 4K). Plafonner par un quota mensuel par personne.

## 73. Dix règles anti-dérapage budgétaire

1. **Toujours convertir les crédits en USD/livrable** avant de comparer.
2. **Ne jamais acheter un gros pack prépayé** sans historique de consommation sur 2-3 mois.
3. **Séparer les budgets** : brouillon (modèles pas chers : schnell, klein, Hailuo, Veo Lite) vs master (pro, 4K).
4. **Noter le fournisseur + le modèle + la date** à côté de chaque comparaison de prix (les grilles bougent indépendamment).
5. **Activer les alertes de dépense** sur chaque console API dès le premier jour.
6. **Une clé API par usage** (dev/prod/veille) pour suivre qui dépense quoi ; rotation en cas de fuite.
7. **Tester à petit format** : 480p/brouillon d'abord, 1080p/4K ensuite (Seedance, Veo et Kling facturent au format).
8. **Couper l'audio natif** quand il n'apporte rien (jusqu'à +50 % à ×5 du coût).
9. **Réutiliser** : une bonne image de référence (Kling O1, FLUX multi-ref) vaut 10 re-générations.
10. **Revue mensuelle** : 15 min pour vérifier le coût/livrable réel vs estimé, et ajuster.

---

# PARTIE I — LICENCES & DROITS : ce que tu as le droit de faire

## 74. Tableau des droits d'usage commercial (vérifié le 27/09/2026)

| Outil | Usage commercial autorisé ? | Conditions / limites |
|-------|----------------------------|----------------------|
| Kling (plans payants) | **Oui** | Plan gratuit : **interdit** (watermark + clause) |
| Kling API | Oui | Selon contrat dev |
| FLUX API (BFL) | Oui | Selon CGU API |
| FLUX.2 [klein] 4B local | **Oui** | Apache 2.0 |
| FLUX.1 [schnell] local | **Oui** | Apache 2.0 |
| FLUX.2 [dev] / klein 9B local | **Non** | Licence commerciale à obtenir auprès de BFL |
| Midjourney | Oui | **>1 M$ de revenus/an → Pro ou Mega obligatoire** ; images publiques par défaut sans Stealth |
| OpenAI (GPT Image / DALL-E, API) | Oui | CGU OpenAI (droits cédés à l'utilisateur) |
| Google (Imagen / Nano Banana, API payante) | Oui | CGU Google Cloud / AI Studio |
| Ideogram (payant) | Oui | Selon plan |
| Adobe Firefly | Oui | Modèle entraîné sur contenu licencié : le plus sûr juridiquement |
| Runway (payant) | Oui | Dès le plan Standard |
| Luma (payant) | Oui | Dès le plan Plus |
| Hailuo / MiniMax | Oui a priori | **Licence territoriale restrictive signalée — à vérifier** |
| Perplexity (contenu des réponses) | Réponses : oui ; **sources citées : leurs propres licences** | Ne pas republier des extraits longs sans vérifier |
| Jev | Oui (sortie = tes décisions) | CGU TypeSafe |

## 75. Le principe juridique de base (pas un avis d'avocat)

En droit français/européen, une image **100 % générée par IA sans apport créatif humain** ne bénéficie en principe pas du droit d'auteur classique (l'originalité suppose une personne physique). Conséquences pratiques :

- Tes visuels IA sont **difficilement protégeables** contre la copie.
- En revanche, **l'assemblage créatif** (direction artistique, sélection, montage, texte ajouté) peut être protégé.
- Pour une doc d'entreprise, ce n'est généralement pas un problème ; pour une identité visuelle de marque, si.
- **Ceci est un principe général, pas un avis juridique** : pour un enjeu réel (marque, contentieux), consulter un juriste.

## 76. Deepfakes, personnes réelles et consentement

- **Ne jamais générer** une personne réelle identifiable (collaborateur, client, personnalité) sans son **consentement écrit**, même « pour une blague de service ». En France, le droit à l'image s'applique ; les montages trompeurs peuvent relever du pénal.
- L'**EU AI Act** impose des **obligations de transparence** pour les contenus générés par IA (mention du caractère artificiel, marquage). Les modalités fines évoluent : **à vérifier** au moment de publier, surtout pour du contenu externe.
- En interne : mentionner « visuel généré par IA » sur les docs qui en contiennent — c'est de l'hygiène, pas de la paranoïa (ça évite qu'on prenne un schéma illustratif pour une photo réelle d'installation).
- Les plateformes (Kling, Luma, Midjourney...) filtrent les contenus sensibles ; les contournements volontaires violent leurs CGU et peuvent faire bannir le compte.

## 77. Données d'entraînement et risque de contentieux

- Des éditeurs attaquent les fournisseurs d'IA (ex. **News Corp, Nikkei/Asahi contre Perplexity**, procédures en cours en 2025-2026). Le risque pèse d'abord sur le **fournisseur**, mais un usage intensif de contenus litigieux dans ta prod n'est pas neutre.
- L'option la plus défendable juridiquement pour une entreprise : **Adobe Firefly** (entraînement licencié) ou des modèles **open weights à licence claire** (Apache 2.0).
- Éviter de générer dans le **style signature d'un artiste vivant nommé** (« à la manière de X ») pour un usage commercial : zone grise juridique et éthique.

## 78. Confidentialité des données envoyées aux API

Règle simple, trois niveaux :

1. **Public / non sensible** → n'importe quelle API.
2. **Interne entreprise** (photos de locaux, schémas, docs) → API avec **résidence UE** quand elle existe (BFL : `api.eu.bfl.ai`), contrats enterprise (SSO, DPA), ou **local**.
3. **Confidentiel / secret** (plans détaillés, données clients, infra critique) → **local uniquement** (FLUX klein 4B/schnell). Aucune API publique.

Ne jamais envoyer : identifiants, clés, données personnelles non nécessaires, documents couverts par un NDA client — sauf validation contractuelle explicite avec le fournisseur.

## 79. Watermarks et métadonnées

- Les offres gratuites **watermarkent** (Kling free, Luma free) : un watermark oublié sur un livrable client = incident d'image.
- Certaines plateformes ajoutent des **métadonnées C2PA** (provenance du contenu). Les conserver : c'est ta preuve de bonne foi en cas de question.
- **Ne jamais retirer un watermark** d'un tiers pour faire passer un visuel gratuit pour un visuel payé : violation de CGU.

## 80. Check-list juridique avant publication d'un visuel/vidéo IA

- [ ] Plan payant / licence qui autorise l'usage commercial ? (vérifié, pas supposé)
- [ ] Aucune personne réelle identifiable sans consentement ?
- [ ] Aucune marque/logo d'un tiers reproduit ?
- [ ] Mention « généré par IA » si diffusion externe (ou interne sensible) ?
- [ ] Source des éléments repris (si montage avec du réel) : droits OK ?
- [ ] Watermark du plan gratuit absent ?
- [ ] Pour FLUX local : variante **Apache 2.0** (klein 4B / schnell), pas dev ?

---

# PARTIE J — PIÈGES TRANSVERSAUX : 20 erreurs à ne pas commettre

## 81. Piège n°1 : les crédits qui expirent en silence

