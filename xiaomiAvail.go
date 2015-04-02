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
		Redmi2				string
		MiPad				string
		Mipowerbank16000	string
		Mipowerbank10400	string
		Mipowerbank5000		string
		Miband				string
}

func init() {
    http.HandleFunc("/", root)
	http.Handle("/css/", http.FileServer(http.Dir(".")))
	http.Handle("/js/", http.FileServer(http.Dir(".")))
}


func root(w http.ResponseWriter, r *http.Request) {
	
	XiaoMiDevice := new(XiaoMiDevice);
	XiaoMiDevice.Note4G = xiaomiSearch("http://www.mi.com/sg/note4g/", "http://store.mi.com/sg/misc/getStarStock/hdid/note4g?jsonpcallback=jQuery18301541439404245466_1422453874281&_=1422453874342", w, r)
	XiaoMiDevice.Redmi2 = xiaomiSearch("http://www.mi.com/sg/redmi2/", "http://store.mi.com/sg/misc/getStarStock/hdid/redmi2?jsonpcallback=getStarStock&_=1427960799062", w, r)
	XiaoMiDevice.MiPad = xiaomiSearch("http://www.mi.com/sg/mipad/", "http://store.mi.com/sg/misc/getStarStock/hdid/mipad?jsonpcallback=getStarStock&_=1427960981544", w, r)
	XiaoMiDevice.Mipowerbank16000 = xiaomiSearch("http://www.mi.com/sg/mipowerbank16000/", "http://store.mi.com/sg/misc/getStarStock/hdid/power16000?jsonpcallback=jQuery18305076766561251134_1422460817629&_=1422460817709", w, r)
	XiaoMiDevice.Mipowerbank10400 = xiaomiSearch("http://www.mi.com/sg/mipowerbank10400/", "http://store.mi.com/sg/misc/getStarStock/hdid/power10400?jsonpcallback=jQuery183011226079403422773_1422460580239&_=1422460580339", w, r)
	XiaoMiDevice.Mipowerbank5000 = xiaomiSearch("http://www.mi.com/sg/mipowerbank5000/", "http://store.mi.com/sg/misc/getStarStock/hdid/power5000?jsonpcallback=getStarStock&_=1427987553449", w, r)
	XiaoMiDevice.Miband = xiaomiSearch("http://www.mi.com/sg/miband/", "http://store.mi.com/sg/misc/getStarStock/hdid/miband?jsonpcallback=jQuery183004000588273629546_1422460598809&_=1422460599020", w, r)
	
	xiaomiAvailForm.ExecuteTemplate(w, "xiaomiAvail.htm", XiaoMiDevice)
}

func xiaomiSearch(url string,jQueryUrl string, w http.ResponseWriter, r *http.Request) string {

	c := appengine.NewContext(r)
	client := urlfetch.Client(c)
    resp, err := client.Get(jQueryUrl)
    
    if err != nil {
        log.Fatal(err)
    }
	
	robots, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		log.Fatal(err)
	}
	re,_ := regexp.Compile("has_packet..false")
	match := re.Match(robots)
	if match {
		return url;
	} else {
		return "Not Available";
	}
	
}


var xiaomiAvailForm = template.Must(template.New("").ParseFiles("xiaomiAvail.htm"))