package xiaomiAvail

import (
    "html/template"
    "net/http"
	"appengine"
	"appengine/urlfetch"
	"io/ioutil"
	"log"
	"regexp"
)

type XiaoMiDevice struct {
		Note4G				string
		Redmi1S				string
		Mi3					string
		Mipowerbank10400	string
		Mipowerbank5200		string
		Miband				string
}

func init() {
    http.HandleFunc("/", root)
	http.Handle("/css/", http.FileServer(http.Dir(".")))
	http.Handle("/js/", http.FileServer(http.Dir(".")))
}


func root(w http.ResponseWriter, r *http.Request) {
	
	XiaoMiDevice := new(XiaoMiDevice);
	XiaoMiDevice.Note4G = xiaomiSearch("http://www.mi.com/sg/note4g/", w, r)
	XiaoMiDevice.Redmi1S = xiaomiSearch("http://www.mi.com/sg/redmi1s/", w, r)
	XiaoMiDevice.Mi3 = xiaomiSearch("http://www.mi.com/sg/mi3/", w, r)
	XiaoMiDevice.Mipowerbank10400 = xiaomiSearch("http://www.mi.com/sg/mipowerbank10400/", w, r)
	XiaoMiDevice.Mipowerbank5200 = xiaomiSearch("http://www.mi.com/sg/mipowerbank5200/", w, r)
	XiaoMiDevice.Miband = xiaomiSearch("http://www.mi.com/sg/miband/", w, r)
	
	xiaomiAvailForm.ExecuteTemplate(w, "xiaomiAvail.htm", XiaoMiDevice)
}

func xiaomiSearch(url string, w http.ResponseWriter, r *http.Request) string {

	c := appengine.NewContext(r)
	client := urlfetch.Client(c)
    resp, err := client.Get(url)
    
    if err != nil {
        log.Fatal(err)
    }
	
	robots, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		log.Fatal(err)
	}
	re := regexp.MustCompile("(J-btn.disabled)|(disabled.J-btn)")
	match := re.MatchString(string(robots))
	if match {
		return "Not Available";
	} else {
		return url;
	}
	
}


var xiaomiAvailForm = template.Must(template.New("").ParseFiles("xiaomiAvail.htm"))