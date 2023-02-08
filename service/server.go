package service

import (
	"context"
	"v3board/global"
	"v3board/models/db"
	"v3board/utils"
)

type serverService struct{}

var Server = serverService{}

// StateUpdate 统计信息更新
func (*serverService) StateUpdate(ctx context.Context, user_id, node_id int64, node_type string, up, down int64) error {
	tempNode := &db.StateNode{RecordAt: utils.TodayStartTime(), RecordType: "d"}
	has, _ := global.DB.Get(tempNode)

	session := global.DB.NewSession()
	defer session.Close()

	if err := session.Begin(); err != nil {
		return err
	}

	if !has {
		stateNode := db.StateNode{
			NodeId:     node_id,
			NodeType:   node_type,
			Up:         up,
			Down:       down,
			RecordType: "d",
			RecordAt:   utils.TodayStartTime(),
		}
		if _, err := session.Insert(&stateNode); err != nil {
			return err
		}

		stateUser := db.StateUser{
			UserId:     user_id,
			Up:         up,
			Down:       down,
			RecordType: "d",
			RecordAt:   utils.TodayStartTime(),
		}
		if _, err := session.Insert(&stateUser); err != nil {
			return err
		}

		return session.Commit()
	}

	if _, err := session.Incr("up", up).Incr("down", down).Update(tempNode); err != nil {
		return err
	}

	tempUser := &db.StateUser{RecordAt: utils.TodayStartTime(), RecordType: "d"}
	if _, err := session.Incr("up", up).Incr("down", down).Update(tempUser); err != nil {
		return err
	}

	return session.Commit()
}
