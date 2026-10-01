---
id: collect-261001-rattrapage/rattrapage/php-guide-6
title: "PHP 8 — Le guide complet du sysadmin qui héberge"
domain: rattrapage
role: reference
task: reference
actors: []
dates: []
keywords: ["decode"]
source: docs/RAG/collect-261001-rattrapage/php_guide.md
source_anchor: ""
source_lines: [1212, 1441]
sha256: e724f82279ca606f239df5bb53b0cf8b0fbb8577fc09514c0f7e2ea60731d081
---

# PHP 8 — Le guide complet du sysadmin qui héberge

```json
{
    "autoload": {
        "psr-4": {
            "App\\": "src/"
        }
    }
}
```

`App\Supervision\Sonde` → `src/Supervision/Sonde.php`. Après ajout d'une classe : `composer dump-autoload` (régénère la table).

🔒 **Sécurité :** l'autoloader ne charge que les classes déclarées — impossible d'inclure un fichier arbitraire via une classe. À comparer avec les `include $_GET['page']` (LFI, section 51).

---

## 30. Composer — le gestionnaire de dépendances

```bash
# Installation (méthode officielle, en root c'est OK pour l'install globale)
php -r "copy('https://getcomposer.org/installer', 'composer-setup.php');"
php composer-setup.php --install-dir=/usr/local/bin --filename=composer
php -r "unlink('composer-setup.php');"
composer --version

# Nouveau projet
mkdir -p /var/www/appli && cd /var/www/appli
composer init --name="entreprise/appli" --type=project --no-interaction

# Ajouter une dépendance (ex. : client HTTP, validateur...)
composer require monolog/monolog
composer require --dev phpunit/phpunit   # dépendance de DEV uniquement

# Installer les dépendances d'un projet existant (déploiement !)
composer install --no-dev --optimize-autoloader --no-interaction

# Mettre à jour
composer update monolog/monolog   # une seule
composer update                   # toutes (⚠️ en dev, jamais en aveugle en prod)

# Voir les paquets obsolètes / vulnérables
composer outdated
composer audit                    # 🔒 audit de sécurité des dépendances
```

**Fichiers clés :**

| Fichier | Rôle | Versionné en git ? |
|---|---|---|
| `composer.json` | Dépendances déclarées | ✅ oui |
| `composer.lock` | Versions exactes installées | ✅ oui (garantit la reproductibilité) |
| `vendor/` | Code des dépendances | ❌ **non** (reconstruit par `composer install`) |

---

## 31. Composer en production — la procédure sûre

```bash
# Sur le serveur de PROD, déploiement type :
cd /var/www/appli

# 1. Jamais "composer update" en prod : on fige via le lock
composer install --no-dev --optimize-autoloader --no-interaction --no-progress

# 2. Vérifier qu'il ne manque rien
composer check-platform-reqs   # extensions PHP requises présentes ?

# 3. Droits : vendor lisible par www-data, jamais inscriptible
sudo chown -R root:www-data vendor
sudo chmod -R 755 vendor
```

`composer.json` — scripts utiles (lancés par `composer <nom>`) :

```json
{
    "scripts": {
        "lint": "find src -name '*.php' -exec php -l {} \\;",
        "tests": "phpunit --colors=always",
        "post-install-cmd": [
            "php bin/cache-clear.php"
        ]
    }
}
```

🔒 **Règles prod :** `--no-dev` (pas de phpunit/xdebug en prod), `vendor/` hors d'atteinte web si possible, `composer audit` dans la CI, et **jamais** de `composer` en root sur les fichiers de l'appli web (propriétaire dédié, ex. `deploy`).

---

## 32. Erreurs vs exceptions — comprendre le modèle PHP 8

Depuis PHP 7, **tout** est throwable : les erreurs fatales sont des objets `Error`.

```
Throwable
├── Error                    → erreurs du moteur (à ne PAS catcher sauf cas précis)
│   ├── TypeError            → mauvais type passé à une fonction typée
│   ├── ParseError           → syntaxe invalide (dans eval/include)
│   ├── DivisionByZeroError  → intdiv(1, 0)
│   └── ...
└── Exception                → erreurs applicatives (à catcher)
    ├── InvalidArgumentException, RuntimeException, LogicException (SPL)
    └── PDOException, JsonException...
```

```php
<?php declare(strict_types=1);

// Hiérarchie d'exceptions métier : un tronc commun = catch ciblé
class AppliException extends Exception {}
class ConfigException extends AppliException {}
class EquipementIntrouvableException extends AppliException {}

function getEquipement(int $id): array
{
    // ...
    throw new EquipementIntrouvableException("Équipement #$id introuvable");
}

try {
    $eq = getEquipement(999);
} catch (EquipementIntrouvableException $e) {
    http_response_code(404);
    echo "Introuvable";
} catch (AppliException $e) {
    // attrape ConfigException et toute autre AppliException
    http_response_code(500);
    error_log($e->getMessage());
}
```

**Philosophie :** les exceptions = cas **exceptionnels** (panne BDD, fichier absent). Pas de contrôle de flux normal avec `try/catch` (ex. : tester l'existence avant plutôt que catcher).

---

## 33. try / catch / finally — le mode d'emploi

```php
<?php declare(strict_types=1);

// --- finally : TOUJOURS exécuté (nettoyage garanti) ---
$fh = fopen('/var/log/import.log', 'a');
try {
    $data = json_decode($json, true, 512, JSON_THROW_ON_ERROR);
    traiter($data);
} catch (JsonException $e) {
    error_log("JSON invalide : " . $e->getMessage());
    $nbErreurs++;
} catch (Throwable $e) {           // filet de sécurité : attrape TOUT
    error_log("Échec import : " . get_class($e) . " : " . $e->getMessage());
} finally {
    fclose($fh);                  // libéré même si exception
}

// --- Multi-catch (PHP 7.1+) ---
try {
    $pdo = connecterBdd($dsn);
} catch (PDOException | ConfigException $e) {
    die("Base inaccessible");     // ⚠️ die() en CLI ; en web : page d'erreur propre (section 52)
}

// --- Relancer après log (enrichir le contexte) ---
try {
    sauvegarder($intervention);
} catch (PDOException $e) {
    error_log("Échec sauvegarde intervention #{$intervention->id}");
    throw $e;  // on laisse remonter après avoir loggué
    // ou : throw new AppliException("Sauvegarde impossible", 0, $e); // chaînage ($e = previous)
}

// --- $e->getPrevious() : remonter la chaîne ---
catch (AppliException $e) {
    $cause = $e->getPrevious(); // l'exception d'origine
}
```

---

## 34. Gestion d'erreurs globale — ne jamais montrer une stack trace

```php
<?php declare(strict_types=1);

// --- bootstrap.php : à inclure en tête de chaque point d'entrée web ---

// 1. Convertir les erreurs PHP (warnings, notices) en exceptions → un seul circuit de gestion
set_error_handler(function (int $errno, string $errstr, string $errfile, int $errline): bool {
    if (!(error_reporting() & $errno)) {
        return false; // erreur masquée par @ (qu'on n'utilise pas de toute façon)
    }
    throw new ErrorException($errstr, 0, $errno, $errfile, $errline);
});

// 2. Gestionnaire d'exceptions non capturées : page 500 PROPRE + log complet
set_exception_handler(function (Throwable $e): void {
    $id = bin2hex(random_bytes(6)); // identifiant pour corréler avec les logs
    error_log("[ERREUR $id] " . get_class($e) . " : " . $e->getMessage()
        . " dans " . $e->getFile() . ":" . $e->getLine()
        . "\n" . $e->getTraceAsString());

    http_response_code(500);
    // En DEV : afficher. En PROD : page générique.
    if (($_ENV['APP_ENV'] ?? 'prod') === 'dev') {
        echo "<h1>Erreur $id</h1><pre>" . htmlspecialchars((string) $e) . "</pre>";
    } else {
        echo "<h1>Erreur interne</h1><p>Référence : $id. L'équipe a été notifiée.</p>";
    }
});

// 3. Erreurs fatales (parse, mémoire) : shutdown function = dernier recours
register_shutdown_function(function (): void {
    $err = error_get_last();
    if ($err !== null && in_array($err['type'], [E_ERROR, E_PARSE, E_CORE_ERROR, E_COMPILE_ERROR], true)) {
        error_log("[FATAL] {$err['message']} dans {$err['file']}:{$err['line']}");
        http_response_code(500);
        echo "<h1>Erreur interne</h1>";
    }
});
```

🔒 **En prod :** `display_errors=Off` (php.ini) + ce trio = **zéro fuite d'information** vers le visiteur, **100 % des erreurs** dans les logs avec un ID de corrélation.

---

## 35. Dates et heures — DateTimeImmutable

```php
<?php declare(strict_types=1);

