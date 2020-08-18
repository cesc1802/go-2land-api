insert into comments(id, content, post_id) values (1, "comment on first post", 1);
insert into comments(id, content, post_id) values (3, "second comment on first post", 1);
insert into comments(id, content, post_id) values (2, "comment on second post", 2);


select posts.id, c.id, posts.content, c.content from posts left join comments c on posts.id = c.post_id where posts.id = 1;


select * from posts join comments c on posts.id = c.post_id where posts.id = 1;