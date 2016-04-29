package main

import (
	"./proto"
	"bytes"
	"fmt"
	"github.com/codegangsta/martini-contrib/render"
	"github.com/go-martini/martini"
	"github.com/golang/protobuf/proto"
	"github.com/martini-contrib/gzip"
	"github.com/martini-contrib/sessions"
	"net/http"
	"strconv"
)

func NotSoClassic() *martini.ClassicMartini {
	r := martini.NewRouter()
	m := martini.New()
	//m.Use(logexcepted)
	m.Use(martini.Recovery())
	m.Use(martini.Static("static"))
	m.MapTo(r, (*martini.Routes)(nil))
	m.Action(r.Handle)
	return &martini.ClassicMartini{m, r}
}

// BOILERPLATE BOILERPLATE BOILERPLATE BOILERPLATE BOILERPLATE BOILERPLATE BOILERPLATE BOILERPLATE
//also boilerplate

func site_initialize() {

	m := NotSoClassic() //martini.Classic()

	store := sessions.NewCookieStore([]byte("jabberwocks, everywhere"))
	store.Options(sessions.Options{
		//Domain:   "localhost",
		Path:     "/",
		MaxAge:   0, //3600 * 8, // 8 hours
		HttpOnly: true,
	})
	session := sessions.Sessions("sess", store)

	m.Use(session)

	m.Use(gzip.All())

	m.Use(render.Renderer(render.Options{
		Layout: "layout", // Specify a layout template. Layouts can call {{ yield }} to render the current template.
	}))

	m.Get("/", DORFGUI)

	//m.Get("/inv", AJAX_GetInventory)
	//m.Post("/connect", AJAX_Connect)

	m.Post("/inv", AJAX_GetInventory)

	m.Post("/equip", AJAX_EquipInventoryItem)
	m.Post("/move", AJAX_MoveInventoryItem)

	m.RunOnAddr(":7777")
}

type pagedata struct {
}

func DORFGUI(r render.Render) {
	{
		r.HTML(200, "public_dash", "sup")
	}
}

var lastinventorystate []byte

func AJAX_GetInventory(r render.Render, req *http.Request) {

	inv_req := &dwarfgui.InventoryRequest{}
	inv_req.UnitId = proto.Int32(int32(0))
	//fmt.Println("wut do")
	reply, text, err := GetInventory(Remote, inv_req)

	//this is a hack. in the future, dwarfgui.plug.dll should do this, maybe.
	inventorystate, _ := proto.Marshal(reply)
	if reply.GetChanged() == false && lastinventorystate != nil && inventorystate != nil {
		if bytes.Equal(lastinventorystate, inventorystate) == false {
			reply.Changed = proto.Bool(true)
		} else {
			reply.Changed = proto.Bool(false)
		}
	} else {
		reply.Changed = proto.Bool(true)
	}
	if lastinventorystate == nil {
		reply.Changed = proto.Bool(false)
	}
	lastinventorystate = inventorystate

	if err != nil {
		fmt.Println("err!!!: ", err)
		fmt.Println("text: ", text)
	} else {
		//fmt.Println(text)
		//fmt.Println("%v", len(reply.GetItem()))

		if err != nil {
			fmt.Println("err: ", err)
		} else {
			r.JSON(200, reply)
		}
	}
}

func AJAX_EquipInventoryItem(r render.Render, req *http.Request) {
	arg := req.FormValue("item_id")
	item_id, _ := strconv.Atoi(arg)

	inv_req := &dwarfgui.InventoryEquipRequest{}
	inv_req.ItemId = proto.Int32(int32(item_id))
	reply, text, err := EquipInventoryItem(Remote, inv_req)
	if err != nil {
		fmt.Println("err: ", err)
		fmt.Println("text: ", text)
	} else {
		//fmt.Println(text)
		r.JSON(200, reply)
	}
}

func AJAX_MoveInventoryItem(r render.Render, req *http.Request) {

	arg := req.FormValue("item_id")
	item_id, _ := strconv.Atoi(arg)
	arg = req.FormValue("dest_id")
	dest_id, _ := strconv.Atoi(arg)

	inv_req := &dwarfgui.InventoryMoveRequest{}
	inv_req.ItemId = proto.Int32(int32(item_id))
	inv_req.DestinationId = proto.Int32(int32(dest_id))

	reply, text, err := MoveInventoryItem(Remote, inv_req)
	if err != nil {
		fmt.Println("err: ", err)
		fmt.Println("text: ", text)
	} else {
		//fmt.Println(text)
		r.JSON(200, reply)
	}
}
