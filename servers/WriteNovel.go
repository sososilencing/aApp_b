package servers

import (
	"exchange_nolves/model"
)

const NovelWords = 2410

//写题目之创建者
func WriteMNovel(novel model.MNovels) *model.MNovels {
	if NovelFilter(novel.Context,novel.Uid){
		novel = InsertMNovel(novel)
		return &novel
	}
	return nil
}
func NovelFilter(context string, uid int) bool {
	if len(context) > NovelWords {
		return false
	}
	if  uid == 0{
		return false
	}
	return true
}

func RandOne() *model.MNovels {
	novel := RandNovel()
	if novel.Uid !=0 {
		return &novel
	}
	return nil
}

func above(novels []model.MNovels) *model.OneNovel {
	lenth := len(novels)
	Mnovel := &model.OneNovel{}
	if lenth == 0{
		return nil
	}
	if lenth == 1{
		Mnovel.MNovel = novels[lenth-1]
		return Mnovel
	}
	if lenth >= 2{
		One := &model.OneNovel{
			MNovel:   novels[lenth-2],
			OneNovel: nil,
		}
		Mnovel.OneNovel = One
		for i:= lenth-3; i>=0;i--{
			Two := &model.OneNovel{
				MNovel:   novels[i],
				OneNovel: nil,
			}
			One.OneNovel = Two
			One = Two
		}
	}

	return Mnovel
}
//一条故事线
func FindOneNovel(novel model.MNovels) (*model.OneNovel, *model.Novels) {

	//上面的章节
	Pnovel := SelectOneNovel(novel.Pid)

	novels := make([]model.MNovels,0)

	if novel.Pid!=0{
		novels = append(novels,*Pnovel)
		for Pnovel.Pid!=0{
			Pnovel = SelectOneNovel(Pnovel.Pid)
			novels = append(novels,*Pnovel)
		}
	}
	//fmt.Println(novels)
	Mnovel := above(novels)

	//找啊找下面的章节
	Cnovels := SelectChildrenNovel(novel.Nid)
	page := model.Novels{
		Nid:            novel.Nid,
		Uid:            novel.Uid,
		Title:          novel.Title,
		Context:        "",
		Pid:            novel.Pid,
		ChildrenNovels: Cnovels,
	}
	findCircle(Cnovels)
	return Mnovel,&page
}


func findCircle(novels []*model.Novels)  {
	for i:=0;i<len(novels);i++ {
		Cnovels := SelectChildrenNovel(novels[i].Nid)
		findCircle(Cnovels)
		novels[i].ChildrenNovels = Cnovels
	}
}

