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
$FUTURE-WEIGHT-SECTION$
</section>

</body>
</html>
`

var ProfileSection = `
MY PROFILE SECTION
`

var FoodsSection = `
MY FOODS SECTION
`

var ExerciseSection = `
MY EXERCISE SECTION
`

var MeditationSection = `
MY MEDITATION SECTION
`

var FutureWeightSection = `
MY FUTURE WEIGHT SECTION
`

func Html() string {
	replacer := strings.NewReplacer(
		"$PROFILE-SECTION$", ProfileSection,
		"$FOODS-SECTION$", FoodsSection,
		"$EXERCISE-SECTION$", ExerciseSection,
		"$MEDITATION-SECTION$", MeditationSection,
		"$FUTURE-WEIGHT-SECTION$", FutureWeightSection,
	)
	html := replacer.Replace(MY_HTML)
	return html
}
