package main

import (
	"./proto"
	"fmt"
	//"github.com/BenLubar/arm_ok/armok_web/proxy"
	"github.com/BenLubar/arm_ok/dfhack"
	"github.com/BenLubar/arm_ok/dfhack/dfproto"
)

var pluginDwarfGUI = "dwarfgui"

func GetInventory(c *dfhack.Conn, req *dwarfgui.InventoryRequest) (*dwarfgui.InventoryReply, []*dfproto.CoreTextNotification, error) {
	var reply dwarfgui.InventoryReply
	text, err := c.RoundTripBind("GetInventory", &pluginDwarfGUI, "dwarfgui.InventoryRequest", "dwarfgui.InventoryReply", req, &reply)
	return &reply, text, err
}

func EquipInventoryItem(c *dfhack.Conn, req *dwarfgui.InventoryEquipRequest) (*dwarfgui.InventoryEquipReply, []*dfproto.CoreTextNotification, error) {
	var reply dwarfgui.InventoryEquipReply
	text, err := c.RoundTripBind("EquipInventoryItem", &pluginDwarfGUI, "dwarfgui.InventoryEquipRequest", "dwarfgui.InventoryEquipReply", req, &reply)
	return &reply, text, err
}

func MoveInventoryItem(c *dfhack.Conn, req *dwarfgui.InventoryMoveRequest) (*dwarfgui.InventoryMoveReply, []*dfproto.CoreTextNotification, error) {
	var reply dwarfgui.InventoryMoveReply
	text, err := c.RoundTripBind("MoveInventoryItem", &pluginDwarfGUI, "dwarfgui.InventoryMoveRequest", "dwarfgui.InventoryMoveReply", req, &reply)
	return &reply, text, err
}

func derp() {

	remote()

	req := &dwarfgui.InventoryRequest{}

	reply, text, err := GetInventory(Remote, req)

	if err != nil {
		fmt.Println("err: ", err)
		fmt.Println("text: ", text)
	} else {
		fmt.Println(text)
		fmt.Println("%v", len(reply.GetItem()))
	}
}
