package process

type UserManager struct {
	userList map[int]*UserProcessor
}

var Umg *UserManager

func InitUserManager(){
	Umg = &UserManager{
		userList: make(map[int]*UserProcessor,1024),
	}
}

func (this *UserManager)AddUserOnline(userId int,up *UserProcessor){
	this.userList[userId] = up
}

func (this *UserManager) DelUserOnline(userId int) {
	delete(this.userList,userId)
}

func (this *UserManager) GetAllOnlineUser() map[int]*UserProcessor {
	return this.userList
}
