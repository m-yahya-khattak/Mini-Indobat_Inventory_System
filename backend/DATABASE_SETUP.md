# Database Setup Guide

Follow these steps once PostgreSQL installation is complete.

## Step 1: Verify PostgreSQL Installation
```powershell
psql --version
```
Expected output: `psql (PostgreSQL) 16.x` or similar

If command not found:
- Restart your PowerShell terminal
- Or add PostgreSQL bin to PATH: `C:\Program Files\PostgreSQL\16\bin`

## Step 2: Start PostgreSQL Service (if not running)
```powershell
# Check if service is running
Get-Service postgresql*

# Start service if needed (run as Administrator)
Start-Service postgresql-x64-16
# (Version number might be different - check what service was created)
```

## Step 3: Create Database

### Option A: Using createdb command
```powershell
createdb -U postgres indobat_inventory
```
You'll be prompted for the PostgreSQL password.

### Option B: Using psql (Recommended)
```powershell
# Connect to PostgreSQL
psql -U postgres
```

Then in the psql prompt, run:
```sql
CREATE DATABASE indobat_inventory;
\q
```

You'll be prompted for the password you set during PostgreSQL installation.

## Step 4: Verify Database Creation
```powershell
psql -U postgres -l
```
Look for `indobat_inventory` in the list.

Or connect directly:
```powershell
psql -U postgres -d indobat_inventory
\q
```

## Step 5: Update .env File
Make sure your `backend/.env` file has the correct credentials:
```env
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=your_actual_password_here
DB_NAME=indobat_inventory
DB_SSLMODE=disable
```

**Important**: Replace `your_actual_password_here` with the password you set during PostgreSQL installation.

## Step 6: Test Database Connection
Once you have the database connection code implemented, you can test it by running your Go application.

## Troubleshooting

### "psql: command not found"
- Restart PowerShell terminal
- Or use full path: `"C:\Program Files\PostgreSQL\16\bin\psql.exe" -U postgres`

### "password authentication failed"
- Make sure you're using the correct password set during installation
- Default user is `postgres`

### "database already exists"
- That's fine! The database is already created, you can skip creation step

### "connection refused" or "could not connect"
- Make sure PostgreSQL service is running
- Check if port 5432 is correct
- Verify firewall settings

