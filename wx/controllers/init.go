package controllers

import (
	"reflect"
	"strings"
	"time"

	"edwardhey.com/football/wx/models"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/toolbox"
)

type BaseController struct {
	beego.Controller `json:"-"`
	IsJSON           bool                   `json:"-"`
	Code             int                    `json:"code"`
	Err              string                 `json:"err,omiempty"`
	JsonData         map[string]interface{} `json:"data"`
}

func (c *BaseController) Prepare() {
	// if c.IsJSON {
	c.JsonData = make(map[string]interface{}, 20)
	c.Data["title"] = ""
	// }
}

func (c *BaseController) Render() error {

	// return nil
	if c.IsJSON {
		c.Data["json"] = c
		c.Ctx.Output.SetStatus(200)
		c.ServeJSON()
		return nil
	}

	c.Layout = "layout.html"
	return c.Controller.Render()
}

func (c *BaseController) ThrowErr(msg string) {
	// c.Ctx.WriteString(msg)
	// log.Fatalln(msg)
	c.CustomAbort(500, msg)
}

func init() {
	beego.AddFuncMap("GetActivityStatusString", models.GetActivityStatusString)
	beego.AddFuncMap("Date", Date)
	beego.AddFuncMap("CallMethod", CallMethod)
	beego.SetStaticPath("/semantic", "static/semantic")
	tk := toolbox.NewTask("createActivity", "0 19 0 * * 3", func() error {
		//创建报名活动，并推送到微信用户
		return nil
	})
	tk.Run()

}

type AuthorizedController struct {
	BaseController
}

func (c *AuthorizedController) Prepare() {
	c.StartSession()
	//验证信息
	if c.GetSession("player") == nil {
		//TODO：上线要改成auth
		c.Redirect("/wx/callback/?code=041LPhzd2bPOrA0T8exd2Zdhzd2LPhzy&state=", 302)
	}
	// fmt.Println(c.GetSession("player"))
	// user := models.GetPlayerWithOpenID(c.Op)
}

//----------------funcs----------------------------------
func GetRequestURIWithUrlFor(path string) string {
	port := beego.AppConfig.String("httpport")
	if port == "80" {
		return beego.AppConfig.String("hostname") + beego.URLFor(path)
	}
	return beego.AppConfig.String("hostname") + ":" + beego.AppConfig.String("httpport") + beego.URLFor(path)
}

//-------------tpl funcs-------------------------------
// DateFormat pattern rules.
var datePatterns = []string{
	// year
	"Y", "2006", // A full numeric representation of a year, 4 digits   Examples: 1999 or 2003
	"y", "06", //A two digit representation of a year   Examples: 99 or 03

	// month
	"m", "01", // Numeric representation of a month, with leading zeros 01 through 12
	"n", "1", // Numeric representation of a month, without leading zeros   1 through 12
	"M", "Jan", // A short textual representation of a month, three letters Jan through Dec
	"F", "January", // A full textual representation of a month, such as January or March   January through December

	// day
	"d", "02", // Day of the month, 2 digits with leading zeros 01 to 31
	"j", "2", // Day of the month without leading zeros 1 to 31

	// week
	"D", "Mon", // A textual representation of a day, three letters Mon through Sun
	"l", "Monday", // A full textual representation of the day of the week  Sunday through Saturday

	// time
	"g", "3", // 12-hour format of an hour without leading zeros    1 through 12
	"G", "15", // 24-hour format of an hour without leading zeros   0 through 23
	"h", "03", // 12-hour format of an hour with leading zeros  01 through 12
	"H", "15", // 24-hour format of an hour with leading zeros  00 through 23

	"a", "pm", // Lowercase Ante meridiem and Post meridiem am or pm
	"A", "PM", // Uppercase Ante meridiem and Post meridiem AM or PM

	"i", "04", // Minutes with leading zeros    00 to 59
	"s", "05", // Seconds, with leading zeros   00 through 59

	// time zone
	"T", "MST",
	"P", "-07:00",
	"O", "-0700",

	// RFC 2822
	"r", time.RFC1123Z,
}

// Date takes a PHP like date func to Go's time format.
func Date(t uint32, format string) string {
	replacer := strings.NewReplacer(datePatterns...)
	format = replacer.Replace(format)
	_t := time.Unix(int64(t), 0)
	return _t.Format(format)
}

func CallMethod(obj interface{}, methodStr string) string {
	return reflect.ValueOf(obj).MethodByName(methodStr).Call(nil)[0].String()
}
