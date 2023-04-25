package model

import (
    "gorm.io/gorm"
    //"path"
    //"os"
    "github.com/Greetlist/CultureWeb/web_admin/server/model/schema"
    //"github.com/Greetlist/CultureWeb/web_admin/server/config"
    //"github.com/Greetlist/CultureWeb/web_admin/server/util"
    //uuid "github.com/nu7hatch/gouuid"
    //LOG "github.com/Greetlist/CultureWeb/web_admin/server/logger"
    ErrorCode "github.com/Greetlist/CultureWeb/web_admin/server/error"
)

type LabelTreeNode struct {
    LabelName string
    LabelID int
    Children []*LabelTreeNode
}

var LabelTree map[int]*LabelTreeNode

func (l *LabelModelStruct) InitTreeFromDB() (*ErrorCode.ResponseError) {
    LabelTree = make(map[int]*LabelTreeNode)
    var totalLabelList []schema.Label
    result := l.DB.Order("parent_id asc").Find(&totalLabelList)
    if result.Error != nil {
        return ErrorCode.DBQueryError
    }
    for _, item := range totalLabelList {
        curNode := new(LabelTreeNode)
        curNode.LabelName = item.LabelName
        curNode.LabelID = item.LabelID
        curNode.Children = make([]*LabelTreeNode, 5)
        LabelTree[item.LabelID] = curNode
        LabelTree[item.ParentID].Children = append(LabelTree[item.ParentID].Children, curNode)
    }
    return nil
}

func AddNode(newNode *LabelTreeNode) {
}

func DeleteNode(labelID int) {
}

type LabelModelStruct struct {
    DB *gorm.DB
}
