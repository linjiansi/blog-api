create table if not exists articles (
    article_id integer unsigned auto_increment primary key,
    title varchar(100) not null,
    contents text not null,
    username varchar(100) not null,
    favorite integer not null,
    created_at datetime not null,
);

create table if not exists comments (
    comment_id integer unsigned auto_increment primary key,
    article_id integer unsigned not null,
    message text not null,
    created_at datetime not null,
    foreign key (article_id) references articles(article_id)
);

insert into articles (title, contents, username, favorite, created_at) values ('first article', 'This is my first blog', 'linjiansi', 8, now());
insert into articles (title, contents, username, favorite, created_at) values ('second article', 'This is my second blog', 'linjiansi', 7, now());
insert into comments (article_id, message, created_at) values (1, 'first comment', now());
insert into comments (article_id, message, created_at) values (1, 'second comment', now());
