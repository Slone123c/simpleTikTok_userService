package ApiServer

import (
	"github.com/cloudwego/hertz/pkg/app/server"
	ApiServer "github.com/cold-runner/simpleTikTok/apiServer/biz/handler/ApiServer"
)

/*
 This file will register all the routes of the services in the master idl.
 And it will update automatically when you use the "update" command for the idl.
 So don't modify the contents of the file, or your code will be deleted when it is updated.
*/

// Register register routes based on the IDL 'api.${HTTP Method}' annotation.
func Register(r *server.Hertz) {

	root := r.Group("/", rootMw()...)
	{
		_douyin := root.Group("/douyin", _douyinMw()...)
		{
			_comment := _douyin.Group("/message", _commentMw()...)
			{
				_action := _comment.Group("/action", _actionMw()...)
				_action.POST("/", append(_commentactionMw(), ApiServer.CommentAction)...)
			}
			{
				_list := _comment.Group("/list", _listMw()...)
				_list.GET("/", append(_commentlistMw(), ApiServer.CommentList)...)
			}
		}
		{
			_favorite := _douyin.Group("/favorite", _favoriteMw()...)
			{
				_action0 := _favorite.Group("/action", _action0Mw()...)
				_action0.POST("/", append(_favoriteactionMw(), ApiServer.FavoriteAction)...)
			}
			{
				_list0 := _favorite.Group("/list", _list0Mw()...)
				_list0.GET("/", append(_favoritelistMw(), ApiServer.FavoriteList)...)
			}
		}
		{
			_feed := _douyin.Group("/feed", _feedMw()...)
			_feed.GET("/", append(_feed0Mw(), ApiServer.Feed)...)
		}
		{
			_message := _douyin.Group("/message", _messageMw()...)
			_message.POST("/action", append(_messageactionMw(), ApiServer.MessageAction)...)
			_message.GET("/chat", append(_messagechatMw(), ApiServer.MessageChat)...)
		}
		{
			_publish := _douyin.Group("/publish", _publishMw()...)
			{
				_action1 := _publish.Group("/action", _action1Mw()...)
				_action1.POST("/", append(_publishactionMw(), ApiServer.PublishAction)...)
			}
			{
				_list1 := _publish.Group("/list", _list1Mw()...)
				_list1.GET("/", append(_publishlistMw(), ApiServer.PublishList)...)
			}
		}
		{
			_relation := _douyin.Group("/relation", _relationMw()...)
			{
				_action2 := _relation.Group("/action", _action2Mw()...)
				_action2.POST("/", append(_relationactionMw(), ApiServer.RelationAction)...)
			}
			{
				_follow := _relation.Group("/follow", _followMw()...)
				{
					_list2 := _follow.Group("/list", _list2Mw()...)
					_list2.GET("/", append(_relationfollowlistMw(), ApiServer.RelationFollowList)...)
				}
			}
			{
				_follower := _relation.Group("/follower", _followerMw()...)
				{
					_list3 := _follower.Group("/list", _list3Mw()...)
					_list3.GET("/", append(_relationfollowerlistMw(), ApiServer.RelationFollowerList)...)
				}
			}
			{
				_friend := _relation.Group("/friend", _friendMw()...)
				_friend.GET("/list", append(_friendlistMw(), ApiServer.FriendList)...)
			}
		}
		{
			_user := _douyin.Group("/user", _userMw()...)
			_user.GET("/", append(_getuserinfoMw(), ApiServer.GetUserInfo)...)
			{
				_login := _user.Group("/login", _loginMw()...)
				_login.POST("/", append(_login0Mw(), ApiServer.Login)...)
			}
			{
				_register := _user.Group("/register", _registerMw()...)
				_register.POST("/", append(_register0Mw(), ApiServer.Register)...)
			}
		}
	}
}
