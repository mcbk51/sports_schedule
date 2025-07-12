This CLI application offers a streamlined way to view daily schedules for major U.S. sports leagues.

It displays all games for a specified date directly in the terminal. If a game is in progress, the current score is shown along with its status. The app also indicates if a game has been delayed or has already finished. You can view schedules for any past, present, or future date. The current win/loss record is displayed.     

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


go run .

📅 Sports Schedule for ALL - Saturday, July 12, 2025
============================================================

⚾ MLB (15 games)
--------------------------------------------------
  10:05 AM  Chicago Cubs @ New York Yankees      In Progress (1-0)

  10:10 AM  Seattle Mariners @ Detroit Tigers    Delayed

  11:10 AM  Pittsburgh Pirates @ Minnesota Twins  Scheduled

  11:15 AM  Atlanta Braves @ St. Louis Cardinals  Scheduled

  1:05 PM  Miami Marlins @ Baltimore Orioles    Scheduled

  1:05 PM  Los Angeles Dodgers @ San Francisco Giants  Scheduled

  1:10 PM  Colorado Rockies @ Cincinnati Reds   Scheduled

  1:10 PM  Tampa Bay Rays @ Boston Red Sox      Scheduled

  1:10 PM  Cleveland Guardians @ Chicago White Sox  Scheduled

  1:10 PM  New York Mets @ Kansas City Royals   Scheduled

  1:10 PM  Washington Nationals @ Milwaukee Brewers  Scheduled

  4:35 PM  Texas Rangers @ Houston Astros       Scheduled

  4:35 PM  Philadelphia Phillies @ San Diego Padres  Scheduled

  6:38 PM  Arizona Diamondbacks @ Los Angeles Angels  Scheduled

  7:05 PM  Toronto Blue Jays @ Athletics        Scheduled

============================================================
Total games: 15

