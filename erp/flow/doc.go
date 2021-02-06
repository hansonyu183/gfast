package flow

type Doc struct{
	Id int
}

func(d *Doc)Get()(*Doc){

	return d
}