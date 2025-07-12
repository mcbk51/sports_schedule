This CLI application offers a streamlined way to view daily schedules for major U.S. sports leagues.

It displays all games for a specified date directly in the terminal. If a game is in progress, the current score is shown along with its status. The app also indicates if a game has been delayed or has already finished. You can view schedules for any past, present, or future date.

In addition to date-based filtering, you can narrow the results by league or team using command-line flags. For example, use --league nba to view only NBA games, or --team "Lakers" to see games involving a specific team. Combine flags like --date 07-14-2025 --league nfl --team "Seahawks" to tailor the output further. By default, the app shows all leagues, all teams, and assumes the current date if no flags are provided.

Output example:
--league NFL --date 09-15-2025

📅 Sports Schedule for NFL - Monday, September 15, 2025
============================================================

🏈 NFL (2 games)
--------------------------------------------------
  4:00 PM  Tampa Bay Buccaneers @ Houston Texans  Scheduled
  7:00 PM  Los Angeles Chargers @ Las Vegas Raiders  Scheduled

============================================================
Total games: 2


--league NFL --date 09-07-2025 --team Seahawks

📅 Sports Schedule for NFL - Sunday, September 7, 2025
============================================================

🏈 NFL (1 games)
--------------------------------------------------
  1:05 PM  San Francisco 49ers @ Seattle Seahawks  Scheduled

============================================================
Total games: 1


📅 Sports Schedule for MLB - Friday, July 11, 2025
============================================================

⚾ MLB (16 games)
--------------------------------------------------
  12:10 PM  Cleveland Guardians @ Chicago White Sox  Final (4-2)

  4:05 PM  Chicago Cubs @ New York Yankees      Final (0-11)

  4:05 PM  Miami Marlins @ Baltimore Orioles    Final (2-5)

  4:10 PM  Colorado Rockies @ Cincinnati Reds   Final (3-2)

  4:10 PM  Seattle Mariners @ Detroit Tigers    Final (12-3)

  4:10 PM  Tampa Bay Rays @ Boston Red Sox      Final (4-5)

  5:10 PM  Cleveland Guardians @ Chicago White Sox  Final (4-5)

  5:10 PM  New York Mets @ Kansas City Royals   Final (8-3)

  5:10 PM  Pittsburgh Pirates @ Minnesota Twins  Final (1-2)

  5:10 PM  Texas Rangers @ Houston Astros       Final (7-3)

  5:10 PM  Washington Nationals @ Milwaukee Brewers  Final (3-8)

  5:15 PM  Atlanta Braves @ St. Louis Cardinals  Final (6-5)

  6:38 PM  Arizona Diamondbacks @ Los Angeles Angels  Final (5-6)

  6:40 PM  Philadelphia Phillies @ San Diego Padres  Final (2-4)

  7:05 PM  Toronto Blue Jays @ Athletics        Final (7-6)

  7:15 PM  Los Angeles Dodgers @ San Francisco Giants  Final (7-8)

============================================================
Total games: 16

