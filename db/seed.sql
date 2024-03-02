-- CREATE DATABASE IF NOT EXISTS inventory;
-- USE inventory;

CREATE TABLE IF NOT EXISTS items (
    ID INTEGER PRIMARY KEY,
    Name TEXT NOT NULL,
    Notes TEXT,
    Description TEXT,
    Stage INTEGER,
    Category INTEGER
);

CREATE TABLE IF NOT EXISTS boxes (
    ID INTEGER PRIMARY KEY,
    Name TEXT NOT NULL,
    Notes TEXT,
    Description TEXT,
    Stage INTEGER,
    Category INTEGER
);

CREATE TABLE IF NOT EXISTS boxitems (
    ItemID INTEGER,
    BoxID INTEGER,
    PRIMARY KEY (ItemID, BoxID),
    FOREIGN KEY (ItemID) REFERENCES items(ID),
    FOREIGN KEY (BoxID) REFERENCES boxes(ID)
);
