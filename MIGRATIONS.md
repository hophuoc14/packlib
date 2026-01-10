# Database Migrations Guide

This document explains how to manage database migrations for the packlib project.

## Overview

We use [golang-migrate](https://github.com/golang-migrate/migrate) for database schema versioning and migration management. This provides:
- **Version control** for database schema changes
- **Rollback capability** to undo migrations
- **Migration history** tracking in the database
- **Both CLI and programmatic** migration execution

## Migration Files

Migrations are stored in `db/migrations/` directory:
```
db/migrations/
├── 000001_initial_schema.up.sql     # Creates tables
├── 000001_initial_schema.down.sql   # Drops tables
├── 000002_add_feature.up.sql        # Next migration
└── 000002_add_feature.down.sql      # Rollback for next migration
```

Each migration has two files:
- **`.up.sql`** - Forward migration (applies changes)
- **`.down.sql`** - Reverse migration (rolls back changes)

## Running Migrations

### Option 1: Manual via CLI Tool (Recommended)

Run migrations manually using the CLI tool:

```bash
# Apply all pending migrations
go run cmd/migrate/main.go up

# Rollback the last migration
go run cmd/migrate/main.go down

# Check current migration version
go run cmd/migrate/main.go version

# Force set version (recovery only - use with caution!)
go run cmd/migrate/main.go force 1
```

### Option 2: Automatic on Application Startup

Uncomment the migration code in `config/postgres.go` to run migrations automatically when the application starts:

```go
// In config/postgres.go
import (
    "log"
    dbmigrate "packlib/db"
)

// In NewPostgresDatabase function
if err := dbmigrate.RunMigrations(sqlDB, dbName); err != nil {
    log.Printf("Warning: Migration failed: %v", err)
    // return nil, err  // Uncomment to fail startup on migration error
}
```

> ⚠️ **Note**: Automatic migrations are convenient for development but consider running migrations manually in production environments for better control.

## Creating New Migrations

### Method 1: Install and Use migrate CLI (Recommended)

1. **Install the migrate CLI** (one-time setup):
   ```bash
   go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest
   ```

2. **Create a new migration**:
   ```bash
   migrate create -ext sql -dir db/migrations -seq add_user_avatar_column
   ```

   This creates two files:
   - `000002_add_user_avatar_column.up.sql`
   - `000002_add_user_avatar_column.down.sql`

3. **Edit the migration files**:
   
   In `000002_add_user_avatar_column.up.sql`:
   ```sql
   ALTER TABLE users ADD COLUMN avatar VARCHAR(255);
   ```
   
   In `000002_add_user_avatar_column.down.sql`:
   ```sql
   ALTER TABLE users DROP COLUMN avatar;
   ```

4. **Run the migration**:
   ```bash
   go run cmd/migrate/main.go up
   ```

### Method 2: Manual File Creation

1. Determine the next sequence number (e.g., if last migration is `000001`, use `000002`)
2. Create two files:
   - `db/migrations/000002_description.up.sql`
   - `db/migrations/000002_description.down.sql`
3. Write your SQL changes in the `.up.sql` file
4. Write the reverse SQL changes in the `.down.sql` file
5. Run the migration

## Migration Best Practices

### 1. **Always Write Down Migrations**
Every `.up.sql` should have a corresponding `.down.sql` that reverses the changes.

### 2. **Make Migrations Atomic**
Each migration should be a single, focused change:
- ✅ Good: `add_user_avatar_column`
- ❌ Bad: `update_user_table_and_add_posts_table`

### 3. **Use IF EXISTS / IF NOT EXISTS**
Make migrations idempotent when possible:
```sql
-- Up migration
ALTER TABLE users ADD COLUMN IF NOT EXISTS avatar VARCHAR(255);

-- Down migration
ALTER TABLE users DROP COLUMN IF EXISTS avatar;
```

### 4. **Test Migrations**
Always test both up and down migrations:
```bash
go run cmd/migrate/main.go up      # Apply migration
go run cmd/migrate/main.go down    # Rollback migration
go run cmd/migrate/main.go up      # Apply again
```

### 5. **Never Modify Existing Migrations**
Once a migration has been applied in any environment (especially production):
- ❌ Never modify it
- ✅ Create a new migration to make changes

### 6. **Keep Migrations Fast**
Avoid operations that lock tables for extended periods. For large tables, consider:
- Adding indexes `CONCURRENTLY` (PostgreSQL)
- Batching data migrations
- Running heavy operations during maintenance windows

## Migration Naming Conventions

Use descriptive, lowercase names with underscores:
- `create_users_table`
- `add_email_index`
- `add_user_avatar_column`
- `remove_deprecated_fields`
- `create_departments_employees_junction`

## Troubleshooting

### "Dirty Database" Error

If a migration fails midway, the database is marked as "dirty":

```bash
# Check version and dirty status
go run cmd/migrate/main.go version

# If dirty, fix the database manually, then force the version
go run cmd/migrate/main.go force 1
```

### Migration Already Applied

If you see "no change" when running `up`, all migrations are already applied. Check with:
```bash
go run cmd/migrate/main.go version
```

### Database Connection Issues

Ensure your `.env` file has correct database credentials:
```
DATABASE_HOST=localhost
DATABASE_PORT=5432
DATABASE_USERNAME=your_user
DATABASE_PASSWORD=your_password
DATABASE_NAME=your_database
```

## Migration Version Tracking

Migrations are tracked in the `schema_migrations` table:
```sql
SELECT * FROM schema_migrations;
```

This table stores:
- `version`: Current migration version number
- `dirty`: Whether the last migration failed (boolean)

## Example Workflow

Here's a complete example of adding a new feature:

```bash
# 1. Create migration files
migrate create -ext sql -dir db/migrations -seq add_user_roles

# 2. Edit 000002_add_user_roles.up.sql
# ALTER TABLE users ADD COLUMN role VARCHAR(50) DEFAULT 'user';

# 3. Edit 000002_add_user_roles.down.sql
# ALTER TABLE users DROP COLUMN role;

# 4. Apply migration
go run cmd/migrate/main.go up

# 5. Verify in database
# psql -d your_database -c "\d users"

# 6. If something goes wrong, rollback
go run cmd/migrate/main.go down
```

## Additional Resources

- [golang-migrate Documentation](https://github.com/golang-migrate/migrate)
- [PostgreSQL ALTER TABLE](https://www.postgresql.org/docs/current/sql-altertable.html)
- [Database Migration Best Practices](https://www.brunton-spall.co.uk/post/2014/05/06/database-migrations-done-right/)
