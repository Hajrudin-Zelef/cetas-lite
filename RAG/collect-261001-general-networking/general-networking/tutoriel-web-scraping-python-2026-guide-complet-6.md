---
id: collect-261001-general-networking/general-networking/tutoriel-web-scraping-python-2026-guide-complet-6
title: "Créer le dossier du projet"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: ["agent", "arr"]
source: docs/RAG/collect-261001-general-networking/tutoriel-web-scraping-python-2026-guide-complet.md
source_anchor: ""
source_lines: [626, 665]
sha256: 4b3365d2fdbedd70d1adb5b18eecfef1a73594a257b94651f43616cb4b672f20
---

# Créer le dossier du projet

**Honeypot Links :** Certains sites intègrent des liens invisibles pour les humains (cachés via CSS : `display:none`, `visibility:hidden`, ou positionnés hors écran) mais visibles pour les scrapers qui ne vérifient pas les propriétés CSS. Cliquer sur ces liens ou les suivre déclenche immédiatement un bannissement permanent de l'IP.

**Rate Limiting et IP Banning :** Trop de requêtes depuis une même adresse IP déclenchent des blocages progressifs : d'abord ralentissement, puis CAPTCHA, puis blocage temporaire, enfin bannissement permanent. La solution : randomiser les délais entre requêtes (1 à 5 secondes), utiliser des pools de proxies rotatifs résidentiels, et respecter scrupuleusement les headers `Retry-After` des réponses 429.

**Behavioral Analysis :** Les systèmes anti-bot modernes analysent les patterns comportementaux : vitesse de navigation trop uniforme, absence de mouvements de souris, accès direct à des URLs profondes sans navigation préalable. Playwright permet de simuler des comportements plus humains via des délais aléatoires, des mouvements de souris, et des patterns de navigation réalistes.

## 8 Pièges Courants en Web Scraping Python à Éviter Absolument

Après avoir analysé des centaines de projets de **web scraping Python** et d'**extraction données web Python**, voici les erreurs les plus fréquentes qui font échouer les projets de scraping, classées de la plus courante à la plus grave.

1. **Ne pas respecter le fichier robots.txt** – Le non-respect du robots.txt peut entraîner des conséquences légales et une interdiction permanente de l'IP. Utilisez toujours`urllib.robotparser.RobotFileParser` pour vérifier les permissions avant de commencer tout projet de scraping. Vérifiez aussi le champ`Crawl-delay` qui indique la fréquence maximale autorisée.
2. **Effectuer des requêtes sans délai** – Scraper à pleine vitesse sans délai est non seulement détectable mais potentiellement assimilable à une attaque par déni de service (DDoS). Implémentez toujours des délais aléatoires entre 0.5 et 3 secondes minimum. Pour les sites sensibles ou les serveurs avec peu de ressources, augmentez à 5-10 secondes.
3. **Utiliser un User-Agent générique ou "python-requests"** – Le User-Agent par défaut de`requests` est "python-requests/2.x.x" – un signal d'alarme immédiat pour tout système anti-bot. Toujours utiliser un User-Agent de navigateur réel, récent, et le mettre à jour régulièrement (les navigateurs sortent de nouvelles versions tous les 6 semaines environ).
4. **Ne pas gérer les erreurs et exceptions** – Un scraper sans gestion d'erreurs robuste s'arrêtera au premier problème réseau, code 503 temporaire, ou structure HTML inattendue. Implémentez systématiquement un retry avec backoff exponentiel et loggez toutes les erreurs pour analyse ultérieure.
5. **Charger toutes les données en mémoire simultanément** – Pour les datasets de plusieurs milliers d'entrées, accumuler toutes les données en RAM avant de sauvegarder peut provoquer des`MemoryError` . Utilisez des générateurs Python et sauvegardez par batches de 500-1000 éléments.
6. **Ne pas nettoyer les données extraites** – Les données brutes du HTML contiennent souvent des espaces insécables (`\xa0` ), des sauts de ligne parasites, des symboles de devises ambigus, et des encodages incorrects. Appliquez systématiquement`.strip()` , des normalisations Unicode (`unicodedata.normalize('NFKC', texte)` ), et des validations de types.
7. **Ignorer les changements de structure HTML** – Les sites web évoluent constamment. Un scraper basé sur des sélecteurs trop spécifiques (`.product-card > div:nth-child(3) > span.price-value` ) échouera silencieusement au prochain déploiement du site. Préférez des sélecteurs sémantiques plus robustes et ajoutez des assertions sur le nombre d'éléments trouvés.
8. **Scraper des données personnelles sans base légale RGPD** – Collecter des données personnelles (emails, noms, photos de profil, numéros de téléphone) sans base légale valide constitue une violation grave du RGPD. En France, la CNIL peut prononcer des sanctions atteignant 4% du chiffre d'affaires mondial annuel ou 20 millions d'euros, selon le montant le plus élevé.

## Dépannage Complet : 10 Erreurs Fréquentes et Leurs Solutions

Voici les erreurs les plus communes rencontrées lors du développement de scrapers Python, avec leurs causes précises et solutions détaillées testées en production.

| Erreur / Symptôme | Cause Probable | Solution Détaillée | 
|---|---|---|
| `HTTPStatusError: 403 Forbidden` | User-Agent détecté, IP bloquée, cookies manquants | Mettre à jour les headers complets, utiliser curl_cffi, ajouter les cookies de session, vérifier robots.txt | 
| `HTTPStatusError: 429 Too Many Requests` | Limite de taux dépassée | Lire le header Retry-After, augmenter les délais à 5-10s, réduire max_concurrent à 3-5 | 
| `TimeoutError / ConnectTimeout` | Serveur lent, réseau instable, timeout trop court | Augmenter timeout à 60s, implémenter retry avec backoff 2^n secondes | 
| `AttributeError: 'NoneType'.text` | Sélecteur CSS incorrect ou structure HTML modifiée | Rouvrir DevTools et re-vérifier le sélecteur, utiliser `if element:` systématiquement | 
| Page HTML vide ou contenu absent | Contenu chargé dynamiquement par JavaScript | Inspecter les requêtes XHR dans DevTools, utiliser Playwright ou intercepter les appels API | 
| `UnicodeDecodeError` | Encodage mal détecté ou déclaré par le serveur | Forcer `response.encoding = 'utf-8'` ou utiliser`chardet.detect(response.content)` | 
| Données incorrectes ou mélangées | A/B testing côté serveur, contenu géolocalisé | Analyser la réponse brute avec `response.text` , tester depuis différentes IPs/locales | 
| Playwright `TimeoutError waiting for selector` | Sélecteur absent ou contenu non chargé dans le délai | Augmenter timeout à 15000ms, utiliser `wait_for_selector(..., state="visible")` | 
| `SSLCertVerificationError` | Certificat SSL expiré ou self-signed sur le serveur cible | Utiliser `verify=False` en dernier recours (risque sécurité), documenter la dérogation | 
| `MemoryError` sur grands datasets | Accumulation de toutes les données en RAM avant sauvegarde | Utiliser des générateurs Python, sauvegarder par batches, augmenter le swap ou utiliser un streaming parser | 

### Stratégie de Retry avec Backoff Exponentiel

Voici un pattern de retry robuste que tout scraper production devrait implémenter. Il gère automatiquement les erreurs temporaires de réseau et les codes 429 avec un délai croissant entre les tentatives, en ajoutant du "jitter" (variation aléatoire) pour éviter les tempêtes de reconnexions simultanées.

