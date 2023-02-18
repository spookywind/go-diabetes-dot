# go-diabetes-dot

## Purpose of this App
* To automate a spreadsheet used to help with efforts to manage my diabetes by:
  * predicting what my long-term future weight would be based on daily impact of food and exercise (using Rosati's equation)
  * contain additional helpful info - e.g. calculate what my A1C would be if I ate yesterday's diet every day
  
## Estimated future weight
* from https://www.aarp.org/health/fitness/info-07-2010/How_Many_Calories_Do_I_Need.html 
* from Rosati: estimated_future_weight = calories / 10
* calories == food_calories - exercise_calories

### exercise-calories/minute and METS
* from https://metscalculator.com/
  * Moderate exercise starts at 3.0 METS
  * a 150-lb person walking @ 2.5 mph (about 3.0 METS) burns 3.4 calories/minute
  * this value goes down with a person's weight, and up with increased intensity 

### exercise calories
* 1 exercise portion == light (20 minutes -- 1-3 METS), moderate (10 minutes -- 3-6 METS), vigorous (5 minutes, 6-9 METS), intense (2.5, 9+ METS)
* 1 exercise portion == ~34 calories
* exercise_calories == exercise portions * 10

## Possible future work
* At some point in the future, I may use Data Science tools to also predict future A1C
