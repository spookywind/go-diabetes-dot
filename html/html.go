package html

import "strings"

var MY_HTML = `
<!DOCTYPE html>
<html>
<head>
	<title>Go Diabetes Dot</title>
</head>
<body>

<section>
$PROFILE-SECTION$
</section>

<section>
$FOODS-SECTION$
</section>

<section>
$EXERCISE-SECTION$
</section>

<section>
$MEDITATION-SECTION$
</section>

<section>
$WEIGHT-SECTION$
</section>

</body>
</html>
`

func Html() string {
	replacer := strings.NewReplacer(
		"$PROFILE-SECTION$", ProfileSection,
		"$FOODS-SECTION$", FoodsSection,
		"$EXERCISE-SECTION$", ExerciseSection,
		"$MEDITATION-SECTION$", MeditationSection,
		"$WEIGHT-SECTION$", WeightSection,
	)
	html := replacer.Replace(MY_HTML)
	return html
}
