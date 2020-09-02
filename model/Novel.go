package model

type Comment struct {
	Id int
	Uid int
	Pid int
	Comment string
}

type MNovels struct {
	Nid int
	Uid int
	Title string
	Context string
	Pid int
}

type Novels struct {
	Nid int
	Uid int
	Title string
	Context string
	Pid int
	ChildrenNovels []*Novels
}

type OneNovel struct {
	MNovel MNovels
	OneNovel *OneNovel
}

type Words struct {
	Id int
	Words string
}

type One struct {
	OneNovel *OneNovel
	Novels *Novels
}
