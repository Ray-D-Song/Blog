package service

import (
	"context"
	"server/db"
	"server/ent"
	"server/ent/blog"
)

func GetBlogList(ctx context.Context) ([]*ent.Blog, error) {
	list, err := db.EC.Blog.Query().Select().All(ctx)
	if err != nil {
		return nil, err
	}
	return list, nil
}

func GetBlog(ctx context.Context, id int) (*ent.Blog, error) {
	val, err := db.EC.Blog.Query().Where(blog.ID(id)).First(ctx)
	if err != nil {
		return nil, err
	}
	return val, nil
}

//func AddBlog(ctx context.Context, b ent.Blog) (*ent.Blog, error) {
//	val, err := db.EC.Blog.Create().SetHits(0).SetCover(b.Cover).SetSubtitle(b.Subtitle).SetTime(strconv.FormatInt(time.Now().Unix(), 10)).SetContentURL(b.ContentURL).SetTitle(b.Title).Save(ctx)
//	if err != nil {
//		return nil, err
//	}
//	return val, nil
//}

func AddHit(ctx context.Context, id int) (*ent.Blog, error) {
	val, err := db.EC.Blog.UpdateOneID(id).AddHits(1).Save(ctx)
	if err != nil {
		return nil, err
	}
	return val, nil
}
