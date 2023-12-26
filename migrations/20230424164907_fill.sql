-- +goose Up

--
-- Заполнение таблицы `News`
--

INSERT INTO News (Title, Content) VALUES ('Заголовок1', 'Наполнение1');
INSERT INTO News (Title, Content) VALUES ('Заголовок2', 'Наполнение2');
INSERT INTO News (Title, Content) VALUES ('Заголовок3', 'Наполнение3');
INSERT INTO News (Title, Content) VALUES ('Заголовок4', 'Наполнение4');
INSERT INTO News (Title, Content) VALUES ('Заголовок5', 'Наполнение5');
INSERT INTO News (Title, Content) VALUES ('Заголовок6', 'Наполнение6');
INSERT INTO News (Title, Content) VALUES ('Заголовок7', 'Наполнение7');
INSERT INTO News (Title, Content) VALUES ('Заголовок8', 'Наполнение8');
INSERT INTO News (Title, Content) VALUES ('Заголовок9', 'Наполнение9');
INSERT INTO News (Title, Content) VALUES ('Заголовок10', 'Наполнение10');
INSERT INTO News (Title, Content) VALUES ('Заголовок11', 'Наполнение11');
INSERT INTO News (Title, Content) VALUES ('Заголовок12', 'Наполнение12');
INSERT INTO News (Title, Content) VALUES ('Заголовок13', 'Наполнение13');
INSERT INTO News (Title, Content) VALUES ('Заголовок14', 'Наполнение14');
INSERT INTO News (Title, Content) VALUES ('Заголовок15', 'Наполнение15');
INSERT INTO News (Title, Content) VALUES ('Заголовок16', 'Наполнение16');
INSERT INTO News (Title, Content) VALUES ('Заголовок17', 'Наполнение17');
INSERT INTO News (Title, Content) VALUES ('Заголовок18', 'Наполнение18');
INSERT INTO News (Title, Content) VALUES ('Заголовок19', 'Наполнение19');

--
-- Заполнение таблицы `NewsCategories`
--

INSERT INTO NewsCategories (newsid, categoryid) VALUES (1, 2);
INSERT INTO NewsCategories (newsid, categoryid) VALUES (2, 3);
INSERT INTO NewsCategories (newsid, categoryid) VALUES (2, 4);
INSERT INTO NewsCategories (newsid, categoryid) VALUES (3, 1);
INSERT INTO NewsCategories (newsid, categoryid) VALUES (3, 26);
INSERT INTO NewsCategories (newsid, categoryid) VALUES (3, 44);
INSERT INTO NewsCategories (newsid, categoryid) VALUES (3, 34);
INSERT INTO NewsCategories (newsid, categoryid) VALUES (3, 4);
INSERT INTO NewsCategories (newsid, categoryid) VALUES (3, 3);
INSERT INTO NewsCategories (newsid, categoryid) VALUES (3, 2);
INSERT INTO NewsCategories (newsid, categoryid) VALUES (5, 2);
INSERT INTO NewsCategories (newsid, categoryid) VALUES (6, 2);
INSERT INTO NewsCategories (newsid, categoryid) VALUES (7, 2);
INSERT INTO NewsCategories (newsid, categoryid) VALUES (8, 2);
INSERT INTO NewsCategories (newsid, categoryid) VALUES (9, 28);
INSERT INTO NewsCategories (newsid, categoryid) VALUES (9, 25);
INSERT INTO NewsCategories (newsid, categoryid) VALUES (9, 22);
INSERT INTO NewsCategories (newsid, categoryid) VALUES (13, 2);
INSERT INTO NewsCategories (newsid, categoryid) VALUES (14, 2);
INSERT INTO NewsCategories (newsid, categoryid) VALUES (15, 2);
INSERT INTO NewsCategories (newsid, categoryid) VALUES (16, 2);
INSERT INTO NewsCategories (newsid, categoryid) VALUES (17, 2);
INSERT INTO NewsCategories (newsid, categoryid) VALUES (18, 2);
INSERT INTO NewsCategories (newsid, categoryid) VALUES (19, 2);