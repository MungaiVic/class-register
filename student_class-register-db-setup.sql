DROP TABLE IF EXISTS album;
CREATE TABLE student (
    id INT AUTO_INCREMENT NOT NULL,
    stuName VARCHAR(128) NOT NULL,
    age INT NOT NULL,
    isInClass BOOLEAN NOT NULL,
    PRIMARY KEY (`id`)
);
INSERT INTO student (stuName, age, isInClass)
VALUES ('John Coltrane', 56, false),
    ('Gerry Mulligan', 17, false),
    ('Sarah Vaughan', 34, false);

DROP TABLE IF EXISTS class;
CREATE TABLE class (
    id INT AUTO_INCREMENT NOT NULL,
    className VARCHAR(128) NOT NULL,
    maxSize INT NOT NULL,
    PRIMARY KEY (`id`),
    UNIQUE KEY `className` (`className`)
);

DROP TABLE IF EXISTS classtime;
CREATE TABLE classtime (
    id INT AUTO_INCREMENT NOT NULL,
    classId INT NOT NULL,
    startTime VARCHAR(128) NOT NULL,
    endTime VARCHAR(128) NULL,
    PRIMARY KEY (`id`),
    FOREIGN KEY (classId) REFERENCES class(id)
);
