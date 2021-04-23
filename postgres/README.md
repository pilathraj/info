## Postgres SQL
 ```
  cd compose
  docker-compose up -d
 ```
 ### pg admin
 1. localhost:5050
    * Add a new server in PgAdmin:
       - Host name/address postgres
       - Port 5432
       - Username as POSTGRES_USER, by default: postgres
       - Password as POSTGRES_PASSWORD, by default changeme
    * Click Database > Query Tool > Select Version() > Run
       - "PostgreSQL 13.2 (Debian 13.2-1.pgdg100+1) on x86_64-pc-linux-gnu, compiled by gcc (Debian 8.3.0-6) 8.3.0, 64-bit"
 3. docker exec -it ae8f6b2bbe4b psql -U postgres
    - Select Version();
      *   version
------------------------------------------------------------------------------------------------------------------
 PostgreSQL 13.2 (Debian 13.2-1.pgdg100+1) on x86_64-pc-linux-gnu, compiled by gcc (Debian 8.3.0-6) 8.3.0, 64-bit
(1 row)
### What is Postgres Sql?
  - RDBMS
  - Advanced, enterprise-class, Open source Database.
  - Support both Relational(SQL) and Non-Relational(JSON) Queries.
  - Primary database for many
     * Web application
     * Mobile Application
     * Data analytics Application

### Feature highlights
   - User defined types
   - Table inheritance
   - Sophisticated Locking mechanism 
   - Foreign Key referential integrity 
   - Views, rules, and subquery
   - Nested transcations (savepoint)
   - Multi-version concurrency control(MVCC)
   - Asynchronous Replication
   - Tablespaces
   - Point in time recovery
   - Native Microsoft windows server version. 
### load sample data
- Get volumn path, look at the volume paths under the key **Destination**.
```psql
docker inspect -f '{{ json .Mounts }}' abc985ddffcf | python -m json.tool
docker cp my_data.dump my_postgres_1:/backups
docker exec my_postgres_1 pg_restore -U postgres -d some_database /backups/my_data.dump
```
  
### PSQL Commands
- list all databases
```psql
 \l
```
- connect database
```psql
\c <databaseName>
```
- quit database connection
```psql
\q
```
### Query 
1. Select
```psql
SELECT first_name FROM customer;
SELECT first_name || ' ' || last_name as full_name FROM customer;
SELECT 5 * 3;
```
2. Order by
  - ORDER BY sort_expresssion [ASC | DESC] [NULLS FIRST | NULLS LAST]
```psql
SELECT first_name, last_name FROM customer ORDER BY first_name ASC, last_name DESC;
SELECT num FROM sort_demo ORDER BY num NULLS LAST;
```
3. Distinct
 - recommended to use order by, when used multiple columns in the distinct
```psql
SELECT DISTINCT bcolor, fcolor FROM distinct_demo ORDER BY bcolor, fcolor;
SELECT DISTINCT ON (bcolor) bicolor, color FROM distinct_demo ORDER BY bcolor,fcolor;
```

