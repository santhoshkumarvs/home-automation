package handler

import (
	"github.com/jakewright/home-automation/libraries/go/database"
	"github.com/jakewright/home-automation/libraries/go/errors"
	"github.com/jakewright/home-automation/service.scene/domain"
	sceneproto "github.com/jakewright/home-automation/service.scene/proto"
)

// HandleReadScene returns the scene with the given ID
func HandleReadScene(req *sceneproto.ReadSceneRequest) (*sceneproto.ReadSceneResponse, error) {
	scene := &domain.Scene{}
	if err := database.Find(&scene, req.SceneId); err != nil {
		return nil, errors.WithMessage(err, "failed to find")
	}

	return &sceneproto.ReadSceneResponse{
		Scene: scene.ToProto(),
	}, nil
}
