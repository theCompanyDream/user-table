#!/usr/bin/env python3
import os
from pathlib import Path
import psycopg2
from ulid import ULID
import random
from faker import Faker
from dotenv import load_dotenv

# Initialize Faker
fake = Faker()
ulid = ULID()

# Example departments list
departments = [
    "accounting", "sales", "engineering", "hr", "marketing", "it", "operations", "marketing", "operations"
]

def get_db_connection():
    # Get database connection parameters from environment variables.
    # Adjust defaults as needed.
    env_path = Path(__file__).resolve().parents[2] / '.env'
    print(env_path)
    # Load environment variables
    load_dotenv(dotenv_path=env_path)
    postgres_user = os.getenv("DATABASE_USERNAME", "postgres")
    postgres_db = os.getenv("DATABASE_NAME")
    postgres_password = os.getenv("DATABASE_PASSWORD")
    postgres_port = os.getenv("DATABASE_PORT")

    conn = psycopg2.connect(
        dbname=postgres_db,
        user=postgres_user,
        password=postgres_password,
        host="localhost",
        port=postgres_port
    )
    return conn

def create_users_table(conn):
    with conn.cursor() as cur:
        cur.execute("""
            CREATE TABLE IF NOT EXISTS USERS (
                ID varchar(26) NOT NULL,
                HASH VARCHAR(64) NOT NULL,
                USER_NAME VARCHAR(50) NOT NULL,
                FIRST_NAME VARCHAR(255) NOT NULL,
                LAST_NAME VARCHAR(255) NOT NULL,
                EMAIL VARCHAR(255) NOT NULL,
                DEPARTMENT VARCHAR(255),
                PRIMARY KEY(ID),
                UNIQUE(EMAIL),
                UNIQUE(USER_NAME)
            );
        """)
        conn.commit()
        print("USERS table created (or already exists).")

def generate_fake_user():
    # Generate fake user data using Faker and Python's uuid
    user_id = str(ulid.generate())
    # Generate a random 64-character hex string (32 bytes * 2 hex digits each)
    hash_value = os.urandom(32).hex()
    user_name = fake.user_name()[:50]
    first_name = fake.first_name()[:255]
    last_name = fake.last_name()[:255]
    email = fake.email()[:255]
    # Example: random user status letter; adjust choices as needed.
    department = random.choice(departments)
    return (user_id, hash_value, user_name, first_name, last_name, email, department)

def insert_fake_users(conn, num_records):
    with conn.cursor() as cur:
        for _ in range(num_records):
            user_data = generate_fake_user()
            cur.execute("""
                INSERT INTO users (id, HASH, USER_NAME, FIRST_NAME, LAST_NAME, EMAIL, DEPARTMENT)
                VALUES (%s, %s, %s, %s, %s, %s, %s);
            """, user_data)
        conn.commit()
        print(f"Inserted {num_records} fake user records.")

def main():
    num_records = 300  # Adjust the number of records you want to generate
    conn = get_db_connection()
    try:
        insert_fake_users(conn, num_records)
    finally:
        conn.close()

if __name__ == '__main__':
    main()
