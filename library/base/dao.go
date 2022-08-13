package base

//
//type Id interface {
//	string | int | int32 | int64
//}
//
//type Do[I Id] interface {
//}

//type IExtendDao[T Do[K], K Id] interface {
//	SelectBatchIds(ctx context.Context, idList []K) ([]*T[K], error)
//	DeleteByMap(ctx context.Context, params map[string]any) (int, error)
//	SelectMaps(ctx context.Context, do T[K]) (map[string]any, error)
//	SelectMapsPage(ctx context.Context, page IPage, do T[K]) (*IPageInfo[map[string]any], error)
//	SelectByMap(ctx context.Context, columnMap map[string]any) ([]*T[K], error)
//	SelectObjs(ctx context.Context, do T[K]) ([]any, error)
//}
//
//type IDao[T Do[K], K Id] interface {
//	Save(ctx context.Context, do T[K]) (int, error)
//	Delete(ctx context.Context, do T[K]) (int, error)
//	DeleteById(ctx context.Context, id K) (int, error)
//	Update(ctx context.Context, do T[K]) (int, error)
//	UpdateById(ctx context.Context, do T[K]) (int, error)
//	DeleteBatchIds(ctx context.Context, idList []K) (int, error)
//	SaveOrUpdate(ctx context.Context, do T[K]) (int, error)
//	SelectById(ctx context.Context, id K) (*T[K], error)
//	SelectOne(ctx context.Context, do T[K]) (*T[K], error)
//	SelectList(ctx context.Context, do T[K]) ([]*T[K], error)
//	SelectCount(ctx context.Context, do T[K]) (int, error)
//	SelectPage(ctx context.Context, page IPage, do T[K]) (*IPageInfo[T[K]], error)
//	Exists(ctx context.Context, do T[K]) (bool, error)
//}

//type Dao[T Do[K], K Id] struct {
//}
