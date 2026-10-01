---
id: collect-261001-rattrapage/rattrapage/php-guide-17
title: "PHP 8 — Le guide complet du sysadmin qui héberge"
domain: rattrapage
role: reference
task: reference
actors: []
dates: []
keywords: ["apache", "gpu"]
source: docs/RAG/collect-261001-rattrapage/php_guide.md
source_anchor: ""
source_lines: [3502, 3617]
sha256: 36d3ee5b5043032d9bf758dcc2dbfffd1055e7d1621199b51a1ae1f1c481e049
---

# PHP 8 — Le guide complet du sysadmin qui héberge

```php
<?php declare(strict_types=1);
// gen-api-key.php — CLI : génère une clé API (affiche la clé UNE fois, stocke le hash)
$cle = bin2hex(random_bytes(32));
$hash = hash('sha256', $cle);
$pdo->prepare('INSERT INTO api_cles (nom, cle_hash, active) VALUES (?, ?, 1)')
    ->execute([$argv[1] ?? 'sans-nom', $hash]);
echo "Clé (à copier MAINTENANT, non récupérable) : $cle\n";
```

---

## 79. Glossaire

| Terme | Définition |
|---|---|
| **SAPI** | Interface d'exécution de PHP : CLI, FPM, (anciennement mod_php). Chaque SAPI a son php.ini. |
| **PHP-FPM** | FastCGI Process Manager : gère des pools de workers PHP pour nginx/Apache. |
| **Pool** | Groupe de workers FPM dédié à un site, avec son utilisateur système et sa config. |
| **Worker** | Processus PHP qui traite une requête à la fois. |
| **OPcache** | Cache de bytecode : évite de recompiler les scripts à chaque requête. |
| **PDO** | PHP Data Objects : couche d'accès BDD unifiée (MySQL, PgSQL, SQLite...). |
| **Requête préparée** | Requête SQL où données et structure sont séparées → anti-injection SQL. |
| **XSS** | Cross-Site Scripting : injection de JS via des données non échappées. |
| **CSRF** | Cross-Site Request Forgery : action exécutée à l'insu de l'utilisateur connecté. |
| **LFI/RFI** | Inclusion de fichier local/distant : exécution via `include` d'un chemin contrôlé par l'attaquant. |
| **Whitelist** | Liste de valeurs autorisées (l'inverse : blacklist = liste d'interdits, fragile). |
| **PRG** | Post/Redirect/Get : rediriger après un POST pour éviter le double envoi. |
| **Hachage (mot de passe)** | `password_hash` : fonction lente et salée (bcrypt/Argon2), à sens unique. |
| **Salage** | Aléa unique ajouté avant hachage (géré automatiquement par `password_hash`). |
| **Timing attack** | Attaque qui déduit un secret en mesurant le temps de comparaison → parée par `hash_equals`. |
| **HMAC** | Code d'authentification par clé (ici : signer les cookies). |
| **Composer** | Gestionnaire de dépendances PHP (`composer.json` = déclaration, `composer.lock` = versions figées). |
| **PSR-4** | Convention d'autoloading : namespace ↔ arborescence de dossiers. |
| **Namespace** | Espace de nom : organise les classes (`App\Supervision\Sonde`). |
| **Trait** | Bloc de méthodes réutilisable entre classes sans héritage. |
| **Enum** | Type énuméré : ensemble fermé de valeurs (PHP 8.1). |
| **Closure** | Fonction anonyme stockable dans une variable. |
| **Arrow function** | Closure courte `fn($x) => ...` avec capture automatique (PHP 7.4). |
| **strict_types** | `declare(strict_types=1)` : refuse les conversions de types silencieuses. |
| **DTO** | Data Transfer Object : objet simple qui transporte des données typées. |
| **Slow log** | Log FPM des requêtes dépassant `request_slowlog_timeout`. |
| **HSTS** | En-tête qui force le navigateur à n'utiliser que HTTPS. |
| **CSP** | Content-Security-Policy : restreint les sources de scripts/styles/images. |
| **2FA/TOTP** | Double authentification par code temporaire (application authenticator). |

---

## 80. Quiz — 10 questions + réponses

**Q1.** Quelle est la différence entre `==` et `===` ? Lequel utiliser par défaut ?
> `==` compare avec conversion de type (`"42" == 42` → true), `===` exige type + valeur identiques. Utiliser `===` par défaut.

**Q2.** Pourquoi `display_errors` doit être à `Off` en production ?
> Pour ne pas exposer aux visiteurs les chemins, requêtes SQL, versions et stack traces. Les erreurs vont dans `log_errors` (fichier), jamais à l'écran.

**Q3.** Qu'est-ce qu'une requête préparée et pourquoi est-ce « le seul moyen » contre les injections SQL ?
> La requête et les données sont envoyées séparément au serveur SQL : les données ne peuvent jamais être interprétées comme du code. Aucun échappement manuel n'est aussi fiable.

**Q4.** Un attaquant peut-il passer `ORDER BY` en paramètre lié (`:tri`) ? Comment trier dynamiquement sans faille ?
> Non, les paramètres liés ne remplacent que des **valeurs**, pas la structure SQL. Solution : whitelist de colonnes autorisées + `in_array(..., true)`.

**Q5.** À quoi sert `session_regenerate_id(true)` et quand l'appeler ?
> Anti-fixation de session : à appeler à chaque changement de privilège (login). Le `true` supprime l'ancien fichier de session.

**Q6.** Pourquoi ne faut-il jamais utiliser `$_FILES['x']['type']` pour valider un upload ?
> C'est le **client** qui l'envoie, donc falsifiable. Il faut détecter le vrai type avec `finfo` côté serveur, avec une whitelist.

**Q7.** `md5($password)` : pourquoi est-ce à bannir, et par quoi le remplacer ?
> MD5 est ultra-rapide et cassé : un GPU teste des milliards de hash/s. Remplacer par `password_hash()` (bcrypt/Argon2 : lent, salé) + `password_verify()`.

**Q8.** Après un déploiement avec `opcache.validate_timestamps=0`, que faut-il faire et pourquoi ?
> `systemctl reload php8.4-fpm` : avec `validate_timestamps=0`, PHP ne revérifie jamais les fichiers, donc l'ancien bytecode resterait servi sans reload.

**Q9.** Comment dimensionner `pm.max_children` ?
> `(RAM totale − RAM OS − RAM BDD − marge) / RAM moyenne d'un worker` (mesurée via `ps`). Trop haut = swap/OOM, trop bas = file d'attente.

**Q10.** Citer 3 choses que fait bien le formulaire de contact de la section 72.
> Token CSRF, validation serveur stricte (whitelist), pattern PRG, honeypot anti-bot, stockage brut + `htmlspecialchars` à l'affichage, en-têtes mail sans injection CRLF. (3 parmi celles-ci.)

---

## 81. Pour aller plus loin

**Qualité de code**
- **PHPStan / Psalm** : analyse statique (détecte les bugs sans exécuter) — `composer require --dev phpstan/phpstan`
- **PHP-CS-Fixer** : formattage automatique au standard PSR-12
- **Rector** : migrations automatiques (7.4 → 8.x) et modernisation du code
- **PHPUnit / Pest** : tests automatisés — viser d'abord les fonctions critiques (calculs, validation)

**Frameworks (quand l'appli grandit)**
- **Symfony** : le standard entreprise (composants réutilisables : Mailer, Validator, Console)
- **Laravel** : productivité maximale, écosystème riche
- Règle : un framework ne dispense pas de comprendre ce guide (config, sécurité, FPM).

**Performance avancée**
- **RoadRunner / FrankenPHP** : serveurs d'applications PHP (workers persistants, sans FPM)
- **Swoole** : PHP asynchrone (websockets, haute concurrence)
- Cache applicatif : **Redis** (sessions, résultats, files avec ` predis/predis`)

**Sécurité avancée**
- **2FA TOTP** : `robthree/twofactorauth`
- **Content-Security-Policy** stricte + `nonce` par requête
- Audit : **OWASP ZAP** en CI sur l'environnement de staging
- Veille : s'abonner aux alertes CVE PHP (`composer audit` en cron hebdo)

**Documentation officielle (référence)**
- php.net/manual/fr — la doc de référence (exemples + commentaires)
- php.net/supported-versions — calendrier de fin de vie des versions (ne jamais rester sur une version EOL)
- OWASP Cheat Sheet Series — fiches sécurité par sujet

**Idée de progression pour Zelef :** prendre une petite appli interne (ex. : suivi d'interventions) et lui appliquer ce guide de bout en bout — installation, durcissement, CRUD PDO, supervision. C'est en hébergeant « pour de vrai » que les réflexes deviennent automatiques.

---

*Fin du guide — bon code, et surtout : logs propres, backups testés, PHP à jour.* 🐘
