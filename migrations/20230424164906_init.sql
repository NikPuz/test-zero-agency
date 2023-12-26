-- +goose Up

--
-- Структура таблицы `News`
--

CREATE TABLE News (
                        Id SERIAL PRIMARY KEY,
                        Title varchar(64) NOT NULL,
                        Content text NOT NULL
);

-- --------------------------------------------------------

--
-- Структура таблицы `NewsCategories`
--

CREATE TABLE NewsCategories (
                                  NewsId SERIAL NOT NULL,
                                  CategoryId SERIAL NOT NULL
);

--
-- Индексы таблицы `News`
--

CREATE INDEX news_id_index ON News (Id);

--
-- Индексы таблицы `NewsCategories`
--

CREATE INDEX news_categories_news_id_category_id_index ON NewsCategories (NewsId, CategoryId);