package service

import (
	"context"
	"server/db"
	"server/ent"
)

func GetRemarkList(ctx context.Context) ([]*ent.Remark, error) {
	list, err := db.EC.Remark.Query().Select().Where().All(ctx)
	if err != nil {
		return nil, err
	}
	return list, nil
}

func AddRemark(ctx context.Context, r ent.Remark) (*ent.Remark, error) {
	val, err := db.EC.Remark.Create().SetContent(r.Content).SetBlogID(r.BlogID).SetTargetID(r.TargetID).SetIsReply(r.IsReply).SetTargetName(r.TargetName).SetContent(r.CreateAt).SetCreateAt(r.CreateAt).SetUserName(r.UserName).Save(ctx)
	if err != nil {
		return nil, err
	}
	return val, err
}

func DelRemark(ctx context.Context, id int) (*ent.Remark, error) {
	val, err := db.EC.Remark.UpdateOneID(id).SetIfDel(true).Save(ctx)
	if err != nil {
		return nil, err
	}
	return val, nil
}
