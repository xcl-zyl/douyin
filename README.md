### Go学习
#### 学习资料
1. [Go 语言上手 - 基础语言](https://juejin.cn/post/7093721879462019102)
2. [Go语言圣经](https://books.studygolang.com/gopl-zh/)
3. [青训营完整手册](https://bytedance.feishu.cn/docs/doccnFRB1TXYJPK6yprPETHLXgd#q8ZYps)
4. [抖音项目-青训营](https://bytedance.feishu.cn/docx/doxcnbgkMy2J0Y3E6ihqrvtHXPg)
5. [Go语言上手PPT](https://bytedance.feishu.cn/file/boxcnQnHXuDOdzd8CqVid7nQLmg)
6. [课后项目实践作业](https://juejin.cn/post/7094452391101071367/)
7. [Go语言工程实践PPT](https://bytedance.feishu.cn/file/boxcnRmlw9MjbtAMBnOW44y8dZd?hash=7cfc75acc80372c08463b622df90a4b5)
8. [高质量与性能调优PPT](https://bytedance.feishu.cn/file/boxcnqqWtT0xgWAIMGWVs7wM6fd?hash=ab6bfba21a54c52073c7341ecb3ab470)

#### 语法学习


#### 实践
1. 视频Feed流、视频投稿、个人信息：   
&emsp;&emsp;支持所有用户刷抖音，按投稿时间倒叙推出，登录用户可自己拍视频投稿，查看自己的基本信息和投稿列表，注册用户流程简化。
- 支持所有用户刷抖音，按投稿时间倒叙推出，单次最多30个视频
  - 限制单次显示视频数量为30，控制顺序为倒叙  
  (需要按时间顺序从后往前排序，然后选取前30个，不足30选取全部，由于参数中并没有携带视频发布时间，目前有两种思路：1.直接通过读取本地视频文件创建时间，理论上与发布时间等价，按照这个排序。2.在后端视频的属性中增加一条发布时间，发布时将视频信息写入数据库，但是需要先实现登录注册功能，因为必须依赖发布。) 理论上刷视频功能不应该首先实现，因为此时没有登录注册没有视频数据。
- 注册登录功能  
  - 
  - 
  
- 实现登录用户投稿功能
- 查看自己基本信息和投稿列表

根据分析，第一步应该设计数据库的表结构：
首先实现基础功能：
应包含用户表、视频表
服务器需要保存用户登录状态，服务器启动时初始化usersLoginInfo？ nonono

**第一部分已全部完成**
存在问题： 1.只有重启程序才能刷新首页，没有刷新功能 2.封面图没有上传


1. 点赞列表，用户评论
- 对视频点赞
- 对视频评论
- 个人主页有视频点赞列表

先完成无数据库版本，添加数据库仅仅是数据的保存。

3. 关注列表，粉丝列表
- 关注功能
- 关注数和粉丝数加入用户信息中
- 能够显示关注列表和粉丝列表



### Go语言数据库操作

