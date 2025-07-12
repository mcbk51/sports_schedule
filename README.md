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
  10:05 AM  Chicago Cubs (56-39) @ New York Yankees (53-42)  Final (5-2)

  10:10 AM  Seattle Mariners (50-45) @ Detroit Tigers (59-37)  Final (15-7)

  11:10 AM  Pittsburgh Pirates (38-58) @ Minnesota Twins (47-48)  Final (4-12)

  11:15 AM  Atlanta Braves (42-52) @ St. Louis Cardinals (50-46)  Final (7-6)

  1:05 PM  Miami Marlins (43-51) @ Baltimore Orioles (43-51)  Final (6-0)

  1:05 PM  Los Angeles Dodgers (57-39) @ San Francisco Giants (52-44)  Final (2-1)

  1:10 PM  Colorado Rockies (22-73) @ Cincinnati Reds (49-47)  Final (3-4)

  1:10 PM  Tampa Bay Rays (50-46) @ Boston Red Sox (52-45)  Final (0-1)

  1:10 PM  Cleveland Guardians (44-49) @ Chicago White Sox (32-63)  In Progress (6-2)

  1:10 PM  New York Mets (55-41) @ Kansas City Royals (46-50)  Final (3-1)

  1:10 PM  Washington Nationals (38-57) @ Milwaukee Brewers (55-40)  Final (5-6)

  4:35 PM  Texas Rangers (47-48) @ Houston Astros (55-39)  Scheduled

  4:35 PM  Philadelphia Phillies (54-40) @ San Diego Padres (51-43)  Scheduled

  6:38 PM  Arizona Diamondbacks (46-49) @ Los Angeles Angels (46-48)  Scheduled

  7:05 PM  Toronto Blue Jays (55-39) @ Athletics (39-57)   Scheduled

============================================================
Total games: 15

