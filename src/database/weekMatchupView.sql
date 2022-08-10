CREATE or REPLACE VIEW weekMatchupView AS
    select week_id, w.name as week_name, w.start as week_start, w.end as week_end, mu.id as matchup_id, mu.away_team_id as away_team_id, mu.home_team_id as home_team_id, ht.city as home_team_city, ht.name as home_team_name, at.city as away_team_city, at.name as away_team_name from matchups mu
    join teams at on mu.away_team_id = at.id
    join teams ht on mu.home_team_id = ht.id
    join weeks w on mu.week_id = w.id;