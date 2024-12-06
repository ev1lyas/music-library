INSERT INTO songs (musicGroup, song, releaseDate, text, link)
VALUES
    ('Adele', 'Someone Like You', '2011-01-24', E'I heard that you''re settled down\n\nThat you found a girl and you''re married now', 'https://example.com/someone-like-you'),
    ('The Weeknd', 'Save Your Tears', '2020-03-20', E'I saw you dancing in a crowded room\n\nYou look so happy when I''m not with you', 'https://example.com/save-your-tears'),
    ('Ed Sheeran', 'Perfect', '2017-09-03', E'I found a love, to carry more than just my secrets\n\nTo carry love, to carry children of our own', 'https://example.com/perfect'),
    ('Taylor Swift', 'Love Story', '2008-09-15', E'We were both young when I first saw you\n\nI close my eyes and the flashback starts', 'https://example.com/love-story'),
    ('Imagine Dragons', 'Demons', '2013-02-17', E'When the days are cold\n\nAnd the cards all fold', 'https://example.com/demons'),
    ('Kendrick Lamar', 'DNA.', '2017-04-14', E'I got, I got, I got, I got\n\nLoyalty, got royalty inside my DNA', 'https://example.com/dna'),
    ('Billie Eilish', 'Everything I Wanted', '2019-11-13', E'I had a dream I got everything I wanted\n\nNot what you''d think', 'https://example.com/everything-i-wanted'),
    ('Lorde', 'Green Light', '2017-03-02', E'I do my makeup in somebody else''s car\n\nWe order different drinks at the same bars', 'https://example.com/green-light'),
    ('Post Malone', 'Rockstar', '2017-09-15', E'I''ve been a, I''ve been a rockstar\n\nRockstar', 'https://example.com/rockstar'),
    ('Sia', 'Elastic Heart', '2013-12-17', E'I''ve got thick skin and an elastic heart\n\nBut your blade, it might be too sharp', 'https://example.com/elastic-heart'),
    ('Travis Scott', 'Goosebumps', '2016-09-16', E'Yeah, I get those goosebumps every time, yeah\n\nYou come around, yeah', 'https://example.com/goosebumps'),
    ('Lady Gaga', 'Poker Face', '2008-09-26', E'I wanna hold ''em like they do in Texas, please\n\nFold ''em, let ''em hit me, raise it, baby, stay with me (I love it)', 'https://example.com/poker-face'),
    ('Drake', 'In My Feelings', '2018-07-10', E'Kiki, do you love me?\n\nAre you riding?', 'https://example.com/in-my-feelings'),
    ('Katy Perry', 'Teenage Dream', '2010-07-23', E'You think I''m pretty without any makeup on\n\nYou think I''m funny when I tell the punchline wrong', 'https://example.com/teenage-dream'),
    ('Beyoncé', 'Single Ladies (Put a Ring on It)', '2008-10-13', E'All the single ladies\n\nAll the single ladies', 'https://example.com/single-ladies'),
    ('The Chainsmokers', 'Don''t Let Me Down', '2016-02-05', E'Hey, I''m waiting for you\n\nDon''t let me down', 'https://example.com/dont-let-me-down'),
    ('Coldplay', 'Adventure of a Lifetime', '2015-11-06', E'Turn your magic on\n\nTo me she''d say', 'https://example.com/adventure-of-a-lifetime'),
    ('Maroon 5', 'Moves Like Jagger', '2011-06-21', E'Just shoot for the stars\n\nIf it feels right', 'https://example.com/moves-like-jagger'),
    ('OneRepublic', 'Secrets', '2009-11-17', E'I need another story\n\nSomething to get off my chest', 'https://example.com/secrets'),
    ('Miley Cyrus', 'Party in the USA', '2009-04-29', E'I hopped on the plane at L.A.X.\n\nWith a dream and my cardigan', 'https://example.com/party-in-the-usa'),
    ('Foo Fighters', 'The Pretender', '2007-07-16', E'Keep you in the dark\n\nYou know they all pretend', 'https://example.com/the-pretender')
    ON CONFLICT (musicGroup, song) DO NOTHING;
