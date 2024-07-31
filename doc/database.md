
## Setup test db on local

Create testdb
```sh
createdb testdb
```

Access to testdb
```sh
psql testdb
```

Run these query
```sql
CREATE USER testuser WITH PASSWORD 'testpassword';

GRANT CREATE ON SCHEMA public TO testuser;
```

Create Migrate 
```sh
./devtools/create_migrate.sh [name_migrate]
```

Run Migrate
```sh
./devtools/migrate.sh [up|down|force|version]
```