// format handles html rendering

package format

import (
	"fmt"
	"mdf/utils"
	"net/http"
)

func ServeHTML(w http.ResponseWriter, r *http.Request, theme string) {
	html := fmt.Sprintf(`
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>mdf - markdown</title>
    <style>
        body {
            font-family: -apple-system, BlinkMacSystemFont, 'Segoe UI', 'Noto Sans', Helvetica, Arial, sans-serif, 'Apple Color Emoji', 'Segoe UI Emoji';
            line-height: 1.5;
            margin: 0;
            padding: 20px;
        }
        .container {
            max-width: 800px;
            margin: 0 auto;
        }
        %s
    </style>
</head>
<body>
    <div class="container">
        <div id="content"></div>
    </div>
    <script>
        fetch('/')
            .then(response => response.text())
            .then(html => {
                document.getElementById('content').innerHTML = html;
            })
            .catch(error => {
                console.error('Error:', error);
                document.getElementById('content').innerHTML = 'Error loading markdown file.';
            });
    </script>
</body>
</html>`, utils.Themes[theme])

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.Write([]byte(html))
}
