package packet

import (
	neteasePacket "Eulogist/core/minecraft/protocol/packet"

	standardPacket "github.com/sandertv/gophertunnel/minecraft/protocol/packet"
)

type Text struct{}

func (pk *Text) ToNetNetEasePacket(standard standardPacket.Packet) neteasePacket.Packet {
	p := neteasePacket.Text{}
	input := standard.(*standardPacket.Text)

	p.TextType = input.TextType
	p.NeedsTranslation = input.NeedsTranslation
	p.SourceName = input.SourceName
	p.Message = input.Message
	p.Parameters = input.Parameters
	p.XUID = input.XUID
	p.PlatformChatID = input.PlatformChatID

	p.NeteaseExtraData = make([]string, 0)
	p.Unknown = ""

	return &p
}

func (pk *Text) ToStandardPacket(netease neteasePacket.Packet) standardPacket.Packet {
	p := standardPacket.Text{}
	input := netease.(*neteasePacket.Text)

	p.TextType = input.TextType
	p.NeedsTranslation = input.NeedsTranslation
	p.SourceName = input.SourceName
	p.Message = input.Message
	p.Parameters = input.Parameters
	p.XUID = input.XUID
	p.PlatformChatID = input.PlatformChatID

	return &p
}