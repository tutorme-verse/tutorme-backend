CREATE TABLE Schools (
    School_ID INTEGER PRIMARY KEY AUTOINCREMENT, -- Auto-incremented unique ID
    School_Name TEXT NOT NULL,                  -- Full name of the organization
    Subdomain TEXT NOT NULL UNIQUE,                  -- Subdomain assigned to the organization
    Created_At TIMESTAMP DEFAULT CURRENT_TIMESTAMP,   -- Timestamp for record creation
    Status TEXT NOT NULL DEFAULT 'active',           -- Status of the organization
    CHECK (Status IN ('active', 'inactive', 'deleted')) -- Constraint to enforce valid status values
);

CREATE TABLE DatabaseDetails (
    Database_ID INTEGER PRIMARY KEY AUTOINCREMENT,    -- Auto-incremented unique ID
    School_ID INTEGER NOT NULL,                 -- Foreign key to Organizations table
    Database_Name TEXT NOT NULL UNIQUE,               -- Unique name of the Turso database
    Connection_URI TEXT NOT NULL,                     -- Connection URI for the database
    Created_At TIMESTAMP DEFAULT CURRENT_TIMESTAMP,   -- Timestamp for record creation
    FOREIGN KEY (School_ID) REFERENCES Schools(School_ID) ON DELETE CASCADE -- Cascade delete
);

