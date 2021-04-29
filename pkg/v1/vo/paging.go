package vo

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"math"
	"reflect"
)

type PagingDataType struct {
	Slice interface{}
}

func NewPagingDataType(sliceType reflect.Type) (ret *PagingDataType) {
	emptySlice := reflect.MakeSlice(reflect.SliceOf(sliceType), 0, 0)
	data := reflect.New(emptySlice.Type())
	data.Elem().Set(emptySlice)
	return &PagingDataType{
		Slice: data.Interface(),
	}
}

type PagingArgs struct {
	Lookup      bson.M
	AddFields   bson.M
	Match       bson.M
	Sort        bson.M
	Page        int
	RowsPerPage int
	DataType    *PagingDataType
}

type PagingData struct {
	Slice       interface{} `json:"slice"`
	Count       int         `json:"count"`
	Page        int         `json:"page"`
	PageCount   int         `json:"pageCount"`
	RowsPerPage int         `json:"rowsPerPage"`
}

func NewPaging(col *mgo.Collection, args *PagingArgs) (ret *PagingData, err error) {
	ret = &PagingData{
		Page:        args.Page,
		RowsPerPage: args.RowsPerPage,
		Count:       0,
		PageCount:   1,
	}

	if ret.RowsPerPage == 0 {
		ret.RowsPerPage = 15
	}

	countQuery := append(basicQuery(args),
		bson.M{"$group": bson.M{"_id": "null", "count": bson.M{"$sum": 1}}},
		bson.M{"$project": bson.M{"_id": 0}},
	)

	var countMap map[string]int
	err = col.Pipe(countQuery).One(&countMap)
	if err != nil && err != mgo.ErrNotFound {
		return
	}

	if err != mgo.ErrNotFound {
		ret.Count = countMap["count"]
		if ret.Count > 0 {
			query := basicQuery(args)
			if args.Sort != nil {
				query = append(query, bson.M{"$sort": args.Sort})
			}

			if ret.Page == 0 {
				ret.Page = 1
				err = col.Pipe(query).All(args.DataType.Slice)
				if err != nil {
					return
				}
			} else {
				ret.PageCount = int(math.Ceil(float64(ret.Count) / float64(ret.RowsPerPage)))
				if ret.Page > ret.PageCount {
					ret.Page = ret.PageCount
				}
				query = append(query,
					bson.M{"$skip": (ret.Page - 1) * ret.RowsPerPage},
					bson.M{"$limit": ret.RowsPerPage},
				)
				err = col.Pipe(query).All(args.DataType.Slice)
				if err != nil {
					return
				}
			}
		}
	}
	ret.Slice = args.DataType.Slice
	return ret, nil
}

func basicQuery(args *PagingArgs) (ret []bson.M) {
	if args.Lookup != nil {
		ret = append(ret, bson.M{"$lookup": args.Lookup})
	}
	if args.AddFields != nil {
		ret = append(ret, bson.M{"$addFields": args.AddFields})
	}
	if args.Match != nil {
		ret = append(ret, bson.M{"$match": args.Match})
	}
	return
}
