// containing some common themes

package utils

var Themes = map[string]string{
	"github-dark": `
		body { color: #c9d1d9; background-color: #0d1117; }
		a { color: #58a6ff; }
		h1, h2 { border-bottom: 1px solid #21262d; }
		code { background-color: rgba(110,118,129,0.4); padding: 0.2em 0.4em; border-radius: 6px; }
		pre { background-color: #161b22; padding: 16px; overflow: auto; border-radius: 6px; }
	`,
	"github-light": `
		body { color: #24292f; background-color: #ffffff; }
		a { color: #0969da; }
		h1, h2 { border-bottom: 1px solid #d0d7de; }
		code { background-color: rgba(175,184,193,0.2); padding: 0.2em 0.4em; border-radius: 6px; }
		pre { background-color: #f6f8fa; padding: 16px; overflow: auto; border-radius: 6px; }
	`,
	"dracula": `
		body { color: #f8f8f2; background-color: #282a36; }
		a { color: #8be9fd; }
		h1, h2 { border-bottom: 1px solid #6272a4; }
		code { background-color: #44475a; padding: 0.2em 0.4em; border-radius: 6px; }
		pre { background-color: #44475a; padding: 16px; overflow: auto; border-radius: 6px; }
	`,
	"nord": `
		body { color: #d8dee9; background-color: #2e3440; }
		a { color: #88c0d0; }
		h1, h2 { border-bottom: 1px solid #4c566a; }
		code { background-color: #3b4252; padding: 0.2em 0.4em; border-radius: 6px; }
		pre { background-color: #3b4252; padding: 16px; overflow: auto; border-radius: 6px; }
	`,
}
