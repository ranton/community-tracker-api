package community

import (
	"github.com/gofiber/fiber/v2"
	"github.com/VncntDzn/community-tracker-api/pkg/common/models"
)

func (h handler) GetCommunityWithmembersPercentage(c *fiber.Ctx) error {
	var community_data []models.CommunityWithMembersPercentage
	sub := h.DB.Table("people").Select("count(people.peopleid)*100/(select count(peopleid) from people where isactive = true) as percentage, community.communityname, community.communitydesc, community.communityicon, community.communityid, community.communitymgrid").Joins("right join community on people.communityid = community.communityid and people.isactive = true").Where("community.isactive = true").Group("community.communityid").Order("lower(community.communityname)")
    if result := h.DB.Table("communityadminandmanager").Select("sub.percentage, sub.communityname, sub.communitydesc,sub.communityicon, sub.communityid, communityadminandmanager.communityadminandmanagername").Joins("right join (?) as sub on communityadminandmanager.communityadminandmanagerid = sub.communitymgrid", sub).Scan(&community_data); result.Error != nil {
        return fiber.NewError(fiber.StatusNotFound, result.Error.Error())
    }

	return c.Status(fiber.StatusOK).JSON(&community_data)
}