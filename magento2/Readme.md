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
```
- **clear cache**
```cmd
php bin/magento cache:flush
```

