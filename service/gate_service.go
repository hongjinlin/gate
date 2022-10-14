/*
@Time : 2022/10/4 08:59
@Author : hongjinlin
@File : gate_service
@Software: GoLand
*/

package service

import (
	"gate/model"
	"gate/sys"
	"github.com/PuerkitoBio/goquery"
	"log"
	"net/http"
	"net/url"
	"time"
)

var GateUrl = "https://www.vpngate.net/cn/"

type GateService struct {
}

func init() {
	initTable()
}

func (g GateService) Gates() {
	gates := g.gatesFromNet()
	gateMgr := model.GateMgr(sys.DB)
	ips := make([]string, 0, len(gates))
	for _, gate := range gates {
		ips = append(ips, gate.IP)
	}
	existGates, err := gateMgr.GetBatchFromIP(ips)
	if err != nil {
		log.Printf("get ips data error: %s", err.Error())
	}
	existGatesMap := make(map[string]*model.Gate, len(existGates))
	updataGates := make([]*model.Gate, 0, len(existGates))
	saveGates := make([]*model.Gate, 0, len(existGates))

	for _, gate := range existGates {
		existGatesMap[gate.IP] = gate
	}
	loc, _ := time.LoadLocation("Asia/Shanghai")
	now := time.Now().In(loc)
	for i, gate := range gates {
		if g := existGatesMap[gate.IP]; g != nil {
			if gate.Runtime == g.Runtime {
				continue
			}
			gates[i].ID = g.ID
			gates[i].Runtime = gate.Runtime
			gates[i].UpdateTime = now
			updataGates = append(updataGates, gates[i])
			continue
		}
		saveGates = append(saveGates, gates[i])
	}
	gateMgr.CreateInBatches(saveGates, len(saveGates))
	for i := range updataGates {
		gateMgr := model.GateMgr(sys.DB)
		gateMgr.Model(updataGates[i]).
			Select(model.GateColumns.Runtime, model.GateColumns.UpdateTime).
			Updates(updataGates[i])
	}
}

func initTable() {
	if sys.DB.Migrator().HasTable(&model.Gate{}) {
		return
	}
	sys.DB.Migrator().CreateTable(&model.Gate{})
}

func (g GateService) gatesFromNet() (gates []*model.Gate) {
	viewstate, viewstategenerator, eventvalidation := g.token()
	resp, err := http.PostForm("https://www.vpngate.net/cn/", url.Values{
		"C_L2TP":               {"on"},
		"Button3":              {"更新服务器列表"},
		"__VIEWSTATE":          {viewstate},
		"__VIEWSTATEGENERATOR": {viewstategenerator},
		"__EVENTVALIDATION":    {eventvalidation},
	})
	defer func() {
		if resp.Body != nil {
			resp.Body.Close()
		}
	}()
	if err != nil {
		log.Fatalf("request error : %s", err.Error())
	}
	//body, err := io.ReadAll(resp.Body)
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	loc, _ := time.LoadLocation("Asia/Shanghai")
	now := time.Now().In(loc)
	doc.Find("table#vg_hosts_table_id tbody tr").Each(func(i int, s *goquery.Selection) {
		// For each item found, get the title
		children := s.Children()
		val, _ := children.Eq(0).Attr("class")
		if val == "vg_table_header" {
			return
		}
		if children.Size() != 10 {
			return
		}
		nation := children.Eq(0).Text()
		ip := children.Eq(1).Children().Eq(2).Text()
		runtime := children.Eq(2).Children().Eq(2).Text()
		gate := &model.Gate{
			Nation:     nation,
			IP:         ip,
			Runtime:    runtime,
			CreateTime: now,
			UpdateTime: now,
		}
		gates = append(gates, gate)
	})
	return
}

// 获取token
func (g GateService) token() (string, string, string) {
	resp, err := http.Get(GateUrl)
	defer func() {
		if resp.Body != nil {
			resp.Body.Close()
		}
	}()
	if err != nil {
		log.Fatalf("request error : %s", err.Error())
	}
	//body, err := io.ReadAll(resp.Body)

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		log.Fatalf("html read error: %s", err.Error())
	}

	return g.domValue(doc, "#__VIEWSTATE"), g.domValue(doc, "#__VIEWSTATEGENERATOR"), g.domValue(doc, "#__EVENTVALIDATION")
}

func (g GateService) domValue(d *goquery.Document, id string) (value string) {
	d.Find(id).Each(func(i int, s *goquery.Selection) {
		// For each item found, get the title
		value = s.Nodes[0].Attr[3].Val
	})
	return
}
