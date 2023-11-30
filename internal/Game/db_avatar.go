package Game

import (
	"time"

	"github.com/gucooing/hkrpg-go/gdconf"
	"github.com/gucooing/hkrpg-go/protocol/proto"
)

type DbAvatar struct {
	Avatar map[uint32]*Avatar
}

type Avatar struct {
	AvatarId          uint32            // 角色id
	Exp               uint32            // 经验
	Level             uint32            // 等级
	FirstMetTimestamp uint64            // 获得时间戳
	Promotion         uint32            // 突破等阶
	Rank              uint32            // 命座
	Hp                uint32            // 血量
	SkilltreeList     map[uint32]uint32 `json:"-"` // 技能等级数据
}

func AddAvatar(avatarId uint32) *Avatar {
	avatar := new(Avatar)
	// TODO
	avatar.AvatarId = avatarId
	avatar.Exp = 0
	avatar.Level = 1
	avatar.FirstMetTimestamp = uint64(time.Now().UnixNano())
	avatar.Promotion = 0
	avatar.Rank = 0
	avatar.Hp = 10000
	return avatar
}
func GetKilltreeList(avatarId string) []*proto.AvatarSkillTree {
	avatarData := gdconf.GetAvatarDataById(avatarId)
	SkilltreeList := make([]*proto.AvatarSkillTree, 0)
	for _, a := range avatarData.SkillList {
		skilltreeList := &proto.AvatarSkillTree{
			PointId: a,
			Level:   1,
		}
		SkilltreeList = append(SkilltreeList, skilltreeList)
	}
	return SkilltreeList
}