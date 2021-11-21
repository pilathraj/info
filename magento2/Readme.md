## Magento 2
- **Installation**
```cmd
php bin/magento setup:install --base-url=http://localhost/magento24/ --db-host=localhost --db-name=yourdbname --db-user=yourdbuser --db-password=yourdbpassword --admin-firstname=admin --admin-lastname=admin --admin-email=admin@admin.com --admin-user=admin --admin-password=admin123 --language=en_US --currency=USD --timezone=America/Chicago --use-rewrites=1 --backend-frontname=admin --search-engine=elasticsearch7 --elasticsearch-host=localhost --elasticsearch-port=9200
```
- **Disabled maintainence**
```cmd
php bin/magento maintenance:disable
```
- **Change the base url**
```cmd
php bin/magento setup:store-config:set --base-url="http://localhost:8080/"
# for secure-base-url
php bin/magento setup:store-config:set --base-url-secure="https://localhost:8080/"
# call clear cache
# rm -rf var/cache var/generation var/page_cache  
```
- **clear cache**
```cmd
php bin/magento cache:flush
```
- **Update static content**
```cmd
php bin/magento setup:static-content:deploy
```
- **Blank page issue**
  - https://meetanshi.com/blog/solved-magento-2-2-7-and-2-3-admin-page-blank-issue/
  - Path lib/internal/Magento path.

- **CSS not loaded**
  - php bin/magento setup:static-content:deploy
  -   php bin/magento indexer:reindex
  -   rm -rf var/cache
- **Change mode**
- bin/magento deploy:mode:set developer
- rm -rf var/cache generated/metadata generated/code var/view_preprocessed pub/static
