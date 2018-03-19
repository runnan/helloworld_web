package route

import (
  "net/http"
  "fmt"
  "html/template"
  "bytes"
)

var layoutFuncs = template.FuncMap{
  "yield": func() (string, error) {
    return "", fmt.Errorf("yield called inappropriately")
  },
}
var layout = template.Must(
  template.
  New("layout.html").
  Funcs(layoutFuncs).
  ParseFiles("templates/layout.html"),
)

var templates = template.Must(template.New("t").ParseGlob("templates/**/*.html"))
func RenderTemplate(w http.ResponseWriter, r *http.Request, name string, data map[string]interface{}) {
  cookie, errCookie := r.Cookie("Token")
  if data == nil {
    data = map[string]interface{}{}
  }
  if errCookie != nil {
    data["CurrentUser"] = nil
  }else{
    info, result := ExtractClaims(cookie.Value)
    if result{
      data["CurrentUser"] = info["username"]
    }else{
      data["CurrentUser"] = nil
    }

  }

  funcs := template.FuncMap{
    "yield": func() (template.HTML, error) {
      buf := bytes.NewBuffer(nil)
      err := templates.ExecuteTemplate(buf, name, data)
      return template.HTML(buf.String()), err
    },
  }
  layoutClone, _ := layout.Clone()
  layoutClone.Funcs(funcs)
  err := layoutClone.Execute(w, data)
  if err != nil {
    http.Error(w,fmt.Sprintf(errorTemplate, name, err),http.StatusInternalServerError,)
  }
}

var errorTemplate = `
<html>
<body>
<h1>Error rendering template %s</h1>
<p>%s</p>
</body>
</html>
`
