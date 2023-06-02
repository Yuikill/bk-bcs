package project

import (
	"context"

	"github.com/Tencent/bk-bcs/bcs-common/pkg/odm/operator"
	"github.com/Tencent/bk-bcs/bcs-services/bcs-project-manager/internal/common/page"
	"github.com/Tencent/bk-bcs/bcs-services/bcs-project-manager/internal/component/bkmonitor"
	"github.com/Tencent/bk-bcs/bcs-services/bcs-project-manager/internal/store"
	pm "github.com/Tencent/bk-bcs/bcs-services/bcs-project-manager/internal/store/project"
	"github.com/Tencent/bk-bcs/bcs-services/bcs-project-manager/internal/util/errorx"
	proto "github.com/Tencent/bk-bcs/bcs-services/bcs-project-manager/proto/bcsproject"
)

// ListForIAMAction xxx
type ListForIAMAction struct {
	ctx   context.Context
	model store.ProjectModel
	req   *proto.ListProjectsForIAMReq
}

// NewListForIAMActionAction new list projects for iam action
func NewListForIAMActionAction(model store.ProjectModel) *ListForIAMAction {
	return &ListForIAMAction{
		model: model,
	}
}

// Do xxx
func (la *ListForIAMAction) Do(ctx context.Context, req *proto.ListProjectsForIAMReq) (
	[]*proto.ListProjectsForIAMResp_Project, error) {
	la.ctx = ctx
	la.req = req

	projects, _, err := la.listProjects()
	if err != nil {
		return nil, errorx.NewDBErr(err.Error())
	}
	spaces, err := bkmonitor.ListSpaces()
	if err != nil {
		return nil, err
	}
	spaceMap := make(map[string]*bkmonitor.Space)
	for _, space := range spaces {
		spaceMap[space.SpaceCode] = space
	}

	filtered := []*proto.ListProjectsForIAMResp_Project{}
	for _, project := range projects {
		if space, ok := spaceMap[project.ProjectID]; ok {
			filtered = append(filtered, &proto.ListProjectsForIAMResp_Project{
				Name:        project.Name,
				ProjectID:   project.ProjectID,
				ProjectCode: project.ProjectCode,
				BusinessID:  project.BusinessID,
				Managers:    project.Managers,
				// 监控BCS项目空间ID注册到权限中心的资源ID需要取负数
				BkmSpaceBizID: int32(-space.ID),
				BkmSpaceName:  space.DisplayName,
			})
		}
	}
	return filtered, nil
}

func (la *ListForIAMAction) listProjects() ([]*pm.Project, int64, error) {

	cond := operator.NewLeafCondition(operator.Eq, operator.M{
		"kind": "k8s",
	})
	// 查询所有开启了容器服务的项目
	projects, total, err := la.model.ListProjects(la.ctx, cond, &page.Pagination{All: true})
	if err != nil {
		return nil, total, err
	}
	projectList := []*pm.Project{}
	for i := range projects {
		projectList = append(projectList, &projects[i])
	}
	return projectList, total, nil
}
