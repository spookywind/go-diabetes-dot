# go-diabetes-dot

## Purpose of this App
* To automate a spreadsheet used to help with efforts to manage my diabetes by:
  * predicting what my long-term future weight would be based on daily impact of food and exercise (using Rosati's equation)
  * contain additional helpful info - e.g. calculate what my A1C would be if I ate yesterday's diet every day
  
## Rosati's equation
* estimated_future_weight = (food_calories - exercise_calories) / 10
* exercise_calories = miles_walked_equiv * 100 calories

### exercise-calories/minute and METS
* from https://metscalculator.com/
  * lying quietly watching television for 1 minute (no exercise, 1 METS, @200 lbs) for 1 minutes = 1.5 calories
  * walking 2.5 mph - firm surface for 1 minute (moderate exercise, 3.0 METS, @200 lbs) = 4.5 calories
  * lying quietly watching television for 1 minute (no exercise, 1 METS, @120 lbs) for 1 minutes = 0.9 calories
  * walking 2.5 mph - firm surface for 1 minute (moderate exercise, 3.0 METS, @120 lbs) = 2.7 calories
  * observation - calories_burned goes down with weight, and up with intensity 
* conclusion
  * generally the moderate exercise I am doing is more intense that walking 2.5 mph
  * I am estimating 3.3 calories / minute for moderate exercise
  
### my variant on Rosati's equation
* 1 exercise serving == 10 minutes of moderate exercise
* at moderate intensity - a mile is walked in 20 minutes (3 mph)
* exercise_calories = miles_walked_equiv * 66 calories
* I am using 2/3 rosati's 100 calorie value (a more conservative value)
* NOTE: 

## Possible future work
* At some point in the future, I may use Data Science tools to also predict future A1C
