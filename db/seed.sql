CREATE TABLE IF NOT EXISTS items (
    ID INTEGER PRIMARY KEY,
    Name TEXT NOT NULL UNIQUE,
    Notes TEXT,
    Description TEXT,
    Stage INTEGER,
    Category INTEGER
);

CREATE TABLE IF NOT EXISTS boxes (
    ID INTEGER PRIMARY KEY,
    Name TEXT NOT NULL UNIQUE,
    Notes TEXT,
    Description TEXT,
    Stage INTEGER,
    Category INTEGER
);

CREATE TABLE IF NOT EXISTS boxitems (
    ItemID INTEGER UNIQUE, -- items can only exist in a single box
    BoxID INTEGER,
    PRIMARY KEY (ItemID, BoxID),
    FOREIGN KEY (ItemID) REFERENCES items(ID),
    FOREIGN KEY (BoxID) REFERENCES boxes(ID)
);
