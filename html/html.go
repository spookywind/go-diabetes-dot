package html

import "strings"

var MY_HTML = `
<!DOCTYPE html>
<html>
<head>
	<title>Go Diabetes Dot</title>
</head>
<body>
$CONTENTS$
</body>
</html>
`

var CONTENTS = `
MY CONTENTS
`

func Html() string {
	replacer := strings.NewReplacer("$CONTENTS$", CONTENTS)
	html := replacer.Replace(MY_HTML)
	return html
}
