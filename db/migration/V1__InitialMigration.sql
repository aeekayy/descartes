CREATE EXTENSION "uuid-ossp" IF NOT EXISTS;

CREATE TABLE builds (
    id uuidv4 NOT NULL PRIMARY KEY DEFAULT uuid_generate_v4(),
    repository varchar(100) NOT NULL, -- GitHub repository names have a character limit of 100
    git_commit varchar(72) NOT NULL, 
    branch varchar(256),
    created timestampz DEFAULT NOW(),
    updated timestampz DEFAULT NOW()
)

CREATE TABLE resources (
    id uuidv4 NOT NULL PRIMARY KEY DEFAULT uuid_generate_v4(),
    name varchar(256) NOT NULL,
    created timestampz DEFAULT NOW(),
    updated timestampz DEFAULT NOW()
)