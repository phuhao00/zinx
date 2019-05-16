package conf

import "sync"

type User struct {
	id         int64 //用户ID
	account    string//登录的账号
	token      string
	username   string//
	nick       string
	lv         int32
	vipLv      int32
	loginTime  int64
	logoutTime int64
	mgrModNum  int32

	channel  string
	platform string
	login    int32

	mountId       int32
	mountLv       int32
	mountExp      int32
	petIllOpenNum int32
	dressNum      int32
	//baseData   *UserBaseData
	//battleData *UserBattleData
	//bag        *UserBagData
	incItemId  int32

	maxPetId int32
	/*pets 		map[int32]*Pet
	sortPowerPets []*Pet*/
	petProto2Count map[int32]int32

	//allItems map[int32]*Item //不考虑堆叠物品
	//
	//battlePetCon *PetBattleContainer
	//defPetCon	*PetBattleContainer

	followPetId      int32
	followPetProtoId int32

	//pos 		Vector3
	isOnline bool //是否在线

	unloadCheckTime int64
	//session         *Session
	//
	//captureData *UserCaptureData
	//petIllData  *UserPetIllData
	//friendData  *UserFriendData
	//titleData   *UserTitleData
	//
	//Ch_Packet      chan *ClientPacket
	//mountContainer *MountContainer
	//
	//gashaponData *GashaponData

	equipFashionDresses     []int32
	initEquipFashionDresses []int32
	//fashionDresses          map[int32]*FashionDress
	//fashionDressAddProp     *Config.PropAddData
	//
	//decorateMap     map[int32](map[int32]*Decorate)
	//decorateAddProp *Config.PropAddData
	//
	//taskStates        map[int32]*TaskState
	//taskStatesByPlot  map[int32]*TaskState
	//taskPlotStates    map[int32]*TaskPlotState
	//taskRingState     *TaskRingState
	//taskDailyState    *TaskDailyState
	//taskSecurityState *TaskSecurityState
	//
	//mails         map[int32]*Mail
	maxMailId     int32
	mailTempDatas map[int64]bool

	//delayExitGrids mapset.Set //TODO:延迟删除功能后面实现
	//fightData *UserFightData
	//fbRecord  bitarray.BitArray
	//
	//probe *Probe
	//
	//fbData      *FbUserData
	//arenaData   *ArenaUserData
	//welfareData *WelFareData

	power                 int64
	maxPetPower           int64
	monsterProtectEndTime int64

	baseUserModMark int32
	complexModTime  map[int32]int64
	teamId          int32
	//fbRoomId 	int64
	//PacketsQueue queue.Queue
	zoneId          int32 //区域id, 纯粹由客户端发协议来设置该值
	triggeredEvents map[int32]bool

	//rpcClient    *RpcClient
	//rpcServer    *RpcServer
	chCloseRecv  chan chan bool
	//recvReqQueue []*ChRecvReq
	mutexReq     sync.Mutex
	isClosing    bool

	//globalMapPos     Vector3
	globalRideMount  int32
	//mysteryStoreData *MysteryStoreData //神秘商店
	//hangupData       *HangupData       //离线收益
	//practiceData     *UserPetPracticeData

	//成就需要维护的数据----------------
	//taskFinishedTotalNum int32 //完成的任务总数
	//identifyTimes int32   //鉴定宠物次数
	//arenaWinTimes int32 	   //竞技场赢的次数
	//wildPVEArenaWinTimes int32 //野外PVE战斗胜利次数
	arenaRank        int32 //竞技场最高排名
	petBattleTotalLv int32 //未存储
	//hadWentTown      *HadWentTown

	headFrameQuantity     int64 //头像框数量
	trackQuantity         int64 //足迹数量
	dialogueFrameQuantity int64 //对话框的数量
	teamLogoQuantity      int64 // 队标的数量

	petEvo2Count          map[int32]int32 //宠物进化阶数对应的数量
	petStage2Count        map[int32]int32 //阶数X宠物数量
	petLv2Count           map[int32]int32 //等级X宠物数量
	petScoreQuality2Count map[int32]int32 //资质X宠物数量
	petPower2Count        map[int32]int32 //超能力X宠物数量
	petTotalStar2Count    map[int32]int32 //阶数*100+星数X宠物数量

	petSkillsLvSum2Count   map[int32]int32           //宠物的技能总等级对应的数量
	petMcSkillsLvSum2Count map[int32]int32           //宠物的秘传技能总等级对应的数量
	petPracticePoint2Count map[int32]int32           //宠物修行点数对应的数量
	zCrystalLv2Star2Count  map[int32]map[int32]int32 //获得的结晶等级对应的星数对应的数两
	friendFavor2Count      map[int32]int32           //好友好感值对应的数量
	//challengeData          *NpcChallengeData         //NPC挑战数据
	////头衔
	//honourAddProp         *Config.PropAddData
	currentHonourLevel    int32   //当前的头衔等级
	NeverReceiveHonourIds []int32 //已完成未领取福利的头衔Id
	//

	pkValue       int32
	isInPvp       bool
	pkProtectTime int64
	//roadHallData  *UserRoadHallData
	//
	//familyData *FamilyData //家族数据
	////帧处理所有的任务状态，统一发包
	//frameChangedTaskStates map[int32]*TaskState
	//frameRemovedTaskStates map[int32]*TaskState
	////任务进度数据
	//TasksHonourStateCache      map[int32]*TaskState
	//TasksAchievementStateCache map[int32]*TaskState
	//TasksRechargeStateCache    map[int32]*TaskState

	//新手引导
	finishedGuideIds []int32
	//充值记录，首充为true，否则False
	RechargeRecords map[int32]bool
	//救援CD
	helpTransferCD int32
	//充值卡需要为护的数据
	//rechargeCardData *RechargeCardData
	//RechargeLimitData *RechargeLimit //每日限购，限领
	//exChangeData *ExChangeUserData
	coldCurrencys map[int32]int64
	//阵营数据
	//campData *UserCampData
	//
	Tasks []*Task
}
//