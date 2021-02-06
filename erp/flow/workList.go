package flow

type workList struct{}

var WorkList workList

type Paras struct {
	WorkId int
}

func (fls *workList) Get(paras Paras) (workList []Work) {

	return

}
