sudo -u postgres psql << EOF
CREATE USER gorm_user WITH PASSWORD 'gorm_password';
CREATE DATABASE gorm_db;
 
GRANT ALL PRIVILEGES ON DATABASE gorm_db TO gorm_user;

EOF
