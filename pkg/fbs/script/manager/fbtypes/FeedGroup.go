// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package fbtypes

import "strconv"

/// type of the feedGroup
type FeedGroup int8

const (
	/// feedGroup x,y,z coordinate without orientation and without auxiliary axes
	FeedGroupFG_XYZ     FeedGroup = 0
	/// feedGroup x,y,z coordinate with orientation and without auxiliary axes
	FeedGroupFG_XYZ_O   FeedGroup = 1
	/// feedGroup x,y,z coordinate without orientation and with auxiliary axes
	FeedGroupFG_XYZ_A   FeedGroup = 2
	/// feedGroup x,y,z coordinate with orientation and with auxiliary axes
	FeedGroupFG_XYZ_O_A FeedGroup = 3
)

var EnumNamesFeedGroup = map[FeedGroup]string{
	FeedGroupFG_XYZ:     "FG_XYZ",
	FeedGroupFG_XYZ_O:   "FG_XYZ_O",
	FeedGroupFG_XYZ_A:   "FG_XYZ_A",
	FeedGroupFG_XYZ_O_A: "FG_XYZ_O_A",
}

var EnumValuesFeedGroup = map[string]FeedGroup{
	"FG_XYZ":     FeedGroupFG_XYZ,
	"FG_XYZ_O":   FeedGroupFG_XYZ_O,
	"FG_XYZ_A":   FeedGroupFG_XYZ_A,
	"FG_XYZ_O_A": FeedGroupFG_XYZ_O_A,
}

func (v FeedGroup) String() string {
	if s, ok := EnumNamesFeedGroup[v]; ok {
		return s
	}
	return "FeedGroup(" + strconv.FormatInt(int64(v), 10) + ")"
}
